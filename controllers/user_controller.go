package controllers

import (
	"catering/config"
	"catering/helpers"
	"catering/middleware"
	"catering/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsers(e echo.Context) error {
	var Users []models.User
	err := config.DB.Find(&Users).Error

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messsage": err.Error(),
		})

	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "success",
		"data":    Users,
	})

}

func CreateUser(e echo.Context) error {
	user := models.User{}
	e.Bind(&user)
	hashPass, errors := helpers.HashPassword(user.Password)
	user.Password = string(hashPass)
	err := config.DB.Save(&user).Error

	if errors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "hashing password",
		})
	}

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "while saving obect user",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "User created",
		"data":    user,
	})
}

func ShowUser(e echo.Context) error {
	User := models.User{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Where("id=?", id).Find(&User).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "User found",
		"data":    User,
	})

}

func UpdateUser(e echo.Context) error {
	User := models.User{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Find(&User, id).Error
	e.Bind(&User)
	config.DB.Save(&User)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "User updates",
		"data":    User,
	})

}

func DeleteUser(e echo.Context) error {

	User := models.User{}
	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Where("id=?", id).Find(&User).Error
	config.DB.Delete(&User)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "User Deleted",
		"data":    User,
	})
}

func Registration(e echo.Context) error {
	user := models.User{}
	// email := e.FormValue("email")
	// pass := e.FormValue("password")
	// name := e.FormValue("full_name")
	// mobile := e.FormValue("mobile")
	// address := e.FormValue("address")

	e.Bind(&user)
	hashPass, errors := helpers.HashPassword(user.Password)

	if errors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "eror generate password",
			"error":   errors.Error,
		})
	}

	newPass := string(hashPass)
	newUser := models.User{
		FullName: user.FullName,
		Mobile:   user.FullName,
		Address:  user.Address,
		Email:    user.Email,
		Password: newPass,
	}
	err := config.DB.Save(&newUser).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "eror save user",
			"error":   err.Error,
		})
	}
	token, errorToken := middleware.CreateToken(user.ID, user.IsAdmin)
	if errorToken != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "eror generate token",
			"error":   errors.Error,
		})
	}
	userResponse := models.UserResponse{
		UserName: user.FullName,
		Mobile:   user.Mobile,
		Email:    user.Email,
		Token:    token,
	}

	e.Response().Header().Set("Authorization", "Bearer "+token)

	response := helpers.Apiresponse("Registration success", http.StatusOK, "success", userResponse)
	return e.JSON(http.StatusOK, response)
}

func Login(e echo.Context) error {
	user := models.User{}
	e.Bind(&user)
	findUser, err := FindUserByEmail(user.Email)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error sql",
			"error":   err.Error(),
		})
	}
	match, error := helpers.CheckPasswordHash(user.Password, findUser.Password)
	if !match {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "your password not match",
			"error":   error.Error(),
		})
	}

	token, errors := middleware.CreateToken(user.ID, user.IsAdmin)

	if errors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error generate token",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{
		UserName: findUser.FullName,
		Mobile:   findUser.Mobile,
		Email:    findUser.Email,
		Token:    token,
	}
	//set header
	e.Response().Header().Set("Authorization", "Bearer "+token)
	response := helpers.Apiresponse("Login Success Now You Can Go to The Dashboard", http.StatusOK, "success", userResponse)
	return e.JSON(http.StatusOK, response)
}

func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := config.DB.Where("email=?", email).Find(&user).Error

	if err != nil {
		return user, nil
	}

	return user, nil

}
