# 웹브라우저 클릭하여 복사

```js
function clickToCopy(ref, name) {
  const range = document.createRange();
  range.selectNode(ref);

  window.getSelection().removeAllRanges();
  window.getSelection().addRange(range);

  try {
    document.execCommand("copy");
    console.log("복사됨");
  } catch (err) {
    console.error(err);
  } finally {
    window.getSelection().removeAllRanges();
  }
}
```
