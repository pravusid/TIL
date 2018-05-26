# Spring Data JPA

## 시작

<https://docs.spring.io/spring-data/jpa/docs/1.7.0.RELEASE/reference/html/>

## 초기설정

- build.gradle에 추가

  ```groovy
  compile('org.springframework.boot:spring-boot-starter-data-jpa')
  runtime('mysql:mysql-connector-java')
  compileOnly('org.projectlombok:lombok')
  ```

- application.properties
  > The data source properties starting with spring.datasource.* will automatically be read by spring boot JPA. To change the hibernate properties we will use prefix spring.jpa.properties.* with hibernate property name. On the basis of given data source URL, spring boot can automatically identify data source driver class. So we need not to configure diver class.

  ```text
  # JDBC 사용할 때
  spring.datasource.driver-class-name=com.mysql.jdbc.Driver

  # DB서버 정보
  spring.datasource.url=jdbc:mysql://localhost/test?autoReconnect=true&useUnicode=true&characterEncoding=utf8
  spring.datasource.username=dbuser
  spring.datasource.password=dbpass

  spring.datasource.tomcat.max-wait=20000
  spring.datasource.tomcat.max-active=50
  spring.datasource.tomcat.max-idle=20
  spring.datasource.tomcat.min-idle=15

  # JPA로 사용할 데이터베이스 명시
  spring.jpa.database=mysql

  # 개발시 테스트용
  spring.jpa.hibernate.ddl-auto=create-drop
  spring.jpa.show-sql=true

  # JPA 설정들
  spring.data.jpa.repositories.enabled: It enables JPA repositories. The default value is true.
  spring.jpa.database: It targets database to operate on. By default embedded database is auto-detected.
  spring.jpa.database-platform: It is used to provide the name of database to operate on. By default it is auto- detected.
  spring.jpa.generate-ddl: It is used to initialize schema on startup. By default the value is false.
  spring.jpa.hibernate.ddl-auto: It is DDL mode used for embedded database. Default value is create-drop.
  spring.jpa.hibernate.naming.implicit-strategy: It is Hibernate 5 implicit naming strategy fully qualified name.
  spring.jpa.hibernate.naming.physical-strategy: It is Hibernate 5 physical naming strategy fully qualified name.
  spring.jpa.hibernate.use-new-id-generator-mappings: It is used for Hibernate IdentifierGenerator for AUTO, TABLE and SEQUENCE.
  spring.jpa.open-in-view: The default value is true. It binds a JPA EntityManager to the thread for the entire processing of the request.
  spring.jpa.properties.*: It sets additional native properties to set on the JPA provider.
  spring.jpa.show-sql: It enables logging of SQL statements. Default value is false.
  ```

## 데이터 입출력

### Annotation

- `@Entity` : Entity임을 표시함. 테이블과 매칭해서 사용한다.
- `@Table(name="articles")` : 기존 테이블과 연결
- `@Id` : PK 정의
- `@GeneratedValue` : Auto_increment, Sequence를 생성해준다
  - `strategy = GenerationType.AUTO`
- `@Column` : 컬럼과 연결, 컬럼 정보 기입
  - `name` : 기존 column 이름
- `@OneToMany(mappedBy="@ManyToOne variable name")`
- `@ManyToOne` : JoinColumn으로 pk를 가져와서 활용
- `@JoinColumn(name="sosi_id", referencedColumnName="id", foreignKey = @ForeignKey(name="fk_sosi_user"))`
- `@ManyToMany`
- `@OrderBy("col1 ASC, col2 ASC")` : order by로 정렬함
- `@Transient` : 테이블의 컬럼과 매핑되지 않고 쓰이는 Attribute 를 정의하고자 할 때

### repository

CRUD 작업을 위해서 `JpaRepository(CrudRepository)<{entity}, {PK_TYPE}>` 인터페이스를 상속받는 `{tablename}Repository` 클래스를 생성한다.

기본적인 CRUD를 위한 method (`findAll`, `findOne`, `save` ...)가 이미 구현되어있다.

#### 핵심 method

- `<S extends T> S save(S entity);`
- `T findOne(ID primaryKey);`
- `Iterable<T> findAll();`
- `Long count();`
- `void delete(T entity);`
- `boolean exists(ID primaryKey);`

#### [JpaRepository Query Creation](https://docs.spring.io/spring-data/jpa/docs/1.7.0.RELEASE/reference/html/#jpa.query-methods.query-creation)

JpaRepository에서 `T findBy[COLUMN][조건](T ColumnName, Pageable Pageable);` 형식의 Method를 자동으로 실행해준다

유용하게도 OneToMany 관계의 Entity 값을 조건으로 쓸 수 있다 `T findBy[ENTITY][COLUMN][조건](T ColumnName, Pageable Pageable);`

Keyword | Sample | JPQL snippet
--- | --- | ---
And | `findByLastnameAndFirstname` | … where x.lastname = ?1 and x.firstname = ?2
Or | `findByLastnameOrFirstname` | … where x.lastname = ?1 or x.firstname = ?2
Is,Equals | `findByFirstname,findByFirstnameIs,findByFirstnameEquals` | … where x.firstname = 1?
Between | `findByStartDateBetween` | … where x.startDate between 1? and ?2
LessThan | `findByAgeLessThan` | … where x.age < ?1
LessThanEqual | `findByAgeLessThanEqual` | … where x.age ⇐ ?1
GreaterThan | `findByAgeGreaterThan` | … where x.age > ?1
GreaterThanEqual | `findByAgeGreaterThanEqual` | … where x.age >= ?1
After | `findByStartDateAfter` | … where x.startDate > ?1
Before | `findByStartDateBefore` | … where x.startDate < ?1
IsNull | `findByAgeIsNull` | … where x.age is null
IsNotNull,NotNull | `findByAge(Is)NotNull` | … where x.age not null
Like | `findByFirstnameLike` | … where x.firstname like ?1
NotLike | `findByFirstnameNotLike` | … where x.firstname not like ?1
StartingWith | `findByFirstnameStartingWith` | … where x.firstname like ?1 (parameter bound with appended %)
EndingWith | `findByFirstnameEndingWith` | … where x.firstname like ?1 (parameter bound with prepended %)
Containing | `findByFirstnameContaining` | … where x.firstname like ?1 (parameter bound wrapped in %)
OrderBy | `findByAgeOrderByLastnameDesc` | … where x.age = ?1 order by x.lastname desc
Not | `findByLastnameNot` | … where x.lastname <> ?1
In | `findByAgeIn(Collection<Age> ages)` | … where x.age in ?1
NotIn | `findByAgeNotIn(Collection<Age> age)` | … where x.age not in ?1
True | `findByActiveTrue()` | … where x.active = true
False | `findByActiveFalse()` | … where x.active = false
IgnoreCase | `findByFirstnameIgnoreCase` | … where UPPER(x.firstame) = UPPER(?1)

결과 개수를 지정하여 조회할 수도 있다.

```java
User findFirstByOrderByLastnameAsc();

User findTopByOrderByAgeDesc();

Page<User> queryFirst10ByLastname(String lastname, Pageable pageable);

Slice<User> findTop3ByLastname(String lastname, Pageable pageable);

List<User> findFirst10ByLastname(String lastname, Sort sort);

List<User> findTop10ByLastname(String lastname, Pageable pageable);
```

#### @Query 사용

@Query 어노테이션으로 직접 query를 작성할 수 있다(JPQL).

  ```java
  @Query("select u from User u where u.firstname like %?1")
  List<User> findByFirstnameEndsWith(String firstname);

  @Query("select u from User u where u.firstname = :firstname or u.lastname = :lastname")
  User findByLastnameOrFirstname(@Param("lastname") String lastname, @Param("firstname") String firstname);

  @Query("select u from #{#entityName} u where u.lastname = ?1")
  List<User> findByLastname(String lastname);
  ```

native query를 사용할 수 있다

  ```java
  @Query(value = "SELECT * FROM USERS WHERE EMAIL_ADDRESS = ?0", nativeQuery = true)
  User findByEmailAddress(String emailAddress);
  ```

#### 쿼리문 생성 후 slice, sort 제어

```java
Page<User> findByLastname(String lastname, Pageable pageable);

Slice<User> findByLastname(String lastname, Pageable pageable);

List<User> findByLastname(String lastname, Sort sort);

List<User> findByLastname(String lastname, Pageable pageable);
```

#### Specification

Repository에서 `JpaSpecificationExecutor` 인터페이스를 추가로 상속받는다

검색조건을 관리하는 `Specification` 클래스를 생성한다. `static`으로 method를 정의하고 `new Specification<T>()`를 return한다.

```java
public static Specification<T> findFoo(final long foo) {
  return new Specification<T>() {
    @Override
    public Predicate toPredicate(Root<T> root, CriteriaQuery<?> query, CriteriaBuilder cb) {
      return cb.equal/like(root.get("COLUMN"), foo/"%"+foo);
    }
  };
}
```

복수의 `Predicate`를 정의할 수도 있다

```java
Specification<Employee> specification = new Specification<Employee>() {
  public Predicate toPredicate(Root<Employee> root, CriteriaQuery<?> query, CriteriaBuilder builder) {
    List<Predicate> predicates = new ArrayList<Predicate>();
    predicates.add(builder.equal(root.get("id"), id));
    predicates.add(builder.equal(root.get("name"), name));
    predicates.add(builder.equal(root.get("address").get("city"), city));
    return builder.and(predicates.toArray(new Predicate[predicates.size()]));
  }
};
```

`@OneToMany` 관계의 데이터를 찾아올 때

```java
public static Specification<Board> findByComment(final String keyword) {
  return (Root<Board> root, CriteriaQuery<?> query, CriteriaBuilder cb) -> {
    Path<Collection<Comment>> comments = root.join("comments");
    return cb.like(comments.get("content"), "%"+keyword+"%");
  };
}
```

`JpaSpecificationExecutor` 인터페이스에 명시된 `Specification<?> spec`를 매개변수로 하는 method를 활용한다

```java
public Page<T> findAll(String keyword, Pageable pageable){
    return someRepository.findAll(Specifications.where(findFoo(keyword)), pageable);
}
```

### Pageable, Page, PageImpl

`PagingAndSortingRepository`에는 페이지 단위 입출력이 이미 구현되어 있다.
`Pageable` interface를 활용하는데 컨트롤러 매개변수 `Pageable pageable`로 구현체를 생성한다.

`Page<T> list = fooRepository.findAll(pageable);` 으로 페이지단위 데이터를 받아온다

페이지조회 조건을 주려면 `@PageableDefault(size = 5, sort = "id", direction = Direction.DESC) Pageable pageable`

사용자정의 조회 method에서도 두번째 인자로 사용할 수 있다.

```java
// repository interface 에 선언
Page<Emp> findFirst10ByJob(String job, Pageable pageable); // 페이징
List<Emp> findFirst10ByJob(String job, Sort sort); // 정렬
```

`Page<T> extends Slice<T>` 인터페이스에서 제공하는 method는 다음과 같다
해당하는 getter를 template engine에서 불러와서 사용할 수 있다. (`content`는 생략가능)

```java
long getTotalElements();  // 전체 데이터 개수
int getTotalPages();  // 전체 페이지 수

List<T> getContent(); // 조회된 데이터 목록
int getNumber();  // 현재 페이지 (0 ~)
int getNumberOfElements();  //현재 페이지 게시물 수
int getSize();  // 가져온 게시물 수
Sort getSort();  // 정렬정보
boolean isFirst();  // 현재 페이지가 첫 페이지인지
boolean isLast();  // 현재 페이지가 마지막 페이지 인지
```

반환받은 페이지 객체를 직렬화한 형태는 다음과 같다.

```json
{
  "links" : [ { "rel" : "next",
                "href" : "http://localhost:8080/persons?page=1&size=20" }
  ],
  "content" : [
     … // 20 Person instances rendered here
  ],
  "pageMetadata" : {
    "size" : 20, // 가져온 게시물 수
    "totalElements" : 30, // 전체 게시물 수
    "totalPages" : 2, // 전체 페이지
    "number" : 0 // 현재 페이지
  }
}
```

### 상위 Entity, Auditing

데이터 객체의 중복을 없애기 위해서 상위 Entity를 생성할 수 있다. `AbstractEntity.class`를 생성한다.

Auditing을 활용하여 변화를 감지하고 자동으로 값을 갱신할 수 있다.

사용예제

  ```java
  @MappedSuperclass
  @EntityListeners(AuditingEntityListener.class)
  public abstract class AbstractEntity {

    @Id
    @GeneratedValue
    private Long id;

    @CreatedDate
    private Date regdate;

    @LastModifiedDate
    private Date moddate;

    public Long getId() {
      return id;
    }

    public String getRegdate() {
      if (regdate == null) {
        return "";
      }
      return LocalDateTime
          .ofInstant(regdate.toInstant(), ZoneId.systemDefault())
          .format(DateTimeFormatter.ofPattern("yyyy.MM.dd HH:mm"));
    }

    public Date getModdate() {
      return moddate;
    }

    @Override
    public int hashCode() {
      final int prime = 31;
      int result = 1;
      result = prime * result + ((id == null) ? 0 : id.hashCode());
      return result;
    }

    @Override
    public boolean equals(Object obj) {
      if (this == obj)
        return true;
      if (obj == null)
        return false;
      if (getClass() != obj.getClass())
        return false;
      AbstractEntity other = (AbstractEntity) obj;
      if (id == null) {
        if (other.id != null)
          return false;
      } else if (!id.equals(other.id))
        return false;
      return true;
    }

    @Override
    public String toString() {
      return "[id=" + id + ", regdate=" + regdate + ", moddate=" + moddate + "]";
    }
  }
  ```

생성한 `AbstractEntity.class`는 데이터객체에서 상속받아 사용한다.

위에서 `@CreatedDate`와 `@LastModifiedDate` 어노테이션을 적용하고 Spring-data가 이를 감지하게 하려면 설정값이 필요하다

**`fooApplication.java` (Spring Boot 설정파일) 클래스 상단에 `@EnableJpaAuditing` 어노테이션을 명시한다.**

#### Auditing 사용시 `LocalDate` `LocalDateTime` 처리

`LocalDate`, `LocalDateTime`을 DB에 입력하려고 하면
`Caused by: com.mysql.jdbc.MsqlDataTruncation: Data truncation: Incorrect dateme value:` 에러가 발생하거나
`tinyblob` 타입으로 저장되는 경우가 발생한다.

`@Convert(converter = Jsr310JpaConverters.LocalDateTimeConverter.class)`

위 애노테이션을 Entity `LocalDate` 필드에 명시하면 변환가능하다.

### Entity에서 JSON 사용처리

`@JsonProperty` : 변환 처리 명시

`@JsonIgnore` : 변환하지 않을 항목

## 영속성 컨텍스트

### 트랜잭션 범위의 영속성 컨텍스트

Service - Repository 범위에서 준영속 -> 영속 유지 트랜잭션 내에서는 동일한 영속성 컨텍스트를 사용한다.
