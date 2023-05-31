package internal

import (
	"fmt"

	"github.com/rs/xid"
)

func notNil(a any, name string) {
	if a == nil {
		msg := fmt.Sprintf("%s is required, but nil pointer received", name)
		panic(msg)
	}
}

func toXid(ids []string) ([]xid.ID, error) {
	x := make([]xid.ID, len(ids))
	var err error
	for i, str := range ids {
		x[i], err = xid.FromString(str)
		if err != nil {
			return nil, err
		}
	}
	return x, nil
}
