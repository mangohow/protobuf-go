package internal_gengo

import (
	"strings"
	"unicode"
)

// ToCamelCase 将变量名转换为驼峰命名
func ToCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	if len(words) == 0 {
		return ""
	}
	if len(words[0]) > 0 {
		words[0] = strings.ToLower(words[0][:1]) + words[0][1:]
	}

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

// ToPascalCase 将变量名转换为帕斯卡命名
func ToPascalCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

// ToSnakeCase 将变量名转换为下划线命名
func ToSnakeCase(s string) string {
	var builder strings.Builder

	for i, char := range s {
		if unicode.IsUpper(char) {
			if i != 0 {
				builder.WriteRune('_')
			}
			builder.WriteRune(unicode.ToLower(char))
		} else {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}
