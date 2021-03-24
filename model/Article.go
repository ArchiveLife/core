package model

import "time"

type Article struct {
	ID            ID
	Title         *string
	Author        *Author
	CoAuthrors    []*Author
	PublishDate   *time.Time
	Content       *string
	Tags          []string
	Type          string
	Medias        []*Media
	References    []*Reference
	ExtAttributes []*Attribtue
}
