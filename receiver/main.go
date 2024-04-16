package receiver

import "fmt"

// add struct
type Task struct {
	Title    string
	Estimate int
}

// add receiver
func (task Task) extendEstimate() {
	task.Estimate += 10
}

// add pointer receiver
func (task *Task) extendEstimatePointer() {
	task.Estimate += 10
}

// add function
func tasks() {
	task1 := Task{
		Title:    "Task 1",
		Estimate: 10,
	}
	task1.extendEstimate()
	fmt.Printf("Task1: %v\n", task1.Estimate) // task1 の Estimate は変わらない
	task1.extendEstimatePointer()             // pointer receiver を使うと Estimate が変わる 20
	fmt.Printf("Task1: %v\n", task1.Estimate)
}
