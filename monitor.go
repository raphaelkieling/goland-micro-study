package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	hello()

	for {
		showMenu()

		valCommand := readCommand()
		switch valCommand {
		case 1:
			initMonitor()
		case 2:
			showLogs()
		case 0:
			os.Exit(0)
		}
	}
}

func hello() {
	fmt.Println("Hello world my friend")
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error on show the logs check the log.txt exists")
		fmt.Println(err)
	}

	fmt.Println(string(file))
}

func initMonitor() {
	fmt.Println("Looking for status")
	webLinks := readCheckFile()

	for _, web := range webLinks {
		res, err := http.Get(web)

		if err != nil {
			log.Fatalln("Error on access:", web)
			writeLog("[Offline]: " + web)
		} else if res.StatusCode == 200 {
			fmt.Println("Success on create log!")
			writeLog("[Online]: " + web)
		}
	}
}

func writeLog(content string) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error on create log")
		fmt.Println(err)
	} else {
		file.WriteString(time.Now().Format("02/01/2006") + ": " + content + "\n")
	}

	file.Close()
}

func showMenu() {
	fmt.Println("1 - Init Monitor")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exits")
}

func readCommand() int {
	var command int
	_, err := fmt.Scan(&command)

	if err != nil {
		log.Println(err)
	}

	return command
}

func readCheckFile() []string {
	file, err := os.Open("check.txt")

	if err != nil {
		log.Println("A error ocurred on get check.txt")
	}

	reader := bufio.NewReader(file)
	lines := []string{}

	for {
		row, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		lines = append(lines, strings.TrimSpace(row))
	}

	return lines
}
