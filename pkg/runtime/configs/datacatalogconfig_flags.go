// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package configs

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (DataCatalogConfig) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (DataCatalogConfig) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in DataCatalogConfig and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg DataCatalogConfig) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("DataCatalogConfig", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "storage-prefix"), *new(string), "StoragePrefix specifies the prefix where DataCatalog stores offloaded ArtifactData in CloudStorage. If not specified,  the data will be stored in the base container directly.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "metrics-scope"), *new(string), "Scope that the metrics will record under.")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "profiler-port"), *new(int), "Port that the profiling service is listening on.")
	return cmdFlags
}
