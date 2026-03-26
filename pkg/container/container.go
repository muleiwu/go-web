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
func (e *entry) resolve(typ reflect.Type) (any, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.done {
		return e.instance, nil
	}

	inst := e.provider.Build()
	if init, ok := inst.(Initializable); ok {
		if err := init.Init(); err != nil {
			return nil, fmt.Errorf("init %v failed: %w", typ, err)
		}
	}

	e.instance = inst
	e.done = true
	fmt.Printf("[Container] Instantiated: %v\n", typ)
	return e.instance, nil
}

type Container struct {
	mu        sync.RWMutex
	providers map[reflect.Type]*entry // type -> entry
	resolving sync.Map                // goroutine ID -> map[reflect.Type]bool，运行时循环检测
}

func NewContainer() *Container {
	return &Container{
		providers: make(map[reflect.Type]*entry),
	}
}

// Register 注册一个Provider（通常在 init() 中调用）
func Register(p Provider) {
	global.Register(p)
}

func (c *Container) Register(p Provider) {
	c.mu.Lock()
	defer c.mu.Unlock()

	typ := p.Type()
	if existing, ok := c.providers[typ]; ok {
		// 优先级高的覆盖低的
		if p.Priority() <= existing.provider.Priority() {
			return
		}
	}
	c.providers[typ] = &entry{provider: p}
	fmt.Printf("[Container] Registered provider: %v (priority=%d)\n", typ, p.Priority())
}

// Get 按类型获取服务实例（懒加载单例）
func Get[T any]() (T, error) {
	return getFromContainer[T](global)
}

func getFromContainer[T any](c *Container) (T, error) {
	var zero T
	typ := reflect.TypeFor[T]()

	c.mu.RLock()
	e, ok := c.providers[typ]
	c.mu.RUnlock()

	if !ok {
		return zero, fmt.Errorf("provider %v not found", typ)
	}

	// 运行时循环依赖检测
	gid := goroutineID()
	val, _ := c.resolving.LoadOrStore(gid, make(map[reflect.Type]bool))
	resolvingSet := val.(map[reflect.Type]bool)
	if resolvingSet[typ] {
		return zero, fmt.Errorf("circular dependency detected at runtime: %v is already being resolved in this call chain", typ)
	}
	resolvingSet[typ] = true
	defer func() {
		delete(resolvingSet, typ)
		if len(resolvingSet) == 0 {
			c.resolving.Delete(gid)
		}
	}()

	// 懒加载，失败可重试
	inst, err := e.resolve(typ)
	if err != nil {
		return zero, err
	}

	result, ok := inst.(T)
	if !ok {
		return zero, fmt.Errorf("provider %v type mismatch", typ)
	}
	return result, nil
}

// MustGet 获取服务，失败则panic（适合确定存在的场景）
func MustGet[T any]() T {
	v, err := Get[T]()

	if err != nil {
		panic(err)
	}
	return v
}

// Inject 通过反射自动注入结构体中带 `inject:""` 标签的字段（按字段类型查找容器）
// inject:"-" 表示跳过该字段
func Inject(target any) error {
	return global.Inject(target)
}

func (c *Container) Inject(target any) error {
	val := reflect.ValueOf(target)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to struct")
	}

	val = val.Elem()
	structType := val.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag, hasTag := field.Tag.Lookup("inject")
		if !hasTag || tag == "-" {
			continue
		}

		fieldType := field.Type

		c.mu.RLock()
		e, ok := c.providers[fieldType]
		c.mu.RUnlock()

		if !ok {
			return fmt.Errorf("inject field %s: provider %v not found", field.Name, fieldType)
		}

		// 触发懒加载
		inst, err := e.resolve(fieldType)
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

// Reset 重置指定类型的服务实例，使其下次 Get 时重新 Build
// 用于 SIGHUP 重载场景
func Reset(typ reflect.Type) {
	global.Reset(typ)
}

func (c *Container) Reset(typ reflect.Type) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if e, ok := c.providers[typ]; ok {
		// 如果旧实例支持 Destroy，先调用
		if d, ok := e.instance.(Destroyable); ok {
			_ = d.Destroy()
		}
		// 重置 entry，保留 provider 但清除实例和 once
		c.providers[typ] = &entry{provider: e.provider}
	}
}

// ResetAll 重置所有服务实例
func ResetAll() {
	global.ResetAll()
}

func (c *Container) ResetAll() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for typ, e := range c.providers {
		if d, ok := e.instance.(Destroyable); ok {
			_ = d.Destroy()
		}
		c.providers[typ] = &entry{provider: e.provider}
	}
}

// ListProviders 查看所有已注册的Provider
func ListProviders() []string {
	global.mu.RLock()
	defer global.mu.RUnlock()

	names := make([]string, 0, len(global.providers))
	for typ := range global.providers {
		names = append(names, typ.String())
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

	for _, typ := range sorted {
		c.mu.RLock()
		e := c.providers[typ]
		c.mu.RUnlock()
		if _, err := e.resolve(typ); err != nil {
			return fmt.Errorf("InitAll: failed to resolve %v: %w", typ, err)
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
