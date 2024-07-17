package scheduler

import (
	"encoding/json"
	"fmt"
	"main/internal/lib/models"
)

type IScheduler interface {
	PromptToSchedule() (models.ResponseSchedule, error)
}

type Converter struct {
}

//Метод, который принимает промт, отправляет его на обработку в chatGPT через API и возвращает ответ в видe строки

func (c *Converter) PromptToSchedule() (models.ResponseSchedule, error) {
	//const op string = "internal.scheduler.scheduler.PromptToSchedule" //op == operation
	//client := openai.NewClient("your token")
	//resp, err := client.CreateChatCompletion(
	//	context.Background(),
	//	openai.ChatCompletionRequest{
	//		Model: openai.GPT3Dot5Turbo,
	//		Messages: []openai.ChatCompletionMessage{
	//			{
	//				Role:    openai.ChatMessageRoleUser,
	//				Content: prompt,
	//			},
	//		},
	//	},
	//)
	//if err != nil {
	//	return nil, fmt.Errorf("%s: execute statement: %w", op, err)
	//}

	//получаем этот json от GPT

	//respBytes, err := json.Marshal(resp.Choices[0].Message.Content)
	jsonData := []byte(`{
  "week_schedule": [
    {
      "day_name": "Понедельник",
      "day_classes": [
        {
          "class_name": "А",
          "class_subjects": [
            {
              "subject_name": "Английский язык",
              "subject_queue": "8:00"
            },
            {
              "subject_name": "Алгебра",
              "subject_queue": "8:50"
            },
            {
              "subject_name": "Русский язык",
              "subject_queue": "9:50"
            },
            {
              "subject_name": "Физкультура",
              "subject_queue": "10:50"
            },
            {
              "subject_name": "География",
              "subject_queue": "11:50"
            },
            {
              "subject_name": "Физика",
              "subject_queue": "12:40"
            },
            {
              "subject_name": "Геометрия",
              "subject_queue": "13:30"
            }
          ]
        },
        {
          "class_name": "Б",
          "class_subjects": [
            {
              "subject_name": "Русский язык",
              "subject_queue": "8:00"
            },
            {
              "subject_name": "Алгебра",
              "subject_queue": "8:50"
            },
            {
              "subject_name": "Английский язык",
              "subject_queue": "9:50"
            },
            {
              "subject_name": "Обществознание",
              "subject_queue": "10:50"
            },
            {
              "subject_name": "Биология",
              "subject_queue": "11:50"
            },
            {
              "subject_name": "Геометрия",
              "subject_queue": "12:40"
            },
            {
              "subject_name": "История",
              "subject_queue": "13:30"
            }
          ]
        }
      ]
    },
    {
      "day_name": "Вторник",
      "day_classes": [
        {
          "class_name": "А",
          "class_subjects": [
            {
              "subject_name": "География",
              "subject_queue": "8:00"
            },
            {
              "subject_name": "Алгебра",
              "subject_queue": "8:50"
            },
            {
              "subject_name": "Химия",
              "subject_queue": "9:50"
            },
            {
              "subject_name": "Физика",
              "subject_queue": "10:50"
            },
            {
              "subject_name": "Обществознание",
              "subject_queue": "11:50"
            },
            {
              "subject_name": "Английский язык",
              "subject_queue": "12:40"
            },
            {
              "subject_name": "Информатика",
              "subject_queue": "13:30"
            }
          ]
        },
        {
          "class_name": "Б",
          "class_subjects": [
            {
              "subject_name": "Литература",
              "subject_queue": "8:00"
            },
            {
              "subject_name": "Информатика",
              "subject_queue": "8:50"
            },
            {
              "subject_name": "Физика",
              "subject_queue": "9:50"
            },
            {
              "subject_name": "Алгебра",
              "subject_queue": "10:50"
            },
            {
              "subject_name": "Английский язык",
              "subject_queue": "11:50"
            },
            {
              "subject_name": "Русский язык",
              "subject_queue": "12:40"
            },
            {
              "subject_name": "Геометрия",
              "subject_queue": "13:30"
            }
          ]
        }
      ]
    },
	{
      "day_name": "Среда",
      "day_classes": [
        {
          "class_name": "А",
          "class_subjects": [
            {
              "subject_name": "Алгебра",
              "subject_queue": "8:00"
            },
            {
              "subject_name": "Литературу",
              "subject_queue": "8:50"
            },
            {
              "subject_name": "Информатика",
              "subject_queue": "9:50"
            },
            {
              "subject_name": "Физика",
              "subject_queue": "10:50"
            },
            {
              "subject_name": "Биология",
              "subject_queue": "11:50"
            },
            {
              "subject_name": "Английский язык",
              "subject_queue": "12:40"
            },
            {
              "subject_name": "История",
              "subject_queue": "13:30"
            }
          ]
        },
        {
          "class_name": "Б",
          "class_subjects": [
            {
              "subject_name": "Литература",
              "subject_queue": "8:00"
            },
            {
              "subject_name": "Физкультура",
              "subject_queue": "8:50"
            },
            {
              "subject_name": "Алгебра",
              "subject_queue": "9:50"
            },
            {
              "subject_name": "История",
              "subject_queue": "10:50"
            },
            {
              "subject_name": "Обществознание",
              "subject_queue": "11:50"
            },
            {
              "subject_name": "Биология",
              "subject_queue": "12:40"
            },
            {
              "subject_name": "География",
              "subject_queue": "13:30"
            }
          ]
        }
      ]
    },
	{
      "day_name": "Четверг",
      "day_classes": [
        {
          "class_name": "А",
          "class_subjects": [
            {
              "subject_name": "Геометрия",
              "subject_queue": "8:00"
            },
            {
              "subject_name": "Английский язык",
              "subject_queue": "8:50"
            },
            {
              "subject_name": "Русский язык",
              "subject_queue": "9:50"
            },
            {
              "subject_name": "Обществознание",
              "subject_queue": "10:50"
            },
            {
              "subject_name": "Литература",
              "subject_queue": "11:50"
            },
            {
              "subject_name": "Биология",
              "subject_queue": "12:40"
            },
            {
              "subject_name": "Алгебра",
              "subject_queue": "13:30"
            }
          ]
        },
        {
          "class_name": "Б",
          "class_subjects": [
            {
              "subject_name": "Русский язык",
              "subject_queue": "8:00"
            },
            {
              "subject_name": "Информатика",
              "subject_queue": "8:50"
            },
            {
              "subject_name": "ОБЖ",
              "subject_queue": "9:50"
            },
            {
              "subject_name": "Химия",
              "subject_queue": "10:50"
            },
            {
              "subject_name": "Английский язык",
              "subject_queue": "11:50"
            },
            {
              "subject_name": "Алгебра",
              "subject_queue": "12:40"
            },
            {
              "subject_name": "Литература",
              "subject_queue": "13:30"
            }
          ]
        }
      ]
    },
	{
      "day_name": "Пятница",
      "day_classes": [
        {
          "class_name": "А",
          "class_subjects": [
            {
              "subject_name": "История",
              "subject_queue": "8:00"
            },
            {
              "subject_name": "Обществознание",
              "subject_queue": "8:50"
            },
            {
              "subject_name": "Алгебра",
              "subject_queue": "9:50"
            },
            {
              "subject_name": "Химия",
              "subject_queue": "10:50"
            },
            {
              "subject_name": "Литература",
              "subject_queue": "11:50"
            },
            {
              "subject_name": "Биология",
              "subject_queue": "12:40"
            },
            {
              "subject_name": "Русский язык",
              "subject_queue": "13:30"
            }
          ]
        },
        {
          "class_name": "Б",
          "class_subjects": [
            {
              "subject_name": "Литература",
              "subject_queue": "8:00"
            },
            {
              "subject_name": "Алгебра",
              "subject_queue": "8:50"
            },
            {
              "subject_name": "Геометрия",
              "subject_queue": "9:50"
            },
            {
              "subject_name": "Физкультура",
              "subject_queue": "10:50"
            },
            {
              "subject_name": "Английский язык",
              "subject_queue": "11:50"
            },
            {
              "subject_name": "История",
              "subject_queue": "12:40"
            },
            {
              "subject_name": "Биология",
              "subject_queue": "13:30"
            }
          ]
        }
      ]
    }
  ]
}
`)

	var resp models.ResponseSchedule
	err := json.Unmarshal(jsonData, &resp)

	if err != nil {

	}

	fmt.Println(resp)
	return resp, nil
}
