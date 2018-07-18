package checkout

import "testing"
import "strings"

func price(t *testing.T, goods string, rules []Rule) int {
	co := New(rules)
	for _, item := range strings.Split(goods, "") {
		co.Scan(item)
	}
	return co.Total()
}

func TestExample(t *testing.T) {
	rules := []Rule{NewGroup("A", 50, 3, 130), NewGroup("B", 30, 2, 45),
		NewFlat("C", 20), NewFlat("D", 15)}

	assertPrice(t, 0, "", rules)
	assertPrice(t, 50, "A", rules)
	assertPrice(t, 80, "AB", rules)
	assertPrice(t, 115, "CDBA", rules)

	assertPrice(t, 100, "AA", rules)
	assertPrice(t, 130, "AAA", rules)
	assertPrice(t, 180, "AAAA", rules)
	assertPrice(t, 230, "AAAAA", rules)
	assertPrice(t, 260, "AAAAAA", rules)

	assertPrice(t, 160, "AAAB", rules)
	assertPrice(t, 175, "AAABB", rules)
	assertPrice(t, 190, "AAABBD", rules)
	assertPrice(t, 190, "DABABA", rules)
}

func assertPrice(t *testing.T, expected int, items string, rules []Rule) {
	actual := price(t, items, rules)
	if expected != actual {
		t.Errorf("%s expected %d, got %d\n", items, expected, actual)
	}
}
