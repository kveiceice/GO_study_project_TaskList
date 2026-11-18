package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	event "study_project_1/events"
	"study_project_1/task"
	"time"
)

var curId int = 1

func nextId() {
	curId++
}

func slice_to_text(arr []string) string {
	new_string := ""
	for i := range arr {
		if i != len(arr)-1 {
			new_string += arr[i]
			new_string += " "
		} else {
			new_string += arr[i]
		}
	}
	return new_string
}

func main() {

	var Events_slice = make([]event.Event, 0, 5)
	var Task_slice = make([]task.Task, 0, 5)
	var cmd string
	var headliner string
	var task_text string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите команду: ")
		a := scanner.Scan()
		if !a {
			fmt.Println("Ошибка ввода -", scanner.Err())
			event.AddEvent(&Events_slice, event.NewEvent("Ошибка ввода", scanner.Text(), time.Now(), time.Now().String()))
			return
		}
		text := scanner.Text()

		sl_text := strings.Fields(text)

		if len(sl_text) == 1 {
			cmd = strings.ToLower(sl_text[0])
		} else if len(sl_text) == 2 {
			cmd = strings.ToLower(sl_text[0])
			headliner = sl_text[1]
		} else if len(sl_text) > 2 {
			cmd = strings.ToLower(sl_text[0])
			headliner = sl_text[1]
			task_text = slice_to_text(sl_text[2:])
		}

		if cmd == "add" {
			if len(sl_text) <= 2 {
				fmt.Println("Введена неверная команда - нужно add {заголовок} {текст задачи}")
				event.AddEvent(&Events_slice, event.NewEvent(text, "Введена неверная команда - нужно add {заголовок} {текст задачи}", time.Now(), time.Now().String()))
				continue
			}
			cur_time := time.Now()
			task.AddTask(&Task_slice, task.NewTask(curId, headliner, task_text, cur_time), &Events_slice)
			event.AddEvent(&Events_slice, event.NewEvent(text, "", time.Now(), time.Now().String()))
			nextId()
		} else if cmd == "list" {
			if len(sl_text) != 1 {
				fmt.Println("Введена неверная команда - ввести list")
				event.AddEvent(&Events_slice, event.NewEvent(text, "Введена неверная команда - ввести list", time.Now(), time.Now().String()))
				continue
			}
			task.GetAllTasks(Task_slice)
			event.AddEvent(&Events_slice, event.NewEvent(text, "", time.Now(), time.Now().String()))
		} else if cmd == "done" {
			if len(sl_text) != 2 {
				fmt.Println("Неверная команда - ввести done {заголовок}")
				event.AddEvent(&Events_slice, event.NewEvent(text, "Неверная команда - ввести done {заголовок}", time.Now(), time.Now().String()))
				continue
			}
			task.FinishTask(&Task_slice, headliner)
			event.AddEvent(&Events_slice, event.NewEvent(text, "", time.Now(), time.Now().String()))
		} else if cmd == "del" {
			if len(sl_text) != 2 {
				fmt.Println("Неверная команда - ввести del {заголовок}")
				event.AddEvent(&Events_slice, event.NewEvent(text, "Неверная команда - ввести del {заголовок}", time.Now(), time.Now().String()))
				continue
			}
			task.DelTask(&Task_slice, headliner)
			event.AddEvent(&Events_slice, event.NewEvent(text, "", time.Now(), time.Now().String()))
		} else if cmd == "help" {
			if len(sl_text) != 1 {
				fmt.Println("неверная команда - help")
				event.AddEvent(&Events_slice, event.NewEvent(text, "неверная команда - help", time.Now(), time.Now().String()))
				continue
			}
			fmt.Println("help / add / del / list / done / exit / events")
			event.AddEvent(&Events_slice, event.NewEvent(text, "", time.Now(), time.Now().String()))
		} else if cmd == "exit" {
			if len(sl_text) != 1 {
				fmt.Println("Неверная команда - exit")
				event.AddEvent(&Events_slice, event.NewEvent(text, "Неверная команда - exit", time.Now(), time.Now().String()))
				continue
			}
			fmt.Println("До свидания!")
			return
		} else if cmd == "events" {
			if len(sl_text) != 1 {
				fmt.Println("введена неверная команда - events")
				event.AddEvent(&Events_slice, event.NewEvent(text, "введена неверная команда - events", time.Now(), time.Now().String()))
				continue
			}
			event.AddEvent(&Events_slice, event.NewEvent(text, "", time.Now(), time.Now().String()))
			event.GetAllEvents(&Events_slice)

		} else {
			fmt.Println("Неизвестная команда")
			event.AddEvent(&Events_slice, event.NewEvent(text, "Неизвестная команда", time.Now(), time.Now().String()))
			continue
		}
	}

}
