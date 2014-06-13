package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

func convertExpDate(date string) int {
	Months := map[string]int{
		"Jan": 1,
		"Feb": 2,
		"Mar": 3,
		"Apr": 4,
		"May": 5,
		"Jun": 6,
		"Jul": 7,
		"Aug": 8,
		"Sep": 9,
		"Oct": 10,
		"Nov": 11,
		"Dec": 12,
	}
	dateSplit := strings.Fields(date)
	month := Months[dateSplit[0]]
	year, _ := strconv.Atoi(dateSplit[1])
	return year*12 + month
}

type Period struct {
	raw    string
	sd, ed int
}

type IntSet map[int]bool

func countExperience(test string) int {
	testSplit := strings.Split(test, "; ")
	//periods := make([]Period, len(testSplit))
	intSet := make(IntSet)
	for _, period := range testSplit {
		periodSplit := strings.Split(period, "-")
		startDate, endDate := periodSplit[0], periodSplit[1]
		sd, ed := convertExpDate(startDate), convertExpDate(endDate)
		for i := sd; i <= ed; i++ {
			intSet[i] = true
		}
		//p := Period{period, sd, ed}
		//periods[ixP] = p
	}

	return len(intSet) / 12
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		test := scanner.Text()
		fmt.Println(countExperience(test))
	}
}
