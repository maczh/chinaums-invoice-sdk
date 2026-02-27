package util

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// Sign 生成签名（遵循文档2.2签名规则）
func Sign(params map[string]interface{}, key string) string {
	// 1. 过滤空值参数，按ASCII排序
	var keys []string
	for k, v := range params {
		if v != nil && v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// 2. 拼接key=value&字符串
	var buf bytes.Buffer
	for i, k := range keys {
		if i > 0 {
			buf.WriteString("&")
		}
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(toString(params[k]))
	}

	// 3. 拼接key，SHA256加密，转大写
	buf.WriteString(key)
	hash := sha256.Sum256(buf.Bytes())
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

// toString 类型转换（支持基础类型和JSON对象）
func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int, uint, int8, int16, int32, int64, float64, float32, bool:
		return fmt.Sprintf("%v", val)
	default:
		// JSON对象序列化（保持字段顺序）
		data, _ := MarshalJSON(val)
		return string(data)
	}
}
