package common

import (
	"fmt"
	"strings"
)

func FormatBytes(data []byte) string {
	var b strings.Builder
	b.WriteString("b\"")
	for _, v := range data {
		b.WriteString(fmt.Sprintf("\\x%02x", v))
	}
	b.WriteString("\"")
	return b.String()
}
