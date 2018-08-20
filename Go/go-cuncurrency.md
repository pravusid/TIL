# 동시성

## 고루틴

고루틴은 경량 스레드로 볼 수 있으며, 현재 수행 흐름과 별개 흐름을 만들어 준다.

고루틴을 생성하는 방법은 매우 간단하다.

```go
go f(x, y, z)
```

함수 앞에 `go`를 붙여서 호출하게 되면 메모리를 공유하는 별개의 논리적 흐름이 생성된다.

### 병렬성과 병행성

물리적으로 별개의 흐름이 수행되는 경우 이를 병렬성(parallelism)이라 한다.
반면 논리적으로 별개의 흐름이 수행되는 경우 이를 동시성 또는 병행성(concurrency)라고 한다.

동시성이 있는 두 루틴은 서로 의존관계가 없다.
동시성은 병렬성과 다르지만 동시성이 있어야 병렬성이 성립하게 된다.

```go
func main() {
    go func() {
        fmt.Println("goroutine")
    }()
    fmt.Println("main routine")
}
```

main과 고루틴은 어느것이 먼저 실행될지 알 수 없으나,
만약 main이 먼저 실행되고 종료되어 버린다면 고루틴은 실행되지도 않을 수 있다.

### 고루틴 기다리기

고루틴을 제어하기 위한 싱크 라이브러리가 제공된다.

[이미지 압축 예제](training/goroutine/imgzip.go)

```go
var wg sync.WaitGroup
wg.Add(len(urls))
for _, url := range urls {
    go func(url string) {
        defer wg.Done()
        if _, err := download(url); err != nil {
            log.Fatal(err)
        }
    }(url)
}
wg.Wait()
```

wg에는 기본값이 0인 카운터가 들어가 있고 `Wait()`는 카운터가 0이 될 때까지 기다린다.

전체 작업이 몇 개일지 알수 없는경우 wg 개수를 동적으로 할당하면 된다.

```go
var wg sync.WaitGroup
for _, url := range urls {
    wg.Add(1)
    go func(url string) {
        defer wg.Done()
        if _, err := download(url); err != nil {
            log.Fatal(err)
        }
    }(url)
}
wg.Wait()
```

이 때 `wg.Add()`를 고루틴 내부에 포함시키면 `Add`가 수행되기 전에 `Wait`을 통과하는 race condition이 발생할 수 있다.

#### 공유 메모리와 병렬 최소값 찾기

앞에서 고루틴은 파일시스템을 공유하였다. 마찬가지로 고루틴들은 메모리를 공유한다.

공유 메모리를 이용하여 병렬화가 되는 예제를 풀어보자.
우선 병렬화 하지 않고 가장 작은 수를 찾는 함수를 작성해보자

```go
func Min(a []int) int {
    if len(a) == 0 {
        return 0
    }
    min := a[0]
    for _, e := range a[1:] {
        if min > e {
            min = e
        }
    }
    return min
}
```

병렬로 작동하도록 작성해보자. 병렬로 데이터를 나누어 가장 작은 수를 찾은 후, 찾은 결과를 모아 한번 더 작은 수를 찾는다.

```go
// n은 고루틴의 개수이다
func ParallelMin(a []int, n int) int {
    if len(a) < n {
        return Min(a)
    }
    mins := make([]int, n)
    size := (len(a) + n - 1) / n
    var wg sync.WaitGroup
    for i := 0; i < n; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            begin, end := i*size, (i+1)*size
            if end > len(a) {
                end = len(a)
            }
            mins[i] = Min(a[begin:end])
        }(i)
    }
    wg.Wait()
    return Min(mins)
}
```

i번째 고루틴은 i번째 배열에 병렬로 값을 넣게 된다.

## 채널

고루틴과 데이터를 주고 받기 위해 메모리나 파일시스템을 공유할 수도 있으나, 채널을 활용하는 것도 가능하다.

채넣은 넣은 데이터를 뽑아낼 수 있는 파이프같은 형태의 자료구조이다.
채널에 데이터를 넣고 뽑아서 다른 고루틴과 통신을 할 수 있다.

채널을 사용하려면 `chan`을 쓰고 주고받을 자료형을 쓰면 된다.
채널은 양방향 채널과 단방향 채널이 있으며 양방향 채널은 단방향 채널로 변환해서 쓸 수 있다.

채널을 서로 복사하는 경우 동일 채널을 가리키는 것이 된다.
따라서 채널은 그 자체로 포인터와 비슷한 레퍼런스 형이라고 할 수 있다.

```go
c1 := make(chan int)
var chan int c2 = c1
var <-chan int c3 = c1
var chan<- int c4 = c1
```

c1을 새 정수 채널로 만들었고 c2는 다른 채널 변수에 c1을 할당한 것이다. c1과 c2는 동일 채널이다.
복사시에 동일한 채널이므로 함수를 호출할 때 채널을 넘기면 함수내에서 동일한 채널에 값을 넣고 뺄 수 있다.

c3의 자료형은 왼쪽 화살표가 있는데 자료를 뺄수만 있는 채널이다(받기 전용).
c4의 자료형은 오른쪽에 화살표가 있고 자료를 넣을 수만 있는 채널이다(보내기 전용).

채널에 자료를 보낼 때 채널 왼쪽 화살표 자료 순으로 작성한다.

```go
c <- 100
```

채널에서 자료를 받을때는 채널 왼쪽에 화살표를 붙이면 된다.

```go
data := <-c
```

### 일대일 단방향 채널 소통

```go
func Example_simpleChannel() {
    c := make(chan int)
    go func() {
        c <- 1
        c <- 2
        c <- 3
    }()
    fmt.Println(<-c)
    fmt.Println(<-c)
    fmt.Println(<-c)
    // Output:
    // 1
    // 2
    // 3
}
```

하지만 보내는 고루틴과 받는 고루틴의 숫자가 맞지 않으면 고루틴이 멈춘다.
보내는 고루틴이 들어왔으나 받는 고루틴이 호출되지 않는다면 고루틴은 멈춰있게 되고,
다른 고루틴으로 문맥 전환(context switching)이 발생한다.

받는 부분과 보내는 부분이 데이터의 개수를 알지 못해도 동작하도록 수정해보자

```go
func Example_simpleChannel() {
    c := make(chan int)
    go func() {
        c <- 1
        c <- 2
        c <- 3
        close(c)
    }()
    for num := range c {
        fmt.Println(num)
    }
    // Output:
    // 1
    // 2
    // 3
}
```

함수가 채널을 반환하게 만드는 패턴을 사용할 수 있다.

```go
func Example_simpleChannel() {
    c := func() <-chan int {
        c := make(chan int)
        go func() {
            defer close(c)
            c <- 1
            c <- 2
            c <- 3
        }()
        return c
    }()
    for num := range c {
        fmt.Println(num)
    }
    // Output:
    // 1
    // 2
    // 3
}
```

### 생성기 패턴

```go
// 피보나치 수열을 max까지 생성
func Fibonacci(max int) <-chan int {
    c := make(chan int)
    go func() {
        defer close(c)
        a, b := 0, 1
        for a <= max {
            c <- a
            a, b = b, a+b
        }
    }()
    return c
}

func ExampleFibonacci() {
    for fib := range Fibonacci(15) {
        fmt.Print(fib, ",")
    }
    // Output:
    // 0, 1, 1, 2, 3, 5, 8, 13
}
```

클로저를 이용하여 생성기를 만든다면 다음과 같다

```go
func FibonacciGenerator(max int) func() int {
    next, a, b := 0, 0, 1
    return func() int {
        next, a, b = a, b, a+b
        if next > max {
            return -1
        }
        return next
    }
}

func ExampleFibonacciGenerator() {
    fib := FibonacciGenerator(15)
    for n := fib(); n >= 0; n = fib() {
        fmt.Print(n, ",")
    }
    // Output:
    // 0, 1, 1, 2, 3, 5, 8, 13
}
```

생성기 패턴을 이용하면 몇 가지 장점이 있다

1. 생성하는 쪽에서 상태 저장 방법을 고민할 필요가 없다
2. 받는 쪽에서는 for의 range를 이용할 수 있다
3. 채널 버퍼를 이용하면 멀티 코어를 활용하거나 입출력 성능상의 장점을 이용할 수 있다.

for의 range를 이용하는 예제를 살펴보자.

```go
func BabyNames(first, second string) <- chan string {
    c := make(chan string)
    go func() {
        defer close(c)
        for _, f := range first {
            for _, s : range second {
                c <- string(f) + string(c)
            }
        }
    }()
    return c
}

func ExampleBabyNames() {
    for n := range BabyNames("성정명재경", "준호우훈진") {
        fmt.Print(n, ", ")
    }
}
```

### 버퍼 있는 채널

버퍼가 없는 채널에 값을 보낼때는 받는쪽도 준비되어 있어야 한다.

받는쪽에 준비가 없어도 보내는 쪽이 미리 보내려면 채널에 버퍼를 잡으면 된다.

```go
// 버퍼의 크기가 10인 정수 채널
c := make(chan int, 10)
```

### 닫힌 채널

채널에서 값을 받을 때 두번째 변수로 채널이 열려있는지 여부를 알수있다

```go
val, ok := <-c
```

채널이 열려있다면 ok는 true가 되고, 채널이 닫혀있다면 ok는 false val은 타입 기본값이 들어온다.

만약 채널에 값이 없는 상태라면 열려있는 채널은 값이 들어올때까지 block 상태로 대기하고,
닫혀있는 채널은 기다리지 않고 기본값을 받고 넘어간다.

만약 닫은 채널을 다시 닫으면 패닉이 발생한다.

## 동시성 패턴

### 파이프라인 패턴

파이프라인은 한 단계의 출력이 다음 단계의 입력으로 이어지는 구조이다.

파이프라인 패턴은 생성기 패턴의 일종이다. 받기 전용 채널을 넘겨받아 입력으로 활용한다.
생성기 패턴과 동일하게 채널은 데이터를 보내는 쪽에서 닫아야 한다.

```go
// PlusOne returns a channel of num + 1 for nums received from in
func PlusOne(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for num := range in {
            out <- num + 1
        }
    }()
    return out
}

func ExamplePlusOne() {
    c := make(chan int)
    go func() {
        defer close(c)
        c <- 5
        c <- 3
        c <- 8
    }()
    for num := range PlusOne(PlusOne(c)) {
        fmt.Println(num)
    }
    // Output:
    // 7
    // 5
    // 10
}
```

PlusOne은 받기 전용 채널을 받아서 다른 받기 전용 채널을 돌려주는 함수이다.
받은 채널에서 숫자를 하나 증가시켜서 반환한다.

서로 다른함수들도 이어 붙일 수 있다.

```go
type IntPipe func(<-chan int) <-chan int

func Chain(ps ..IntPipe) IntPipe {
    return func(in <-chan int) <-chan int {
        c := in
        for _, p := range ps {
            c = p(c)
        }
        return c
    }
}
```

작성한 Chain을 이용하면 `A(B(c))`를 `Chain(B, A)(c)`와 같이 표현할 수 있다.

### 채널 공유로 팬아웃 하기

팬아웃(fan-out: 게이트 하나의 출력이 게이트 여러 입력으로 들어가는 경우)

앞의 작업은 빠르게 이루어지나 후속작업에 시간이 걸려 결과물을 여러곳에 분배해야 하는 경우가 있다.
이런경우 채널 하나를 여럿에게 공유하면 된다.

```go
func main() {
    c := make(chan int)
    for i := 0; i < 3; i++ {
        go func(i int) {
            for n:= range c {
                time.Sleep(1) // 고루틴 3개에서 모두 출력해보기 위해서 sleep
                fmt.Println(i, n)
            }
        }(i) // 값을 넘겨 사용하지 않으면 외부의 i값과 고루틴내의 i값이 달라 결과에 문제가 발생할 수 있다
    }
    for i := 0; i < 10; i++ {
        c <- 1
    }
    close(c)
}
```

### 팬인 하기

Fan-in으로 결과값을 합치려해도 채널을 공유하면 된다.
같은 채널에 여러 고루틴이 값을 보내도 받아가는 곳에서는 모든 값을 받아갈 수 있다.

다만 여기에서 채널을 닫을 때는 주의해야 한다. 고루틴에서 채널을 닫아버리면 채널 닫기가 반복수행 되어 패닉이 발생한다.

```go
func FanIn(ins ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    wg.Add(len(ins))
    for _, in := range ins {
        go func(in <-chan int) {
            defer wg.Done()
            for num := range in {
                out <- num
            }
        }(in)
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

c := FanIn(c1, c2, c3)
```

### 분산처리

팬아웃 해서 파이프라인을 통과시키고 다시 팬인시키면 분산처리를 할 수 있다.

```go
// Distribute 함수는 IntPipe를 받은 뒤 n개로 분산처리 하는 함수로 돌려준다
func Distribute(p IntPipe, n int) IntPipe {
    return func(in <-chan int) <-chan int {
        cs := make([]<-chan int, n)
        for i := 0; i < n; i++ {
            cs[i] = p(in)
        }
        return FanIn(cs...)
    }
}
```

Distribute와 Chain을 함께 이용해서 파이프라인을 구성할 수 있다.

```go
// in으로 들어온 자료가 Cut 고루틴에서 처리되고 그 결과값이 10개로 나누어진 Draw Pain Decorate
// Chain을 통과한후 합쳐져 Box 고루틴을 통과한다
out := Chain(Cut, Distribute(Chain(Draw, Paint, Decorate), 10) Box)(in)
```

Go에서는 고루틴마다 스레드를 모두 할당하지 않으며,
동시에 수행될 필요가없는 고루틴은 하나의 스레드에서 순차적으로 수행되며 이는 컴파일 타임에 예측가능한 경우가 많다.

### select

`select`를 사용하면 동시에 여러 채널과 통신할 수 있다.
`select`의 형태는 `switch`문과 비슷하지만 동시성 프로그래밍에 사용되며 다른 특성들이 있다.

- 모든 case가 계산된다. 함수 호출이 있으면 select를 수행할 때 모두 호출된다.
- 각 case는 채널에 입출력하는 형태가 되며 막히지 않고 입출력이 가능한 case가 있으면 그중에 하나가 선택되어 입출력이 수행되고 해당 case의 고드만 수행된다.
- default가 있으면 모든 case에 입출력이 불가능할 때 코드가 수행된다. default가 없고 모든 case에 입출력이 불가능하면 가능한 case가 발생할 때까지 기다린다.

```go
select {
    case n := <-c1:
        fmt.Println(n, "is from c1")
    case n := <-c2:
        fmt.Println(n, "is from c2")
    default:
        fmt.Println("default")
}
```

#### 팬인하기

`select`를 사용하면 고루틴을 여러개 이용하지 않아도 팬인 할 수 있다.

```go
select {
    case n := <-c1:
        c <- n
    case n := <-c2:
        c <- n
    case n := <-c3:
        c <- n
}
```

위 코드는 c1, c2, c3 중 어느 채널이라도 데이터가 준비되면 그것을 c로 보내는 코드이다.
select문을 반복하면 팬인을 할 수 있다.

그러나 일부 닫혀있는 채널이 있다면 의도하지 않은 기본값이 계속 들어올 수 있다.
이를 수정한 예제를 살펴보자

```go
func FanIn(in1, in2, in3 <-chan int) <-chan int {
    out := make(chan int)
    openCount := 3
    closeChan := func(c *<-chan int) bool {
        *c = nil
        openCount--
        return openCount == 0
    }
    go func() {
        defer close(out)
        for {
            select {
                case n, ok := <-in1:
                if ok {
                    out <- n
                } else if closeChan(&in1) {
                    return
                }
                case n, ok := <-in2:
                    if ok {
                        out <- n
                    } else if closeChan(&in2) {
                        return
                    }
                case n, ok := <-in3:
                    if ok {
                        out <- n
                    } else if closeChan(&in3) {
                        return
                    }
            }
        }
    }()
    return out
}

func main() {
    c1, c2, c3 := make(chan int), make(chan int), make(chan int)
    sendInts := func(c chan<- int, begin, end int) {
        defer close(c)
        for i := begin; i < end; i++ {
            c <- i
        }
    }
    go sendInts(c1, 11, 14)
    go sendInts(c2, 21, 23)
    go sendInts(c3, 31, 35)
    for n := range FanIn(c1, c2, c3) {
        fmt.Print(n, ",")
    }
}
```

이 경우 닫힌 채널을 `nil`로 바꾸어준다. `nil` 채널은 받기 보내기가 모두 막힌다.
이를 반복수행하면서 열려 있는 채널이 0개가 되면 함수를 반환하고 out채널을 닫는다.

#### 채널을 기다리지 않고 받기

채널에 값이 준비되지 않으면 준비될때 까지 기다리는 것이 기본동작이다.
select를 이용해서 채널값이 있으면 받고 없으면 스킵하는 흐름을 구성해보자

```go
select {
    case n := <-c:
        fmt.Println(n)
    default:
        fmt.Println("Data is not ready")
}
```

#### 시간제한

채널과 통신을 기다리는 시간의 상한선을 지정하려면 `time.After` 함수를 이용할 수 있다.
이 함수는 지정된 시간이 지나면 현재시간이 전달되는 채널을 반환한다.

```go
select {
    case n := <-recv:
        fmt.Println(n)
    case send <- 1:
        fmt.Println("sent 1")
    case <-time.After(5 * time.Second):
        fmt.Println("5초간 통신이 없음")
}
```

전체 제한시간을 지정하려면 타이머 채널을 반복문 밖에서 할당하면 된다.

```go
timeout := time.After(5 * time.Second)
for {
    select {
        case n := <-recv:
            fmt.Println(n)
        case send <- 1:
            fmt.Println("sent 1")
        case <-timeout:
            fmt.Println("5초간 통신이 없음")
    }
}
```

### 파이프라인 중단하기

채널을 받는 쪽에서 필요한 자료를 획득하여 중간에서 채널사용을 끝내고 싶은 경우가 있다.
만약 채널에서 데이터를 받는 루프를 끝내버리면 고루틴은 여전히 메모리상에서 작동하므로 자원 낭비가 발생한다.

그렇다고 받는 쪽에서 채널을 닫아버리면 보내는쪽에서 채널에 접근할 때 패닉이 발생한다.

이때 done 채널을 하나더 만들어 보내는 고루틴에서 종료를 감지하게 만들면 된다.

```go
func PlusOne(done <-chan struct{}, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for num := range in {
            select {
                case out <- num + 1:
                case <-done:
                    return
            }
        }
    }()
    return out
}

func main() {
    c := make(chan int)
    go func() {
        defer close(c)
        for i := 3; i < 103; i += 10 {
            c <- i
        }
    }()
    done := make(chan struct{})
    nums := plusOne(done, PlusOne(done, Plusone(done, PlusOne(done, PlusOne(done, c)))))
    for num := range nums {
        fmt.Println(num)
        if num == 18 {
            break
        }
    }
    close(done)
    time.Sleep(100 * time.Millisecond)
    fmt.Println("고루틴수: ", runtime.NumGoroutine())
    for _ = range nums {
        // consume all nums
    }
    time.Sleep(100 * time.Millisecond)
    fmt.Println("고루틴수: ", runtime.NumGoroutine())
}
```

`close(done)`을 호출해서 done채널을 select로 받고있는 곳(보내는 채널)에 기본값이 전송되고 채널을 닫게 된다.

### 컨텍스트(context.Context) 활용

위와 같이 done 채널을 운영해도 되지만 복잡한 상황을 대비해 context 패턴을 이용하는 것이 좋다.

사용을 위해서는 라이브러리를 설치해야 한다: `go get golang.org/x/net/context`

```go
func PlusOne(ctx context.Context, in <-chan int) <-chan it {
    out := make(chan int)
    go func() {
        defer close(out)
        for num := range in {
            select {
                case out <- num + 1:
                case <-ctx.Done():
                    return
            }
        }
    }()
    return out
}

func main() {
    c := make(chan int)
    go func() {
        defer close(c)
        for i := 3; i < 103; i += 10 {
            c <- 1
        }
    }()
    ctx, cancel := context.WithCancel(context.Background())
    nums := PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, c)))))
    for num := range nums {
        fmt.Println(num)
        if num == 18 {
            cancel()
            break
        }
    }
}
```

`context.Background()`가 가장 상위에 존재하여 프로그램 종료시까지 계속 살아있다.
`context.WithCancel()`로 최사우이 컨텍스트에 취소 기능을 붙이고, ctx에 컨텍스트를 cancel에 취소 호출함수를 할당한다.

`WithDeadline`, `WithTimeout`을 이용하여 시간이 지나면 취소되게 할 수도 있다.

컨텍스트는 관례상 함수의 첫 번째 인자로 넘겨주고 받는다.

### 요청과 응답 매칭

데이터를 받았을 때 어느요청에 대한 응답인지 알아야 하는 경우가 있다.

한 가지 방법은 채널로 자료를 넘겨주고 받을 때 id 번호를 같이 넘겨 확인하는 것이다.
그러나 요청에 대한 응답을 다른 고루틴이 가져가면 의도하지 않은 결과가 발생할 수 있다.

이런경우 보내는 Request 구조체에 받을 채널도 함께 넣어서 보내면 된다.

```go
type Request struct {
    Num     int
    Resp    chan Response
}

type Response struct {
    Num         int
    WorkerId    int
}

func PlusOneService(reqs <-chan Request, workerId int) {
    for req := range reqs {
        go func(req Request) {
            defer close(req.Resp)
            req.Resp <- Response(req.Num + 1, workerId)
        }(req)
    }
}

func main() {
    reqs := make(chan Request)
    defer close(reqs)
    for i := 0; i < 3; i++ {
        go PlusOneService(reqs, i)
    }
    var wg sync.WaitGroup
    for i := 3; i < 53; i += 10 {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            resps := make(chan Response)
            reqs <- Request{i, resps}
            fmt.Println(i, "=>", <-resps)
        }(i)
    }
    wg.Wait()
}
```

### 동적으로 고루틴 이어붙이기

prime의 배수를 걸러내는 고루틴을 계속해서 이어붙이는 예제

2부터 숫자를 하나씩 증가시켜 채널에 보내고, 다른 고루틴에서는 채널에서 숫자를 받을 때마다 출력한다.
출력된 숫자의 배수가 되는 숫자들을 걸러내는 필터 고루틴을 이어붙이면 출력된 숫자들은 모두 소수가 된다.

```go
// Range returns a channel and sends ints
// (start, start+step, start+2*step)
func Range(ctx context.Context, start, step int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := start; ; i += step {
            select {
                case out <- i
                case <-ctx.Done:
                    return
            }
        }
    }()
    return out
}

//FilterMultiple returns a IntPipe that filters multiple of n
func FilterMultiple(n int) IntPipe {
    return func(ctx context.Context, in <-chan int) <-chan int {
        out := make(chan int)
        go func() {
            defer close(out)
            for x := range in {
                if x%n == 0 {
                    continue
                }
                select {
                    case out <- x:
                    case <-ctx.Done():
                        return
                }
            }
        }()
        return out
    }
}

func Primes(ctx context.Context) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        c := Range(ctx, 2, 1)
        for {
            select {
                case i := <-c:
                    c = FilterMultiple(i)(ctx, c)
                    select {
                        case out <- i:
                        case <-ctx.Done():
                            return
                    }
                case <-ctx.Done():
                    return
            }
        }
    }()
    return out
}
```

`<-c`를 받는 부분에서 막혀있으면 ctx가 취소될 수 있고
여기에서 받은 값을 `out <- i` 처럼 보낼 때 막혀있다 ctx가 취소될 수 있으므로 select를 다중으로 만들어야 한다.

Primes를 이용하는 코드 예제 이다

```go
func PrintPrimes(max int) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    for prime := range Primes(ctx) {
        if prime > max {
            break
        }
        fmt.Print(prime, " ")
    }
    fmt.Println()
}
```

### 주의점

- 자료를 보내는 채널은 보내는 쪽에서 닫는다
- 보내는 쪽에서 반복문 등을 활용하여 보냈다가 중간에 return을 할 수 있으므로 닫을 때에는 defer를 이용하는 것이 좋다. 그렇지 않으면 중간에 return 했을 때 채널을 닫지 않고 종료할 수 있다.
- 모든 자료처리가 끝나는 시점까지 기다리는 방법으로는 받는 쪽이 끝날때까지 기다리는 것이 더 안정적이다. 즉 생산자가 아닌 소비자 쪽에서 `done <- true`를 호출하고 생산자는 끝났다는 신호를 받아 채널을 닫아야 한다.
- 특별한 이유가 없다면 받는 쪽에서는 range를 이용하는 것이 좋다. 생산자가 채널을 닫은 경우에 반복문을 빠져나오므로 편리하다.
- 루틴이 끝났음을 알리고 다른 쪽에서 기다리는 것은 `sync.WaitGroup`을 이용하는 것이 나은 경우가 많다.
- 끝났음을 알리는 `done` 채널은 자료를 보내는 쪽에서 결정하지말고 차라리 채널을 닫아서 끝났음을 알리는 것이 낫다.
- `done` 채널에 자료를 보내어 신호를 주는 것 보다 `close(done)`으로 채널을 닫는것이 나은 경우가 많다.

## 경쟁 상태

고루틴들을 사용하다보면 서로 막혀 교착 상태(dead lock)이 발생하는 경우 오류가 출력되지만,
경쟁 상태(race condition)의 경우 쉽게 발견할 수 없는 버그가 될 수 있다.

경쟁 상태는 공유된 자원에 둘 이상의 프로세스가 동시에 접근하여 잘못된 결과가 나올 수 있는 상태이다.
타이밍에 따라서 결과가 달라지므로 알아채기도, 고치기도 번거로운 버그가 될 수 있다.

### 동시성 디버그

Go 도구에서 `-race` 옵션을 주면 경쟁 상태를 탐지할 수 있다.

```go
go test -race mypkg
go run -race mysrc.go
go build -race mycmd
go install -race mypkg
```

동적으로 고루틴을 이어붙일때 메모리 누수가 일어나는지 알아보기 위해서 `runtime.NumGoroutine()`을 호출하면 된다.

`runtime.NumCPU()`와 `runtime.GOMAXPROCS()`를 이용하여 현재 사용가능한 CPU 수와 얼마의 CPU를 이용할지 통제할 수도 있다.

### atomic과 sync.WaitGroup

`sync/atomic` 패키지에는 여러 경쟁상태를 대비하기 위한 함수들이 있다.

채널을 이용해도 복잡한 문제를 쉽개 해결할 수 있다.

```go
func main() {
    req, resp := make(chan struct{}), make(chan int64)
    cnt := int64(10)
    go func(cnt int64) {
        defer close(resp)
        for _ = range req {
            cnt--
            resp <- cnt
        }
    }(cnt)
    for i := 0; i < 10; i++ {
        go func() {
            // do somthing
            req <- struct{}{}
        }()
    }
    for cnt = <-resp; cnt > 0; cnt = <-resp {
    }
    close(req)
    fmt.Println(cnt)
}
```

코드에 의도를 명확히 나타내기 위해 WaitGroup을 사용할 수도 있다

```go
func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // do something
        }()
    }
    wg.Wait()
}
```

### sync.Once

```go
func main() {
    done := make(chan struct{})
    go func() {
        defer close(done)
        fmt.Println("initialized")
    }()
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            <-done
            fmt.Println("Goroutine:", i)
        }(i)
    }
    wg.Wait()
}
```

이 경우 고루틴의 실행순서는 알 수 없으나 고루틴들은 `<-done`으로 먼저 기다린 후 실행되므로 초기화 코드가 우선 수행된다.

이렇게 한 번만 코드를 수행할때 쓸 수 있는 것이 `sync.Once`이다.

```go
func main() {
    var once sync.Once
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            once.Do(func() {
                fmt.Println("initialized")
            })
            fmt.Println("Goroutine:", i)
        }(i)
    }
    wg.Wait()
}
```

### Mutex / RWMutex

뮤텍스(mutex)는 상호 배타 잠금기능이 있다.
동시에 둘 이상의 고루틴에서 코드 흐름을 제어할 수 있다. 외부자원에 접근하는 경우 활용하면 효과적이다.

뮤텍스를 활용하는 방법 중 하나는 접근하고자 하는 자원 포인터와 뮤텍스 포인터를 하나의 구조체에 넣어 사용하는 것이다.

```go
type Accessor struct {
    R   *Resource
    L   *sync.Mutex
}
```

생성할 때 `Accessor{&resource, &sync.Mutex{}}`와 같이 할당해주고 자원에 접근하는 메소드에서 `Lock`을 활용한다.

```go
func (acc *Accessor) Use() {
    // do something
    acc.L.Lock()
    // Use acc.R
    acc.L.Unlock()
    //do something else
}
```

`sync.RWMutex`는 조금 더 복잡하다.
어떤 자원에 여러 프로세스가 쓰기 접근을 한다면 다른 프로세스 모두 그 동안 접근할 수 없는 경우 이용된다.

```go
type ConcurrentMap struct {
    M   map[string]string
    L   *sync.RWMutex
}

func (m ConcurrentMap) Get(key string) string {
    m.L.RLock()
    defer m.L.RUnlock()
    return m.M[key]
}

func (m ConcurrentMap) Set(key, value string) {
    m.L.Lock()
    m.M[key] = value
    m.L.Unlock()
}

func main() {
    m := ConcurrentMap{map[string]string{}, &sync.RWMutex{}}
    // ...
}
```

RWMutex도 Mutex의 일종이므로 RLock과 RUnlock을 사용하지 않고 Lock과 Unlock만 사용하면 Mutex와 동일하다.

## 문맥 전환

문맥 전환 (context switching)이란 프로그램이 여러 프로세스 혹은 스레드에서 동작할 때
기존에 하던 작업들을 메모리에 보관해두고 다른 작업을 싲가하는 것을 말한다.

문맥 전환을 이용하여 프로그램이 병행적으로 수행될 수 있지만 이때 비용이 발생한다.

고루틴은 스레드보다 비용이 낮다. 고루틴을 어러개 만들어도 스레드는 그것보다 적에 만들어지고
어러 개의 고루틴이 하나의 스레드에 대응된다.

Go 컴파일러가 어떤 고루틴들이 하나의 스레드에 묶이는 것이 좋은지를 분석하면 스레드의 문맥전환을 하지 않도록 코드를 생성한다.

Go 컴파일러는 주로 다음의 경우에 문맥 전환을 하는 코드를 생성할 수도 있다.

- 파일이나 네트워크 연산처럼 시간이 오래 걸리는 입출력 연산이 있을 때
- 채널에 보내거나 받을 때
- go로 고루틴이 생성될 때
- 가비지 컬렉션 사이클이 지난 뒤

Go의 입출력은 주로 블럭킹이다. 이때 다른 고루틴으로 문맥 전환하는 것은 매우 좋은 방식이다.

Go 언어를 만든 사람들은 논블록킹 비동기 입출력이 혼란스럽다고 생각하여 동시성이 있는 고루틴을 생성하고
고루틴에서 동기화 입출력을 하면서 필요한 문맥전환을 하는 방식을 이용하여 설계 하였다.

문맥 전환을 강제로 시키기 위한 코드 중 하나는 `time.Sleep(0)`이다.
문맥 전환 없이 끝나지 않는 연산을 수행하는 코드가 있다면 다른 고루틴은 처리 기회를 받지 못할수도 있기 때문이다.
