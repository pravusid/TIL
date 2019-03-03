# Spring Boot 배포

## Build 및 실행

- 빌드: `./gradlew clean build`
- 빌드 결과물: `build/libs/*.jar`
- 실행: `nohup java -jar -Dspring.profiles.active=<활성화프로필> <파일명>.jar &`

## build executable jar

<https://spring.io/guides/gs/spring-boot/#scratch>

```groovy
buildscript {
    ext {
        springBootVersion = '2.0.1.RELEASE'
    }
    repositories {
        mavenCentral()
    }
    dependencies {
        classpath('org.springframework.boot:spring-boot-gradle-plugin:${springBootVersion}')
    }
}

apply plugin: 'org.springframework.boot'
apply plugin: 'io.spring.dependency-management'
apply plugin: 'java'
apply plugin: 'idea'
```

Boot 2.x 에서는 `bootJar` / `bootWar` task가 빌드와 관련있다

또한 `bootRepackage`가 `jar` task를 확장한 `bootJar`와 `war` task를 확장한 `bootWar`로 분리되었으며
각각 `jar` / `war` task를 비활성화 시킨다

```groovy
bootJar {
    baseName = 'spring-boot-vue'
    version = version
    mainClassName = 'kr.pravusid.Application'
}
```

Boot 1.x 에서는 `bootRepackage` task가 빌드와 관련있다

```groovy
jar {
    baseName = 'spring-boot-vue'
    version = version
}

bootRepackage {
    mainClass = 'kr.pravusid.WebApplication'
}
```

## 운영환경 분리

운영환경의 설정파일(properties)은 개발소스와 분리하여 서버에 별도 보관하는 것이 좋음
