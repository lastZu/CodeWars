package main

import (
	"fmt"
	"strings"
)

/*
For seconds = 3662, your function should return
    "1 hour, 1 minute and 2 seconds"
*/

type duration struct {
	name    string
	count   int
	seconds int
}

func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	var order = [5]duration{
		{
			name:    "second",
			count:   0,
			seconds: 1,
		},
		{
			name:    "minute",
			count:   0,
			seconds: 60,
		},
		{
			name:    "hour",
			count:   0,
			seconds: 60 * 60,
		},
		{
			name:    "day",
			count:   0,
			seconds: 60 * 60 * 24,
		},
		{
			name:    "year",
			count:   0,
			seconds: 60 * 60 * 24 * 365,
		},
	}

	temp := int(seconds)
	for i := 4; i >= 0; i-- {
		order[i].count = temp / order[i].seconds
		temp = temp % order[i].seconds
	}

	var res []string
	for i := 4; i >= 0; i-- {
		count := order[i].count
		name := order[i].name
		if count > 1 {
			name += "s"
		}
		if count != 0 {
			res = append(res, fmt.Sprintf("%d %s", count, name))
		}
	}
	answer := strings.Join(res, ", ")
	answer = strings.Replace(answer, ",", " and", -1)
	if len(res)-2 > 0 {
		answer = strings.Replace(answer, " and", ",", len(res)-2)
	}

	return answer
}
