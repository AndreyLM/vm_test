package models

import (
	"fmt"
	"strings"
	"time"
)

// Tuesday 20 Oct 1990

const ctLayout = "Monday 2 Jan 2006"

var nilTime = (time.Time{}).UnixNano()

// CustomTime - custom time
type CustomTime struct {
	time.Time
}

// UnmarshalJSON - unmarshalling json
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

// MarshalJSON - marshalling json
func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}
