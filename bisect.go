package bisect

// Package bisect implements common bisection algorithms

// Interface is a type, typically a collection can be
// sorted in-place by the routines in this package
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Compares a given item with the element x at index i
	// Should return a bool:
	//    negative , if a[i] < than
	//    positive , if a[i] > than
	Less(i int, than interface{}) bool
}

// Return the index where to insert item x in a, assuming a is sorted.
// The return value i is such that all e in a[:i] have e <= x, and all e in
// a[i:] have e > x.  So if x already appears in the list, a.insert(x) will
// insert just after the rightmost x already there.
func BisectRight(a Interface, x interface{}) int {
	lo, hi := 0, a.Len()
	for lo < hi {
		m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
		if a.Less(m, x) {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return lo
}

// Return the index where to insert item x in a, assuming a is sorted.
// The return value i is such that all e in a[:i] have e < x, and all e in
// a[i:] have e >= x.  So if x already appears in the list, a.insert(x) will
// insert just before the leftmost x already there.
func BisectLeft(a Interface, x interface{}) int {
	lo, hi := 0, a.Len()
	for lo < hi {
		m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
		if a.Less(m, x) {
			lo = m + 1
		} else {
			hi = m
		}
	}
	return lo
}

// Convenience types for common cases

// IntSlice attaches the methods of Items to []int
type IntSlice []int

func (p IntSlice) Len() int                          { return len(p) }
func (p IntSlice) Less(i int, with interface{}) bool { return p[i] < with.(int) }
func (p *IntSlice) Insert(i int, x interface{}) {
	*p = append(*p, 0)
	copy((*p)[i+1:], (*p)[i:])
	(*p)[i] = x.(int)
}

var _ Interface = (*IntSlice)(nil)

// Convenience methods.
func (p *IntSlice) InsortRight(x int) { p.Insert(BisectRight(p, x), x) }
func (p *IntSlice) InsortLeft(x int)  { p.Insert(BisectLeft(p, x), x) }
func (p *IntSlice) BisectRight(x int) { BisectRight(p, x) }
func (p *IntSlice) BisectLeft(x int)  { BisectLeft(p, x) }

// StringSlice attaches the methods of Items to []string
type StringSlice []string

func (p StringSlice) Len() int                          { return len(p) }
func (p StringSlice) Less(i int, with interface{}) bool { return p[i] < with.(string) }
func (p *StringSlice) Insert(i int, x interface{}) {
	*p = append(*p, "")
	copy((*p)[i+1:], (*p)[i:])
	(*p)[i] = x.(string)
}

// Convenience methods.
func (p *StringSlice) InsortRight(x string) { p.Insert(BisectRight(p, x), x) }
func (p *StringSlice) InsortLeft(x string)  { p.Insert(BisectLeft(p, x), x) }
func (p *StringSlice) BisectRight(x string) { BisectRight(p, x) }
func (p *StringSlice) BisectLeft(x string)  { BisectLeft(p, x) }

var _ Interface = (*StringSlice)(nil)

// Float64Slice attaches the methods of items to []float64
type Float64Slice []float64

func (p Float64Slice) Len() int { return len(p) }
func (p Float64Slice) Less(i int, with interface{}) bool {
	return p[i] < with.(float64) || isNaN(p[i]) && !isNaN(with.(float64))
}
func (p *Float64Slice) Insert(i int, x interface{}) {
	*p = append(*p, float64(0))
	copy((*p)[i+1:], (*p)[i:])
	(*p)[i] = x.(float64)
}

// isNaN is a copy of math.IsNaN to avoid a dependency on the math package.
func isNaN(f float64) bool {
	return f != f
}

// InsortRight is a convenience method.
func (p *Float64Slice) InsortRight(x float64) { p.Insert(BisectRight(p, x), x) }
func (p *Float64Slice) InsortLeft(x float64)  { p.Insert(BisectLeft(p, x), x) }
func (p *Float64Slice) BisectRight(x float64) { BisectRight(p, x) }
func (p *Float64Slice) BisectLeft(x float64)  { BisectLeft(p, x) }

var _ Interface = (*Float64Slice)(nil)
