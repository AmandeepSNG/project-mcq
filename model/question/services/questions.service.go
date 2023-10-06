package questionService

import (
	"Plateform-MCQ/model/question/dto"
	"Plateform-MCQ/model/question/schema"
	"Plateform-MCQ/providers"
	"context"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetQuestionList() ([]*schema.QuestionSchema, error) {
	var questionlist []*schema.QuestionSchema
	client, err := providers.GetMongoDBClient()
	if err != nil {
		return questionlist, err
	}
	defer client.Disconnect(context.Background())
	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection(schema.QUESTION_SCHEMA_NAME)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return questionlist, err
	}

	for cursor.Next(context.TODO()) {
		var question schema.QuestionSchema
		err := cursor.Decode(&question)
		if err != nil {
			return nil, err
		}
		questionlist = append(questionlist, &question)
	}
	return questionlist, nil
}

func GetQuestionDetails(questionId string) (*schema.QuestionSchema, error) {
	var questionDetails *schema.QuestionSchema
	client, err := providers.GetMongoDBClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())
	whereCondtion := bson.D{bson.E{Key: "questionId", Value: questionId}}
	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection(schema.QUESTION_SCHEMA_NAME)
	collection.FindOne(context.TODO(), whereCondtion).Decode(&questionDetails)
	return questionDetails, err
}

func CreateQuestion(question dto.CreateQuestionDto) (*schema.QuestionSchema, error) {
	client, err := providers.GetMongoDBClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	// Create a new question schema
	createdQuestion := &schema.QuestionSchema{
		Id:           primitive.NewObjectID(),
		QuestionId:   uuid.New().String(),
		Question:     question.Question,
		QuestionType: question.QuestionType,
		Status:       question.Status,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Options:      question.Options,
	}

	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection(schema.QUESTION_SCHEMA_NAME)

	_, err = collection.InsertOne(context.TODO(), createdQuestion) // Assuming userService.context is valid

	if err != nil {
		return nil, err
	}
	return createdQuestion, nil
}

// CORRECT

func UpdateQuestion(questionId string, question *dto.UpdateQuestionDto) (*schema.QuestionSchema, error) {
	client, err := providers.GetMongoDBClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())
	checkExistingQuestion, err := GetQuestionDetails(questionId) // Added error handling here
	if err != nil {
		return nil, err
	}
	if checkExistingQuestion == nil {
		return nil, errors.New("question doesn't exist")
	}
	// Create a new updatedQuestion instance and populate it with the updated values
	updatedQuestion := &schema.QuestionSchema{
		Id:           checkExistingQuestion.Id,
		QuestionId:   checkExistingQuestion.QuestionId,
		QuestionType: question.QuestionType,
		Question:     question.Question,
		Options:      question.Options,
		Status:       question.Status,
		UpdatedAt:    time.Now(),
		CreatedAt:    checkExistingQuestion.CreatedAt,
	}
	whereCondition := bson.D{bson.E{Key: "questionId", Value: questionId}}
	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection(schema.QUESTION_SCHEMA_NAME)
	_, updationError := collection.UpdateOne(context.TODO(), whereCondition, bson.D{bson.E{Key: "$set", Value: updatedQuestion}})
	if updationError != nil {
		return nil, updationError
	}
	return updatedQuestion, nil
}

func DeleteQuestion(questionId string) error {
	client, err := providers.GetMongoDBClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	checkExistingQuestion, fetchingError := GetQuestionDetails(questionId) // Added error handling here
	if fetchingError != nil {
		return fetchingError
	}
	if checkExistingQuestion == nil {
		return errors.New("question doesn't exist")
	}
	whereCondition := bson.D{bson.E{Key: "questionId", Value: questionId}}
	collection := client.Database(os.Getenv("DATABASE_NAME")).Collection(schema.QUESTION_SCHEMA_NAME)
	_, deletionErr := collection.DeleteOne(context.TODO(), whereCondition)
	return deletionErr
}
