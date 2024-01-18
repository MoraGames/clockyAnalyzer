package sets

import (
	"math"
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(PerfectSquareSet)

type PerfectSquareSet struct {
	LastTime time.Time
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_PerfectSquare(label string) *PerfectSquareSet {
	lstTime := time.NewTime(23, 04)
	return &PerfectSquareSet{lstTime, label, 0, gaps.NewGaps(lstTime)}
}

func (s *PerfectSquareSet) GetName() string {
	return "PerfectSquare"
}

func (s *PerfectSquareSet) GetLabel() string {
	return s.Label
}

func (s *PerfectSquareSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *PerfectSquareSet) GetNums() int {
	return s.Nums
}

func (s *PerfectSquareSet) IncreaseNums() {
	s.Nums++
}

func (s *PerfectSquareSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *PerfectSquareSet) Verify(h1, h2, m1, m2 int) bool {
	total := (h1 * 1000) + (h2 * 100) + (m1 * 10) + (m2)
	sqrt := int(math.Sqrt(float64(total)))
	return (sqrt * sqrt) == total
}
