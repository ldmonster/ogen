package validate

import (
	"strings"
)

// Error represents validation error.
type Error struct {
	Fields []FieldError
}

func (e *Error) Error() string {
	var b strings.Builder
	b.WriteString("invalid:")
	for i, f := range e.Fields {
		if i != 0 {
			b.WriteRune(',')
		}
		b.WriteRune(' ')
		b.WriteString(f.Name)
		b.WriteString(" (")
		b.WriteString(f.Error.Error())
		b.WriteString(")")
	}

	return b.String()
}

// FieldError is failed validation on field.
type FieldError struct {
	Name  string
	Error error
}