package sets

import (
	"math"
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(SymmetricalDifferenceSet)

type SymmetricalDifferenceSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_SymmetricalDifference(name, label string) *SymmetricalDifferenceSet {
	var set = SymmetricalDifferenceSet{Name: name, Label: label, Nums: 0}

	lstTime := time.NewTime(23, 54)
	if !set.Verify(lstTime.SplitTime()) {
		panic("sets: NewSet_Mirror: lstTime is not valid")
	}
	set.LastTime = lstTime
	set.Gaps = gaps.NewGaps(lstTime)

	return &set
}

func (s *SymmetricalDifferenceSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *SymmetricalDifferenceSet) GetName() string {
	return s.Name
}

func (s *SymmetricalDifferenceSet) GetLabel() string {
	return s.Label
}

func (s *SymmetricalDifferenceSet) GetNums() int {
	return s.Nums
}

func (s *SymmetricalDifferenceSet) IncreaseNums() {
	s.Nums++
}

func (s *SymmetricalDifferenceSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *SymmetricalDifferenceSet) Verify(h1, h2, m1, m2 int) bool {
	diff1, diff2 := math.Abs(float64(h1-m2)), math.Abs(float64(h2-m1))
	return diff1 == diff2
}
