package controller

import (
	"net/http"
	"strconv"
	"web-service/model"
	"web-service/utils"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    users, err := model.GetAllUsers()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    utils.SuccessResponse(c, users, "Users fetched successfully")
}

func GetUserByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
        return
    }

    user, err := model.GetUserByID(id)
    if err != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "User not found")
        return
    }
    utils.SuccessResponse(c, user, "User fetched successfully")
}

func CreateUser(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    id, err := model.CreateUser(user)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    user.ID = int(id)
    utils.SuccessResponse(c, user, "User created successfully")
}

func UpdateUser(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
        return
    }
    user.ID = id

    if err := model.UpdateUser(user); err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    utils.SuccessResponse(c, user, "User updated successfully")
}

func DeleteUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
        return
    }

    if err := model.DeleteUser(id); err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    utils.SuccessResponse(c, nil, "User deleted successfully")
}