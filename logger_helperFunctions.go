package log

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog"
)

func putValuesToArray(arrs *zerolog.Array, val any) {
	switch t := val.(type) {
	case map[string]any:
		arrs.Dict(createDict(t))
	case []any:
		newArrs := zerolog.Arr()
		for _, v := range t {
			putValuesToArray(newArrs, v)
		}
		arrs.MarshalZerologArray(newArrs)
	case []map[string]any:
		newArrs := zerolog.Arr()
		for _, v := range t {
			putValuesToArray(newArrs, v)
		}
		arrs.MarshalZerologArray(newArrs)
	case string, int, bool, int64, int32, float32, float64:
		arrs.Str(fmt.Sprintf("%v", t))
	default:
		arrs.Str(fmt.Sprintf("%v", t))
	}
}

func createDict(t map[string]any) *zerolog.Event {
	d := zerolog.Dict()
	for k, v := range t {
		putValueToEvent(d, k, v)
	}
	return d
}

func createArrayFromAny(t []any) *zerolog.Array {
	arrs := zerolog.Arr()
	for _, v := range t {
		putValuesToArray(arrs, v)
	}
	return arrs
}

func createArrayFromMap(t []map[string]any) *zerolog.Array {
	arrs := zerolog.Arr()
	for _, v := range t {
		putValuesToArray(arrs, v)
	}
	return arrs
}

func fuzzyPutValueToEvent(event *zerolog.Event, key string, t any) {
	var m map[string]any
	var a []any
	if dat, err := json.Marshal(t); err != nil {
		event.Str(key, fmt.Sprintf("%v", t))
	} else if err := json.Unmarshal(dat, &m); err != nil {
		if err := json.Unmarshal(dat, &a); err != nil {
			event.Str(key, fmt.Sprintf("%v", t))
		} else {
			arrs := zerolog.Arr()
			for _, v := range m {
				putValuesToArray(arrs, v)
			}
			event.Array(key, arrs)
		}
	} else {
		for k, v := range m {
			putValueToEvent(event, k, v)
		}
	}
}
