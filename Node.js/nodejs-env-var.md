# Node.js 환경변수

## 환경변수 사용

- 런타임 실행할 때 환경변수 정의: `FOO=BAR node index.js`
- `dotenv` 패키지 사용: <https://github.com/motdotla/dotenv>
- 런타임 실행할 때 [[nodejs#Node.js CLI|`--env-file=config`]] 인자 사용 (v20.6.0 버전에서 추가됨)

## TypeScript 환경변수 타입 정의

`env.d.ts`

```ts
declare global {
  namespace NodeJS {
    interface ProcessEnv {
      FOO: string
    }
  }
}

export {};
```
