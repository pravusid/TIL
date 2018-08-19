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

## 직렬화와 역직렬화

직렬화(serialization)란 객체의 상태를 보관이나 전송 가능한 상태로 변환하는 것을 말한다.
직렬화의 반대로 보관되거나 전송받은 데이터를 다시 객체로 복원하는 것을 역직렬화(deserialization)라 한다.

직렬화는 보조 기억장치에 저장/불러오기, 네트워크로 데이터 전송, Remote Procedure Call등에 사용된다.

### JSON

#### JSON 직렬화 및 역직렬화

Go의 `json.Marshal` 함수로 직렬화를 할 수 있다

```go
func Example_arshallJSON() {
    t := Task{
        "Laundry",
        DONE,
        NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
    }
    b, err := json.Marshal(t)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(string((b)))
    // Output:
    // {"Title":"Laundry", "Status":2,"Deadline":"2018-08-19T16:40:00Z"}
}
```

`json.Unmarshal()` 함수로 역직렬화를 할 수 있다

```go
func Example_unmarshalJSON() {
    b := []byte(`{"Title":"Laundry", "Status":2,"Deadline":"2018-08-19T16:40:00Z"}`)
    t := Task()
    err := json.Unmarshal(b, &t)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(t.Title)
    fmt.Println(t.Status)
    fmt.Println(t.Deadline.UTC())
    // Output:
    // Laundry
    // 2
    // 2018-08-19T16:40 +0000 UTC
}
```

#### JSON 태그

기본 직렬화 필드이름을 변경하려면 json 태그를 사용하면된다. json태그를 JSON라이브러리가 읽어서 처리한다.

```go
type MyStruct struct {
    Title       string  `json:"title"`
    Internal    string  `json:"-"`
    Value       int64   `json:"value,omitempty"`
    Id          int64   `json:",string"`
}
```

자바스크립트에서 숫자형은 8바이트 실수형이므로 정수값은 53비트를 넘어가면 정확도가 떨어진다.
따라서 64비트 정수형을 JSON으로 주고 받을 때는 `string`으로 주고받는것이 좋다.

#### 구조체가 아닌 자료형 처리

구조체가 아닌 배열을 직렬화/역직렬화 하기 위해서 JSON 라이브러리를 사용할 수 있다.

```go
func Example_mapMarshalJSON() {
    b, _ := json.Marshal(map[string]string{
        "Name": "Kim",
        "Age": "20",
    })
    fmt.Println(string(b))
    // Output:
    // {"Age":"20","Name":"Kim"}
}
```

맵은 순서가 없지만 JSON 라이브러리는 키를 정렬하여 출력하므로 Age가 먼저 출력된다.

JSON에 이용되는 맵은 키가 문자열형이어야 한다.
값으로 아무 자료형을 담으려면 `interface{}` 자료형을 쓰면된다.

```go
func Example_mapMarshalJSON() {
    b, _ := json.Marshal(map[string]interface{}{
        "Name": "Kim",
        "Age": 20,
    })
    fmt.Println(string(b))
    // Output:
    // {"Age":20,"Name":"Kim"}
}
```

#### JSON 필드 조작

구조체에서 특정 필드를 빼고 직렬화 할 때: 구조체 내장 이용

```go
type Fields struct {
    VisibleField    string  `json:"visibleField"`
    InvisibleField  string  `json:"invisibleField"`
}

func ExampleOmitFields() {
    f := &Fields{"a", "b"}
    b, _ := json.Marshal(struct {
        *Fields
        InvisibleField  string  `json:"invisibleField,omitempty"`
        Additional      string  `json:"additional,omitempty"`
    }{Fields: f, Additional: "c"})
    fmt.Println(string(b))
    // Output:
    // {"VisibleField":"a","Additional":"c"}
}
```

### Gob

Gob은 Go언어에서 기본으로 제공하는 직렬화 방식이다.
주고 받는 코드가 모두 Go로 되어 있다면 사용을 고려해볼 수 있다.

다음은 맵을 인코딩 한 다음 한 줄에 16바이트씩 16진수로 출력하고 이를 다시 복원하는 예제이다

```go
func Example_gob() {
    var b bytes.Buffer
    enc := gob.NewEncoder(&b)
    data := map[string]string{"N": "J"}
    if err := enc.Encode(data); err != nil {
        fmt.Println(err)
    }
    const width = 16
    for start := 0; start < len(b.Bytes()); start += width {
        end := start + width
        if end > len(b.Bytes()) {
            end = len(b.Bytes())
        }
        fmt.Printf("%x\n", b.Bytes()[start:end])
    }
    dec := gob.NewDecoder(&b)
    var restored map[string]string
    if err := dec.Decode(&restored); err != nil {
        fmt.Println(err)
    }
    fmt.Println(restored)
}
```

## 인터페이스

구조체가 자료의 묶음이라면, 인터페이스는 메소드의 묶음이다.

Go에서 인터페이스에 이름을 붙일때는 주로 인터페이스 메소드 이름에 `er`을 붙인다.

### 인터페이스 정의

인터페이스로 정의한 두 메소드를 정의하고 있는 자료형은 해당 인터페이스로 사용할 수 있다.

```go
type MyInterface interface {
    Method1()
    Method2(i int) error
}
```

구조체의 내장과 비슷하에 여러 인터페이스를 합칠 수 있다

```go
type ReadWriter {
    io.Reader
    io.Writer
}
```

### 커스텀 프린터

`Print` 계열 함수들이 문자열이 아닌 자료형을 출력하게 하려면 `String()` 함수를 정의해주면 된다.

`fmt.Stringer` 인터페이스는 `func String() string` 메소드를 갖고 있다.

```go
fuc (t Task) String() string {
    check := "v"
    if t.Status != DONE {
        check = " "
    }
    return fm.tSprint("[%s] %s %s", check, t.Title, t.Deadline)
}
```

이와 같이 정의하면, `Print` 계열의 함수들이 `Stringer` 인터페이스인 경우 `String` 메소드를 호출한다

```go
func ExampleTask_String() {
    fmt.Println(Task{"Laundry", DONE, nil})
    // Output:
    // [v] Laundry <nil>
}
```

자료형 변환을 이용하면 다른 구현의 출력을 하게 할 수 있다.

작업 안에 여러개의 작업이 있는 구조이다.

```go
type Task struct {
    Title       string      `json:"title,omitempty"`
    Status      status      `json:"status,omitempty"`
    Deadline    *Deadline   `json:"deadline,omitempty"`
    Priority    int         `json:"priority,omitempty"`
    SubTasks    []Task      `json:"subTasks,omitempty"`
}

type IncludeSubTasks Task

func (t IncludeSubTasks) indentedString(prefix string) string {
    str := prefix + Task(t).String()
    for _, st := range t.SubTasks {
        str += "\n" + IncludeSubTasks(st).indentedString(prefix+"  ")
    }
    return str
}

func (t IncludeSubTasks) String() string {
    return t.indentedString("")
}

func ExampleIncludeSubTasks_String() {
    fmt.Println(IncludeSubTasks(Task{
        Title:      "Laundry",
        Status:     TODO,
        Deadline:   nil,
        Priority:   2,
        subTasks:   []Task{{
            title:      "Wash",
            Status:     TODO,
            Deadline:   nil,
            Priority:   2,
            SubTasks:   []Task{
                {"Put", DONE, nil, 2, nil},
                {"Detergent", TODO, nil, 2, nil},
            },
        }, {
            Title:      "Dry",
            Status:     TODO,
            Deadline:   nil,
            Priority:   2,
            SubTasks:   nil,
        }},
    }))
    // Output:
    // [ ] Laundry <nil>
    //  [ ] Wash <nil>
    //    [v] Put <nil>
    //    [ ] Detergent <nil
    //  [ ] Dry <nil>
}
```

### 정렬과 힙

Go 언어의 `sort.Sort`에서 이용하는 정렬은 비교정렬/불안정정렬이다.

두 자료를 비교하여 어느 자료가 먼저와야하는지를 작성하면 나머지 부분은 내장된 정렬 알고리즘으로 정렬해준다.

#### 정렬 인터페이스의 구현

Go는 제네릭을 지원하지 않지만 인터페이스를 지원하므로 다양한 형태의 정렬을 수행할 수 있다.

`sort.Interface`에 정의된 인터페이스를 따르면 정렬을 할 수 있다.

```go
type Interface interface {
    // Len is the number of elements in the collection
    Len() int
    // Less reports whether the element with index i should sort before the element with index j
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j
    Swap(i, j int)
}
```

해당 자료의 자료형은 정해져 있지 않지만 인덱스인 i와 j는 정수형으로 고정되어 있다.

[대소문자 구분없는 정렬](training/sort/sorting.go)

#### 정렬 알고리즘

싱글 스레드에서 비교정렬 알고리즘은 Quicksort가 좋다.
많은 비교 정렬 알고리즘들이 O(nlog n)의 시간 복잡도를 가지지만 잘 구현된 퀵소트는 평균적으로 가장 효율적이다.

퀵소트는 평균 O(nlog n)의 복잡도를 가지지만 최악의 경우 O(n^2)가 될 수 있다.
이를 극복하기 위해 random pivot을 뽑거나 sample pivot을 사용하기도 한다.

7개 이하의 값에 대해서는 삽입 정렬이 가장 효율적이다.
삽입 정렬은 O(n^2)의 시간복잡도를 가지지만 작은 크기의 자료에 대해서는 빠른정렬보다 효과적이다.

Go 언의 `sort.Sort`에서는 기본적으로 퀵소트를 이용한다.
퀵소트의 최악의 경우를 피하기 위해서 pivot 3개를 골라서 가운데 값을 고르는 중위법을 이용한다.

그렇게 했지만 깊은 빠른정렬의 경우가 된다면 힙 정렬을 이용하며, 7개 이하의 자료에 대해서는 삽입 정렬을 이용한다.

#### 힙

힙은 자료 중에 가장 작은 값을 O(log N)의 시간복잡도로 꺼낼 수 있는 자료구조이다.

힙 알고리즘을 이용하기 위한 인터페이스는 다음과 같다.

```go
type Interface interface {
    sort.Interface
    Push(x interface{}) // add x as element Len()
    Pop() interface{}   // remove and return element Len() - 1
}
```

`heap.Interface`는 `sort.Interface`를 내장하고 있으므로 총 5개의 메소드가 구현되어야 한다.

[대소문자 구분없는 정렬된 Heap](training/sort/heap.go)

단순히 정렬을 하는 것이라면 굳이 힙 정렬을 쓸필요는 없다.
힙 정렬의 시간 복잡도는 O(nlog n)이지만 일반적으로 퀵소트보다 느리고 메모리를 랜덤 엑세스 하므로 캐시를 효율적으로 사용할 수 없다.

그러나 무작위 퀵소트나 합병 정렬은 정렬이 모두 끝난 뒤에 자료를 이용할 수 있지만 힙 정렬은 보다 일찍 자료를 사용할 수 있다.

정렬이 다 끝나지 않은 상태에서 졍렬이 완료된 자료를 사용할 수 있는 것은 선택정렬의 특징인데, 힙 정렬또한 선택 정렬의 일종으로 볼 수 있다.

### 외부 의존성 줄이기

많은 경우 테스트를 할 때 외부 리소스 접근을 막고싶은 경우가 있다.
가능하면 만들어져 있는 인터페이스를 받아서 동작하게 코드를 작성하면 유연하게 코드를 작성할 수 있다.

만약 파일 이름변경과 삭제를 하는 인터페이스를 작성하고 구현한다면 다음과 같을 것이다.

```go
type FileSystem interface {
    func Rename(oldpath, newpath string) error
    func Remove(name string) error
}

type OSFileSystem struct{}

func (fs OSFileSystem) Rename(oldpath, newpath string) error {
    return os.Rename(oldpath, newpath)
}

func (fs OSFileSystem) Remove(name string) error {
    return os.Remove(name)
}

// 사용
func ManageFiles(fs FileSystem) {
    // ...
}
```

이런 방식으로 OSFileSystem을 이용하여 호출하면 실제 파일시스템을 이용하고,
테스트 용도로 가짜 파일시스템을 만들어 이용할 수 있다.

### 빈 인터페이스와 형 단언

나열된 메소드를 정의하고 있는 자료형은 인터페이스로 취급될 수 있다는 점을 생각해본다면,
비어있는 인터페이스는 아무 자료형이나 취급할 수 있다는 뜻이 된다.

그렇다면 `interface{}` 타입을 원래 자료형으로 변환하려면 어떻게 해야 할까

인터페이스는 실제로 자료형과 값을 갖고있는 구조체로 표현된다.
따라서 형변환을 할 때 자료형이 맞는지 실행시간에 검사가 일어나야 한다. 이를 형 단언(type assertion)이라 한다.

```go
func ExampleCaseInsensitive_heapString() {
    apple := CaseInsensitive([]string{
        "iPhone", "iPad", "MacBook", "AppStore",
    })
    heap.Init(&apple)
    for apple.Len() > 0 {
        popped := heap.Pop(&apple)
        s := popped.(string)
        fmt.Println(s)
    }
    // Output:
    // AppStore
    // iPad
    // iPhone
    // MacBook
}
```

`heap.Pop`은 `interface{}`형을 반환한다.
반환된 `popped`에 `.(string)`을 붙여서 값의 타입을 `string`이라고 단언한 것이다.

만일 실행시간에 단언한 타입과 다른 결과가 나온다면 panic이 발생한다.

형 단언은 빈 인터페이스 이외에 인터페이스를 실제 자료형으로 받을때도 쓸 수 있다.

```go
var r io.Reader = NewReader()
f, ok := r.(os.File)
```

`r`이 실제로 `os.File`인 경우 `f`를 해당 자료형으로 이용할 수 있다.
만약 자료형이 맞지 않으면 패닉이 발생하기 때문에 두 개의 값으로 받아 검사할 수 있다.

### 인터페이스 변환 스위치

인터페이스들이 지정하는 범위는 다양할 수 있다.

`io.WriterAt`, `io.WriterSeeker`, `io.WriterCloser` 모두 `io.Writer`를 내장하고 있는 인터페이스이다.

인터페이스를 받아서 특정자료형일 때, 혹은 조금 더 좁은 범위의 자료형일 때에 맞춰 구현을 달리 하려는 경우
자료형 스위치(type switch)를 활용할 수 있다.

`strings.Join`을 확장 구현해보자

```go
func Join(sep string, a ...interface{}) string {
    if len(a) == 0 {
        return ""
    }
    t := make([]string, len(a))
    for i := range a {
        switch x := a[i].(type) {
            case string:
                t[i] = x
            case int:
                t[i] = strconv.Itoa(x)
            case fmt.Stringer:
                t[i] = x.String()
        }
    }
    return strings.Join(t, sep)
}
```

`a[i]`를 `.(type)` 으로 형 단언을 하여 `x`에 할당하였다.
