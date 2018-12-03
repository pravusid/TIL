# 코틀린 제네릭스

## 제네릭 타입 파라미터

제네릭스를 사용하면 type parameter를 받는 타입을 정의할 수 있다.

코틀린 컴파일러는 보통 타입과 마찬가지로 타입 인자도 추론할 수 있다.
`val authors = listOf("Dmitry", "Svetlana")`

빈 리스트를 만든다면 타입인자를 추론할 근거가 없으므로 직접 타입인자를 명시해야 한다.
`val readers: MutableList<String> = mutableListOf()` 또는 `val readers = mutableListOf<String>()`

### 제네릭 함수와 프로퍼티

특정 타입을 저장하는 리스트뿐 아니라 제네릭 리스트를 다룰 수 있는 함수가 필요하다.
제네릭을 다루는 함수의 예를 살펴보자

```kotlin
fun <T> List<T>.slice(indices: IntRange): List<T>
```

함수의 타입 파라미터 `T`가 수신 객체와 반환타입에 쓰인다.
이런 함수를 호출할 때 타입 인자를 명시적으로 지정할 수 있지만, 컴파일러가 추론하므로 불필요 한 경우가 많다.

```kotlin
>>> val letters = ('a'..'z').toList()
>>> println(letters.slice<Char>(0..2))
[a, b, c]
-- 추론을 사용한다면
>>> println(letters.slice(10..13))
[k, l, m, n]
```

마찬가지로 제네릭 확장 프로퍼티를 선언할 수도 있다.

```kotlin
val <T> List<T>.penultimate: T
  get() = this[size -2]

>>> println(listOf(1, 2, 3, 4).penultimate)
3
```

하지만 확장 프로퍼티만 제네릭하게 만들 수 있고, 일반 프로퍼티는 불가능하다.

### 제네릭 클래스

클래스에서도 제네릭을 선언할 수 있다.

```kotlin
class StringList: List<String> {
  override fun get(index: Int): String = ...
}
class ArrayList<T> : List<T> {
  override fun get(index:Int): T = ...
}
```

클래스가 자기 자신을 타입인자로 참조할 수도 있다.

```kotlin
interface Comparable<T> {
  fun compareTo(other: T): Int
}
class String : Comparable<String> {
  override fun compareTo(other: String): Int = ...
}
```

### 타입파라미터 제약

Type parameter constraint는 클래스나 함수에 사용할 수 있는 타입 인자를 제한하는 기능이다.
어떤 타입을 제네릭 타입의 upper bound로 지정하면 제네릭 타입을 사용할 때의 타입인자는 반드시 upper bound와 같거나 그 타입의 하위타입이어야 한다.

타입 파라미터 제약을 위해서는 `:`를 쓰고 upper bound를 뒤에 쓰면 된다.

upper bound를 설정하면 제네릭을 선언한 함수 내부에서 해당 타입은 upper bound 타입에 명시된 method를 호출할 수 있다.

```kotlin
fun <T : Number> oneHalf(value: T): Double {
  return value.toDouble()
}
```

드물게 타입 파라미터에 둘 이상의 제약을 가해야 하는 경우도 있다. 그럴 땐 약간 다른 구문을 사용한다.

```kotlin
fun <T> ensureTrailingPeriod(seq: T) where T : CharSequence, T : Appendable {
  if (!seq.endsWith('.')) {
    seq.append('.')
  }
}
```

### 타입 파라미터를 널이 될 수 없는 타입으로 한정

아무런 상한을 정하지 않은 타입 파라미터는 결과적으로 `Any?`를 상한으로 정한 파라미터와 같다

```kotlin
class Processor<T> {
  fun process(value: T) {
    value?.hashCode() // T는 널이 될 수 있는 타입이므로 안전한 호출을 사용해야 함
  }
}
```

널이 되지 않으면서 제약이 필요없다면 `Any`를 상한으로 사용하면 된다.

```kotlin
class Processor<T : Any> {
  fun process(value: T) {
    value.hashCode()
  }
}
```

## 실행시 제네릭스의 동작 : 소거된 타입 파라미터와 실체화된 타입 파라미터

JVM의 제네릭스는 타입 소거(type erasure)를 사용해서 구현된다. 즉 실행시점에 제네릭 클래스의 인스턴스에 타입 인자 정보는 들어있지 않다.

코틀린에서는 함수를 inline으로 선언하면 타입인자가 지워지지 않게 할 수 있고 이를 실체화(reify)라고 한다.

### 실행시점의 제네릭: 타입 검사와 캐스트

자바와 마찬가지로 코틀린 제네릭 타입 인자 정보는 런타임에 지워진다.
이는 제네릭 클래스 인스턴스가 그 인스턴스를 생성할 때 쓰인 타입인자에 대한 정보를 유지하지 않는다는 뜻이다.

```kotlin
>>> if (value is List<String>) { ... }
ERROR: Cannot check for instance of erased type
```

즉 다음 두 리스트는 실행시점에 완전히 같은 타입이 된다.

```kotlin
val list1: List<String> = listOf("a", "b")
val list2: List<Int> = listOf(1, 2)
```

물론 컴파일 단계에서 컴파일러가 타입인자를 인식하고 맞는 값만 넣도록 보장하기 때문에 내부의 값은 선언한 타입과 같다.

만약 타입 파라미터가 2개 이상이라면 `*` (star projection)을 사용한다.
즉, 인자의 타입을 알 수 없는 경우 `if (value is List<*>) { ... }` 와 같이 사용하게 된다.

`as` 나 `as?` 캐스팅에서도 제네릭 타입을 사용할 수 있다.
실행시점에는 제네릭의 타입인자를 알 수 없으므로 캐스팅은 항상 성공하고, 컴파일러가 unchecked cast라는 경고를 해준다. (컴파일은 진행된다)

### 실체화한 타입 파라미터를 사용한 함수

코틀린 제네릭 타입인자 정보는 실행시점에 삭제되지만, 인라인 함수의 타입 파라미터는 실체화 되므로 실행시점에 알 수 있다.
컴파일 이후 타입 파라미터가 지워지지 않음을 표시하기 위해 제네릭 앞에 `reified` 키워드를 붙인다.

```kotlin
inline fun<reified T> isA(value: Any) = value is T
>>> print(isA<String>)("abc")
true
```

실체화한 타입 파라미터를 사용하는 간단한 예제 중 하나는 표준 라이브러리 함수인 `filterIsInstance`이다

```kotlin
inline fun <reified T> iterable<*>.filterIsInstance(): List<T> {
  val destination = mutableListOf<T>
  for (element in this) {
    if (element is T) {
      destination.add(element)
    }
  }
  return destination
}

val items = listOf("one", 2, "three")
>>> println(items.filterIsInstance<String>())
[one, three]
```

자바 코드에서는 `reified` 타입 파라미터를 사용하는 inline 함수를 호출할 수 없다.

### 실체화한 타입 파라미터로 클래스 참조

`java.lang.Class` 타입인자를 파라미터로 받는 API에 대한 코틀린 어댑터를 구축하는 경우 실체화한 타입파라미터를 자주 사용한다.

`java.lang.Class`를 사용하는 API의 예로 JDK의 `ServiceLoader`가 있다.

표준 자바 API인 `ServiceLoader`를 사용해 서비스를 읽으려면 다음과 같다.

`val serviceImpl = ServiceLoader.load(Service::class.java)`

위의 예를 구체화한 타입 파라미터를 사용하면 다음과 같다.

`val serviceImpl = loadService<Service>()`

타입파라미터를 사용하기 위해 `loadService` 함수를 정의해 보자

```kotlin
inline fun <reified T> loadService() {
  return ServiceLoader.load(T::class.java)
}
```

### 실체화한 타입 파라미터의 제약

현재 코틀린에서 구현된 실체화된 타입파라미터의 명세로는 다음과 같은 경우 사용 가능하다

- 타입 검사와 캐스팅 (is, !is, as, as?)
- 코틀린 리플렉션 API (`::class`)
- 코틀린 타입에 대응하는 `java.lang.Class` 얻기 (`::class.java`)
- 다른 함수를 호출 할 때 타입 인자로 사용

하지만 다음 경우 사용할 수 없다

- 타입 파라미터 클래스의 인스턴스 생성하기
- 타입 파라미터 클래스의 companion object method 호출
- 실체화한 타입 파라미터를 요구하는 함수 호출시, 실체화 하지 않은 타입 파라미터로 받은 타입을 인자로 넘기는 것
- 클래스, 프로퍼티, 인라인 함수가 아닌 함수의 타입 파라미터를 `reified`로 지정 하기

실체화한 타입 파라미터를 인라인 함수에만 사용할 수 있으므로,
실체화한 타입 파라미터를 사용하는 함수는 자신에게 전달되는 모든 람다를 인라이닝 한다.
람다를 인라이닝 하지 않으려면 `noinline` 변경자를 함수 타입 파라미터에 붙이면 된다.

## 변성(variance)

변성은 Subtyping 이론의 영역이다.

객체지향 원칙 중 A 타입이 B 타입의 하위타입이면 B타입의 객체를 A타입의 객체로 교체할 수 있다는 리스코프 치환원칙이 존재한다.
변성은 supertype / subtype을 바라보는 관점에 대한 문제이다.

치환될 타입(대체 이후에도 type-safe 해야함)의 방향성이 어떠한지를 알아보는 것이 변성이다.
코틀린에서 치환해서 사용할 법한 타입은, 클래스의 타입, 제너릭 클래스의 타입, 함수의 타입 정도가 있을 것이다.

각 타입의 객체가 사용되는 상황과 위치에 따라 치환할 수 있는 타입의 관계가 달라진다.

### 변성이 있는 이유

만약 `List<Any>` 타입의 파라미터를 받는 함수에 `List<String>`을 넘기면 안전할까?

`String`은 `Any`를 확장하므로 `Any` 타입의 파라미터에 `String` 값을 넘기면 안전하겠지만,
`List`의 인자로 들어가는 경우는 상황이 다르다.

```kotlin
fun addAnswer(list: MutableList<Any>) {
  list.add(42)
}

>>> val strings = mutableListof("abc", "bac")
>>> addAnswer(Strings)
>>> println(strings.maxBy { it.length }) // 현 라인에서 실행시점 예외가 발생
ClassCastException: Integer cannot be cast to String
```

따라서 `List<Any>` 타입의 파라미터를 받는 함수에 `List<String>`을 넘기는 경우는 두 가지로 나누어 볼 수 있다.

함수가 리스트의 원소를 변경하는 경우 타입 불일치 가능성 때문에 `List<String>`을 넘길 수 없다.
즉, mutable인 경우 제네릭 타입은 무공변적(invariance)이어야 한다.

하지만 원소 추가나 변경이 없는 경우에는 `List<String>`을 넘길 수 있다.
immutable인 경우 제네릭의 타입은 공변적(variance)이라고 할 수 있다.

### 클래스, 타입, 하위타입

제네릭 클래스가 아닌 클래스에서는 클래스 이름을 바로 타입으로 쓸 수 있다. `String`, `String?`

제네릭 클래스에서는 보다 복잡하다. 올바른 타입을 얻으려면 제네릭 타입의 타입 파라미터를 구체적인 타입 인자로 바꿔줘야 한다.
`List`는 클래스이지만 타입은 아니다. 하지만 타입인자를 치환한 `List<Int>`, `List<String?>` 등은 타입이다.

어떤 타입 A의 값이 필요한 장소에 어떤 타입 B의 값을 넣어도 아무 문제가 없다면, B는 A위 subtype 이다.
supertype은 subtype의 반대 개념이다.

컴파일러는 변수 대입이나 함수 인자 전달 시 subtype 검사를 매번 수행한다.

간단한 경우 subtype은 subclass와 근본적으로 같다.
널이 될 수 없는 타입은 널이 될 수 있는 타입의 하위 타입이다. 그러나 두 타입 모두 같은 클래스에 해당한다.

제네릭 타입에 대해 이야기할 때 하위 클래스와 하위 타입의 차이가 중요해진다.

제네릭 타입을 인스턴스화할 때 타입 인자로 서로 다른 타입이 들어가고
서로 다른 타입으로 생성된 인스턴스 사이에 타입관계가 성립되지 않으면 그 제네릭타입을 무공변(invariant)라고 말한다.

Java에서는 모든 제네릭 클래스가 무공변이지만, Kotlin에서는 읽기 전용 컬렉션을 표현하면 공변적(covariant)이 될 수있다.

### 공변성 (Covariance: 하위 타입 관계를 유지)

A가 B의 하위 타입일 때 `Producer<A>`가 `Producer<B>`의 하위타입이면 `Producer`는 공변적이다.
코틀린에서 제네릭 클래스가 타입 파라미터에 대해 공변적임을 표시하려면 타입 파라미터 이름 앞에 `out`을 넣어야 한다.

```kotlin
interface Producer<out T> {
  fun produce(): T
}
```

클래스의 타입 파라미터를 공변적으로 만들면 함수 정의에 사용한 파라미터 타입과 타입인자의 타입이 정확히 일치하지 않더라도
그 클래스의 인스턴스를 함수 인자나 반환 값으로 사용할 수 있다.

모든 클래스를 공변적으로 만들 수는 없다. 공변적으로 만들면 안전하지 못한 클래스도 있다.

타입 안전성을 보장하기 위해 공변적 파라미터는 항상 out 위치에 있어야만 한다.
이는 클래스가 `T` 타입의 값을 생산할 수는 있지만(반환타입) `T` 타입의 값을 소비할 수는 없다(파라미터 타입)는 뜻이다.

`out` 키워드는 `T` 타입의 사용을 제한하며 `T`로 인해 생기는 하위 타입관계의 타입 안전성을 보장한다.

앞에서 살펴본 `List<T>` 타입을 다시 살펴보자. 코틀린 `List`는 읽기 전용이므로
List에 `T` 타입의 값을 추가하거나 기존 값을 변경하는 메소드는 없다. 따라서 `List`는 `T`에 대해 공변적이다.

타입 파라미터를 함수의 파라미터 타입이나 반환 타입에만 쓸수있는 것은 아니다.
타입 파라미터를 다른 타입의 타입 인자로 사용할 수도 있다.

```kotlin
interface List<out T> : Collection<T> {
  fun subList (fromIndex: Int, toIndex: Int): List<T> // T는 out 위치에 있다
}
```

### 반공변성 (contravariance)

반공변성의 메소드는 `T`는 `in` 위치에서만 사용된다

`Consumer<T>`를 예로 들어 설명하자면,
타입 B가 타입 A의 하위타입인경우 `Consumer<A>`가 `Consumer<B>`의 하위타입인 관계가 성립할 때,
제네릭 클래스 `Consumer<T>`는 타입인자 `T`에 대해 반공변이다.

`<? super B>`와 같은 제네릭으로 생각해 볼 수 있다.

`in` 위치에서 반공변성이 적용되는 이유는 `T` 타입을 인자로 받는 함수를 사용되는 경우 문제가 발생할 수 있기 때문이다.
만약 `B extends A` 임과 동시에 `C extends A` 라면 `in`위치에서 사용되는 B와 C는 A의 인터페이스를 구현하기는 하지만,
의도한 올바른 동작을 보장할 수 없다.

### 공변성 / 반공변성

클래스나 인터페이스가 어떤 타입 파라미터에 대해서는 공변적이면서 다른 타입 파라미터에 대해서는 반공변적일 수도 있다.
코틀린의 `Function` 인터페이스가 대표적인 예이다.
다음 선언은 파라미터가 하나뿐인 `Function` 인터페이스인 `Function1`이다.

```kotlin
interface Function1<in P, out R> {
  operator fun invoke(p: P): R
}
```

해당 인터페이스를 사용하는 예제를 보자

`fun enumerateCats(f: Function1<Cat, Number>) { ... }`

코틀린 표기에서 (P) -> R은 `Function1<P, R>`을 알아보기 쉽게 적은 것이다.
이를 코틀린 문법에 따라 다시 쓰면 다음과 같다

```kotlin
fun enumerateCats(f: (Cat) -> Number) { ... }
fun Animal.getIndex(): Int = ...

>>> enumerateCats(Animal::getIndex)
```

이는 함수 타입의 공변성과 반공변성을 살펴보는 문제이기 때문이다.

함수 `Function1`의 타입관계는 첫 번째 타입인자의 (함수 인자타입) 하위 타입관계와는 반대지만 (Cat => Animal)
두 번째 타입인자(함수 반환타입) 하위 타입 관계와는 같다 (Number <= Int)

앞에서 설명한 대로 함수의 입력타입은 반공변적이고 반환타입은 공변적이고, 이를 동시에 적용한 상황이다.

### 생성자 및 private 메소드

컴파일러는 타입 파라미터가 쓰이는 위치를 제한한다.
클래스가 공변적으로 선언된 경우 "Type parameter T is declared as 'out' but occurs in 'in' position" 이라는 오류가 발생한다.
생성자 파라미터는 in이나 out 어느쪽도 아니므로 타입 파라미터가 out이라 해도 여전히 타입을 생성자 파라미터 선언에 사용할 수 있다.

변성은 코드에서 위험 여지가 있는 메소드를 호출할 수 없게 만들어 제네릭 타입의 인스턴스를 잘못 사용하는 일이 없도록 방지한다.
생성자는 인스턴스 생성 뒤 호출할 수 있는 메소드가 아니므로 위험하지 않지만,
`val`이나 `var` 키워드를 생성자 파라미터에 쓴다면 `getter`나 `setter`를 정의하는 것과 같다.

위치 규칙은 외부에서 볼 수 있는 (public, protected, internal) 클래스 API에 적용할 수 있다.
`private` 메소드의 파라미터는 in / out이 아닌 위치이므로, 클래스 내부구현에는 위치가 적용되지 않는다.

### 사용지점 변성

클래스를 선언하면서 변성을 지정하는 것을 declaration site variance 라고 부른다.

자바에서는 타입 파라미터가 있는 타입을 사용할 때마다 해당 타입 파라미터를 하위 타입이나 상위 타입중 어떤 타입으로 대치할 수 있는지 명시해야 한다.
이를 use-site variance라고 부른다.

자바8 표준 라이브러리 `Function` 인터페이스를 살펴보면 자바에서는 use-site variance를 사용함을 알 수 있다.

```java
public interface Stream {
  <R> Stream<R> map(Function<? super T, ? extends R> mapper);
}
```

declaration site variance를 사용하면 매번 변성을 지정하지 않아도 되므로 간결한 코드를 작성할 수 있다.
물론 코틀린에서도 use-site variance를 지원한다.
코틀린의 use-site variance는 자바의 bounded wildcard와 동일한 역할을 수행한다.
`<out T>`는 `<? extends T>`와 같고 `<in T>`는 `<? super T>`와 같다.

```kotlin
fun <T: R, R> copyData(source: MutableList<out T>, dest: MutableList<in R>) {
  for (item in source) {
    dest.add(item)
  }
}
```

### 스타 프로젝션 `*`

`MutableList<*>`는 `MutableList<Any?>`와 같지 않다. (`MutableList<T>`는 `T`에 대해 무공변성이다.)
`MutableList<Any?>`는 모든 타입의 원소를 담는 리스트이지만, `MutableList<*>`는 아직 정해지지 않은 구체적인 타입의 원소만을 담는 리스트이다.

따라서 `MutableList<*>`는 아웃 프로젝션 타입으로 처리된다. (`MutableList<out Any?>` 처럼 처리됨)
`*` 타입이 어떤 타입인지는 모르지만 `Any?`의 하위타입이라는 사실은 분명하므로 안전하게 `Any?` 타입의 원소를 꺼낼 수는 있으나 넣을 수는 없다.

코틀린의 `Type<*>`는 자바의 `Type<?>`에 대응한다.

타입 파라미터를 시그니처에서 전혀 언급하지 않거나,
데이터를 읽기는 하지만 타입에 관심없는 경우처럼 타입 인자정보가 중요하지 않을때 스타 프로젝션을 사용한다.

```kotlin
fun printFirst(list: List<*>) {
  if (list.isNotEmpty()) {
    println(list.first())
  }
}
```

### 스타 프로젝션 예제

사용자 입력검증을 위한 `FieldValidator`라는 인터페이스를 정의했다고 가정하자.
`FieldValidator`에 in 파라미터를 정의해 반공변성을 부여한다.
반공변성이므로 `String` 타입의 필드검증을 위해 `Any`타입을 검증하는 `FieldValidator`를 사용할 수 있다.

```kotlin
interface FieldValidator<in T> {
  fun validate(input: T): Boolean
}

object DefaultStringValidator : FieldValidator<String> {
  override fun validate(input: String) = input.isNotEmpty()
}

object DefaultIntValidator : FieldValidator<Int> {
  override fun validate(input: Int) = input >= 0
}

object Validators {
  private val validators = mutableMapOf<KClass<*>, FieldValidator<*>>()
  fun <T : Any> registerValidator(kClass: KClass<T>, fieldValidator: FieldValidator<T>) {
    validators[kClass] = fieldValidator
  }

  @Suppress("UNCHECKED_CAST")
  operator fun <T : Any> get(kClass: KClass<T>): FieldValidator<T> =
    validators[kClass] as? FieldValidator<T> ?: throw IllegalArgumentException(
      "No validator for ${kClass.simpleName}")
}

>>> Validators.registerValidator(String::class, DefaultStringValidator)
>>> Validators.registerValidator(Int::class, DefaultIntValidator)
>>> println(Validators[String::class].validate("Kotlin"))
true
>> println(Validators[Int::class].validate(42))
true
```

타입 안정성을 보장하지 못하는 부분을 API 내부에 감추고 외부에서 안전하지 못한 부분에 접근하지 못하게 처리한다.
