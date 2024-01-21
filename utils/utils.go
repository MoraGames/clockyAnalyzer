package utils

import "fmt"

func DecimalToSexagesimal(minutes float64) (int, int) {
	seconds := int(minutes * 60)
	return seconds / 60, seconds % 60
}

func FloatWidth(f float64, precision int) int {
	return len(fmt.Sprintf("%.0f\n", f)) + precision
}
