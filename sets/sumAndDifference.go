package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(SumAndDifferenceSet)

type SumAndDifferenceSet struct {
	LastTime time.Time
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_SumAndDifference(label string) *SumAndDifferenceSet {
	lstTime := time.NewTime(23, 00)
	return &SumAndDifferenceSet{lstTime, label, 0, gaps.NewGaps(lstTime)}
}

func (s *SumAndDifferenceSet) GetName() string {
	return "SumAndDifference"
}

func (s *SumAndDifferenceSet) GetLabel() string {
	return s.Label
}

func (s *SumAndDifferenceSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *SumAndDifferenceSet) GetNums() int {
	return s.Nums
}

func (s *SumAndDifferenceSet) IncreaseNums() {
	s.Nums++
}

func (s *SumAndDifferenceSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *SumAndDifferenceSet) Verify(h1, h2, m1, m2 int) bool {
	h, m := (h1*10)+h2, (m1*10)+m2
	return (h + m) == (h - m)
}
