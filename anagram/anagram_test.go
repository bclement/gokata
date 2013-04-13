package anagram

import "os"
import "fmt"
import "net/http"
import "container/list"
import "bufio"
import "testing"

func TestSmall(t *testing.T) {
	file, err := os.Open("small.txt")
	assertErr(t, err)
	defer file.Close()
	reader := bufio.NewReader(file)
	res, err := Collect(reader)
	assertErr(t, err)
	for _, l := range res.Dict {
		logList(t, l)
		t.Logf("\n")
	}
	assert(t, len(res.Dict) == 7, "anagram groups")
	assert(t, len(res.Longest) == 7, "longest anagram")
	t.Logf(res.Most)
	l := res.Dict[res.Most]
	assert(t, l.Len() == 4, fmt.Sprint("most anagrams", l))
}

func TestLarge(t *testing.T) {
	response, err := http.Get("http://web.archive.org/web/20070704025714/http://pragmaticprogrammer.com/katadata/wordlist.txt")
	assertErr(t, err)
	reader := bufio.NewReader(response.Body)
	res, err := Collect(reader)
	assertErr(t, err)
	assert(t, len(res.Dict) == 2006, fmt.Sprintf("anagram groups %d", len(res.Dict)))
	assert(t, len(res.Longest) == 15, fmt.Sprintf("longest anagram %d", len(res.Longest)))
	l := res.Dict[res.Most]
	assert(t, l.Len() == 7, fmt.Sprintf("most anagrams %d", l.Len()))
}

func logList(t *testing.T, l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		t.Logf("%v", e.Value)
	}
}

func assertErr(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error %s", err)
	}
}

func assert(t *testing.T, result bool, msg string) {
	if !result {
		t.Errorf(msg)
	}
}
