package container

import (
	"fmt"
	"sort"
	"strings"
)

// Sortable 任何具有名称和依赖声明的类型均可参与拓扑排序。
type Sortable interface {
	Name() string
	DependsOn() []string
}

// SortByDependency 使用 Kahn 算法对实现 Sortable 接口的元素进行拓扑排序。
// 返回按依赖顺序排列的列表（被依赖者在前）。
// 若存在循环依赖、缺失依赖或重复名称，返回错误。
func SortByDependency[T Sortable](items []T) ([]T, error) {
	byName := make(map[string]T, len(items))
	inDegree := make(map[string]int, len(items))
	graph := make(map[string][]string, len(items))

	for _, item := range items {
		name := item.Name()
		if _, exists := byName[name]; exists {
			return nil, fmt.Errorf("duplicate name: %q", name)
		}
		byName[name] = item
		inDegree[name] = 0
	}

	for _, item := range items {
		name := item.Name()
		for _, dep := range item.DependsOn() {
			if _, ok := byName[dep]; !ok {
				return nil, fmt.Errorf("%q depends on %q, which is not registered", name, dep)
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

	if len(order) != len(items) {
		cycle := extractSortableCycle(byName, inDegree)
		return nil, fmt.Errorf("circular dependency detected: %s", cycle)
	}

	result := make([]T, len(order))
	for i, name := range order {
		result[i] = byName[name]
	}
	return result, nil
}

// extractSortableCycle 从剩余未排序的节点中提取一个循环路径，返回可读字符串
func extractSortableCycle[T Sortable](byName map[string]T, inDegree map[string]int) string {
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

		item := byName[current]
		deps := item.DependsOn()
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

// getDeps 获取 provider 的依赖列表，未实现 DependencyAware 的返回 nil
func getDeps(p Provider) []string {
	if da, ok := p.(DependencyAware); ok {
		return da.DependsOn()
	}
	return nil
}

// providerSortable 将 entry 包装为 Sortable，用于 topoSort 内部复用泛型排序
type providerSortable struct {
	name     string
	provider Provider
}

func (p *providerSortable) Name() string        { return p.name }
func (p *providerSortable) DependsOn() []string { return getDeps(p.provider) }

// topoSort 使用 Kahn 算法对 providers 进行拓扑排序。
// 返回按依赖顺序排列的 name 列表（被依赖者在前）。
func topoSort(providers map[string]*entry) ([]string, error) {
	items := make([]*providerSortable, 0, len(providers))
	for name, e := range providers {
		items = append(items, &providerSortable{name: name, provider: e.provider})
	}

	sorted, err := SortByDependency(items)
	if err != nil {
		return nil, err
	}

	order := make([]string, len(sorted))
	for i, item := range sorted {
		order[i] = item.name
	}
	return order, nil
}
