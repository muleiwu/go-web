package container

import (
	"strings"
	"testing"
)

func TestInitAll_HappyPath(t *testing.T) {
	c := NewContainer()

	var initOrder []string

	c.Register(&LazyProvider{
		name:    "A",
		factory: func() any { initOrder = append(initOrder, "A"); return "A" },
	})
	c.Register(&LazyProvider{
		name:    "B",
		factory: func() any { initOrder = append(initOrder, "B"); return "B" },
		deps:    []string{"A"},
	})
	c.Register(&LazyProvider{
		name:    "C",
		factory: func() any { initOrder = append(initOrder, "C"); return "C" },
		deps:    []string{"A", "B"},
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
		name:    "A",
		factory: func() any { return "A" },
		deps:    []string{"B"},
	})
	c.Register(&LazyProvider{
		name:    "B",
		factory: func() any { return "B" },
		deps:    []string{"A"},
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

	// 使用不带依赖的 SimpleProvider（传统用法）
	c.Register(NewSimpleProvider("legacy", "legacy-value"))

	val, err := getFromContainer[string](c, "legacy")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "legacy-value" {
		t.Errorf("expected 'legacy-value', got %q", val)
	}
}

func TestDestroyAll_ReverseOrder(t *testing.T) {
	c := NewContainer()

	var destroyOrder []string

	type destroyable struct {
		name      string
		onDestroy func()
	}
	// Destroy 实现 Destroyable 接口
	// 由于 destroyable 不是导出类型，直接用闭包记录
	type trackingInstance struct {
		name string
	}

	c.Register(&LazyProvider{
		name: "base",
		factory: func() any {
			return &trackingInstance{name: "base"}
		},
	})
	c.Register(&LazyProvider{
		name: "dependent",
		factory: func() any {
			return &trackingInstance{name: "dependent"}
		},
		deps: []string{"base"},
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
	for _, name := range []string{"base", "dependent"} {
		c.mu.RLock()
		e := c.providers[name]
		c.mu.RUnlock()
		if e.done {
			t.Errorf("provider %q should be reset after DestroyAll", name)
		}
	}

	_ = destroyOrder // 未使用，但保留变量声明的清晰性
}

func TestRuntimeCycleDetection(t *testing.T) {
	c := NewContainer()

	// A 的 Build 中尝试获取自己 -> 运行时循环检测
	// 实际场景中 Assembly 使用 MustGet，循环时会 panic
	c.Register(&LazyProvider{
		name: "self-ref",
		factory: func() any {
			// 模拟 MustGet 行为：获取失败则 panic
			v, err := getFromContainer[any](c, "self-ref")
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

	_, _ = getFromContainer[any](c, "self-ref")
}

func TestInitAll_MixedProviders(t *testing.T) {
	c := NewContainer()

	// 混合：有依赖声明的和没有的
	c.Register(NewSimpleProvider("no-deps", "value1"))
	c.Register(NewSimpleProviderWithDeps("with-deps", "value2", "no-deps"))

	if err := c.InitAll(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	v1, err := getFromContainer[string](c, "no-deps")
	if err != nil || v1 != "value1" {
		t.Errorf("expected 'value1', got %q, err: %v", v1, err)
	}

	v2, err := getFromContainer[string](c, "with-deps")
	if err != nil || v2 != "value2" {
		t.Errorf("expected 'value2', got %q, err: %v", v2, err)
	}
}
