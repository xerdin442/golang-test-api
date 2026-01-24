package tasks

import (
	"github.com/hibiken/asynq"
	"github.com/xerdin442/api-practice/internal/config"
)

type TasksClient interface {
	Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
}

type TaskHandler struct {
	cfg *config.Config
}

func NewTaskHandler(c *config.Config) *TaskHandler {
	return &TaskHandler{cfg: c}
}
