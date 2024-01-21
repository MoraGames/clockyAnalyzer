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
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_PerfectSquare(name, label string) *PerfectSquareSet {
	var set = PerfectSquareSet{Name: name, Label: label, Nums: 0}

	lstTime := time.NewTime(23, 04)
	if !set.Verify(lstTime.SplitTime()) {
		panic("sets: NewSet_Mirror: lstTime is not valid")
	}
	set.LastTime = lstTime
	set.Gaps = gaps.NewGaps(lstTime)

	return &set
}

func (s *PerfectSquareSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *PerfectSquareSet) GetName() string {
	return s.Name
}

func (s *PerfectSquareSet) GetLabel() string {
	return s.Label
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
