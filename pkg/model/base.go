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
	err := dbs.MDB.Collection(makeName(typ.Name())).FindOne(query, ptr)
	return err
}

func Save(ptr interface{}) error {
	val := reflect.ValueOf(ptr).Elem()
	if val.Type().Kind() == reflect.Slice {
		typ := val.Type().Elem().Elem()
		for i := 0; i < val.Len(); i++ {
			err := dbs.MDB.Collection(makeName(typ.Name())).Save(val.Index(i).Interface().(bongo.Document))
			if err != nil {
				return err
			}
		}
		return nil
	}
	err := dbs.MDB.Collection(makeName(val.Type().Name())).Save(ptr.(bongo.Document))
	return err
}
