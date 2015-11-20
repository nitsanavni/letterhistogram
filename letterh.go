package letterhistogram

import (
	"bytes"
	"strings"
)

func Hist(s string) string {
	var letterHist bytes.Buffer
	for _, r := range "abcdefghijklmnopqrstvwxyz" {
		for i := 0; i < strings.Count(s, string(r)); i++ {
			letterHist.WriteString(string(r))
		}
	}
	return letterHist.String()
}