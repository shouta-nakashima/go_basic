package structDir

import (
	"fmt"
	"unsafe"
)

type Task struct {
	Title    string
	Estimate int
}

func tasks() {
	task1 := Task{
		Title:    "Task 1",
		Estimate: 10,
	}

	// Update the title of task1
	task1.Title = "Task 1 updated"
	fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)

	var Task2 Task = task1 // Copy the value of task1 to task2
	Task2.Title = "Task 2"
	fmt.Printf("Task1: %v Task2 %v\n", task1.Title, Task2.Title)

	task1p := &Task{ // Create a pointer to a Task struct
		Title:    "Task 1p",
		Estimate: 10,
	}

	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Offsetof(task1p))
	task1p.Title = "Task 1p updated" // Update the title of task1p
	fmt.Printf("task1p: %+v\n", *task1p)

	var task2p *Task = task1p // Copy the value of task1p to task2p
	task2p.Title = "Task 2p"
	fmt.Printf("task1p: %v task2p %v\n", task1p.Title, task2p.Title) // task1p and task2p point to the same memory address
}
