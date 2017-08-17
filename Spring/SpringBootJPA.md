# JPA - hibernate

## 시작

## 초기설정

- build.gradle에 추가
  ```groovy
  compile('org.springframework.boot:spring-boot-starter-data-jpa')
  runtime('mysql:mysql-connector-java')
  compileOnly('org.projectlombok:lombok')
  ```

- application.properties
  > The data source properties starting with spring.datasource.* will automatically be read by spring boot JPA. To change the hibernate properties we will use prefix spring.jpa.properties.* with hibernate property name. On the basis of given data source URL, spring boot can automatically identify data source driver class. So we need not to configure diver class.
  ```
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

- @Entity
- @Table(name="articles") : 기존 테이블과 연결
- @Id
- @GeneratedValue
  - strategy = GenerationType.AUTO
- @Column
  - name : 기존 column 이름
- @OneToMany
- @JoinColumn(name="sosi_id", referencedColumnName="id")

### repository

CRUD 작업을 위해서 JpaRepository<{tablename}, {PK_TYPE}> 클래스를 상속받는 {tablename}Repository 클래스를 생성한다.  
기본적인 CRUD를 위한 method (findAll, findOne, save ...)가 이미 구현되어있다.

### Pageable

Repository에는 페이지 단위 입출력이 이미 구현되어 있다.