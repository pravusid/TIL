# JAVA GUI

## 배치관리자

모든 컨테이너는 배치관리자가 적용되어야 하며 만일 개발자가 배치관리자를 결정하지 않으면 디폴트 배치관리자가 적용된다.

자바 GUI 5가지 배치관리자

- FlowLayout(컨테이너에 따라 유동적)
- BorderLayout(동서남북, 중앙의 방위를 갖는다)
- GridLayout(행과 열을 지원하는 레이아웃)
- GridBagLayout(x, y 좌표로 배치가능한 레이아웃) Netbean IDE 같은 툴 이용
- CardLayout

## Component

GUI 컴포넌트 큰 2가지 분류

- 컨테이너 - 다른 컴포넌트를 포함할 수 있는 컴포넌트
  > Frame, Panel, Applet(X)

  - Frame >> 기본 BorderLayout

  - Panel >> 기본 FlowLayout (유동적 + 컴포넌트 본래의 크기를 보존)

- 비주얼 컴포넌트(일반 컴포넌트)
  > Button, Checkbox, TextArea...

## 그래픽처리

- 그래픽 주체 : 사람 vs 컴포넌트
- 그리는 행위 : 사람의 붓놀림 vs 컴포넌트의 paint()
- 다양한 표현을 위한 도구 : 팔레트 vs Graphics 객체
- 그래픽의 대상 : 캔버스 vs 컴포넌트 자기자신(Canvas, Panel.... 컨테이너류)

### 자바의 그래픽 처리과정

그래픽처리는 시스템이 처리하므로, 개발자는 그래픽의 갱신등을 원할 때 직접 paint를 호출하면 안되고, repaint를 통해 갱신을 요청해야 한다. 그래야 안정적이다.
  > repaint -> update -> paint

## 이벤트 발생시 처리 과정

1. 리스너
1. 재정의
1. 컴포넌트와 리스너 연결

### 어댑터

이벤트 구현시 사용되는 인터페이스의 추상메서드 수가 너무 많으며, 사용되지도 않는 메서드를 코드에 둘 필요가 있을까?? : 비효율적이다

해결책) 어댑터라는 클래스를 지원하고 있다. 개발자 대신 미리 추상메서드들을 구현해 놓은 객체를 가리킨다.

주의) 어댑터의 자료형은?? 리스너 : implements 리스너 관계에 있으므로

## 자바 Swing

자바swing은 awt의 단점을 개선하기 위해서 등장했다. `javax.swing` 패키지에서 지원한다.

awt의 컴포넌트에 J를 붙인 JFrame, JButton, JPanel 등의 클래스를 사용할 수 있다.