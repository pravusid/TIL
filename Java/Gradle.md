# Gradle

## Gradle 이란

## 명령어

버전 업데이트

`./gradlew wrapper --gradle-version {version}`

gradle을 실행하면 버전 변경사항을 확인하고 새로운 wrapper를 다운로드 함

`./gradlew tasks`

## tips

### dependency 포함한 jar build

`build.gradle`에 다음 내용 추가 후 `fatJar` 실행

```groovy
jar {
    manifest {
        attributes 'Main-Class': 'kr.pravusid'
    }
}

task fatJar(type: Jar) {
    manifest.from jar.manifest
    classifier = 'all'
    from {
        configurations.runtime.collect { it.isDirectory() ? it : zipTree(it) }
    } {
        exclude "META-INF/*.SF"
        exclude "META-INF/*.DSA"
        exclude "META-INF/*.RSA"
    }
    with jar
}
```
