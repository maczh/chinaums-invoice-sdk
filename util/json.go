package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// MarshalJSON 保持字段顺序的JSON序列化
// 核心逻辑：
// 1. 对结构体：按字段定义顺序序列化
// 2. 对Map：按key的ASCII升序序列化
// 3. 对基础类型：直接序列化
func MarshalJSON(v interface{}) ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	val := reflect.ValueOf(v)
	// 处理指针类型
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return []byte("null"), nil
		}
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		return marshalStruct(val)
	case reflect.Map:
		return marshalMap(val)
	case reflect.Slice, reflect.Array:
		return marshalSlice(val)
	default:
		// 基础类型（string/int/bool等）
		return json.Marshal(v)
	}
}

// marshalStruct 按结构体字段定义顺序序列化
func marshalStruct(val reflect.Value) ([]byte, error) {
	var buf strings.Builder
	buf.WriteString("{")

	typ := val.Type()
	fieldCount := val.NumField()
	written := 0

	for i := 0; i < fieldCount; i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		// 跳过未导出字段（首字母小写）
		if field.PkgPath != "" {
			continue
		}

		// 获取JSON tag（优先用tag，无tag则用字段名）
		jsonTag := field.Tag.Get("json")
		if jsonTag == "-" {
			continue // 跳过指定忽略的字段
		}

		// 解析JSON tag（处理 omitempty 等选项）
		tagParts := strings.Split(jsonTag, ",")
		fieldName := tagParts[0]
		if fieldName == "" {
			fieldName = field.Name // 无tag时使用字段名
		}

		// 处理 omitempty：字段值为零值则跳过
		if len(tagParts) > 1 && tagParts[1] == "omitempty" && isZeroValue(fieldVal) {
			continue
		}

		// 序列化字段值
		fieldBytes, err := MarshalJSON(fieldVal.Interface())
		if err != nil {
			return nil, fmt.Errorf("marshal field %s error: %w", fieldName, err)
		}

		// 拼接字段（非第一个字段需加逗号）
		if written > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(`"` + fieldName + `":`)
		buf.Write(fieldBytes)
		written++
	}

	buf.WriteString("}")
	return []byte(buf.String()), nil
}

// marshalMap 按key的ASCII升序序列化Map
func marshalMap(val reflect.Value) ([]byte, error) {
	var buf strings.Builder
	buf.WriteString("{")

	// 获取所有key并按ASCII升序排序
	keys := val.MapKeys()
	sort.Slice(keys, func(i, j int) bool {
		// 将key转为字符串后比较
		keyI := keyToString(keys[i])
		keyJ := keyToString(keys[j])
		return keyI < keyJ
	})

	written := 0
	for _, key := range keys {
		// 序列化key
		keyBytes, err := MarshalJSON(key.Interface())
		if err != nil {
			return nil, fmt.Errorf("marshal map key error: %w", err)
		}

		// 序列化value
		valBytes, err := MarshalJSON(val.MapIndex(key).Interface())
		if err != nil {
			return nil, fmt.Errorf("marshal map value error: %w", err)
		}

		// 拼接键值对
		if written > 0 {
			buf.WriteString(",")
		}
		buf.Write(keyBytes)
		buf.WriteString(":")
		buf.Write(valBytes)
		written++
	}

	buf.WriteString("}")
	return []byte(buf.String()), nil
}

// marshalSlice 序列化切片/数组
func marshalSlice(val reflect.Value) ([]byte, error) {
	var buf strings.Builder
	buf.WriteString("[")

	length := val.Len()
	for i := 0; i < length; i++ {
		elemBytes, err := MarshalJSON(val.Index(i).Interface())
		if err != nil {
			return nil, fmt.Errorf("marshal slice element %d error: %w", i, err)
		}

		if i > 0 {
			buf.WriteString(",")
		}
		buf.Write(elemBytes)
	}

	buf.WriteString("]")
	return []byte(buf.String()), nil
}

// keyToString 将Map的key转为字符串（用于排序）
func keyToString(key reflect.Value) string {
	switch key.Kind() {
	case reflect.String:
		return key.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(key.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(key.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(key.Float(), 'f', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(key.Bool())
	default:
		// 复杂类型（如结构体）序列化后作为key
		bytes, _ := MarshalJSON(key.Interface())
		return string(bytes)
	}
}

// isZeroValue 判断字段值是否为零值（用于omitempty）
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	case reflect.Struct:
		// 结构体零值：所有字段都是零值
		typ := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := typ.Field(i)
			if field.PkgPath != "" {
				continue // 跳过未导出字段
			}
			if !isZeroValue(v.Field(i)) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// StructToMap 将结构体转为Map（保持字段顺序）
// 用于签名时的参数拼接
func StructToMap(v interface{}) (map[string]interface{}, error) {
	jsonStr := ToJsonString(v)
	if jsonStr == "" {
		return nil, errors.New("marshal struct to json string error")
	}
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, fmt.Errorf("unmarshal json string to map error: %w", err)
	}
	return result, nil
}

func ToJsonString(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}
