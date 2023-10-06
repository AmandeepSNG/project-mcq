package questionController

import (
	"Plateform-MCQ/common"
	"Plateform-MCQ/helpers"
	"Plateform-MCQ/model/question/dto"
	questionService "Plateform-MCQ/model/question/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getQuestionList(ctx *gin.Context) {
	// Retrieve the list of questions
	questionList, err := questionService.GetQuestionList()
	if err != nil {
		log.Println("Error while fetching question list:", err)

		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}
	helpers.GenerateSuccessResponse(ctx, common.SUCCESS, http.StatusOK, questionList)
}

func getQuestionDetails(ctx *gin.Context) {
	questionDetails, err := questionService.GetQuestionDetails(ctx.Param("questionId"))
	if err != nil {
		log.Println("Error while fetching details of question with questionId => "+ctx.Param("questionId")+"and the error is =>", err)

		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}
	helpers.GenerateSuccessResponse(ctx, common.SUCCESS, http.StatusOK, questionDetails)
}

func createQuestion(ctx *gin.Context) {
	var requestBody dto.CreateQuestionDto
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		log.Println("error occurred while binding request body body", err)
		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}
	createdQuestion, err := questionService.CreateQuestion(requestBody)
	if err != nil {
		log.Println("error occurred while creating new question", err)

		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}
	helpers.GenerateSuccessResponse(ctx, common.SUCCESS, http.StatusOK, createdQuestion)
}

func updateQuestion(ctx *gin.Context) {
	var requestBody dto.UpdateQuestionDto
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		log.Println("error occurred while binding request body", err)
		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}

	updatedQuestion, err := questionService.UpdateQuestion(ctx.Param("questionId"), &requestBody)
	if err != nil {
		log.Println("error occurred while updating question with ID =>"+ctx.Param("questionId"), err)
		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}
	helpers.GenerateSuccessResponse(ctx, common.SUCCESS, http.StatusOK, updatedQuestion)
}
func deleteQuestion(ctx *gin.Context) {
	err := questionService.DeleteQuestion(ctx.Param("questionId"))
	if err != nil {
		log.Println("error occurred while deleting question with ID =>"+ctx.Param("questionId"), err)

		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}
	helpers.GenerateSuccessResponse(ctx, common.SUCCESS, http.StatusOK, nil)
}

func RegisterQuestionRoutes(routerGroup *gin.RouterGroup) {
	questionRouter := routerGroup.Group("/questions")

	questionRouter.GET("/list", getQuestionList)
	questionRouter.POST("/", createQuestion)
	questionRouter.GET("/:questionId", getQuestionDetails)
	questionRouter.PATCH("/:questionId", updateQuestion)
	questionRouter.DELETE("/:questionId", deleteQuestion)
}
