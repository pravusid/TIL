# Go 구조체 및 인터페이스

## 구조체

필드의 모음을 구조체라고 한다.

배열이 서로 같은 자료형의 자료들을 묶어놓은 것이라면, 구조체는 서로 다른 자료형의 자료들도 묶을 수 있다.

### 구조체 사용법

```go
type Task struct {
    title   string
    done    bool
    due     *time.Time
}

myTask := Task{
    "laundry",
    false,
    nil,
}
```

구조체 값 할당시 원하는 필드만 값을 넣을 수도 있다. 값이 없는 필드는 기본값으로 설정된다.

```go
myTask := Task{
    title: "laundry",
}
```

구조체 선언과 동시에 값을 할당할 수도 있다.

```go
var task = struct {
    title   string
    done    bool
    due     *time.Time
}{"laundry", false, nil}
```

### const와 iota

Go에서는 enum 자료형이 없고 대신 상수를 사용한다.

```go
const {
    UNKNOWN status = 0
    TODO    status = 1
    DONE    status = 2
}
```

0, 1, 2를 순서대로 붙이는 것이 번거롭다면 `iota`를 사용할 수 있다.

```go
const {
    UNKNOWN status = iota
    TODO    status = iota
    DONE    status = iota
}
```

`iota`를 사용할 때는 한번만 써도 된다

```go
const {
    UNKNOWN status = iota
    TODO
    DONE
}
```

단순히 0부터 순차적으로 사용하는 것 이외의 방법으로 사용할 수도 있다.

```go
type ByteSize float64

const {
    _   = iota // ignore first value
    KB  ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
}
```

KB는 2^10 인 1024가 되고 다음 MB는 2^20이 된다

### 테이블 기반 테스트

Go 언어의 testing 패키지에서는 다른 언어들 처럼 assertion을 이용한 유닛테스트를 지원하지 않는다.

이럴 때 구조체와 배열을 이용하여 테이블 기반 테스트를 할 수 있다.

```go
func TestFib(t *testing.T) {
    cases := []struct {
        in, want int
    }{
        {0, 0},
        {5, 5},
        {6, 8},
    }
    for _, c := range cases {
        got := seq.Fib(c.in)
        if got != c.want {
            t.Errorf("Fib(%d) == %d, want %d", c.in, got, c.want)
        }
    }
}
```

fmt 패키지에 있는 `Sprintf` 함수의 테이블 기반 테스트 예제이다.

```go
var flagtests = []struct {
    in  string
    out string
}{
    {"%a", "[%a]"},
    {"%-a", "[%-a]"},
    {"%+a", "[%+a]"},
    {"%#a", "[%#a]"},
    {"% a", "[% a]"},
    {"%0a", "[%0a]"},
    {"%1.2a", "[%1.2a]"},
    {"%-1.2a", "[%-1.2a]"},
    {"%+1.2a", "[%+1.2a]"},
    {"%-+1.2a", "[%-+1.2a]"},
    {"%-+1.2abc", "[%-+1.2abc]"},
    {"%-1.2abc", "[%-1.2abc]"},
}

func TestFlagParser(t *testing.T) {
    var flagprinter flagPrinter
    for _, tt := range flagtests {
        s := Sprintf(tt.in, &flagprinter)
        if s != tt.out {
            t.Errorf("Sprintf(%q, &flagprinter) => %q, want %q", tt.in, s, tt.out)
        }
    }
}
```

위의 예제는 총 12가지 사례에 대한 테스트를 하고 있다. 복잡한 테스트의 경우 인덱스 번호를 붙여 추적하는 것이 좋다.

### 구조체 내장

Go 언어의 구조체는 여러 자료형의 필드를 가질 수 있다는 점이 가장 중요하다.

Deadline 이라는 자료형 예제를 만들어보자

```go
type Deadline time.Time

func (d *Deadline) OverDue() bool {
    return d != nil && time.Time(*d).Before(time.Now())
}

func ExampleDeadline_OverDue() {
    d1 := Deadline(time.Now().Add(-4 * time.Hour))
    d2 := Deadline(time.Now().Add(4 * time.Hour))
    fmt.Println(d1.OverDue())
    fmt.Println(d2.OverDue())
    // Output:
    // true
    // false
}

type Task struct {
    Title       string
    Status      status
    DeadLine    *DeadLine
}

func (t Task) OverDue() bool {
    return t.Deadline.OverDue()
}

func Example_taskTestAll() {
    d1 := Deadline(time.Now().Add(-4 * time.Hour))
    d2 := Deadline(time.Now().Add(4 * time.Hour))
    t1 := Task{"4h ago", TODO, &d1}
    t2 := Task{"4h later", TODO, &d2}
    t3 := Task{"no due", TODO, nil}
    fmt.Println(t1.OverDue())
    fmt.Println(t2.OverDue())
    fmt.Println(t3.OverDUe())
    // Output:
    // true
    // false
    // false
}
```

Task 구조체 필드중 Deadline 자료형이 있어서 메소드도 이용할 수 있다.
그러나 메소드마다 같은 이름의 메소드를 호출하는 코드를 작성해야 하는데, 이런 반복작업을 덜어주는 것이 내장기능이다.

```go
type Task struct {
    Title       string
    Status      status
    *Deadline
}
```

Task 구조체에서 Deadline 필드 이름을 생략하면, Task에 대하여 OverDue 메소드를 작성할 필요가 없다.
Task가 내장하고 있는 *Deadline 자료형은 자료형의 이름과 같은 Deadline 이라는 필드가 되고, 정의된 메소드도 바로 호출할 수 있다.

구조체 내장은 장점이 있지만, 직렬화/역직렬화 시 결과가 바뀌어 버리는 문제점도 있다.

다음과 같이 Deadline 자료형을 구조체로 바꾸어 time.Time을 내장해보자

```go
type Deadline struct {
    time.Time
}

func NewDeadline(t time.Time) *Deadline {
    return &Deadline{t}
}

type Task struct {
    Title       string
    Status      status
    Deadline    *Deadline
}
```

구조체를 내장하면 내장된 구조체에 들어있는 필드에 바로 접근가능하게 된다.
따라서 구조체 내장으로 개념적으로 여러 구조체 필드를 합친 새 구조체를 만들 수 있게 된다.

이는 상속과는 달리 내부에 필드로 내장하고 있으면서 편의를 제공하는 기능일 뿐이다.

```go
type Address struct {
    City    string
    State   string
}

type Telephone struct {
    Mobile string
    Direct string
}

type Contact struct {
    Address
    Telephone
}

func ExampleContact() {
    var c Contact
    c.Mobile = "123-1234-1234"
    fmt.Println(c.Telephone.Mobile)
    c.Address.City = "Mapo"
    c.State = "Seoul"
    c.Direct = "N/A"
    fmt.Println(c)
    // Output:
    // 123-1234-1234
    // {{Mapo Seoul} {123-1234-1234 N/A}}
}
```
