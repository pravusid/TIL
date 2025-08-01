# TailwindCSS

- <https://tailwindcss.com/>
- <https://github.com/aniftyco/awesome-tailwindcss>

## Origin

- <https://adamwathan.me/css-utility-classes-and-separation-of-concerns/>
- <https://news.ycombinator.com/item?id=21553496>
- [The evolution of scalable CSS](https://frontendmastery.com/posts/the-evolution-of-scalable-css/)
  - [(번역) 확장 가능한 CSS의 진화](https://ykss.netlify.app/translation/the_evolution_of_scalable_css/)

## Refs

- [Grouping variants together](https://github.com/tailwindlabs/tailwindcss/discussions/8337)
- <https://darkghosthunter.medium.com/tailwind-the-base-the-components-and-the-utilities-a81137c52534>
- [The Pros and Cons of TailwindCSS](https://webartisan.info/the-pros-and-cons-of-tailwindcss)
  - [(번역) Tailwind CSS의 장점과 단점](https://ykss.netlify.app/translation/the_pros_and_cons_of_tailwindcss/)
- <https://fe-developers.kakaoent.com/2022/220303-tailwind-tips/>
- <https://fe-developers.kakaoent.com/2022/221013-tailwind-and-design-system/>
- [Tailwind CSS에서 혼란을 방지하기 위한 5가지 모범 사례](https://velog.io/@lky5697/5-best-practices-for-preventing-chaos-in-tailwind-css)

## Libraries

### Components

- <https://github.com/tailwindlabs/tailwindcss-forms>
- <https://github.com/shadcn-ui/ui>
- <https://github.com/saadeghi/daisyui>
- <https://github.com/mdbootstrap/TW-Elements>

### Plugins & Tools

- <https://github.com/ben-rogerson/twin.macro>
- <https://github.com/dcastil/tailwind-merge>
- <https://github.com/barvian/fluid-tailwind>
- <https://github.com/lukeed/clsx>
- <https://github.com/joe-bell/cva>
- <https://github.com/nextui-org/tailwind-variants> (cva + tailwind-merge)
- <https://github.com/gregberge/twc>
- <https://github.com/tailwindlabs/prettier-plugin-tailwindcss>

## install

<https://tailwindcss.com/docs/installation/framework-guides>

## tailwind directive

<https://tailwindcss.com/docs/functions-and-directives#tailwind>

```css
/**
 * This injects Tailwind's base styles and any base styles registered by
 * plugins.
 */
@tailwind base;

/**
 * This injects Tailwind's component classes and any component classes
 * registered by plugins.
 */
@tailwind components;

/**
 * This injects Tailwind's utility classes and any utility classes registered
 * by plugins.
 */
@tailwind utilities;

/**
 * Use this directive to control where Tailwind injects the hover, focus,
 * responsive, dark mode, and other variants of each class.
 *
 * If omitted, Tailwind will append these classes to the very end of
 * your stylesheet by default.
 */
@tailwind variants;
```

### base

<https://tailwindcss.com/docs/preflight>

> 브라우저간의 불일치를 완화하고, 직접작성한 디자인 시스템을 쉽게 적용하기 위해서 tailwind에서 정의한 기본 스타일을 말한다.
> 보통은 브라우저 기본 스타일을 제거하는 형태로 정의되어 있고, tailwind 설정에서 비활성화 할 수도 있다.

## VisualStudioCode Extension

> 자동완성 기능을 다양한 상황에서 사용하려면 별도의 설정이 필요하다

- <https://github.com/tailwindlabs/tailwindcss/issues/7553>
- <https://github.com/tailwindlabs/tailwindcss/discussions/7554>
- <https://github.com/lukeed/clsx?tab=readme-ov-file#tailwind-support>
- <https://cva.style/docs/getting-started/installation#intellisense>
- <https://github.com/ben-rogerson/twin.macro/discussions/227>
- <https://github.com/paolotiu/tailwind-intellisense-regex-list>
