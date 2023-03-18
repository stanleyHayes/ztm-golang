package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour, Minute, Second uint
}

type TimeParseError struct {
	Message string
	Input   string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.Message, t.Input)
}

func Parse(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid number of time components", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{
				fmt.Sprintf("Error parsing hour: %v", err),
				input,
			}
		}

		minute, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{
				fmt.Sprintf("Error parsing minute: %v", err),
				input,
			}
		}

		second, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParseError{
				fmt.Sprintf("Error parsing second: %v", err),
				input,
			}
		}

		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{
				"hour out of range: 0 <= hour <= 23",
				fmt.Sprintf("%v", hour),
			}
		}
		if minute > 59 || minute < 0 {
			return Time{}, &TimeParseError{
				"minute out of range: 0 <= minute <= 59",
				fmt.Sprintf("%v", minute),
			}
		}
		if second > 59 || second < 0 {
			return Time{}, &TimeParseError{
				"second out of range: 0 <= second <= 59",
				fmt.Sprintf("%v", second),
			}
		}
		return Time{Hour: uint(hour), Minute: uint(minute), Second: uint(second)}, nil
	}
}

func main() {

}
