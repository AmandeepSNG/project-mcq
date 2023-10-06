package dto

import "Plateform-MCQ/model/question/schema"

type CreateQuestionDto struct {
	QuestionType string            `json:"questionType"`
	Question     string            `json:"question"`
	Options      []*schema.Options `json:"options"`
	Status       bool              `json:"status"`
}
