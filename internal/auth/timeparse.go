package auth

import (
	"errors"
	"time"
)

func timeParse(timeout string) (time.Duration, error) {
	var err error
	var t time.Duration

	length := len(timeout)
	if length > 1 {
		suffix := timeout[length-1]

		switch string(suffix) {
		case "h":
			t, err = time.ParseDuration(timeout)
		case "m":
			t, err = time.ParseDuration(timeout)
		case "d": //day
			t, err = time.ParseDuration(timeout[:length-1] + "h")
			t = 24 * t
		case "M": // month
			t, err = time.ParseDuration(timeout[:length-1] + "h")
			t = 730 * t
		default:
			err = errors.New("ERROR: TimeParse: wrong time format")
		}
	} else {
		err = errors.New("ERROR: TimeParse: wrong time format")
	}

	return t, err
}

// ToTime - converts string (example: 3d) to time.Duration
func ToTime(s string) time.Duration {

	t, err := timeParse(s)
	if err != nil {
		t, _ = timeParse("7d")
	}

	return t
}
