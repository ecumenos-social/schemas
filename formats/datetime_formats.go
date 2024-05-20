package formats

import "time"

var DateTimeFormat = "2006-01-02T15:04:05.999Z"

func FormatDateTime(v time.Time) string {
	return v.UTC().Format(DateTimeFormat)
}
