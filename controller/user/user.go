package user

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/SW-117/project2/library"
	"github.com/SW-117/project2/model/user"
	"github.com/gin-gonic/gin"
)

func RegisteruUser(c *gin.Context) {
	resData := &library.ResData{}

	reqData, err := getUserData(c)

	if err != nil {

		log.Default().Println("[get user data error:%v]", err)
		resData = library.ResJson(2001, "params error:"+err.Error(), nil)
		c.JSON(http.StatusOK, resData)
		return
	}

	_, err = user.RegisteruUser(reqData, c)
	if err != nil {
		log.Default().Println("[insert user to mysql error:%v]", err)
		resData = library.ResJson(5001, "mysql error:"+err.Error(), nil)
		c.JSON(http.StatusOK, resData)
		return
	}

	resData = library.ResJson(0, "", nil)
	c.JSON(http.StatusOK, resData)
	return

}

func getUserData(ctx *gin.Context) (*library.UserData, error) {
	reqData := new(library.UserData)

	reqData.Name = ctx.DefaultQuery("name", "")
	if reqData.Name == "" {
		return nil, errors.New("name can not be null")
	}

	if reqData.PassWard = ctx.DefaultQuery("pass", ""); reqData.PassWard == "" {
		return nil, errors.New("pass can not be null")
	}
	Sex := ctx.DefaultPostForm("sex", "0")
	reqData.Sex, _ = strconv.ParseInt(Sex, 10, 64)
	log.Default().Println("[user data :]", reqData)

	return reqData, nil
}
