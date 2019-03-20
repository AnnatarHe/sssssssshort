package src

import (
	"fmt"
	"strings"
)

var codes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"

func Encode(id int64) string {
	str := make([]byte, 0, 12)
	if id == 0 {
		return "0"
	}
	for id > 0 {
		ch := codes[id%64]
		str = append(str, byte(ch))
		id /= 64
	}
	return string(str)
}

func Decode(url string) (int64, error) {
	res := int64(0)

	for i := len(url); i > 0; i-- {
		ch := url[i-1]
		res *= 64
		mod := strings.IndexRune(codes, rune(ch))
		if mod == -1 {
			return -1, fmt.Errorf("Invalid url character: '%c'", ch)
		}
		res += int64(mod)
	}
	return res, nil
}
