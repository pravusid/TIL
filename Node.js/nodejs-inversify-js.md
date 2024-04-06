# InversifyJS

> A powerful and lightweight inversion of control container for JavaScript & Node.js apps powered by TypeScript.

<https://github.com/inversify/InversifyJS>

## Docs

- <https://github.com/inversify/InversifyJS/blob/master/wiki/readme.md>
- <https://github.com/inversify/InversifyJS/blob/master/wiki/recipes.md>

## Utils

### <https://github.com/inversify/inversify-binding-decorators>

> An utility that allows developers to declare InversifyJS bindings using ES2016 decorators

```ts
import { injectable, Container } from 'inversify';
import 'reflect-metadata';

@injectable()
class Katana implements Weapon {
  public hit() {
    return 'cut!';
  }
}

@injectable()
class Shuriken implements ThrowableWeapon {
  public throw() {
    return 'hit!';
  }
}

let container = new Container();
container.bind<Katana>('Katana').to(Katana);
container.bind<Shuriken>('Shuriken').to(Shuriken);
```

상단의 코드를 제공된 데코레이터를 사용하여 동일하게 표현할 수 있다

```ts
import { injectable, Container } from 'inversify';
import { provide, buildProviderModule } from 'inversify-binding-decorators';
import 'reflect-metadata';

@provide(Katana)
class Katana implements Weapon {
  public hit() {
    return 'cut!';
  }
}

@provide(Shuriken)
class Shuriken implements ThrowableWeapon {
  public throw() {
    return 'hit!';
  }
}

let container = new Container();
// Reflects all decorators provided by this package and packages them into
// a module to be loaded by the container
container.load(buildProviderModule());
```

### <https://github.com/inversify/inversify-inject-decorators>

> The decorators included in this library will allow you to lazy-inject properties even when the instances of a class are not created by InversifyJS.
> This library allows you to integrate InversifyJS with any library or framework that takes control over the creation of instances of a given class.

### <https://github.com/inversify/inversify-logger-middleware>

```ts
let module = new ContainerModule((bind: inversify.interfaces.Bind) => {
  bind<Weapon>('Weapon').to(Katana).whenInjectedInto(Samurai);
  bind<Weapon>('Weapon').to(Shuriken).whenInjectedInto(Ninja);
  bind<Warrior>('Warrior').to(Samurai).whenTargetTagged('canSneak', false);
  bind<Warrior>('Warrior').to(Ninja).whenTargetTagged('canSneak', true);
});
```

> This middleware will display the InversifyJS resolution plan in console in the following format.

```text
// container.getTagged<Warrior>("Warrior", "canSneak", true);

SUCCESS: 0.41 ms.
    └── Request : 0
        └── serviceIdentifier : Warrior
        └── bindings
            └── Binding<Warrior> : 0
                └── type : Instance
                └── implementationType : Ninja
                └── scope : Transient
        └── target
            └── serviceIdentifier : Warrior
            └── name : undefined
            └── metadata
                └── Metadata : 0
                    └── key : canSneak
                    └── value : true
        └── childRequests
            └── Request : 0
                └── serviceIdentifier : Weapon
                └── bindings
                    └── Binding<Weapon> : 0
                        └── type : Instance
                        └── implementationType : Shuriken
                        └── scope : Transient
                └── target
                    └── serviceIdentifier : Weapon
                    └── name : shuriken
                    └── metadata
                        └── Metadata : 0
                            └── key : name
                            └── value : shuriken
                        └── Metadata : 1
                            └── key : inject
                            └── value : Weapon
```

## Examples

- <https://github.com/inversify/inversify-basic-example>
- <https://github.com/inversify/inversify-express-example>
