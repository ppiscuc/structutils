package structutils

import (
	"reflect"
	"strings"
)

//Methods lists the methods implemented by a struct
func Methods(v interface{}) []string {
	var methods []string
	typ := reflect.TypeOf(v)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i).Name
		methods = append(methods, method)
	}
	return methods
}

//StructJSONTags gets a list of JSON tags
func StructJSONTags(item interface{}) []string {
	var vals []string
	val := reflect.ValueOf(item)
	for i := 0; i < val.Type().NumField(); i++ {
		tags := val.Type().Field(i).Tag.Get("json")
		csvtags := strings.Split(tags, ",")
		vals = append(vals, csvtags[0])
	}
	return vals
}
