package main

import (
	"fmt"
	"other/bp/cuwu_setstest/generalSets"
	"other/bp/cuwu_setstest/sets"
	"other/bp/cuwu_setstest/time"
)

var (
	PrintEveryTime  bool = false
	PrintFinalStats bool = true
	PrintAllGaps    bool = false
)

func main() {
	generalSets := map[string]generalSets.GeneralSet{
		// General sets
		"Any":          generalSets.NewGeneralSet_Any("Any", nil),
		"Approved":     generalSets.NewGeneralSet_Approved("Approved", nil),
		"Under-Review": generalSets.NewGeneralSet_UnderReview("Under-Review", nil),
	}
	allSets := []sets.Set{
		// Approved sets
		sets.NewSet_Equal("Equal", "Approved"),
		sets.NewSet_ShortEqual("ShortEqual", "Approved"),
		sets.NewSet_Repeat("Repeat", "Approved"),
		sets.NewSet_Mirror("Mirror", "Approved"),
		sets.NewSet_Rise("Rise", "Approved"),
		sets.NewSet_ShortRise("ShortRise", "Approved"),
		sets.NewSet_ShortFall("ShortFall", "Approved"),
		sets.NewSet_RapidRise("RapidRise", "Approved"),
		sets.NewSet_ShortRapidRise("ShortRapidRise", "Approved"),
		sets.NewSet_ShortRapidFall("ShortRapidFall", "Approved"),
		sets.NewSet_Double("Double", "Approved"),

		// In-Review sets
		sets.NewSet_ShortTriple("ShortTriple", "Under-Review"),
		sets.NewSet_SymmetricalDifference("SymmetricalDifference", "Under-Review"),
		sets.NewSet_SymmetricalSum("SymmetricalSum", "Under-Review"),
		sets.NewSet_PerfectSquare("PerfectSquare", "Under-Review"),
		sets.NewSet_SumAndDifference("SumAndDifference", "Under-Review"),

		// Rejected sets
		// -- none --
	}

	// Find and update the last set of the day (the set with the smallest gap) for each general set
	for _, set := range allSets {
		if generalSets["Any"].GetLastSet() == nil || set.GetGaps().Cur < generalSets["Any"].GetLastSet().GetGaps().Cur {
			generalSets["Any"].SetLastSet(set)
		}
		if generalSets[set.GetLabel()].GetLastSet() == nil || set.GetGaps().Cur < generalSets[set.GetLabel()].GetLastSet().GetGaps().Cur {
			generalSets[set.GetLabel()].SetLastSet(set)
		}
	}

	// lastSet := allSets[0]
	// for _, set := range allSets[1:] {
	// 	if set.GetGaps().Cur < lastSet.GetGaps().Cur {
	// 		lastSet = set
	// 	}
	// }

	totalTimes := 24 * 60

	for hour := 0; hour < 24; hour++ {
		for minute := 0; minute < 60; minute++ {
			currentTime := time.NewTime(hour, minute)

			// Loop through all sets
			for _, set := range allSets {
				// Check if the time is valid for the set
				verification := set.Verify(currentTime.SplitTime())
				if verification {
					// Increase the number of times the set was verified
					set.IncreaseNums()

					// Update the gaps of the set
					set.GetGaps().UpdateGaps(currentTime)

					// Enable the flag for time verified by at least one set.
					generalSets["Any"].SetAtLeastOne(true)
					generalSets[set.GetLabel()].SetAtLeastOne(true)
				} else {
					set.GetGaps().IncreaseCur(currentTime)
				}
			}

			// Loop through all general sets
			for _, generalSet := range generalSets {
				// Check if the time was verified by at least one set (for this general set)
				if generalSet.GetAtLeastOne() {
					// Increase the number of times the general set was verified
					generalSet.IncreaseNums()

					// Update the gaps of the general set
					generalSet.GetGaps().UpdateGaps(currentTime)

					// Reset the flag for time verified by at least one set
					generalSet.SetAtLeastOne(false)
				} else {
					generalSet.GetGaps().IncreaseCur(currentTime)
				}
			}

			if PrintEveryTime {
				fmt.Printf(">>> TIME = %v\n", currentTime.Stringify())
				for _, set := range allSets {
					fmt.Printf(" | %v = %v/%v (Gap = %v)\n", set.GetName(), set.GetNums(), totalTimes, set.GetGaps().Cur)
				}
				fmt.Printf("\n")
			}
		}
	}

	// Update the gaps averages for all sets
	for _, set := range allSets {
		set.GetGaps().UpdateAvg()
	}
	for _, generalSet := range generalSets {
		generalSet.GetGaps().UpdateAvg()
	}

	// Print the final stats for all sets
	if PrintFinalStats {
		fmt.Printf(">>> FINAL STATS:\n")
		for _, set := range allSets {
			fmt.Printf(" |- %v = %v/%v\n", set.GetName(), set.GetNums(), totalTimes)
			if PrintAllGaps {
				fmt.Printf(" |   Stats:\n |    |- All: %v\n |    |- Min: %v (%v - %v)\n |    |- Max: %v (%v - %v)\n |    |- Avg: %.3f\n\n", set.GetGaps().All, set.GetGaps().Min.Val, set.GetGaps().Min.Prv.Stringify(), set.GetGaps().Min.Cur.Stringify(), set.GetGaps().Max.Val, set.GetGaps().Max.Prv.Stringify(), set.GetGaps().Max.Cur.Stringify(), set.GetGaps().Avg)
			} else {
				fmt.Printf(" |   Stats:\n |    |- Min: %v (%v - %v)\n |    |- Max: %v (%v - %v)\n |    |- Avg: %.3f\n\n", set.GetGaps().Min.Val, set.GetGaps().Min.Prv.Stringify(), set.GetGaps().Min.Cur.Stringify(), set.GetGaps().Max.Val, set.GetGaps().Max.Prv.Stringify(), set.GetGaps().Max.Cur.Stringify(), set.GetGaps().Avg)
			}
		}
		for _, generalSet := range generalSets {
			fmt.Printf(" |- %v = %v/%v\n", generalSet.GetName(), generalSet.GetNums(), totalTimes)
			if PrintAllGaps {
				fmt.Printf(" |   Stats:\n |    |- All: %v\n |    |- Min: %v (%v - %v)\n |    |- Max: %v (%v - %v)\n |    |- Avg: %.3f\n\n", generalSet.GetGaps().All, generalSet.GetGaps().Min.Val, generalSet.GetGaps().Min.Prv.Stringify(), generalSet.GetGaps().Min.Cur.Stringify(), generalSet.GetGaps().Max.Val, generalSet.GetGaps().Max.Prv.Stringify(), generalSet.GetGaps().Max.Cur.Stringify(), generalSet.GetGaps().Avg)
			} else {
				fmt.Printf(" |   Stats:\n |    |- Min: %v (%v - %v)\n |    |- Max: %v (%v - %v)\n |    |- Avg: %.3f\n\n", generalSet.GetGaps().Min.Val, generalSet.GetGaps().Min.Prv.Stringify(), generalSet.GetGaps().Min.Cur.Stringify(), generalSet.GetGaps().Max.Val, generalSet.GetGaps().Max.Prv.Stringify(), generalSet.GetGaps().Max.Cur.Stringify(), generalSet.GetGaps().Avg)
			}
		}
	}
}
