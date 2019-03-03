# Gradle

## Gradle 이란

## 명령어

버전 업데이트

`./gradlew wrapper --gradle-version {version}`

gradle을 실행하면 버전 변경사항을 확인하고 새로운 wrapper를 다운로드 함

`./gradlew tasks`

## build

`build.gradle`에 다음 내용 추가 후 `build` 실행한다

> Kotlin의 경우 MainClass의 끝에 `Kt`를 붙인다

```groovy
jar {
    manifest {
        attributes 'Main-Class': 'kr.pravusid.Application'
    }
    from {
        configurations.compile.collect { it.isDirectory() ? it : zipTree(it) }
    }
}
```
