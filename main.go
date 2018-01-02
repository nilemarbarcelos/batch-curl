package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

const url = "http://www.foobar.com/"

func main() {
	data := readData()
	for _, d := range data {
		fmt.Println("Migrating", d)

		payload := "value=true"
		cmd := exec.Command("curl", "-X", "PUT", "-d", payload, "foobar"+d)
		p, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(p))
	}

}

func readData() []string {
	fmt.Println("Reading data from file...")
	var data []string
	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		data = append(data, line)

		if err == io.EOF {
			break
		}
	}
	file.Close()
	return data
}
