package set

import "sort"

type Set struct {
	// To ensure unique data is present; not exposed
	hash map[int]bool

	// To hold data; not exposed
	data []int
}

// Get a new Set instance
func NewSet() *Set {
	s := &Set{}
	s.hash = make(map[int]bool)
	s.data = []int{}

	return s
}

// Add an integer to the set
func (s *Set) Add(v int) {
	_, ok := s.hash[v]
	if !ok { // new Add
		s.hash[v] = true
		s.data = append(s.data, v)
	}
	return
}

// Add a list of integers to the set
func (s *Set) AddList(l []int) {
	for _, v := range l {
		s.Add(v)
	}
}

// Sort the set
func (s *Set) Sort() {
	sort.Ints(s.data)
}

// return the set as a list
func (s *Set) Set() []int {
	return s.data
}
