package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(message string) []string {
	out := make([]string, 0)
	pos := 0
	group_start := 0
	for pos < len(message) {
		if message[pos] == ' ' {
			group_start = pos + 1
			out = append(out, " ")
		} else {
			if (pos-group_start)%2 == 0 {
				out = append(out, message[pos:pos+2])
			}
		}
		pos += 1
	}
	return out
}

func main() {
	message := "012222 1114142503 0313012513 03141418192102 0113 2419182119021713 06131715070119"
	keyed_alphabet := "BHISOECRTMGWYVALUZDNFJKPQX"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	msg := parse(message)
	out := make([]string, 0, len(msg))
	for _, ch := range msg {
		if ch == " " {
			out = append(out, " ")
		} else {
			di, _ := strconv.ParseInt(ch, 10, 0) // 1, 22, 22, 11, 14 ...
			out = append(out, string(alphabet[strings.Index(keyed_alphabet, string(alphabet[di]))]))
		}
	}
	fmt.Println(strings.Join(out, ""))
}
