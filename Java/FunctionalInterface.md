# Functional Interface

메소드를 first-class-citizen으로 활용하기 위해
추상 메소드가 하나만 선언된 인터페이스를 함수형 인터페이스로 정의한다.

해당 인터페이스가 함수형 인터페이스임을 명시적으로 나타내기 위해 `@FucntionalInterface` 애노테이션을 붙인다.

`java.util.function` 패키지의 함수형 인터페이스는
크게 `Consumer`, `Supplier`, `Function`, `Operator`, `Predicate로` 구분된다.

## Consumer 함수형 인터페이스

`Consumer` 함수형 인터페이스는 매개 값을 받아 사용만 하고 리턴값이 없는 `accept()` 메소드를 가지고 있다.

매개 변수의 타입과 수에 따라서 아래와 같은 Consumer 인터페이스가 있다

- `Consumer<T>`: `void accept(T t)`
- `BiConsumer<T, U>`: `void accept(T t, U u)`
- `DoubleConsumer`: `void accept(double value)`
- `IntConsumer`: `void accept(int value)`
- `LongConsumer`: `void accept(long value)`
- `ObjDoubleConsumer<T>`: `void accept(T t, double value)`
- `ObjIntConsumer<T>`: `void accept(T t, int value)`
- `ObjLongConsumer<T>`: `void accept(T t, long value)`

## Supplier 함수형 인터페이스

`Supplier` 함수형 인터페이스의는 매개 값이 없고 리턴값이 있는 `getXXX()` 메소드를 가지고 있다.

리턴 타입에 따라서 아래와 같은 Supplier 함수형 인터페이스들이 있다.

- `Supplier<T>`: `T get()`
- `BooleanSupplier`: `boolean getAsBoolean()`
- `DoubleSupplier`: `double getAsDouble()`
- `IntSupplier<T>`: `int getAsInt()`
- `LongSupplier<T>`: `long getAsLong()`

## Function 함수형 인터페이스

`Function` 함수형 인터페이스는 매개 값과 리턴값이 있는 `applyXXX()` 메소드를 가지고 있다.

매개 변수 타입과 리턴 타입에 따라서 아래와 같은 Function 함수형 인터페이스가 있다.

- `Function<T, R>`: `R apply(T t)`
- `BiFunction<T, U, R>`: `R apply(T t, U u)`
- `DoubleFunction<R>`: `R apply(double value)`
- `IntFunction<R>`: `R apply(int value)`
- `IntToDoubleFunction`: `double applyAsDouble(int value)`
- `IntToLongFunction`: `long applyAsLong(int value)`
- `LongToDoubleFunction`: `double applyAsDouble(long value)`
- `LongToIntFunction`: `int applyAsDouble(long value)`
- `ToDoubleFunction<T>`: `double applyAsDouble(T t)`
- `ToDoubleBiFunction<T, U>`: `double applyAsDouble(T t, U u)`
- `ToIntFunction<T>`: `int applyAsInt(T t)`
- `ToIntBiFunction<T, U>`: `int applyAsInt(T t, U u)`
- `ToLongFunction<T>`: `long applyAsLong(T t)`
- `ToLongBiFunction<T, U>`: `long applyAsLong(T t, U u)`

## Operator 함수형 인터페이스

`Operator` 함수형 인터페이스는 매개 값와 리턴값이 있는 `applyXXX()` 메소드를 가지고 있다.
`Function` 함수형 인터페이스와 동일하게 매개 변수를 받아 실행한 후 반환하지만, 타입이 유지된다는 점에서 차이가 있다.

- `UnaryOperator<T>`: `Function<T, R>` 확장
- `BinaryOperator<T>`: `BiFunction<T, U, R>` 확장
- `DoubleUnaryOperator`: `double applyAsDouble(double value)`
- `DoubleBinaryOperator`: `double applyAsDouble(double val1, double val2)`
- `IntUnaryOperator`: `int applyAsInt(int value)`
- `IntBinaryOperator`: `int applyAsInt(int val1, int val2)`
- `LongUnaryOperator`: `long applyAsLong(long value)`
- `LongBinaryOperator`: `long applyAsLong(long val1, long val2)`

## Predicate 함수형 인터페이스

`Predicate` 함수형 인터페이스는 매개 변수와 boolean 리턴값이 있는 `testXXX()` 메소드를 가지고 있다.
이 메소드들은 조건에 맞춰 비교한 결과를 반환한다.

매개 변수 타입과 수에 따라서 아래와 같은 Predicate 함수형 인터페이스들이 있다.

- `Predicate<T>`: `boolean test(T t)`
- `BiPredicate<T, U>`: `boolean test(T t, U u)`
- `DoublePredicate`: `boolean test(double value)`
- `IntPredicate<T>`: `boolean test(int value)`
- `LongPredicate<T>`: `boolean test(long value)`

## andThen() / compose()

함수형 인터페이스의 조건은 하나의 추상메소드만 존재하는 것이다.
따라서 인터페이스에 default 메소드 및 static 메소드가 정의되어 있더라도 함수형 인터페이스의 성질을 유지할 수 있다.

`java.util.function` 패키지의 함수형 인터페이스들도 하나 이상의 default 및 static 메소드를 가지고 있다.

`Comsumer`, `Function`, `Operator` 함수형 인터페이스는 `andThen()`과 `compose()` default 메소드를 가지고 있다.

- `andThen()`: `consumer`, `function`, `operator` 모두지원
- `compose()`: `Function<T, R>`을 상속하는 인터페이스에서 지원

`andThen()`과 `compose()` 메소드는 두 개의 함수형 인터페이스를 연결하는데 사용한다.
`andThen()`과 `compose()`의 차이점은 어떤 함수형 인터페이스부터 먼저 처리하느냐에 따라 다르다.

`andThen()`은 앞의 인터페이스에서 값을 처리하고 이어진 `andThen(interface)`의 매개함수로 값을 넘겨 처리한다.

`compose()`의 경우 함수의 합성과 동일하며 `f(g(x))`와 같은 형태이다.
`compose(interface)`의 매개함수에서 처리한 값을 `compose()`를 호출한 인터페이스로 넘겨 최종 값을 계산한다.

## Consumer의 순차적 연결

`Consumer` 함수형 인터페이스는 처리 결과를 반환하지 않으므로 `andThen()` 메소드는 함수형 인터페이스의 호출 순서만 정하게 된다.

## Function과 Operator의 순차적 연결

Function과 Operator 종류의 함수형 인터페이스는 먼저 실행한 함수형 인터페이스의 결과 값을
다음 함수형 인터페이스의 매개값으로 넘겨주고, 최종 처리 결과를 리턴한다.

## and(), or(), negate() 디폴트 메소드와 isEqual() 정적 메소드

`Predicate` 함수형 인터페이스는 `and()`, `or()`, `negate()`, `isEqual()` 메소드를 가지고 있다.

이 메소드들은 각각 논리 연산자인 `&&`, `||`, `!`, `==`과 대응된다.

`isEqual()` 메소드는 `test()` 매개 값인 `sourceObject`와
`isEqual()`의 매개 값인 `targetObject`를 `java.util.Objects` 클래스의 `eqauls()`의 매개값으로 제공하고,
Objects.equals(source, targetObject)의 리턴값을 얻는 새로운 `Predicate<Object>`를 생성한다.

```java
Predicate<Object> predicate = Predicate.isEqual(targetObject);
boolean result = predicate.test(sourceObject);
```

## minBy(), maxBy() 정적 메소드

`BinaryOperator<T>` 함수형 인터페이스는 `minBy()`와 `maxBy()` 메소드를 제공한다.

이 두 메소드는 매개 값으로 제공되는 `Comparator`를 이용하여 최대와 최소를 얻는 `BinaryOperator<T>`를 리턴한다.

o1과 o2를 비교해서 o1이 작으면 음수를, o1과 o2가 동일하면 0을, o1이 크면 양수를 리턴하는 `compare()` 메소드가 있을 때

```java
@FunctionalInterface
public interface Comparator<T> {
    public int compare(T o1, T o2);
}
```

`Comparator<T>`를 타겟 타입으로 하고 int타입의 크기를 비교하는 람다식을 다음과 같이 작성할 수 있다.

```java
(o1, o2) -> {
    // ...
    return Integer.compare(o1, o2);
}
```

```java
public class OperatorMinByMaxByExam {

    public static void main(String[] args) {
        BinaryOperator<Fruit> binaryOperator;
        Fruit fruit;

        binaryOperator = BinaryOperator.minBy((f1, f2) -> Integer.compare(f1.price, f2.price));
        fruit = binaryOperator.apply(new Fruit("Strawberry", 5000), new Fruit("Graph", 9000));
        System.out.println(fruit.name);

        binaryOperator = BinaryOperator.maxBy((f1, f2) -> Integer.compare(f1.price, f2.price));
        fruit = binaryOperator.apply(new Fruit("Strawberry", 5000), new Fruit("Graph", 9000));
        System.out.println(fruit.name);

    }
}
```

## Method Reference

일반적으로 함수형 메소드 (accept, get, apply...)가 실행할 본문을 람다식으로 구현하여 넘긴다.
단순히 메소드를 호출하는 작업이라면, 직접 메소드를 참조하도록 하는 방법이다.

### 활용 case

1. `(arg) -> Class.staticMethod(arg)` / `Class::staticMethod` / 정적 메소드 참조
2. `(arg, rest) -> arg.instanceMethod(rest)` / `Class::instanceMethod` / 매개변수 참조
3. `(arg) -> instance.instanceMethod(arg)` / `instance::instanceMethod` / 외부 인스턴스 참조 - Lambda Capturing
4. `(arg) -> new Class(arg)` / `Class::new` / 생성자 참조
