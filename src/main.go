package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func displayFile(file *os.File) {
	scaner := bufio.NewScanner(file)

	for scaner.Scan() {
		fmt.Println(scaner.Text())
	}
}

func dispAndWrite(input string, file *os.File) {

	t := time.Now()

	tlMsg := fmt.Sprintf("%v:%s\n", t.Format("15:04:05"), input)

	fmt.Print(tlMsg)
	file.WriteString(tlMsg)
}

func main() {

	file, _ := os.OpenFile(`./TL.txt`, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()

	flag.Parse()

	if flag.NArg() != 0 {

		arg0 := flag.Arg(0)
		dispAndWrite(arg0, file)

	}

	displayFile(file)

	var input string

	for {
		fmt.Scan(&input)

		dispAndWrite(input, file)

		//Ctrl+C to exit?
	}
}
