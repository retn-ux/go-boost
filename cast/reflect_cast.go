package cast

import "reflect"

func DynamicCast(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	ok = false
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
