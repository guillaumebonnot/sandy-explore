package main

import (
	"fmt"
	"map_explorer/point"
	"map_explorer/trie"
)

func main() {
	// BONUS first we need to find the boundaries, in order to find which type we should use
	find_boundaries()
	// we determine that should use int16 be -499 < x < 499
	explore()
}

// we explore map boundaries
func find_boundaries() {
	for i := int16(0); i < 1000; i++ {
		if point.New(i, int16(0)).IsMine() {
			fmt.Println("The boundaries are : ", i)
			return
		}
	}
}

// we explore them map for real this time
func explore() {
	// the tree to store which points we already checked
	tree := trie.New()
	// the queue to plainify which points we want to check next
	queue := make([]*point.Point, 0, 100)

	queue = planify(tree, queue, point.New(0, 0))

	count := 0
	for len(queue) > 0 {

		// dequeue
		p := queue[0]
		queue = queue[1:]

		// fmt.Println("Check ", p)

		if p.IsMine() {
			count++
		} else {
			for _, neighbour := range p.Neighbours() {
				// say hello to your neighbour
				queue = planify(tree, queue, neighbour)
			}
		}
	}

	fmt.Println("Number of points Sandy can explore (including zero) :", count)
}

func planify(tree *trie.Trie, queue []*point.Point, p *point.Point) []*point.Point {
	// check in the tree it has already been planified
	_, ok := tree.Search(p.Path())
	// already planified
	if ok {
		return queue
	}
	tree.Add(p)
	return append(queue, p)
}
