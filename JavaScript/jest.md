# Jest

<https://jestjs.io/docs/en/getting-started.html>

Facebook 에서 주도하는 자바스크립트 테스트 툴 (React와 함께 성장)

## 기존의 JavaScript test stack

- test runner : Mocha
- assertion(문법) : Chai
- mocking : Sinon
- coverage : Istanbul

## NODE_ENV

jest 기본 `NODE_ENV`는 `test`임

> <https://github.com/facebook/jest/blob/master/packages/jest-cli/bin/jest.js#L12-L14>

## jest설정 (ts기준)

<https://jest-bot.github.io/jest/docs/configuration.html>

`npm i -D @types/jest jest ts-jest`

`packages.json`

```json
{
  "scripts": {
    "test": "jest",
    "test:watch": "jest --watch",
    "test:cov": "jest --coverage"
  }
}
```

`jest.config.js`

```js
module.exports = {
  setupFiles: ["dotenv/config"],
  moduleFileExtensions: ["js", "jsx", "json", "ts", "tsx"],
  testRegex: "^.+\\.spec\\.(js|jsx|ts|tsx)$",
  transform: {
    "^.+\\.(ts|tsx)$": "ts-jest"
  },
  coverageDirectory: "coverage",
  collectCoverageFrom: ["src/**/*.{js,jsx,ts,tsx}"],
  testEnvironment: "node"
};
```

### tsconfig 별도 적용

`jest.config.js`

```js
module.exports = {
  // ...
  globals: {
    "ts-jest": {
      tsConfig: "tsconfig.spec.json"
    }
  }
  // ...
};
```

`tsconfig.spec.json`

```json
{
  "extends": "./tsconfig.base.json",
  "compilerOptions": {
    "types": ["jest", "node"],
    "strict": false
  },
  "include": ["src/**/*.spec.*", "*.js"]
}
```

## 실행

유닛 테스트를 위한 파일명이 `___.spec.js/ts`로 끝나게 한다

- 전체 실행: `jest`
- 코드 변경시 테스트 실행: `jest --watch`
- 부분 실행 `--testNamePattern(-t) '<regexp>'`
- 부분 실행: `--watch '<regexp>'`

## 기본 문법

### Test Case

```js
describe("my beverage", () => {
  test("is delicious", () => {
    expect(myBeverage.delicious).toBeTruthy();
  });

  it("is not sour", () => {
    expect(myBeverage.sour).toBeFalsy();
  });
});
```

- `describe`: 테스트 묶음

  - `describe.each(table)(name, fn, timeout)`
  - `describe.only(name, fn)`
  - `describe.only.each(table)(name, fn)`
  - `describe.skip(name, fn)`
  - `describe.skip.each(table)(name, fn)`

- `test`: 개별 test case

  - `test(name, fn, timeout)`
  - `test.each(table)(name, fn, timeout)`
  - `test.only(name, fn, timeout)`
  - `test.only.each(table)(name, fn)`
  - `test.skip(name, fn)`
  - `test.skip.each(table)(name, fn)`

- `it`: test의 alias (BDD 문법)

### matcher

- `expect({object})`
- `toBe({object})`: 타입과 값 검사
- `toEqual({object})`: 값 검사
- `toBeNull()`: null임을 검사
- `toBeUndefined()`: undefined임을 검사
- `toThrow({Error})`: Error가 발생하는지 확인한다
- ...
- `not`: 위의 matcher의 부정형
  - `not.toBeNull()`: null이 아님을 검사
  - ...

### teardown

#### 여러 테스트를 위한 반복적 setup

```js
beforeEach(() => {
  initializeCityDatabase();
});

afterEach(() => {
  clearCityDatabase();
});

test("city database has Vienna", () => {
  expect(isCity("Vienna")).toBeTruthy();
});

test("city database has San Juan", () => {
  expect(isCity("San Juan")).toBeTruthy();
});
```

setup들은 테스트와 마찬가지로 비동기를 코드를 다룬다.
만약 `initializeCityDatabase()`가 promise를 반환할 때 다음처럼 쓰면 setup은 비동기 코드를 처리해준다.

```js
beforeEach(() => {
  return initializeCityDatabase();
});
```

#### 1회성 setup

만약 파일의 시작에서 한 번의 setup만 필요하다면 다음처럼 기술하면 된다.
마찬가지로 비동기를 처리하려면 promise를 반환하면 setup 메소드가 처리해준다.

```js
beforeAll(() => {
  return initializeCityDatabase();
});

afterAll(() => {
  return clearCityDatabase();
});

test("city database has Vienna", () => {
  expect(isCity("Vienna")).toBeTruthy();
});

test("city database has San Juan", () => {
  expect(isCity("San Juan")).toBeTruthy();
});
```

#### scoping

```js
beforeAll(() => console.log("1 - beforeAll"));
afterAll(() => console.log("1 - afterAll"));
beforeEach(() => console.log("1 - beforeEach"));
afterEach(() => console.log("1 - afterEach"));
test("", () => console.log("1 - test"));
describe("Scoped / Nested block", () => {
  beforeAll(() => console.log("2 - beforeAll"));
  afterAll(() => console.log("2 - afterAll"));
  beforeEach(() => console.log("2 - beforeEach"));
  afterEach(() => console.log("2 - afterEach"));
  test("", () => console.log("2 - test"));
});

// 1 - beforeAll
// 1 - beforeEach
// 1 - test
// 1 - afterEach
// 2 - beforeAll
// 1 - beforeEach
// 2 - beforeEach
// 2 - test
// 2 - afterEach
// 1 - afterEach
// 2 - afterAll
// 1 - afterAll
```

## Mocking

### 생성

```js
const mock = jest.fn();
```

- `mockFn.getMockName()`: get mockname
- `mockFn.mockName(value)`: set mockname

- `mockFn.mock.calls`: mock 함수가 호출될 때 마다 인자들이 배열로 누적됨

  ```js
  [
    ["arg1", "arg2"],
    ["arg3", "arg4"]
  ];
  ```

- `mockFn.mock.results`: 호출의 결과 object들이 배열로 누적됨

  ```js
  [
    {
      isThrow: false,
      value: "result1"
    },
    {
      isThrow: true,
      value: {
        /* Error instance */
      }
    },
    {
      isThrow: false,
      value: "result2"
    }
  ];
  ```

- `mockFn.mock.instances`: `new jest.fn()`으로 인스턴스화된 mock function의 배열

- `mockFn.mockClear()`: mockFn.mock을 초기화
- `mockFn.mockReset()`: clear의 기능 + mock return value or implementation 삭제
- `mockFn.mockRestore()`: jest.spyOn으로 생성된 스파이객체를 초기상태(non-mocked implementation)로 되돌림

- `mockFn.mockImplementation(fn)`: mock에게 함수구현을 부여한다. 클래스 생성자로 사용될 수도 있다.

  ```js
  const mockFn = jest.fn().mockImplementation(scalar => 42 + scalar);
  // or: jest.fn(scalar => 42 + scalar);

  const a = mockFn(0);
  const b = mockFn(1);

  a === 42; // true
  b === 43; // true

  mockFn.mock.calls[0][0] === 0; // true
  mockFn.mock.calls[1][0] === 1; // true
  ```

- `mockFn.mockImplementationOnce(fn)`: 위와 같지만, mock이 한번 호출될 때 까지만 작동한다

- `mockFn.mockReturnThis()`: this를 반환하는 mock 함수를 생성한다

  ```js
  jest.fn(function() {
    return this;
  });
  ```

- `mockFn.mockReturnValue(value)`: 입력해둔 값을 반환하는 mock 함수를 생성한다
- `mockFn.mockReturnValueOnce(value)`: 위와 같지만 mock이 한번 호출될 때 까지만 작동한다.

- `mockFn.mockResolvedValue(value)`: Promise.resolve로 wrapping된 값을 반환한다
- `mockFn.mockResolvedValueOnce(value)`: 위의 행동을 한번만 수행한다

- `mockFn.mockRejectedValue(value)`: Promise.reject로 wrapping된 값을 반환한다
- `mockFn.mockRejectedValueOnce(value)`: 위의 행동을 한번만 수행한다

## Mocking & Testing 예제

```ts
import * as typeorm from "typeorm";
import TokenData from "../interfaces/tokenData.interface";
import UserDto from "../user/user.dto";
import AuthenticationService from "./authentication.service";

(typeorm as any).getRepository = jest.fn();

describe("AuthenticationService", () => {
  describe("when creating a cookie", () => {
    it("should return a string", () => {
      const tokenData: TokenData = { token: "", expiresIn: 1 };
      (typeorm as any).getRepository.mockReturnValue({});
      const authenticationService = new AuthenticationService();
      expect(typeof authenticationService.createCookie(tokenData)).toEqual(
        "string"
      );
    });
  });

  describe("when registering a user", () => {
    describe("if the email is already taken", () => {
      it("should throw an error", async () => {
        const userData: UserDto = {
          name: "Hong Gildong",
          email: "gdhong@chosun.com",
          password: "somepassword"
        };
        (typeorm as any).getRepository.mockReturnValue({
          findOne: () => Promise.resolve(userData)
        });
        const authenticationService = new AuthenticationService();
        await expect(
          authenticationService.register(userData)
        ).rejects.toMatchObject(new Error(userData.email));
      });
    });

    describe("if the email is not taken", () => {
      it("should not throw an error", async () => {
        const userData: UserDto = {
          name: "Hong Gildong",
          email: "gdhong@chosun.com",
          password: "somepassword"
        };
        (typeorm as any).getRepository.mockReturnValue({
          findOne: () => Promise.resolve(undefined),
          create: () => ({ ...userData, id: 0 }),
          save: () => Promise.resolve()
        });
        const authenticationService = new AuthenticationService();
        await expect(
          authenticationService.register(userData)
        ).resolves.toBeDefined();
      });
    });
  });
});
```
