package tu_test

import (
	"testing"
	"time"

	"github.com/iporsut/tu"
)

func TestTimeUtil(t *testing.T) {
	currentDate := time.Date(2021, 12, 19, 20, 36, 30, 123456789, time.UTC)
	testcases := []struct {
		now      time.Time
		timeFunc tu.TimeFunc
		expected time.Time
	}{
		{currentDate, tu.Tomorrow.Of, time.Date(2021, 12, 20, 20, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.Yesterday.Of, time.Date(2021, 12, 18, 20, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.BeginningOfDay.Of, time.Date(2021, 12, 19, 0, 0, 0, 0, time.UTC)},
		{currentDate, tu.EndOfDay.Of, time.Date(2021, 12, 19, 23, 59, 59, 999999999, time.UTC)},
		{currentDate, tu.BeginningOfMonth.Of, time.Date(2021, 12, 1, 0, 0, 0, 0, time.UTC)},
		{currentDate, tu.EndOfMonth.Of, time.Date(2021, 12, 31, 23, 59, 59, 999999999, time.UTC)},
		{currentDate, tu.BeginningOfYear.Of, time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		{currentDate, tu.EndOfYear.Of, time.Date(2021, 12, 31, 23, 59, 59, 999999999, time.UTC)},
		{currentDate, tu.Tomorrow.Tomorrow().Of, time.Date(2021, 12, 21, 20, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.Yesterday.Yesterday().Of, time.Date(2021, 12, 17, 20, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.Tomorrow.Yesterday().Of, time.Date(2021, 12, 19, 20, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.Yesterday.Tomorrow().Of, time.Date(2021, 12, 19, 20, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.N(2).SecondsAgo().Of, time.Date(2021, 12, 19, 20, 36, 28, 123456789, time.UTC)},
		{currentDate, tu.N(2).SecondsLater().Of, time.Date(2021, 12, 19, 20, 36, 32, 123456789, time.UTC)},
		{currentDate, tu.N(2).MinutesAgo().Of, time.Date(2021, 12, 19, 20, 34, 30, 123456789, time.UTC)},
		{currentDate, tu.N(2).MinutesLater().Of, time.Date(2021, 12, 19, 20, 38, 30, 123456789, time.UTC)},
		{currentDate, tu.N(2).HoursAgo().Of, time.Date(2021, 12, 19, 18, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.N(2).HoursLater().Of, time.Date(2021, 12, 19, 22, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.N(2).DaysAgo().Of, time.Date(2021, 12, 17, 20, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.N(2).DaysLater().Of, time.Date(2021, 12, 21, 20, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.N(2).MonthsAgo().Of, time.Date(2021, 10, 19, 20, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.N(2).MonthsLater().Of, time.Date(2022, 2, 19, 20, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.N(2).YearsAgo().Of, time.Date(2019, 12, 19, 20, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.N(2).YearsLater().Of, time.Date(2023, 12, 19, 20, 36, 30, 123456789, time.UTC)},
		{currentDate, tu.N(2).DaysAgo().BeginningOfDay().Of, time.Date(2021, 12, 17, 0, 0, 0, 0, time.UTC)},
		{time.Date(2024, 1, 30, 0, 0, 0, 0, time.UTC), tu.EndOfMonth.Of, time.Date(2024, 1, 31, 23, 59, 59, 999999999, time.UTC)},
		{time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC), tu.EndOfMonth.Of, time.Date(2024, 3, 31, 23, 59, 59, 999999999, time.UTC)},
	}
	for _, tc := range testcases {
		if r := tc.timeFunc(tc.now); r != tc.expected {
			t.Errorf("expect %v but got %v", tc.expected, r)
		}
	}
}
