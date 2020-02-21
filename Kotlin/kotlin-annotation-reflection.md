# 애노테이션과 리플렉션

## 애노테이션 선언과 적용

### 애노테이션 적용

애노테이션 적용은 `@`과 애노테이션 이름으로 이루어진다.

코틀린의 `@Deprecated` 애노테이션은 `replaceWith` 파라미터를 통해 옛버전을 대신할 수 있는 패턴을 제시할 수 있다.

```kotlin
@Deprecated("Use removeAt(index) instead.", ReplaceWith("removeAt(index)"))
fun remove(index: Int) { ... }
```

애노테이션에 인자를 넘길 때는 일반 함수와 마찬가지로 괄호 안에 인자를 넣는다.
애노테이션의 인자로는 원시 타입의 값, 문자열, enum, 클래스 참조, 다른 애노테이션 클래스, 그리고 앞의 요소들로 이루어진 배열이 들어갈 수 있다.

애노테이션 인자를 지정하는 문법은 자바와 약간 다르다.

- 클래스를 애노테이션 인자로 지정할 때는 `@MyAnnotation(MyClass::class)` 처럼 `::class`를 클래스 이름 뒤에 넣어야 한다.
- 다른 애노테이션을 인자로 지정할 때는 인자로 들어가는 애노테이션 앞에 `@`를 넣지 않아야 한다. 위 예제의 `ReplaceWith`는 애노테이션이다.
- 배열을 인자로 지정하려면 `@RequestMapping(path=arrayOf("/foo", "/bar"))` 처럼 `arrayOf` 함수를 사용한다.
  자바에서 선언한 애노테이션 클래스를 사용한다면 `value`라는 이름의 파라미터가 필요에 따라 자동으로 가변 길이 인자로 변환된다.
  따라서 그런 경우에는 `@JavaAnnotationWithArrayValue("abc", "foo", "bar")` 처럼 `arrayOf` 함수를 쓰지 않아도 된다.

애노테이션 인자를 컴파일 시점에 알 수 있어야 하므로 임의의 프로퍼티를 인자로 지정할 수는 없다.
프로퍼티를 애노테이션 인자로 사용하려면 그 앞에 `const` 변경자를 붙여야한다.

```kotlin
const val TEST_TIMEOUT = 100L
@Test(timeout = TEST_TIMEOUT) fun testMethod() { ... }
```

### 애노테이션 대상

애노테이션을 붙일 때 어떤 요소에 애노테이션을 붙일지 표시할 필요가 있다.
use-site target 선언으로 애노테이션을 붙일 요소를 정할 수 있다.

다음은 `@Rule` 애노테이션을 프로퍼티 getter에 적용하는 예이다 : `@get:Rule`

애노테이션을 사용하는 예를 보자

```kotlin
class HasTempFolder {
  @get:Rule
  val folder = TemporaryFolder()

  @Test
  fun testUsingTempFolder() {
    val createdFile = folder.newFile("myfile.txt")
    val createdFolder = folder.newFolder("subfolder")
    ...
  }
}
```

자바에 선언된 애노테이션을 사용해 프로퍼티에 애노테이션을 붙이는 경우 기본적으로 프로퍼티의 필드에 애노테이션을 붙인다.
하지만 코틀린으로 애노테이션을 선언하면 프로퍼티에 직접 적용할 수 있는 애노테이션을 만들 수 있다.
지원하는 대상 목록은 다음과 같다.

- `property`: 프로퍼티 전체, 자바에 선언된 애노테이션에는 이 사용 지점 대상을 사용할 수 없다
- `field`: 프로퍼티에 의해 생성되는 필드
- `get`: 프로퍼티 게터
- `set`: 프로퍼티 세터
- `receiver`: 확장함수나 프로퍼티 수신 객체 파라미터
- `param`: 생성자 파라미터
- `setparam`: 세터 파라미터
- `delegate`: 위임 프로퍼티의 위임 인스턴스를 담아둔 필드
- `file`: 파일안에 선언된 최상위 함수와 프로퍼티를 담아두는 클래스, `package` 선언 앞에서 파일의 최상위 수준에만 적용가능 하다.

자바와 달리 코틀린에서는 애노테이션 인자로 클래스나 함수 선언이나 타입 외에 임의의 식을 허용한다.
다음은 안전하지 못한 캐스팅 경고를 무시하는 로컬 변수 선언이다

```kotlin
fun test(list: List<*>) {
  @Suppress("UNCHECKED_CAST")
  val strings = list as List<String>
  ...
}
```

### 자바 API를 코틀린 애노테이션으로 제어

다음의 애노테이션을 사용하면 코틀린 선언을 자바에 노출시키는 방법을 변경할 수 있다.

- `@JvmName`은 코틀린 선언이 만들어내는 자바 필드나 메소드 이름을 변경한다
- `@JvmStatic`을 메소드, 객체선언, 동반 객체에 적용하면 그 요소가 자바 정적 메소드로 노출된다
- `@JvmOverloads`를 사용하면 디폴트 파라미터 값이 있는 함수에 대해 컴파일러가 자동으로 오버로딩한 함수를 생성해준다
- `@JvmField`를 프로퍼티에 사용하면 게터나 세터가 없는 `public` 자바 필드로 프로퍼티를 노출시킨다

### 애노테이션을 활용한 JSON 직렬화

애노테이션을 사용하는 고전적인 예제로 객체 직렬화 제어를 들 수 있다.
직렬화는 객체럴 저장장치에 저장하거나 네트워크를 통해 전송하기 위해 텍스트나 이진 형식으로 변환하는 것이다.

JSON 직렬화를 위한 제이키드 라이브러리 예제를 살펴보자
객체를 JSON으로 직렬화할 때 제이키드 라이브러리는 기본적으로 모든 프로퍼티를 직렬화 하며 프로퍼티 이름을 키로 사용한다.
애노테이션을 사용하면 이런 동작을 변경할 수 있다.

- `@JsonExclude` 애노테이션을 사용하면 직렬화나 역직렬화 시 그 프로퍼티를 무시할 수 있다.
- `@JsonName` 애노테이션을 사용하면 프로퍼티를 표현하는 키/값 쌍의 키로 프로퍼티 이름 대신 애노테이션이 지정한 이름을 쓸 수 있다.

### 애노테이션 선언

애노테이션을 선언하려면 `class` 키워드 앞에 `annotation`이라는 변경자를 붙인다.
파라미터가 있는 애노테이션을 정의하려면 애노테이션 클래스 주 생성자에 파라미터를 선언해야 하고 모든 파라미터에 `val`을 붙여야 한다.

자바 애노테이션에는 `vale`라는 특별 메소드가 있고, 코틀린 애노테이션에는 `name`이라는 프로퍼티가 있다.

### 메타애노테이션

애노테이션 클래스에 적용할 수 있는 애노테이션을 메타 애노테이션이라고 부른다.

`@Target` 메타애노테이션은 애노테이션을 적용할 수 있는 요소의 유형을 지정한다.
구체적 `@Target`을 지정하지 않으면 모든 선언에 적용할 수 있다.
애노테이션 대상이 정의된 `enum`은 `AnnotationTarget`이다: `@Target(AnnotationTarget.ANNOTATION_CLASS)`

`@Retention` 메타애노테이션은 애노테이션 클래스를 소스 수준에서만 유지할지 `.class`파일에 저장할지,
실행 시점에 리플렉션을 사용해 접근하게 할지를 지정한다.
자바 컴파일러는 기본적으로 애노테이션을 `.class` 파일에는 저장하지만 런타임에는 사용할 수 없게한다.
하지만 대부분의 애노테이션은 런타임에도 사용할 수 있어야 하므로 코틀린에서는 기본적으로 애노테이션의 Retention을 `RUNTIME`으로 지정한다.

### 애노테이션 파라미터로 클래스 사용

제이키드의 `@DeserializeInterface(MyClass::class)` 처럼 클래스 참조를 인자로 받는 애노테이션을 정의해보자

`annotation class DeserializeInterface(val targetClass: KClass<out Any>)`

`KClass`의 타입 파라미터를 쓸 때 `out` 변경자 없이 `KClass<Any>`라고 쓰면
`DeserializeInterface`에 `MyClass::class`를 인자로 넘길 수 없고 `Any::class`만 넘길 수 있다.
`out` 키워드가 있으면 모든 타입 `T`에 대해 `KClass<T>`가 `KClass<out Any>`의 하위타입이 된다(공변성)

### 애노테이션 파라미터로 제네릭 클래스 받기

제이키드의 `@CustomSerializer` 애노테이션은 커스텀 직렬화 클래스에 대한 참조를 인자로 받는다.
이 직렬화 클래스는 `ValueSerializer` 인터페이스를 구현해야 한다.

```kotlin
interface ValueSerializer<T> {
  fun toJsonValue(value: T): Any?
  fun fromJsonValue(jsonValue: Any?): T
}
```

날짜를 직렬화 한다고 가정하자. 이때 `ValueSerializer<Date>`를 구현하는 `DateSerializer`를 `Person` 클래스에 적용한다면

```kotlin
data class Person(
  val name: String,
  @CustomSerializer(DateSerializer::class) val birthDate: Date
)
```

`@CustomSerializer` 애노테이션을 구현하는 방법을 살펴보자
`ValueSerializer` 타입을 참조하려면 항상 타입인자를 제공해야 한다.
하지만 이 애노테이션이 어떤 타입에 쓰일지 알 수 없으므로 스타 프로젝션을 사용할 수 있다.

```kotlin
annotation class CustomSerializer(
  val serializerClass: KClass<out ValueSerializer<*>>
)
```

## 리플렉션

리플렉션은 실행 시점에 동적으로 객체의 프로퍼티와 메소드에 접근할 수 있게 해주는 방법이다.
타입과 관계없이 객체를 다뤄야 하거나 객체가 제공하는 메소드나 프로퍼티 이름을 오직 실행시점에만 알 수 있는 경우가 있다.

코틀린 리플렉션을 사용하려면 두 가지 리플렉션 API를 다뤄야 한다.
하나는 `java.lang.reflect`이고 다른 하나는 `kotlin.reflect` API이다.

코틀린 API에서는 자바에는 없는 프로퍼티나 널이 될 수 있는 타입과 같은 코틀린 고유개념에 대한 리플렉션을 제공한다.

### 코틀린 리플렉션 API

`java.lang.Class`에 해당하는 `KClass`를 사용하면
클래스 안에 있는 모든 선언을 열거 하고 각 선언에 접근하거나 클래스의 상위 클래스를 얻는등의 작업이 가능하다.
`MyClass::class`라는 식을 쓰면 `KClass`의 인스턴스를 얻을 수 있다.
실행시점에 객체의 클래스를 얻으려면 먼저 객체의 `javaClass` 프로퍼티를 사용해 객체의 자바 클래스를 얻어야한다.
`javaClass`는 자바의 `java.lang.Object.getClass()`와 같다.
일단 자바 클래스를 얻었으면 `.kotlin` 확장 프로퍼티를 통해 자바에서 코틀린 리플렉션 API로 옮겨올 수 있다.

`KClass` 에서는 클래스 내부를 살펴볼 때 사용할 수 있는 다양한 메소드가 있다.

```kotlin
interface KClass<T : Any> {
  val simpleName: String?
  val qualifiedName: String?
  val members: Collection<KCallable<*>>
  val constructors: Collection<KFunction<T>>
  val nestedClasses: Collection<KClass<*>>
  ...
}
```

클래스의 모든 멤버 목록은 `KCallable` 인스턴스의 컬렉션이다.
`KCallable`은 함수와 프로퍼티를 아우르는 공통 상위 인터페이스이고 그 안에는 `call` 메소드가 들어있다.
`call` 메소드를 사용하면 함수나 프로퍼티의 getter를 호출할 수 있다.

```kotlin
interface KCallable<out R> {
  fun call(vararg args: Any?): R
  ...
}
```

`call`을 사용할 때는 함수 인자를 `vararg` 리스트로 전달한다.
다음 코드는 `call`을 사용해 함수를 호출하는 예제이다.

```kotlin
fun foo(x: Int) = println(x)
>>> val kFunction = ::foo
>>> kFunction.call(42)
42
```

`::foo` 식의 값은 `KFunction` 클래스의 인스턴스이다. 이 함수 참조가 가리키는 함수를 호출하려면 `KCallable.call` 메소드를 호출한다.
`call`에 넘긴 인자 개수와 원래 함수에 정의된 파라미터 개수가 맞아야 한다.

함수를 호출하기 위해 더 구체적인 메소드를 사용할 수 있다.
`::foo`의 타입 `KFunction1<Int, Unit>`에는 파라미터와 반환 값 타입 정보가 들어있다.
1은 함수의 파라미터가 1개라는 의미이다. `KFunction1` 인터페이스를 통해 함수를 호출하려면 `invoke` 메소드를 사용해야한다.
`invoke`는 정해진 개수의 인자만을 받아들이며, 인자 타입은 `KFunction1` 제네릭 인터페이스의 첫 번째 타입 파라미터와 같다.
또한 `kfunction`을 직접호출할 수도 있다.

```kotlin
fun sum(x: Int, y: Int) = x + y
>>> val kFunction: KFunction2<Int, Int, Int> = ::sum
>>> println(kFunction.invoke(1, 2) + kFunction(3, 4))
10
```

`KFunction`의 인자 타입과 반환 타입을 모두 다 안다면 `invoke` 메소드를 호출하는게 낫다.
`call` 메소드는 타입 안전성을 보장해주지 않기 때문이다.

> `KFunctionN` 타입은 `KFunction`을 확장하며 `N`과 파라미터 개수가 같은 `invoke`를 추가로 포함한다.
> 이런 함수 타입들은 컴파일러가 생성한 합성타입이다. 따라서 `kotlin.reflect` 패키지에서 이런 타입의 정의를 찾을 수 는 없다.

`KProperty`의 `call` 메소드를 호출할 수도 있다. `KProperty`의 `call`은 프로퍼티의 getter를 호출한다.
최상위 프로퍼티는 `KProperty0` 인터페이스의 인스턴스로 표현되며, `KProperty0` 안에는 인자가 없는 `get` 메소드가 있다.

> 최상위 수준이나 클래스 안에 정의된 프로퍼티만 리플렉션으로 접근할 수 있고 함수의 로컬 변수에는 접근할 수 없다.

### 리플렉션을 사용한 객체 직렬화 구현

우선 제이키드의 직렬화 함수 선언을 살펴보자

`fun serialize(obj: Any): String`

이 함수는 객체를 받아서 그 객체에 대한 JSON 표현을 문자열로 돌려준다.
값을 직렬화 하면서 `StringBuilder` 객체뒤 직렬화한 문자열을 추가한다.
`append` 호출을 더 간결하게 수행하기 위해 직렬화 기능을 `StringBuilder`의 확장함수로 구현한다.

`serialize`는 대부분의 작업을 `serializeObject`에 위임한다.

`fun serialize(obj: Any): String = builderString { serializeObject(obj) }`

이제 `serializeObject` 구현을 살펴보자

```kotlin
private fun StringBuilder.serializeObject(obj: Any) {
  val kClass = obj.javaClass.kotlin
  val properties = kClass.memberProperties

  properties.joinToStringBuilder(this, prefix = "{", postfix = "}") { prop ->
    serializeString(prop.name)
    append(": ")
    serializePropertyValue(prop.get(obj))
  }
}
```

이 예제에서는 어떤 객체의 클래스에 정의된 모든 프로퍼티를 열거하기 때문에 각 프로퍼티가 어떤 타입인지 알 수 없다.
따라서 `prop` 변수의 타입은 `KProperty1<Any, *>`이며, `prop.get(obj)` 메소드 호출은 `Any` 타입의 값을 반환한다.
하지만 이 코드에서는 어떤 프로퍼티의 `get`에 넘기는 객체가 바로 그 프로퍼티를 얻어온 객체인 `obj`이므로 항상 프로퍼티 값이 제대로 반환된다.

### 애노테이션을 활용한 직렬화 제어

애노테이션을 `serializeObject` 함수가 어떻게 처리하는지 살펴보자.

어떻게 `@JsonExclude`가 붙은 프로퍼티를 제외할 수 있을지 살펴보자

`KAnnotatedElement` 인터페이스에는 `annotations` 프로퍼티가 있다.
이 프로퍼티는 소스코드상에서 해당요소에 적용된 (`@Retention`이 `RUNTIME`인) 모든 애노테이션 인스턴스의 컬렉션이다.
`KProperty`는 `KAnnotatedElement`를 확장하므로 `property.annotations`를 통해 프로퍼티의 모든 애노테이션을 얻을 수 있다.

```kotlin
inline fun <reified T> KAnnotatedElement.findAnnotation(): T? {
  annotations.filterIsInstance<T>().firstOrNull()
}
```

`findAnnotation` 함수는 인자로 전달받은 타입의 애노테이션이 있으면 반환하고, 타입 파라미터를 `reified`로 만들어 클래스를 타입인자로 전달한다.

```kotlin
val properties = kClass.memberProperties.filter { it.findAnnotation<JsonExclude>() == null }
```

`@JsonName`의 경우 애노테이션 존재여부와 함께 애노테이션 인자도 알아야 한다.

```kotlin
val jsonNameAnn = prop.findAnnotation<JsonName>()
val propName = jsonNameAnn?.name ?: prop.name
```

`Person` 클래스 인스턴스를 직렬화 하는 과정을 살펴보자.
`firstName` 프로퍼티를 직렬화 하는 동안 `jsonNameAnn`에는 `JsonName` 애노테이션 클래스에 해당하는 인스턴스가 들어있다.

프로퍼티 필터링을 포함하는 객체 직렬화

```kotlin
private fun StringBuilder.serializeObject(obj: Any) {
  obj.javaClass.kotlin.memberProperties
      .filter { it.findAnnotation<JsonExclude>() == null}
      .joinToStringBuilder(this, prefix = "{", postfix = "}") {
        serializeProperty(it, obj)
      }
}

private fun StringBuilder.serializeProperty(prop: KProperty1<Any, *>, obj: Any) {
  val jsonNameAnn = prop.findAnnotation<JsonName>()
  val propName = jsonNameAnn?.name ?: prop.name
  serializeString(propName)
  append(": ")
  serializePropertyValue(prop.get(obj))
}
```

마지막으로 `@CustomSerializer` 애노테이션을 구현해보자

```kotlin
annotaion class CustomSerializer {
  val serializerClass: KClass<out ValueSerializer<*>>
}
```

이를 사용하기 위한 `getSerializer`를 구현해보자

```kotlin
fun KProperty<*>.getSerializer(): ValueSerializer<Any?>? {
  val customSerializerAnn = findAnnotation<CustomSerializer>() ?: return null
  val serializerClass = customSerializerAnn.serializerClass
  val valueSerializer = serializerClass.objectInstance ?: serializerClass.createInstance()

  @Suppress("UNCHECKED_CAST")
  return valueSerializer as ValueSerializer<Any?>
}
```

`getSerializer`는 `KProperty`의 확장함수 이다.

`@CustomSerializer` 애노테이션으로 처리되는 클래스와 객체는 모두 `KClass`로 표현된다.
`KClass`에서 `objectInstance` 호출로 싱글톤 인스턴스를 생성한다.

### JSON 파싱과 객체 역직렬화

역직렬화 API는 직렬화와 마찬가지로 함수 하나로 이뤄져 있다.

`inline fun <reified T: Any> deserialize(json: String): T`

JSON 역직렬화기는 3단계로 구현되어 있다.

1. 어휘 분석기 : lexical analyzer
2. 문법분석기 : syntax analyzer
3. 파서 : parser

최상위 역직렬화 함수를 정의해보자

```kotlin
fun <T : Any> deserialize(json: Reader, targetClass: KClass<T>): T {
  val seed = ObjectSeed(targetClass, CallsInfoCache())
  Parser(json, seed).parse()
  return seed.spawn()
}
```

객체 역직렬화 하기

```kotlin
class ObjectSeed<out T: Any>(targetClass: KClass<T>, val classInfoCache: CallsInfoCache): Seed {
  private val classInfo: ClassInfo<T> = classInfoCache[targetClass]
  private val valueArguments = mutableMapOf<KParameter, Any?>()
  private val seedArguments = mutableMapOf<KParameter, seed>()
  private val arguments: Map<KParameter, Any?>
    get() = valueArguments + seedArguments.mapValues { it.value.spawn() }

  override fun setSimpleProperty(propertyName: String, value: Any?) {
    val param = classInfo.getConstructorParameter(propertyName)
    valueArguments[param] = classInfo.deserializeConstructorArgument(param, value)
  }

  override fun createCompositeProperty(propertyName: String, isList: Boolean): Seed {
    val param = classInfo.getConstructorParameter(propertyName)
    val deserializeAs = classInfo.getDeserializeClass(propertyName)
    val seed = createSeedForType(deserializeAs ?: param.type.javaType, isList)
    return seed.apply { seedArguments[param] = this }
  }

  override fun sapwn(): T = classInfo.createInstance(arguments)
}
```

### 최종 역직렬화 단계

`KCallable.call` 은 인자 리스트를 받아서 함수나 생성자를 호출해준다.
하지만 디폴트 파라미터 값을 지원하지 않는다는 한계가 있다.
디폴트 파라미터를 사용하기 위해서는 `KCallable.callBy`를 사용할 수 있다.

```kotlin
interface KCallable<out R> {
  fun callBy(args: Map<KParameter, Any?>): R
  ...
}
```

`callBy` 메소드는 파라미터와 파라미터에 해당하는 값을 연결해주는 맵을 인자로 받는다.
인자로 받은 맵에서 파라미터를 찾을수 없을 때 디폴트 값이 정의되어 있다면 디폴트 값을 사용한다.
여기서는 `args` 맵에 들어있는 각 값의 타입이 생성자의 파라미터 타입과 일치해야 한다. 그렇지 않으면 `Exception`이 발생한다.
따라서 파라미터 타입이 어떤것인지를 확인하기 위해 `KParameter.type` 프로퍼티를 활용해야 한다.

```kotlin
fun serializerForType(type: Type): ValueSerializer<out Any?>? =
  when(type) {
    Byte::class.java -> ByteSerializer
    Int::class.java -> IntSerializer
    Boolean::class.java -> BooleanSerializer
    ...
    else -> null
  }
```

타입별 `ValueSerializer` 구현은 필요한 타입 검사나 변환을 수행한다

```kotlin
object BooleanSerializer : ValueSerializer<Boolean> {
  override fun fromJsonValue(jsonValue: Any?): Boolean {
    if (jsonValue !is Boolean) throw JKidException("Boolean expected")
    return jsonValue
  }

  override fun toJsonValue(value: Boolean) = value
}
```

`ClassInfoCache`는 리플렉션 연산의 비용을 줄이기 위한 클래스이다.
직렬화와 역직렬화에 사용되는 애노테이션들은 파라미터가 아니라 프로퍼티에 적용된다.
하지만 객체를 역직렬화할 때는 프로퍼티가 아니라 생성자 파라미터를 다뤄야 한다.

애노테이션을 꺼내려면 파라미터에 해당하는 프로퍼티를 찾아야 하는데 JSON에서 모든 키/값 쌍을 읽을 때마다 검색을 수행하면 느려진다.
따라서 클래스별로 한 번만 검색을 수행하고 검색결과를 캐시에 넣는 기능을 수행한다.

```kotlin
class ClassInfoCache {
  private val cacheData = mutableMapOf<KClass<*>, ClassInfo<*>>()
  @Suppress("UNCHECKED_CAST")
  operator fun <T : Any> get(cls: KClass<T>): ClassInfo<T> =
    cacheData.getOrPut(cls) { ClassInfo(cls) } as ClassInfo<T>
}
```

맵에 값을 저장할 때는 타입정보가 사라지지만,
맵에서 받은 값의 타입인 `ClassInfo<T>` 타입인자가 항상 올바른 값이 되게 `get` 메소드 값이 구현이 보장한다.

`ClassInfo` 클래스는 대상 클래스의 새 인스턴스를 만들고 필요한 정보를 캐시해 둔다.

```kotlin
class ClassInfo<T : Any>(cls: KClass<T>) {
  private val constructor = cls.primaryConstructor!!
  private val jsonNameToParam = hashMapOf<String, KParameter>()
  private val paramToSerializer = hashMapOf<KParameter, ValueSerializer<out Any?>>()
  private val jsonNameToDeserializeClass = hashMapOf<String, Class<out Any?>>()

  init {
    constructor.parameters.forEach { cacheDataForParameter(cls, It) }
  }

  fun getConstructorParameter(propertyName: String): KParameter =
    jsonNameToParam[propertyName]!!

  fun deserializeConstructorArgument(param: KParameter, value: Any?): Any {
    val serializer = paramToSerializer[param]
    if (serializer != null) return serializer.fromJsonValue(value)
    validateArgumentType(param, value)
    return value
  }

  fun createInstance(arguments: Map<KParameter, Any?>): T {
    ensureAllParametersPresent(arguments)
    return constructor.callBy(arguments)
  }
  ...
}
```

`jsonNameToParam`은 JSON 파일의 각 키에 해당하는 파라미터를 저장하고, `paramToSerializer`는 각 파라미터에 대한 직렬화기를 저장한다.
`jsonNameToDeserializeClass`는 `@DeserializeInterface` 애노테이션 인자로 지정한 클래스를 저장한다.
