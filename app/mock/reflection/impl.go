package reflection

import (
	"fmt"
	"reflect"
)

/*
GenerateReflectionCopy
recursively generates reflected copy of t via property pattern
*/
func GenerateReflectionCopy(t map[string]interface{}) map[string]interface{} {
	for key, v := range t {

		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			if mv, ok := v.(map[string]interface{}); ok {
				GenerateReflectionCopy(mv)
			} else {
				fmt.Println("unknown property from reflect.Map")
			}

		default:
			if mv, ok := v.(PropertyPattern); ok {
				fmt.Println(mv)
				if mv.Kind == single {
					GenerateValueFromReflection(mv.Reflect, t, key)
				} else {
					a := make([]interface{}, mv.Len)
					m := make(map[string]interface{}, 1)
					for i := range a {
						GenerateValueFromReflection(mv.Reflect, m, "k")
						a[i] = m["k"]
					}
					t[key] = a
				}
			} else {
				fmt.Println("unknown property from reflect.Map")
			}
		}
	}
	return t
}

/*
ReflectAndDescribe
recursively reflects data types from t
if array reflection found, items will be recast to interfaces
*/

func ReflectAndDescribe(t map[string]interface{}) map[string]interface{} {
	for key, v := range t {
		rv := reflect.ValueOf(v)

		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			t[key] = PropertyPattern{
				Reflect:         reflect.String,
				ReflectionLabel: reflect.String.String(),
				Kind:            single,
				Len:             rv.Len(),
			}
			fmt.Println(reflect.String)
		case reflect.Bool:
			t[key] = PropertyPattern{
				Reflect:         reflect.Bool,
				ReflectionLabel: reflect.Bool.String(),
				Kind:            single,
			}
		case reflect.Float64:
			t[key] = PropertyPattern{
				Reflect:         reflect.Float64,
				ReflectionLabel: reflect.Float64.String(),
				Kind:            single,
			}
		case reflect.Map:
			if mv, ok := v.(map[string]interface{}); ok {
				ReflectAndDescribe(mv)
			} else {
				fmt.Println("unknown property from reflect.Map")
			}

		case reflect.Array, reflect.Slice:
			//test := reflect.TypeOf(v).Elem()
			ret := make([]interface{}, rv.Len())
			for i := range ret {
				ret[i] = rv.Index(i).Interface()
			}
			kind := reflect.TypeOf(ret[0]).Kind()
			t[key] = PropertyPattern{
				Reflect:         kind,
				ReflectionLabel: reflect.Float64.String(),
				Kind:            slice,
				Len:             rv.Len(),
			}
		}
	}

	fmt.Println(t)
	return t
}

/*
GenerateValueFromReflection
generates random values by types
*/

func GenerateValueFromReflection(k reflect.Kind, t map[string]interface{}, key string) {
	switch k {
	case reflect.String:
		t[key] = "pes"
	case reflect.Float64:
		t[key] = 999
	case reflect.Bool:
		t[key] = true
	}
}



func GenerateReflectionData(t map[string]interface{}, l int) []map[string]interface{} {
	data := make([]map[string]interface{}, l)
	for i := 0; i < l; i++ {
		copy := GenerateReflectionCopy(t)
		data = append(data, copy)
	}
	return data
}