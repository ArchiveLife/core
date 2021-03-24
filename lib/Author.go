package lib

type Author struct {
	ID            string
	Age           *int
	FullName      string
	Address       *string
	Mobile        *string
	ExtAttributes []*Attribtue
}
