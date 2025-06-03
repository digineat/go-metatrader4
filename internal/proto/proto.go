package proto

import (
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
	"unicode"

	"golang.org/x/text/encoding/charmap"
)

// EncodeParams converts params map into a sorted base64-encoded string using Windows-1251 encoding.
func EncodeParams(params map[string]string) (string, error) {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for i, k := range keys {
		if i > 0 {
			sb.WriteByte('|')
		}
		sb.WriteString(k)
		sb.WriteByte('=')
		sb.WriteString(params[k])
	}
	sb.WriteByte('|')

	enc := charmap.Windows1251.NewEncoder()
	encoded, err := enc.String(sb.String())
	if err != nil {
		return "", fmt.Errorf("encode params: %w", err)
	}
	return base64.StdEncoding.EncodeToString([]byte(encoded)), nil
}

// DecodeResponse decodes base64-encoded Windows-1251 text to UTF-8 and removes control characters.
func DecodeResponse(data string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(strings.TrimSpace(data))
	if err != nil {
		return "", fmt.Errorf("base64 decode: %w", err)
	}
	decoded, err := charmap.Windows1251.NewDecoder().Bytes(raw)
	if err != nil {
		return "", fmt.Errorf("decode charset: %w", err)
	}
	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) || r == '\n' || r == '\r' || r == '\t' {
			return r
		}
		return -1
	}, string(decoded))
	return cleaned, nil
}

// BuildRequest returns byte slice representing the command and parameters.
func BuildRequest(command, encodedParams string, quit bool) []byte {
	if quit {
		return []byte(fmt.Sprintf("%s %s\nQUIT\n", command, encodedParams))
	}
	return []byte(fmt.Sprintf("%s %s\n", command, encodedParams))
}
