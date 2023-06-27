# nodejs performance analyzer

## clinic.js

nodejs 성능 문제 분석도구

- <https://clinicjs.org/>
- [Node.js에서 CPU-Intensive한 코드를 찾아내는 방법](https://youtu.be/pMo8M5cqxfQ?t=3569)

### Doctor

<https://clinicjs.org/documentation/doctor/>

다음 지표의 시계열 데이터 출력

- CPU usage
- Memory Usage
- Event Loop Delay ms
- Active Handles

#### Doctor 실행

```sh
clinic doctor --on-port 'autocannon localhost:$PORT' -- node dist/main.js
```

- `--on-port`: 서버에서 포트를 수신하기 시작하면 스크립트를 실행함
- `$PORT`: 서버가 첫 번째로 수신하는 포트가 환경변수로 지정됨
- `-- <command>`: 더블대시 뒤에 프로파일링을 하려는 서버를 실행하는 명령어를 지정

### Flame

> [FlameGraph](https://github.com/brendangregg/FlameGraph) 출력에 사용
>
> FlameGraph == Stack trace visualizer

<https://clinicjs.org/documentation/flame/>

### Bubbleprof

<https://clinicjs.org/documentation/bubbleprof/>

### HeapProfiler

<https://clinicjs.org/documentation/heapprofiler/>

## 부하테스트 도구

- <https://github.com/mcollina/autocannon>
- <https://github.com/grafana/k6>
- <https://github.com/apache/jmeter>
- <https://github.com/naver/ngrinder>

## deoptexplorer-vscode

- <https://devblogs.microsoft.com/typescript/introducing-deopt-explorer/>
- <https://github.com/microsoft/deoptexplorer-vscode>
