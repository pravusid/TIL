# Go 활용

## 에러처리

error 자료형은 인터페이스 자료형이다

```go
type error interface {
    Error() string
}
```

### 반복된 에러 처리 피하기

Must와 같은 이름의 함수를 만들어 err이 nil 값이 아닐 때 panic을 발생시킬 수 있다.

```go
func Must(i int64, err error) int64 {
    if err != nil {
        panic(err)
    }
    return i
}
```

이런 방식은 에러가 나면 프로그램이 종료되어야 하거나 반드시 에러가 발생하지 않는 경우에 사용하면 편리하다.

### panic과 recover

패닉이 발생하면 호출 스택을 타고 역순으로 올라가서 결국 프로그램이 종료된다.

이때 호출 스택에 있는 함수 f1 내부에서 defer f2()와 같이 등록이 되었다면 f1 함수가 종료되기 전에 f2()가 호출된다.
마찬가지로 이 함수를 호출한 부분으로 가서 defer를 처리하면서 패닉을 계속 상위호출자로 전파한다.

이 과정에서 더이상 패닉이 전파되지 않도록 하는 방법이 recover이다. recover는 defer 안에서만 효력이 있다.

```go
func f() int {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    g() // function panics
    return 100
}

func g() {
    panic("Panic")
}

func Example_f() {
    fmt.Println("f() =", f())
    // Output:
    // Recovered in f
    // 0
}
```

defer는 패닉이 발생했을 때도 실행되지만 패닉이 없어도 실행된다.
따라서 defer 내의 recover가 nil인경우 패닉이 발생하지 않은 경우이다.

패닉이 발생했으므로 기본값인 0이 반환된다.
만약 defer에서 원하는 값을 할당하고 싶다면 익명함수에 명명된 결과인자를 붙이면 된다.

```go
defer func() (i int) {
    if r := recover(); r != nil {
        fmt.Println("Recovered in f", r)
        i = -1
    }
}
```

## 오버로딩

Go 언어에서는 오버로딩이 지원되지 않는다.

우선 오버로딩이 필요한 유형을 나누어보자

- 자료형에 따라 다른이름 붙이기: 오버로딩을 반드시 하지 않아도 된다
- 동일한 자료형의 자료개수에 따른 오버로딩: `max(a, b)`, `max(a, b, c)`와 같은경우인데 가변인자를 사용하면 된다
- 자료형 스위치 활용: 오버로딩을 반드시 해야하는 경우에는 인터페이스로 인자를 받고 메소드 내에서 자료형 스위치로 다른 자료형에 맞춘다
- 다양한 인자 넘기기: 편의를 위한 오버로딩이다. 필요하다면 여러 값을 묶은 구조체를 넘기ㅡㄴ 것을 고려하자

인터페이스를 활용하는 것이 나은 경우도 있다.

```go
func String(int i) { .. }
func String(double d) { ... }

// 인터페이스 사용시
type Stringer interface {
    String()    string
}

type Int int
type Double float64

func (i Int) String() string { ... }
func (d Double) String() string { ... }

func ExampleString() {
    fmt.Println(Int(5).String(), Double(3.7).String())
    // Output:
    // 5
    // 3.7
}
```

### 연산자 오버로딩

Go 언어는 연산자 오버로딩을 지원하지 않는다.

연산자 오버로딩은 문제를 풀기위해서라기 보다 편의상의 기능으로 볼 수 있다.

## 템플릿 및 제네릭 프로그래밍

제네릭은 알고리즘을 표현하면서 자료형을 배제할 수 있는 프로그래밍 패러다임이다.

하지만 Go 언어는 제네릭을 지원하지 않는다.

### 유닛테스트

xUnit 스타일의 테스트에는 `assertEqual`과 같은 함수를 이용하여 값이 같은지 비교한다.

이를 대체하기 위한 방법으로 우선 if를 이용할 수 있을 것이다

```go
if expected != actual {
    t.Error("Not equal")
}
```

그러나 한 줄로 표현할 수 있는 것을 여러줄로 표현해야 하는 문제가 있다.

`reflect.DeepEqual`을 활용해서 범용적인 `assertEqual` 함수를 작성할 수도 있을 것이다.

앞에서 살펴본대로 테이블 기반 테스트를 진행할 수도 있다.

### 컨테이너 알고리즘

제네릭은 주로 컨테이너에 많이 이용된다. 이런경우 인터페이스를 활용할 수 있다.

### 자료형 메타 데이터

넘어온 자료형에 따라 다른 동작을 하게 하려면 자료형 스위치를 활용하면 된다.

그러나 자료형인지를 알아보는 것 뿐만아니라 자료형에 대한 메타데이터를 처리하고 싶을수도 있다.
이럴때 `reflect` 패키지를 이용해서 자료형의 메타데이터를 이용할 수 있다.

예를 들어 넘어온 값을 보고 해당 자료형에 따라 알맞은 자료형의 맵을 반환하는 함수를 보자

```go
func NewMap(key, value interface{}) interface{} {
    keyType := reflect.TypeOf(key)
    valueType := reflect.TypeOf(value)
    mapType := reflect.MapOf(KeyType, valueType)
    mapValue := reflect.MakeMap(mapType) // 값에 대한 메타데이터
    return mapValue.Interface()
}

fumc main() {
    m := NewMap("", 0).(map[string]int) // 사용시 형 단언해야함
}
```

`reflect`를 활용해 구조체의 필드를 확인할 수도 있다.

```go
func FieldNames(s interface{}) ([]string, error) {
    t := reflect.TypeOf(s)
    if t.Kind() != reflect.Struct {
        return nil, errors.New("FieldNames: s is not a struct!")
    }
    names := []string{}
    n := t.NumField()
    for i := 0; i < n; i++ {
        names = append(names, t.Field(i).Name)
    }
    return names, nil
}

func main() {
    s := struct {
        id      int
        Name    string
        Age     int
    }{}
    fmt.Println(FieldNames(s))
    // Output:
    // [id Name Age] <nil>
}
```

`reflect` 패키지에서 함수나 메소드도 다룰 수 있으므로 함수를 받아 다른 함수형으로 변경하여 반환할 수도 있다.

아무것도 반환하지 않는 함수가 에러는 반드시 돌려줘야 하는 경우 nil 에러를 돌려주는 함수를 서로 다른 자료형애 대해 만들 수 있다.

```go
func AppendNilError(f interface{}, err error) (interface{}, error) {
    t := reflect.TypeOf(f)
    if t.Kind() != reflect.Func {
        return nil, error.New("AppendNilError: f is not a function")
    }
    in, out := []reflect.Type{}, []reflect.Type{}
    for i := 0; i < t.NumIn(); i++ {
        in = append(in, t.In(i))
    }
    for i := 0; i< t.NumOut(); i++ {
        out = append(out, t.Out(i))
    }
    out = append(out, reflect.TypeOf((*error)(nil)).Elem())
    funcType := reflect.FuncOf(in, out, t.IsVariadic())
    v := reflect.ValueOf(f)
    funcValue := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
        results := v.Call(args)
        results = append(results, reflect.ValueOf(&err).Elem())
        return results
    })
    return funcValue.Interface(), nil
}

func main() {
    f := func() {
        fmt.Println("called")
    }
    f2, err := AppendNilError(f, errors.New("test error"))
    fmt.Println("AppendNilError.err:", err)
    fmt.Println(f2.(func() error)())
}
```

**`reflect`를 사용하면 정적인 자료형 검사를 할 수 없으므로 필요한 경우에만 이용해야 한다**

### go generate

C 언어등에서 제공하는 전처리기 통해 소스코드를 확장해서 컴파일하는 매크로와 비슷한 기능의 도구가 있다.

go generator를 이용하면 임의의 명령을 수행하여 프로그램 코드를 생성할 수 있다.

만약 enum 값을 const를 이용하여 지정했는데 이것을 문자열로 바꾸려면 go generator의 stringer를 활용하면 된다.

우선 stringer를 설치한다: `go get golang.org/x/tools/cmd/stringer`

stringer를 설치하였으면 소스코드에 다음 명령어를 추가한다: `//go:generate stringer -type=Pill` (Pill 자료형을 문자열로 변경)

## 객체지향

Go는 객체지향을 완전하게 지원하지 않는다

### 다형성

객체에 메소드를 호출했을 때 그 객체가 메소드에 대한 다양한 구현을 갖고 있을 수 있다.
이는 Go의 인터페이스로 구현이 가능하다.

```go
type Shape interface {
    Area()  float32
}

type Square struct {
    Size    float32
}

func (s Square) Area() float32 {
    return s.Size * s.Size
}

type Rectangle struct {
    Widht, Height, float32
}

func (r Rectangle) Area() float32 {
    return r.Width * r.Height
}

func TotalArea(shapes []Shape) float32 {
    var total float 32
    for _, shape := range shape {
        total += shape.Area()
    }
    return total
}

func ExampleTotalArea() {
    fmt.Println(TotalArea([]Shape{
        Square(3)
        Rectangle(4, 5)
    }))
}
```

### 인터페이스

Go 언어에서는 자바와 같이 `implements` 예약어를 사용하여 명시적으로 인터페이스를 구현을 선언하지 않는다.

Go 언어에서는 인터페이스 내의 메소드들을 구현하기만 하면 그 인터페이스를 구현하는 것이 된다.

### 상속

객체지향의 Has A 관계의 경우 상속보다는 object composition이며 이는 재사용하고자 하는 자료형의 변수를 구조체에 내장하면 된다.

Is A 관계의 상속은 많은 경우 추상클래스를 상속한다. 이는 Go 언어의 인터페이스를 사용하면 된다.

구상클래스를 상속받아야 하는경우 인터페이스와 구현을 함께 상속받아야 할 것이다. 이 경우를 살펴보자.

#### 메소드 추가

기존의 코드를 재사용 하면서 기능 추가를 하려는 경우 상속할 수 있다.

이러한 용도로 상속이라는 복잡한 개념 대신 메소드 추가를 활용한다.
아래의 코드에 `Area()`외에 둘레를 구하는 기능을 추가하려고 한다.

```go
type Rectangle struct {
    Widht, Height, float32
}

func (r Rectangle) Area() float32 {
    return r.Width * r.Height
}
```

이를 위해 구조체 내장을 사용할 수 있다.

```go
type RectangleCircum struct{ Rectangle }

func (r RectangleCircum) Circum() float32 {
    return 2 * (r.Width + r.Height)
}

func ExampleRectangleCircum() {
    r := RectangleCircum{Rectangle{3, 4}}
    fmt.Println(r.Area())
    fmt.Println(r.Circum())
    // Output:
    // 12
    // 14
}
```

필요하다면 상속과 함께 생성자도 만들어줄 수 있다

```go
func NewRectangleCircum(width, height float32) *RectangleCircum {
    return &RectangleCircum{Rectangle{width, height}}
}
```

#### 오버라이딩

기존의 구현을 다른 구현으로 대체하고자 하는 경우에도 상속을 사용할 수 있다.
이 역시 구조체 내장으로 해결 가능하다.

#### 서브 타입

기존 객체가 쓰이던 곳에 상속받은 객체를 쓰려고 상속하기도 한다.
이 경우 인터페이스와 구조체 내장을 모두 사용할 수 있다.

기본적으로 구조체 내장을 사용하면 기존 메소드를 사용할 수 있으므로 자동으로 인터페이스를 구현한 것이 된다.
따라서 실제 사용할 때도 해당 타입의 서브타입으로 처리된다.

자료형이 주어진 인터페이스를 구현하고 있는지 알아보려면 `reflect.Type.Implements` 메소드를 이용하면 된다
내장된 구조체가 있는지 알아보려면, 구조체에서 내장된 구조체의 이름으로 필드를 찾은 다음 `Anonymous` 필드를 찾아보면 된다.

```go
impl := reflect.TypeOf((*RectangleCircum)(nil)).Elem().Implements(
    reflect.TypeOf((*Shape)(nil)).Elem(),
)
field, ok := reflect.TypeOf(RectangleCircum{}).FieldByName("Rectangle")
emb := ok && field.Anonymous && field.Type == reflect.TypeOf(Rectangle{})
```

impl은 RectangleCircum이 Shape 인터페이스를 구현하는지 여부가
emb는 Rectangle이 RectangleCircum에 내장되어 있는지 여부가 기록된다.

### 캡슐화

Go에서는 상속이 없기 때문에 `protected`는 없지만 패키지 단위의 `public`과 `private`는 존재한다.

대문자로 시작하는 이름은 다른 패키지에서 참조 가능하고, 소문자로 시작되는 이름은 다른 패키지에서 참조 불가능하다.

Go에서는 getter에 get을 붙이지는 않지만 setter에 set을 붙이는 관례가 있다.
또한 자료형 이름이 소문자여서 외부에서 생성하지 못하는 경우 `New~()` 형태의 생성자를 함수로 구현하여 해당자료형이나 인터페이스를 반환할 수 있다.

내부 패키지를 이용하면 작성한 다른 패키지에서는 접근가능하나 외부에서는 접근 불가능하다.
이경우 패키지 경로에 `internal`을 넣으면 `internal`이 있는 경로에 있는 패키지를 포함한 범위에서만 접근 가능하다.

## 디자인 패턴

### 반복자 패턴

클로저를 이용하여 호출하는 반복자

콜백을 넘겨주어 함수가 모든 원소에 대하여 호출되는 반복자

인터페이스를 이용한 반복자

채널을 이용한 반복자 등을 작성하였다

### 추상 팩토리 패턴

추상 팩토리 패턴은 팩토리들을 여럿 묶어 놓은 팩토리를 추상화하는 패턴이다.

윈도우와 맥의 UI 구현을 다르게 하기위하여 추상팩토리를 이용하는 예제이다.

```go
type Button interface {
    Paint()
    OnClick()
}

type Label interface {
    Paint()
}

type WinButton struct{}

func (WinButton) Paint() {
    fmt.Println("win button paint")
}
func (WinButton) OnClick() {
    fmt.Println("win button click")
}

type WinLabel struct{}

func (WinLabel) Paint() {
    fmt.Println("win label paint")
}

type MacButton struct{}

func (MacButton) Paint() {
    fmt.Println("mac button paint")
}
func (MacButton) OnClick() {
    fmt.Println("mac button click")
}

type MacLabel struct{}

func (MacLabel) Paint() {
    fmt.Println("mac label paint")
}

// UI factory can create buttons and labels
type UIFactory interface {
    CreateButton()  Button
    CreateLabel()   Label
}

type WinFactory struct{}

func (WinFactory) CreateButton() Button {
    return WinButton{}
}

func (WinFactory) CreateLabel() Label {
    return WinLabel{}
}

type MacFactory struct{}

func (MacFactory) CreateButton() Button {
    return MacButton{}
}

func (MacFactory) CreateLabel() Label {
    return MacLabel{}
}

// CreateFactory returns a UIFactory of the given os
func CreateFactory(os string) UIFactory {
    if os == "win" {
        return WinFactory{}
    } else {
        return MacFactory{}
    }
}

func Run(f UIFactory) {
    button := f.CreateButton()
    butotn.Paint()
    button.OnClick()
    label := f.CreateLabel()
    label.Paint()
}

func main() {
    f := CreateFactory("win")
    Run(f)
}
```

### 비지터 패턴

비지터 패턴(visitor pattern)은 알고리즘을 객체 구조에서 분리시키기 위한 디자인 패턴이다.

Visitor를 `accept()`로 받아서(Visitor가 데이터 구조를 돌아다니면서) 필요한 알고리즘을 수행(`visit()`)하도록 한다.

인터페이스를 이용하여 구현할 수 있다.

```go
type CarElementVisitor interface {
    VisitWheel(wheel Wheel)
    VisitEngine(engine Engine)
    VisitBody(body body)
    VisitCar(car Car)
}

type Acceptor interface {
    Accept(visitor CarElementVisitor)
}

type Wheel string

func (w Wheel) Name() string {
    return string(w)
}

func (w Wheel) Accept(visitor CarElementVisitor) {
    visitor.VisitWheel(w)
}

type Engine string

func (e Engine) Accept(visitor CarElementVisitor) {
    visitor.VisitEngine(e)
}

type Body string

func (b Body) Accept(visitor CarElementVisitor) {
    visitor.VisitBody(b)
}

type Car []Acceptor

func (c Car) Accept(visitor CarElementVisitor) {
    for _, e := range c {
        e.Accept(visitor)
    }
    visitor.VisitCar(c)
}

type CarElementPrintVisitor struct{}

func (CarElementPrintVisitor) VisitWheel(wheel Wheel) {
    fmt.Println("Visiting " + wheel.Name() + " wheel.")
}

func (CarElementPrintVisitor) VisitEngine(engine Engine) {
    fmt.Println("Visiting engine")
}

func (CarElementPrintVisitor) VisitBody(body Body) {
    fmt.Println("Visiting body")
}

func (CarElementPrintVisitor) VisitCar(car Car) {
    fmt.Println("Visiting car")
}

type CarElementDoVisitor struct{}

func (CarElementDoVisitor) VisitWheel(wheel Wheel) {
    fmt.Println("Kicking my " + wheel.Name() + " wheel.")
}

func (CarElementDoVisitor) VisitEngine(engine Engine) {
    fmt.Println("Starting my engine.")
}

func (CarElementDoVisitor) VisitBody(body Body) {
    fmt.Println("Moving my body")
}

func (CarElementDoVisitor) VisitCar(car Car) {
    fmt.Println("Starting my car")
}

func main() {
    car := Car{
        Wheel("front left"),
        Wheel("front right"),
        Wheel("back left"),
        Wheel("back right"),
        Body{},
        Engine{},
    }
    car.Accept(CarElementPrintVisitor{})
    car.Accept(CarElementDoVisitor{})
}
```
