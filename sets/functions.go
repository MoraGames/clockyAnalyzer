package sets

import "math"

func Equal(h1, h2, m1, m2 int) bool {
	return h1 == h2 && h2 == m1 && m1 == m2
}

func ShortEqual(_, h2, m1, m2 int) bool {
	return h2 == m1 && m1 == m2
}

func Repeat(h1, h2, m1, m2 int) bool {
	return h1 == m1 && h2 == m2 && h1 != h2
}

func Mirror(h1, h2, m1, m2 int) bool {
	return h1 == m2 && h2 == m1 && h1 != h2
}

func Rise(h1, h2, m1, m2 int) bool {
	return h2 == h1+1 && m1 == h2+1 && m2 == m1+1
}

func ShortRise(_, h2, m1, m2 int) bool {
	return m1 == h2+1 && m2 == m1+1
}

func ShortFall(_, h2, m1, m2 int) bool {
	return m1 == h2-1 && m2 == m1-1
}

func RapidRise(h1, h2, m1, m2 int) bool {
	return h2 == h1+2 && m1 == h2+2 && m2 == m1+2
}

func ShortRapidRise(_, h2, m1, m2 int) bool {
	return m1 == h2+2 && m2 == m1+2
}
func ShortRapidFall(_, h2, m1, m2 int) bool {
	return m1 == h2-2 && m2 == m1-2
}

func Double(h1, h2, m1, m2 int) bool {
	return 2*((h1*10)+h2) == (m1*10)+m2
}

func ShortTriple(_, h2, m1, m2 int) bool {
	return 3*h2 == (m1*10)+m2
}

func SymmetricalDifference(h1, h2, m1, m2 int) bool {
	diff1, diff2 := math.Abs(float64(h1-m2)), math.Abs(float64(h2-m1))
	return diff1 == diff2
}

func SymmetricalSum(h1, h2, m1, m2 int) bool {
	sum1, sum2 := h1+m2, h2+m1
	return sum1 == sum2
}

func PerfectSquare(h1, h2, m1, m2 int) bool {
	total := (h1 * 1000) + (h2 * 100) + (m1 * 10) + (m2)
	sqrt := int(math.Sqrt(float64(total)))
	return (sqrt * sqrt) == total
}

func SumAndDifference(h1, h2, m1, m2 int) bool {
	h, m := (h1*10)+h2, (m1*10)+m2
	return (h + m) == (h - m)
}
