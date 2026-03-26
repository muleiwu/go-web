package container

import (
	"strings"
	"testing"

	"cnb.cool/mliev/open/go-web/pkg/interfaces"
)

type mockAssembly struct {
	name string
	deps []string
}

func (m *mockAssembly) Name() string           { return m.name }
func (m *mockAssembly) DependsOn() []string    { return m.deps }
func (m *mockAssembly) Assembly() (any, error) { return m.name, nil }

func assemblyNames(assemblies []interfaces.AssemblyInterface) []string {
	result := make([]string, len(assemblies))
	for i, a := range assemblies {
		result[i] = a.Name()
	}
	return result
}

func assemblyIndexOf(order []interfaces.AssemblyInterface, name string) int {
	for i, a := range order {
		if a.Name() == name {
			return i
		}
	}
	return -1
}

func TestSortByDependency_FrameworkOrder(t *testing.T) {
	assemblies := []interfaces.AssemblyInterface{
		&mockAssembly{name: "cache", deps: []string{"config", "logger", "redis"}},
		&mockAssembly{name: "redis", deps: []string{"config"}},
		&mockAssembly{name: "database", deps: []string{"config"}},
		&mockAssembly{name: "logger", deps: []string{"config"}},
		&mockAssembly{name: "config", deps: []string{"env"}},
		&mockAssembly{name: "env", deps: nil},
	}

	sorted, err := SortByDependency(assemblies)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(sorted) != 6 {
		t.Fatalf("expected 6, got %d", len(sorted))
	}

	if assemblyIndexOf(sorted, "env") != 0 {
		t.Errorf("env should be first, got order: %v", assemblyNames(sorted))
	}
	if assemblyIndexOf(sorted, "config") < assemblyIndexOf(sorted, "env") {
		t.Errorf("config should come after env")
	}
	for _, svc := range []string{"logger", "database", "redis"} {
		if assemblyIndexOf(sorted, svc) < assemblyIndexOf(sorted, "config") {
			t.Errorf("%s should come after config", svc)
		}
	}
	if assemblyIndexOf(sorted, "cache") < assemblyIndexOf(sorted, "logger") ||
		assemblyIndexOf(sorted, "cache") < assemblyIndexOf(sorted, "redis") {
		t.Errorf("cache should come after logger and redis")
	}
}

func TestSortByDependency_CircularDependency(t *testing.T) {
	assemblies := []interfaces.AssemblyInterface{
		&mockAssembly{name: "A", deps: []string{"B"}},
		&mockAssembly{name: "B", deps: []string{"A"}},
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
	assemblies := []interfaces.AssemblyInterface{
		&mockAssembly{name: "A", deps: []string{"missing"}},
	}

	_, err := SortByDependency(assemblies)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "not registered") {
		t.Errorf("expected 'not registered' error, got: %v", err)
	}
}

func TestSortByDependency_DuplicateName(t *testing.T) {
	assemblies := []interfaces.AssemblyInterface{
		&mockAssembly{name: "A", deps: nil},
		&mockAssembly{name: "A", deps: nil},
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
	// 使用独立容器避免污染全局状态
	assemblies := []interfaces.AssemblyInterface{
		&mockAssembly{name: "b", deps: []string{"a"}},
		&mockAssembly{name: "a", deps: nil},
	}

	err := RegisterAssemblies(assemblies)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// 验证注册到全局容器
	valA, err := Get[string]("a")
	if err != nil {
		t.Fatalf("failed to get 'a': %v", err)
	}
	if valA != "a" {
		t.Errorf("expected 'a', got %q", valA)
	}

	valB, err := Get[string]("b")
	if err != nil {
		t.Fatalf("failed to get 'b': %v", err)
	}
	if valB != "b" {
		t.Errorf("expected 'b', got %q", valB)
	}

	// 清理
	ResetAll()
}
