# Travis CI

Github에서 관리하는 Continuous Integration 프로젝트

<https://docs.travis-ci.com/>

## Java

### gradle wrapper 권한

gradle 사용시 권한설정에 유의해야 한다

`gradlew`, `gradle-wrapper.jar` 파일의 권한이 755여야 한다

`chmod +x gradlew`, `chmod +x gradle/wrapper/gradle-wrapper.jar`

### 설정파일 누락 확인

Database를 사용하는 경우 (test cases) 설정파일(application.yml)이 빠져있지 않은지 확인해야 한다
