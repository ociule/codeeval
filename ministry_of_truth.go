package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

const IMPOSSIBLE = "I cannot fix history"

func encode(a, b string) (string) {
    aSplit := strings.Fields(a)
    bSplit := strings.Fields(b)

    if len(bSplit) > len(aSplit) {
        return IMPOSSIBLE
    }
    outSlice := make([]string, 0, len(aSplit))
    posA, posB := 0, 0
    for posB < len(bSplit) && posA < len(aSplit) {
        aWord, bWord := aSplit[posA], bSplit[posB]
        fmt.Println("====", aWord, bWord)
        // can we match bWord inside aWord ?
        if strings.Index(aWord, bWord) > -1 {
            outSlice = append(outSlice, aSplit[posA] + bWord)
            // Let's continue to the next bWord
            posA += 1
            posB += 1
        } else { // no ? move to next aWord
            outSlice = append(outSlice, strings.Repeat("_", len(aWord)))
            posA += 1
        }
        //fmt.Println("+++", posA, posB)
        // Did we get to the end of A without finishing b ? IMPOSSIBLE
        if posA >= len(aSplit) && posB < (len(aSplit) - 1) {
            return IMPOSSIBLE
        }

        // Did we get to the end of B without finishing A ? Finish censoring A
        if posB >= len(bSplit) && posA < len(aSplit) {
            for posA < len(aSplit) {
                aWord := aSplit[posA]
                outSlice = append(outSlice, strings.Repeat("_", len(aWord)))
                posA += 1
            }
        }
    }
	return strings.Join(outSlice, " ")
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
        inS := strings.Split(line, ";")
        expected := ""
		a, b := inS[0], inS[1]
        if len(inS) > 2 {
            expected = inS[2]
        }
		fmt.Println(">>>", a, b, expected)
		fmt.Println(encode(a, b))
	}
}
