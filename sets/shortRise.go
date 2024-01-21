package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(ShortRiseSet)

type ShortRiseSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_ShortRise(name, label string) *ShortRiseSet {
	var set = ShortRiseSet{Name: name, Label: label, Nums: 0}

	lstTime := time.NewTime(23, 45)
	if !set.Verify(lstTime.SplitTime()) {
		panic("sets: NewSet_Mirror: lstTime is not valid")
	}
	set.LastTime = lstTime
	set.Gaps = gaps.NewGaps(lstTime)

	return &set
}

func (s *ShortRiseSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *ShortRiseSet) GetName() string {
	return s.Name
}

func (s *ShortRiseSet) GetLabel() string {
	return s.Label
}

func (s *ShortRiseSet) GetNums() int {
	return s.Nums
}

func (s *ShortRiseSet) IncreaseNums() {
	s.Nums++
}

func (s *ShortRiseSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *ShortRiseSet) Verify(_, h2, m1, m2 int) bool {
	return m1 == h2+1 && m2 == m1+1
}
