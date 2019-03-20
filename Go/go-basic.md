# Go 기초

## 변수와 상수

### 변수

Go 언어는 변수 선언시 `var 변수명 자료형` 순서로 쓴다

선언한 변수는 자료형에 맞는 값을 할당할 수 있다.
또한 선언과 동시에 할당할 수도 있다.

```go
var num int
num = 10

var num int = 10
```

동일한 타입의 변수를 여러개 선언할 때 한줄로 쓸 수 있다.
마찬가지로 선언과 동시에 값을 할당할 수 있다.

```go
var a, b, c int

var a, b, c int = 1, 2, 3
```

마지막으로 변수 타입을 기재하지 않고 할당하는 값으로 타입추론을 할 수 있다.

```go
num := 10       // (int 자료형)
str := "Hello"  // (string 자료형)
```

### 상수

상수는 `const` 키워드를 사용하여 선언한다

```go
const num int = 10
```

Go에서는 `enum` 자료형이 없고 대신 상수를 사용한다.

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

## 키워드 (예약어)

Go 언어에서는 예약어가 25개 존재하며 다른 언어에 비해 적은 편이다

```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

## 데이터 타입

- 불리언 타입: bool
- 문자열 타입: string (Immutable 타입)
- 정수형 타입: int int8 int16 int32 int64
- 부호없는 정수형 타입: uint uint8 uint16 uint32 uint64 uintptr
- Float 및 복소수 타입: float32 float64 complex64 complex128
- 기타 타입
  - byte: uint8과 동일하며 바이트 코드에 사용
  - rune: int32과 동일하며 유니코드 코드포인트에 사용한다

문자열 리터럴은 Back Quote(``) 혹은 이중인용부호(" ")를 사용하여 표현한다.

Back Quote에서는 여러줄에 걸쳐 문자열을 쓸 수있고 escape 문자열이 변환되지 않고 그대로 출력된다.

### 데이터 타입 변환

데이터 타입 변환을 위해서는 일종의 생성자와 비슷한 형태의 문법을 사용한다

`변환할타입(변환전데이터)`

## 연산자

### 산술연산자

사칙연산자(+, -, *, /, % (나머지))

증감연산자(++, --)

### 관계연산자

- 같다: a == b
- 같지 않다: a != c
- 크거나 같다: a >= b
- 작거나 같다: a <= b
- 크다: a > b
- 작다: a < b

### 논리연산자

- AND: `&&`
- OR: `||`
- NOT: `!`

### Bitwise 연산자

Bitwise 연산자는 비트단위 연산을 위해 사용된다.

- AND: `&`
- OR: `|`
- XOR: `^`
- 쉬프트: `<<`, `>>`

### 할당연산자

- 할당연산: `=`
- 더하고 할당: `+=`
- 빼고 할당: `-=`
- 곱하고 할당: `*=`
- 나누고 할당: `/=`
- 나머지 구하고 할당: `%=` (C %= A is equivalent to C = C % A)
- 왼쪽 쉬프트 후 할당: `<<=` (C <<= 2 is same as C = C << 2)
- 오른쪽 쉬프트 후 할당: `>>=` (C >>= 2 is same as C = C >> 2)
- 비트 AND 연산후 할당: `&=` (C &= 2 is same as C = C & 2)
- 비트 XOR 연산후 할당: `^=` (C ^= 2 is same as C = C ^ 2)
- 비트 OR 연산후 할당: `|=` (C |= 2 is same as C = C | 2)

### 포인터연산자

```go
var p = &k  // k의 주소를 할당
println(*p) // p가 가리키는 주소에 있는 실제 내용을 출력
```

## 조건문

### if문

if 문의 조건은 괄호없이 사용하고 else if, else의 문법은 다른 언어와 동일하다

```go
if c == 1 {
    // something
} else if c == 2 {
    // something
} else {
    // something
}
```

조건문을 사용하기 전 Optional Statement를 함께 실행할 수 있다

```go
if val := i; val == 5 {
    // something
}
```

### switch문

```go
switch category {
case 1:
    // something
case 2, 3:
    // something
default:
    // something
}
```

스위치문에서도 Optional Statement를 사용할 수 있다

```go
switch x := i; x + 1 {
    //...
}
```

Go언어의 switch문은 case문에 복잡한 expression을 가질 수 있다

Go언어의 switch문은 변수의 Type에 따라 case로 분기할 수 있다

## 반복문

Go언어의 반복문은 `while`이 없고 `for` 루프 하나밖에 없다

```go
for i := 1; i <= 100; i++ {
    fmt.Println(i)
}
```

초기값과 증감식을 생략하고 조건식만 작성하면 `while` 루프 처럼 사용할 수 있다

```go
n := 1
for n < 100 {
    fmt.Println(n)
    n++
}
```

조건식도 생략하면 무한루프로 사용할 수 있다

```go
for {
    println("Infinite loop")
}
```

### for range

for range 문은 컬렉션으로 부터 한 요소씩 가져와 차례로 for 문장내의 블럭에서 실행한다.

```go
names := []string{"김씨", "이씨", "박씨"}
for index, name := range names {
    fmt.Println(index, name)
}
```

### break, continue

- break: 즉시 루프 탈출시 사용
- continue: 해당 키워드 아래 블럭을 건너뛰고 다시 루프 초기로 돌아간다
