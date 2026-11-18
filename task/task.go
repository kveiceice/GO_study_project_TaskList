package task

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	event "study_project_1/events"
	"time"

	"github.com/k0kubun/pp"
)

type Task struct {
	id               int
	headline         string
	text             string
	start            string
	actual_start     time.Time
	finished         bool
	finishTime       time.Time
	human_finishtime string
	duration         string
}

func NewTask(id int, headline string, text string, start time.Time) Task {
	if headline == "" {
		fmt.Println("Нет заголовка")
		return Task{}
	}
	if text == "" {
		fmt.Println("Нет текста")
		return Task{}
	}
	return Task{
		id:           id,
		headline:     headline,
		text:         text,
		start:        start.Format(time.RFC3339),
		actual_start: start,
		finished:     false,
	}
}

func AddTask(Tasks_slice *[]Task, task Task, ev_slice *[]event.Event) {
	*Tasks_slice = append(*Tasks_slice, task)
	fmt.Println("Был добавлен элемент с заголовком-", task.headline)
}

func DelTask(Tasks_slice *[]Task, header string) {
	var sameTasks []int
	var choiceDeleteTasks []Task
	if len(*Tasks_slice) > 0 {
		for i := 0; i < len(*Tasks_slice); i++ {
			if (*Tasks_slice)[i].headline == header {
				sameTasks = append(sameTasks, i)
				choiceDeleteTasks = append(choiceDeleteTasks, (*Tasks_slice)[i])
			}
		}
		if len(sameTasks) == 1 {
			fmt.Println("Был удален элемент с заголовком-", (*Tasks_slice)[sameTasks[0]].headline)
			*Tasks_slice = append((*Tasks_slice)[:sameTasks[0]], (*Tasks_slice)[sameTasks[0]+1:]...)
		} else if len(sameTasks) == 0 {
			fmt.Println("Не было найдена такой задачи!")
		} else {
			pp.Println(choiceDeleteTasks)
			fmt.Print("Было найдено несколько задач с похожим заголовком - введите id той задачи которую хотите удалить: ")
			scanner := bufio.NewScanner(os.Stdin)
			a := scanner.Scan()
			if !a {
				fmt.Println("Ошибка ввода")
				return
			}
			text := scanner.Text()
			defText := strings.Fields(text)
			val, error := strconv.Atoi(defText[0])
			if len(defText) != 1 {
				fmt.Println("Введено больше чем 1 значение!")
				return
			}
			if error != nil {
				fmt.Println("Введено не целочисленное значение!")
				return
			}
			for i := 0; i < len(*Tasks_slice); i++ {
				if (*Tasks_slice)[i].id == val {
					fmt.Println("Удалена задача с id =", val)
					*Tasks_slice = append((*Tasks_slice)[:i], (*Tasks_slice)[i+1:]...)
				}
			}
		}
	} else {
		fmt.Println("Сначала добавьте хотя бы одну задачу!")
	}

}

func GetAllTasks(Tasks_slice []Task) {
	if len(Tasks_slice) > 0 {
		for _, v := range Tasks_slice {
			if v.duration == "" {

				pp.Println("Id", v.id)
				pp.Println("Headline -", v.headline)
				pp.Println("Text -", v.text)
				pp.Println("Start time -", v.actual_start)
				pp.Println("Finished -", v.finished)
			} else {
				pp.Println("Id -", v.id)
				pp.Println("Headline -", v.headline)
				pp.Println("Text -", v.text)
				pp.Println("Start time -", v.actual_start)
				pp.Println("Finished -", v.finished)
				pp.Println("Finish time -", v.finishTime)
				pp.Println("Duration -", v.duration)
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Сначала добавьте хотя бы одну задачу!")
	}
}

func FinishTask(Tasks_slice *[]Task, header string) {
	if len(*Tasks_slice) > 0 {
		for i := 0; i < len(*Tasks_slice); i++ {
			if (*Tasks_slice)[i].headline == header {
				(*Tasks_slice)[i].finished = true
				t := time.Now()
				act_start := (*Tasks_slice)[i].actual_start
				(*Tasks_slice)[i].duration = t.Sub(act_start).String()[:5]
				(*Tasks_slice)[i].finishTime = t
				(*Tasks_slice)[i].human_finishtime = t.Format(time.RFC3339)
				fmt.Println("Завершена задача с заголовком " + header)
			}
		}
	} else {
		fmt.Println("Сначала добавьте хотя бы одну задачу!")
	}
}
