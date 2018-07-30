# LocalDateTime

자바의 복잡한 날짜처리를 간편하게 하기 위해 1.8에서 도입된 API

우선 Spring 5 (Spring boot 2) 버전 이상에서는 LocalDate/Time 처리방식이 변화하였다.
이전 버전에서 어떤 방식으로 처리할지를 정리해 둔 것이다.

## Spring Data JPA에서 LocalDateTime 사용

`LocalDate`, `LocalDateTime`을 DB에 입력하려고 하면
`Caused by: com.mysql.jdbc.MsqlDataTruncation: Data truncation: Incorrect dateme value:` 에러가 발생하거나
`tinyblob` 타입으로 저장되는 경우가 발생한다.

### Jsr310JpaConverters (Spring Data JPA 1.8 이상)

`@Convert(converter = Jsr310JpaConverters.LocalDateTimeConverter.class)`

위 애노테이션을 Entity `LocalDate` 필드에 명시하면 변환가능하다.

이 경우 필드는 `Date` 자료형이 아니므로 `@Temporal(TemporalType.TIMESTAMP)` 애노테이션을 붙일 수 없다.
`@Temporal`은 `java.util.Date`이나 `java.util.Calendar`에만 붙일 수 있게 되어 있기 때문이다.

### AttributeConverter를 상속하고 변환 로직을 구현 (Spring Data JPA 1.8 미만, JPA 2.1 이상)

JPA 2.1에 도입된 AttributeConverter 클래스를 상속받는,
자체 Converter를 사용하여 LocalDate/Time을 Date타입으로 자동으로 변환할 수 있다.

```java
@Converter
public class LocalDateTimeConverter implements AttributeConverter<LocalDateTime, Date> {

    @Override
    public Date convertToDatabaseColumn(LocalDateTime localDateTime) {
        return 변환코드
    }

    @Override
    public LocalDateTime convertToEntityAttribute(Date date) {
        return 변환코드
    }
}
```

`org.springframework.data.convert.Jsr310Converters` 클래스는 LocalDate, LocalTime, LocalDateTime, Instant, ZoneId 변환을 제공한다.

Converter를 만들었다면 Converter로 자동 변환할 필드에 `@Convert` 애노테이션을 지정해준다.

```java
@EntityListeners(AuditingEntityListener.class)
@MappedSuperclass
public abstract class BaseEntity implements Serializable {

    @CreatedDate
    @Convert(converter = LocalDateTimePersistenceConverter.class)
    @Column(name = "created_at", updatable = false)
    private LocalDateTime createdAt;

    @LastModifiedDate
    @Convert(converter = LocalDateTimePersistenceConverter.class)
    @Column(name = "modified_at", updatable = true)
    private LocalDateTime modifiedAt;
}
```

### JSR-310 상호 변환 로직을 getter/setter 안에 직접 구현

Jsr310Converters이나 AttributeConverter에 포함된 로직을 직접 구현하는 방식이다.
번거롭지만, 다른곳에서 별도의 Helper Class를 사용하고 있다면 나쁘지 않은 방법일 수 있다.

```java
@EntityListeners(AuditingEntityListener.class)
@MappedSuperclass
public abstract class BaseEntity implements Serializable {

    @CreatedDate
    @Temporal(TemporalType.TIMESTAMP)
    @Column(name = "created_at", updatable = false)
    private Date createdAt;

    @LastModifiedDate
    @Temporal(TemporalType.TIMESTAMP)
    @Column(name = "modified_at", updatable = true)
    private Date modifiedAt;

    // 변환 로직 직접구현

}
```

변환 로직은 [참고문서](/Java/LocalDateTime.md)

## jackson에서 LocalDate/Time 처리

### Spring 4 (ISO 8601)

jackson에서 jsr310 관련 처리를 해주는 모듈을 의존성에 등록한다

```groovy
dependencies { compile group: 'com.fasterxml.jackson.datatype', name: 'jackson-datatype-jsr310', version: '2.9.6' }
```

`application.yml`에 serialize 설정을 추가한다.

```yml
spring.jackson.serialization.WRITE_DATES_AS_TIMESTAMPS: false
```

다음의 결과를 얻을 수 있다

```json
{  
  "localDate":"2018-01-01",
  "localTime":"10:24",
  "localDateTime":"2018-01-01T10:24:00",
  "zonedDateTime":"2018-01-01T10:24:00+09:00"
}
```

### Spring 5 (ISO 8601)

별다른 설정없이 `LocalDateTime` 반환값을 제공한다.

```json
{
  "localDateTime": "2018-01-01T10:24:00.445428",
  "offsetDateTime": "2018-01-01T10:24:00.445428+09:00",
}
```
