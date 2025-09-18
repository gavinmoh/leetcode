package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskManager(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		tasks := [][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}}
		taskManager := Constructor(tasks)

		taskManager.Add(4, 104, 5)
		taskManager.Edit(102, 8)

		assert.Equal(t, 3, taskManager.ExecTop())

		taskManager.Rmv(101)
		taskManager.Add(5, 105, 15)

		assert.Equal(t, 5, taskManager.ExecTop())
	})

	t.Run("test case 2", func(t *testing.T) {
		tasks := [][]int{{1, 101, 8}, {2, 102, 20}, {3, 103, 5}}
		taskManager := Constructor(tasks)

		taskManager.Add(4, 104, 5)
		taskManager.Edit(102, 9)

		assert.Equal(t, 2, taskManager.ExecTop())

		taskManager.Rmv(101)
		taskManager.Add(50, 101, 8)

		assert.Equal(t, 50, taskManager.ExecTop())
	})
}
