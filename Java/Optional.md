# Optional

`java.util.Optional<T>`은 값이 있을 수도 있고, `null`이 될 수도 있는 객체을 감싼 wrapper 클래스이다.

## 선언

처리할 타입을 `Optional<T>`의 제너릭으로 선언한다.
변수명에 `maybe`를 붙여 Optional 타입의 변수라는 것을 더 명확히 나타내기도 한다.

## 생성

Optional 클래스는 세 가지의 static 팩토리 메소드를 제공한다

### Optional.empty()

null을 담고 있는 빈 Optional 객체를 얻는다. `Optional.empty()` 객체는 미리생성된 싱글톤 인스턴스이다.

`Optional<User> maybeUser = Optional.empty();`

### Optional.of(value)

null이 아닌 객체를 담고 있는 Optional 객체를 생성한다. 값이 null인경우 `NPE`가 발생한다.

`Optional<User> maybeUser = Optional.of(user);`

### Optional.ofNullable(value)

null여부가 불확실한 값을 담고있는 Optional 객체를 생성한다.
값이 null인경우 `empty()`와 다르게 `NPE`가 발생하지 않고 비어 있는 Optional 객체(`Optional.empty()`)를 얻는다.

`Optional<User> maybeUser = Optional.ofNullable(user);`

## Optional 내부 객체 접근

아래 메소드들은 모두 Optional이 null이 아닌 경우 해당 값을 반환한다.
Optional이 null인 경우 해당 메소드마다 다르게 작동한다.

### get()

null인 경우 `NoSuchElementException`을 던진다

### orElse(T other)

null인 경우 명시한 인자를 반환한다

### orElseGet(Supplier<? extends T> other)

null인 경우 함수형 인자로 반환값을 정의한다. null인 경우에만 함수가 호출되므로 orElse(T other) 대비 성능상 이점이 있다

### orElseThrow(Supplier<? extends X> exceptionSupplier)

null인 경우 함수형 인자를 통해 예외를 생성하여 던진다

### ifPresent(Consumer<? super T> consumer)

Optional 객체가 감싸고 있는 값이 존재할 경우 그 값을 받아 실행할 내용을 함수형 인자로 정의할 수 있다.

```java
Optional<User> maybeUser = userRepository.findOne(1);
maybeUser.ifPresent(user -> {
    System.out.println(user.getUsername());
});
```

## 함수형 method 지원

Optional에서 함수형으로 데이터를 처리하는 Stream API 메소드인 `map()`, `flatMap()`, `filter()`등을 사용할 수 있다.
