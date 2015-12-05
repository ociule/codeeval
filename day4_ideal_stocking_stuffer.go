package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "crypto/md5"
import "encoding/hex"

func solve(in string) int {
	var s string
	var ms string
	h := md5.New()
	for n := 1; n < 1000000000; n++ {
		s = fmt.Sprintf("%s%d", in, n)
		h.Write([]byte(s))
		ms = hex.EncodeToString(h.Sum(nil))
		if strings.HasPrefix(ms, "000000") {
			return n
		}
		h.Reset()
	}
	return -1
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Println(solve(line))
	}
}
