package main

import (
	"fmt"
)

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

func newCache() Cache {
	return Cache{
		Queue: newQueue(),
		Hash:  Hash{},
	}
}

func newQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head

	return Queue{
		Head:   head,
		Tail:   tail,
		Length: 0,
	}
}

func (c *Cache) Check(str string) {
	var node *Node

	if val, ok := c.Hash[str]; ok {

		node = c.Remove(val)
	} else {

		node = &Node{Val: str}
	}

	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(node *Node) *Node {
	node.Left.Right = node.Right
	node.Right.Left = node.Left
	c.Queue.Length--
	return node
}

func (c *Cache) Add(node *Node) {
	node.Left = c.Queue.Tail.Left
	node.Right = c.Queue.Tail
	c.Queue.Tail.Left.Right = node
	c.Queue.Tail.Left = node
	c.Queue.Length++
}

func (c *Cache) Display() {
	current := c.Queue.Head.Right
	fmt.Print("Cache contents: ")
	for current != c.Queue.Tail {
		fmt.Printf("%s ", current.Val)
		current = current.Right
	}
	fmt.Println()
}

func main() {
	fmt.Println("Start cache")
	cache := newCache()
	for _, word := range []string{"dog", "cat", "bird", "fish", "mug", "man"} {
		cache.Check(word)
		cache.Display()
	}
}
