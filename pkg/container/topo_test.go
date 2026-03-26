package container

import (
	"strings"
	"testing"
)

// testProvider 用于测试的 Provider 实现
type testProvider struct {
	name string
	deps []string
}

func (p *testProvider) Name() string        { return p.name }
func (p *testProvider) Build() any          { return p.name }
func (p *testProvider) Priority() int       { return 0 }
func (p *testProvider) DependsOn() []string { return p.deps }

func makeEntries(providers ...*testProvider) map[string]*entry {
	m := make(map[string]*entry)
	for _, p := range providers {
		m[p.name] = &entry{provider: p}
	}
	return m
}

// indexOf 返回 name 在 order 中的位置
func indexOf(order []string, name string) int {
	for i, n := range order {
		if n == name {
			return i
		}
	}
	return -1
}

func TestTopoSort_LinearChain(t *testing.T) {
	// C -> B -> A（A 无依赖，B 依赖 A，C 依赖 B）
	entries := makeEntries(
		&testProvider{name: "A", deps: nil},
		&testProvider{name: "B", deps: []string{"A"}},
		&testProvider{name: "C", deps: []string{"B"}},
	)

	order, err := topoSort(entries)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(order) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(order))
	}

	// A 必须在 B 前面，B 必须在 C 前面
	if indexOf(order, "A") > indexOf(order, "B") {
		t.Errorf("A should come before B, got %v", order)
	}
	if indexOf(order, "B") > indexOf(order, "C") {
		t.Errorf("B should come before C, got %v", order)
	}
}

func TestTopoSort_Diamond(t *testing.T) {
	// D 无依赖，B 和 C 依赖 D，A 依赖 B 和 C
	entries := makeEntries(
		&testProvider{name: "D", deps: nil},
		&testProvider{name: "B", deps: []string{"D"}},
		&testProvider{name: "C", deps: []string{"D"}},
		&testProvider{name: "A", deps: []string{"B", "C"}},
	)

	order, err := topoSort(entries)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(order) != 4 {
		t.Fatalf("expected 4 entries, got %d", len(order))
	}

	// D 在 B 和 C 前面，B 和 C 在 A 前面
	if indexOf(order, "D") > indexOf(order, "B") {
		t.Errorf("D should come before B, got %v", order)
	}
	if indexOf(order, "D") > indexOf(order, "C") {
		t.Errorf("D should come before C, got %v", order)
	}
	if indexOf(order, "B") > indexOf(order, "A") {
		t.Errorf("B should come before A, got %v", order)
	}
	if indexOf(order, "C") > indexOf(order, "A") {
		t.Errorf("C should come before A, got %v", order)
	}
}

func TestTopoSort_NoDependencies(t *testing.T) {
	entries := makeEntries(
		&testProvider{name: "X", deps: nil},
		&testProvider{name: "Y", deps: nil},
		&testProvider{name: "Z", deps: nil},
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
	entries := makeEntries(
		&testProvider{name: "A", deps: []string{"B"}},
		&testProvider{name: "B", deps: []string{"A"}},
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
	entries := makeEntries(
		&testProvider{name: "A", deps: []string{"B"}},
		&testProvider{name: "B", deps: []string{"C"}},
		&testProvider{name: "C", deps: []string{"A"}},
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
	entries := makeEntries(
		&testProvider{name: "A", deps: []string{"X"}},
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
	// noDepsProvider 不实现 DependencyAware
	type noDepsProvider struct {
		name string
	}

	entries := makeEntries(
		&testProvider{name: "A", deps: nil},
		&testProvider{name: "B", deps: []string{"A"}},
	)
	// 添加一个不实现 DependencyAware 的 provider（使用 SimpleProvider，deps 为 nil）
	entries["C"] = &entry{provider: &SimpleProvider{name: "C"}}

	order, err := topoSort(entries)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(order) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(order))
	}

	if indexOf(order, "A") > indexOf(order, "B") {
		t.Errorf("A should come before B, got %v", order)
	}
}

func TestTopoSort_FrameworkAssemblyOrder(t *testing.T) {
	// 模拟框架实际的依赖关系
	entries := makeEntries(
		&testProvider{name: "env", deps: nil},
		&testProvider{name: "config", deps: []string{"env"}},
		&testProvider{name: "logger", deps: []string{"config"}},
		&testProvider{name: "database", deps: []string{"config"}},
		&testProvider{name: "redis", deps: []string{"config"}},
		&testProvider{name: "cache", deps: []string{"config", "logger", "redis"}},
	)

	order, err := topoSort(entries)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(order) != 6 {
		t.Fatalf("expected 6 entries, got %d", len(order))
	}

	// env 最先
	if indexOf(order, "env") != 0 {
		t.Errorf("env should be first, got %v", order)
	}
	// config 在 env 之后
	if indexOf(order, "config") < indexOf(order, "env") {
		t.Errorf("config should come after env, got %v", order)
	}
	// cache 在 config、logger、redis 之后
	if indexOf(order, "cache") < indexOf(order, "config") ||
		indexOf(order, "cache") < indexOf(order, "logger") ||
		indexOf(order, "cache") < indexOf(order, "redis") {
		t.Errorf("cache should come after config, logger, and redis, got %v", order)
	}
}
