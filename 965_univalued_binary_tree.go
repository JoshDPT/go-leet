package main

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
  stack := []*TreeNode{root}
  count := 1
  var value int;

  for len(stack) > 0 {
    cur := stack[len(stack)-1]
    stack = stack[:len(stack)-1]

    if count == 1 {
      value = cur.Val
      count += 1
    } else if cur.Val != value {
      return false
    }

    if cur.Left != nil {
      stack = append(stack, cur.Left)
    }
    if cur.Right != nil {
      stack = append(stack, cur.Right)
    }
  }

  return true
}
