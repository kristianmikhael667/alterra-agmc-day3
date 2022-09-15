package controllers

import (
	"fmt"
	"main/middleware"
	"main/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	books = map[int]*models.Book{}
	seq   = 1
)

func CreateBooks(c echo.Context) error {
	u := &models.Book{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	books[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func GetBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, books[id])
}

func UpdateBook(c echo.Context) error {
	u := new(models.Book)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	books[id].Author = u.Author
	books[id].Title = u.Title
	return c.JSON(http.StatusOK, books[id])
}

func DeleteBooks(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(books, id)
	return c.JSON(http.StatusOK, "Success Delete Book")
}

func GetAllBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, e := middleware.BasicAuth(user.Email, user.Password)
	token, e := middleware.CreateToken(int(user.ID))

	if e != nil {
		fmt.Println(e)
	}

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	if users == false {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Your Email Password is wrong",
			"status":  users,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success Login",
		"users":  users,
		"token":  token,
	})
}
