package main

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/labstack/echo/v4"
)

type model struct {
	ID int64
}

type todo struct {
	ID         int64      `json:"id"`      // json id라는 이름으로 마샬링시켜라
	Content    string     `json:"content"` // 소문자로 적으면 export가 되지 않는 타입이라 toJson으로 parsing을 할 때 ignore 된다.
	CreateAt   time.Time  `json:"createAt"`
	CompleteAt *time.Time `json:"completeAt,omitempty"` // 값이  null이면 무시해라 | *는 Optional의 의도
	/**
	*로 설정하지 않으면 예를 들어
	{
		a := todo{}
		a.CompleteAt == null -> uncompleted
	}
	이런 식으로 표현할 수 없게 된다.
	*/
}

/**
{
	"id": 1,
	"content": "빨래하기" // 소문자로 type을 적게되면 "Content": "빨래하기" 이렇게 되어버린다.
}
*/

var (
	// sequence
	value = make(map[int64]*todo) // 자료구조 map
	seq   = int64(0)
)

func createTodo(ctx echo.Context) error {
	var binder struct {
		Content string `json:"content"`
	}
	err := ctx.Bind(&binder) // 제대로 파싱을 못했다면
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	seq++
	value[seq] = &todo{
		ID:       seq,
		Content:  binder.Content,
		CreateAt: time.Now(),
	}

	fmt.Println(value)
	return ctx.NoContent(http.StatusOK)
}

/**
{
	"items": List<Todo>
}
*/
func getAllTodo(ctx echo.Context) error {
	res := make([]todo, 0, len(value)) // type, 미리 생성하는 크기, capacity (물리적인 사이즈)
	// make : map or slice 생성하는 함수
	// len : map or slice 값의 크기
	// java arraylist와 같음
	for key := range value {
		res = append(res, *value[key]) // 값 복사
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].CreateAt.Before(res[j].CreateAt)
	})

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

	if value[id] == nil {
		return ctx.NoContent(http.StatusNotFound)
	}

	if binder.IsComplete {
		now := time.Now()
		value[id].CompleteAt = &now
	} else {
		value[id].CompleteAt = nil
	}

	return ctx.NoContent(http.StatusOK)
}

func deleteTodo(ctx echo.Context) error {
	var id int64
	echo.PathParamsBinder(ctx).Int64("id", &id)

	if value[id] == nil {
		return ctx.NoContent(http.StatusNotFound)
	}

	delete(value, id)
	return ctx.NoContent(http.StatusOK)
}

func main() {
	e := echo.New()
	e.Static("/", "./public") // '/'를 들어갔을 때 ./public으로 serving 해라

	e.POST("/api/todo", createTodo) // func(echo.Context) error
	e.GET("/api/todo", getAllTodo)
	e.PATCH("/api/todo/:id", updateTodo)
	e.DELETE("/api/todo/:id", deleteTodo)
	// POST /api/todo -> create
	/**
	Request Body
	{
		"content" : string
	}
	*/

	// GET /api/todo -> fetch
	/**
	Response Body
	{
		"id": int64,
		"content": string,
		"createAt" : datetime,
		"completeAt": *datetime
	}
	*/

	// PATCH /api/todo/:id -> update (일부 변경)
	/**
	Request Body
	{
		isComplete: boolean
	}
	*/

	// DELETE /api/todo/:id -> delete

	e.Logger.Fatal(e.Start(":1323"))
}
