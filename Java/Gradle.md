# Gradle

## Gradle 이란

## 명령어

## tips

### dependency 포함한 jar build

`build.gradle`에 다음 내용 추가 후 `fatJar` 실행

```groovy
jar {
    manifest {
        attributes 'Main-Class': 'kr.co.iot4health.Main'
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
