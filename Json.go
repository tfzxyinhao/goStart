package main

import (
	"bytes"
	"container/list"
	"errors"
	"reflect"
	"strconv"
)

// every level info,now only object count
type ObjecInfo struct {
	count int
}

// json data
type YObject struct {
	buffer bytes.Buffer
	info   *list.List
}

// create new json object and init it
func NewJson() *YObject {
	j := &YObject{}
	j.init()
	return j
}

func (obj *YObject) init() {
	obj.buffer.WriteByte('{')
	obj.info = list.New()
}

// start an json array
func (obj *YObject) BeginArray(key string) {
	obj.info.PushBack(&ObjecInfo{count: 0})
	obj.buffer.WriteString(key)
	obj.buffer.WriteString(":[")
}

// end write json array
func (obj *YObject) EndArray() {
	obj.buffer.WriteString("]")
	info := obj.info.Back()
	if info != nil {
		obj.info.Remove(info)
	}
}

// start to write json object
func (obj *YObject) BeginObject(key string) {
	obj.info.PushBack(&ObjecInfo{count: 0})
	if len(key) != 0 {
		obj.buffer.WriteString(key)
		obj.buffer.WriteString(":{")
	} else {
		obj.buffer.WriteString("{")
	}
}

// end to write json object
func (obj *YObject) EndObject() {
	obj.buffer.WriteString("}")
	info := obj.info.Back()
	if info != nil {
		obj.info.Remove(info)
	}
}

// add object to current level
func (obj *YObject) Add(key string, value interface{}) error {
	info := obj.info.Back().Value.(*ObjecInfo)
	if info == nil {
		return errors.New("Add method must after BeginObject or BeginArray")
	}

	vType := reflect.TypeOf(value).Kind()
	if vType == reflect.String {
		if info.count > 0 {
			obj.buffer.WriteByte(',')
		}
		obj.buffer.WriteString(key)
		obj.buffer.WriteString(":\"")
		obj.buffer.WriteString(value.(string))
		obj.buffer.WriteByte('"')
		info.count++
	} else if vType == reflect.Float64 {
		if info.count > 0 {
			obj.buffer.WriteByte(',')
		}
		obj.buffer.WriteString(key)
		obj.buffer.WriteByte(':')
		obj.buffer.WriteString(strconv.FormatFloat(value.(float64), 'f', 6, 64))
		info.count++
	} else if vType == reflect.Float32 {
		if info.count > 0 {
			obj.buffer.WriteByte(',')
		}
		obj.buffer.WriteString(key)
		obj.buffer.WriteByte(':')
		obj.buffer.WriteString(strconv.FormatFloat(value.(float64), 'f', 6, 32))
		info.count++
	} else if vType == reflect.Uint {
		if info.count > 0 {
			obj.buffer.WriteByte(',')
		}
		obj.buffer.WriteString(key)
		obj.buffer.WriteByte(':')
		obj.buffer.WriteString(strconv.FormatUint(value.(uint64), 10))
		info.count++
	} else if vType == reflect.Int {
		if info.count > 0 {
			obj.buffer.WriteByte(',')
		}
		obj.buffer.WriteString(key)
		obj.buffer.WriteByte(':')
		obj.buffer.WriteString(strconv.Itoa(value.(int)))
		info.count++
	}

	return nil
}

// serialization the json data
func (obj *YObject) ToString() (string, error) {
	if obj.info.Len() != 0 {
		return "", errors.New("Begin operation not match End opertion")
	}

	obj.buffer.WriteByte('}')
	return obj.buffer.String(), nil
}
