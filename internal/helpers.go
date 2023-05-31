package internal

import "fmt"

func notNil(a any, name string) {
	if a == nil {
		msg := fmt.Sprintf("%s is required, but nil pointer received", name)
		panic(msg)
	}
}
