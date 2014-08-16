package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"
import "sort"
import "regexp"

type IPParser struct {
	Regexp *regexp.Regexp
	Name   string
}

// Compile the expression once, usually at init time.
// Use raw strings to avoid having to quote the backslashes.
// These regexps will accept some invalid IPs, for example validDec will accept 4999999999 > 255.255.255.255
// This means we'll need a isValid function after canonization
var validDottedDec = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[1]?[0-9][0-9]?)\.` +
	`(25[0-5]|2[0-4][0-9]|[1]?[0-9][0-9]?)\.` +
	`(25[0-5]|2[0-4][0-9]|[1]?[0-9][0-9]?)\.` +
	`(25[0-5]|2[0-4][0-9]|[1]?[0-9][0-9]?)`)
var validDottedHex = regexp.MustCompile(`0x[0-9a-f]{2}\.0x[0-9a-f]{2}\.0x[0-9a-f]{2}\.0x[0-9a-f]{2}`)
var validDottedOct = regexp.MustCompile(`0[0-7]{3}\.0[0-7]{3}\.0[0-7]{3}\.0[0-7]{3}`)
var validDottedBin = regexp.MustCompile(`[0-1]{8}\.[0-1]{8}\.[0-1]{8}\.[0-1]{8}`)
var validBin = regexp.MustCompile(`[01]{24,31}`)
var validOct = regexp.MustCompile(`[0-3][0-7]{2}[0-3][0-7]{2}[0-3][0-7]{2}[0-3][0-7]{2}`)
var validHex = regexp.MustCompile(`0x[0-9A-F]{8}`)
var validDec = regexp.MustCompile(`[1-4][0-9]{7,9}`)

var parsers = []IPParser{
	{validDottedDec, "dottedDec"}, {validDottedHex, "dottedHex"},
	{validDottedOct, "dottedOct"}, {validDottedBin, "dottedBin"},
	{validBin, "bin"}, {validOct, "oct"},
	{validHex, "hex"}, {validDec, "dec"},
}

func parseFromDottedHex(in string) (out string) {
	inSplit := strings.Split(in, ".")
	outSlice := make([]string, 4)
	for ix, hexGroup := range inSplit {
		intGroup, _ := strconv.ParseInt(hexGroup[2:], 16, 64)
		outSlice[ix] = strconv.FormatInt(intGroup, 10)
	}
	return strings.Join(outSlice, ".")
}

func parseFromDottedOct(in string) (out string) {
	inSplit := strings.Split(in, ".")
	outSlice := make([]string, 4)
	for ix, octGroup := range inSplit {
		intGroup, _ := strconv.ParseInt(octGroup, 8, 64)
		outSlice[ix] = strconv.FormatInt(intGroup, 10)
	}
	return strings.Join(outSlice, ".")
}

func parseFromDottedBin(in string) (out string) {
	inSplit := strings.Split(in, ".")
	outSlice := make([]string, 4)
	for ix, binGroup := range inSplit {
		intGroup, _ := strconv.ParseInt(binGroup, 2, 64)
		outSlice[ix] = strconv.FormatInt(intGroup, 10)
	}
	return strings.Join(outSlice, ".")
}

func parseFromBase(in string, base int) (out string) {
	intIP, _ := strconv.ParseInt(in, base, 64)
	return stringIPFromInt(int(intIP))
}

const POW224 = 16777216
const POW216 = 65536

// This is like a conversion to base 256, actually
func stringIPFromInt(in int) (out string) {
	g1 := in / POW224
	g2 := in % POW224 / POW216
	g3 := in % POW216 / 256
	g4 := in % 256
	return fmt.Sprintf("%d.%d.%d.%d", g1, g2, g3, g4)
}

func (p *IPParser) toCanonicalForm(in string) (out string) {
	switch p.Name {
	case "dottedDec":
		out = in
	case "dottedHex":
		out = parseFromDottedHex(in)
	case "dottedOct":
		out = parseFromDottedOct(in)
	case "dottedBin":
		out = parseFromDottedBin(in)
	case "bin":
		out = parseFromBase(in, 2)
	case "oct":
		out = parseFromBase(in, 8)
	case "hex":
		out = parseFromBase(in, 16)
	case "dec":
		out = parseFromBase(in, 10)
	}
	return
}

func (p *IPParser) isValid(canonical string) (out bool) {
	out = true
	canSplit := strings.Split(canonical, ".")
	if len(canSplit) != 4 {
		return false
	}
	for ix, strGroup := range canSplit {
		intGroup, _ := strconv.ParseInt(strGroup, 10, 64)
		switch {
		case ix == 0 && (intGroup < 1 || intGroup > 255):
			return false
		case ix > 0 && (intGroup < 0 || intGroup > 255):
			return false
		default:
		}
	}
	return
}

func (p *IPParser) parseRaw(in string) (out []string) {
	rawOut := p.Regexp.FindAllString(in, -1)
	out = make([]string, 0, 10)
	for _, strIP := range rawOut {
		canonical := p.toCanonicalForm(strIP)
		if p.isValid(canonical) {
			out = append(out, canonical)
		}
	}
	return
}

func findIPs(in string) (out []string) {
	out = make([]string, 0, 10)
	for _, p := range parsers {
		strIPs := p.parseRaw(in)
		out = append(out, strIPs...)
	}
	return out
}

func parseTest(in string, counts map[string]int) {
	found := findIPs(in)
	for _, ip := range found {
		val, present := counts[ip]
		if present {
			counts[ip] = val + 1
		} else {
			counts[ip] = 1
		}
	}
}

func printMax(counts map[string]int) {
	maxC := 0
	maxIPs := make([]string, 0, 10)

	for ip, count := range counts {
		switch {
		case count > maxC:
			maxC = count
			maxIPs = make([]string, 0, 10)
			maxIPs = append(maxIPs, ip)
		case count == maxC:
			maxIPs = append(maxIPs, ip)
		}
	}
	if len(maxIPs) > 1 {
		sort.Sort(sort.StringSlice(maxIPs))
	}

	fmt.Println(strings.Join(maxIPs, ""))
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	counts := make(map[string]int, 100)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parseTest(line, counts)
	}
	printMax(counts)
}
