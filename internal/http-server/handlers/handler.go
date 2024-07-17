package handlers

import (
	"encoding/json"
	"main/internal/scheduler"
	"net/http"
)

type handlerScheduler struct {
	scheduler scheduler.IScheduler
}

func (h *handlerScheduler) GetScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var promptString string
	if err := json.NewDecoder(r.Body).Decode(&promptString); err != nil {

	}

	//TODO: добавить запись в базу данных сформированных расписаний для чтобы потом отдавать от туда расписание по дате

	newSchedule, err := h.scheduler.PromptToSchedule()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(newSchedule)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
