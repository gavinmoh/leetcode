package main

import "container/heap"

type Task struct {
	UserId   int
	TaskId   int
	Priority int
	Index    int
}

type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Priority == pq[j].Priority {
		return pq[i].TaskId > pq[j].TaskId
	}

	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	task := x.(*Task)
	task.Index = n
	*pq = append(*pq, task)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	task := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	task.Index = -1
	*pq = old[0 : n-1]
	return task
}

type TaskManager struct {
	pq    PriorityQueue
	tasks map[int]*Task
}

func Constructor(tasks [][]int) TaskManager {
	taskManager := TaskManager{
		pq:    PriorityQueue{},
		tasks: map[int]*Task{},
	}

	heap.Init(&taskManager.pq)

	for _, task := range tasks {
		taskManager.Add(task[0], task[1], task[2])
	}

	return taskManager
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	task := &Task{
		UserId:   userId,
		TaskId:   taskId,
		Priority: priority,
	}

	this.tasks[taskId] = task
	heap.Push(&this.pq, task)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	task := this.tasks[taskId]
	task.Priority = newPriority
	heap.Fix(&this.pq, task.Index)
}

func (this *TaskManager) Rmv(taskId int) {
	task := this.tasks[taskId]
	delete(this.tasks, taskId)
	heap.Remove(&this.pq, task.Index)
}

func (this *TaskManager) ExecTop() int {
	for this.pq.Len() > 0 {
		task := heap.Pop(&this.pq).(*Task)
		delete(this.tasks, task.TaskId)

		return task.UserId
	}

	return -1
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
