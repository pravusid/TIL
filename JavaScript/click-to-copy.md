# 웹브라우저 클릭하여 복사

case1

```js
function clickToCopy(ref) {
  const range = document.createRange();
  range.selectNode(ref);

  window.getSelection().removeAllRanges();
  window.getSelection().addRange(range);

  try {
    document.execCommand('copy');
    console.log('복사됨');
  } catch (err) {
    console.error(err);
  } finally {
    window.getSelection().removeAllRanges();
  }
}
```

case2

```js
function clickToCopy(ref) {
  const temp = document.createElement('input');
  temp.value = ref.target?.value;
  document.body.appendChild(temp);
  temp.select();
  document.execCommand('copy');
  document.body.removeChild(temp);
}
```
