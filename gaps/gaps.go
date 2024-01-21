package gaps

import "github.com/MoraGames/clockyAnalyzer/time"

type (
	Gaps struct {
		Lst time.Time
		Cur int
		Sum int
		All []int
		Mma bool
		Min GapsTriad
		Max GapsTriad
		Avg float64
	}

	GapsTriad struct {
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
		Min: GapsTriad{0, time.Time{}, time.Time{}},
		Max: GapsTriad{0, time.Time{}, time.Time{}},
		Avg: 0.0,
	}
}

func (g *Gaps) UpdateGaps(cur time.Time) {
	// Update the total gaps sum
	g.Sum += g.Cur
	// Append the current gap to the gaps slice
	g.All = append(g.All, g.Cur)
	// Update the minimum and maximum gaps (if necessary)
	if g.Mma {
		if g.Cur < g.Min.Val {
			g.Min = GapsTriad{g.Cur, g.Lst, cur}
		}
		if g.Cur > g.Max.Val {
			g.Max = GapsTriad{g.Cur, g.Lst, cur}
		}
	} else {
		g.Mma = true
		g.Min = GapsTriad{g.Cur, g.Lst, cur}
		g.Max = GapsTriad{g.Cur, g.Lst, cur}
	}
	// Update the last time
	g.Lst = cur
	// Reset the current gap
	g.Cur = 1
}

func (g *Gaps) UpdateAvg() {
	// Update the average gap
	g.Avg = float64(g.Sum) / float64(len(g.All))
}
