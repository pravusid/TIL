# Spring Framework with Kotlin

## Refs

- <https://docs.spring.io/spring-framework/reference/languages/kotlin.html>
- [코프링 프로젝트 투입 일주일 전: 주니어 개발자의 코틀린 도입 이야기](https://www.inflearn.com/course/lecture?courseSlug=인프콘2023-다시보기&unitId=177910)

## Starter

Kotlin Spring MVC 프로젝트 생성은 <https://start.spring.io/>에서 가능하다

`build.gradle` 생성 예시

```groovy
buildscript {
  ext {
    kotlinVersion = '1.2.51'
    springBootVersion = '2.0.4.RELEASE'
  }
  repositories {
    mavenCentral()
  }
  dependencies {
    classpath("org.springframework.boot:spring-boot-gradle-plugin:${springBootVersion}")
    classpath("org.jetbrains.kotlin:kotlin-gradle-plugin:${kotlinVersion}")
    classpath("org.jetbrains.kotlin:kotlin-allopen:${kotlinVersion}")
  }
}

apply plugin: 'kotlin'
apply plugin: 'kotlin-spring'
apply plugin: 'eclipse'
apply plugin: 'org.springframework.boot'
apply plugin: 'io.spring.dependency-management'

group = 'kr.pravusid'
version = '0.0.1-SNAPSHOT'
sourceCompatibility = 1.8
compileKotlin {
  kotlinOptions {
    freeCompilerArgs = ["-Xjsr305=strict"]
    jvmTarget = "1.8"
  }
}
compileTestKotlin {
  kotlinOptions {
    freeCompilerArgs = ["-Xjsr305=strict"]
    jvmTarget = "1.8"
  }
}

repositories {
  mavenCentral()
}


dependencies {
  compile('org.springframework.boot:spring-boot-starter-data-jpa')
  compile('org.springframework.boot:spring-boot-starter-security')
  compile('org.springframework.boot:spring-boot-starter-web')
  compile('com.fasterxml.jackson.module:jackson-module-kotlin')
  compile("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
  compile("org.jetbrains.kotlin:kotlin-reflect")
  runtime('org.springframework.boot:spring-boot-devtools')
  runtime('com.h2database:h2')
  testCompile('org.springframework.boot:spring-boot-starter-test')
  testCompile('org.springframework.security:spring-security-test')
}
```

Kotlin 컴파일러에서 JSR-305를 위한 옵션은 다음이 있다

`-Xjsr305=strict`: produce compilation error (experimental feature).
`-Xjsr305=warn`: produce compilation warnings (default behaviour)
`-Xjsr305=ignore`: do nothing.

JSR-305(소프트웨어 결함 탐지를 위한 애노테이션)에서는 다음내용을 처리한다

- Nonnull
- Nullable
- CheckReturnValue
- OverridingMethodsMustInvokeSuper
- ParametersAreNullableByDefault
- ParametersAreNonnullByDefault

## data class 사용

Kotlin으로 spring data JPA를 사용하다 보면 기본설정으로 오류 발생함

`org.springframework.orm.jpa.JpaSystemException: No default constructor for entity: : kr.pravusid.User; nested exception is org.hibernate.InstantiationException: No default constructor for entity: : kr.pravusid.User`

이는 기본생성자가 없기 때문임

```kotlin
@Entity
data class User(...) {
    constructor() : this(null)
}
```

기본생성자를 만들어주면 되지만 모든 Entity data class에 처리하는 것보다 나은 방법이 있다.
no argument constructor를 자동생성하는 `apply plugin: "kotlin-noarg"`을 사용하면 된다.

kotlin-jpa plugin은 kotlin-noarg plugin을 포함하며 jpa를 지원한다.

`build.gradle`

```groovy
buildscript {
    dependencies {
        classpath "org.jetbrains.kotlin:kotlin-noarg:$kotlin_version"
    }
}

apply plugin: "kotlin-jpa"
```

## final class 문제

Kotlin의 클래스는 기본적으로 final이다

lazy loading등에 활용되는 프록시 객체를 생성하려면 상속을 사용해야 하는데 불가능해진다.

이를 처리하기 위해 kotlin-allopen plugin을 사용한다

`build.gradle`

```groovy
buildscript {
    dependencies {
        classpath "org.jetbrains.kotlin:kotlin-allopen:$kotlin_version"
    }
}

apply plugin: "kotlin-allopen"

allOpen {
    annotation "javax.persistence.Entity"
}
```

## 연관 객체의 id에 직접 접근

`Post`와 연관객체 `User`(N:1 관계) 사이에서 `id`값을 조회하는 경우를 보자

`Post.User.Id`는 사실 `User`를 조회해보지 않아도 `post.user_id` 컬럼으로 알아낼 수 있다.
데이터베이스 테이블이 그렇게 되어있고, Hibernate에서도 지원하는 기능이다.

Kotlin data class 필드에 애노테이션을 달면, getter 메소드가 아니라 JVM 필드 애노테이션이 붙는다.

그러면 프로퍼티 접근 모드가 아니라 필드 접근 모드가 되고, 이 경우 Hibernate는 지연 로딩을 지원하지 않게 된다.
이 경우 getter에 애노테이션을 붙이면 된다.

```kotlin
@Entity
data class User(
        @get:Id
        @get:GeneratedValue(strategy = GenerationType.AUTO)
        var id: Int? = null,
        var name: String
)
```
