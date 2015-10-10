package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

type BFMachine struct {
	IP      int
	DP      int
	Program string
	Cells   [50000]int8
	Jumps   map[int]int
	Output  []string
}

func InitBFMachine(program string) BFMachine {
	bfm := BFMachine{}
	// Start position will be middle of data array, to allow "negative" DP via multiple < instructions
	bfm.DP = len(bfm.Cells) / 2
	bfm.Program = program
	bfm.Output = make([]string, 0, 10000)
	bfm.Jumps = bfm.GetJumps()
	return bfm
}

func (bfm *BFMachine) GetJumps() map[int]int {
	// Build a jump table
	nJumps := strings.Count(bfm.Program, "[") + strings.Count(bfm.Program, "]")
	jumps := make(map[int]int, nJumps)

	for pos, r := range bfm.Program {
		if r == ']' {
			countLoop := 0
			for i := pos - 1; i >= 0; i-- {
				currentInst := bfm.Program[i]
				if currentInst == '[' {
					if countLoop == 0 {
						jumps[pos] = i
						jumps[i] = pos
						break
					} else {
						countLoop -= 1
					}
				}
				if currentInst == ']' {
					countLoop += 1
				}
			}
		}
	}
	return jumps
}

func (bfm *BFMachine) CurrentCell() int8 {
	return bfm.Cells[bfm.DP]
}

func (bfm *BFMachine) Jumpback() {
	if bfm.CurrentCell() != 0 {
		bfm.IP = bfm.Jumps[bfm.IP]
	}
}

func (bfm *BFMachine) Jumpforward() {
	if bfm.CurrentCell() == 0 {
		bfm.IP = bfm.Jumps[bfm.IP]
	}
}

func (bfm *BFMachine) Run() {
	pl := len(bfm.Program)
	for bfm.IP < pl {
		currentInstruction := bfm.Program[bfm.IP]
		//fmt.Println("will execute", string(currentInstruction))
		switch currentInstruction {
		case '+':
			bfm.Cells[bfm.DP] += 1
		case '-':
			bfm.Cells[bfm.DP] -= 1
		case '>':
			bfm.DP += 1
		case '<':
			bfm.DP -= 1
		case '.':
			bfm.Output = append(bfm.Output, string(bfm.Cells[bfm.DP]))
		case '[':
			bfm.Jumpforward()
		case ']':
			bfm.Jumpback()
		}
		bfm.IP += 1
	}
}

func (bfm *BFMachine) PPrint() {
	temp := make([]string, 0, 200)
	for n := 0; n < 200; n += 1 {
		temp = append(temp, string(bfm.Cells[n]))
	}
	fmt.Println(bfm.Jumps)
	//fmt.Println(strings.Join(temp, ""))
}

func solve(in string) string {
	bfm := InitBFMachine(in)
	bfm.Run()
	//bfm.PPrint()
	return strings.Join(bfm.Output, "")
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
