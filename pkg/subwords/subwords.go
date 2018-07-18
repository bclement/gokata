package main

import "bufio"
import "os"
import "io"
import "fmt"
import "strings"
import "container/list"
import "github.com/bclement/gokata/trie"

/*
storage for a compound word
Str contains the word, Pivot is the start of the second subword
*/
type Subwords struct {
	Str   string
	Pivot int
}

/*
find subwords that together make a length-long word and return through results channel
*/
func FindSubwords(reader *bufio.Reader, length int, results chan *Subwords) error {
	root, words, err := readInput(reader, length)
	if err != nil {
		return err
	}
	for e := words.Front(); e != nil; e = e.Next() {
		word := e.Value.(string)
		node := root
		wordLen := len(word)
		for i, index := 0, 0; index < wordLen; {
			node, i, err = node.FindNext(word[index:])
			if err != nil || node == nil {
				break
			}
			index += i + 1
			if root.Contains(word[index:]) {
				results <- &Subwords{word, index}
				break
			}
		}
	}
	return nil
}

/*
read words from reader. length-long words are returned in list, possible subwords are
used to populate the trie
*/
func readInput(reader *bufio.Reader, length int) (*trie.Node, *list.List, error) {
	root := trie.New()
	words := list.New()
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return nil, nil, err
			}
			break
		}
		line = strings.TrimSpace(line)
		lineLen := len(line)
		if lineLen == length {
			words.PushBack(line)
		} else if lineLen < length {
			root.Insert(line)
		}
	}
	return root, words, nil
}

/*
output routine, drain channel and format to stdout
*/
func output(ch chan *Subwords) {
	for sub := range ch {
		fmt.Printf("%s + %s => %s\n", sub.Str[:sub.Pivot], sub.Str[sub.Pivot:], sub.Str)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("missing argument\n")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	ch := make(chan *Subwords)
	go output(ch)
	err = FindSubwords(reader, 6, ch)
	if err != nil {
		fmt.Printf("%s", err)
	}
}
