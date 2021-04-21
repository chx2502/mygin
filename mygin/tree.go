package mygin

import "strings"
/**
	node 是一棵字典树，gin 的实际实现中用的是基数树。
	相比于字典树，基数树优化了叶子节点，将无分叉的单一路径压缩到一个叶子节点上，
	减少了树的高度，从而实现更高的查找效率和更少的空间使用。
*/
type node struct {
	pattern string 	// pattern to be matched, ex: /p/:lang
	part string		// a part of route	ex: :lang
	children []*node	// children route of current route
	isWild bool	// weather is wild match
}
// 在节点 n 的 children 数组中匹配当前 part 的节点，用于 添加 route 时的精确匹配
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 在节点 n 的 children 数组中匹配所有当前 part 的节点，用于搜索 route 时的模糊匹配
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert a given route
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// search a given route
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
