package topaz

import "fmt"

func errorf(err error, values ...any) error {
	return fmt.Errorf(err.Error(), values...)
}
