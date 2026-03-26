package container

import (
	"fmt"
	"sort"
	"strings"
)

// getDeps 获取 provider 的依赖列表，未实现 DependencyAware 的返回 nil
func getDeps(p Provider) []string {
	if da, ok := p.(DependencyAware); ok {
		return da.DependsOn()
	}
	return nil
}

// topoSort 使用 Kahn 算法对 providers 进行拓扑排序。
// 返回按依赖顺序排列的 name 列表（被依赖者在前）。
// 若存在循环依赖或缺失依赖，返回错误。
func topoSort(providers map[string]*entry) ([]string, error) {
	// 构建邻接表和入度表
	inDegree := make(map[string]int, len(providers))
	// graph[A] = [B, C] 表示 A 被 B、C 依赖（A 初始化后才能初始化 B、C）
	graph := make(map[string][]string, len(providers))

	for name := range providers {
		inDegree[name] = 0
	}

	for name, e := range providers {
		deps := getDeps(e.provider)
		for _, dep := range deps {
			if _, ok := providers[dep]; !ok {
				return nil, fmt.Errorf("provider %q depends on %q, which is not registered", name, dep)
			}
			graph[dep] = append(graph[dep], name)
			inDegree[name]++
		}
	}

	// 入度为 0 的节点入队（排序保证确定性）
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

		// 对邻居排序保证确定性
		neighbors := graph[name]
		sort.Strings(neighbors)
		for _, next := range neighbors {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
		// 重新排序队列保证确定性
		sort.Strings(queue)
	}

	if len(order) != len(providers) {
		cycle := extractCycle(providers, inDegree)
		return nil, fmt.Errorf("circular dependency detected: %s", cycle)
	}

	return order, nil
}

// extractCycle 从剩余未排序的节点中提取一个循环路径，返回可读字符串
func extractCycle(providers map[string]*entry, inDegree map[string]int) string {
	// 找到一个入度不为 0 的节点开始
	remaining := make(map[string]bool)
	for name, deg := range inDegree {
		if deg > 0 {
			remaining[name] = true
		}
	}

	// 从任意剩余节点开始 DFS 找循环
	var start string
	for name := range remaining {
		start = name
		break
	}

	// 沿依赖链走，直到找到重复节点
	visited := make(map[string]bool)
	var path []string
	current := start

	for {
		if visited[current] {
			// 从 current 开始截取循环部分
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

		// 找 current 的一个依赖（也在 remaining 中）
		deps := getDeps(providers[current].provider)
		found := false
		sort.Strings(deps)
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

	// fallback：列出所有剩余节点
	var names []string
	for name := range remaining {
		names = append(names, name)
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}
