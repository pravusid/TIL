# Tornado FX

JavaFx의 Kotlin Wrapper Framework

## 배포

`build.gradle`에 manifest를 명시해야 한다

```groovy
jar {
    manifest {
        attributes(
                'Class-Path': configurations.compile.collect { it.getName() }.join(' '),
                'Main-Class': 'kr.pravusid.app.MyApp'
        )
    }
    from(configurations.compile.collect { entry -> zipTree(entry) }) {
        exclude 'META-INF/MANIFEST.MF'
        exclude 'META-INF/*.SF'
        exclude 'META-INF/*.DSA'
        exclude 'META-INF/*.RSA'
    }
}
```
