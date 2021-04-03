package main

import (
	"net/http"
	"time"
	"github.com/labstack/echo/v4"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type todo struct {
	ID int64 `json:"id" gorm:"primaryKey"`
	Content string `json:"content"`
	CreateAt time.Time `json:"createAt" gorm:"index"`
	CompleteAt *time.Time `json:"completeAt,omitempty" gorm:"index"`
}

var (
	db *gorm.DB
)

func createTodo(ctx echo.Context) error {
	var binder struct {
		Content string `json:"content"`
	}
	err := ctx.Bind(&binder)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	item := todo{
		Content: binder.Content,
		CreateAt: time.Now(),
	}
	db.Create(&item)
	return ctx.NoContent(http.StatusOK)
}

func getAllTodo(ctx echo.Context) error {
	var res []todo
	db.Find(&res)

	return ctx.JSON(http.StatusOK, echo.Map{
		"items": res,
	})
}

func updateTodo(ctx echo.Context) error {
	var binder struct {
		IsComplete bool `json:"isComplete"`
	}
	err := ctx.Bind(&binder)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	var id int64
	echo.PathParamsBinder(ctx).Int64("id", &id)

	var item todo
	db.First(&item, id)
	if item.ID == 0 {
		return ctx.NoContent(http.StatusNotFound)
	}

	if binder.IsComplete {
		now := time.Now()
		item.CompleteAt = &now
	} else {
		item.CompleteAt = nil
	}
	db.Save(&item)

	return ctx.NoContent(http.StatusOK)
}

func deleteTodo(ctx echo.Context) error {
	var id int64
	echo.PathParamsBinder(ctx).Int64("id", &id)

	var item todo
	db.First(&item, id)
	if item.ID == 0 {
		return ctx.NoContent(http.StatusNotFound)
	}

	db.Delete(&todo{}, id)
	return ctx.NoContent(http.StatusOK)
}

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&todo{})
	e := echo.New()
	e.Static("/", "./public")
	e.POST("/api/todo", createTodo)
	e.GET("/api/todo", getAllTodo)
	e.PATCH("/api/todo/:id", updateTodo)
	e.DELETE("/api/todo/:id", deleteTodo)
	e.Logger.Fatal(e.Start(":8081"))
}