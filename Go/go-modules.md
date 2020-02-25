# Go Modules

<https://github.com/golang/go/wiki/Modules>

Go Modules는 1.11 버전에서 새로 추가된 의존성 관리도구이다.

`go.mod`가 존재하는 경로가 모듈 루트이며, 다음 네 가지 디렉티브가 있다: module, require, replace, exclude

## 시작

### 모듈 생성

`go mod init [module-repository]`

모듈을 생성하면 `go.mod` 파일과 `go.sum` 파일이 생성된다

`go.sum` 파일의 경우 관리하는 패키지를 검증하기 위한 방법이다.

패키지의 유효성 검증을 위해 매번 checksum을 얻어오지 않고, 미리 계산된 checksum으로 현재 연결된 패키지의 유효성을 검사한다.

`// indirect` 의존성은 직접 사용하는 의존성에 의해 간접적으로 참조하는 의존성이다

### 모듈에 의존성 추가

해당 모듈을 인식하는 범위(루트의 서브경로)에서 의존성을 추가한다.

`-u` 커맨드를 사용하여 존재하는 패키지의 경우 버전 업데이트를 할 수 있다.

`go get [-u] <package-repository>`

## 명령어

`go help mod [command]`

> `go mod <command> [arguments]`

- `go mod edit`: go.mod 파일 수정
- `go mod download`: 로컬 캐시에 모듈 다운로드
- `go mod tidy`: 로컬 캐시에 누락된 모듈은 다운로드, 사용되지 않는 모듈은 삭제
- `go mod verify`: 로컬 캐시에 있는 의존성 검증
- `go mod vendor`: 의존성 복사본이 저장된 vendor 디렉토리 생성
- `go mod graph`: 모듈 의존 그래프 출력
- `go mod why`: 특정 모듈 필요 이유 출력

## 프록시

Go moudles는 프록시를 지원한다

Go modules에서는 패키지를 찾을때 $GOPATH를 우선 검색하고, 패키지가 없다면 패키지명에 명시된 원격 repository에서 패키지를 찾는다.

원격 repository를 사용하는 대신 임의의 프록시를 지정하여 패키지를 관리할 수 있다.

프록시를 설정하려면 환경변수에 지정하면 된다: `export GOPROXY=http://goproxy:8080`

또는 의존성 추가시 프록시를 매번 지정할 수 있다: `http_proxy=http://goproxy:8080 go get <package-repository>`
