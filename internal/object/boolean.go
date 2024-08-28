package object

import "fmt"

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) HashKey() HashKey {
	var hash uint64

	if b.Value {
		hash = 1
	} else {
		hash = 0
	}

	return HashKey{Type: BOOLEAN_OBJ, Value: hash}
}
