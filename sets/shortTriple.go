package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(ShortTripleSet)

type ShortTripleSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_ShortTriple(name, label string) *ShortTripleSet {
	lstTime := time.NewTime(23, 56)
	return &ShortTripleSet{lstTime, name, label, 0, gaps.NewGaps(lstTime)}
}

func (s *ShortTripleSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *ShortTripleSet) GetName() string {
	return s.Name
}

func (s *ShortTripleSet) GetLabel() string {
	return s.Label
}

func (s *ShortTripleSet) GetNums() int {
	return s.Nums
}

func (s *ShortTripleSet) IncreaseNums() {
	s.Nums++
}

func (s *ShortTripleSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *ShortTripleSet) Verify(_, h2, m1, m2 int) bool {
	return 3*h2 == (m1*10)+m2
}
