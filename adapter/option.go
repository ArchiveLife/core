package adapter

type OptionValueType string

const (
	String OptionValueType = "String"
)

type Option struct {
	Order     int
	Name      string
	Optional  bool
	ValueType OptionValueType
}

type OptionValue struct {
	Option
	Value interface{}
}
