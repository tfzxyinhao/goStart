package main

import (
	"bytes"
	"container/list"
	"errors"
	"reflect"
	"strconv"
)

const (
	YES = iota
	NEEDRB
	NEEDRM
)

type ObjecInfo struct {
	count int
}

type YObject struct {
	buffer bytes.Buffer
	info   *list.List
}

func NewJson() *YObject {
	j := &YObject{}
	j.init()
	return j
}

func (obj *YObject) init() {
	obj.buffer.WriteByte('{')
	obj.info = list.New()
}

func (obj *YObject) BeginArray(key string) {
	obj.info.PushBack(&ObjecInfo{count: 0})
	obj.buffer.WriteString(key)
	obj.buffer.WriteString(":[")
}

func (obj *YObject) EndArray() {
	obj.buffer.WriteString("]")
	info := obj.info.Back()
	if info != nil {
		obj.info.Remove(info)
	}
}

func (obj *YObject) BeginObject(key string) {
	obj.info.PushBack(&ObjecInfo{count: 0})
	if len(key) != 0 {
		obj.buffer.WriteString(key)
		obj.buffer.WriteString(":{")
	} else {
		obj.buffer.WriteString("{")
	}
}

func (obj *YObject) EndObject() {
	obj.buffer.WriteString("}")
	info := obj.info.Back()
	if info != nil {
		obj.info.Remove(info)
	}
}

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
		info.count += 1
	} else if vType == reflect.Float64 {
		if info.count > 0 {
			obj.buffer.WriteByte(',')
		}
		obj.buffer.WriteString(key)
		obj.buffer.WriteByte(':')
		obj.buffer.WriteString(strconv.FormatFloat(value.(float64), 'f', 6, 64))
		info.count += 1
	} else if vType == reflect.Float32 {
		if info.count > 0 {
			obj.buffer.WriteByte(',')
		}
		obj.buffer.WriteString(key)
		obj.buffer.WriteByte(':')
		obj.buffer.WriteString(strconv.FormatFloat(value.(float64), 'f', 6, 32))
		info.count += 1
	} else if vType == reflect.Uint {
		if info.count > 0 {
			obj.buffer.WriteByte(',')
		}
		obj.buffer.WriteString(key)
		obj.buffer.WriteByte(':')
		obj.buffer.WriteString(strconv.FormatUint(value.(uint64), 10))
		info.count += 1
	} else if vType == reflect.Int {
		if info.count > 0 {
			obj.buffer.WriteByte(',')
		}
		obj.buffer.WriteString(key)
		obj.buffer.WriteByte(':')
		obj.buffer.WriteString(strconv.Itoa(value.(int)))
		info.count += 1
	}

	return nil
}

func (obj *YObject) ToString() (string, error) {
	if obj.info.Len() != 0 {
		return "", errors.New("Begin operation not match End opertion")
	}

	obj.buffer.WriteByte('}')
	return obj.buffer.String(), nil
}
