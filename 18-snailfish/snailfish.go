package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, snailfishNumbers := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	magnitude := calculateMagnitude(snailfishNumbers)
	fmt.Println("Magnitude:", magnitude)

	greatestMagnitude := findGreatestMagnitude(&snailfishNumbers)
	fmt.Println("Greatest magnitude:", greatestMagnitude)

}

func findGreatestMagnitude(snailfishNumbers *[]string) (maxMagnitude int) {
	for i := 0; i < len(*snailfishNumbers); i++ {
		for j := 1; j < len(*snailfishNumbers); j++ {
			snailfishSum := add(newNode((*snailfishNumbers)[i], nil, 0), newNode((*snailfishNumbers)[j], nil, 0))
			magnitude := snailfishSum.magnitude()
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
			}
		}
	}

	return
}

func calculateMagnitude(snailfishNumbers []string) int {
	number := newNode(snailfishNumbers[0], nil, 0)
	for i := 1; i < len(snailfishNumbers); i++ {
		number = add(number, newNode(snailfishNumbers[i], nil, 0))
	}
	return number.magnitude()
}

func add(a, b *Node) (node *Node) {
	node = &Node{
		left:   a,
		right:  b,
		leaf:   false,
		parent: nil,
		depth:  0,
	}
	node.left.parent = node
	node.right.parent = node
	node.left.deepen()
	node.right.deepen()
	node.reduce()
	return
}

type Node struct {
	value               int
	leaf                bool
	left, right, parent *Node
	depth               int
}

func newNode(input string, parent *Node, depth int) (node *Node) {
	if number, err := strconv.Atoi(input); err == nil {
		return &Node{
			value:  number,
			leaf:   true,
			parent: parent,
			depth:  depth,
		}
	}

	if input[1] != '[' {
		index := strings.Index(input, ",")
		node = &Node{
			parent: parent,
			depth:  depth,
		}

		node.left = newNode(input[1:index], node, depth+1)
		node.right = newNode(input[index+1:len(input)-1], node, depth+1)
		return
	}

	counter := 0
	for i := 1; i < len(input)-1; i++ {
		switch input[i] {
		case '[':
			counter++
		case ']':
			counter--
		}

		if counter == 0 {
			node = &Node{
				parent: parent,
				depth:  depth,
			}

			node.left = newNode(input[1:i+1], node, depth+1)
			node.right = newNode(input[i+2:len(input)-1], node, depth+1)
			return
		}
	}

	return
}

func (node *Node) deepen() {
	node.depth++
	if node.left != nil {
		node.left.deepen()
	}
	if node.right != nil {
		node.right.deepen()
	}
}

func (node *Node) reduce() {
	exploded := true
	splitted := true
	for exploded || splitted {
		for exploded {
			exploded = node.explode()
		}
		splitted = node.split()
		exploded = node.explode()
	}
}

func (node *Node) explode() bool {
	n := node.findDepth(4)

	if n == nil {
		return false
	}

	leftNode := n.findLeft()
	if leftNode != nil {
		leftNode.value += n.left.value
	}

	rightNode := n.findRight()
	if rightNode != nil {
		rightNode.value += n.right.value
	}

	*n = Node{
		value:  0,
		leaf:   true,
		parent: n.parent,
		depth:  n.depth,
	}

	return true
}

func (node *Node) findLeft() *Node {
	var base *Node
	current := node
	parent := current.parent
	for parent != nil {
		if parent.right == current {
			base = parent
			break
		}
		current = parent
		parent = parent.parent
	}

	if base != nil {
		base = base.left
		for !base.leaf {
			base = base.right
		}
	}

	return base
}

func (node *Node) findRight() *Node {
	var base *Node
	current := node
	parent := current.parent
	for parent != nil {
		if parent.left == current {
			base = parent
			break
		}
		current = parent
		parent = parent.parent
	}

	if base != nil {
		base = base.right
		for !base.leaf {
			base = base.left
		}
	}

	return base
}

func (node Node) split() bool {
	n := node.findValueGreaterThan(10)

	if n == nil {
		return false
	}

	leftValue := n.value / 2
	rightValue := leftValue + n.value%2

	*n = Node{
		leaf:   false,
		parent: n.parent,
		depth:  n.depth,
	}

	n.left = &Node{
		value:  leftValue,
		leaf:   true,
		parent: n,
		depth:  n.depth + 1,
	}

	n.right = &Node{
		value:  rightValue,
		leaf:   true,
		parent: n,
		depth:  n.depth + 1,
	}

	return true
}

func (node *Node) findDepth(depth int) *Node {
	if !node.leaf && node.depth == depth {
		return node
	}

	if node.left != nil {
		if n := node.left.findDepth(depth); n != nil {
			return n
		}
	}

	if node.right != nil {
		if n := node.right.findDepth(depth); n != nil {
			return n
		}
	}

	return nil
}

func (node *Node) findValueGreaterThan(n int) *Node {
	if node.leaf && node.value >= n {
		return node
	}

	if node.left != nil {
		if n := node.left.findValueGreaterThan(n); n != nil {
			return n
		}
	}

	if node.right != nil {
		if n := node.right.findValueGreaterThan(n); n != nil {
			return n
		}
	}

	return nil
}

func (node *Node) magnitude() int {
	if node.leaf {
		return node.value
	}
	return node.left.magnitude()*3 + node.right.magnitude()*2
}
