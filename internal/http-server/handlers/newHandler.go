package handlers

import (
	"main/internal/scheduler"
)

func NewHandler(scheduler scheduler.IScheduler) *handlerScheduler {
	return &handlerScheduler{
		scheduler: scheduler,
	}
}
