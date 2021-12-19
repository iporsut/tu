## tu - Time Util package

tu is a time util for golang

[![go report card](https://goreportcard.com/badge/github.com/iporsut/tu "go report card")](https://goreportcard.com/report/github.com/iporsut/tu)
[![Go Reference](https://pkg.go.dev/badge/github.com/iporsut/tu.svg)](https://pkg.go.dev/github.com/iporsut/tu)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Install

```
go get github.com/iporsut/tu
```

## Usage

Get datetime before or after specific datetime

```go
import "github.com/iporsut/tu"

now := time.Date(2021, 12, 19, 20, 36, 30, 123456789, time.UTC)

tu.Tomorrow.Of(now)                           // 2021-12-20 20:36:30.123456789
tu.Yesterday.Of(now)                          // 2021-12-18 20:36:30.123456789
tu.BeginningOfDay.Of(now)                     // 2021-12-19 00:00:00
tu.EndOfDay.Of(now)                           // 2021-12-19 23:59:59.999999999
tu.BeginningOfMonth.Of(now)                   // 2021-12-01 00:00:00
tu.EndOfMonth.Of(now)                         // 2021-12-31 23:59:59.999999999
tu.BeginningOfYear.Of(now)                    // 2021-01-01 00:00:00
tu.EndOfYear.Of(now)                          // 2021-12-31 23:59:59.999999999

// Get datetime from N value duration
tu.N(2).SecondsAgo().Of(now)                  // 2021-12-19 20:36:28.123456789
tu.N(2).SecondsLater().Of(now)                // 2021-12-19 20:36:32.123456789
tu.N(2).MinutesAgo().Of(now)                  // 2021-12-19 20:34:30.123456789
tu.N(2).MinutesLater().Of(now)                // 2021-12-19 20:38:30.123456789
tu.N(2).HoursAgo().Of(now)                    // 2021-12-19 18:36:30.123456789
tu.N(2).HoursLater().Of(now)                  // 2021-12-19 22:36:30.123456789
tu.N(2).DaysAgo().Of(now)                     // 2021-12-17 20:36:30.123456789
tu.N(2).DaysLater().Of(now)                   // 2021-12-21 20:36:30.123456789
tu.N(2).MonthsAgo().Of(now)                   // 2019-10-19 20:36:30.123456789
tu.N(2).MonthsLater().Of(now)                 // 2022-02-19 20:36:30.123456789
tu.N(2).YearsAgo().Of(now)                    // 2019-12-19 20:36:30.123456789
tu.N(2).YearsLater().Of(now)                  // 2023-12-19 20:36:30.123456789

// Compose Time Function
tu.Tomorrow.Tomorrow().Of(now)                // 2021-12-21 20:36:30.123456789
tu.Yesterday.Yesterday().Of(now)              // 2021-12-17 20:36:30.123456789
tu.N(2).DaysAgo().BeginningOfDay().Of(now)    // 2021-12-17 00:00:00
```

# Author

**Weerasak Chongnguluam**

* <http://github.com/iporsut>
* <singpor@gmail.com>
* <http://twitter.com/iporsut>

## License

Released under the [MIT License](http://www.opensource.org/licenses/MIT).
