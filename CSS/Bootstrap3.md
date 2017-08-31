# 부트스트랩(Bootstrap)

## 미디어 쿼리(CSS3)

- 화면크기에 따라 출력을 다르게 한다(반응형 웹)
- css에서 @media 지시어 사용
  ```css
  /* (화면출력/인쇄)에서만 적용 */
  @media screen/print {
    div {
      width: 800px;
      height: 800px;
    }
  }

  /* width가 800보다 작은경우에만 적용 */
  @media(max-width:800px) {
    div{
      width: 400px;
      height: 400px;
    }
  }

  /* max와 min 함께 사용 */
  @media(min-width:800px) and (max-width:1200px) {
    div{
      width: 1000px;
      height: 1000px;
    }
  }
  ```

## Respond.js, html5shiv

- 미디어쿼리와 html5는 구형 웹브라우저에서 지원되지 않기 때문에 추가 라이브러리를 이용해서 해당 기능을 구현한다.
- Respond.js : 미디어쿼리 기능
- html5shiv : html5 미지원 브라우저에서 html5 사용

## Grid, bootstrap CSS

- grid class 계층 (div, table등에 적용)
  - .container(고정) / .container-fluid(유동)
  - .row : 한 줄당 col 12칸
  - .col-크기-개수
    - .col-xs-* < 768
    - .col-sm-* < 992
    - .col-md-* < 1200
    - .col-lg-*
    - .col 내부에 row 생성 가능 -> 내부의 전체 col수는 다시 12

- responsive util
  - visible-mode : .visible-md -> md일때만 나타남, 다른크기에도 적용가능
  - hidden-mode : .hidden-md -> md일때 사라짐, 다른크기에도 적용가능
  - 컬럼 영역의 높이가 다른경우 레이아웃이 어긋남 : div 기본속성이 float left
    - clearfix적용 : 비어있는 div 삽입후 .clearfix 적용
  - 크기를 차지하는 빈 컬럼 삽입 : .col-크기-offset-개수
  - 크기를 차지하지 않는 빈 컬럼 삽입(크기는 개행 여부의 차이)
    - 현재 위치에서 이동
    - .col-크기-push-개수 : 들여쓰기
    - .col-크기-pull-개수 : 내어쓰기

### 반응형 적용하지 않기

- viewport 메타태그 삭제
- .container의 크기 재정의 : width: 970px !important;
- navbars를 사용하는 경우 생략
- .col-xs-*만 사용

### 스타일 설정

- Font 설정
  - Heading tag : h1(36px) ~ (6px간격) ~ h6(12px)
  - Body : 14px
  - (class) align : .text-left .text-center .text-right

- Table 설정 (class) .table .table-border .table-hover

## 컴포넌트

- .active : 활성화 (컴포넌트가 다수이면 하나 적용해놓아야 함)
- disabled : 입력형 태그에 적용가능 한 사용불가 명령

### Form 설정

- form.form-horizontal
  - div.form-group
    - `<label for="이름" class="control-label col-xs-2">`
    - input.form-control .col-xs-10
- form.form-inline
  - label.sr-only
- p.form-control-static : 텍스트 수직중앙정렬

### button class

- 버튼 모양
  - .btn .btn-default
  - .btn .btn-primary
  - .btn .btn-success
  - .btn .btn-info
  - .btn .btn-warning
  - .btn .btn-danger
- 버튼 크기
  - .btn-xs
  - .btn sm
  - .btn-lg
  - .btn-block : 부모태그의 크기에 버튼크기를 맞춤
- 버튼 기능
  ```javascript
  // data loading
  $("#btn").click(function() {
    $(this).button("loading");
    setTimeout(function() {
      $("#btn").button("reset");
    }, 2000);
  });
  ```

### 이미지

- `<img data-src="holder.js/가로x세로" alt="설명" class="스타일">`
- 스타일 종류 : .img-rounded, .img-circle, .img-thumbnail

### glyphicons : 클래스에서 단독지정(다른 클래스와 함께사용X)

- `<span class="glyphicon glyphicon-종류"></span>`

### dropdown

- div.dropdown/.dropup .text-right(정렬)
  - button.btn .btn-primary .dropdown-toggle data-toggle="dropdown"
    - span.caret
  - ul.dropdown-menu .dropdown-menu-right(정렬)
    - `<li class="dropdown-header></li>`
    - `<li> <a> 내용 </a> </li>`
    - `<li class="divider"></li>` : 구분선
- dropdown 기능
  ```javascript
  $("#버튼id").click(function(e) {
    $(".dropdown-toggle").dropdown("toggle");
    return false;
  });
  ```
  - event
    ```javascript
      $("#menu1").on("show/shown.bs.dropdown", function(e) {
      });
    ```

### button group

- div.btn-group
  - `<button/a class="btn btn-primary">` 버튼 여러개를 div가 묶는다
- .btn-group-vertical
- .btn-group-justified
- .btn-toolbar : 버튼 그룹을 묶는다
  - .btn-toolbar > .btn-group

### button group dropdown

- dropdown div를 button group에 넣고 class를 .btn-group으로 변경
- 드랍다운 글자를 지울경우 : .caret만 남겨둔다
- `<span class="sr-only">내용</span>` 을 추가한다

### input group

- div.input-group
  - input.form-control
  - span.input-group-addon .input-group-btn : span 사이에 내용을 넣는다

### navs(네비게이션)

- ul.nav .nav-tabs
- ul.nav .nav-tabs .jav-justified (균등배분)
- ul.nav .nav-pills
- ul.nav .nav-pills .nav-stacked
  - li.active data-toggle="tab" > a[href="#contentID"] {내용}
  - li > a[href="#contentID"] {내용}
  - li.dropdown : 드랍다운 메뉴
- div.tab-content
  - div.tab-pane fade in active id="home"
  - div.tab-pane fade id="profile"

### navbar(네비바)

- nav.navbar .navbar-default / .navbar-inverse / .navbar-fixed-top|bottom (상|하단 고정)
  - div.navbar-header
    - button.navbar-toggle data-toggle="collapse" data-target="#menu1"
      - `<span.sr-only>Toggle</span>`
      - span.icon-bar
      - span.icon-bar
    - a.navbar-brand

  -div.collapse .navbar-collapse id="menu1" (메뉴, 폼)
    - ul.nav .navbar-nav
      - li > a {메뉴명}
    - ul.nav .navbar-nav .navbar-right
      - li > a {메뉴명}
    - form.navbar-form
      -div.form-group
        - input.form-control :text
        - input.btn .btn-default .navbar-btn :button

### breadcrumbs

- ol.breadcrumbs
  - li.active > {내용}
  - li > a {내용}

### pagination

- ul.pagination .pagination-크기
  - li.disabled > a &laquo;
  - li.active > a
  - li > a &raquo;

### pager

- ul.pager
  - li > a {내용}
  - li.previous > a {내용} : 좌측 정렬
  - li.next > a {내용} : 우측 정렬

### transitions (전환효과)

  ```javascript
  $(".tab-pane").on($.support.transition.end, function() {
    console.log("콜백함수 내용");
  });
  ```

### modal

- button.btn .btn-priamry data-toggle="modal" data-target="#myModal"
- div.modal .fade #myModal data-backdrop="static"(배경눌러 끄기X)
  - div.modal-dialog > div.modal-content
    - div.modal-header
      - {내용}
      - button.close data-dismiss="modal" &times;
    - div.modal-body {내용}
    - div.modal-footer
      - {내용}
      - button.btn .btn-primary data-dismiss="modal"

- 스크립트로 modal 생성
  ```javascript
    $("#btnShow").click(function() {
        $("#myModal").modal("show"); // modal 보이기
        $("#myModal").modal("hide"); // modal 숨기기
    });
  ```

- modal Event
  ```javascript
    $("#myModal").on("show.bs.modal", function(e) {
      console.log("event", e.type);
    });
  ```

### scrollspy (스크롤 감시)

- div.nav-container의 위치가 relative로 지정되어야 함
- `<body data-spy="scroll" data-taget="#navbarID" data-offset="활성화위치설정(탭 높이만큼 줄여준다)">`
  ```javascript
    $("body").scrollspy({
      target:"#navbarID",
      "data-offset":55
    });
  ```

### tab (navs)

  ```javascript
    // 탭에 content 이벤트걸기
    $("#tab1 a").click(function(e) {
      e.preventDefault();
      $(this).tab("show');
    });
    // 버튼으로 content show
    $("#btn").click(function(e) {
      $("#tab1 li").eq(1).tab("show");
    });
  ```

### tooltip

- `<span data-toggle="tooltip" data-title="툴팁내용" data-animation="true" data-placement="auto" data-html="true" data-trigger="hover focus" data-delay="1000" data-container="#ID">`
  ```javascript
    $("[data-toggle=tooltip]").tooltip();
    $("#대상").tooltip({
      title:"툴팁내용"
    });
    // 버튼으로 툴팁
    $("btn").click(function() {
      $("#tip1").tooltip("hide/show/toggle");
    });
  ```
- 동적 생성시
  ```javascript
    $("body").tooltip({
      selector:"[data-toggle=tooltip]"
    })

    $("#btnAdd").click(fucntion() {
      $("#row1").append('<button class="btn btn-primary" data-toggle="tooltip">');
    })
  ```

### popover (툴팁과 유사한 기능)

- 툴팁과 거의 동일함
- data-content="팝업 내용"
  ```javascript
  $("[data-toggle=popover]").popover();
    $("#대상").popover({
      title:"팝오버 제목",
      content:"팝오버 내용"
    });
  ```

### alert

- div.alert .alert-success {내용}
  - button.close data-dismiss="alert" {&times;}
  ```javascript
  // 호출
  $("#id").click(function() {
    $("#alert").alert("close");
  });
  // 이벤트
  $("#alert").on("close/closed.bs.alert", function() {
    console.log("이벤트 내용");
  });
  ```

### collapse

- `data-toggle="collapse" data-target="#id"`
- `div#id.collapse.in`
- collapse 대상 묶기 : `data-parent="#id"`

### carousel (이미지 슬라이더)

- div.carousel .slide #slider
  - ol.carousel-indicators
    - li.active [data-target="#img"] [data-slide-to="0"]
    - li [data-target="#img"] [data-slide-to="1"]
  - div.carousel-inner
    -div.item .active {`<img src="">`} > div.carousel-caption {캡션}
    -div.item {`<img src="">`} > div.carousel-caption {캡션}
  - a.left .carousel-control [href="#slider"] [data-slide="prev"]
    - span.glyphicon .glyphicon-chevron-left
  - a.right .carousel-control [href="#slider"] [data-slide="next"]
    - span.glyphicon .glyphicon-chevron-right
  ```javascript
  $("#slider").carousel("cycle/pause/prev/next")
  // 이벤트
  $("#slider").on("slide/slid.bs.carousel", function(e) {
    console.log("내용", e);
  });
  ```

### affix (고정 기능)

- tag.affix data-spy="affix" data-offset-top/bottom="300"
  ```javascript
  $("#target").affix({
    offset:{
      top:300
    }
  });
  ```