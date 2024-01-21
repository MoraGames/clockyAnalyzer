package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(DoubleSet)

type DoubleSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_Double(name, label string) *DoubleSet {
	var set = DoubleSet{Name: name, Label: label, Nums: 0}

	lstTime := time.NewTime(23, 46)
	if !set.Verify(lstTime.SplitTime()) {
		panic("sets: NewSet_Double: lstTime is not valid")
	}
	set.LastTime = lstTime
	set.Gaps = gaps.NewGaps(lstTime)

	return &set
}

func (s *DoubleSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *DoubleSet) GetName() string {
	return s.Name
}

func (s *DoubleSet) GetLabel() string {
	return s.Label
}

func (s *DoubleSet) GetNums() int {
	return s.Nums
}

func (s *DoubleSet) IncreaseNums() {
	s.Nums++
}

func (s *DoubleSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *DoubleSet) Verify(h1, h2, m1, m2 int) bool {
	return 2*((h1*10)+h2) == (m1*10)+m2
}
