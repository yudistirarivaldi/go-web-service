package controller

import (
	"net/http"
	"strconv"
	"web-service/formatter"
	"web-service/model"
	"web-service/utils"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} formatter.UserFormatter
// @Failure 500 {object} utils.ErrorResponseFormat
// @Router /users [get]
func GetUsers(c *gin.Context) {
    users, err := model.GetAllUsers()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }
    utils.SuccessResponse(c, formatter.FormatUsers(users), "Users fetched successfully")
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get a user by their unique ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"  // Parameter path ID pengguna
// @Success 200 {object} model.User // Tipe data respon sukses (sesuaikan dengan tipe data yang dikembalikan)
// @Failure 400 {object} utils.ErrorResponseFormat // Tipe data respon error untuk ID yang tidak valid
// @Failure 404 {object} utils.ErrorResponseFormat // Tipe data respon error untuk pengguna yang tidak ditemukan
// @Router /users/{id} [get]
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

    validationErrors := utils.ValidateStruct(&user)
    if validationErrors != nil {
        c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
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