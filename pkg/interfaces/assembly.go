package interfaces

import (
	"fmt"
	"sort"
	"strings"
)

type AssemblyInterface interface {
	// Name 返回此 Assembly 的唯一标识（即它注册到容器中的服务名）
	Name() string
	// DependsOn 声明此 Assembly 依赖的其他 Assembly 名称
	DependsOn() []string
	// Assembly 构建服务实例并返回，由框架统一注册到容器
	Assembly() (any, error)
}

// SortAssemblies 对 Assembly 列表进行拓扑排序，按依赖顺序返回。
// 被依赖的 Assembly 排在前面。若存在循环依赖或缺失依赖，返回错误。
func SortAssemblies(assemblies []AssemblyInterface) ([]AssemblyInterface, error) {
	byName := make(map[string]AssemblyInterface, len(assemblies))
	inDegree := make(map[string]int, len(assemblies))
	graph := make(map[string][]string, len(assemblies))

	for _, a := range assemblies {
		name := a.Name()
		if _, exists := byName[name]; exists {
			return nil, fmt.Errorf("duplicate assembly name: %q", name)
		}
		byName[name] = a
		inDegree[name] = 0
	}

	for _, a := range assemblies {
		name := a.Name()
		for _, dep := range a.DependsOn() {
			if _, ok := byName[dep]; !ok {
				return nil, fmt.Errorf("assembly %q depends on %q, which is not registered", name, dep)
			}
			graph[dep] = append(graph[dep], name)
			inDegree[name]++
		}
	}

	var queue []string
	for name, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, name)
		}
	}
	sort.Strings(queue)

	var order []string
	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		order = append(order, name)

		neighbors := graph[name]
		sort.Strings(neighbors)
		for _, next := range neighbors {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
		sort.Strings(queue)
	}

	if len(order) != len(assemblies) {
		cycle := extractAssemblyCycle(byName, inDegree)
		return nil, fmt.Errorf("circular dependency detected: %s", cycle)
	}

	result := make([]AssemblyInterface, len(order))
	for i, name := range order {
		result[i] = byName[name]
	}
	return result, nil
}

func extractAssemblyCycle(byName map[string]AssemblyInterface, inDegree map[string]int) string {
	remaining := make(map[string]bool)
	for name, deg := range inDegree {
		if deg > 0 {
			remaining[name] = true
		}
	}

	var start string
	for name := range remaining {
		start = name
		break
	}

	visited := make(map[string]bool)
	var path []string
	current := start

	for {
		if visited[current] {
			cycleStart := -1
			for i, name := range path {
				if name == current {
					cycleStart = i
					break
				}
			}
			if cycleStart >= 0 {
				cyclePath := append(path[cycleStart:], current)
				return strings.Join(cyclePath, " -> ")
			}
			break
		}
		visited[current] = true
		path = append(path, current)

		a := byName[current]
		deps := a.DependsOn()
		sort.Strings(deps)
		found := false
		for _, dep := range deps {
			if remaining[dep] {
				current = dep
				found = true
				break
			}
		}
		if !found {
			break
		}
	}

	var names []string
	for name := range remaining {
		names = append(names, name)
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}
