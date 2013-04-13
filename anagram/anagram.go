package anagram

import "io"
import "sort"
import "bufio"
import "strings"
import "container/list"

/*
anagram collection results
Dict is a map of ordered letters to lists of anagrams
Longest is the key with the most letters
Most is the key to the biggest list of anagrams
*/
type Results struct {
	Dict    map[string]*list.List
	Longest string
	Most    string
}

/*
Collect anagrams from readers, one per line
*/
func Collect(reader *bufio.Reader) (*Results, error) {
	rval := &Results{make(map[string]*list.List), "", ""}
	most := -1
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		line = strings.TrimSpace(line)
		sorted := getkey(line)
		l, ok := rval.Dict[sorted]
		if !ok {
			l = list.New()
			rval.Dict[sorted] = l
		}
		l.PushBack(line)
		if l.Len() > most {
			rval.Most = sorted
			most = l.Len()
		}
	}
	cull(rval)
	return rval, nil
}

/*
remove one-entry results
finalize stats
*/
func cull(res *Results) {
	longest := -1
	for key, l := range res.Dict {
		if l.Len() < 2 {
			delete(res.Dict, key)
			continue
		}
		if len(key) > longest {
			res.Longest = key
			longest = len(key)
		}
	}
}

/*
get key from word
*/
func getkey(s string) string {
	letters := strings.Split(s, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}
