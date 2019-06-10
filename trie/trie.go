package trie

// TODO Make the trie compact

// Trier is the interface required for the objects we put in the trie
type Trier interface {
	Path() uint32
}

type Node struct {
	item     Trier
	children [2]*Node
}

type Trie struct {
	root *Node
}

func New() *Trie {
	return &Trie{root: &Node{}}
}

func (trie *Trie) Add(trier Trier) {
	node := trie.root
	path := trier.Path()

	for height := 0; height < 32; height++ {
		shifted := path >> uint(31-height)
		index := shifted % 2
		// create the node
		if node.children[index] == nil {
			node.children[index] = &Node{}
		}
		node = node.children[index]
	}
	node.item = trier
}

func (trie *Trie) Search(path uint32) (Trier, bool) {
	node := trie.root

	for height := 0; height < 32; height++ {
		shifted := path >> uint(31-height)
		index := shifted % 2
		node = node.children[index]
		if node == nil {
			return nil, false
		}
	}

	return node.item, true
}
