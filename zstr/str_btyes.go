package zstr

import (
	"encoding/json"
	"reflect"
	"unsafe"
)

//Str2Bytes is 字符串转bytes不需要通过拷贝
func Str2Bytes(str string) []byte {
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	return b
}

//Bytes2Str is bytes转字符串不需要通过拷贝
func Bytes2Str(b []byte) (s string) {
	pBytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pString := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pString.Data = pBytes.Data
	pString.Len = pBytes.Len
	return
}

func Serialize(data interface{}) []byte {
	res, err := json.Marshal(data)
	if err != nil {
		return []byte{}
	}
	return res
}

func Unserialize(b []byte, dst interface{}) {
	if err := json.Unmarshal(b, dst); err != nil {
		dst = nil
	}
}

func SerializeStr(data interface{}, arg ...interface{}) string {
	return string(Serialize(data))
}
