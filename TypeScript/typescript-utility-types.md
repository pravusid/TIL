# TypeScript - Utility Types

- 문서: <https://www.typescriptlang.org/docs/handbook/utility-types.html>
- 구현: <https://github.com/microsoft/TypeScript/blob/main/src/lib/es5.d.ts>

TypeScript에서는 타입 변환을 편리하게 할 수 있는 유틸리티 타입을 global scope로 사용할 수 있다

## Refs

- <https://github.com/type-challenges/type-challenges>
- <https://github.com/sindresorhus/type-fest>

## `Partial<T>`

`T` 타입의 모든 프로퍼티를 optional로 설정하여 해당 타입의 subset을 표현하는 타입을 반환한다.

```ts
interface Todo {
  title: string;
  description: string;
}

function updateTodo(todo: Todo, fieldsToUpdate: Partial<Todo>) {
  return { ...todo, ...fieldsToUpdate };
}

const todo1 = {
  title: 'organize desk',
  description: 'clear clutter',
};

const todo2 = updateTodo(todo1, {
  description: 'throw out trash',
});
```

## `Required<T>`

`T` 타입의 모든 프로퍼티를 required(!optional)로 설정한 타입을 반환한다.

```ts
interface Props {
  a?: number;
  b?: string;
}

const obj: Props = { a: 5 }; // OK

const obj2: Required<Props> = { a: 5 }; // Error: property 'b' missing
```

## `Readonly<T>`

`T` 타입의 모든 프로퍼티를 `readonly`로 설정한(재할당 불가) 타입을 반환한다.

```ts
interface Todo {
  title: string;
}

const todo: Readonly<Todo> = {
  title: 'Delete inactive users',
};

todo.title = 'Hello'; // Error: cannot reassign a readonly property
```

런타임에 재할당이 되지 않아야 하는 경우를 나타내는데 유용하다. (`Object.freeze`)

```ts
function freeze<T>(obj: T): Readonly<T>;
```

## `Record<K,T>`

`T` 타입의 프로퍼티[세트] `K`를 가지는 타입을 반환한다.

`Record<K,T>`는 특정 타입의 프로퍼티를 다른 타입으로 매핑하기 위해 사용할 수 있다.

```ts
interface PageInfo {
  title: string;
}

type Page = 'home' | 'about' | 'contact';

const x: Record<Page, PageInfo> = {
  about: { title: 'about' },
  contact: { title: 'contact' },
  home: { title: 'home' },
};
```

## `Pick<T,K>`

`T` 타입으로부터 프로퍼티[세트] `K`를 선택한 타입(subset)을 반환한다.

```ts
interface Todo {
  title: string;
  description: string;
  completed: boolean;
}

type TodoPreview = Pick<Todo, 'title' | 'completed'>;

const todo: TodoPreview = {
  title: 'Clean room',
  completed: false,
};
```

`Todo`에서 `title`과 `description?`을 사용하는 타입의 예제는 다음과 같다.

```ts
type NoStatus = Pick<Todo, 'title'> & Pick<Partial<Todo>, 'description'>;

const picked1: NoStatus = { title: '포켓몬스터' }; // OK
const picked2: NoStatus = {
  title: '포켓몬스터',
  description: '피카츄는 내친구',
}; // OK
```

## `Omit<T, K>`

`Omit<T, K>` 타입은 포함되지 않았는데 `Pick<T, Exclude<keyof T, K>>`타입으로 사용할 수 있기 때문이다.

```ts
type Person = {
  name: string;
  age: number;
  location: string;
};

type RemainingKeys = Exclude<keyof Person, 'location'>;
type QuantumPerson = Pick<Person, RemainingKeys>;

// equivalent to
type QuantumPerson = {
  name: string;
  age: number;
};
```

TypeScript 3.5에서 자주 발생하는 이러한 유형의 작업을 처리하기 위한 Helper 타입이 추기되었다.

```ts
type Omit<T, K extends keyof any> = Pick<T, Exclude<keyof T, K>>;
```

위와 같이 별도로 `Omit` 타입을 정의할 필요 없이 `lib.d.ts`에 포함된 타입을 사용하면 된다.

컴파일러는 `Omit` 타입을 통해 제네릭에서 object rest destructuring 선언을 통해 생성된 타입을 표현한다.

## `Exclude<T,U>`

`T` 타입으로부터 프로퍼티[세트] `U`와 `T` 타입의 공통 프로퍼티들을 제외한 타입을 반환한다.
(`T` 타입에서 `U` 타입에 할당가능한 모든 프로퍼티를 제외함)

```ts
type T0 = Exclude<'a' | 'b' | 'c', 'a'>; // "b" | "c"
type T1 = Exclude<'a' | 'b' | 'c', 'a' | 'b'>; // "c"
type T2 = Exclude<string | number | (() => void), Function>; // string | number
```

> `Exclude` 타입은 정확히는 `Diff` 타입의 구현이다.
> `Diff`가 정의되어 있는 코드와 충돌을 회피하기 위해서 `Exclude`로 명명하였다.
> 또한 의미론적으로 더 나은 느낌을 전달한다.

## `Extract<T,U>`

`T` 타입으로부터 프로퍼티[세트] `U`와 `T` 타입의 공통 프로퍼티들을 추출한 타입을 반환한다.
(`T` 타입으로부터 `U` 타입에 할당할 수 있는 모든 프로퍼티를 추출)

```ts
type T0 = Extract<'a' | 'b' | 'c', 'a' | 'f'>; // "a"
type T1 = Extract<string | number | (() => void), Function>; // () => void
```

## `NonNullable<T>`

`T` 타입에서 `null` 과 `undefined`를 제외한 타입을 반환한다.

```ts
type T0 = NonNullable<string | number | undefined>; // string | number
type T1 = NonNullable<string[] | null | undefined>; // string[]
```

## `Parameters<T>`

함수 타입의 모든 파라미터 타입들을 추출한다.

모든 파라미터 타입들을 튜플 타입 형태로 제공한다.
대상이 함수가 아니면 `never` 타입을 반환한다.
여러 파라미터가 단일 타입으로 구성된다면 해당타입의 배열로 출력된다.

```ts
type A = Parameters<() => void>; // []
type B = Parameters<typeof Array.isArray>; // [any]
type C = Parameters<typeof parseInt>; // [string, (number | undefined)?]
type D = Parameters<typeof Math.max>; // number[]
```

## `ConstructorParameters<T>`

생성자 함수 타입의 모든 파라미터 타입들을 추출한다.

모든 파라미터 타입들을 튜플 타입 형태로 제공한다.
대상이 (생성자)함수가 아니면 `never` 타입을 반환한다.

```ts
type A = ConstructorParameters<ErrorConstructor>; // [(string | undefined)?]
type B = ConstructorParameters<FunctionConstructor>; // string[]
type C = ConstructorParameters<RegExpConstructor>; // [string, (string | undefined)?]
```

## `ReturnType<T>`

함수형 타입 `T`의 반환형을 반환한다.

```ts
type T0 = ReturnType<() => string>; // string
type T1 = ReturnType<(s: string) => void>; // void
type T2 = ReturnType<<T>() => T>; // {}
type T3 = ReturnType<<T extends U, U extends number[]>() => T>; // number[]
type T4 = ReturnType<typeof f1>; // { a: number, b: string }
type T5 = ReturnType<any>; // any
type T6 = ReturnType<never>; // any
type T7 = ReturnType<string>; // Error
type T8 = ReturnType<Function>; // Error
```

## `InstanceType<T>`

생성자 함수 타입 `T`의 인스턴스 타입을 반환한다.

```ts
class C {
  x = 0;
  y = 0;
}

type T0 = InstanceType<typeof C>; // C
type T1 = InstanceType<any>; // any
type T2 = InstanceType<never>; // any
type T3 = InstanceType<string>; // Error
type T4 = InstanceType<Function>; // Error
```

## `ThisParameterType<T>`

함수 타입의 `this` 파라미터 타입을, `this` 파라미터가 없는 경우 `unknown` 타입을 추출한다

```ts
function toHex(this: Number) {
  return this.toString(16);
}

function numberToString(n: ThisParameterType<typeof toHex>) {
  return toHex.apply(n);
}
```

## `OmitThisParameter<T>`

타입에서 `this` 파라미터를 제외한다

만약 명시적으로 선언된 `this` 파라미터가 없다면 결과는 그대로의 타입이다.
명시적으로 선언된 `this` 파라미터가 있으면 `this` 파리미터가 제외된 새 타입이 `T`로 부터 생성된다.

제너릭은 지워지고, 마지막 함수의 overload signature가 새로운 함수타입으로 전파된다.

```ts
function toHex(this: Number) {
  return this.toString(16);
}

const fiveToHex: OmitThisParameter<typeof toHex> = toHex.bind(5);

console.log(fiveToHex());
```

## `ThisType<T>`

이 유틸리티 타입은 변환된 타입을 반환하지 않는다. 대신 `this`의 문맥적 marker를 제공한다.

`ThisType<T>`를 사용하기 위해서는 컴파일러의 `--noImplicitThis` 옵션을 사용해야 한다.

```ts
type ObjectDescriptor<D, M> = {
  data?: D;
  methods?: M & ThisType<D & M>; // Type of 'this' in methods is D & M
};

function makeObject<D, M>(desc: ObjectDescriptor<D, M>): D & M {
  let data: object = desc.data || {};
  let methods: object = desc.methods || {};
  return { ...data, ...methods } as D & M;
}

let obj = makeObject({
  data: { x: 0, y: 0 },
  methods: {
    moveBy(dx: number, dy: number) {
      this.x += dx; // Strongly typed this
      this.y += dy; // Strongly typed this
    },
  },
});

obj.x = 10;
obj.y = 20;
obj.moveBy(5, 5);
```

위의 예제에서 `makeObject` 인수의 methods 객체는 `thisType <D & M>`을 포함하는 문맥 타입을 가지므로,
methods 객체 내 메소드에서의 this 타입은 `{ x: number, y: number } & { moveBy (dx: number, dy: number): number }`이다.
methods 프로퍼티의 타입이 추론 대상임과 동시에 메소드의 `this` 타입의 source임을 확인하라.

`ThisType<T>` 마커 인터페이스는 `lib.d.ts`에 선언 된 비어있는 인터페이스이다.
인터페이스는 객체 리터럴의 문맥적 타입에서 인식되는 것 이외에도 비어있는 인터페이스처럼 작동한다.

## Intrinsic String Manipulation Types

<https://www.typescriptlang.org/docs/handbook/2/template-literal-types.html#uppercasestringtype>

template string literals 타입의 유틸리티 타입으로 추가되었으며, 다음 네 가지가 있다

- `Uppercase<StringType>`
- `Lowercase<StringType>`
- `Capitalize<StringType>`
- `Uncapitalize<StringType>`

## ClassType

<https://www.typescriptlang.org/docs/handbook/2/generics.html#using-class-types-in-generics>

```ts
type Clazz<T> = {
  new (...args: any[]): T;
};

// 다음 타입도 Clazz<T> 타입과 같은 결과를 출력함
type Construct<T> = new (...args: any[]) => T;
```

ClassType 활용

```ts
class Test {
  constructor() {}
}

const a: Clazz<Test> = Test;
const b: Construct<Test> = Test;
```

## `NoInfer<T>`

> TypeScript 5.4에서 추가되었다

<https://devblogs.microsoft.com/typescript/announcing-typescript-5-4/#the-noinfer-utility-type>

문제가 되는 상황

- `C` 타입은 `colors` 변수의 추론된 타입만 사용하는 것이 의도이다
- 그러나 `C` 타입을 추론할 때 `defaultColor` 변수 타입도 추론에 포함된다

```ts
function createStreetLight<C extends string>(colors: C[], defaultColor?: C) {
  // ...
}

// Oops! This undesirable, but is allowed!
createStreetLight(['red', 'yellow', 'green'], 'blue');
```

NoInfer 사용하지 않고 해결

```ts
function createStreetLight<C extends string, D extends C>(colors: C[], defaultColor?: D) {}

createStreetLight(['red', 'yellow', 'green'], 'blue');
//                                            ~~~~~~
// error!
// Argument of type '"blue"' is not assignable to parameter of type '"red" | "yellow" | "green" | undefined'.
```

NoInfer 사용해서 해결

```ts
function createStreetLight<C extends string>(colors: C[], defaultColor?: NoInfer<C>) {
  // ...
}

createStreetLight(['red', 'yellow', 'green'], 'blue');
//                                            ~~~~~~
// error!
// Argument of type '"blue"' is not assignable to parameter of type '"red" | "yellow" | "green" | undefined'.
```

## 타입활용

### `infer NonFunctionProperties`

```ts
type NonFunctionPropertyNames<T> = {
  [K in keyof T]: T[K] extends Function ? never : K;
}[keyof T];

type NonFunctionProperties<T> = Pick<T, NonFunctionPropertyNames<T>>;
```

### `infer OnlyTypeProperties`

```ts
export type OnlyTypePropertyNames<T, O> = {
  [K in keyof T]: T[K] extends O ? K : never;
}[keyof T];
```

### `ExtractProps<T>`

<https://x.com/colinhacks/status/1818047762891506050?s=46&t=zQaFJ1iBc-ZF5GKv5CVZpQ>

```ts
// extract non-function properties from the prototype
type IsProp<T, K extends keyof T> = T[K] extends (...args: any[]) => any ? never : K;

type ExtractProps<T extends { prototype: unknown }> = {
  [k in keyof T['prototype'] as IsProp<T['prototype'], k>]: T['prototype'][k];
};
```

> `NonFunctionProperties<T>` 타입과는 다르게 다른 클래스를 상속하는 SubClass에서도 활용가능

```ts
class Player {
  name!: string;
  points!: number;

  constructor(props: ExtractProps<typeof Player>) {
    Object.assign(this, props);
  }
}

class FooPlayer extends Player {
  level!: number;

  constructor(props: ExtractProps<typeof FooPlayer>) {
    super(props);
  }
}
```

### `Mutable<T>`, `Immutable<T>`

```ts
// eslint-disable-next-line @typescript-eslint/ban-types
type Primitive = undefined | null | boolean | string | number | Function | symbol | bigint;

export type Mutable<T> = T extends Primitive
  ? T
  : T extends ReadonlyArray<infer U>
  ? Array<Mutable<U>>
  : T extends ReadonlyMap<infer K, infer V>
  ? Map<Mutable<K>, Mutable<V>>
  : T extends ReadonlySet<infer S>
  ? Set<Mutable<S>>
  : { -readonly [P in keyof T]: Mutable<T[P]> };

export type Immutable<T> = T extends Primitive
  ? T
  : T extends Array<infer U>
  ? ReadonlyArray<Immutable<U>>
  : T extends Map<infer K, infer V>
  ? ReadonlyMap<Immutable<K>, Immutable<V>>
  : T extends Set<infer S>
  ? ReadonlySet<Immutable<S>>
  : { readonly [P in keyof T]: Immutable<T[P]> };
```

### `Awaited<T>`

> TypeScript 4.5 버전에서 추가됨 <https://www.typescriptlang.org/docs/handbook/release-notes/typescript-4-5.html#the-awaited-type-and-promise-improvements>

```ts
export type Awaited<T> = T extends Promise<infer U> ? U : T;
```

### `Nullable<T>`

`Partial<T>`과 유사하지만 `undefined` 대신 `null`을 사용하는 경우

```ts
export type Nullable<T> = {
  [P in keyof T]: T[P] | null;
};
```

### `Optional<T>`

```ts
export type Optional<T> = T | undefined;

// 함수형 스타일로 결과를 처리할 수 있다
export const unwrap = <T>(optional: Optional<T>): T => {
  if (optional !== undefined) {
    return optional;
  }
  throw new Error('optional object is undefined');
};
```

### `DeepPartial<T>`

```ts
export type DeepPartial<T> = T extends object
  ? {
      [P in keyof T]?: DeepPartial<T[P]>;
    }
  : T;
```

### `Array.filter` with type-guard

<https://github.com/microsoft/TypeScript/issues/16655>

```ts
// type narrowing issue in TS Array.filter
Array.of(1, undefined, 2).filter(Boolean); // (number | undefined)[]

// 타입가드가 포함된 HOF 사용시
export function isTruthy<T>(obj: T | undefined | null): obj is T {
  return !!obj;
}

Array.of(1, undefined, 2).filter(isTruthy); // number[]
```

type-guard를 인스턴스 타입확인에 적용할 수도 있다

```ts
/**
 * @interface predicate
 */
export const typeOf =
  <In, Out extends In>(Proto: ClassType<Out>) =>
  (obj: In | undefined | null): obj is Out =>
    obj instanceof Proto;
```

type-guard를 객체의 프로퍼티에 적용할 수도 있다

```ts
/** 선택한 속성만 Required 적용 */
export type Convinced<T, R extends keyof T> = {
  [K in R]: T[K] extends infer I | null | undefined ? I : T[K];
} & Omit<T, R>;

export function hasTruthyValueIn<K extends keyof T, T>(key: K) {
  return (obj: T): obj is T & Convinced<T, K> => !!obj[key];
}

Array.of({ id: 1 }, { id: undefined }, { id: 2 }).filter(hasTruthyValueIn('id')); // { id: number }[]
```

#### `Array.filter` with overloading

<https://www.karltarvas.com/2021/03/11/typescript-array-filter-boolean.html>

`types/lib.es5.d.ts`

```ts
/** @link https://stackoverflow.com/a/51390763/1470607 */
type Falsy = false | 0 | '' | null | undefined;

interface Array<T> {
  /**
   * Returns the elements of an array that meet the condition specified in a callback function.
   * @param predicate A function that accepts up to three arguments. The filter method calls the predicate function one time for each element in the array.
   * @param thisArg An object to which the this keyword can refer in the predicate function. If thisArg is omitted, undefined is used as the this value.
   */
  filter<S extends T>(predicate: BooleanConstructor, thisArg?: any): Exclude<S, Falsy>[];
}
```

`tsconfig.json`

```json
{
  "include": ["lib.es5.d.ts"]
}
```

### WritableKeys

- <https://stackoverflow.com/a/52473108>
- <https://stackoverflow.com/a/49579497>
- <https://github.com/Microsoft/TypeScript/issues/27024#issuecomment-421529650>

```ts
type IfEquals<X, Y, A = X, B = never> = (<T>() => T extends X ? 1 : 2) extends <T>() => T extends Y ? 1 : 2 ? A : B;

type WritableKeys<T> = {
  [P in keyof T]-?: IfEquals<{ [Q in P]: T[P] }, { -readonly [Q in P]: T[P] }, P>;
}[keyof T];

type ReadonlyKeys<T> = {
  [P in keyof T]-?: IfEquals<{ [Q in P]: T[P] }, { -readonly [Q in P]: T[P] }, never, P>;
}[keyof T];
```

### Discriminated unions with helper

<https://www.typescriptlang.org/docs/handbook/2/narrowing.html#discriminated-unions>

> 식별자로 필터링한 결과에 타입적용

```ts
export type Id<I, T> = { _id: I } & T;

export const findById = <T extends Id<unknown, unknown>, I extends T['_id']>(list: T[], id: I & T['_id']) =>
  list.find((e): e is Extract<T, { _id: I }> => e._id === id);
```

### Union to Intersection

- <https://stackoverflow.com/questions/50374908/transform-union-type-to-intersection-type>
- <https://github.com/type-challenges/type-challenges/issues/122>
- <https://github.com/sindresorhus/type-fest/blob/main/source/union-to-intersection.d.ts>

```ts
type UnionToIntersection<U> = (U extends unknown ? (k: U) => void : never) extends (m: infer I) => void ? I & U : never;
```

### Tuple to Union

<https://ghaiklor.github.io/type-challenges-solutions/en/medium-tuple-to-union.html>

```ts
type Arr = ['1', '2', '3'];

const a: TupleToUnion<Arr>; // expected to be '1' | '2' | '3'

type TupleToUnion<T extends unknown[]> = T[number];
```

### Unique Array Elements

- <https://ja.nsommer.dk/articles/type-checked-unique-arrays.html>
- <https://stackoverflow.com/questions/57016728/is-there-a-way-to-define-type-for-array-with-unique-items-in-typescript>
- <https://stackoverflow.com/questions/71235152/exhaustive-list-of-keys-of-type>

> ~~컴파일타임에 tuple을 인자로 전달해야 하므로 유용하지는 않음~~
>
> TypeScript 5.0에 도입된 [const Type Parameters](https://github.com/microsoft/TypeScript/pull/51865)를 사용해보면 좋을듯

### Exhaustive Array Elements

> class, closure, recursive type 등을 활용

#### class

<https://github.com/gvergnaud/ts-pattern/blob/main/src/types/Match.ts>

#### closure

- [Array containing all options of type value in Typescript](https://stackoverflow.com/a/58508661)
- [enforce-that-an-array-is-exhaustive-over-a-union-type](https://stackoverflow.com/a/55266531)

```ts
type Furniture = 'chair' | 'table' | 'lamp' | 'ottoman';

type AtLeastOne<T> = [T, ...T[]];

const exhaustiveStringTuple =
  <T extends string>() =>
  <L extends AtLeastOne<T>>(
    ...x: L extends any ? (Exclude<T, L[number]> extends never ? L : Exclude<T, L[number]>[]) : never
  ) =>
    x;

const missingFurniture = exhaustiveStringTuple<Furniture>()('chair', 'table', 'lamp');
// error, Argument of type '"chair"' is not assignable to parameter of type '"ottoman"'

const extraFurniture = exhaustiveStringTuple<Furniture>()('chair', 'table', 'lamp', 'ottoman', 'bidet');
// error, "bidet" is not assignable to a parameter of type 'Furniture'

const furniture = exhaustiveStringTuple<Furniture>()('chair', 'table', 'lamp', 'ottoman');
// okay
```

#### typing

[how-to-exhaustive-check-the-elements-in-an-array-in-typescript](https://stackoverflow.com/a/74354311)

```ts
type Country = 'uk' | 'france' | 'india';

type MapOfKeysOf<U extends string> = {
  [key in U]: MapOfKeysOf<Exclude<U, key>>;
};

type ExhaustiveArrayOfObjects<
  Keys extends { [key: string]: {} },
  T = {},
  KeyName extends string = 'value'
> = {} extends Keys
  ? []
  : {
      [key in keyof Keys]: [T & Record<KeyName, key>, ...ExhaustiveArrayOfObjects<Keys[key], T, KeyName>];
    }[keyof Keys];

type Option = {
  label: string;
  id: Country;
};

const data: ExhaustiveArrayOfObjects<MapOfKeysOf<Country>, Option, 'id'> = [
  {
    label: 'United Kingdom',
    id: 'uk',
  },
  {
    label: 'France',
    id: 'france',
  },
  {
    label: 'India',
    id: 'india',
  },
];
```

### Enum Value

```ts
/**
 * @link https://stackoverflow.com/questions/72050271/check-if-value-exists-in-string-enum-in-typescript
 */
export function isValueInEnum<E extends string>(strEnum: Record<string, E>) {
  const enumValues = Object.values(strEnum) as string[];
  return (value: string | null | undefined): value is E => !!value && enumValues.includes(value);
}
```

### Prettify

<https://www.totaltypescript.com/concepts/the-prettify-helper>

```ts
type Prettify<T> = {
  [K in keyof T]: T[K];
} & {};
```

Intersected Object Type (`{ a: string } & { b: string }`)을 정리해서 보여줌

> `{ a: string; b: string }`

### 자동완성가능한 literal union

<https://github.com/microsoft/TypeScript/issues/29729>

```ts
type LiteralUnion<T extends U, U = string> = T | (U & {});
```

### Omit(Pick) only optional keys

- <https://github.com/type-challenges/type-challenges/blob/main/questions/00057-hard-get-required/README.md>
- <https://github.com/type-challenges/type-challenges/blob/main/questions/00059-hard-get-optional/README.md>

```ts
type RequiredFieldsOnly<T> = {
  [K in keyof T as T[K] extends Required<T>[K] ? K : never]: T[K];
};

type PartialFieldsOnly<T> = {
  [K in keyof T as T[K] extends Required<T>[K] ? never : K]: T[K];
};
```
