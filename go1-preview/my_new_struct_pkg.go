package my_new_struct_pkg

import (
	"fmt"
)

type Struct struct {
	Public int
	secret int
}

func NewStruct(a int) Struct {  // Note: not a pointer.
	return Struct{a, a*2}
}

func (s Struct) String() string {
	return fmt.Sprintf("{%d (secret %d)}", s.Public, s.secret)
}

