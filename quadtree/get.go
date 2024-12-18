package quadtree

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.

func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {

	for i := range contentHolder {
		for j := range contentHolder[i] {
			contentHolder[i][j] = -1
		}
	}

	q.getContentRecursive(q.root, topLeftX, topLeftY, contentHolder)

}

func (q Quadtree) getContentRecursive(n *node, topLeftX, topLeftY int, contentHolder [][]int) {
	if n == nil {
		return
	}

	if n.topLeftX+n.width <= topLeftX || n.topLeftY+n.height <= topLeftY ||
		n.topLeftX >= topLeftX+len(contentHolder[0]) || n.topLeftY >= topLeftY+len(contentHolder) {
		return
	}

	if n.isLeaf {

		startY := max(n.topLeftY, topLeftY)
		startX := max(n.topLeftX, topLeftX)
		endY := min(n.topLeftY+n.height, topLeftY+len(contentHolder))
		endX := min(n.topLeftX+n.width, topLeftX+len(contentHolder[0]))

		for i := startY; i < endY; i++ {
			for j := startX; j < endX; j++ {
				contentHolder[i-topLeftY][j-topLeftX] = n.content
			}
		}
	} else {

		q.getContentRecursive(n.topLeftNode, topLeftX, topLeftY, contentHolder)
		q.getContentRecursive(n.topRightNode, topLeftX, topLeftY, contentHolder)
		q.getContentRecursive(n.bottomLeftNode, topLeftX, topLeftY, contentHolder)
		q.getContentRecursive(n.bottomRightNode, topLeftX, topLeftY, contentHolder)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
