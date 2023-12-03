# Semantic HTML

- <https://web.dev/learn/html/semantic-html/>
- <https://web.dev/learn/html/headings-and-sections?hl=ko>
- <https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Document_and_website_structure>
- <https://developer.mozilla.org/ko/docs/Web/HTML/Element>

## 검증

브라우저 개발자도구의 접근성 패널에서 구조를 확인해볼 수 있다

## section, article

문제는 `section` 태그와 `article` 태그의 사용방식이다.
이에 대해서는 [`<section>` 버리고 HTML5 `<article>` 써야 하는 이유](https://webactually.com/2020/03/03/<section>을-버리고-HTML5-<article>을-써야-하는-이유/) 같은 의견이 있다.

## web.dev 예시

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
    <header>
      <h1>five words</h1>
    </header>
    <section>
      <h2>three words</h2>
      <p>forty-six words</p>
      <p>forty-four words</p>
    </section>
    <section>
      <h2>seven words</h2>
      <p>sixty-eight words</p>
      <p>forty-four words</p>
    </section>
  </main>
  <footer>
    <p>five words</p>
  </footer>
</body>
```

## next.js 공식문서 예시

구조를 단순화한 예시이다.
시맨틱 태그는 정확한 위치에 사용하고 레이아웃을 위한 태그는 `div`를 사용한 것을 알 수 있다.

```html
<header>
  <nav>
    <!-- 내용 -->
  </nav>
</header>
<main>
  <div> <!-- container layout -->
    <div> <!-- sidebar layout -->
      <!-- ... -->
      <nav>
        <!-- sidebar links -->
      </nav>
    </div>
    <article>
      <!-- 본문 -->
    </article>
    <nav>
      <!-- TOC -->
    </nav>
  </div>
</main>
```
