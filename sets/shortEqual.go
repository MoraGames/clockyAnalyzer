package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(ShortEqualSet)

type ShortEqualSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_ShortEqual(name, label string) *ShortEqualSet {
	var set = ShortEqualSet{Name: name, Label: label, Nums: 0}

	lstTime := time.NewTime(23, 33)
	if !set.Verify(lstTime.SplitTime()) {
		panic("sets: NewSet_Mirror: lstTime is not valid")
	}
	set.LastTime = lstTime
	set.Gaps = gaps.NewGaps(lstTime)

	return &set
}

func (s *ShortEqualSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *ShortEqualSet) GetName() string {
	return s.Name
}

func (s *ShortEqualSet) GetLabel() string {
	return s.Label
}

func (s *ShortEqualSet) GetNums() int {
	return s.Nums
}

func (s *ShortEqualSet) IncreaseNums() {
	s.Nums++
}

func (s *ShortEqualSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *ShortEqualSet) Verify(_, h2, m1, m2 int) bool {
	return h2 == m1 && m1 == m2
}
