package main

import (
	"strconv"
	"strings"
)

// Unpack function unpacks strings with numbers to strings
//  "a4bc2d5e" => "aaaabccddddde"
//  "abcd" => "abcd"
//  "45" => "" (некорректная строка)
func Unpack(source string) string {
	var result strings.Builder
	var character string
	var escaped bool
	var count int64
	for i, c := range source {
		s := string(c)

		if !escaped && (s == "\\") {
			escaped = true
			continue
		}

		if (i == 0) && strings.ContainsAny(s, "0123456789") {
			result.WriteString("incorrect string")
			break
		}

		if escaped {
			character = s
			result.WriteString(s)
			escaped = false
		} else {
			if strings.ContainsAny(s, "0123456789") {
				if count == 0 {
					count, _ = strconv.ParseInt(s, 10, 0)
				} else {
					tmp, _ := strconv.ParseInt(s, 10, 0)
					count = 10*count + tmp
				}
			} else {
				if count != 0 {
					for i := int64(0); i < count-1; i++ {
						result.WriteString(character)
					}
				}

				character = s
				result.WriteString(character)

				count = 0
			}
		}

	}

	if count != 0 {
		for i := int64(0); i < count-1; i++ {
			result.WriteString(character)
		}
	}

	return result.String()
}
