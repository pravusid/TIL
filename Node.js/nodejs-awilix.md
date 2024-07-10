# Awilix

<https://github.com/jeffijoe/awilix>

## Injection modes

Proxy (기본설정)

```ts
class UserService {
  constructor(di: { emailService: EmailService; logger: Logger }) {
    this.emailService = di.emailService;
    this.logger = di.logger;
  }
}
```

Classic (코드에 minified를 적용하면 사용할 수 없다)

```ts
class UserService {
  constructor(emailService: EmailService, logger: Logger) {
    this.emailService = emailService;
    this.logger = logger;
  }
}
```

## 변수명

- 생성자에 주입되는 종속성의 변수명은 파일이름을 설정의 formatName을 사용하여 변환한 것임
- `RESOLVER` 설정을 통해 임의 이름을 사용할 수 있음

```ts
export class Foo {
  static [RESOLVER]: BuildResolverOptions<unknown> = { name: 'bar' };
}
```

## 종속성 중복등록

- 동일한 이름의 종속성 (별도의 `RESOLVER` 설정이 없다면 동일한 파일명)을 중복(여러번) 등록하면 마지막으로 등록한 종속성을 사용함
- 종속성 등록순서는 [`listModules`](https://github.com/jeffijoe/awilix/blob/65b4d4246aafec5f31d760398ed644abc3fb48ba/src/load-modules.ts#L107) 순서를 따름
- `node:path.resolve` 실행결과의 순서임

## Auto Loading Modules

<https://github.com/jeffijoe/awilix?tab=readme-ov-file#auto-loading-modules>

자동으로 불러올 종속성은 다음 설정 중 하나가 적용되어야 한다

- `[RESOLVER]`
- `export default`

그리고 나서 `container.loadModules()`를 호출한다
