package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/omaciel/gobutsu/db/sqlc"
)

type createTestCaseRequest struct {
	Classname    string  `json:"classname" binding:"required"`
	Filename     string  `json:"filename"`
	Linenumber   int32   `json:"linenumber"`
	Testcasename string  `json:"testcasename" binding:"required"`
	Duration     float64 `json:"duration"`
}

func(server *Server) createTestCase(ctx *gin.Context) {
	var req createTestCaseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTestCaseParams{
		Classname:    req.Classname,
		Filename:     req.Filename,
		Linenumber:   req.Linenumber,
		Testcasename: req.Testcasename,
		Duration:     req.Duration,
	}

	testcase, err := server.app.CreateTestCase(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, testcase)
}

type getTestCaseRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func(server *Server) getTestCase(ctx *gin.Context) {
	var req getTestCaseRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	testcase, err := server.app.GetTestCase(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, testcase)
}

type listTestCaseRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func(server *Server) listTestCase(ctx *gin.Context) {
	var req listTestCaseRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTestCasesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID -1) * req.PageSize,
	}

	testcases, err := server.app.ListTestCases(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, testcases)
}