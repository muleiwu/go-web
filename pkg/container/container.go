package container

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"sync"
)

var global = NewContainer()

type entry struct {
	mu       sync.Mutex
	done     bool
	provider Provider
	instance any
}

// resolve 执行懒加载初始化，失败时允许重试
func (e *entry) resolve(name string) (any, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.done {
		return e.instance, nil
	}

	inst := e.provider.Build()
	if init, ok := inst.(Initializable); ok {
		if err := init.Init(); err != nil {
			return nil, fmt.Errorf("init %s failed: %w", name, err)
		}
	}

	e.instance = inst
	e.done = true
	fmt.Printf("[Container] Instantiated: %s\n", name)
	return e.instance, nil
}

type Container struct {
	mu        sync.RWMutex
	providers map[string]*entry       // name -> entry
	aliases   map[reflect.Type]string // 接口类型 -> name
	resolving sync.Map                // goroutine ID -> map[string]bool，运行时循环检测
}

func NewContainer() *Container {
	return &Container{
		providers: make(map[string]*entry),
		aliases:   make(map[reflect.Type]string),
	}
}

// Register 注册一个Provider（通常在 init() 中调用）
func Register(p Provider) {
	global.Register(p)
}

func (c *Container) Register(p Provider) {
	c.mu.Lock()
	defer c.mu.Unlock()

	name := p.Name()
	if existing, ok := c.providers[name]; ok {
		// 优先级高的覆盖低的
		if p.Priority() <= existing.provider.Priority() {
			return
		}
	}
	c.providers[name] = &entry{provider: p}
	fmt.Printf("[Container] Registered provider: %s (priority=%d)\n", name, p.Priority())
}

// BindInterface 将接口类型绑定到具体服务名
// 例如: BindInterface((*cache.Cache)(nil), "redis-cache")
func BindInterface(iface any, name string) {
	global.BindInterface(iface, name)
}

func (c *Container) BindInterface(iface any, name string) {
	t := reflect.TypeOf(iface)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	c.mu.Lock()
	c.aliases[t] = name
	c.mu.Unlock()
}

// Get 按名称获取服务实例（懒加载单例）
func Get[T any](name string) (T, error) {
	return getFromContainer[T](global, name)
}

func getFromContainer[T any](c *Container, name string) (T, error) {
	var zero T
	c.mu.RLock()
	e, ok := c.providers[name]
	c.mu.RUnlock()

	if !ok {
		return zero, fmt.Errorf("provider %q not found", name)
	}

	// 运行时循环依赖检测
	gid := goroutineID()
	val, _ := c.resolving.LoadOrStore(gid, make(map[string]bool))
	resolvingSet := val.(map[string]bool)
	if resolvingSet[name] {
		return zero, fmt.Errorf("circular dependency detected at runtime: %q is already being resolved in this call chain", name)
	}
	resolvingSet[name] = true
	defer func() {
		delete(resolvingSet, name)
		if len(resolvingSet) == 0 {
			c.resolving.Delete(gid)
		}
	}()

	// 懒加载，失败可重试
	inst, err := e.resolve(name)
	if err != nil {
		return zero, err
	}

	result, ok := inst.(T)
	if !ok {
		return zero, fmt.Errorf("provider %q type mismatch", name)
	}
	return result, nil
}

// MustGet 获取服务，失败则panic（适合确定存在的场景）
func MustGet[T any](name string) T {
	v, err := Get[T](name)

	if err != nil {
		panic(err)
	}
	return v
}

// Inject 通过反射自动注入结构体中带 `inject:"name"` 标签的字段
func Inject(target any) error {
	return global.Inject(target)
}

func (c *Container) Inject(target any) error {
	val := reflect.ValueOf(target)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to struct")
	}

	val = val.Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("inject")
		if tag == "" {
			continue
		}

		c.mu.RLock()
		e, ok := c.providers[tag]
		c.mu.RUnlock()

		if !ok {
			return fmt.Errorf("inject field %s: provider %q not found", field.Name, tag)
		}

		// 触发懒加载
		inst, err := e.resolve(tag)
		if err != nil {
			return fmt.Errorf("inject field %s: %w", field.Name, err)
		}

		fv := val.Field(i)
		if !fv.CanSet() {
			return fmt.Errorf("field %s is not settable", field.Name)
		}

		iv := reflect.ValueOf(inst)
		if !iv.Type().AssignableTo(fv.Type()) {
			return fmt.Errorf("field %s: type %v not assignable to %v", field.Name, iv.Type(), fv.Type())
		}
		fv.Set(iv)
	}
	return nil
}

// Reset 重置指定名称的服务实例，使其下次 Get 时重新 Build
// 用于 SIGHUP 重载场景
func Reset(name string) {
	global.Reset(name)
}

func (c *Container) Reset(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if e, ok := c.providers[name]; ok {
		// 如果旧实例支持 Destroy，先调用
		if d, ok := e.instance.(Destroyable); ok {
			_ = d.Destroy()
		}
		// 重置 entry，保留 provider 但清除实例和 once
		c.providers[name] = &entry{provider: e.provider}
	}
}

// ResetAll 重置所有服务实例
func ResetAll() {
	global.ResetAll()
}

func (c *Container) ResetAll() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for name, e := range c.providers {
		if d, ok := e.instance.(Destroyable); ok {
			_ = d.Destroy()
		}
		c.providers[name] = &entry{provider: e.provider}
	}
}

// ListProviders 查看所有已注册的Provider
func ListProviders() []string {
	global.mu.RLock()
	defer global.mu.RUnlock()

	names := make([]string, 0, len(global.providers))
	for name := range global.providers {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// InitAll 按拓扑序初始化所有已注册的 Provider。
// 在所有 Provider 注册完成后调用，可检测循环依赖并按正确顺序初始化。
func InitAll() error {
	return global.InitAll()
}

func (c *Container) InitAll() error {
	c.mu.RLock()
	sorted, err := topoSort(c.providers)
	c.mu.RUnlock()
	if err != nil {
		return err
	}

	for _, name := range sorted {
		c.mu.RLock()
		e := c.providers[name]
		c.mu.RUnlock()
		if _, err := e.resolve(name); err != nil {
			return fmt.Errorf("InitAll: failed to resolve %q: %w", name, err)
		}
	}
	return nil
}

// DestroyAll 按反向拓扑序销毁所有已解析的实例（先销毁依赖者，再销毁被依赖者）。
func DestroyAll() error {
	return global.DestroyAll()
}

func (c *Container) DestroyAll() error {
	c.mu.RLock()
	sorted, err := topoSort(c.providers)
	c.mu.RUnlock()
	if err != nil {
		// 拓扑排序失败时回退到无序销毁
		c.ResetAll()
		return err
	}

	// 反向销毁：先销毁依赖者，再销毁被依赖者
	for i := len(sorted) - 1; i >= 0; i-- {
		c.Reset(sorted[i])
	}
	return nil
}

// goroutineID 获取当前 goroutine 的 ID，用于运行时循环检测
func goroutineID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// "goroutine 123 [..."
	field := bytes.Fields(buf[:n])[1]
	id, _ := strconv.ParseInt(string(field), 10, 64)
	return id
}
