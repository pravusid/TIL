# Go 언어 함수

코드 덩어리를 만든 다음 그것을 호출하고 귀환할 수 있는 구조를 서브루틴이라고 한다.

Go에서 이런 서브루틴을 함수라고 부른다.
함수에서는 인자로 값을 넘겨줄 수 있고 값을 돌려 받을 수 있다.

내부적으로 서브루틴은 주로 스택으로 구현한다.
일반적으로 호출이 이루어지면 스택에 현재 프로그램 카운터와 넘겨줄 인자들을 넣은 뒤
프로그램 카운터 값을 변경하여 호출될 서브루틴으로 건너 뛴다.

Go 언어는 값에 의한 호출(Call by value)만을 지원한다.
만약 함수 내에서 넘겨받은 값을 변경하더라도 함수 밖의 변수에는 영향을 주지 않는다.

함수 밖의 값을 변경하려면 해당 값이 들어 있는 주소를 넘겨서 주소에 있는 값을 변경하여
Call by reference와 비슷한 효과를 낼 수 있다.

## 값 넘겨주고 넘겨 받기

### 값 넘겨주기

[ReadFrom](training/list_id/list_id.go) 함수에서 `*[]string` 자료형으로 `lines`를 받았다.

슬라이스는 배열에 대한 포인터, 길이, 용량을 포함한 구조체이므로,
`ReadFrom` 함수가 `lines` 변수의 값을 변경하고자 하려면 슬라이스의 포인터를 받아야 한다.

물론 슬라이스의 배열에 대한 포인터를 따라가서 값을 변경할 수도 있을 것이다.

```go
func AddOne(nums []int) {
    for i := range nums {
        nums[i]++
    }
}

func ExampleAddOne() {
    n := []int{1, 2, 3, 4}
    AddOne(n)
    fmt.Println(n)
    // Output:
    // [2 3 4 5]
}
```

`ReadFrom` 함수를 살펴보면 다음과 같다

`func ReadFrom(r io.Read, lines *[]string) error`

여기에서 슬라이스 포인터를 넘겨주는 이유는 슬라이스의 값(배열 포인터, 길이, 크기)을 변경해야 하기 때문이다.

슬라이스에 새로운 값을 추가하기 위해서는 슬라이스 길이를 변경해야 한다. 용량이 부족한 경우에는 배열을 확장해야 한다.

포인터로 넘어온 값은 `*`을 앞에 붙여서 값을 참조 할 수 있고, 변수 앞에 `&`를 붙이면 해당 변수의 포인터 값을 얻을 수 있다.

### 둘 이상의 반환값

Go 언어의 함수에서 특이한 점은 둘 이상의 반환값을 둘 수 있다는 것이다.

반환값이 둘 이상일 때는 반환 타입을 괄호로 둘러싸고 쉼표로 구분한다

```go
func WriteTo(w io.Writer, lines []string) (int64, error) { }
```

값을 받을 때는 쉼표로 구분하여 반환값의 수에 맞게 받으면 된다.
값을 받을 때 버리고 싶은 값은 underscore(`_`) 문자를 이용하면 된다

```go
_, err := WriteTo(w, lines)
```

### 에러값 주고 받기

에러를 주고 받기 위한 어색한 방법 사용

1. 정상적이지 않은 결과값 반환 (개수를 돌려줄 때 음수 반환)
2. 호출하는 쪽에서 에러값을 받을 변수의 포인터나 레퍼런스를 함수로 넘겨주는 방법

Go의 관례상 에러는 가장 마지막 값으로 반환함

Panic이라는 다른 언어의 Exception과 같은 것을 제공해준다.
그러나 Go 언어의 패닉은 일반적인 에러보다 심각한 상황에서 주로 쓰인다.

다만 에러를 돌려받아 이용하는 코드를 작성하다보면 반복적인 코드가 많아지는 문제가 있다.

```go
if err := MyFunc(); err != nil { }
```

Go에서 예외를 현재 문맥에서 처리할 수 없을 때 에러를 그대로 반환할 수 있다

```go
if err := MyFunc(); err != nil {
    return nil, err
}
```

새로운 에러를 생성해아 하는 경우 `errors.New` 와 `fmt.Errorf`를 이용할 수 있다

### 명명된 결과 인자

Go에서는 돌려주는 값들 역시 넘겨받는 인자와 같은 형태로 쓸 수 있다.
돌려주는 인자들은 기본값으로 초기화 된다.

반환할 때에는 기존의 방식대로 결과값들을 `return`뒤에 쉼표로 구분하여 나열할 수도 있고,
생략하고 `return`만 쓰면 돌려주는 인자들의 값이 반환된다.

```go
func WriteTo(w io.Writer, lines []string) (n int64, err error) {
    for _, line := range lines {
        var nw int
        nw, err = fmt.Fprintln(w, lines)
        n += int64(nw)
        if err != nil {
            return
        }
        return
    }
}
```

많은 경우 명명된 결과 인자는 코드를 읽기 어렵게 만드므로, 필요하지 않을 때는 사용하지 않는것이 좋다.

### 가변인자

가변인자를 위해서는 `...` 키워드를 사용한다

```go
func WriteTo(w io.Writer, lines... string) (n int64, err error) { }
```

`lines`를 가변인자로 변경하여도 `lines`는 슬라이스가 된다

 이미 자료형이 슬라이스인 경우 가변인자를 받는 함수로 넘기려면 함수 호출시 `...` 키워드를 사용하면 된다

 ```go
lines := []string{"hello", "world", "Go"}
WriteTo(w, lines...)
 ```

## 값으로 취급되는 함수

Go 언어에서 함수는 Frist class citizen으로 분류된다.
즉, 함수역시 값으로 변수에 담길 수 있고 인자로 사용하거나 반환값이 될 수 있다는 것이다.

### 함수 리터럴

순수하게 함수의 값만 표현하려면 이름을 뺄 수 있다.
이를 함수 리터럴, 익명함수라고 부를 수 있고, 다른 언어에서 람다식으로 표현하는 것과 유사하다.

```go
func Example_funcLiteralVar() {
    printHello := func() {
        fmt.Println("hello")
    }
    printHello()
    // Output:
    // hello
}
```

### 고계함수

고계 함수 (higher order function)는 함수를 넘기고 받는 함수이다.

[예제](training/hof/hof_test.go): `ExampleReadFrom_Print()`

### 클로저 (closure)

클로저는 외부에서 선언한 변수를 함수 리터럴 내에서 접근할 수 있는 코드를 의미한다.

[예제](training/hof/hof_test.go): `ExampleReadFrom_append()`

클로저를 이용하여 함수를 호출할 때마다 증가한 값을 받을 수 있는 생성기를 만들어 보자

```go
func NewIntGenerator() func() int {
    var next int
    return func() int {
        next++
        return enxt
    }
}

func ExampleNewIntGenerator() {
    gen := NewIntGenerator()
    fmt.Println(gen(), gen(), gen(), gen(), gen())
    fmt.Println(gen(), gen(), gen(), gen(), gen())
    // Output:
    // 1 2 3 4 5
    // 6 7 8 9 10
}

func ExampleNewIntGenerator_multiple() {
    gen1 := NewIntGenerator()
    gen2 := NewIntGenerator()
    fmt.Println(gen1(), gen1(), gen1())
    fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())
    fmt.Println(gen1(), gen1(), gen1(), gen1()
    // Output:
    // 1 2 3
    // 1 2 3 4 5
    // 4 5 6 7
}
```

`NewIntGenerator` 함수는 정수를 반환하는 클로저 함수를 반환한다.
반환하는 함수 리터럴이 속해있는 범위 내의 `next` 변수에 접근하고 있다.
따라서 이 함수는 `next` 변수와 세트가 되어 작동한다.

같은 방식을 사용하여 lazy evaluation을 구현하거나 무한한 크기의 자료구조를 만들 수 있다.

### 명명된 자료형

자료형에 새로운 이름을 붙일 수 있다.

`rune` 타입은 사실 `int32`의 별칭인데, 다음과 같이 다른 이름을 붙일 수 있다.
이런 자료형을 Named Type이라고 한다.

```go
type runes []rune
type MyFunc func() int
```

정수형으로 ID를 사용하는 정점과 간선으로 이루어진 그래프를 다루는 코드를 작성해보자

```go
func NewVertexIdGenerator() func() int {
    var next int
    return func() int {
        next++
        return next
    }
}

func NewEdgeIdGenerator() func() int {
    var next int
    return func() int {
        next++
        return next
    }
}

// 간선의 ID를 받아서 새로운 간선을 생성하는 함수
func NewEdge(eid int) {
    // ...
}

func main() {
    gen1 := NewVertexIdGenerator()
    gen2 := NewEdgeIdGenerator()
    // ...
    e := NewEdge(gen1())  // 간선을 넘겨야 되는데 정점을 넘겼다면??
}
```

`main()` 함수에서의 실수를 방지하기 위해서 정점과 간선의 ID의 자료형에 서로 다른 이름을 붙일 수 있다.

```go
type VertexId int
type EdgeId int
```

위의 함수를 고쳐쓰면 다음과 같다

```go
func NewVertexIdGenerator() func() VertexId { ... }

func NewEdgeIdGenerator() func() EdgeId { ... }

func NewEdge(eid EdgeId) { ... }
```

### 명명된 함수형

Go 언어에서 함수는 일급시민으로 분류되므로 함수의 자료형 역시 사용자가 정의할 수 있다.

```go
type BinOp func(int, int) int

func OpThreeAndFour(f BinOp) {
    fmt.Println(f(3, 4))
}
```

함수를 호출할 때 익명 함수를 넘겨준다면 컴파일 오류가 발생하지 않는다.
즉, 호출시 둘다 명명되어 있는 타입일때 다른 타입이라면 오류가 발생한다.

```go
OpThreeAndFour(func (a, b int) int {
    return a + b
})
```

이를 다른말로 하면 표현이 같은 함수형이라도 이름이 다른경우 호환되지 않는다.

```go
type BinSub func(int, int) int

func BinOpToBinSub(f BinOp) BinSub {
    var count int
    return func(a, b int) int {
        fmt.Println(f(a, b))
        count++
        return count
    }
}

func ExampleBinOpToBinSub() {
    sub := BinOpToBinSub(func(a, b, int) int {
        return a + b
    })
    sub(5, 7)
    sub(5, 7)
    count := sub(5, 7)
    fmt.Println("count:", count)
    // Output:
    // 12
    // 12
    // 12
    // count: 3
}

func ExampleBinOpToBinSub_error() {
    // 컴파일 에러 발생
    sub := BinOpToBinSub(BinOpToBinSub(func(a, b int) int {
        return a + b
    }))
    // ...
}
```

### 인자고정

함수의 인자를 고정하고 싶을 때가 있다.

```go
type MultiSet map[string]int
type SetOp func(m MultiSet, val string)

// 집합에 val을 추가함
func Insert(m MultiSet, val string)
```

앞에서 만든 `ReadFrom`에 사용해 보자

```go
func BindMap(f SetOp, m MultiSet) func(val string) {
    return func(val string) {
        f(m, val)
    }
}

m: = NewMultiSet()
ReadFrom(r, BindMap(Insert, m))
```

### 패턴의 추상화

고계 함수를 이용하면 높은 수준의 추상화를 사용할 수 있다.

이전에 만들었던 생성기 예제에서 NewVertexIdGenerator, NewEdgeIdGenerator의 패턴이 동일했다.

```go
func NewVertexIdGenerator() func() VertexId {
    var next int
    return func() int {
        next++
        return next
    }
}

func NewEdgeIdGenerator() func() EdgeId {
    var next int
    return func() int {
        next++
        return next
    }
}
```

반복되는 패턴이 있다면 이를 추상화 할 수 있다.

```go
func NewIntGenerator() func() int {
    var next int
    return func() int {
        next++
        return next
    }
}

func NewVertexIdGenerator() func() VertexId {
    gen := NewIntGenerator()
    return func() VertexId {
        return VertexId(gen())
    }
}
```

[제곱근의 NewtonTransform](training/newton/newton.go)

### 자료구조에 담은 함수

Go에서 함수는 일급 시민이므로 자료구조에 담을 수도 있다.

[계산기 예제](training/calculator/calculator_func.go)

## 메소드 (Method)

코드 덩어리를 만든 다음 그것을 호출하고 반환하는 구조를 서브루틴이라고 했다.
여기에 리시버가 붙으면 Method가 된다.

자료형 T에 대하여 메소드를 호출할 때 자료형 T에 대한 리시버가 함수 이름, 즉 메소드 이름앞에 붙는다.

```go
func (recv T) MethodName(p1 T1, p2 T2) R1
```

### 단순 자료형 메소드

Go 언어에서는 모든 명명된 자료형에서 메소드를 정의할 수 있다.

```go
type VertexId int

func ExampleVertexId_print() {
    i := VertexId(100)
    fmt.Println(i)
    // Output:
    // 100
}
```

`i`는 `VertexId`형이지만 화면에 출력하면 정수형과 마찬가지로 출력된다.
다음과 같은 코드를 작성해서 이를 다른 형태로 출력해보자.

```go
func (id VertexId) String() string {
    return fmt.Sprintf("VertexId(%d)", id)
}

func ExampleVertexId_String() {
    i := VertexId(100)
    fmt.Println(i)
    // Output:
    // VertexId(100)
}
```

`i`가 `VertexId` 자료형이면 `i.String()`과 같이 메소드를 호출할 수 있다.
여기에서 `i.String()`을 호출한 것도 아닌데 결과가 나온것은 인터페이스 기능때문이다.

### 문자열 다중 집합

메소드를 활용해서 문자열 다중 집합을 구현해보자

```go
type MultiSet map[string]int

func (m MultiSet) Insert(val string) {
    m[val]++
}

func (m MultiSet) Erase(val string) {
    if m[val] <= 1 {
        delete(m, val)
    } else {
        m[val]--
    }
}

func (m MultiSet) Count(val string) int {
    return m[val]
}

func (m MultiSet) String() string {
    s := "{ "
    for val, count := range m {
        s += strings.Repeat(val+" ", count)
    }
    return s + "}"
}
```

`map[string]int`로 표현되는 명명된 자료형이 여러개일 수 있고 이들은 모두 다른 메소드를 가질 수 있다.

### 포인터 리시버

포인터 리시버는 자료형이 포인터형인 리시버이다.
리시버 역시 함수의 다른 인자들과 같이 값으로 전달되는데, 포인터로 전달해야 할 경우 포인터 리시버를 사용해야 한다.

```go
type Graph [][]int

func WriteTo(w io.Writer, adjList [][]int) error
func ReadFrom(r io.Reader, adjList *[][] error)
```

다음과 같이 메소드의 자료형을 정의할 수 있다

```go
func (adjList Graph) WriteTo(w io.Writer) error
func (adjList *Graph) ReadFrom(r io.Reader) error
```

### 공개 및 비공개

객체지향 언어들은 메소드를 `public` 혹은 `private`으로 지정하여 접근제한을 할 수 있다.

Go 언어에서는 메소드 이름을 대문자/소문자로 구분할 수 있다.
메소드 이름이 대문자로 시작하면 다른 모듈에서 보이고, 소문자로 시작하면 다른 모듈에서 보이지 않는다.

대/소문자 구분은 메소드 뿐만 아니라 모듈 전역에 정의된 자료형, 변수, 상수, 함수에 모두 적용된다.

공개된 요소에는 주석을 달도록 되어있다 (컨벤션상, godoc이 문서를 자동생성함)

## 활용

### 타이머 활용

프로그램의 수행을 잠시 멈추고 싶을 때 `time.Sleep` 함수를 사용할 수 있다. [예제](training/sleep/sleep.go)

위의 예제는 blocking 타이머이다.
`time.Timer`를 이용하면 non-blocking 타이머를 이용할 수 있다.

이런 비동기적 상황에서 사용되는 테크닉이 콜백이다. 고계 함수를 이용한 콜백은 다음과 같을 것이다.

```go
timer := time.AfterFunc(5*time.Second, func() {
    // 수행
})
timer.Stop()
```

### path/filepath 패키지

path/filepath 패키지는 파일 이름 경로를 다루는 패키지 이다.
