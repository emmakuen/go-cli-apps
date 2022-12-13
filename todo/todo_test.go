package todo_test

import (
	"github.com/emmakuen/go-cli-apps/todo"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l = todo.Add(l, taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l = todo.Add(l, taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("new task should not be completed")
	}

	l, err := todo.Complete(l, 1)

	if err != nil {
		t.Errorf("couldn't mark the task complete.")
	}

	if !l[0].Done {
		t.Errorf("new task should be completed.")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}

	for _, task := range tasks {
		l = todo.Add(l, task)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("expected %q, got %q instead", tasks[0], l[0].Task)
	}

	l, err := todo.Delete(l, 2)

	if err != nil {
		t.Errorf("failed to delete a task")
	}

	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("expected %q, got %q instead", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	sourceList := todo.List{}
	targetList := todo.List{}

	taskName := "New Task"
	sourceList = todo.Add(sourceList, taskName)

	if sourceList[0].Task != taskName {
		t.Errorf("expected %q, got %q instead.", taskName, sourceList[0].Task)
	}

	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating temp file: %s", err)
	}

	defer os.Remove(tempFile.Name())

	if err := todo.Save(sourceList, tempFile.Name()); err != nil {
		t.Fatalf("error saving list to file: %s", err)
	}

	if err := todo.Get(&targetList, tempFile.Name()); err != nil {
		t.Fatalf("error getting list from file: %s", err)
	}

	if len(targetList) != 1 {
		t.Errorf("target list is empty %v", targetList)
		return
	}

	if sourceList[0].Task != targetList[0].Task {
		t.Errorf("task %q should match %q task.", sourceList[0].Task, targetList[0].Task)
	}
}
