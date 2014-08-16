package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "regexp"

type IPParsers struct {
    Regexp *regexp.Regexp
    Name string
}
// Compile the expression once, usually at init time.
// Use raw strings to avoid having to quote the backslashes.
// These regexps will accept some invalid IPs, for example validDec will accept 4999999999 > 255.255.255.255
// This means we'll need a isValid function after cannonization
var validDottedDec = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[1]?[0-9][0-9]?)\.`+
`(25[0-5]|2[0-4][0-9]|[1]?[0-9][0-9]?)\.`+
`(25[0-5]|2[0-4][0-9]|[1]?[0-9][0-9]?)\.`+
`(25[0-5]|2[0-4][0-9]|[1]?[0-9][0-9]?)`)
var validDottedHex = regexp.MustCompile(`0x[0-9a-f]{2}\.0x[0-9a-f]{2}\.0x[0-9a-f]{2}\.0x[0-9a-f]{2}`)
var validDottedOct = regexp.MustCompile(`0[0-7]{3}\.0[0-7]{3}\.0[0-7]{3}\.0[0-7]{3}`)
var validDottedBin = regexp.MustCompile(`[0-1]{8}\.[0-1]{8}\.[0-1]{8}\.[0-1]{8}`)
var validBin = regexp.MustCompile(`[01]{24,31}`)
var validOct = regexp.MustCompile(`[0-3][0-7]{2}[0-3][0-7]{2}[0-3][0-7]{2}[0-3][0-7]{2}`)
var validHex = regexp.MustCompile(`0x[0-9A-F]{8}`)
var validDec = regexp.MustCompile(`[1-4][0-9]{7,9}`)

var parsers = []IPParsers{
    {validDottedDec, "dottedDec"}, {validDottedHex, "dottedHex"},
    {validDottedOct, "dottedOct"}, {validDottedBin, "dottedBin"},
    {validBin, "bin"}, {validOct, "oct"},
    {validHex, "hex"}, {validDec, "dec"},
}

func findIPs(in string) (out []string) {
    out = make([]string, 0, 10)
    for _, r := range parsers {
        s := r.Regexp.FindAllString(in, -1)
        out = append(out, s...)
    }
    return out
}

func parseTest(in string, counts map[string]int) {
    fmt.Println(findIPs(in))
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

    counts := make(map[string]int, 300)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		//fmt.Println(line)
		parseTest(line, counts)
	}
    fmt.Println(counts)
}
