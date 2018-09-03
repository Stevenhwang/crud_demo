package controllers

import (
	"awesomeProject/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func ListUsers(c echo.Context) error {
	users := new([]models.User)
	models.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func FindUser(c echo.Context) error {
	user := new(models.User)
	uid, _ := strconv.Atoi(c.Param("id"))
	models.DB.First(&user, uid)
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	models.DB.Create(&user)
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	newUser := new(models.User)
	oldUser := new(models.User)
	if err := c.Bind(newUser); err != nil {
		return err
	}
	uid, _ := strconv.Atoi(c.Param("id"))
	models.DB.First(&oldUser, uid)
	models.DB.Model(&oldUser).Update("Name", newUser.Name)
	return c.JSON(http.StatusOK, oldUser)
}

func DeleteUser(c echo.Context) error {
	user := new(models.User)
	uid, _ := strconv.Atoi(c.Param("id"))
	models.DB.First(&user, uid)
	models.DB.Delete(&user)
	return c.JSON(http.StatusOK, "hello, world")
}
