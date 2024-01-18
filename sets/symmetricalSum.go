package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(SymmetricalDifferenceSet)

type SymmetricalSumSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_SymmetricalSum(name, label string) *SymmetricalSumSet {
	lstTime := time.NewTime(23, 56)
	return &SymmetricalSumSet{lstTime, name, label, 0, gaps.NewGaps(lstTime)}
}

func (s *SymmetricalSumSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *SymmetricalSumSet) GetName() string {
	return s.Name
}

func (s *SymmetricalSumSet) GetLabel() string {
	return s.Label
}

func (s *SymmetricalSumSet) GetNums() int {
	return s.Nums
}

func (s *SymmetricalSumSet) IncreaseNums() {
	s.Nums++
}

func (s *SymmetricalSumSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *SymmetricalSumSet) Verify(h1, h2, m1, m2 int) bool {
	sum1, sum2 := h1+m2, h2+m1
	return sum1 == sum2
}
