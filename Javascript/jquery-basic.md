# JQuery

## selector

- $(selector) : css selector를 적용함
- 태그명
- ID
- Class
- 내장객체

제이쿼리의 선택자는 엘리먼트셋을 반환한다. 따라서 결과값이 하나인 경우 0번째 배열 인자를 선택하여 사용할 수도 있다.

`$(this)[0] === this`, `$("#myDiv")[0] === document.getElementById("myDiv")`

## Events

Mouse Events | Keyboard Events | Form Events | Document/Window Events
-------------|-----------------|-------------|-----------------------
click | keypress | submit | load
dblclick | keydown | change | resize
mouseenter | keyup | focus | scroll
mouseleave |   | blur | unload

### 이벤트 처리

e.preventDefault();

### The on() Method

The on() method attaches one or more event handlers for the selected elements.

### Dropdown collapse 처리 예시

```javascript
var login = function(e) {
  e.stopPropagation(e);
  $.ajax({
    type : 'POST',
    url : '/login',
    data : {
      'email' : $('#login-email').val(),
      'pwd' : $('#login-password').val()
    },
    success : function(resp) {
      if (resp.result=="no") {
        $('#login-alert').slideDown(500).delay(2000).slideUp(500);
      } else {
        alert("로그인성공");
        $('#login-form').dropdown("toggle");
      }
    }
  });
};

$('#login-form').bind('click', function(e) {
  e.stopPropagation()
});

$('#login-alert').hide();

$('#login-btn').click(function(e) {
  login(e);
});

$('#login-email, #login-password').keypress(function(e) {
  if (e.keyCode == '13') {
    login(e)
  }
});
```

## JQuery AJAX

```javascript
$ajax({
  type: "get" or "post" or "put" or "delete" ,
  contentType: "multipart/form-data" or "application/json; charset=UTF-8",
  url: "",
  data: {id:value, pw:value} or JSON.stringify(json),
  dataType: "json", // 서버에서 돌아오는 데이터(response)가 어떤 형식일지
  success: function(data, status) {
    /* 결과값으로 기능구현 */
  },
  error: function(xhr, status) {
    /* 결과값으로 기능구현 */
  }
});
```
