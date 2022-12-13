package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func Add(l List, task string) List {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	return append(l, t)
}

func Complete(l List, taskNumber int) (List, error) {
	if taskNumber <= 0 || taskNumber > len(l) {
		return []item{}, fmt.Errorf("item %d does not exist", taskNumber)
	}

	l[taskNumber-1].Done = true
	l[taskNumber-1].CompletedAt = time.Now()

	return l, nil
}

func Delete(l List, taskNumber int) (List, error) {
	if taskNumber <= 0 || taskNumber > len(l) {
		return []item{}, fmt.Errorf("item %d does not exist", taskNumber)
	}

	l = append(l[:taskNumber-1], l[taskNumber:]...)

	return l, nil
}

func Save(l List, filename string) error {
	js, err := json.Marshal(l)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func Get(l *List, filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}
