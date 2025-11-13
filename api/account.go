package api

import (
	"net/http"

	db "menribardhi/micro-go-psql/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}
type getAccountReq struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) createAccount(context *gin.Context) {
	var req createAccountRequest
	if err := context.ShouldBindBodyWithJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	params := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	account, err := s.store.CreateAccount(context.Request.Context(), params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	context.JSON(http.StatusCreated, account)
}

func (s *Server) getAccountById(context *gin.Context) {
	var req getAccountReq
	if err := context.ShouldBindUri(&req); err != nil {
		context.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := s.store.GetAccount(context.Request.Context(),
		req.ID)
	if err != nil {
		context.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	context.JSON(http.StatusOK, account)
}

type listAccountsRequest struct {
	Page  int32 `form:"page" binding:"required,min=1"`
	Count int32 `form:"count" binding:"required,min=1,max=10"`
}

func (s *Server) listAccounts(context *gin.Context) {
	var req listAccountsRequest
	if err := context.ShouldBindQuery(&req); err != nil {
		context.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	params := db.ListAccountsParams{
		Limit:  req.Count,
		Offset: (req.Page - 1) * req.Count,
	}
	accounts, err := s.store.ListAccounts(context.Request.Context(), params)
	if err != nil {
		context.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	context.JSON(http.StatusOK, accounts)
}
