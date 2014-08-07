package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"
import "errors"

func parseIPFromString(in string) []byte {
	inSplit := strings.Split(in, ".")

	out := make([]byte, 4)
	for ix, byteString := range inSplit {
		byte_, _ := strconv.ParseUint(byteString, 10, 16)
		out[ix] = byte(byte_)
	}
	return out
}

func parseTest(in string) ([]byte, []byte, []byte) {
	inSplit := strings.Fields(in)
	source := parseIPFromString(inSplit[0])
	dest := parseIPFromString(inSplit[1])
	rawHeader := inSplit[2:22]
	header := make([]byte, len(rawHeader))
	for ix, strByte := range rawHeader {
		intByte, _ := strconv.ParseUint(strByte, 16, 8)
		header[ix] = byte(intByte)
	}
	return source, dest, header
}

func sum16Bits(in []byte) (int, error) {
	if len(in) < 4 || len(in)%2 != 0 {
		return 0, errors.New("len(in) must be even and longer than 4 - must have at least 2 16 bits number to sum!")
	}
	sum := 0

	for i := 0; i < len(in); i += 2 {
		sum += int(in[i])<<8 + int(in[i+1])
	}

	return sum, nil
}

// See http://www.roman10.net/how-to-calculate-iptcpudp-checksumpart-1-theory/
func headerChecksum(header []byte) int {
	headerWithoutChecksum := make([]byte, len(header))
	_ = copy(headerWithoutChecksum, header)
	// Make sure the zone containing the checksum is zeroed
	headerWithoutChecksum[10] = 0
	headerWithoutChecksum[11] = 0
	//fmt.Println(headerWithoutChecksum[10:12])

	sum, _ := sum16Bits(headerWithoutChecksum)
	//fmt.Println(strconv.FormatInt(int64(sum), 16))

	// Fold the result into 16bits by adding the carry to the result
	carry := sum >> 16
	carryLSh16 := carry << 16
	result := sum - carryLSh16
	//fmt.Println(strconv.FormatInt(int64(carryLSh16), 16))
	//fmt.Println(strconv.FormatInt(int64(result), 16))

	sum16 := result + carry
	//fmt.Println(strconv.FormatInt(int64(sum16), 16))

	sum16not := uint16(^sum16)
	//fmt.Println(sum16not, strconv.FormatInt(int64(sum16not), 16))
	return int(sum16not)
}

func replaceAdresses(source, dest, header []byte) {
	_ = copy(header[12:17], source)
	_ = copy(header[16:20], dest)
	//fmt.Println(header)
	cs := headerChecksum(header)
	csByte1 := byte(cs >> 8)
	csByte2 := byte(cs - (int(csByte1) << 8))
	//fmt.Println(cs, csByte1, csByte2)
	_ = copy(header[10:12], []byte{csByte1, csByte2})
	//fmt.Println(header)
	printHeaderHex(header)
}

func printHeaderHex(header []byte) {
	for ix, byte_ := range header {
		if ix == len(header) { // Last byte ?
			fmt.Printf("%02x ", byte_)
		} else {
			fmt.Printf("%02x ", byte_)
		}
	}
	fmt.Print("\n")
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
		//fmt.Println(line)
		source, dest, header := parseTest(line)
		replaceAdresses(source, dest, header)
		//fmt.Println(source, dest, header)
	}
}
