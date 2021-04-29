# Tsyringe

<https://github.com/microsoft/tsyringe>

> A lightweight dependency injection container for TypeScript/JavaScript for constructor injection.

## Getting Started

```ts
import 'reflect-metadata';
import { autoInjectable, container, injectable, singleton } from 'tsyringe';

// @scoped(Lifecycle.ContainerScoped)
// @scoped(Lifecycle.ResolutionScoped)
// @scoped(Lifecycle.Singleton) -> @deprecated == @singleton()
// @scoped(Lifecycle.Transient) -> @deprecated == @injectable()

// @injectable 에서 new Ctor()로 호출된다
// Lifecycle.Transient 상태와 같음
class FooHello {
  constructor() {
    console.log('foo hello generated');
  }

  hello() {
    console.log('hello');
  }
}

@singleton()
class FooWorld {
  constructor() {
    console.log('foo world generated');
  }

  hello() {
    console.log('world');
  }
}

@injectable()
class Bar {
  constructor(public fooHello: FooHello, public fooWorld: FooWorld) {}
}

// 생성자를 대체하는 것이므로 new Ctor() 대신 호출됨
@autoInjectable()
class FooBar {
  constructor(public fooHello?: FooHello, public fooWorld?: FooWorld) {}
}

const bar1 = container.resolve(Bar);
const bar2 = container.resolve(Bar);
const fooBar = new FooBar();
const fooBarFromContainer = container.resolve(FooBar);
// foo hello generated
// foo world generated
// foo hello generated
// foo hello generated
// foo hello generated

console.log('identical Bar', bar1 === bar2); // identical Bar false
console.log('identical FooHello', bar1.fooHello === bar2.fooHello); // identical FooHello false
console.log('identical FooWorld', bar1.fooWorld === bar2.fooWorld); // identical FooWorld true

// 컨테이너 registred 상태인 것은 라이프사이클내 동일 인스턴스를 관리하는 경우만
console.log('FooHello registered', container.isRegistered(FooHello)); // FooHello registered false
console.log('FooWorld registered', container.isRegistered(FooWorld)); // FooWorld registered true
console.log('Bar registered', container.isRegistered(Bar)); // Bar registered false
console.log('FooBar registered', container.isRegistered(FooBar)); // FooBar registered false

console.log('fooBar from ctor', fooBar); // fooBar from ctor FooBar { fooHello: FooHello {}, fooWorld: FooWorld {} }
fooBar.fooHello?.hello(); // hello
fooBar.fooWorld?.hello(); // world

console.log('fooBar from container', fooBarFromContainer); // fooBar from container FooBar { fooHello: FooHello {}, fooWorld: FooWorld {} }
fooBarFromContainer.fooHello?.hello(); // hello
fooBarFromContainer.fooWorld?.hello(); // world
```

## custom decorator

### `singleton` 소스 참조

<https://github.com/microsoft/tsyringe/blob/master/src/decorators/singleton.ts>

```ts
import constructor from '../types/constructor';
import injectable from './injectable';
import { instance as globalContainer } from '../dependency-container';

function singleton<T>(): (target: constructor<T>) => void {
  return function (target: constructor<T>): void {
    injectable()(target);
    globalContainer.registerSingleton(target);
  };
}

export default singleton;
```

<https://github.com/microsoft/tsyringe/blob/master/src/dependency-container.ts>

```ts
export const instance: DependencyContainer = new InternalDependencyContainer();
export default instance;
```

<https://github.com/microsoft/tsyringe/blob/master/src/index.ts>

```ts
export { instance as container } from './dependency-container';
```

### `controller`

> Ts.ED import 코드 참조
> <https://github.com/TypedProject/tsed/blob/master/packages/di/src/utils/importFiles.ts>

```ts
import * as globby from 'globby';

const files = globby.sync('dist/**/api/**/*-api.js');

for (const file of files) {
  await import(`${process.cwd()}/${file}`);
}
```

> 호출한 `container`는 tsyringe 내부 `globalContainer`와 동일

```ts
import { container, injectable } from 'tsyringe';

type Ctor<T> = new (...args: any[]) => T;

export function controller<T>(): (target: Ctor<T>) => void {
  return function (target: Ctor<T>): void {
    injectable()(target);
    container.registerSingleton('controller', target);
  };
}
```
