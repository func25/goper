package development

import (
	"reflect"
)

type UniqueKey interface {
	ToKey() string
}

func getConfigName(confModel any) string {
	if c, ok := confModel.(UniqueKey); ok {
		return c.ToKey()
	}

	return reflect.TypeOf(confModel).Elem().Name()
}
