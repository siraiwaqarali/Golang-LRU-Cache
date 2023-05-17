// LRU Cache: Maintains a specific length of items in cache
// LRU - Only the most recently used items are maintained in the cache
// Item goes out from tail, when the cache is full

// For a true LRU Cache
// 1. If an item is already exists, we need to remove it and add it to the beginning
// 2. An order of items is maintained
// 3. Deletion happens at the tail and addition happens at the head

package main

import "fmt"

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail, Length: 0}
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()

	for _, str := range []string{"STR1", "STR2", "STR3", "STR4", "STR5", "STR6"} {
		cache.Check(str)
		cache.Display()
	}
}

func (c *Cache) Check(str string) {
	node := &Node{}
	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Add(node *Node) {
	fmt.Println("ADD:", node.Val)
	rightOfHead := c.Queue.Head.Right // right of head

	c.Queue.Head.Right = node // right of head will be new node
	node.Left = c.Queue.Head  // left of new node will be head
	node.Right = rightOfHead  // right of new node will be right of head
	rightOfHead.Left = node   // left of right of head will be new node

	c.Queue.Length++ // increment length
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
	c.Hash[node.Val] = node // add to hash
}

func (c *Cache) Remove(node *Node) *Node {
	fmt.Println("REMOVE:", node.Val)
	left := node.Left   // left of (node to remove)
	right := node.Right // right of (node to remove)

	left.Right = right // right of left will be right of (node to remove)
	right.Left = left  // left of right will be left of (node to remove)

	c.Queue.Length -= 1      // decrement length
	delete(c.Hash, node.Val) // remove from hash

	return node
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	fmt.Printf("%d - [", q.Length)

	node := q.Head.Right
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right // move to right
	}
	fmt.Println("]")
}
