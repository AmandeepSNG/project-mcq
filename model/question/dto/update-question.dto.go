package dto

import "Plateform-MCQ/model/question/schema"

type UpdateQuestionDto struct {
	QuestionId   string            `json:"questionId"`
	QuestionType string            `json:"questionType"`
	Question     string            `json:"question"`
	Options      []*schema.Options `json:"options"`
	Status       bool              `json:"status"`
}
