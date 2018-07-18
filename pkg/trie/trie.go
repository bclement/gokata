package trie

import "strings"
import "fmt"

type Node struct {
	arr  [26]*Node
	leaf bool
}

/*
create a new trie node
*/
func New() *Node {
	return &Node{[26]*Node{}, false}
}

/*
returns *Node, int < len(key) if trie doesn't contain key as prefix
returns *Node, int == len(key) if trie contains key as prefix
returned node needs to be checked as leaf
*/
func (n *Node) Find(key string) (*Node, int, error) {
	count := 0
	key = strings.ToLower(key)
	keyLen := len(key)
	curr := n
	for ; count < keyLen; count += 1 {
		i := key[count] - 'a'
		if i < 0 || i > 25 {
			return nil, 0, fmt.Errorf("input out of range")
		}
		next := curr.arr[i]
		if next == nil {
			break
		}
		curr = next
	}
	return curr, count, nil
}

/*
finds the next subword in key
returns node that ends that word or nil if none found
*/
func (n *Node) FindNext(key string) (*Node, int, error) {
	count := 0
	key = strings.ToLower(key)
	keyLen := len(key)
	curr := n
	for ; count < keyLen; count += 1 {
		i := key[count] - 'a'
		if i < 0 || i > 25 {
			return nil, 0, fmt.Errorf("input out of range")
		}
		next := curr.arr[i]
		if next == nil {
			return nil, count, nil
		}
		if next.leaf {
			return next, count, nil
		}
		curr = next
	}
	return nil, count, nil
}

/* 
return true if trie contains key
*/
func (n *Node) Contains(key string) bool {
	node, i, err := n.Find(key)
	if err != nil {
		return false
	}
	return i == len(key) && node.leaf
}

/*
insert string into trie
*/
func (n *Node) Insert(str string) error {
	strLen := len(str)
	if strLen == 0 {
		return fmt.Errorf("cannot insert empty string")
	}
	lower := strings.ToLower(str)
	curr, index, err := n.Find(lower)
	if err != nil {
		return err
	}
	if index == strLen {
		curr.leaf = true
		return nil
	}
	for ; index < strLen; index++ {
		i := lower[index] - 'a'
		if i < 0 || i > 25 {
			return fmt.Errorf("input out of range")
		}
		tmp := New()
		curr.arr[i] = tmp
		curr = tmp
	}
	curr.leaf = true
	return nil
}
