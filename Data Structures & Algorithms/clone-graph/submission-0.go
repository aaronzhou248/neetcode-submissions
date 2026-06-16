func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }

    // Maps original node -> its clone
    clones := map[*Node]*Node{}

    var dfs func(n *Node) *Node
    dfs = func(n *Node) *Node {
        // If we already cloned this node, return the clone
        if clone, exists := clones[n]; exists {
            return clone
        }

        // Create the clone (without neighbors yet)
        clone := &Node{Val: n.Val}
        clones[n] = clone // Mark as visited BEFORE recursing (prevents infinite loops)

        // Recursively clone all neighbors
        for _, neighbor := range n.Neighbors {
            clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
        }

        return clone
    }

    return dfs(node)
}