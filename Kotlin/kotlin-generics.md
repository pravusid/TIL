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

```kt
fun <T> List<T>.slice(indices: IntRange): List<T>
```

함수의 타입 파라미터 `T`가 수신 객체와 반환타입에 쓰인다.
이런 함수를 호출할 때 타입 인자를 명시적으로 지정할 수 있지만, 컴파일러가 추론하므로 불필요 한 경우가 많다.

```kt
>>> val letters = ('a'..'z').toList()
>>> println(letters.slice<Char>(0..2))
[a, b, c]
-- 추론을 사용한다면
>>> println(letters.slice(10..13))
[k, l, m, n]
```

마찬가지로 제네릭 확장 프로퍼티를 선언할 수도 있다.

```kt
val <T> List<T>.penultimate: T
  get() = this[size -2]

>>> println(listOf(1, 2, 3, 4).penultimate)
3
```

하지만 확장 프로퍼티만 제네릭하게 만들 수 있고, 일반 프로퍼티는 불가능하다.

### 제네릭 클래스

클래스에서도 제네릭을 선언할 수 있다.

```kt
class StringList: List<String> {
  override fun get(index: Int): String = ...
}
class ArrayList<T> : List<T> {
  override fun get(index:Int): T = ...
}
```

클래스가 자기 자신을 타입인자로 참조할 수도 있다.

```kt
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

```kt
fun <T : Number> oneHalf(value: T): Double {
  return value.toDouble()
}
```

드물게 타입 파라미터에 둘 이상의 제약을 가해야 하는 경우도 있다. 그럴 땐 약간 다른 구문을 사용한다.

```kt
fun <T> ensureTrailingPeriod(seq: T) where T : CharSequence, T : Appendable {
  if (!seq.endsWith('.')) {
    seq.append('.')
  }
}
```

### 타입 파라미터를 널이 될 수 없는 타입으로 한정

아무런 상한을 정하지 않은 타입 파라미터는 결과적으로 `Any?`를 상한으로 정한 파라미터와 같다

```kt
class Processor<T> {
  fun process(value: T) {
    value?.hashCode() // T는 널이 될 수 있는 타입이므로 안전한 호출을 사용해야 함
  }
}
```

널이 되지 않으면서 제약이 필요없다면 `Any`를 상한으로 사용하면 된다.

```kt
`class Processor<T : Any> {
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

```kt
>>> if (value is List<String>) { ... }
ERROR: Cannot check for instance of erased type
```

즉 다음 두 리스트는 실행시점에 완전히 같은 타입이 된다.

```kt
val list1: List<String> = listOf("a", "b")
val list2: List<Int> = listOf(1, 2)
```

물론 컴파일 단계에서 컴파일러가 타입인자를 인식하고 맞는 값만 넣도록 보장하기 때문에 내부의 값은 선언한 타입과 같다.
