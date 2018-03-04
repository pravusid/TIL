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

## 객체 매핑 : 기본 어노테이션

- `@Entity` : 매핑 대상이 되는 클래스를 알려준다
- `@Table` : Entity와 매핑되는 테이블을 알려준다. 생략시 클래스 이름(Entity 이름)을 테이블 이름으로 매핑한다.
- `@Id` : Entity 클래스의 필드를 테이블의 Primary Key와 매핑한다. @Id가 사용된 필드를 식별자 필드라 한다.
- `@Column` : 필드를 컬럼에 매핑한다. 대소문자를 구분하는 DB에서는 생략하지 말고 명시해야한다. 생략시 필드명이 컬럼명에 매핑되고 camelcase의 경우 snakecase로 변환된다.
- `@Enumerated`: 자바 enum과 매핑

### 기본키 매핑

`@id` 어노테이션 사용

#### 기본키 직접할당

적용 가능 타입 : 기본형(primary type), wrapper형, String, Date, BigDecimal, BigInteger

#### 기본키 자동생성

DMBS마다 지원방식이 다르므로 사용에 유의, (`@GeneratedValue(strategy = {전략명}`)

- AUTO: DBMS에 따라 자동으로 방식을 선택한다 (Oracle:SEQUENCE, MYSQL:IDENTITY ...)
- IDENTITY: 기본키 생성을 DBMS에 위임 (MYSQL AUTO_INCREMENT 해당), 영속상태를 위해서는 식별자가 필요하므로 트랜잭션을 지원하는 쓰기지연에 사용할 수 없음. (데이터베이스에 Entity 저장하여 식별자 획득 후 영속성 컨텍스트에 저장함)
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
  - allocationSize 기본값은 50인데 이는 시퀀스에 접근하는 수를 줄이기 위해서이다. allocationSize 크기만큼 JPA가 메모리에서 식별자를 할당하고 INSERT를 일괄 진행하는 형태이다.
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

## 연관관계 매핑 기초

## 다양한 연관관계 매핑

## 고급 매핑

## 프록시와 연관관계 관리

## 값 타입

## 객체지향 쿼리언어
