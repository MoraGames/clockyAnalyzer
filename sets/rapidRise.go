package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(RapidRiseSet)

type RapidRiseSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_RapidRise(name, label string) *RapidRiseSet {
	lstTime := time.NewTime(23, 56)
	return &RapidRiseSet{lstTime, name, label, 0, gaps.NewGaps(lstTime)}
}

func (s *RapidRiseSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *RapidRiseSet) GetName() string {
	return s.Name
}

func (s *RapidRiseSet) GetLabel() string {
	return s.Label
}

func (s *RapidRiseSet) GetNums() int {
	return s.Nums
}

func (s *RapidRiseSet) IncreaseNums() {
	s.Nums++
}

func (s *RapidRiseSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *RapidRiseSet) Verify(h1, h2, m1, m2 int) bool {
	return h2 == h1+2 && m1 == h2+2 && m2 == m1+2
}
