package container

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Sortable 任何具有类型标识和依赖声明的类型均可参与拓扑排序。
type Sortable interface {
	Type() reflect.Type
	DependsOn() []reflect.Type
}

// SortByDependency 使用 Kahn 算法对实现 Sortable 接口的元素进行拓扑排序。
// 返回按依赖顺序排列的列表（被依赖者在前）。
// 若存在循环依赖、缺失依赖或重复类型，返回错误。
func SortByDependency[T Sortable](items []T) ([]T, error) {
	byType := make(map[reflect.Type]T, len(items))
	inDegree := make(map[reflect.Type]int, len(items))
	graph := make(map[reflect.Type][]reflect.Type, len(items))

	for _, item := range items {
		t := item.Type()
		if _, exists := byType[t]; exists {
			return nil, fmt.Errorf("duplicate type: %v", t)
		}
		byType[t] = item
		inDegree[t] = 0
	}

	for _, item := range items {
		t := item.Type()
		for _, dep := range item.DependsOn() {
			if _, ok := byType[dep]; !ok {
				return nil, fmt.Errorf("%v depends on %v, which is not registered", t, dep)
			}
			graph[dep] = append(graph[dep], t)
			inDegree[t]++
		}
	}

	queue := collectZeroDegree(inDegree)

	var order []reflect.Type
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		order = append(order, t)

		neighbors := graph[t]
		sortTypes(neighbors)
		for _, next := range neighbors {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
		sortTypes(queue)
	}

	if len(order) != len(items) {
		cycle := extractSortableCycle(byType, inDegree)
		return nil, fmt.Errorf("circular dependency detected: %s", cycle)
	}

	result := make([]T, len(order))
	for i, t := range order {
		result[i] = byType[t]
	}
	return result, nil
}

// extractSortableCycle 从剩余未排序的节点中提取一个循环路径，返回可读字符串
func extractSortableCycle[T Sortable](byType map[reflect.Type]T, inDegree map[reflect.Type]int) string {
	remaining := make(map[reflect.Type]bool)
	for t, deg := range inDegree {
		if deg > 0 {
			remaining[t] = true
		}
	}

	var start reflect.Type
	for t := range remaining {
		start = t
		break
	}

	visited := make(map[reflect.Type]bool)
	var path []reflect.Type
	current := start

	for {
		if visited[current] {
			cycleStart := -1
			for i, t := range path {
				if t == current {
					cycleStart = i
					break
				}
			}
			if cycleStart >= 0 {
				cyclePath := append(path[cycleStart:], current)
				names := make([]string, len(cyclePath))
				for i, t := range cyclePath {
					names[i] = t.String()
				}
				return strings.Join(names, " -> ")
			}
			break
		}
		visited[current] = true
		path = append(path, current)

		item := byType[current]
		deps := item.DependsOn()
		sortTypes(deps)
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
	for t := range remaining {
		names = append(names, t.String())
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}

// getDeps 获取 provider 的依赖列表，未实现 DependencyAware 的返回 nil
func getDeps(p Provider) []reflect.Type {
	if da, ok := p.(DependencyAware); ok {
		return da.DependsOn()
	}
	return nil
}

// providerSortable 将 entry 包装为 Sortable，用于 topoSort 内部复用泛型排序
type providerSortable struct {
	typ      reflect.Type
	provider Provider
}

func (p *providerSortable) Type() reflect.Type        { return p.typ }
func (p *providerSortable) DependsOn() []reflect.Type { return getDeps(p.provider) }

// topoSort 使用 Kahn 算法对 providers 进行拓扑排序。
// 返回按依赖顺序排列的 reflect.Type 列表（被依赖者在前）。
func topoSort(providers map[reflect.Type]*entry) ([]reflect.Type, error) {
	items := make([]*providerSortable, 0, len(providers))
	for t, e := range providers {
		items = append(items, &providerSortable{typ: t, provider: e.provider})
	}

	sorted, err := SortByDependency(items)
	if err != nil {
		return nil, err
	}

	order := make([]reflect.Type, len(sorted))
	for i, item := range sorted {
		order[i] = item.typ
	}
	return order, nil
}

// sortTypes 对 reflect.Type 切片按字符串表示排序，保证确定性
func sortTypes(types []reflect.Type) {
	sort.Slice(types, func(i, j int) bool {
		return types[i].String() < types[j].String()
	})
}

// collectZeroDegree 收集入度为 0 的节点并排序
func collectZeroDegree(inDegree map[reflect.Type]int) []reflect.Type {
	var queue []reflect.Type
	for t, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, t)
		}
	}
	sortTypes(queue)
	return queue
}
