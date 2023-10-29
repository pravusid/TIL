# Chrome DevTools

## console utilities

<https://developer.chrome.com/docs/devtools/console/utilities/>

- `$_`: 마지막으로 평가된 표현식의 반환 값
- `$0`~`$4`: Elements 탭에서 -n번째 선택한 Dom element
- `$(selector [, startNode])`: `document.querySelector(selector)`의 alias
- `$$(selector [, startNode])`: `document.querySelectorAll('li')`의 alias
- `$x(selector)`: 해당 요소의 xpath를 반환
- `copy(...val)`: OS 클립보드로 값을 복사한다
