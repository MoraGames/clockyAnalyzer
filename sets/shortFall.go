package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(ShortFallSet)

type ShortFallSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_ShortFall(name, label string) *ShortFallSet {
	var set = ShortFallSet{Name: name, Label: label, Nums: 0}

	lstTime := time.NewTime(23, 21)
	if !set.Verify(lstTime.SplitTime()) {
		panic("sets: NewSet_Mirror: lstTime is not valid")
	}
	set.LastTime = lstTime
	set.Gaps = gaps.NewGaps(lstTime)

	return &set
}

func (s *ShortFallSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *ShortFallSet) GetName() string {
	return s.Name
}

func (s *ShortFallSet) GetLabel() string {
	return s.Label
}

func (s *ShortFallSet) GetNums() int {
	return s.Nums
}

func (s *ShortFallSet) IncreaseNums() {
	s.Nums++
}

func (s *ShortFallSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *ShortFallSet) Verify(_, h2, m1, m2 int) bool {
	return m1 == h2-1 && m2 == m1-1
}
