package interfaces

import (
	"strings"
	"testing"
)

type mockAssembly struct {
	name string
	deps []string
}

func (m *mockAssembly) Name() string           { return m.name }
func (m *mockAssembly) DependsOn() []string    { return m.deps }
func (m *mockAssembly) Assembly() (any, error) { return m.name, nil }

func indexOf(order []AssemblyInterface, name string) int {
	for i, a := range order {
		if a.Name() == name {
			return i
		}
	}
	return -1
}

func TestSortAssemblies_FrameworkOrder(t *testing.T) {
	assemblies := []AssemblyInterface{
		&mockAssembly{name: "cache", deps: []string{"config", "logger", "redis"}},
		&mockAssembly{name: "redis", deps: []string{"config"}},
		&mockAssembly{name: "database", deps: []string{"config"}},
		&mockAssembly{name: "logger", deps: []string{"config"}},
		&mockAssembly{name: "config", deps: []string{"env"}},
		&mockAssembly{name: "env", deps: nil},
	}

	sorted, err := SortAssemblies(assemblies)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(sorted) != 6 {
		t.Fatalf("expected 6, got %d", len(sorted))
	}

	// env 必须最先
	if indexOf(sorted, "env") != 0 {
		t.Errorf("env should be first, got order: %v", names(sorted))
	}
	// config 在 env 之后
	if indexOf(sorted, "config") < indexOf(sorted, "env") {
		t.Errorf("config should come after env")
	}
	// logger, database, redis 在 config 之后
	for _, svc := range []string{"logger", "database", "redis"} {
		if indexOf(sorted, svc) < indexOf(sorted, "config") {
			t.Errorf("%s should come after config", svc)
		}
	}
	// cache 在 logger 和 redis 之后
	if indexOf(sorted, "cache") < indexOf(sorted, "logger") ||
		indexOf(sorted, "cache") < indexOf(sorted, "redis") {
		t.Errorf("cache should come after logger and redis")
	}
}

func TestSortAssemblies_CircularDependency(t *testing.T) {
	assemblies := []AssemblyInterface{
		&mockAssembly{name: "A", deps: []string{"B"}},
		&mockAssembly{name: "B", deps: []string{"A"}},
	}

	_, err := SortAssemblies(assemblies)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "circular dependency detected") {
		t.Errorf("expected circular dependency error, got: %v", err)
	}
}

func TestSortAssemblies_MissingDependency(t *testing.T) {
	assemblies := []AssemblyInterface{
		&mockAssembly{name: "A", deps: []string{"missing"}},
	}

	_, err := SortAssemblies(assemblies)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "not registered") {
		t.Errorf("expected 'not registered' error, got: %v", err)
	}
}

func TestSortAssemblies_DuplicateName(t *testing.T) {
	assemblies := []AssemblyInterface{
		&mockAssembly{name: "A", deps: nil},
		&mockAssembly{name: "A", deps: nil},
	}

	_, err := SortAssemblies(assemblies)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "duplicate") {
		t.Errorf("expected duplicate error, got: %v", err)
	}
}

func names(assemblies []AssemblyInterface) []string {
	result := make([]string, len(assemblies))
	for i, a := range assemblies {
		result[i] = a.Name()
	}
	return result
}
