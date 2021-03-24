package lib

import "time"

type Article struct {
	Title      *string
	Author     *string
	CoAuthrors []string
	Date       *time.Time
	Content    *string
	Tags       []string
	References []*Reference
}
