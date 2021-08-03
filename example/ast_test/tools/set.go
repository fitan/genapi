package tools

type Seter interface {
	GetKey() string
}

type Set struct {
	hash       map[string]struct{}
	containers []interface{}
}

func NewSet() *Set {
	return &Set{hash: map[string]struct{}{}, containers: []interface{}{}}
}

func (s *Set) Add(sets ...Seter) {
	for _, set := range sets {
		key := set.GetKey()

		_, ok := s.hash[key]
		if !ok {
			s.hash[key] = struct{}{}
			s.containers = append(s.containers, set)
		}
	}
}

func (s *Set) Get() []interface{} {
	return s.containers
}
