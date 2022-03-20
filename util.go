package topaz

import "fmt"

func errorf(err error, values ...interface{}) error {
	return fmt.Errorf(err.Error(), values...)
}
