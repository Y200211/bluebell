package controller

import (
	"bluebell/logic"
	"bluebell/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//type VoteData struct {
//	PostID int64 `json:"post_id,string" binding:"required"`
//	Direction int `json:"direction,string" binding:"required"`
//}

func PostVoteContrtoller(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
	}
	logic.PostVote()
	ResponseSuccess(c, nil)
}
