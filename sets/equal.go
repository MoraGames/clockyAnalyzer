package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(EqualSet)

type EqualSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_Equal(name, label string) *EqualSet {
	var set = EqualSet{Name: name, Label: label, Nums: 0}

	lstTime := time.NewTime(22, 22)
	if !set.Verify(lstTime.SplitTime()) {
		panic("sets: NewSet_Equal: lstTime is not valid")
	}
	set.LastTime = lstTime
	set.Gaps = gaps.NewGaps(lstTime)

	return &set
}

func (s *EqualSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *EqualSet) GetName() string {
	return s.Name
}

func (s *EqualSet) GetLabel() string {
	return s.Label
}

func (s *EqualSet) GetNums() int {
	return s.Nums
}

func (s *EqualSet) IncreaseNums() {
	s.Nums++
}

func (s *EqualSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *EqualSet) Verify(h1, h2, m1, m2 int) bool {
	return h1 == h2 && h2 == m1 && m1 == m2
}
