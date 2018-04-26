# Spring Boot 배포

## Jar로 배포

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
        attributes  'Title': 'boot-vue', 'Version': 1.0, 'Main-Class': 'com.talsist.TalsistApplication'
    }
    dependsOn configurations.runtime
    from {
        configurations.compile.collect {it.isDirectory()? it: zipTree(it)}
    }
}
```

gradle `build > build` 실행
