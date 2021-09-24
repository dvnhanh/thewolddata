package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type message struct {
	ReturnCode int         `JSON:"returncode"`
	Data       interface{} `JSON:"data"`
	Timestamp  int64       `JSON:"timestamp"`
}

func buidSuccessMessage(data interface{}) message {
	if data == nil {
		data = ""
	}
	return message{
		ReturnCode: 0,
		Data:       data,
		Timestamp:  time.Now().Unix(),
	}
}

func buidErrorMessage(statusCode int, err error) (int, message) {
	returncode := -1
	data := err.Error()

	sqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		statusCode = http.StatusInternalServerError
		data = "internal error"
		returncode = int(sqlErr.Number)
	}

	return statusCode, message{
		ReturnCode: returncode,
		Data:       data,
		Timestamp:  time.Now().Unix(),
	}
}

func (p *httpServer) ping(c *gin.Context) {
	c.JSON(http.StatusOK, buidSuccessMessage("pong"))
}

func (p *httpServer) register(c *gin.Context) {
	var req ReqRegister
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if len(req.Email) == 0 {
		c.JSON(buidErrorMessage(http.StatusBadRequest, errors.New("invalid email")))
		return
	}

	if len(req.Password) > 15 || len(req.Password) == 0 {
		c.JSON(
			buidErrorMessage(
				http.StatusBadRequest,
				errors.New("maximum is 15 and minimum is 10. Please try to again!"),
			),
		)
		return
	}

	if err := p.svc.Register(req.Email, req.Password); err != nil {
		c.JSON(buidErrorMessage(
			http.StatusInternalServerError,
			err,
		))
		return
	}
	c.JSON(http.StatusOK, buidSuccessMessage(nil))
}
