package controllers

import (
	"aws-wiper/types"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func transformError(e error) map[string]interface{} {
	switch e.(type) {
	case validator.ValidationErrors:
		ve := e.(validator.ValidationErrors)

		fields := make(map[string][]string)
		for _, e := range ve {
			fieldName := toSnakeCase(e.Field())
			if _, exists := fields[fieldName]; !exists {
				fields[fieldName] = make([]string, 0)
			}
			fields[fieldName] = append(fields[fieldName], e.Tag())
		}

		return types.Map{"error": fields}
	default:
		return types.Map{"error": e.Error()}
	}
}

func toSnakeCase(s string) string {
	if s == "" {
		return s
	}
	if len(s) == 1 {
		return strings.ToLower(s)
	}
	source := []rune(s)
	dist := strings.Builder{}
	dist.Grow(len(s) + len(s)/3)
	skipNext := false
	for i := 0; i < len(source); i++ {
		cur := source[i]
		switch cur {
		case '-', '_':
			dist.WriteRune('_')
			skipNext = true
			continue
		}
		if unicode.IsLower(cur) || unicode.IsDigit(cur) {
			dist.WriteRune(cur)
			continue
		}

		if i == 0 {
			dist.WriteRune(unicode.ToLower(cur))
			continue
		}

		last := source[i-1]
		if (!unicode.IsLetter(last)) || unicode.IsLower(last) {
			if skipNext {
				skipNext = false
			} else {
				dist.WriteRune('_')
			}
			dist.WriteRune(unicode.ToLower(cur))
			continue
		}
		// last is upper case
		if i < len(source)-1 {
			next := source[i+1]
			if unicode.IsLower(next) {
				if skipNext {
					skipNext = false
				} else {
					dist.WriteRune('_')
				}
				dist.WriteRune(unicode.ToLower(cur))
				continue
			}
		}
		dist.WriteRune(unicode.ToLower(cur))
	}

	return dist.String()
}
