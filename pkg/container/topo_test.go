package container

import (
	"reflect"
	"strings"
	"testing"
)

// 测试用哨兵类型
type typeA struct{}
type typeB struct{}
type typeC struct{}
type typeD struct{}
type typeX struct{}
type typeY struct{}
type typeZ struct{}

// 模拟框架实际的服务类型
type typeEnv struct{}
type typeConfig struct{}
type typeLogger struct{}
type typeDatabase struct{}
type typeRedis struct{}
type typeCache struct{}

// testProvider 用于测试的 Provider 实现
type testProvider struct {
	typ  reflect.Type
	deps []reflect.Type
}

func (p *testProvider) Type() reflect.Type        { return p.typ }
func (p *testProvider) Build() any                { return nil }
func (p *testProvider) Priority() int             { return 0 }
func (p *testProvider) DependsOn() []reflect.Type { return p.deps }

func makeEntries(providers ...*testProvider) map[reflect.Type]*entry {
	m := make(map[reflect.Type]*entry)
	for _, p := range providers {
		m[p.typ] = &entry{provider: p}
	}
	return m
}

// indexOfType 返回 typ 在 order 中的位置
func indexOfType(order []reflect.Type, typ reflect.Type) int {
	for i, t := range order {
		if t == typ {
			return i
		}
	}
	return -1
}

func TestTopoSort_LinearChain(t *testing.T) {
	tA := reflect.TypeFor[typeA]()
	tB := reflect.TypeFor[typeB]()
	tC := reflect.TypeFor[typeC]()

	// C -> B -> A（A 无依赖，B 依赖 A，C 依赖 B）
	entries := makeEntries(
		&testProvider{typ: tA, deps: nil},
		&testProvider{typ: tB, deps: []reflect.Type{tA}},
		&testProvider{typ: tC, deps: []reflect.Type{tB}},
	)

	order, err := topoSort(entries)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(order) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(order))
	}

	// A 必须在 B 前面，B 必须在 C 前面
	if indexOfType(order, tA) > indexOfType(order, tB) {
		t.Errorf("A should come before B, got %v", order)
	}
	if indexOfType(order, tB) > indexOfType(order, tC) {
		t.Errorf("B should come before C, got %v", order)
	}
}

func TestTopoSort_Diamond(t *testing.T) {
	tA := reflect.TypeFor[typeA]()
	tB := reflect.TypeFor[typeB]()
	tC := reflect.TypeFor[typeC]()
	tD := reflect.TypeFor[typeD]()

	// D 无依赖，B 和 C 依赖 D，A 依赖 B 和 C
	entries := makeEntries(
		&testProvider{typ: tD, deps: nil},
		&testProvider{typ: tB, deps: []reflect.Type{tD}},
		&testProvider{typ: tC, deps: []reflect.Type{tD}},
		&testProvider{typ: tA, deps: []reflect.Type{tB, tC}},
	)

	order, err := topoSort(entries)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(order) != 4 {
		t.Fatalf("expected 4 entries, got %d", len(order))
	}

	// D 在 B 和 C 前面，B 和 C 在 A 前面
	if indexOfType(order, tD) > indexOfType(order, tB) {
		t.Errorf("D should come before B, got %v", order)
	}
	if indexOfType(order, tD) > indexOfType(order, tC) {
		t.Errorf("D should come before C, got %v", order)
	}
	if indexOfType(order, tB) > indexOfType(order, tA) {
		t.Errorf("B should come before A, got %v", order)
	}
	if indexOfType(order, tC) > indexOfType(order, tA) {
		t.Errorf("C should come before A, got %v", order)
	}
}

func TestTopoSort_NoDependencies(t *testing.T) {
	entries := makeEntries(
		&testProvider{typ: reflect.TypeFor[typeX](), deps: nil},
		&testProvider{typ: reflect.TypeFor[typeY](), deps: nil},
		&testProvider{typ: reflect.TypeFor[typeZ](), deps: nil},
	)

	order, err := topoSort(entries)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(order) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(order))
	}
}

func TestTopoSort_SimpleCycle(t *testing.T) {
	tA := reflect.TypeFor[typeA]()
	tB := reflect.TypeFor[typeB]()

	entries := makeEntries(
		&testProvider{typ: tA, deps: []reflect.Type{tB}},
		&testProvider{typ: tB, deps: []reflect.Type{tA}},
	)

	_, err := topoSort(entries)
	if err == nil {
		t.Fatal("expected cycle error, got nil")
	}

	if !strings.Contains(err.Error(), "circular dependency detected") {
		t.Errorf("expected circular dependency error, got: %v", err)
	}
}

func TestTopoSort_ComplexCycle(t *testing.T) {
	tA := reflect.TypeFor[typeA]()
	tB := reflect.TypeFor[typeB]()
	tC := reflect.TypeFor[typeC]()

	entries := makeEntries(
		&testProvider{typ: tA, deps: []reflect.Type{tB}},
		&testProvider{typ: tB, deps: []reflect.Type{tC}},
		&testProvider{typ: tC, deps: []reflect.Type{tA}},
	)

	_, err := topoSort(entries)
	if err == nil {
		t.Fatal("expected cycle error, got nil")
	}

	if !strings.Contains(err.Error(), "circular dependency detected") {
		t.Errorf("expected circular dependency error, got: %v", err)
	}
}

func TestTopoSort_MissingDependency(t *testing.T) {
	tA := reflect.TypeFor[typeA]()
	tX := reflect.TypeFor[typeX]()

	entries := makeEntries(
		&testProvider{typ: tA, deps: []reflect.Type{tX}},
	)

	_, err := topoSort(entries)
	if err == nil {
		t.Fatal("expected missing dependency error, got nil")
	}

	if !strings.Contains(err.Error(), "not registered") {
		t.Errorf("expected 'not registered' error, got: %v", err)
	}
}

func TestTopoSort_Mixed(t *testing.T) {
	tA := reflect.TypeFor[typeA]()
	tB := reflect.TypeFor[typeB]()
	tC := reflect.TypeFor[typeC]()

	entries := makeEntries(
		&testProvider{typ: tA, deps: nil},
		&testProvider{typ: tB, deps: []reflect.Type{tA}},
	)
	// 添加一个不实现 DependencyAware 的 provider（使用 SimpleProvider，deps 为 nil）
	entries[tC] = &entry{provider: &SimpleProvider{typ: tC}}

	order, err := topoSort(entries)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(order) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(order))
	}

	if indexOfType(order, tA) > indexOfType(order, tB) {
		t.Errorf("A should come before B, got %v", order)
	}
}

func TestTopoSort_FrameworkAssemblyOrder(t *testing.T) {
	tEnv := reflect.TypeFor[typeEnv]()
	tConfig := reflect.TypeFor[typeConfig]()
	tLogger := reflect.TypeFor[typeLogger]()
	tDatabase := reflect.TypeFor[typeDatabase]()
	tRedis := reflect.TypeFor[typeRedis]()
	tCache := reflect.TypeFor[typeCache]()

	// 模拟框架实际的依赖关系
	entries := makeEntries(
		&testProvider{typ: tEnv, deps: nil},
		&testProvider{typ: tConfig, deps: []reflect.Type{tEnv}},
		&testProvider{typ: tLogger, deps: []reflect.Type{tConfig}},
		&testProvider{typ: tDatabase, deps: []reflect.Type{tConfig}},
		&testProvider{typ: tRedis, deps: []reflect.Type{tConfig}},
		&testProvider{typ: tCache, deps: []reflect.Type{tConfig, tLogger, tRedis}},
	)

	order, err := topoSort(entries)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(order) != 6 {
		t.Fatalf("expected 6 entries, got %d", len(order))
	}

	// env 最先
	if indexOfType(order, tEnv) != 0 {
		t.Errorf("env should be first, got %v", order)
	}
	// config 在 env 之后
	if indexOfType(order, tConfig) < indexOfType(order, tEnv) {
		t.Errorf("config should come after env, got %v", order)
	}
	// cache 在 config、logger、redis 之后
	if indexOfType(order, tCache) < indexOfType(order, tConfig) ||
		indexOfType(order, tCache) < indexOfType(order, tLogger) ||
		indexOfType(order, tCache) < indexOfType(order, tRedis) {
		t.Errorf("cache should come after config, logger, and redis, got %v", order)
	}
}
