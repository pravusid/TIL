# Java Persistence API

JPA(Java Persistence API)는 자바진영의 ORM(Object Relational Mapping) 기술표준이다.

객체지향 언어와 관계형 데이터베이스는 목적이 다르기 때문에 둘을 함께 사용하는 상황에서는 패러다임 불일치가 발생할 수 밖에 없다.
JPA는 패러다임 불일치 문제를 간편하게 해소해 주고 객체지향 모델링을 유지하도록 돕는다.

과거 엔터프라이즈 자바 빈즈의 복잡성 문제로 부터 Hibernate가 탄생했고 하이버네이트를 토대로 자바 표준 ORM인 JPA가 탄생했다.

JPA는 자바 애플리케이션과 JDBC사이에서 동작한다.

## JPA 시작

JPA와 하이버네이트를 사용하기 위해서는 다음의 라이브러리를 사용한다.

- `hibernate-core` (`hibernate-jpa-2.1-api` 포함)

### 기본설정 : `persistence.xml`

<https://docs.jboss.org/hibernate/orm/5.2/userguide/html_single/Hibernate_User_Guide.html>

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<persistence xmlns="http://xmlns.jcp.org/xml/ns/persistence"
             xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
             xsi:schemaLocation="http://xmlns.jcp.org/xml/ns/persistence
             http://xmlns.jcp.org/xml/ns/persistence/persistence_2_1.xsd"
             version="2.1">
  <!-- 영속성유닛 등록 / transaction-type 설정 -->
  <persistence-unit name="unitname">
    <provider>org.hibernate.jpa.HibernatePersistenceProvider</provider>
    <!-- @Entity 등록 -->
    <class>kr.pravusid.domain.User</class>
    <class>kr.pravusid.domain.Board</class>
    <properties>
      <!-- 필수속성 -->
      <property name="javax.persistence.jdbc.driver" value="org.h2.Driver"/>
      <property name="javax.persistence.jdbc.url" value="jdbc:h2:mem:mydb"/>
      <property name="javax.persistence.jdbc.user" value="sa"/>
      <property name="javax.persistence.jdbc.password" value=""/>
      <property name="hibernate.dialect" value="org.hibernate.dialect.H2Dialect" />
      <!-- 스키마 자동생성 -->
      <property name="hibernate.hbm2ddl.auto" value="create-drop"/>
      <!-- 기본키 생성 전략: 기본은 false -->
      <property name="hibernate.id.new_generator_mappings" value="true"/>
      <!-- 선택사항 -->
      <property name="hibernate.show_sql" value="true" />
      <property name="hibernate.format_sql" value="true" />
      <property name="hibernate.use_sql_comments" value="true" />
    </properties>
  </persistence-unit>
</persistence>
```

설정파일은 `classpath:META-INF/persistence.xml`에 위치하면 자동으로 인식된다.

`persistence-unit`에서는 `transaction-type="JTA||RESOURCE_LOCAL"`을 명시할 수 있다.

`javax.persistence.jdbc`에서는 데이터베이스 접속 설정을 한다.

`hibernate.dialect`에서는 특정 DBMS에 종속되지 않고 다양한 DBMS활용을 위한 Dialect(방언)을 명시한다.

기타 `hibernate` 설정에서는 개발 편의나 디버깅을 위한 옵션을 제공한다.

#### hibernate.hbm2ddl.auto 속성

- `create`: 기존 테이블 삭제 후 생성
- `create-drop`: create 속성에 종료시 drop
- `update`: 변경사항만 수정한다
- `validate`: 매핑정보를 비교해 차이가 있으면 경고 출력 후 application을 실행하지 않는다.

## 영속성 관리

### EntityManagerFactory

- EntityManagerFactory는 보통 하나를 생성하여 공유한다.
- 생성될 때 커넥션풀도 생성한다.
- Thread safe 하다.

### EntityManager

- EntityManager는 EntityManagerFactory로 부터 생성한다.
- EntityManager는 연결이 실제로 필요한 시점까지 커넥션을 획득하지 않는다.
- EntityManager는 Thread safe하지 않다.

### 영속성 컨텍스트 (Persistence context)

Java 내에서 데이터베이스 계층(persistence layer)과 동기화하는 Entity 저장환경이라고 볼 수 있다.

#### 영속성 컨텍스트와 관련한 Entity의 생명주기

- 비영속(transient): 영속성 컨텍스트와 전혀 관계가 없는 상태
- 영속(managed): 영속성 컨텍스트에 저장된 상태
  - 영속화 : `em.persist({Object});`
  - `em.find()`나 JPQL을 사용해서 조회한 엔티티도 영속성 컨텍스트에 의해 관리된다.
- 준영속/분리됨(detached): 영속성 컨텍스트에 저장되었다 분리된 상태
  - 특정 Entity를 준영속 상태로: `em.detach({Object})`
  - 영속성 컨텍스트를 완전히 초기화(영속성 컨텍스트의 모든 엔티티를 준영속 상태로): `em.clear()`
  - 영속성 컨텍스트를 종료한다(영속성 컨텍스트의 모든 엔티티를 준영속 상태로): `em.close()`
- 삭제(removed): 영속성 컨텍스트에서 삭제된 상태
  - 제거: `em.remove({Object})`

#### 영속성 컨텍스트의 특징

- 영속성 컨텍스트는 엔티티를 식별자 값(@Id)으로 구분한다. 따라서 영속상태는 반드시 식별자 값이 있어야 한다.
- 영속성 컨텍스트는 `flush()`를 하는 순간 데이터베이스에 저장된다. (영속성 컨텍스트와 데이터베이스를 동기화)
  - Flush가 호출 되는 경우
    - `em.flush()`를 호출시 플러시: 사용빈도 낮음
    - 트랜잭션 커밋시 플러시 자동호출
    - JPQL 쿼리 실행시 플러시 자동호출
  - 플러시 옵션(`em.setFlushMode({Option}`))
    - `FlushModeType.AUTO`: 기본값
    - `FlushModeType.COMMIT`: 커밋할 때만 플러시
- 영속성 컨텍스트의 장점
  - 1차 캐시: 영속상태의 엔티티는 이곳에 저장된다. 1차 캐시의 키는 식별자 값이다. 엔티티 조회를 시도하면 우선 1차 캐시에서 찾고 없다면 데이터베이스를 조회한다.
  - 동일성 보장: 동일한 식별자로 조회한 엔티티 인스턴스는 조회 시점과 관계없이 동일성(identity)을 보장한다. - (영속성 컨텍스트에 포함된 상태)
  - 트랜잭션을 지원하는 쓰기 지연: 트랜잭션 상황에서 1차 캐시의 Entity 수정에 대응되는 쿼리를 모아두고 트랜잭션 커밋시 한번에 처리한다.
  - 변경 감지: 엔티티의 변경사항을 데이터베이스에 자동으로 반영하는 기능. 엔티티를 영속성 컨텍스트에 보관할 때 최초상태가 스냅샷으로 저장된다. 플러시 시점에 스냅샷과 엔티티를 비교하여 변경점을 찾는다.
    - 변경 감지 과정
      - 트랜잭션을 커밋하면 flush가 호출된다
      - 현재 Entity들과 snapshot을 비교하여 변경된 Entity를 찾는다.
      - 변경된 Entity가 있으면 수정쿼리를 생성하여 쓰기 지연 저장소에 보낸다.
      - 쓰기지연 저장소의 SQL을 데이터베이스로 보낸다.
      - 데이터베이스 트랜잭션을 커밋한다.
    - 변경 감지 이점
      - 수정 쿼리가 항상 같다 (바인딩 데이터만 다르다)
      - 따라서 데이터베이스는 이전에 파싱된 쿼리를 재사용할 수 있다.
      - 전체 필드를 사용해서 수정쿼리를 같게 만들지 않고 변경점만 반영하려면 `@DynamicUpdate` 어노테이션을 사용한다.
  - 지연 로딩: 실제 객체 대신 프록시 객체를 불러두고 해당객체가 실제 사용될 때 영속성 컨텍스트에 해당 데이터를 불러오는 방법이다.

#### 준영속 상태의 특징

- 영속성 컨텍스트가 관리하지 않으므로, 영속성 상태의 기능인 1차캐시, 쓰기지연, 변경감지, 지연로딩 등이 동작하지 않는다.
- 준영속 상태는 영속상태에서 분리되었기 때문에 식별자 값을 가지고 있다.
- 지연로딩을 사용할 수 없다.
- `merge()`를 호출하여 준영속 인스턴스를 다시 영속성 컨텍스트에 병합할 수 있다. `merge()`는 비영속을 영속 컨텍스트에 병합하는 기능도 있다.

## 객체 매핑

- `@Entity` : 매핑 대상이 되는 클래스를 알려준다
- `@Table` : Entity와 매핑되는 테이블을 알려준다. 생략시 클래스 이름(Entity 이름)을 테이블 이름으로 매핑한다.

### 기본키 매핑

`@Id` : Entity 클래스의 필드를 테이블의 Primary Key와 매핑한다. @Id가 사용된 필드를 식별자 필드라 한다.

#### 기본키 직접할당

적용 가능 타입 : 기본형(primary type), wrapper형, String, Date, BigDecimal, BigInteger

#### 기본키 자동생성

DMBS마다 지원방식이 다르므로 사용에 유의, (`@GeneratedValue(strategy = {전략명}`)

- AUTO: DBMS에 따라 자동으로 방식을 선택한다 (Oracle:SEQUENCE, MYSQL:IDENTITY ...)
- IDENTITY: 기본키 생성을 DBMS에 위임 (MYSQL AUTO_INCREMENT 해당),
  영속상태를 위해서는 식별자가 필요하므로 트랜잭션을 지원하는 쓰기지연에 사용할 수 없음.
  (데이터베이스에 Entity 저장하여 식별자 획득 후 영속성 컨텍스트에 저장함)
  ```java
  @Entity
  public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
  }
  ```
- SEQUENCE: 데이터베이스 시퀀스를 사용하여 기본키 할당 (오라클...)
  ```java
  @Entity
  public class User {
    @Id
    @SequenceGenerator(name = "{NAME}", sequenceName = "{데이터베이스의 시퀀스}", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "{NAME}")
    private Long id;
  }
  ```
  - allocationSize 기본값은 50인데 이는 시퀀스에 접근하는 수를 줄이기 위해서이다.
    allocationSize 크기만큼 JPA가 메모리에서 식별자를 할당하고 INSERT를 일괄 진행하는 형태이다.
- TABLE: 키 생성 테이블 사용
  ```java
  @Entity
  @TableGenerator(name = "{NAME}", table = "{TABLE_NAME}", pkColumnValue = "USER_SEQ", allocationSize = 1)
  public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY, generator = "{NAME}")
    private Long id;
  }
  ```
  - Table 전략은 키 조회, 키 증가시 각각 DB에 접근하는 단점이 있다 (2회)

### 필드와 컬럼 매핑

#### `@Column`

필드를 컬럼에 매핑한다. 대소문자를 구분하는 DB에서는 생략하지 말고 명시해야한다.
생략시 필드명이 컬럼명에 매핑되고 camelcase의 경우 snakecase로 변환된다

속성 정리
| 속성 | 기능 | 기본값 |
| --- | --- | --- |
| name | 필드와 매핑할 테이블의 컬럼이름 | 객체의 필드 이름 |
| table | 하나의 엔티티를 두 개 이상의 테이블에 매핑 | 현재 매핑 테이블 |
| precision, scale | DDL 생성시 소수 포함 전체 자릿수, 소수의 자릿수 지정. BigDecimal 에서 사용 | |
| nullable | DDL 생성시 not null 제약조건 | true |
| unique | DDL 생성시, 한 컬럼에 Unique 제약조건시, 두 개 이상은 @Table에서 지정 | |
| length | 문자 길이 제약조건, String에 사용 | 255 |
| columnDefinition | DDL 생성시 옵션 직접 기입 | |

#### `@Enumerated`

자바 enum과 매핑할 때 사용.
기본값은 enum의 순서가 저장되는 ORDINAL이나 값 변경 추적이 불가하므로 STRING 사용이 권장됨

- `EnumType.ORDINAL`: enum 순서를 데이터베이스에 저장 (기본 값)
- `EnumType.STRING`: enum 이름을 데이터베이스에 저장

#### `@Temporal`

날짜 타입(java.util.Date, java.util.Calendar) 매핑시 사용
`@Temporal` 생략 시 timestamp(datetime)과 매핑 됨

- `TemporalType.DATE`: (날짜) 데이터베이스 date 타입과 매핑
- `TemporalType.TIME`: (시각) 데이터베이스 time 타입과 매핑
- `TemporalType.TIMESTAMP`: (날짜와 시각) 데이터베이스 timestamp 타입과 매핑

#### `@Lob`

BLOB, CLOB 타입 매핑

- CLOB과 매핑되는 Java Type: String, Char[], java.sql.CLOB
- BLOB과 매핑되는 Java Type: byte[], java.sql.BLOB

#### `@Transient`

해당 필드를 데이터베이스와 매핑하지 않는다 (저장, 조회 둘다 하지 않는다)

#### `@Access`

JPA가 Entity에 접근하는 방식을 지정한다

- `AccessType.FIELD`: 필드에 직접 접근 (private여도 접근함)
- `AccessType.PROPERTY`: getter로 접근

`@Access`를 명시 하지 않으면 `@Id`의 위치에 따라 달라진다.
(변수에 위치: FIELD, getter에 위치: PROPERTY)

## 연관관계 매핑 기초

## 다양한 연관관계 매핑

## 고급 매핑

## 프록시와 연관관계 관리

## 값 타입

## 객체지향 쿼리언어
