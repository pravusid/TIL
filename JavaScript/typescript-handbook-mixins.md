# TypeScript HandBook: Mixins

## 소개

전통적인 객체지향 계층구조와 함께 재사용 가능한 구성 요소에서 클래스를 만드는 또 다른 보편적인 방법은 클래스의 단순한 부분을 결합하여 클래스를 만드는 것이다.
스칼라(scala)와 같은 언어에서 mixins 또는 traits에 대한 이이디어에 익숙할 수도 있으며 이와 같은 패턴은 자바스크립트 커뮤니티에서도 인기를 얻고있다.

## Mixin 샘플

아래에 mixins로 작동할 두 클래스가 있다.
각각의 클래스는 하나의 활동 혹은 기능에 초점을 맞추고 있다.

```ts
// Disposable Mixin
class Disposable {
  isDisposed: boolean;
  dispose() {
    this.isDisposed = true;
  }
}

// Activatable Mixin
class Activatable {
  isActive: boolean;
  activate() {
    this.isActive = true;
  }
  deactivate() {
    this.isActive = false;
  }
}
```

다음으로 두 개의 클래스를 조합하여 사용할 클래스를 생성한다.

우선 클래스를 대상으로 `extends` 대신 `implements`를 사용하고 있음을 알 수 있다.
이는 클래스를 인터페이스로 취급하여 구현이 아닌 `Disposable`, `Activatable` 타입만을 사용하는 것이다.
즉, 사용하는 클래스에서 구현을 제공해야 한다는 것을 의미하지만, 믹스인을 사용하여 이를 피하고자 한다.

요구조건을 충족하기 위해 믹스인에서 사용할 예정인 멤버에 대해 stand-in 프로퍼티와 타입을 만든다.
이는 컴파일러가 런타임에 이러한 멤버를 사용할 수 있음을 확인하도록 한다.

```ts
class SmartObject implements Disposable, Activatable {
  constructor() {
    setInterval(() => console.log(this.isActive + " : " + this.isDisposed), 500);
  }

  interact() {
    this.activate();
  }

  // Disposable
  isDisposed: boolean = false;
  dispose: () => void;

  // Activatable
  isActive: boolean = false;
  activate: () => void;
  deactivate: () => void;
}
```

마지막으로, 두 클래스를 mixin하여 전체 구현을 만든다.

```ts
applyMixins(SmartObject, [Disposable, Activatable]);

let smartObj = new SmartObject();
setTimeout(() => smartObj.interact(), 1000);
```

일반적으로 mixins를 수행할 헬퍼함수를 다음과 같이 정의할 수 있다.
이는 각각의 믹스인 프로퍼티를 거쳐 믹스인의 타겟으로 복사하고 stand-in 프로퍼티를 구현체로 채운다.

```ts
////////////////////////////////////////
// In your runtime library somewhere
////////////////////////////////////////

function applyMixins(derivedCtor: any, baseCtors: any[]) {
  baseCtors.forEach(baseCtor => {
    Object.getOwnPropertyNames(baseCtor.prototype).forEach(name => {
      Object.defineProperty(derivedCtor.prototype, name, Object.getOwnPropertyDescriptor(baseCtor.prototype, name));
    });
  });
}
```
