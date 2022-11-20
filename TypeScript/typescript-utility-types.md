# TypeScript - Utility Types

- 문서: <https://www.typescriptlang.org/docs/handbook/utility-types.html>
- 구현: <https://github.com/microsoft/TypeScript/blob/main/lib/lib.es5.d.ts#L1530>

TypeScript에서는 타입 변환을 편리하게 할 수 있는 유틸리티 타입을 global scope로 사용할 수 있다.

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

## 타입활용

<https://github.com/type-challenges/type-challenges>

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
  // eslint-disable-next-line @typescript-eslint/ban-types
  [K in keyof T]: T[K] extends O ? K : never;
}[keyof T];
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

### Discriminated unions with Helper

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

```ts
type UnionToIntersection<U> = (U extends unknown ? (k: U) => void : never) extends (k: infer I) => void ? I : never;
```
