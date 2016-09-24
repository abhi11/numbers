package set

import "sort"

type Set struct {
	// To make sure unique data in present
	Hash map[int]bool

	// To hold data
	Data []int
}

func NewSet() *Set {
	s := &Set{}
	s.Hash = make(map[int]bool)
	s.Data = []int{}

	return s
}

func (s *Set) Add(v int) {
	_, ok := s.Hash[v]
	if !ok { // new Add
		s.Hash[v] = true
		s.Data = append(s.Data, v)
	}
	return
}

func (s *Set) AddList(l []int) {
	for _, v := range l {
		s.Add(v)
	}
}

func (s *Set) Sort() []int {
	// sort data and return, creates a side effect in the set order
	sort.Ints(s.Data)
	return s.Data
}
