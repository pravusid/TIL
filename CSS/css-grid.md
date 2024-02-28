# CSS Grid

- <https://developer.mozilla.org/ko/docs/Web/CSS/CSS_grid_layout/Basic_concepts_of_grid_layout>
- <https://yamoo9.gitbook.io/css-grid/>
- <https://www.joshwcomeau.com/css/interactive-guide-to-grid/>

## grid 선언

> 그리드 컨테이너는 요소에 `display: grid` 또는 `display: inline-grid`를 선언하여 만듭니다.
> 이렇게 하면 해당 요소 바로 밑에 있는 모든 자식 요소는 그리드 아이템이 됩니다.

## grid 트랙

- `columns`, `rows` 고정값 또는 비율을 정의할 수 있다
- `fr` 단위는 컨테이너의 남은 공간을 비율로 지정한 것이다

```css
.wrapper {
  display: grid;
  grid-template-columns: 1fr 2fr 1fr;
  grid-template-rows: 1fr 1fr 1fr;
}
```

- `fr` 선언을 반복하는 대신 `repeat` 키워드를 사용할 수 있다
- `1fr 1fr 1fr == repeat(3, 1fr)`

### grid 크기 자동설정

- grid에서는 행, 열 각각의 크기(비율)를 정해진 횟수만큼 선언하게 된다
- 선언한 크기를 넘어서는 그리드 요소가 발생했을 때 적용될 크기를 지정할 수 있다
- 이 때 자동설정 크기는 고정 값 대신 `minmax`를 사용할 수 있다

```css
.wrapper {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-auto-rows: 200px;
  grid-auto-rows: minmax(100px, auto); /* 최소 100px, 최대크기 자동설정(콘텐츠 크기에 따라) */
}
```

## grid 라인
