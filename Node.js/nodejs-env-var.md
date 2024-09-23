# Node.js 환경변수

## 환경변수 사용

- 런타임 실행할 때 환경변수 정의: `FOO=BAR node index.js`
- `dotenv` 패키지 사용: <https://github.com/motdotla/dotenv>
- 런타임 실행할 때 [[nodejs#Node.js CLI|`--env-file=config`]] 인자 사용 (v20.6.0 버전에서 추가됨)
- [`process.loadEnv(path)`](https://nodejs.org/docs/latest/api/process.html#processloadenvfilepath) 사용 (Added in: v21.7.0)

## TypeScript 환경변수 타입 정의

[[typescript-declaration-files#process.env]]
