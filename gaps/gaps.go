package gaps

import "other/bp/cuwu_setstest/time"

type (
	Gaps struct {
		Lst time.Time
		Cur int
		Sum int
		All []int
		Mma bool
		Min GapsPair
		Max GapsPair
		Avg float64
	}

	GapsPair struct {
		Val int
		Prv time.Time
		Cur time.Time
	}
)

func NewGaps(lastTime time.Time) *Gaps {
	return &Gaps{
		Lst: lastTime,
		Cur: lastTime.UntilMidnight(),
		Sum: 0,
		All: make([]int, 0),
		Mma: false,
		Min: GapsPair{0, time.Time{}, time.Time{}},
		Max: GapsPair{0, time.Time{}, time.Time{}},
		Avg: 0.0,
	}
}

func (g *Gaps) UpdateGaps(cur time.Time) {
	g.UpdateSum()
	g.UpdateAll()
	g.UpdateMma(cur)
	g.ResetCur(cur)
}

func (g *Gaps) IncreaseCur(cur time.Time) {
	g.Cur++
}

func (g *Gaps) ResetCur(cur time.Time) {
	g.Lst = cur
	g.Cur = 1
}

func (g *Gaps) UpdateSum() {
	g.Sum += g.Cur
}

func (g *Gaps) UpdateAll() {
	g.All = append(g.All, g.Cur)
}

func (g *Gaps) UpdateMma(cur time.Time) {
	if g.Mma {
		if g.Cur < g.Min.Val {
			g.Min = GapsPair{g.Cur, g.Lst, cur}
		}
		if g.Cur > g.Max.Val {
			g.Max = GapsPair{g.Cur, g.Lst, cur}
		}
	} else {
		g.Mma = true
		g.Min = GapsPair{g.Cur, g.Lst, cur}
		g.Max = GapsPair{g.Cur, g.Lst, cur}
	}
}

func (g *Gaps) UpdateAvg() {
	g.Avg = float64(g.Sum) / float64(len(g.All))
}
