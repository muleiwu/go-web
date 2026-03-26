package container

import (
	"fmt"

	"cnb.cool/mliev/open/go-web/pkg/interfaces"
)

// RegisterAssemblies 按依赖顺序排序 assemblies，逐个调用 Assembly() 并注册到全局容器。
func RegisterAssemblies(assemblies []interfaces.AssemblyInterface) error {
	sorted, err := SortByDependency(assemblies)
	if err != nil {
		return fmt.Errorf("assembly dependency error: %w", err)
	}
	for _, a := range sorted {
		instance, err := a.Assembly()
		if err != nil {
			return fmt.Errorf("assembling %s failed: %w", a.Name(), err)
		}
		Register(NewSimpleProvider(a.Name(), instance))
	}
	return nil
}

// ReloadAssemblies 重置所有容器服务后重新注册 assemblies。
func ReloadAssemblies(assemblies []interfaces.AssemblyInterface) error {
	ResetAll()
	return RegisterAssemblies(assemblies)
}
