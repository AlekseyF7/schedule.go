package models

type ResponseSchedule struct {
	WeekSchedule []DaySchedule `json:"week_schedule"`
}

type DaySchedule struct {
	DayName    string          `json:"day_name"`
	DayClasses []ClassSchedule `json:"day_classes"`
}

type ClassSchedule struct {
	ClassName     string    `json:"class_name"`
	ClassSubjects []Subject `json:"class_subjects"`
}

type Subject struct {
	SubjectName  string `json:"subject_name"`
	SubjectQueue string `json:"subject_queue"`
}
