# HTML Layout

- 문서 구조관련 HTML 태그는 [[html-semantic]] 참고
- 레이아웃을 위한 주요 CSS 기능은 [[css-position]], [[css-flexbox]], [[css-grid]] 참고

## 헤더, 본문형

- 기본 전략은 다음과 같다
- `main` 태그는 `display: flex`, `max-width: 64rem`, `width: 100%` 적용
- `main > nav` 태그는 `고정 너비`, `flex-shrink: 0`
- `main > article` 태그는 `flex: 1 1 auto`, `min-width: 0`

```html
<body>
  <header>
    <h1>Three words</h1>
    <nav>
      <a>one word</a>
      <a>one word</a>
      <a>one word</a>
      <a>one word</a>
    </nav>
  </header>
  <main>
    <nav>
      <a>one word</a>
      <a>one word</a>
    </nav>
    <article>
      <h1>five words</h1>
      <p>forty-six words</p>
      <p>forty-four words</p>
    </article>
  </main>
  <footer>
    <p>five words</p>
  </footer>
</body>
```
