# TypeScript Inferred Type Predicates

- <https://github.com/microsoft/TypeScript/pull/57465>
- <https://github.com/microsoft/TypeScript/pull/57847>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-5-5-rc/#inferred-type-predicates>

## abstract

```ts
function f(p: trueType, p2: T2, ...) {
  // ...
  if (expr) {
    p1;  // trueType
  } else {
    p1;  // never?
}
```

## check for truthiness

> We don't infer a type guard here if you check for truthiness, only if you check for non-nullishness:

```ts
const numsTruthy = [0, 1, 2, null, 3].filter((x) => !!x);
//    ^? const numsTruthy: (number | null)[]
const numsNonNull = [0, 1, 2, null, 3].filter((x) => x !== null);
//    ^? const numsNonNull: number[]
```

> This is because of the false case: if the truthiness test returns false, then x could be 0.
> Until TypeScript can represent "numbers other than 0" or it has a way to return distinct type predicates for the true and false cases,
> there's nothing that can be inferred from the truthiness test here.
>
> If you're working with object types, on the other hand, there is no footgun and a truthiness test will infer a predicate:

```ts
const datesTruthy = [new Date(), null, new Date(), null].filter((d) => !!d);
//    ^? const datesTruthy: Date[]
```

> This provides a tangible incentive to do non-null checks instead of truthiness checks in the cases where you should be doing that anyway,
> so I call this a win. Notably the example in the original issue tests for truthiness rather than non-null.
