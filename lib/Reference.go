package lib

type ReferenceType int

const (
	Author ReferenceType = 1
)

type Reference struct {
	Type        ReferenceType
	ReferenceId string
}
