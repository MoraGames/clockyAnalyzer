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
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_SymmetricalDifference(label string) *SymmetricalDifferenceSet {
	lstTime := time.NewTime(23, 54)
	return &SymmetricalDifferenceSet{lstTime, label, 0, gaps.NewGaps(lstTime)}
}

func (s *SymmetricalDifferenceSet) GetName() string {
	return "SymmetricalDifference"
}

func (s *SymmetricalDifferenceSet) GetLabel() string {
	return s.Label
}

func (s *SymmetricalDifferenceSet) GetLastTime() time.Time {
	return s.LastTime
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
