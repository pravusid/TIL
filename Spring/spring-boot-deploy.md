# Spring Boot 배포

## Build 및 실행

- 빌드: `./gradlew clean build`
- 빌드 결과물: `build/libs/*.jar`
- 실행: `nohup java -jar -Dspring.profiles.active=<활성화프로필> <파일명>.jar &`

## Fully executable jar

`build.gradle` 수정

```groovy
buildscript {
    ext {
        springBootVersion = '1.5.6.RELEASE'
    }
    repositories {
        mavenCentral()
    }
    dependencies {
        classpath("org.springframework.boot:spring-boot-gradle-plugin:${springBootVersion}")
        classpath 'io.spring.gradle:dependency-management-plugin:1.0.5.RELEASE'
    }
}

jar {
    manifest {
        attributes  'Title': 'boot-vue', 'Version': 1.0, 'Main-Class': 'kr.pravusid.WebApplication'
    }
    dependsOn configurations.runtime
    from {
        configurations.compile.collect {it.isDirectory()? it: zipTree(it)}
    }
}
```

## 운영환경 분리

운영환경의 설정파일(properties)은 개발소스와 분리하여 서버에 별도 보관하는 것이 좋음
