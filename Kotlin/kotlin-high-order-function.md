# 코틀린 고차함수

고차함수(high order function)는 람다를 인자로 받거나 반환하는 함수이다.
또한 코틀린에서는 람다를 사용함에 있어 성능상 부가비용을 없애고 유연하게 흐름을 제어할 수 있는 inline 함수도 존재한다.

## 고차함수

### 함수 타입

람다를 인자로 받는 함수를 저으이하려면 먼저 람다 인자의 타입을 어떻게 선언할 수 있는지 알아야 한다.

```kt
val sum: (Int, Int) -> Int = { x, y -> x + y }
val action: () -> Unit = { println(40) }
```

1번 예시는 Int 파라미터를 2개 받아 Int를 반환하는 함수이다.
2번 예시는 파라미터를 받지 않고 아무 값도 반환하지 않는 함수이다.

함수 타입을 선언할 때는 반환 타입을 반드시 명시해야 하므로 반환값이 없는 경우라도 `Unit`을 명시해야 한다.

다른 함수와 마찬가지로 함수 타입에서도 반환 타입을 널이 될 수 있는 타입으로 지정할 수 있다.
`var canReturnNull: (Int, Int) -> Int? = { null }`

널이 될 수 있는 함수 타입 변수를 정의할 수도 있다. 이를 위해서 함수 타입을 괄호로 감싸고 그 뒤에 물음표를 붙이면 된다.
`var funOrNull: ((Int, Int) -> Int)? = null`

함수 타입에서 파라미터 이름을 지정할 수도 있다.

```kt
fun performRequest(
  url: String,
  callback: (code: Int, content: String) -> Unit) {
    ...
  }
```

### 인자로 받은 함수 호출

간단한 고차 함수를 정의해보자

```kt
fun twoAndThree(operation: (Int, Int) -> Int) {
  val result = operation(2, 3)
  println("${result}")
}
```

인자로 받은 함수를 호출하는 구문은 일반 함수를 호출하는 구문과 같다.

### 자바에서 코틀린 함수 타입 사용

컴파일된 코드 안에서 함수 타입은 일반 인터페이스로 바뀐다.
즉 함수 타입의 변수는 `FunctionN` 인터페이스를 구현하는 객체를 저장한다.
코틀린 표준 라잉브러리는 인자의 개수에 따라 `Function0<R>`, `Function1<P1, R>` 등의 인터페이스를 제공한다.
각 인터페이스에는 `invoke` 메소드 정의가 들어있고, 호출하면 함수를 실행할 수 있다.

함수 타입인 변수는 인자 개수에 따라 적당한 `FunctionN` 인터페이스를 구현하는 클래스의 인스턴스를 저장하며,
그 클래스의 invoke 메소드 본문에는 람다의 본문이 들어간다.

자바8 이전의 자바에서는 필요한 `FunctionN` 이넡페이스의 `invoke` 메소드를 구현하는 익명 클래스를 넘기면 된다.
자바에서 코틀린 표준 라이브러리가 제공하는 람다를 인자로 받는 확장함수를 쉽게 호출 할 수 있지만,
수신 객체를 확장 함수의 첫 번째 인자로 명시적으로 넘겨야 하므로 코드가 깔끔하지는 않다.

```java
List<String> strings = new ArrayList();
strings.add("42");
CollectionsKt.forEach(strings, s -> {
  System.out.println(s);
  return Unit.INSTANCE;
});
```

### 디폴트 값을 지정한 함수 파라미터, 널이 될 수 있는 함수 파라미터

파라미터를 함수 타입으로 선언할 때도 디폴트 값을 지정할 수 있다.

```kt
fun <T> Collection<T>.joinToString(
    separator: String = ", ",
    prefix: String = "",
    postfix: String = "",
    transform: (T) -> String = { it.toString() }): String { // 디폴트 값을 람다로 지정
  val result = StringBuilder(prefix)

  for ((index, element) in this.withIndex()) {
    if (index >0) result.append(separator)
    result.append(transform(element))
  }

  result.append(postfix)
  return result.toString()
}

>>> val letters = listOf("Alpha", "Beta")
>>> println(letters.joinToString { it.toLowerCase() })
alpha, beta
```

널이 될 수 있는 함수 타입을 사용할 수 있다. 널이 될 수 있는 함수 타입으로 함수를 받으면 그 함수를 직접 호출할 수 없게 된다.

```kt
fun <T> Collection<T>.joinToString(
    separator: String = ", ",
    prefix: String = "",
    postfix: String = "",
    transform: ((T) -> String)? = null): String { // 널이 될 수 있는 함수의 파라미터 선언
  val result = StringBuilder(prefix)

  for ((index, element) in this.withIndex()) {
    if (index >0) result.append(separator)
    val str = transform?.invoke(element) ?: element.toString() // 널이 될 경우를 위해 안전호출과 엘비스 연산자 사용
    result.append(transform(element))
  }

  result.append(postfix)
  return result.toString()
}
```

### 함수를 함수에서 반환

함수를 반환하는 함수 정의하기

```kt
enum class Delivery { STANDARD, EXPEDITED }

class Order(val itemCount: Int)

fun getShippingCostCalculator(
    delivery: Delivery): (Order) -> Double { // 함수를 반환하는 함수
  if (delivery == Delivery.EXPEDITED) {
    return { oreder -> 6 + 2.1 * order.itemCount }
  }
  return { order -> 1.2 * order.itemCount }
}
```

### 람다를 활용한 중복제거

고차 함수를 사용해 중복을 제거한 예

```kt
data class SiteVisit (
    val path: String,
    val duration: Double,
    val os: OS)

enum class OS { WINDOWS, LINUX, MAC, IOS, ANDROID}

val log = listOf(
  SiteVisit("/", 34.0, OS.WINDOWS)
  SiteVisit("/", 22.0, OS.MAC)
  SiteVisit("/login", 12.0, OS.WINDOWS)
  SiteVisit("/signup", 8.0, OS.IOS)
  SiteVisit("/", 16.3, OS.ANDROID)
)

fun List<SiteVisit>.averageDurationFor(predicate: (SiteVisit) -> Boolean) =
    filter(predicate).map(SiteVisit::duration).average()
```

## 인라인 함수

람다를 호출하면 만들어진 익명클래스를 호출하지만, 지역변수를 capture하면 람다가 생성되는 시점마다 익명 클래스가 만들어져 부가비용이 발생한다.
`inline` 변경자를 어떤 함수에 붙이면 컴파일러는 그 함수를 호출하는 모든 문장을 함수 본문에 해당하는 바이트코드로 바꿔치기 해준다.

### 인라이닝이 작동하는 방식

어떤 함수를 `inline`으로 선언하면 함수를 호출하는 코드가 있을때 함수호출 대신 함수 본문이 바이트 코드로 컴파일 된다.

```kt
inline fun<T> synchronized(lock: Lock, action: () -> T): T {
  lock.lock()
  try {
    return action()
  } finally {
    lock.unlock()
  }
}

fun foo(l: Lock) {
  println("Before sync")
  synchronized(1) {
    println("Action")
  }
  println("After sync")
}
```

위의 foo 함수는 다음 코틀린 코드와 동등하다

```kt
fun __foo__(1: Lock) {
  println("Before sync")
  1.lock()
  try {
    println("Action")
  } finally {
    1.unlock()
  }
  println("After sync")
}
```

`synchronized` 함수 본문뿐만 아니라 전달된 람다의 본문도 함께 인라이닝 된다는 점에 유의해야 한다.
바이트코드는 람다를 호출한 코드의 일부분으로 간주하므로, 코틀린 컴파일러는 람다를 익명 클래스로 감싸지 않는다.

만약 인라인 함수를 호출하면서 람다대신 함수 타입의 변수를 넘기면 인라인 함수를 호출하는 코드위치에서는 변수에 저장된 람다의 코드를 알 수없다.
따라서 람다 본문은 인라이닝 되지 않고, `synchronized` 함수 본문만 인라이닝 된다.

```kt
class LockOwner(val lock: Lock) {
  fun runUnderLock(body: () -> Unit) {
    synchronized(lock, body)
  }
}
```

위의 코드는 아래와 동등하다

```kt
class LockOwner(val lock: Lock) {
  fun __runUnderLock__(body: () -> Unit) {
    lock.lock()
    try {
      body()
    } finally {
      lock.unlock()
    }
  }
}
```

### 인라인 함수의 한계

인라이닝을 사용하면 람다가 본문에 직접 펼쳐지기 때문에 함수가 파라미터로 전달받은 람다를 본문에 사용하는 방식이 한정될 수 밖에 없다.
파라미터로 받은 람다를 다른 변수에 저장하고 나중에 그 변수를 사용한다면 람다를 표현하는 객체가 어딘가는 존재해야 하기 때문에 람다를 인라이닝 할 수 없다.

예를 들어 시퀀스에 대해 동작하는 메소드 중에는 람다를 받아서 모든 시퀀스 원소에 그 람다를 적용한 새 시퀀스를 반환하는 함수가 많다.
그런 함수는 인자로 받은 람다를 시퀀스 객체 생성자의 인자를 넘기곤 한다.

```kt
fun <T, R> Sequence<T>.map(transform: (T) -> R): Sequence<R> {
  return TransformingSequence(this, transform)
}
```

위의 `map` 함수는 `transform` 파라미터로 전달받은 함수값을 호출하지 않고, `TransformingSequence` 클래스 생성자에 함수 값을 넘긴다.
`TransformingSequence` 클래스는 전달 받은 람다를 프로퍼티로 저장되는데, 이런기능을 지원하려면 `transform`을 인라이닝 하지 않는 함수표현으로 만들 수 밖에 없다.

이런식으로 인라이닝하면 안 되는 람다를 파라미터로 받는다면 `noinline` 변경자를 파라미터 이름앞에 붙여서 인라이닝을 금지할 수 있다.

```kt
inline fun foo(inlined: () -> Unit, noinline notInlined: () -> Unit) {
  ...
}
```

### 컬렉션 연산 인라이닝

코틀린의 `filter`, `map` 함수는 인라인 함수이다. 따라서 성능상에 큰 신경을 쓸 필요가 없다.

만약 처리할 원소가 많아지면 중간리스트에 대한 부가비용이 발생한다.
이때 `asSequence`를 통해 리스트 대신 시퀀스를 사용하면 중간 리스트로 인한 부가비용이 줄어든다.
이때 각 중간 시퀀스는 람다를 필드에 저장하는 객체로 펴현되며, 최종 연산은 중간 시퀀스에 있는 여러 람다를 연쇄 호출한다.
따라서 시퀀스는 람다를 인라인 하지 않는다.

따라서 지연계산을 통해 성능을 향상시키려는 이유로 모든 컬렉션 연산에 `asSequence`를 붙여서는 안된다.
시퀀스 연산에서는 람다가 인라이닝 되지 않기 때문에 크기가 작은 컬렉션은 오히려 성능이 떨어질 수 있다.

### 함수를 인라인으로 선언해야 하는 경우

`inline` 키워드를 사용해도 람다를 인자로 받는 함수만 성능이 좋아질 가능성이 높다.

일반 함수 호출의 경우 JVM은 이미 강력하게 인라이닝을 지원한다.
이런 과정은 바이트코드를 실제 기계어 코드로 번역하는 과정(JIT)에서 일어난다.
이런 JVM의 최적화를 활용한다면 바이트코드에서는 각 함수 구현이 정확히 한 번만 있으면 되고 그 함수를 호출 하는 부분에서 따로 함수코드를 중복할 필요가 없다.

그러나 코틀린 인라인 함수는 바이트 코드에서 각 함수 호출 지점을 함수 본문으로 대치하기 때문에 코드 중복이 발생한다.

반면 람다를 인자로 받는 함수를 인라이닝하면 이익이 더 많다.
함수 호출 비용을 줄일 수 있을 뿐 아니라 람다를 표현하는 클래스와 람다 인스턴스에 해당하는 객체를 만들 필요도 없어진다.
인라이닝을 사용하면 일반 람다에는 사용할 수 없는 몇 가지 기능을 사용할 수 있다. (non-local 반환 등)

`inline` 변경자를 함수에 붙이면 바이트코드가 아주 커질 수 있다. 코틀린 표준 라이브러리가 제공하는 `inline`함수를 보면 모두 크기가 아주작다는 사실을 알 수 있다.

### 자원 관리를 위해 인라인된 람다 사용

람다로 중복을 없앨 수 있는 일반적인 패턴 중 하나는 어떤 작업을 시작하기전 자원을 획득하고 작업을 마친 후 자원을 해제하는 자원관리이다.
코틀린에서는 앞에서 살펴본 `Lock` 인터페이스의 확장함수인 `withLock`을 사용하는 방식도 이와 같다.

자바7부터 지원하는 `try-with-resource` 문과 코틀린 코드를 비교해보자

```java
static String readFirstLineFromFile(String path) throws IOException {
  try (BufferedReader br = new BufferedReader(new FileReader(path))) {
    return br.readLine();
  }
}
```

코틀린에서는 자바 `try-with-resource`와 같은 기능을 제공하는 `use`라는 함수가 표준 라이브러리에 들어있다

```kt
fun readFirstLineFromFile(path: String): String {
  BufferedReader(FileReader(path)).use { br ->
    return br.readLine()
  }
}
```

`use` 함수는 `closeable` 자원에 대한 확장 함수이며 람다를 인자로 받는다. 물론 `use` 함수도 인라인 함수이다.
람다의 본문 안에서 사용한 `return`은 non-local `return`이다.
이 `return`문은 람다가 아니라 `readFirstLineFromFile` 함수를 끝내면서 값을 반환한다.

## 고차 함수 안에서 흐름 제어

### 람다 안의 return 문

람다안에서 `return`을 사용하면 람다로부터만 반환되는 것이 아니라 그 람다를 호출하는 함수가 실행을 끝내고 반환된다.
자신을 둘러싸고 있는 블록보다 더 바깥에 있는 다른 블록을 반환하게 만드는 `return`문을 non local return 이라고 부른다.

이렇게 `return`이 바깥쪽 함수를 반환시킬 수 있는 때는 람다를 인자로 받는 함수가 인랑니 함수인 경우 뿐이다.
하지만 인라이닝 되지 않는 함수에 잔달되는 람다 안에서 return을 사용할 수 는 없다.

인라이닝 되지 않는 함수는 람다를 변수에 저장할 수 있고, 바깥쪽 함수로부터 반환된 뒤에 저장해 둔 람다가 호출될 수도 있다.
그런 경우 람다 안의 `return`이 실행되는 시점이 바깥쪽 함수를 반환하기엔 너무 늦은 시점일 수 있다.

### 레이블을 사용한 return

람다식에서도 local return을 사용할 수 있다.
람다안에서 local return은 `for` 루프의 `break`와 비슷한 역할을 한다.
local return 은 람다의 실행을 끝내고 람다를 호출했던 코드의 실행을 계속 이어간다.

local return 과 non-local return을 구분하기 위해 label을 사용해야 한다.
`return`으로 실행을 끝내고 싶은 람다 식 앞에 레이블을 붙이고, `return` 키워드 뒤에 그 레이블을 추가하면 된다.

```kt
fun lookForAlice(people: List<Person>) {
  people.forEach {
    if (it.name == "Alice") return@forEach
  }
  println("Alice might be somewhere") // 항상 이 줄이 실행된다
}
```

위의 예제를 보면 `forEach`는 `inline` 람다임에도 불구하고 `return` 이후 함수가 끝나지 않는다.
`label`을 추가하면 local return을 사용할 수 있다.

하지만 non-local 반환문은 장황하고, 람다 안의 여러 위치에 return 식이 들어가야 하는 경우 사용하기 불편하다.

익명함수를 사용하면 non-local 반환문을 쉽게 작성할 수 있다.
익명함수 안에서 레이블이 붙지 않은 return 식은 무명함수 자체를 반환시킬 뿐 무명함수를 둘러싼 다른 함수를 반환시키지 않는다.

```kt
fun lookForAlice(people: List<Person>) {
  people.forEach(fun (person) {
    if (person.name == "Alice") return
    println("${person.name} is not Alice")
  })
}
```

즉 `return` 식은 `fun` 키워드로 정의된 함수를 반환하는 것이라 할 수 있다.
