package module

import (
	"fmt"
)

func LogDebug(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}