package lib

import "time"

type Article struct {
	ID            string
	Title         *string
	Author        *Author
	CoAuthrors    []*Author
	PublishDate   *time.Time
	Content       *string
	Tags          []string
	References    []*Reference
	Type          string
	ExtAttributes []*Attribtue
}
