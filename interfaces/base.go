package interfaces

import (
	"database/sql"
	"unicode"
)

// Logger interface
type Logger interface {
	Log(message string) error
}

// DBHandler interface
type DBHandler interface {
	Query(statement string) (*sql.Rows, error)
}

func toSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}
