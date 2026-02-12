// https://leetcode.com/problems/binary-tree-maximum-path-sum/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Approach:
// 1. Use DFS (postorder traversal) to process children before the current node.
// 2. At each node, calculate the maximum gain from the left and right child.
//    - Ignore negative gains by taking max(childGain, 0).
// 3. The maximum path sum **through the current node** is:
//       node.Val + leftGain + rightGain
//    - Update globalMax if this is larger than the current value.
// 4. Return the **maximum gain to parent**, which is:
//       node.Val + max(leftGain, rightGain)
//    - Only one side can be continued upward to keep the path linear.

func maxPathSum(root *TreeNode) int {
    globalMax := root.Val

    // create a recursive method
    var maxGain func(node *TreeNode) int
    maxGain = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftGain := max(maxGain(node.Left), 0)
        rightGain := max(maxGain(node.Right), 0)

        currentPathSum := node.Val + leftGain + rightGain
        globalMax = max(globalMax, currentPathSum)
        return node.Val + max(leftGain, rightGain)
    }

    maxGain(root)
    return globalMax
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}