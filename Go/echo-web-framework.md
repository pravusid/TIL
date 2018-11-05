# Echo web Framework

## 시작

`go get -u github.com/labstack/echo/...`

<https://echo.labstack.com/guide>

## 동작 방식

`http.Server.Handler` 인터페이스를 구현하는 `func (e *Echo) ServeHTTP(w http.ResponseWriter, r *http.Request)`
메소드가 실질적으로 Echo에서 작업을 처리함

`serveHTTP` 메소드 내에서 각 라우트의 Handler Function을 Router의 Find 메소드로 찾아
최초에 재사용가능한 Pool로 생성된 Context중 하나의 포인터를 가져와 핸들러 함수에 넣어 실행시킴

`c := e.pool.Get().(*context) (sync.Pool)`

## HelloWorld

```go
package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}
```

## Routing

HTTP methods

```go
e.POST("/users", saveUser)
e.GET("/users/:id", getUser)
e.PUT("/users/:id", updateUser)
e.DELETE("/users/:id", deleteUser)
```

Path Variable

```go
e.GET("/users/:id", getUser)

func getUser(c echo.Context) error {
    // User ID from path `users/:id`
    id := c.Param("id")
    return c.String(http.StatusOK, id)
}
```

Query Parameter

```go
e.GET("/show", show)

func show(c echo.Context) error {
    // Get team and member from the query string
    team := c.QueryParam("team")
    member := c.QueryParam("member")
    return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}
```

## Request

## Response
