package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(ShortRapidRiseSet)

type ShortRapidRiseSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_ShortRapidRise(name, label string) *ShortRapidRiseSet {
	lstTime := time.NewTime(23, 56)
	return &ShortRapidRiseSet{lstTime, name, label, 0, gaps.NewGaps(lstTime)}
}

func (s *ShortRapidRiseSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *ShortRapidRiseSet) GetName() string {
	return s.Name
}

func (s *ShortRapidRiseSet) GetLabel() string {
	return s.Label
}

func (s *ShortRapidRiseSet) GetNums() int {
	return s.Nums
}

func (s *ShortRapidRiseSet) IncreaseNums() {
	s.Nums++
}

func (s *ShortRapidRiseSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *ShortRapidRiseSet) Verify(_, h2, m1, m2 int) bool {
	return m1 == h2+2 && m2 == m1+2
}
