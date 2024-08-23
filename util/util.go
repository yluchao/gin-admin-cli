package util

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/jinzhu/inflection"
)

// Copied from golint
var commonInitialisms = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
var commonInitialismsReplacer *strings.Replacer

func init() {
	var commonInitialismsForReplacer []string
	for _, initialism := range commonInitialisms {
		commonInitialismsForReplacer = append(commonInitialismsForReplacer, initialism, strings.Title(strings.ToLower(initialism)))
	}
	commonInitialismsReplacer = strings.NewReplacer(commonInitialismsForReplacer...)
}

// ToLowerUnderlinedNamer 转换为小写下划线命名
func ToLowerUnderlinedNamer(name string) string {
	const (
		lower = false
		upper = true
	)

	if name == "" {
		return ""
	}

	var (
		value                                    = commonInitialismsReplacer.Replace(name)
		buf                                      = bytes.NewBufferString("")
		lastCase, currCase, nextCase, nextNumber bool
	)

	for i, v := range value[:len(value)-1] {
		nextCase = bool(value[i+1] >= 'A' && value[i+1] <= 'Z')
		nextNumber = bool(value[i+1] >= '0' && value[i+1] <= '9')

		if i > 0 {
			if currCase == upper {
				if lastCase == upper && (nextCase == upper || nextNumber == upper) {
					buf.WriteRune(v)
				} else {
					if value[i-1] != '_' && value[i+1] != '_' {
						buf.WriteRune('_')
					}
					buf.WriteRune(v)
				}
			} else {
				buf.WriteRune(v)
				if i == len(value)-2 && (nextCase == upper && nextNumber == lower) {
					buf.WriteRune('_')
				}
			}
		} else {
			currCase = upper
			buf.WriteRune(v)
		}
		lastCase = currCase
		currCase = nextCase
	}

	buf.WriteByte(value[len(value)-1])

	s := strings.ToLower(buf.String())
	return s
}

// SnakeToPascalCase 将蛇形命名转换为首字母大写的命名方式（PascalCase）
func SnakeToPascalCase(s string) string {
	// 将字符串按下划线分割
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if part == "" {
			continue
		}
		// 将每部分的首字母转为大写
		runes := []rune(part)
		runes[0] = unicode.ToUpper(runes[0])
		parts[i] = string(runes)
	}
	// 将所有部分拼接起来
	return strings.Join(parts, "")
}

// SnakeToCamelCase 将蛇形命名转换为驼峰命名（CamelCase）
func SnakeToCamelCase(s string) string {
	pascalCase := SnakeToPascalCase(s)
	if pascalCase == "" {
		return ""
	}
	// 将第一个字符转为小写
	runes := []rune(pascalCase)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// ToPascalCase 统一处理小写驼峰命名 (camelCase) 和下划线命名 (snake_case) 转换为 PascalCase
func ToPascalCase(s string) string {
	if s == "" {
		return ""
	}

	// 创建一个字符串构建器
	var result strings.Builder
	result.Grow(len(s))

	// 标志下一个字符是否应该大写
	shouldCapitalize := true

	for _, r := range s {
		if r == '_' { // 遇到下划线
			shouldCapitalize = true
			continue
		}

		if shouldCapitalize {
			result.WriteRune(unicode.ToUpper(r))
			shouldCapitalize = false
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}

// ToPlural 转换为复数
func ToPlural(v string) string {
	return inflection.Plural(v)
}
