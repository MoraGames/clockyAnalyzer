package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(RepeatSet)

type RepeatSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_Repeat(name, label string) *RepeatSet {
	lstTime := time.NewTime(23, 56)
	return &RepeatSet{lstTime, name, label, 0, gaps.NewGaps(lstTime)}
}

func (s *RepeatSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *RepeatSet) GetName() string {
	return s.Name
}

func (s *RepeatSet) GetLabel() string {
	return s.Label
}

func (s *RepeatSet) GetNums() int {
	return s.Nums
}

func (s *RepeatSet) IncreaseNums() {
	s.Nums++
}

func (s *RepeatSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *RepeatSet) Verify(h1, h2, m1, m2 int) bool {
	return h1 == m1 && h2 == m2 && h1 != h2
}
