package model

import (
	"reflect"
	"strings"

	"github.com/go-bongo/bongo"
	"github.com/sunho/engbreaker/pkg/dbs"
	"gopkg.in/mgo.v2/bson"
)

func makeName(name string) string {
	return strings.ToLower(name) + "s"
}

func Get(ptr interface{}, query bson.M) error {
	val := reflect.ValueOf(ptr) // *model or *[]model
	typ := val.Elem().Type()    // model or []model
	if typ.Kind() == reflect.Slice {
		typ = typ.Elem() // model
		if typ.Kind() == reflect.Ptr {
			panic("This should not be pointer")
		}

		slice := reflect.MakeSlice(reflect.SliceOf(typ), 0, 0)
		iter := reflect.New(typ).Interface() // *model

		results := dbs.MDB.Collection(makeName(typ.Name())).Find(query)
		for results.Next(iter) {
			slice = reflect.Append(slice, reflect.ValueOf(iter).Elem())
		}

		slicePtr := reflect.New(slice.Type()) // *[]model
		slicePtr.Elem().Set(slice)
		val.Elem().Set(slicePtr.Elem())
		return nil
	}
	err := dbs.MDB.Collection(makeName(typ.Name())).FindOne(query, ptr)
	return err
}

func toStructPtr(obj interface{}) interface{} {
	val := reflect.ValueOf(obj)
	vp := reflect.New(val.Type())
	vp.Elem().Set(val)
	return vp.Interface()
}
func Save(ptr interface{}) error {
	val := reflect.ValueOf(ptr).Elem() // model or []model
	typ := val.Type()
	if typ.Kind() == reflect.Slice {
		typ = typ.Elem() // model
		for i := 0; i < val.Len(); i++ {
			err := dbs.MDB.Collection(makeName(typ.Name())).Save(toStructPtr(val.Index(i).Interface()).(bongo.Document))
			if err != nil {
				return err
			}
		}
		return nil
	}
	err := dbs.MDB.Collection(makeName(typ.Name())).Save(ptr.(bongo.Document))
	return err
}
