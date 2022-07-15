package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	nomes := []string{"João", "Maria", "José"}

	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
		time.Sleep(1 * time.Second)
	}

	for index, nome := range nomes {
		fmt.Println("Index: ", index, "Nome: ", nome)
	}

	names := readNamesInArchive()
	fmt.Println(names[0])

	fmt.Println("------------------------------------------")

	for _, e := range names {
		registerLogs(e)
	}

	showLogs()
}

func readNamesInArchive() []string {
	archive, err := ioutil.ReadFile("nomes.txt")
	arquivo, _ := os.Open("nomes.txt")

	var names []string

	if err != nil {
		fmt.Println("Arquivo: ", err)
		return nil
	}

	reader := bufio.NewReader(arquivo)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		names = append(names, line)
		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	fmt.Println(string(archive))
	return names
}

func registerLogs(log string) {
	archive, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	archive.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + log + "\n")
}

func showLogs() {
	archive, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(archive))
}
