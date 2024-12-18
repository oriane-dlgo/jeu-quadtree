package quadtree

func MakeFromArray(floorContent [][]int) (q Quadtree) {

	/* initialise la taille du terrain avec height et width */

	height := len(floorContent)
	width := len(floorContent[0])

	root := creanode(floorContent, 0, 0, height, width)

	return Quadtree{
		height: height,
		width:  width,
		root:   root,
	}
}

func creanode(floorContent [][]int, topLeftX int, topLeftY int, height int, width int) *node {

	/*les 2 if vérifient si le noeud est une leaf. Dans ces cas la condition leaf de &node (le noeud) est à True*/

	if height == 1 && width == 1 { /* regarde si le terrain fais 1x1*/
		return &node{
			topLeftX: topLeftX,
			topLeftY: topLeftY,
			width:    width,
			height:   height,
			content:  floorContent[topLeftY][topLeftX],
			isLeaf:   true,
		}
	}

	if isHomogeneous(floorContent, topLeftX, topLeftY, height, width) { /*regarde si toutes les cases comportent la meme valeur*/
		return &node{
			topLeftX: topLeftX,
			topLeftY: topLeftY,
			width:    width,
			height:   height,
			content:  floorContent[topLeftY][topLeftX],
			isLeaf:   true,
		}
	}

	/* divise le terrain en deux en fonction de la ou est le noeud */

	midHeight := topLeftY + height/2
	midWidth := topLeftX + width/2

	/* divise le terrain en 4, en bas à droite/gauche et en haut à gauche/droite*/

	topLeftNode := creanode(floorContent, topLeftX, topLeftY, height/2, width/2)
	topRightNode := creanode(floorContent, midWidth, topLeftY, height/2, width-width/2)     /* width/2 et width-width/2 donne mathématiquement la même valeur mais cela permet de faire comprendre à la machine qu'on parle du coté droit et pas du coté gauche*/
	bottomLeftNode := creanode(floorContent, topLeftX, midHeight, height-height/2, width/2) /* pareil pour height-height/2 qui désigne le carré d'en bas*/
	bottomRightNode := creanode(floorContent, midWidth, midHeight, height-height/2, width-width/2)

	/* comme le noeud n'est pas une leaf alors la condition leaf est à false et on renseigne donc les donné qui suivent*/

	return &node{
		topLeftX:        topLeftX,
		topLeftY:        topLeftY,
		width:           width,
		height:          height,
		isLeaf:          false,
		topLeftNode:     topLeftNode,
		topRightNode:    topRightNode,
		bottomLeftNode:  bottomLeftNode,
		bottomRightNode: bottomRightNode,
	}
}

func isHomogeneous(floorContent [][]int, topLeftX, topLeftY, height, width int) bool {
	value := floorContent[topLeftY][topLeftX]     /* on prend la valeur de la première case et on regarde si toutes les autres sont égaux ou non */
	for y := topLeftY; y < topLeftY+height; y++ { /* important de garder les valeurs de topleft pour garder la position du noeud par rapport au terrain */
		for x := topLeftX; x < topLeftX+width; x++ {
			if floorContent[y][x] != value {
				return false
			}
		}
	}
	return true
}
