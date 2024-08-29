package utils

import "unicode"

// CamelToSnake 将 CamelCase 转换为 snake_case
func CamelToSnake(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

// GetStringFromMap 通用的字符串提取函数
func GetStringFromMap(m map[string]interface{}, key string) (string, bool) {
	value, ok := m[key].(string)
	if !ok || value == "" {
		return "", false
	}
	return value, true
}
