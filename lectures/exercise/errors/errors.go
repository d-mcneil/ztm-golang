//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"strconv"
	"strings"
)

type Time struct {
	hours, minutes, seconds int
}

type TimeConvertError struct {
	isErrorHours     bool
	isErrorMinutes   bool
	isErrorSeconds   bool
	isInvalidHours   bool
	isInvalidMinutes bool
	isInvalidSeconds bool
}

func (e *TimeConvertError) Error() string {
	if !e.isErrorHours && !e.isErrorMinutes && !e.isErrorSeconds {
		return "unable to parse date string"
	}
	var str string
	if e.isErrorHours || e.isErrorMinutes || e.isErrorSeconds {
		str += "error parsing following time fields: "
		if e.isErrorHours {
			str += "hours"
		}
		if e.isErrorMinutes {
			if e.isErrorHours {
				str += ", minutes"
			} else {
				str += "minutes"
			}
		}
		if e.isErrorSeconds {
			if e.isErrorHours || e.isErrorMinutes {
				str += ", seconds"
			} else {
				str += "seconds"
			}
		}
	}
	if (e.isErrorHours || e.isErrorMinutes || e.isErrorSeconds) &&
		(e.isInvalidHours || e.isInvalidMinutes || e.isInvalidSeconds) {
		str += "\n"
	}
	if e.isInvalidHours || e.isInvalidMinutes || e.isInvalidSeconds {
		str += "invalid amount following time fields: "
		if e.isInvalidHours {
			str += "hours"
		}
		if e.isInvalidMinutes {
			if e.isInvalidHours {
				str += ", minutes"
			} else {
				str += "minutes"
			}
		}
		if e.isInvalidSeconds {
			if e.isInvalidHours || e.isInvalidMinutes {
				str += ", seconds"
			} else {
				str += "seconds"
			}
		}
	}
	return str
}

func ParseDate(dateString string) (Time, error) {
	var time Time

	dateParts := strings.Split(strings.Replace(dateString, " ", "", -1), ":")
	if len(dateParts) != 3 {
		return time, &TimeConvertError{}
	}

	var isErrorHours, isErrorMinutes, isErrorSeconds, isInvalidHours, isInvalidMinutes, isInvalidSeconds bool

	hours, err := strconv.Atoi(dateParts[0])
	if err != nil {
		isErrorHours = true
	} else if hours > 23 || hours < 0 {
		isInvalidHours = true
	}

	minutes, err := strconv.Atoi(dateParts[1])
	if err != nil {
		isErrorMinutes = true
	} else if minutes > 59 || minutes < 0 {
		isInvalidMinutes = true
	}

	seconds, err := strconv.Atoi(dateParts[2])
	if err != nil {
		isErrorSeconds = true
	} else if seconds > 59 || seconds < 0 {
		isInvalidSeconds = true
	}

	if isErrorHours || isErrorMinutes || isErrorSeconds || isInvalidHours || isInvalidMinutes || isInvalidSeconds {
		return time, &TimeConvertError{isErrorHours, isErrorMinutes, isErrorSeconds, isInvalidHours, isInvalidMinutes, isInvalidSeconds}
	}

	time.hours = hours
	time.minutes = minutes
	time.seconds = seconds

	return time, nil
}
