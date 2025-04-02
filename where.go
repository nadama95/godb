package godb

import "fmt"

type Comparator int

type Where struct {
	field      string
	comparator Comparator
	value      any
}

func (w Where) String() string {
	return fmt.Sprintf("%s %s '%v'", w.field, w.comparator.String(), w.value)
}

const (
	EqualTo Comparator = iota
	GreaterThan
	LessThan
	GreaterOrEqual
	LessOrEqual
	NotEqual
	Like
)

func (c Comparator) String() string {
	switch c {
	case EqualTo:
		return "="
	case GreaterThan:
		return ">"
	case LessThan:
		return "<"
	case GreaterOrEqual:
		return ">="
	case LessOrEqual:
		return "<="
	case NotEqual:
		return "<>"
	case Like:
		return "LIKE"
	default:
		return ""
	}
}
