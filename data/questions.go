package data

import "math/rand"

type Question struct {
	ID          int              `json:"id"`
	AnswerCount int              `json:"answerCount"`
	Prompt      string           `json:"prompt"`
	Answers     []QuestionAnswer `json:"answers"`
}

type QuestionAnswer struct {
	ID     int    `json:"id"`
	Desc   string `json:"desc"`
	Points int    `json:"points"`
}

func GetQuestions() []*Question {
	return questionList
}

func GetNewRandQuestion() *Question {
	randQ := rand.Intn(len(questionList))
	return questionList[randQ]
}

func AddQuestion(q *Question) {
	q.ID = findNextId()
	questionList = append(questionList, q)
}

func findNextId() int {
	return questionList[len(questionList)-1].ID + 1
}

var questionList = []*Question{
	{
		ID:          1,
		AnswerCount: 4,
		Prompt:      "Cual es la respuesta?",
		Answers: []QuestionAnswer{
			{
				ID:     1,
				Desc:   "Respuesta 1",
				Points: 28,
			},
			{
				ID:     2,
				Desc:   "Respuesta 2",
				Points: 14,
			},
			{
				ID:     3,
				Desc:   "Respuesta 3",
				Points: 10,
			},
			{
				ID:     4,
				Desc:   "Respuesta 4",
				Points: 8,
			},
		},
	},
	{
		ID:          2,
		AnswerCount: 5,
		Prompt:      "Pregunta numero 2?",
		Answers: []QuestionAnswer{
			{
				ID:     1,
				Desc:   "Respuesta 1",
				Points: 28,
			},
			{
				ID:     2,
				Desc:   "Respuesta 2",
				Points: 14,
			},
			{
				ID:     3,
				Desc:   "Respuesta 3",
				Points: 10,
			},
			{
				ID:     4,
				Desc:   "Respuesta 4",
				Points: 8,
			},
			{
				ID:     5,
				Desc:   "Respuesta 5",
				Points: 3,
			},
		},
	},
}
