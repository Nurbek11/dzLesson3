package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func Sort(dir string) {

	f, err := os.Open(dir)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var idx int
	for index, element := range lines {
		if element == ")" {
			idx = index
		}
	}

	var imports = lines[3:idx]
	sort.Strings(imports)

	input, err := ioutil.ReadFile(dir)
	if err != nil {
		log.Fatalln(err)
	}

	linesToWrite := strings.Split(string(input), "\n")

	var index = 0
	for i, line := range linesToWrite {
		//if strings.Contains(line, "]") {
		//	lines[i] = "LOL"
		//}
		if strings.Contains(line, ")") {
			break
		}
		if i > 2 {
			linesToWrite[i] = imports[index]
			index++
		}
	}
	output := strings.Join(linesToWrite, "\n")
	err = ioutil.WriteFile(dir, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {
	Sort("sortImport")

}
