package checkout

/*
Pricing rule
*/
type Rule interface {

	/*
	   Return price for added items
	   seen is the number of items previously seen
	   added in the number of new items
	*/
	GetPrice(seen, added int) int

	/*
	   return the id for item this rule defines price for
	*/
	GetId() string
}

/*
flat rate price
id of item
unit price for item
*/
type Flat struct {
	id   string
	unit int
}

/*
create new flat rate price rule
*/
func NewFlat(id string, unit int) *Flat {
	return &Flat{id, unit}
}

/*
see Rule.GetPrice
*/
func (f *Flat) GetPrice(seen, added int) int {
	return added * f.unit
}

/*
see Rule.GetId
*/
func (f *Flat) GetId() string {
	return f.id
}

/*
pricing rule for group discount
flat rate for individual
special price for count
*/
type Group struct {
	Flat
	count   int
	special int
}

/*
create a new group price rule
unit price for individual
special price for count
*/
func NewGroup(id string, unit, count, special int) *Group {
	// naive implementation, breaks for large discounts
	return &Group{Flat{id, unit}, count, special % unit}
}

/*
see Rule.GetPrice
*/
func (g *Group) GetPrice(seen, added int) int {
	rval := 0
	totalCount := seen + added
	for i := seen + 1; i <= totalCount; i++ {
		if i%g.count == 0 {
			rval += g.special
		} else {
			rval += g.Flat.unit
		}
	}
	return rval
}

/*
keeps track of running total
*/
type Register struct {
	rules  map[string]Rule
	counts map[string]int
	total  int
}

/*
create a new register with set of price rules
*/
func New(rules []Rule) *Register {
	rmap := make(map[string]Rule, len(rules))
	cmap := make(map[string]int, len(rules))
	for _, rule := range rules {
		id := rule.GetId()
		rmap[id] = rule
	}
	return &Register{rmap, cmap, 0}
}

/*
add item to register total
*/
func (r *Register) Scan(id string) {
	seen := r.counts[id]
	r.total += r.rules[id].GetPrice(seen, 1)
	r.counts[id] += 1
}

/*
get register total so far
*/
func (r *Register) Total() int {
	return r.total
}
