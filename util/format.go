package util

import "strconv"
import "github.com/shopspring/decimal"

func ToFloat32(v interface{}) float32 {
	if v == nil {
		return 0
	}
	switch v.(type) {
	case float32:
		return v.(float32)
	case string:
		vs, _ := strconv.ParseFloat(v.(string), 32)
		return float32(vs)
	case int:
		return float32(v.(int))
	case float64:
		return float32(v.(float64))
	case int32:
		return float32(int(v.(int32)))
	case int64:
		return float32(int(v.(int64)))
	}
	return 0
}

func ToFloat64(v interface{}) float64 {
	if v == nil {
		return 0
	}
	switch v.(type) {
	case float64:
		return v.(float64)
	case string:
		vs, _ := strconv.ParseFloat(v.(string), 64)
		return vs
	case int:
		return float64(v.(int))
	case float32:
		return float64(v.(float32))
	case int32:
		return float64(int(v.(int32)))
	case int64:
		return float64(int(v.(int64)))
	}
	return 0
}

func ToInt(v interface{}) int {

	if v == nil {
		return 0
	}
	switch v.(type) {
	case int:
		return v.(int)
	case string:
		vs, _ := strconv.Atoi(v.(string))
		return vs
	case float64:
		return int(decimal.NewFromFloat(v.(float64)).IntPart())
	case float32:
		return int(decimal.NewFromFloat32(v.(float32)).IntPart())
	case int32:
		return int(v.(int32))
	case int64:
		return int(v.(int64))
	}
	return 0
}

func ToInt32(v interface{}) int32 {
	return int32(ToInt(v))
}

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch v.(type) {
	case string:
		return v.(string)
	case int:
		return strconv.Itoa(v.(int))
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32)
	case int32:
		return strconv.Itoa(int(v.(int32)))
	case int64:
		return strconv.Itoa(int(v.(int64)))
	case uint8:
		return strconv.Itoa(int(v.(uint8)))
	}
	return ""
}
