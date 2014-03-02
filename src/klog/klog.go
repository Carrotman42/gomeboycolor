// klog is a package that allows people to selectively log different parts of the emulator,
//   for uses spanning from benchmarking to error checking
package klog

import (
	"fmt"
)

type logkind bool
const (
	Halt logkid = true
)

func (l logkind) If(todo func()) {
	if l {
		todo()
	}
}

func (l logkind) Log(s string, i...interface{}) {
	if l {
		fmt.Println(s, i...)
	}
}