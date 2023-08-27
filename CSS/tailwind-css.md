# TailwindCSS

<https://tailwindcss.com/>

## Refs

- <https://darkghosthunter.medium.com/tailwind-the-base-the-components-and-the-utilities-a81137c52534>
- <https://ykss.netlify.app/translation/the_pros_and_cons_of_tailwindcss/>
- <https://fe-developers.kakaoent.com/2022/220303-tailwind-tips/>
- <https://fe-developers.kakaoent.com/2022/221013-tailwind-and-design-system/>

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
