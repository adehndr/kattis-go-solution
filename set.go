package main

type Set struct {
	Data []int
	mapData map[int]bool
}

func InitSet() Set {
	return Set{
		Data: []int{},
		mapData: make(map[int]bool),
	}
}

func (s *Set) Get() []int{
	return s.Data
}

func (s *Set) Add(input int) {
	if _, ok := s.mapData[input]; !ok {
		s.mapData[input] = true
		s.Data = append(s.Data, input)
	}
}
