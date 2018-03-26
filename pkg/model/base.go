package model

import (
	"reflect"
	"strings"

	"github.com/sunho/engbreaker/pkg/dbs"
	"gopkg.in/mgo.v2/bson"
)

func makeName(name string) string {
	return strings.ToLower(name) + "s"
}

func Get(ptr interface{}, query bson.M) error {
	val := reflect.ValueOf(ptr)
	typ := val.Elem().Type()
	if typ.Kind() == reflect.Slice {
		slice := reflect.MakeSlice(typ, 0, 0)
		typ = typ.Elem().Elem()
		iter := reflect.New(typ).Interface()
		results := dbs.MDB.Collection(makeName(typ.Name())).Find(query)
		for results.Next(iter) {
			slice = reflect.Append(slice, reflect.ValueOf(iter))
		}
		slicePtr := reflect.New(slice.Type())
		slicePtr.Elem().Set(slice)
		val.Elem().Set(slicePtr.Elem())
		return nil
	}
	return nil
}
