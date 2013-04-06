package trie

import "testing"
import "fmt"

func TestInsert(t *testing.T) {
	root := New()
	root.Insert("f")
	root.Insert("foo")
	root.Insert("foobar")
	root.Insert("bar")
	find(t, root, "foo")
	find(t, root, "foobar")
	find(t, root, "bar")
}

func find(t *testing.T, root *Node, str string) {
	node, i, _ := root.Find(str)
	assert(t, i == len(str), "find parsed "+str)
	assert(t, node.leaf, str+" is leaf")
}

func TestContains(t *testing.T) {
	root := New()
	root.Insert("foo")
	assert(t, root.Contains("foo"), "contains foo")
	assert(t, !root.Contains("bar"), "doesn't have bar")
}

func TestFindNext(t *testing.T) {
	root := New()
	root.Insert("fu")
	root.Insert("fuo")
	root.Insert("bar")
	word := "fuobar"
	n, i, _ := root.FindNext(word)
	assert(t, i == 1, fmt.Sprintf("found fu at %d", i))
	assert(t, n.leaf, "found leaf")
	n, i, _ = n.FindNext(word[2:])
	assert(t, i == 0, fmt.Sprintf("found fuo at %d", i))
	assert(t, n.leaf, "found leaf")

}

func assert(t *testing.T, cond bool, msg string) {
	if !cond {
		t.Errorf("%s", msg)
	}
}
