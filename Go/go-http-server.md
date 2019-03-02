# Golang HTTP Server

## Simple HTTP Server

```go
package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net"
)

func main() {
    l, err := net.Listen("tcp", "0.0.0.0:8080")
    if err != nil {
        panic(err)
    }

    for {
        conn, _ := l.Accept()
        go handle(conn)
    }
}

var get []byte
var sleep []byte

func init() {
    get = []byte("GET / HTTP/1.1\r\n")
}

func handle(conn net.Conn) {
    // net.http 패키지에서는 32kb(32*1024 bytes) sync.Pool을 사용함
    buffer := make([]byte, 1024)
    conn.Read(buffer)

    var statusline, filename string
    if bytes.HasPrefix(buffer, get) {
        statusline, filename = "HTTP/1.1 200 OK\r\n\r\n", "hello.html"
    } else {
        statusline, filename = "HTTP/1.1 404 NOT FOUND\r\n\r\n", "404.html"
    }

    content, _ := ioutil.ReadFile(filename)
    response := fmt.Sprintf("%s%s", statusline, content)
    conn.Write([]byte(response))
    conn.Close()
}
```

## `net/http` package

Go 표준 패키지 `net/http`에서 웹서버 기능을 제공한다: <https://godoc.org/net/http>

### ListenAndServe()

`ListenAndServe()` 메소드는 지정한 포트로 웹 서버를 열고 Request를 고루틴에 할당한다.

`ListenAndServe()` 메소드는 두 개의 인자를 받는다

- Listen을 수행할 포트
- ServeMux (nil인 경우 DefaultServeMux)

ServeMux는 기본적으로 HTTP Request Router (혹은 Multiplexor)인데, 별도의 ServeMux를 만들어 Routing 부분을 다르게 정의할 수 있다.

### http.Handle()

DefaultServeMux를 사용하는 경우, `Handle()` 혹은 `HandleFunc()`을 사용하여 라우팅을 정의한다.

`http.Handle()` 메소드 인자는 두 개이다

- URL/URL 패턴
- http.Handler 인터페이스를 구현한 객체

http.Handler 인터페이스의 `ServeHTTP()` 메소드는 HTTP Response.Writer와 HTTP Request를 파라미터로 받는다.

```go
package main

import (
    "net/http"
)

func main() {
    http.Handle("/", new(testHandler))
    http.ListenAndServe(":8080", nil)
}

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type testHandler struct {
    http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte(req.URL.Path))
}
```

### Static 파일 핸들러

Request URL 패스에 표시된 정적 파일을 서버 상의 특정 폴더(wwwroot) 에서 읽어 들여 파일 내용을 전달한다.

파일내용을 그냥 전달하면 기본값인 text/plain으로 전송되므로, Content-Type을 Response 헤더에 추가한 후 응답한다.

```go
package main

import (
    "io/ioutil"
    "net/http"
    "path/filepath"
)

func main() {
    http.Handle("/", new(staticHandler))
    http.ListenAndServe(":8080", nil)
}

type staticHandler struct {
    http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    localPath := "wwwroot" + req.URL.Path
    content, err := ioutil.ReadFile(localPath)
    if err != nil {
        w.WriteHeader(404)
        w.Write([]byte(http.StatusText(404)))
        return
    }

    contentType := getContentType(localPath)
    w.Header().Add("Content-Type", contentType)
    w.Write(content)
}

func getContentType(localPath string) string {
    var contentType string
    ext := filepath.Ext(localPath)

    switch ext {
    case ".html":
        contentType = "text/html"
    case ".css":
        contentType = "text/css"
    case ".js":
        contentType = "application/javascript"
    case ".png":
        contentType = "image/png"
    case ".jpg":
        contentType = "image/jpeg"
    default:
        contentType = "text/plain"
    }

    return contentType
}
```
