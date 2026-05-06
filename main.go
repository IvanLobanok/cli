package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var tm TaskManager
	var exitFlag = true
	var filename string = "tasks.json"

	tm.LoadFromFile(filename)
	scanner := bufio.NewScanner(os.Stdin)

	for exitFlag {
		fmt.Println("\nЧто будем делать?")
		fmt.Println("1 - вывести все задачи")
		fmt.Println("2 - вывести невыполненные задачи")
		fmt.Println("3 - вывести выполненные задачи")
		fmt.Println("4 - добавить задачу")
		fmt.Println("5 - изменить статус задачи")
		fmt.Println("6 - удалить задачу")
		fmt.Println("0 - ВЫХОД")

		scanner.Scan()
		choise := strings.TrimSpace(scanner.Text())

		switch choise {
		case "1":
			tm.ViewAll()
		case "2":
			tm.ViewAllInProgress()
		case "3":
			tm.ViewAllDone()
		case "4":
			fmt.Println("Введите название новой задачи:")
			scanner.Scan()
			str := strings.TrimSpace(scanner.Text())
			tm.Add(str)
			tm.SaveToFile(filename)
		case "5":
			fmt.Println("Введите ID задачи:")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			idInt, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}
			err = tm.UpdateStatus(idInt)
			if err != nil {
				log.Fatal(err)
			}
			tm.SaveToFile(filename)

		case "6":
			fmt.Println("Введите ID задачи:")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			idInt, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}
			err = tm.Delete(idInt)
			if err != nil {
				log.Fatal(err)
			}
			tm.SaveToFile(filename)

		case "0":
			fmt.Println("Досвидания!")
			exitFlag = !exitFlag
		}
	}
}
