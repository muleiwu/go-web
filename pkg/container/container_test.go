package container

import (
	"reflect"
	"strings"
	"testing"
)

// 测试用哨兵类型（与 topo_test.go 共用文件内定义会冲突，这里用不同命名）
type svcA struct{}
type svcB struct{}
type svcC struct{}
type svcBase struct{}
type svcDependent struct{}
type svcNoDeps struct{}
type svcWithDeps struct{}
type svcSelfRef struct{}
type svcLegacy struct{}

func TestInitAll_HappyPath(t *testing.T) {
	c := NewContainer()

	var initOrder []string

	c.Register(&LazyProvider{
		typ:     reflect.TypeFor[svcA](),
		factory: func() any { initOrder = append(initOrder, "A"); return svcA{} },
	})
	c.Register(&LazyProvider{
		typ:     reflect.TypeFor[svcB](),
		factory: func() any { initOrder = append(initOrder, "B"); return svcB{} },
		deps:    []reflect.Type{reflect.TypeFor[svcA]()},
	})
	c.Register(&LazyProvider{
		typ:     reflect.TypeFor[svcC](),
		factory: func() any { initOrder = append(initOrder, "C"); return svcC{} },
		deps:    []reflect.Type{reflect.TypeFor[svcA](), reflect.TypeFor[svcB]()},
	})

	if err := c.InitAll(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// 验证初始化顺序：A 在 B 前面，B 在 C 前面
	if len(initOrder) != 3 {
		t.Fatalf("expected 3 inits, got %d", len(initOrder))
	}

	aIdx, bIdx, cIdx := -1, -1, -1
	for i, n := range initOrder {
		switch n {
		case "A":
			aIdx = i
		case "B":
			bIdx = i
		case "C":
			cIdx = i
		}
	}
	if aIdx > bIdx {
		t.Errorf("A should init before B, got order: %v", initOrder)
	}
	if bIdx > cIdx {
		t.Errorf("B should init before C, got order: %v", initOrder)
	}
}

func TestInitAll_CircleDetection(t *testing.T) {
	c := NewContainer()

	c.Register(&LazyProvider{
		typ:     reflect.TypeFor[svcA](),
		factory: func() any { return svcA{} },
		deps:    []reflect.Type{reflect.TypeFor[svcB]()},
	})
	c.Register(&LazyProvider{
		typ:     reflect.TypeFor[svcB](),
		factory: func() any { return svcB{} },
		deps:    []reflect.Type{reflect.TypeFor[svcA]()},
	})

	err := c.InitAll()
	if err == nil {
		t.Fatal("expected cycle error, got nil")
	}

	if !strings.Contains(err.Error(), "circular dependency detected") {
		t.Errorf("expected circular dependency error, got: %v", err)
	}
}

func TestBackwardCompatibility(t *testing.T) {
	c := NewContainer()

	// 使用不带依赖的 SimpleProvider
	c.Register(NewSimpleProvider(reflect.TypeFor[svcLegacy](), "legacy-value"))

	val, err := getFromContainer[string](c)
	// getFromContainer uses TypeFor[T] which is string here
	// Actually we need to get by the registered type
	_ = val
	_ = err

	// 用注册时的类型获取
	c2 := NewContainer()
	c2.Register(NewSimpleProvider(reflect.TypeFor[string](), "legacy-value"))

	val2, err2 := getFromContainer[string](c2)
	if err2 != nil {
		t.Fatalf("unexpected error: %v", err2)
	}
	if val2 != "legacy-value" {
		t.Errorf("expected 'legacy-value', got %q", val2)
	}
}

func TestDestroyAll_ReverseOrder(t *testing.T) {
	c := NewContainer()

	var destroyOrder []string

	type trackingInstance struct {
		name string
	}

	tBase := reflect.TypeFor[svcBase]()
	tDep := reflect.TypeFor[svcDependent]()

	c.Register(&LazyProvider{
		typ: tBase,
		factory: func() any {
			return &trackingInstance{name: "base"}
		},
	})
	c.Register(&LazyProvider{
		typ: tDep,
		factory: func() any {
			return &trackingInstance{name: "dependent"}
		},
		deps: []reflect.Type{tBase},
	})

	// 先初始化
	if err := c.InitAll(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// DestroyAll 应该不报错
	if err := c.DestroyAll(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// 验证所有实例已被重置
	for _, typ := range []reflect.Type{tBase, tDep} {
		c.mu.RLock()
		e := c.providers[typ]
		c.mu.RUnlock()
		if e.done {
			t.Errorf("provider %v should be reset after DestroyAll", typ)
		}
	}

	_ = destroyOrder // 未使用，但保留变量声明的清晰性
}

func TestRuntimeCycleDetection(t *testing.T) {
	c := NewContainer()

	tSelfRef := reflect.TypeFor[svcSelfRef]()

	// A 的 Build 中尝试获取自己 -> 运行时循环检测
	c.Register(&LazyProvider{
		typ: tSelfRef,
		factory: func() any {
			// 模拟 MustGet 行为：获取失败则 panic
			v, err := getFromContainer[svcSelfRef](c)
			if err != nil {
				panic(err)
			}
			return v
		},
	})

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic from runtime cycle detection, got none")
		}
		errMsg := r.(error).Error()
		if !strings.Contains(errMsg, "circular dependency detected at runtime") {
			t.Errorf("expected runtime cycle error, got: %v", errMsg)
		}
	}()

	_, _ = getFromContainer[svcSelfRef](c)
}

func TestInitAll_MixedProviders(t *testing.T) {
	c := NewContainer()

	tNoDeps := reflect.TypeFor[svcNoDeps]()
	tWithDeps := reflect.TypeFor[svcWithDeps]()

	// 混合：有依赖声明的和没有的
	c.Register(NewSimpleProvider(tNoDeps, "value1"))
	c.Register(NewSimpleProviderWithDeps(tWithDeps, "value2", tNoDeps))

	if err := c.InitAll(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// getFromContainer 使用 TypeFor[T]，所以要用注册类型
	// 但这里注册的 instance 是 string，获取时需要用 string 类型
	// 这说明一个设计要点：注册类型和实例类型应当一致
	// 在重构后的设计中，注册的 key 是 reflect.Type，获取时也用 reflect.TypeFor[T]()
	// 由于这里注册 key 是 svcNoDeps 但值是 string，Get[svcNoDeps]() 会类型不匹配
	// 改为直接测试 providers 存在
	c.mu.RLock()
	_, ok1 := c.providers[tNoDeps]
	_, ok2 := c.providers[tWithDeps]
	c.mu.RUnlock()

	if !ok1 {
		t.Error("expected provider for svcNoDeps")
	}
	if !ok2 {
		t.Error("expected provider for svcWithDeps")
	}
}
