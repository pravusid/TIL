# 람다

## 람다 식과 멤버 참조

일련의 동작을 변수에 저장하거나 다른 함수에 넘겨야 하는 경우가 자주 있다.
예전 자바에서는 익명 내부 클래스를 통해 이런 목적을 달성했지만, 문법이 상당히 번거롭ㄴ다.

이와 달리 함수형 프로그래밍에서는 함수를 값 처럼 다루는 접근 방법을 택함으로써 이 문제를 해결한다.
클래스를 선언하고 그 클래스의 인스턴스를 함수에 넘기는 대신 함수형 언어에서는 함수를 직접 다른 함수에 전달할 수 있다.

람다식을 사용하면 함수를 선언할 필요가 없고 코드블록을 직접 함수의 인자로 전달할 수 있다.

## 람다와 컬렉션

람다가 없다면 컬렉션을 편리하게 처리할 수 잇는 좋은 라이브러리를 제공하기 힘들다.

`data class Person(val name: String, val age: Int)`
사람들로 이루어진 리스트가 있고 그중 가장 연장자를 찾으려 한다. 람다를 사용하지 않는다면 루프를 써서 직접 검색을 구현할 것이다.

하지만 람다를 사용한 라이브러리 함수를 쓰면 간결하게 결과를 구할 수 있다. 모든 컬렉션에 대해 maxBy 함수를 호출할 수 있다.

```kotlin
>>> val people = listOf(Person("Alice", 29), Person("Bob", 31))
>>> println(people.maxBy { it.age })
Person(name=Bob, age=31)
```

위의 코드를 멤버 참조를 활용하여 다음과 같이 사용할 수도 있다.
`people.maxBy(Person::age)`

## 람다식의 문법

`{ x: Int, y: Int -> x + y }`

코틀린 람다식은 항상 중괄호로 둘러싸여 있다. 또한 인자 목록 주변에 괄호가 없다.
람다식을 변수에 저장하여, 다른 일반 함수와 마찬가지로 다룰 수 있다.

람다식을 직접 호출할 수도 있지만 읽기도 불편하고 쓸모도 없다.
코틀린에서 `run()` 함수를 사용하면 인자로 받은 람다를 실행할 수 있다.

```kotlin
>>> { println(42) }()
42
>>> run { println(42) }
42
```

앞에서 작성한 연장자를 찾는 코드를 람다로 다시 쓰면 다음과 같다
`people.maxBy({ p: Person -> p.age })`

코틀린에서는 함수호출시 맨뒤에 있는 인자가 람다식이라면 해당 람다를 괄호 밖으로 빼낼 수 있다.
`people.maxBy() { p: Person -> p.age }`

람다가 함수의 유일한 인자이고 괄호 뒤에 람다를 썼다면 호출시 빈 괄호는 생략해도 된다.
`people.maxBy { p: Person -> p.age }`

로컬 변수처럼 컴파일러는 람다 파라미터의 타입도 추론한다.
`people.maxBy { p -> p.age }`

람다의 파라미터가 하나뿐이고 타입을 컴파일러가 추론할 수 있을경우 디폴트 파라미터인 `it`을 쓸 수 있다.
`people.maxBy { it.age }`

> 람다를 변수에 저장할 때는 파라미터의 타입을 추론할 문맥이 없으므로, 타입을 명시해야 한다.

**본문이 여러 줄로 이루어진 람다의 경우 본문의 맨 마지막에 있는 식이 람다의 결과 값이 된다.**

### 현재 영역에 있는 변수에 접근

중첩 클래스와 동일하게 람다에서도 파이널 변수가 아닌 로컬변수에 접근하여 변경할 수 있다.
마찬가지로 람다를 함수 안에서 정의하면 함수의 파라미터뿐 아니라 람다 정의에 앞서 선언된 로컬 변수까지 람다에서 접근가능하다.

```kotlin
fun printMessageWithPrefix(messages: Collection<String>, prefix: String) {
  message.forEach { // 람다식
    println("${prefix} ${it}")
  }
}
```

이와 같이 람다 안에서 사용하는 외부 변수를 람다가 capture 한 변수라고 부른다.
기본적으로 함수 안에 정의된 로컬변수의 생명주기는 함수 반환과 함께 끝나지만,
capture 변수가 있는 람다를 저장해서 함수 종료후에 실행해도 람다의 본문코드는 capture 변수를 사용할 수 있다.

> 코틀린에서 파이널변수(val)을 capture하면 변수의 값을 복사하고, 변경가능변수(var)를 capture하면 변수를 Ref 클래스 인스턴스에 넣는다. 그 Ref 인스턴스에 대한 참조를 파이널로 만들면 쉽게 람다로 capture할 수 있고, 람다에서는 Ref 인스턴스의 필드를 변경할 수 있다.

### 멤버 참조

넘기려는 코드가 이미 함수로 선언된 경우, 함수를 호출하는 람다를 사용하지 않고 함수를 직접 넘기는 방법을 사용한다.

코틀린에서는 자바8과 마찬가지로 함수를 값으로 바꿀 수 있다.
이중 콜론 `::`을 사용한 식을 Member reference 라고 부른다. 멤버참조는 프로퍼티나 메소드를 단 하나만 호출하는 함수 값을 만들어 준다.

Member reference: `클래스::멤버`

참조대상이 함수/프로퍼티와 관계없이 참조 뒤에는 괄호를 넣으면 안된다.
멤버 참조는 그 멤버를 호출하는 람다와 같은 타입이다. 따라서 다음과 같이 바꿔 쓸 수 있다. (에타변환)

```kotlin
people.maxBy(Person::age)
people.maxBy { p -> p.age }
people.maxBy { it.age }
```

또한 멤버가 아닌 최상위 함수나 프로퍼티를 참조할 수도 있다.

```kotlin
fun salute() = println("Salute!")
>>> run(::salute)
Salute!
```

람다로 인자가 여럿인 다른 함수에게 작업을 위임하는 경우 람다를 정의하지 않고 직접 위임함수에 대한 참조를 제공할 수 있다.

```kotlin
// case1
val action = { person: Person, message: String ->
  sendEmail(person, message)
}
// case2
val nextAction = ::sendEmail
```

생성자 참조를 사용하면 클래스 생성 작업을 연기하거나 저장해둘 수 있다.
`::` 뒤에 클래스 이름을 넣으면 생성자 참조를 만들 수 있다.

```kotlin
data class Person(val name, String, val age: Int)

>>> val createPerson = ::Person // 생성자 참조를 변수에 할당한다
>>> val p = createPerson("Alice", 29)
>>> println(p)
Person(name=Alice, age=29)
```

확장함수도 멤버변수와 똑같은 방식으로 참조할 수 있다.

```kotlin
fun Person.isAdult() = age > 21
val predicate = Person::isAdult
```

코틀린 1.1 부터는 참조를 얻은 다음 참조를 호출할 때 인스턴스 객체 제공을 바운드 멤버참조로 할 수 있다.

```kotlin
val p = Person.("Dmitry", 34)
val personAgeFunction = Person::age

// 바운드 멤버참조로 고쳐쓰면 다음과 같다
val dmitryAgeFunction = p::age
```

## 컬렉션 함수형 API

함수형 프로그래밍 스타일을 사용하면 컬렉션을 다룰 때 편리하다.
이런 스타일은 C#, 그루비, 스칼라 등 람다를 지원하는 대부분의 언어에서 찾아볼 수 있다.

### filter와 map

filter 함수는 컬렉션을 이터레이션 하면서 주어진 람다에 각 원소를 넘기고 람다가 true를 반환하는 원소만 모은다.

```kotlin
val list = listOf(1, 2, 3, 4)

>>> println(list.filter { it % 2 == 0 })
[2, 4]
```

filter 함수는 컬렉션에서 원치 않는 원소를 제거하지만 원소를 변환할 수는 없다.
원소를 변환하기 위해서 map 함수를 사용해야 한다.

map 함수는 주어진 람다를 컬렉션의 각 원소에 적용한 결과를 모아서 새 컬렉션을 만든다 (mapping)

```kotlin
val list = listOf(1, 2, 3, 4)

>>> println(list.map { it * it })
[1, 4, 9, 16]
```

맵(dictionary 자료형)의 경우 key와 value를 처리하는 함수가 따로 존재한다.
`filterKeys`와 `mapKeys`는 키를 걸러내거나 변환하고 `filterValues`와 `mapValues`는 값을 걸러내거나 변환한다.

### all, any, count, find

컬렉션의 원소가 어떤 조건을 만족하는지 판단하는 연산이 있다.
코틀린에서 `all`과 `any`가 그런 연산이다.
`count`함수는 조건을 만족하는 원소의 개수를 반환하며, `find`함수는 조건을 만족하는 첫 번째 원소를 반환한다.

```kotlin
val canBeInClub27 = { p: Person -> p.age <= 27 }
val people = listOf(Person("Alice", 27), Person("Bob", 31))

>>> println(people.all(canBeInClub27))
false
>>> println(people.any(canBeInClub27))
true
```

어떤 조건에 대해 `!all`을 수행한 결과와 그 조건의 부정에 대해 `any`를 수행한 결과는 같다.
또 어떤 조건에 대해 `!any`를 수행한 결과와 그 조건의 부정에 대해 `all`을 수행한 결과도 같다.
가독성을 높이려면 `any`와 `all`앞에 `!`를 붙이지 않는 편이 좋다.

조건을 만족하는 원소의 개수를 구하려면 count를 사용한다.

```kotlin
val people = listOf(Person("Alice", 27), Person("Bob", 31))
>>> println(people.count(canBeInClub27))
1
```

> `count`를 사용하지 않고 컬렉션을 필터링한 결과의 크기를 가져오는 경우가 있다. 하지만 그렇게 처리를 하면 조건을 만족하는 모든 원소가 들어가는 중간 컬렉션이 생성된다. 중간 연산결과가 필요한 것이 아니라면 `count`가 더 효율적이다.

조건을 만족하는 원소를 하나 찾고 싶으면 `find` 함수를 사용한다.

```kotlin
val people = listOf(Person("Alice", 27), Person("Bob", 31))
>>> println(people.find(canBeInClub27))
Person(name=Alice, age=27)
```

### groupBy

`groupBy`는 특성을 파라미터로 전달하면 컬렉션을 자동으로 구분해주는 함수이다.

```kotlin
val list = listOf("a", "ab", "b")
>>> println(list.groupBy(String::first))
{a=[a, ab], b=[b]}
```

### flatMap과 flatten

`flatMap` 함수는 먼저 인자로 주어진 람다를 컬렉션의 모든 객체에 적용하고 (map),
 람다를 적용한 결과 얻어지는 여러 리스트를 한 리스트로 모은다 (flatten).

 ```kotlin
val strings = listOf("abc", "def")
>>> println(strings.flatMap { it.toList() })
[a, b, c, d, e, f]
```

flatMap으로 리스트에 들어있는 원소를 단일 리스트로 변환하여 반환한다. (flatten: string to char, map: char to list)

반환할 내용 없이 리스트를 펼치기만 하면 된다면 `flatten()` 함수를 사용하면 된다.

### lazy 컬렉션 연산

앞에서 사용한 `map` 이나 `filter` 같은 컬렉션 함수들은 결과를 즉시(eagerly) 생성한다.
sequence를 사용하여 연쇄적으로 컬렉션 연산을 할 수 있다.

`people.map(Person::name).filter { it.startWith("A") }`

두 번의 컬렉션 연산을 하면 두 개의 결과 리스트가 생성된다.
결과로 하나의 컬렉션만 필요하다면 보다 효율적인 연산방법이 필요하다.
이를 위해 지연계산을 할 수 있고, 지연 계산은 `Sequence` 인터페이스를 활용한다.

우선 컬렉션을 시퀀스로 변환하고 결과를 구한 뒤 컬렉션으로 변환하는 과정이 필요하다.

`people.asSequence().map(Person::name).filter { it.startWith("A") }. toList()`

시퀀스의 원소는 필요할 때 계산된다(지연계산).
따라서 시퀀스에서 계산을 실행하게 만드려면 최종 시퀀스의 원소를 하나씩 이터레이션 하거나 시퀀스를 리스트로 변환해야 한다.

지연계산을 위해서 시퀀스는 중간연산과 최종연산으로 나누어진다.
중간연산은 다른 시퀀스를 반환하고, 최종연산은 결과를 반환한다. 중간연산은 항상 지연 계산된다.
결과는 최초 컬렉션을 변환한 시퀀스로부터 일련의 계산을 수행해 얻을 수 있는 컬렉션이나 원소, 숫자, 객체이다.

> 자바8의 스트림의 개념이 시퀀스와 같다. 코틀린에서 같은 개념을 별도로 구현해서 제공하는 이유는 자바6 기반이기 때문이다. 자바8에서는 코틀린 컬렉션과 시퀀스에서 제공하지 않는 스트림 병렬실행이 있으므로 적절한 쪽을 선택해서 사용하는 것이 좋다.

#### generateSequence()

시퀀스를 생성할 때 `asSequence()` 함수 뿐만 아니라 `generateSequence()` 함수를 사용할 수 있다.
이런 시퀀스를 사용하는 용례중 하나는 객체의 조상으로 이루어진 시퀀스를 만들어 내는 것이다.

다음은 어떤 파일의 상위 디렉토리를 찾으면서 숨김 속성을 가진 디렉토리에 파일이 들어있는지 검사하는 것이다.

```kotlin
fun File.isInsideHiddenDirectory() =
  generateSequence(this) { it.parentFile }.any { it.isHidden }

>>> val file = File("Users/svtk/.HiddenDir/a.txt")
>>> println(file.isInsideHiddenDirectory())
true
```

## 자바 함수형 인터페이스 활용

자바로 작성된 API에 코틀린 람다를 사용해도 문제가 없다.
함수형 인터페이스(functional interface) 또는 SAM(single abstract method) 인터페이스에 람다를 사용할 수 있다.

```java
void postponeComputation(int delay, Runnable computation);
```

코틀린에서 람다를 위 함수에 넘길 수 있다. 또한 객체식을 함수형 인터페이스 구현으로 넘길 수도 있다.

```kotlin
postponeComputation(1000) { println(42) }

postponeComputation(1000, object: Runnable {
  override fun run() {
    println(42)
  }
})
```

하지만 람다와 익명 객체 사이에는 차이가 있다. 객체를 명시적으로 선언하는 경우 메소드를 호출할 때 마다 새로운 객체가 생성된다.
람다의 경우 람다에 대응하는 무명객체를 메소드 호출마다 반복 사용한다.
만약 람다가 지역변수를 capture하는 경우 매 호출마다 새로운 인스턴스를 생성해준다.

컬렉션을 확장한 메소드에 람다를 넘기는 경우(코틀린 자체의 함수),
즉 코틀린 inline으로 표시된 코틀린 함수에 람다를 넘기면 아무런 무명 클래스도 만들어지지 않는다.
대부분의 코틀린 확장 함수에는 inline 표시가 붙어있다.

### 람다를 함수형 인터페이스로 명시적 변경

SAM 생성자는 람다를 함수형 인터페이스 인스턴스로 변환할 수 있게 컴파일러가 자동으로 생성한 함수이다.
컴파일러가 자동으로 람다를 함수형 인터페이스 익명 클래스로 바꾸지 못하는 경우 SAM 생성자를 사용할 수 있다.

```kotlin
fun createAllDoneRunnable(): Runnable {
  return Runnable {
    println("All done!")
  }
}

>>> createAllDoneRunnable().run()
All done!
```

SAM 생성자의 이름은 사용하려는 함수형 인터페이스 이름과 같다.
SAM 생성자는 함수형 인터페이스의 유일한 추상 메소드의 본문에 사용할 람다만을 인자로 받아서 함수형 인터페이스를 구현하는 클래스의 인스턴스를 반환한다.

람다로 생성한 함수형 인터페이스 인스턴스를 변수에 저장해야 할 때도 SAM 생성자를 사용할 수 있다.
여러 버튼에 같은 리스너를 적용하는 상황에서 SAM 생성자를 통해 람다를 함수형 인터페이스 인스턴스로 만들어 변수에 저장해 활용할 수 있다.

```kotlin
val listener = onClickListener { view ->
  val text = when (view.id) {
    R.id.button1 -> "first button"
    R.id.button2 -> "second button"
    else -> "unknown button"
  }
  toast(text)
}
button1.setOnClickListener(listener)
button2.setOnClickListener(listener)
```

`OnClickListener`를 구현하는 객체 선언을 할 수도 있지만 SAM 생성자를 쓰는쪽이 더 간결한다.

람다에는 익명 객체와는 달리 인스턴스 자신을 가리키는 `this`가 없다는 사실에 유의해야 한다. 컴파일러 입장에서 보면 람다는 코드 블록일 뿐이고, 객체가 아니므로 객체처럼 람다를 참조할 수는 없다. 람다 안에서 `this`는 람다를 둘러싼 클래스의 인스턴스를 가리킨다.

이벤트 리스너가 이벤트를 처리하다가 자신의 리스너 등록을 해제해야 한다면, 익명 객체(object)를 사용하여 리스너를 구현하면된다. 익명객체 인스턴스 내부에서는 `this`가 객체 자신을 가리키기 때문이다.

람다를 사용하면 대부분의 SAM 변환을 컴파일러가 자동으로 수행하지만, 가끔 오버로드 한 메소드 중 어떤 타입의 메소드를 선택해 변환후 넘겨줘야 할찌 모호한 때가 있다.
그럴경우 명시적으로 SAM 생성자를 적용하면 컴파일러 오류를 피할 수 있다.

## Lambda with receiver: let, with, run, apply

수신객체 지정 람다는 수신 객체를 명시하지 않고 람다의 본문 안에서 다른 객체의 메소드를 호출할 수 있는 방식이다.

### let 함수

`let` 함수를 호출한 객체를 이어지는 블록의 인자로 넘기고, 블록의 결과값을 반환함: `fun <T, R> T.let(block: (T) -> R): R`

커스텀 뷰에서 Padding 값을 지정할 때 let을 사용해서 다음과 같이 작성할 수 있다.

```kotlin
TypedValue.applyDimension(TypedValue.COMPLEX_UNIT_DIP, 16f,
    resources.displayMetrics).toInt().let { padding ->
      setPadding(padding, 0, padding, 0)
    }
```

인자가 하나인 경우 `it`을 사용할 수 있다.

let()을 안전한 호출(Safe Calls - ?.)과 함께 사용하면 `if (null != obj) ...` 를 대체할 수 있다.

### with 함수

어떤 객체의 이름을 반복하지 않고 그 객체에 대한 다양한 연산을 수행할 수 있는 함수이다: `fun <T, R> with(receiver: T, block: T.() -> R): R`

```kotlin
fun alphabet(): String {
  val result = StringBuilder()
  for (letter in 'A'..'Z') {
    result.append(letter)
  }
  return result.toString()
}
```

위의 예제를 `with`를 사용하여 작성해 보자

```kotlin
fun alphabet(): String {
  return with(StringBuilder()) {
    for (letter in 'A'..'Z') {
      this.append(letter) // this 생략 가능
    }
    this.toString() // 람다에서 값을 반환한다 (마지막 줄)
  }
}
```

> 람다는 일반함수와 비슷한 동작을 정의하는 한 방법이다. 수신 객체 지정 람다는 확장 함수와 비슷한 동작을 정의하는 한 방법이다.

with에게 인자로 넘긴 객체의 클래스와 with를 사용하는 코드가 들어있는 클래스 안에 이름이 같은 메소드가 있으면, 호출할 대상을 명시해야 한다.
바깥쪽 클래스(Outer)에 정의된 메소드를 호출하고싶다면 `this@Outer.method()` 처럼 써야한다.

### run 함수

`run` 함수는 인자가 없는 익명 함수처럼 동작하는 형태와 객체에서 호출하는 형태 총 두 가지가 있다.

객체 없이 `run` 함수를 사용하면 인자 없는 익명 함수처럼 사용할 수 있다
`fun <R> run(block: () -> R): R`

이어지는 블럭 내에서 처리할 작업들을 넣어줄 수 있으며, 일반 함수와 마찬가지로 반환값을 지정할 수 있다.

객체에서 `run` 함수를 호출할 경우, 호출하는 객체를 이어지는 블록의 리시버로 전달하고, 블록의 결과값을 반환한다.
`fun <T, R> T.run(block: T.() -> R): R`

객체에서 이 함수를 호출하는 경우 객체를 리시버로 전달받으므로, 특정 객체의 메서드나 필드를 연속적으로 호출하거나 값을 할당할 때 사용한다.

`run` 메서드에서도 안전한 호출(Safe Calls)를 사용할 수 있다.

```kotlin
override fun onCreate(savedInstanceState: Bundle?) {
  supportActionBar?.run {
    setDisplayHomeAsUpEnabled(true)
    setHomeAsUpIndicator(R.drawable.ic_clear_white)
  }
}
```

`with` 함수는 사실상 `run` 함수와 기능이 거의 동일하며, 리시버로 전달할 객체가 어디에 위치하는지만 다르다.
`run` 함수는 `let` 함수와 `with` 함수를 합쳐놓은 형태로 볼 수 있다.

```kotlin
supportActionBar?.let {
  with(it) {
    setDisplayHomeAsUpEnabled(true)
    setHomeAsUpIndicator(R.drawable.ic_clear_white)
  }
}
```

`with` 함수는 안전한 호출을 지원하지 않으므로 `run` 함수를 통해 해당 기능을 활용할 수 있다.

### apply 함수

`apply` 함수는 `with`와 거의 같지만, `apply`는 항상 자신에게 전달된 수신객체를 반환한다: `fun <T> T.apply(block: T.() -> Unit): T`

앞의 예제를 `apply`를 사용하여 고쳐보자

```kotlin
fun alphabet() = StringBuilder().apply {
  for (letter in 'A'..'Z') {
    append(letter)
  }
}.toString()
```

`apply`는 확장함수로 정의돼 있다. 이런 `apply`함수는 객체의 인스턴스를 만드는 즉시 프로퍼티 일부를 초기화 하는 경우 유용하다.

앞의 예제는 표준 라이브러리의 `buildString` 함수를 사용하여 단순화 할 수 있다.
`buildString`의 인자는 수신객체 지정 람다이며 수신객체는 항상 `StringBuilder`가 된다

```kotlin
fun alphabet() = buildString {
  for (letter in 'A'..'Z') {
    append(letter)
  }
}
```
