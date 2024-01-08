// Package tu provides time utility functions to easily get time before or after of specific date time
package tu

import (
	"time"
)

var (
	// Yesterday get yesterday datetime of given datetime
	Yesterday TimeFunc = func(t time.Time) time.Time {
		return t.AddDate(0, 0, -1)
	}

	// Tomorrow get tomorrow datetime of given datetime
	Tomorrow TimeFunc = func(t time.Time) time.Time {
		return t.AddDate(0, 0, 1)
	}

	// BeginningOfDay get beginning of day of given datetime
	BeginningOfDay TimeFunc = func(t time.Time) time.Time {
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	}

	// EndOfDay get end of day of given datetime
	EndOfDay TimeFunc = func(t time.Time) time.Time {
		return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
	}

	// BeginningOfWeek get beginning of week of given datetime
	BeginningOfWeek TimeFunc = func(t time.Time) time.Time {
		tt := BeginningOfDay(t).AddDate(0, 0, int(t.Weekday())*-1)
		return time.Date(t.Year(), t.Month(), tt.Day(), 0, 0, 0, 0, t.Location())
	}

	// EndOfWeek get end of week of given datetime
	EndOfWeek TimeFunc = func(t time.Time) time.Time {
		tt := EndOfDay(t).AddDate(0, 0, 6-int(t.Weekday()))
		return time.Date(t.Year(), t.Month(), tt.Day(), 23, 59, 59, 999999999, t.Location())
	}

	// BeginningOfMonth get beginning of month of given datetime
	BeginningOfMonth TimeFunc = func(t time.Time) time.Time {
		return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	}

	// EndOfMonth get end of month of given datetime
	EndOfMonth TimeFunc = func(t time.Time) time.Time {
		tt := BeginningOfMonth(t.AddDate(0, 1, 0)).AddDate(0, 0, -1)
		return time.Date(t.Year(), t.Month(), tt.Day(), 23, 59, 59, 999999999, t.Location())
	}

	// BeginningOfYear get beginning of year of given datetime
	BeginningOfYear TimeFunc = func(t time.Time) time.Time {
		return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	}

	// EndOfYear get end of year of given datetime
	EndOfYear TimeFunc = func(t time.Time) time.Time {
		return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, t.Location())
	}
)

// TimeFunc types of time utility function provides methods to compose utility functions
type TimeFunc func(time.Time) time.Time

// Yesterday composes Yesterday function with f TimeFunc function
func (f TimeFunc) Yesterday() TimeFunc {
	return func(t time.Time) time.Time {
		return Yesterday(f(t))
	}
}

// Tomorrow composes Tomorrow function with f TimeFunc function
func (f TimeFunc) Tomorrow() TimeFunc {
	return func(t time.Time) time.Time {
		return Tomorrow(f(t))
	}
}

// BeginningOfDay composes BeginningOfDay function with f TimeFunc function
func (f TimeFunc) BeginningOfDay() TimeFunc {
	return func(t time.Time) time.Time {
		return BeginningOfDay(f(t))
	}
}

// EndOfDay composes EndOfDay function with f TimeFunc function
func (f TimeFunc) EndOfDay() TimeFunc {
	return func(t time.Time) time.Time {
		return EndOfDay(f(t))
	}
}

// BeginningOfWeek composes BeginningOfWeek function with f TimeFunc function
func (f TimeFunc) BeginningOfWeek() TimeFunc {
	return func(t time.Time) time.Time {
		return BeginningOfWeek(f(t))
	}
}

// EndOfWeek composes EndOfWeek function with f TimeFunc function
func (f TimeFunc) EndOfWeek() TimeFunc {
	return func(t time.Time) time.Time {
		return EndOfWeek(f(t))
	}
}

// BeginningOfMonth composes BeginningOfMonth function with f TimeFunc function
func (f TimeFunc) BeginningOfMonth() TimeFunc {
	return func(t time.Time) time.Time {
		return BeginningOfMonth(f(t))
	}
}

// EndOfMonth composes EndOfMonth function with f TimeFunc function
func (f TimeFunc) EndOfMonth() TimeFunc {
	return func(t time.Time) time.Time {
		return EndOfMonth(f(t))
	}
}

// BeginningOfYear composes BeginningOfYear function with f TimeFunc function
func (f TimeFunc) BeginningOfYear() TimeFunc {
	return func(t time.Time) time.Time {
		return BeginningOfYear(f(t))
	}
}

// EndOfYear composes EndOfYear function with f TimeFunc function
func (f TimeFunc) EndOfYear() TimeFunc {
	return func(t time.Time) time.Time {
		return EndOfYear(f(t))
	}
}

// Of apply f TimeFunc function with given datetime
func (f TimeFunc) Of(t time.Time) time.Time {
	return f(t)
}

// Num provides method to easily get time before or after of specific date time with N number time duration
type Num int

// SecondsAgo provides TimeFunc function to get n seconds ago
func (n Num) SecondsAgo() TimeFunc {
	return func(t time.Time) time.Time {
		return t.Add(time.Duration(-int(n)) * time.Second)
	}
}

// SecondsLater provides TimeFunc function to get n seconds later
func (n Num) SecondsLater() TimeFunc {
	return func(t time.Time) time.Time {
		return t.Add(time.Duration(int(n)) * time.Second)
	}
}

// MinutesAgo provides TimeFunc function to get n minutes ago
func (n Num) MinutesAgo() TimeFunc {
	return func(t time.Time) time.Time {
		return t.Add(time.Duration(-int(n)) * time.Minute)
	}
}

// MinutesLater provides TimeFunc function to get n minutes later
func (n Num) MinutesLater() TimeFunc {
	return func(t time.Time) time.Time {
		return t.Add(time.Duration(int(n)) * time.Minute)
	}
}

// HoursAgo provides TimeFunc function to get n hours ago
func (n Num) HoursAgo() TimeFunc {
	return func(t time.Time) time.Time {
		return t.Add(time.Duration(-int(n)) * time.Hour)
	}
}

// HoursLater provides TimeFunc function to get n hours later
func (n Num) HoursLater() TimeFunc {
	return func(t time.Time) time.Time {
		return t.Add(time.Duration(int(n)) * time.Hour)
	}
}

// DaysAgo provides TimeFunc function to get n days ago
func (n Num) DaysAgo() TimeFunc {
	return func(t time.Time) time.Time {
		return t.AddDate(0, 0, -int(n))
	}
}

// DaysLater provides TimeFunc function to get n days later
func (n Num) DaysLater() TimeFunc {
	return func(t time.Time) time.Time {
		return t.AddDate(0, 0, int(n))
	}
}

// MonthsAgo provides TimeFunc function to get n months ago
func (n Num) MonthsAgo() TimeFunc {
	return func(t time.Time) time.Time {
		return t.AddDate(0, -int(n), 0)
	}
}

// MonthsLater provides TimeFunc function to get n months later
func (n Num) MonthsLater() TimeFunc {
	return func(t time.Time) time.Time {
		return t.AddDate(0, int(n), 0)
	}
}

// YearsAgo provides TimeFunc function to get n years ago
func (n Num) YearsAgo() TimeFunc {
	return func(t time.Time) time.Time {
		return t.AddDate(-int(n), 0, 0)
	}
}

// YearsLater provides TimeFunc function to get n years later
func (n Num) YearsLater() TimeFunc {
	return func(t time.Time) time.Time {
		return t.AddDate(int(n), 0, 0)
	}
}

// N alias of Num
type N = Num
