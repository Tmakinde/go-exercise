// https://leetcode.com/problems/binary-tree-level-order-traversal/

// Use BFS (level-order traversal)
// Start with the root node in a queue
// While the queue is not empty:
//   - Freeze the current queue size (this is one level)
//   - Process exactly those nodes
//   - Append their children to the queue (for the next level)
//   - Remove the processed nodes from the queue
 
func levelOrder(root *TreeNode) [][]int {
    queue := []*TreeNode{root}

    var result [][]int
    if root == nil {
        return [][]int{}
    }

    for len(queue) > 0 {
        levelSize := len(queue) // loop queue once
        level := []int{}
        for i := 0; i < levelSize; i++ {
            element := queue[i]
            level = append(level, element.Val)
            if element.Left != nil {
                queue = append(queue, element.Left)
            }

            if element.Right != nil {
                queue = append(queue, element.Right)
            }
        }

        queue = queue[levelSize:]
        result = append(result, level)
    }
    return result
}
