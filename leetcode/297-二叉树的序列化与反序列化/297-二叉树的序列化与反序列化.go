package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 297-二叉树的序列化与反序列化
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")

	var dfs func(*[]string) *TreeNode
	dfs = func(s *[]string) *TreeNode {
		val := (*s)[0]
		*s = (*s)[1:]
		if val == "#" {
			return nil
		}
		valInt, _ := strconv.Atoi(val)
		node := &TreeNode{Val: valInt}
		node.Left = dfs(s)
		node.Right = dfs(s)
		return node
	}

	return dfs(&vals)

}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
