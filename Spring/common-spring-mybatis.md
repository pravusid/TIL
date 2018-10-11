# MyBatis

## Common MyBatis

### 기본 환경설정

db.properties

```text
driver = oracle.jdbc.driver.OracleDriver
url = jdbc:oracle:thin:@localhost:1521:ORCL
username = scott
password = tiger
maxActive = 10
maxIdle = 5
maxWait = -1
```

Config.xml

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE configuration
  PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
  "http://mybatis.org/dtd/mybatis-3-config.dtd">
<properties resource="db.properties"/>
<configuration>
  <environments default="development">
    <environment id="development">
      <transactionManager type="JDBC"/> <!--JDBC or MANAGED-->
      <dataSource type="POOLED"> <!--UNPOOLED or POOLED or JNDI-->
        <property name="driver" value="${driver}"/>
        <property name="url" value="${url}"/>
        <property name="username" value="${username}"/>
        <property name="password" value="${password}"/>
      </dataSource>
    </environment>
  </environments>
  <mappers>
    <mapper resource="org/mybatis/example/BlogMapper.xml"/>
  </mappers>
</configuration>
```

### mybatis에서 JNDI(Java Naming and Directory Interface) 사용

config.xml

```xml
<transactionManager type="JDBC"/>
  <dataSource type="JNDI">
    <property name="data_source" value="java:comp/env/jdbc/myoracle"/>
  </dataSource>
```

### Mapper.xml

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper
  PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
  "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="org.mybatis.example.BlogMapper">
  <select id="selectBlog" resultType="Blog">
    select * from Blog where id = #{id}
  </select>
</mapper>
```

### mybatis에서 mapper의 namespace 줄여쓰기

config.xml

```xml
<configuration>
  <typeAliases>
    <typeAlias type="app.emp.Emp" alias="Emp"/>
  </typeAliases>
```

### 중복되는 쿼리 모아쓰기

Mapper.xml

```xml
정의
<sql id="baseEmp">
    select * from emp
</sql>

사용
<include refid="baseEmp"></include>
```

### selectkey

#### 마지막 insert된 레코드의 시퀀스를 조회 (Mapper에서 사용)

```xml
<insert id="insert">
  insert into guest(guest_id, name, phone, addr, email)
    values(seq_guest.nextval, #{name}, #{phone}, #{addr}, #{email})
  <selectKey keyProperty="guest_id" order="AFTER" resultType="int">
    select seq_guest.currval from dual
  </selectKey>
</insert>
```

#### sequence 없이 insert

```xml
<insert id="boardInsert" parameterType="Board">
  <selectKey keyProperty="no" resultType="int" order="BEFORE">
    SELECT NVL(MAX(no)+1, 1) as no FROM mvcBoard
  </selectKey>
  INSERT INTO mvcBoard(no, name, subject, content, pwd, group_id)
  VALUES(#{no},#{name},#{subject},#{content},#{pwd},(SELECT NVL(MAX(group_id)+1, 1) FROM mvcBOARD))
</insert>
```

#### 컬럼명 수동지정과 JOIN

``` xml
<resultMap type="Emp" id="myMap">
  <result column="empno" property="empno"/>
  <result column="ename" property="ename"/>
  <!-- DTO간 1:1 관계를 위한 association -->
  <association property="dept" javaType="Dept" select="**selectDept**" column="deptno"/>
  <!-- DTO간 1:N 관계를 위한 collection -->
  <collection property="list" ofType="Comments" column="news_id" select="**selectComments**">
    <result column="comments_id" property="comments_id"/>
    <result column="msg" property="msg"/>
  </collection>
</resultMap>

<select id="empDeptJoin" resultMap="myMap">
  select * from emp
</select>
<select id="selectDept" resultType="Dept">
  select * from dept where deptno=#{deptno}
</select>
```

### SqlSession

MyBatis에서 쿼리문을 실행하기 위해서 sqlSession이 필요한데 이를 얻기위해서는 sqlSessionFactory객체가 필요하며 환경설정xml을 스트림을 이용하여 읽어야한다.
코드를 DAO에 일일이 작성하지 않기 위해 공통 클래스를 정의한다.

```java
String resource = "org/mybatis/example/mybatis-config.xml";
InputStream inputStream = Resources.getResourceAsStream(resource);
SqlSessionFactory sqlSessionFactory = new SqlSessionFactoryBuilder().build(inputStream);
```

## Spring MyBatis

### MyBatis 연결 설정 (XML only)

application-context.xml

```xml
<!-- 오라클 정보를 모아서 전송 : driver, url, username, password -->
<bean id="ds" class="org.apache.commons.dbcp.BasicDataSource"
  p:driverClassName="oracle.jdbc.driver.OracleDriver"
  p:url="jdbc:oracle:thin:@localhost:1521:ORCL"
  p:username="scott"
  p:password="tiger" />
<!-- MyBatis : SqlSessionFactory -->
<bean id="ssf" class="org.mybatis.spring.SqlSessionFactoryBean"
  p:dataSource-ref="ds"
  p:configLocation="classpath:Config.xml" />
<!-- 동시에 모아서 처리 : SqlSessionTemplate -->
<bean id="sst" class="org.mybatis.spring.SqlSessionTemplate">
  <constructor-arg ref="ssf"/>
</bean>
<!-- DAO로 전송 -->
<bean id="dao" class="com.idpravus.dao.EmpDAO"
  p:sqlSessionTemplate-ref="sst" />
<!-- DAO를 비즈니스 로직 클래스에 DI -->
```

### MyBatis 설정 (XML + Annotation)

application-context.xml

```xml
<!-- 사용자 등록 -->
<context:annotation-config/>
<context:component-scan base-package="com.idpravus.*"/>

<!-- 오라클 정보 : DataSource -->
<bean id="ds" class="org.apache.commons.dbcp.BasicDataSource"
  p:driverClassName="oracle.jdbc.driver.OracleDriver"
  p:url="jdbc:oracle:thin:@localhost:1521:ORCL"
  p:username="scott"
  p:password="tiger" />

<bean id="ssf" class="org.mybatis.spring.SqlSessionFactoryBean"
  p:dataSource-ref="ds" />

<bean id="mapper" class="org.mybatis.spring.mapper.MapperFactoryBean"
  p:sqlSessionFactory-ref="ssf"
  p:mapperInterface="com.idpravus.dao.EmpMapper" />
```

### Interface + Annotation으로 구현한 Mapper

```java
@Select("SELECT empno,ename,job,hiredate,sal FROM emp")
public List<EmpVO> empAllData();

@Select("SELECT empno,ename,job,hiredate,sal FROM emp WHERE empno=#{empno}")
public EmpVO empFindData(int empno);

@SelectKey(keyProperty="empno", resultType=int.class, before=true,
    statement="SELECT NVL(MAX(empno)+1,1) as empno FROM emp")
@Insert("INSERT INTO emp VALUES(#{empno},#{ename},#{job},7788,SYSDATE,#{sal},100,10)")
public void empInsertData(EmpVO vo);

@Options(useGeneratedKeys=true, keyProperty="id", keyColumn="id")
@Insert("INSERT INTO ingredient VALUES(ingredient_seq.nextval,#{name},#{cal})")
public int insertIngr(IngredientVO vo);
```

### MyBatis Annotation @Many 1:n 연결(Join)

```java
@Results(value= {
    @Result(property="id", column="id"),
    @Result(property="name", column="name"),
    @Result(property="religion", column="id", javaType=List.class,
      many=@Many(select="selectIngrReligion"))
})
@Select("SELECT *"
      + " FROM ("
        + " SELECT X.*, ROWNUM as rnum"
        + " FROM ("
          + " SELECT id, name, cal"
          + " FROM ingredient"
          + " ORDER BY id ASC"
          + " ) X"
        + " WHERE ROWNUM <= #{end}"
        + " )"
      + " WHERE rnum > #{start}")
public List<IngredientVO> selectIngrList(Map map);

@Select("SELECT id, name, ingredient_id"
  + " FROM ingr_religion i, religion r"
  + " WHERE i.religion_id=r.id AND ingredient_id=#{id}")
public List<ReligionVO> selectIngrReligion();
```

### 활용

```xml
<!-- 해당하는 회원ID 숫자 -->
<select id="checkId" parameterType="Map" resultType="int">
  SELECT count(member_idx) FROM member WHERE member_id=#{memberId}
</select>

<!-- 회원 조회 / 동적쿼리 -->
<select id="mList" parameterType="Map" resultType="Map">
  SELECT
    member_idx,
    member_id AS memberId,
    member_nick AS memberNick,
    member_name AS memberName,
    email, DATE_FORMAT(create_date, '%Y-%m-%d') AS createDate
  FROM member
  WHERE 1=1
  <if test="searchType!=null and searchType==1">
    AND ( member_id like concat('%',#{searchText},'%') OR email like concat('%',#{searchText},'%') )
  </if>
  <if test="searchType!=null and searchType==2">
    AND member_id like concat('%',#{searchText},'%')
  </if>
  <if test="searchType!=null and searchType==3">
    AND email like concat('%',#{searchText},'%')
  </if>
  ORDER BY ${sidx} ${sord}
  LIMIT ${startIdx}, ${rows}
</select>

<!-- 자동증가하는 PK의 값을 boardSeq를 Key로 반환하여 돌려줌.  -->
<insert id="write" parameterType="Map" useGeneratedKeys="true" keyProperty="boardSeq">
  INSERT INTO board (`type_seq`,`member_idx`,`member_id`, `member_nick`, `title`, `content`, `has_file`, `create_date`)
  VALUES (#{typeSeq},#{memberIdx},#{memberId},#{memberNick},#{title},#{contents}, '0' ,now())
</insert>

<!-- 조회수 1 증가 -->
<update id="updateHit" parameterType="int">
  UPDATE board SET hits=hits+1
  WHERE `type_seq`=#{0} AND `board_seq`=#{1}
</update>

<!-- 해당 PK를 가지는 게시글의 내용과 수정 날짜를 변경 -->
<update id="updateBoard" parameterType="Map">
  UPDATE board SET title=#{title}, content = #{contents}, update_date = now()
  WHERE `type_seq`= #{typeSeq} AND `board_seq`= #{boardSeq}
</update>
```
