package time

import "fmt"

type (
	Time struct {
		Hour   int
		Minute int
	}
)

func NewTime(hour, minute int) Time {
	return Time{
		hour,
		minute,
	}
}

func (t *Time) SplitTime() (int, int, int, int) {
	return t.Hour / 10, t.Hour % 10, t.Minute / 10, t.Minute % 10
}

func (t *Time) Stringify() string {
	return fmt.Sprintf("%d%d:%d%d", t.Hour/10, t.Hour%10, t.Minute/10, t.Minute%10)
}

func (t *Time) UntilMidnight() int {
	return (24-t.Hour)*60 - t.Minute
}
