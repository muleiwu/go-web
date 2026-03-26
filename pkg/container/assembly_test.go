package container

import (
	"reflect"
	"strings"
	"testing"

	"cnb.cool/mliev/open/go-web/pkg/interfaces"
)

// 测试用 mock 类型
type mockEnv struct{}
type mockConfig struct{}
type mockLogger struct{}
type mockDatabase struct{}
type mockRedis struct{}
type mockCache struct{}
type mockSvcA struct{}
type mockSvcB struct{}
type mockMissing struct{}

type mockAssembly struct {
	typ  reflect.Type
	deps []reflect.Type
	val  any
}

func (m *mockAssembly) Type() reflect.Type        { return m.typ }
func (m *mockAssembly) DependsOn() []reflect.Type { return m.deps }
func (m *mockAssembly) Assembly() (any, error)    { return m.val, nil }

func assemblyTypes(assemblies []interfaces.AssemblyInterface) []string {
	result := make([]string, len(assemblies))
	for i, a := range assemblies {
		result[i] = a.Type().String()
	}
	return result
}

func assemblyIndexOfType(order []interfaces.AssemblyInterface, typ reflect.Type) int {
	for i, a := range order {
		if a.Type() == typ {
			return i
		}
	}
	return -1
}

func TestSortByDependency_FrameworkOrder(t *testing.T) {
	tEnv := reflect.TypeFor[mockEnv]()
	tConfig := reflect.TypeFor[mockConfig]()
	tLogger := reflect.TypeFor[mockLogger]()
	tDatabase := reflect.TypeFor[mockDatabase]()
	tRedis := reflect.TypeFor[mockRedis]()
	tCache := reflect.TypeFor[mockCache]()

	assemblies := []interfaces.AssemblyInterface{
		&mockAssembly{typ: tCache, deps: []reflect.Type{tConfig, tLogger, tRedis}, val: "cache"},
		&mockAssembly{typ: tRedis, deps: []reflect.Type{tConfig}, val: "redis"},
		&mockAssembly{typ: tDatabase, deps: []reflect.Type{tConfig}, val: "database"},
		&mockAssembly{typ: tLogger, deps: []reflect.Type{tConfig}, val: "logger"},
		&mockAssembly{typ: tConfig, deps: []reflect.Type{tEnv}, val: "config"},
		&mockAssembly{typ: tEnv, deps: nil, val: "env"},
	}

	sorted, err := SortByDependency(assemblies)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(sorted) != 6 {
		t.Fatalf("expected 6, got %d", len(sorted))
	}

	if assemblyIndexOfType(sorted, tEnv) != 0 {
		t.Errorf("env should be first, got order: %v", assemblyTypes(sorted))
	}
	if assemblyIndexOfType(sorted, tConfig) < assemblyIndexOfType(sorted, tEnv) {
		t.Errorf("config should come after env")
	}
	for _, svc := range []reflect.Type{tLogger, tDatabase, tRedis} {
		if assemblyIndexOfType(sorted, svc) < assemblyIndexOfType(sorted, tConfig) {
			t.Errorf("%v should come after config", svc)
		}
	}
	if assemblyIndexOfType(sorted, tCache) < assemblyIndexOfType(sorted, tLogger) ||
		assemblyIndexOfType(sorted, tCache) < assemblyIndexOfType(sorted, tRedis) {
		t.Errorf("cache should come after logger and redis")
	}
}

func TestSortByDependency_CircularDependency(t *testing.T) {
	tA := reflect.TypeFor[mockSvcA]()
	tB := reflect.TypeFor[mockSvcB]()

	assemblies := []interfaces.AssemblyInterface{
		&mockAssembly{typ: tA, deps: []reflect.Type{tB}, val: "a"},
		&mockAssembly{typ: tB, deps: []reflect.Type{tA}, val: "b"},
	}

	_, err := SortByDependency(assemblies)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "circular dependency detected") {
		t.Errorf("expected circular dependency error, got: %v", err)
	}
}

func TestSortByDependency_MissingDependency(t *testing.T) {
	tA := reflect.TypeFor[mockSvcA]()
	tMissing := reflect.TypeFor[mockMissing]()

	assemblies := []interfaces.AssemblyInterface{
		&mockAssembly{typ: tA, deps: []reflect.Type{tMissing}, val: "a"},
	}

	_, err := SortByDependency(assemblies)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "not registered") {
		t.Errorf("expected 'not registered' error, got: %v", err)
	}
}

func TestSortByDependency_DuplicateType(t *testing.T) {
	tA := reflect.TypeFor[mockSvcA]()

	assemblies := []interfaces.AssemblyInterface{
		&mockAssembly{typ: tA, deps: nil, val: "a1"},
		&mockAssembly{typ: tA, deps: nil, val: "a2"},
	}

	_, err := SortByDependency(assemblies)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "duplicate") {
		t.Errorf("expected duplicate error, got: %v", err)
	}
}

func TestRegisterAssemblies(t *testing.T) {
	tA := reflect.TypeFor[mockSvcA]()
	tB := reflect.TypeFor[mockSvcB]()

	assemblies := []interfaces.AssemblyInterface{
		&mockAssembly{typ: tB, deps: []reflect.Type{tA}, val: "b"},
		&mockAssembly{typ: tA, deps: nil, val: "a"},
	}

	err := RegisterAssemblies(assemblies)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// 验证注册到全局容器
	valA, err := Get[string]()
	// 注意：注册的 key 是 mockSvcA 类型，但 Get[string]() 用的 key 是 string 类型
	// 在新系统中需要直接检查 providers
	_ = valA
	_ = err

	global.mu.RLock()
	_, okA := global.providers[tA]
	_, okB := global.providers[tB]
	global.mu.RUnlock()

	if !okA {
		t.Error("expected provider for mockSvcA")
	}
	if !okB {
		t.Error("expected provider for mockSvcB")
	}

	// 清理
	ResetAll()
}
