package main

import (
	"fmt"

	"github.com/MoraGames/clockyAnalyzer/groups"
	"github.com/MoraGames/clockyAnalyzer/prints"
	"github.com/MoraGames/clockyAnalyzer/sets"
	"github.com/MoraGames/clockyAnalyzer/time"
)

const (
	TotalTimes = 24 * 60
)

var (
	AllGroups = map[string]*groups.Group{
		"Any":          groups.NewGroup("Any"),
		"Approved":     groups.NewGroup("Approved"),
		"Under-Review": groups.NewGroup("Under-Review"),
		"Rejected":     groups.NewGroup("Rejected"),
	}
	AllSets = []*sets.Set{
		// Approved sets
		sets.NewSet("Equal", "Approved", sets.Equal, time.NewTime(22, 22)),
		sets.NewSet("ShortEqual", "Approved", sets.ShortEqual, time.NewTime(23, 33)),
		sets.NewSet("Repeat", "Approved", sets.Repeat, time.NewTime(23, 23)),
		sets.NewSet("Mirror", "Approved", sets.Mirror, time.NewTime(23, 32)),
		sets.NewSet("Rise", "Approved", sets.Rise, time.NewTime(23, 45)),
		sets.NewSet("ShortRise", "Approved", sets.ShortRise, time.NewTime(23, 45)),
		sets.NewSet("ShortFall", "Approved", sets.ShortFall, time.NewTime(23, 21)),
		sets.NewSet("RapidRise", "Approved", sets.RapidRise, time.NewTime(13, 57)),
		sets.NewSet("ShortRapidRise", "Approved", sets.ShortRapidRise, time.NewTime(23, 57)),
		sets.NewSet("ShortRapidFall", "Approved", sets.ShortRapidFall, time.NewTime(17, 53)),
		sets.NewSet("Double", "Approved", sets.Double, time.NewTime(23, 46)),
		sets.NewSet("ShortTriple", "Approved", sets.ShortTriple, time.NewTime(23, 9)),
		sets.NewSet("PerfectSquare", "Approved", sets.PerfectSquare, time.NewTime(23, 04)),

		// In-Review sets
		sets.NewSet("SumAndDifference", "Under-Review", sets.SumAndDifference, time.NewTime(23, 00)),

		// Rejected sets
		sets.NewSet("SymmetricalDifference", "Rejected", sets.SymmetricalDifference, time.NewTime(23, 54)),
		sets.NewSet("SymmetricalSum", "Rejected", sets.SymmetricalSum, time.NewTime(23, 56)),
	}
)

func main() {
	// Find and update the last set of the day (the set with the smallest gap) for each group
	for _, set := range AllSets {
		if AllGroups["Any"].LastSet == nil || set.Gaps.Cur < AllGroups["Any"].LastSet.Gaps.Cur {
			AllGroups["Any"].SetLastSet(set)
		}
		if AllGroups[set.Label].LastSet == nil || set.Gaps.Cur < AllGroups[set.Label].LastSet.Gaps.Cur {
			AllGroups[set.Label].SetLastSet(set)
		}
	}
	// Update the number of sets for each group
	for _, group := range AllGroups {
		setNums := 0
		for _, set := range AllSets {
			if group.Name == "Any" || set.Label == group.Name {
				fmt.Printf("[DEBUG] SetNums++ : Set %v (%v) for group %v\n", set.Name, set.Label, group.Name)
				setNums++
			}
		}
		fmt.Printf("[DEBUG] SetNums = %v for group %v\n", setNums, group.Name)
		group.SetSetNums(setNums)
	}

	for hour := 0; hour < 24; hour++ {
		for minute := 0; minute < 60; minute++ {
			currentTime := time.NewTime(hour, minute)

			// Loop through all sets
			for _, set := range AllSets {
				// Check if the time is valid for the set
				verification := set.Verify(currentTime.SplitTime())
				if verification {
					// Increase the number of times the set was verified
					set.Nums++

					// Update the gaps of the set
					set.Gaps.UpdateGaps(currentTime)

					// Enable the flag for time verified by at least one set.
					AllGroups["Any"].AtLeastOne = true
					AllGroups[set.Label].AtLeastOne = true
				} else {
					set.Gaps.Cur++
				}
			}

			// Loop through all groups
			for _, group := range AllGroups {
				// Check if the time was verified by at least one set (for this general set)
				if group.AtLeastOne {
					// Increase the number of times the general set was verified
					group.EventNums++

					// Update the gaps of the general set
					group.Gaps.UpdateGaps(currentTime)

					// Reset the flag for time verified by at least one set
					group.AtLeastOne = false
				} else {
					group.Gaps.Cur++
				}
			}

			prints.PrintTime(AllSets, currentTime, TotalTimes)
		}
	}

	// Update the gaps averages for all sets
	for _, set := range AllSets {
		set.Gaps.UpdateAvg()
	}
	// Update the gaps averages for all groups
	for _, group := range AllGroups {
		group.Gaps.UpdateAvg()
	}

	// Print the final stats for all sets
	prints.PrintFinal(AllSets, AllGroups, TotalTimes)
}
