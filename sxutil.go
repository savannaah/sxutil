package sxutil

import (
	"encoding/base64"
	"fmt"
	"math"
	"reflect"
)

// Base64Encode takes in a string and returns a base 64 encoded string
func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

// Base64Decode takes in a base 64 encoded string and returns the //actual string or an error of it fails to decode the string
func Base64Decode(src string) (string, error) {
	if len(src) == 0 {
		return "", fmt.Errorf("cannot decode empty string, occurred in sxutil package")
	}
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//returns true if value exits  in the list and false otherwise
func Contains(list interface{}, value interface{}) bool {
	arr := reflect.ValueOf(list)
	val := reflect.ValueOf(value)
	if arr.Kind() != reflect.Slice {
		panic("invalid data-type, occurred in sxutil package")
	}
	if val.Kind() == reflect.Slice {
		panic("invalid data-type, occurred in sxutil package")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == value {
			return true
		}
	}
	return false
}

//returns common elements between a and b
func Common(a, b interface{}) (c interface{}) {
	arrA := reflect.ValueOf(a)
	arrB := reflect.ValueOf(b)

	if arrA.Kind() != reflect.Slice {
		panic("invalid data-type, occurred in sxutil package")
	}
	if arrB.Kind() != reflect.Slice {
		panic("invalid data-type, occurred in sxutil package")
	}

	var arrC []interface{}
	m := make(map[interface{}]bool)

	for i := 0; i < arrA.Len(); i++ {
		m[arrA.Index(i).Interface()] = true
	}

	for i := 0; i < arrB.Len(); i++ {
		if _, ok := m[arrB.Index(i).Interface()]; ok {
			arrC = append(arrC, arrB.Index(i).Interface())
		}
	}

	return arrC
}

//returns true every value in a and b matches exactly and false otherwise
func Equal(a, b interface{}) bool {
	arrA := reflect.ValueOf(a)
	arrB := reflect.ValueOf(b)

	if arrA.Kind() != reflect.Slice {
		panic("invalid data-type, occurred in sxutil package")
	}
	if arrB.Kind() != reflect.Slice {
		panic("invalid data-type, occurred in sxutil package")
	}

	if arrA.Len() != arrB.Len() {
		return false
	}

	m := make(map[interface{}]bool)

	for i := 0; i < arrA.Len(); i++ {
		m[arrA.Index(i).Interface()] = true
	}

	for i := 0; i < arrB.Len(); i++ {
		if _, ok := m[arrB.Index(i).Interface()]; !ok {
			return false
		}
	}

	m = make(map[interface{}]bool)

	for i := 0; i < arrB.Len(); i++ {
		m[arrB.Index(i).Interface()] = true
	}

	for i := 0; i < arrA.Len(); i++ {
		if _, ok := m[arrA.Index(i).Interface()]; !ok {
			return false
		}
	}

	return true
}

//returns non duplicate elements
func Unique(a interface{}) interface{} {
	arrA := reflect.ValueOf(a)

	if arrA.Kind() != reflect.Slice {
		panic("invalid data-type, occurred in sxutil package")
	}

	var arrC []interface{}
	m := make(map[interface{}]bool)

	for i := 0; i < arrA.Len(); i++ {
		if _,ok := m[arrA.Index(i).Interface()]; ok {
			continue
		}
		m[arrA.Index(i).Interface()] = true
		arrC = append(arrC, arrA.Index(i).Interface())
	}

	return arrC
}

//returns missing elements in b compared to a
func Missing(a, b interface{}) interface{} {
	arrA := reflect.ValueOf(a)
	arrB := reflect.ValueOf(b)

	if arrA.Kind() != reflect.Slice {
		panic("invalid data-type, occurred in sxutil package")
	}
	if arrB.Kind() != reflect.Slice {
		panic("invalid data-type, occurred in sxutil package")
	}

	var arrC []interface{}
	m := make(map[interface{}]bool)

	for i := 0; i < arrA.Len(); i++ {
		m[arrA.Index(i).Interface()] = true
	}

	for i := 0; i < arrB.Len(); i++ {
		if _, ok := m[arrB.Index(i).Interface()]; !ok {
			arrC = append(arrC, arrB.Index(i).Interface())
		}
	}

	return arrC
}

//returns uncommon elements
func Unmatched(a, b interface{}) interface{} {
	arrA := reflect.ValueOf(a)
	arrB := reflect.ValueOf(b)

	if arrA.Kind() != reflect.Slice {
		panic("invalid data-type, occurred in sxutil package")
	}
	if arrB.Kind() != reflect.Slice {
		panic("invalid data-type, occurred in sxutil package")
	}

	var arrC []interface{}
	m := make(map[interface{}]bool)
	for i := 0; i < arrA.Len(); i++ {
		m[arrA.Index(i).Interface()] = true
	}
	for i := 0; i < arrB.Len(); i++ {
		if _, ok := m[arrB.Index(i).Interface()]; !ok {
			arrC = append(arrC, arrB.Index(i).Interface())
		}
	}

	m = make(map[interface{}]bool)
	for i := 0; i < arrB.Len(); i++ {
		m[arrB.Index(i).Interface()] = true
	}
	for i := 0; i < arrA.Len(); i++ {
		if _, ok := m[arrA.Index(i).Interface()]; !ok {
			arrC = append(arrC, arrA.Index(i).Interface())
		}
	}

	return arrC
}

func CheckDecimalPlaces(place int, value float64) bool {
	valueF := value * math.Pow(10.0, float64(place))
	extra := valueF - float64(int(valueF))

	return extra == 0
}
