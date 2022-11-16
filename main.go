package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var mode int
	fmt.Println("Выберите режим работы. Введите 1 для обработки http, 2 для обработки https: ")
	fmt.Scan(&mode)
	file, err := os.Open("Free_Proxy_List.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fscaner := bufio.NewScanner(file)
		for fscaner.Scan() {
			line := fscaner.Text()
			arr_from_line := strings.Split(line, "\",")
			if len(arr_from_line) == 1 {
				continue
			}
			ip := strings.Replace(arr_from_line[0], "\"", "", -1)
			port := strings.Replace(arr_from_line[7], "\"", "", -1)
			proxy := ip + ":" + port + "\n"
			if mode == 1 {
				result_file, err := os.OpenFile("http.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				if err != nil {
					log.Fatal(err)
				} else {
					if _, err := result_file.WriteString(proxy); err != nil {
						log.Fatal(err)
					}
				}
				result_file.Close()
			}
			if mode == 2 {
				result_file, err := os.OpenFile("https.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				if err != nil {
					log.Fatal(err)
				} else {
					if _, err := result_file.WriteString(proxy); err != nil {
						log.Fatal(err)
					}
				}
				result_file.Close()
			}
		}
	}
	file.Close()
}
