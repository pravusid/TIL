# Expose Getters in TypeScript

```ts
toJSON() {
  const proto = Object.getPrototypeOf(this);
  const jsonObj: any = Object.assign({}, this);

  Object.entries(Object.getOwnPropertyDescriptors(proto))
    .filter(([key, descriptor]) => typeof descriptor.get === 'function')
    .forEach(([key, descriptor]) => {
      if (descriptor && key[0] !== '_') {
        jsonObj[key] = (this as any)[key];
      }
    });

  return jsonObj;
}
```

기존 `properties`를 출력하지 않고 `getters`만 출력하려면

```ts
// before
const jsonObj: any = Object.assign({}, this);

// after
const jsonObj: any = {};
```
