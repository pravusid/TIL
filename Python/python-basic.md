# 파이썬 기초

## 자료형

### 숫자형

1. 정수형
2. 실수형: 소수점
3. 8진수: `0o`로 시작
4. 16진수: `0x`로 시작

숫자형 연산자

1. 사칙연산: `+`, `-`, `*`, `/`
2. 제곱: `**`
3. 나머지: `%`
4. 나눈 후 정수반환: `//`

### 문자열

1. 큰따옴표, 작은따옴표: 한 줄 문자열
2. 연속된 따옴표 3개: 여러 줄 문자열

#### 이스케이프 코드

1. `\n`: 줄바꿈
2. `\t`: 수평 탭
3. `\\`: 역슬래시
4. `\'`: 작은 따옴표
5. `\"`: 큰 따옴표

#### 문자열 연산

1. 문자열 더하기(`"hello" + "python"`): concatenation
2. 문자열 곱하기(`"hello" * n`: n번 반복 출력

#### 문자열 인덱싱과 슬라이싱

```py
>>> strng = "hello, my friends"
>>> strng[0]
h

# -1은 뒤에서부터 접근 (0에는 -가 붙지 않으므로 1부터 시작)
>>> strng[-1]
s
```

문자열의 일부만 가져올 수 있음(슬라이싱), **슬라이싱은 자료의 복사로 작동한다 (copy)**

```py
>>> strng = "yes, my lord"
>>> strng[0:5]
yes,
```

파이썬의 slice는 반폐구간(뒤의 숫자는 포함되지 않음)이다.

- 끝 번호가 없다면 끝까지: `strng[5:]`
- 시작 번호가 없다면 처음부터: `strng[:5]`
- 모두 생략하면 처음부터 끝까지: `strng[:]`
- `-`기호를 사용할 수 있다: `strng[2:-2]`: 2부터 -3까지의 문자열을 반환한다

#### 문자열 포매팅

- 숫자 대입: `"hello %d wolrd" % 5`
- 문자열 대입: `"hello %s world" % "my"`
- 변수 대입: `"hello %d world" % num`
- 2개 이상의 값: `"hello %d world %s" % (num, strng)`

##### 문자열 포맷 코드

- `%s`: string
- `%c`: character
- `%d`: 정수
- `%f`: 부동소수
- `%o`: 8진수
- `%x`: 16진수
- `%%`: %문자

`%s` 포맷은 어떤 자료형을 대입해도 문자열로 변환되어 들어간다

##### `format()` 함수

문자열 내부에 `{0}, {1} ...` 순서로 위치를 지정하고 `.format(strng0, strng1)` 로 문자를 지정한다.

key-value의 형태로 사용할 수도 있다. `"hello {strng}".format(strng = "my")`

##### 크기와 정렬

문자열의 총 자릿수와 문자열을 정렬할 수 있다

- 왼쪽 정렬: `"{0:<10}".format("hello")`
- 오른쪽 정렬: `"{0:>10}".format("hello")`
- 가운데 정렬: `"{0:^10}".format("hello")`
- 공백 채우기: `"{0:=^10}".format("hello")` -> `"==hello=="`

소수점 크기

- 단위절사: `"{0:0.2f}".format(3.141592)` -> `3.14`
- 공백: `"{0:10.2f}".format(3.141592)` -> `      3.14`

curly brace를 사용하려면 두개를 연속으로 사용하면 된다: `"{{ this }}"` -> `{ this }`

#### f 문자열 포매팅

파이썬 3.6이상부터 f 문자열 포매팅 기능을 사용할 수 있다.

```py
>>> name = "둘리"
>>> f"내 이름은 {name}"
내 이름은 둘리
```

- 표현식 지원: `f"나이는 {age+1}"`
- 딕셔너리 사용: `f"이름은 {d['name']} 이다"`
- 정렬: `f"{'hello':<10}"` -> 'hello     '
- 공백채우기 `f"{'hello':=<10}"` -> 'hello====='
- 소수점 표기: `f"{'3.1415':10.4f}"` -> '      3.14'

#### 문자열 관련 함수

- 문자 개수: `strng.count("a")`: strng 중 a의 개수 반환
- 문자 위치: `strng.find("a")`: strng 중 a가 처음 나온 위치, 없다면 -1 반환
- 문자 위치: `strng.index("a")`: strng 중 a가 처음 나온 위치, 없다면 오류 발생
- 문자열 삽입: `"=".join("abcd")`: 문자열 사이에 앞의 값을 넣는다 -> `a=b=c=d`
- 대문자로 변환: `strng.upper()`
- 소문자로 변환: `strng.lower()`
- 왼쪽 공백 삭제: `strng.lstrip()`
- 오른쪽 공백 삭제: `strng.rstrip()`
- 좌우 공백 삭제: `strng.strip()`
- 문자열 치환: `"Hello, world".replace("world", "town")` -> `Hello, town`
- 문자열 나누기: `strng.split()`, `strng.split(",")`: 값이 없다면 공백, 있다면 해당 문자를 기준으로 나눈다

### 리스트

리스트 생성 `lst = []` 또는 `lst = list()`

#### 리스트의 인덱싱과 슬라이싱

인덱싱과 슬라이싱 방식은 문자열과 동일하다.
중첩된 리스트 역시 인덱싱과 슬라이싱 할 수 있다. (다차원 배열로 접근)

```py
>>> lst = [1, 2, 3]
>>> lst[0]
1
>>> lst[-1]
3
```

슬라이싱 역시 `[시작점:종료점]`이며 종료점이 포함되지 않는 반폐구간이다.
또한 아무 숫자도 없으면 시작점/종료점까지를 가리킨다.

```py
>>> lst = [1, 2, 3]
>>> lst[0:-1]
[1, 2]
```

#### 리스트 연산자

- 결합: `+`: `[1, 2, 3] + [4, 5, 6]` -> `[1, 2, 3, 4, 5, 6]`
- 반복: `*`: `[1, 2, 3] * 2` -> `[1, 2, 3, 1, 2, 3]`

#### 리스트 수정과 삭제

리스트를 여러 방법으로 수정할 수 있다

```py
>>> lst = [1, 2, 3]

# 하나 수정(단일 원소)
>>> lst[1] = 4
[1, 4, 3]

# 하나 수정(리스트)
>>> lst[1] = [1, 2]
[1, [1, 2], 3]

# 연속된 범위 값 수정
>>> lst[1:2] = [5, 6]
[1, 5, 6, 3]
```

다음 방식으로 리스트 원소를 삭제할 수 있다

```py
>>> lst = [1, 2, 3]

# slice로 삭제
>>> lst[1:2] = []
[1, 3]

# del 함수 사용 (리스트 뿐만 아니라 전체 객체를 삭제할 수 있는 방법이다)
>>> del lst[1]
[1, 3]
```

#### 리스트 관련 함수

- 요소추가: `lst.append(5)`: 리스트의 마지막에 해당 원소를 추가한다
- 정렬: `lst.sort()`: 리스트를 오름차순으로 정렬한다
- 뒤집기: `lst.reverse()`: 현재 리스트 원소를 역순으로 배열한다
- 위치: `lst.index(x)`: x원소의 위치를 반환한다
- 삽입: `lst.insert(x, y)`: x번째 위치에 y를 삽입한다
- 요소제거: `lst.remove(x)`: 첫 번째 원소 x를 삭제한다
- 요소 꺼내기: `lst.pop()`: 리스트의 맨 마지막 원소를 반환하고, 목록에서 삭제한다
- 개수: `lst.count(x)`: 리스트에서 x의 개수를 반환한다
- 확장: `lst.extend([x, y, z ...])`: 리스트에 `[x, y, z ...]`를 덧붙인다

### 튜플

튜플은 리스트와 속성이 비슷하지만 수정 불가능하다 (immutable)

값을 바꿀 수 없으므로 선언과 동시에 값을 할당해야 한다: `tupl = (1, 2, 3)`

괄호를 생략할 수도 있다: `tupl = 1, 2, 3`

#### 튜플 인덱싱과 슬라이싱, 연산 ( + * )

인덱싱과 슬라이싱, 연산은 (더하기, 곱하기) 리스트와 동일하다

### 딕셔너리

key-value 여러 개로 구성된 자료형이다

`dic = {"name": "kim", "city": "seoul"}`

#### 딕셔너리 추가, 삭제

```py
# 추가
>>> dic = {"name": "kim"}
>>> dic["city"] = "seoul"
{"name": "kim", "city": "seoul"}

# 삭제
>>> del dic[0]
{}
```

> 딕셔너리 Key는 고유값이므로 중복되는 key값이 있다면 마지막 요소를 제외하고 무시된다

#### 딕셔너리 사용

key를 사용해서 value 얻기

```py
>>> dic = {"name": "kim", "city": "seoul"}
>>> dic["name"]
kim
```

#### 딕셔너리 관련 함수

- 키 목록: `dic.keys()`: `dict_keys` 객체가 반환된다
- 값 목록: `dic.values()`: `dict_values` 객체가 반환된다
- 키,값 쌍 얻기: `dic.items()`: key-value 쌍을 튜플로 묶은 리스트를 반환함
- 모두 삭제: `dic.clear()`
- 키로 값 얻기: `dic.get(key, default-value)`: key에 대응하는 value 반환, 값이 없다면 기본값 반환 가능
- 존재여부: `x in dic`: 해당 key가 딕셔너리 안에 있는지 여부

### 집합(set)

집합 자료형은 `set` 키워드로 생성할 수 있다.
생성시 괄호안에 리스트 자료형(리스트에 준하는-string...)을 넣을 수 있다: `set([1, 2, 3])`

#### 집합 자료형의 특징

- 중복을 허용하지 않음
- 순서가 없음

#### 집합 자료형의 활용

- 교집합: `setA & setB`: `setA.intersection(setB)`
- 합집합: `setA | setB`: `setA.union(setB)`
- 차집합: `setA - setB`: `setA.difference(setB)`

#### 집합 자료형 관련 함수

- 값 추가(1 개): `myset.add(5)`
- 값 추가(여러 개): `myset.update([1, 2, 3])`
- 값 제거: `myset.remove(5)`

### Boolean 자료형

참 거짓을 나타내는 자료형이다. 파이썬에서는 대문자로 시작한다: `True`, `False`

#### 다른 자료형의 bool 값

- `"문자열"`: True
- `""` : False
- `[1, 2, 3]`: True
- `[]`: False
- `()`: False
- `{}`: False
- `1`: True
- `0`: False
- `None`: False

bool 값을 얻기 위해서 내장함수 `bool(x)`을 사용하면 된다

### 변수 생성 Trick

- 튜플로 여러 변수를 동시 할당할 수 있다: `x, y = ("hello", "world")`
- 튜플은 괄호 생략이 가능하다: `(x, y) = "hello", "world"`
- 리스트로 변수 생성이 가능하다: `[x, y] = ["hello", "world"]`
- 여러 변수에 같은 값을 할당 가능하다: `x = y = "hello"`
- 튜플로 두 변수의 값을 간단히 교환 가능하다: `x, y = y, x`

## 제어문

### if문

기본구조는 다음과 같다

```py
if 조건:
    # 수행문
elif 조건:
    # 수행문
else:
    # 수행문
```

#### 비교연산자

| 연산자 | 내용 |
| --- | --- |
| `x<y` | x가 y보다 작다 |
| `x<=y` | x가 y보다 작거나 같다 |
| `x>y` | x가 y보다 크다 |
| `x>=y` | x가 y보다 크거나 같다 |
| `x==y` | x가 y와 같다 |
| `x!=y` | x가 y와 같지 않다 |
| `x and y` | x와 y 둘다 참이어야 한다 |
| `x or y` | x 또는 y 둘중 하나 이상 참이다 |
| `not x` | x는 참이 아니다 |
| `x in y` | y 원소 중 x가 있다: y는 iterable = 리스트, 튜플, 문자열 |
| `x not in y` | y 원소 중 x가 없다 |
| `x is y` | x는 y 자료형이다 |
| `x is not y` | x는 y 자료형이 아니다 |

### while문

반복해서 문장을 수행할 때 사용한다

```py
while 조건문:
    # 수행문
```

`while`문은 `break` 예약어로 탈출 할 수 있다.

`continue` 예약어를 사용하면 예약어 아래의 문장은 건너뛰고 반복문 처음으로 돌아간다.

### for문

```py
for 변수 in iterable:
    # 수행문
```

`while`과 마찬가지로 `break`와 `continue`를 사용할 수 있다

#### range 함수

`for`문을 사용하기 위해 숫자 범위 객체를 만들어주는 `range()` 함수를 사용한다.

- `range(시작, 끝)` / `range(시작, 끝, step)`

range 함수는 끝을 포함하지 않는 반폐구간 이다.

#### list내부 for문: list comprehension

```py
>>> lst = [1, 2, 3]
>>> result = []
>>> for x in lst:
        result.append(x)
>>> print(lst)
[1, 2, 3]

# 위의 코드를 list comprehension을 사용했을 때 다음과 같다
>>> result = [x for x in lst]
>>> print(lst)
[1, 2, 3]
```

일반형은 다음과 같다: `[표현식 for 항목 in 반복가능객체 if 조건]`

`for`문을 연속으로 사용하면 다중 `for`문으로 사용가능 하다.

## 함수

### 함수 정의

```py
def 함수이름(매개변수):
    # 수행문
    return 반환값
```

### 가변 인자

매개변수 앞에 `*`을 붙이면 된다

```py
def 함수명(*매개변수):
    # 수행문
```

### 키워드 인자 (kwargs)

함수의 인수로 key=value 형태의 값을 주면 키워드 인자 딕셔너리에 저장된다

```py
def fun(**kwargs):
    print(kwargs)

>>> fun(x = 1)
{"x": 1}
```

### 초기값 설정

```py
def fun(name, city = "seoul"):
    print("name")
    print("city")
```

만약 `city` 입력값이 없으면 기본값인 seoul로 할당된다

### lambda

사용 예는 다음과 같다: `lambda x, y: x + y`

## 입력과 출력

### 입력

`strng = input("안내문")`

사용자 입력을 받아 `strng` 변수에 할당한다

### 출력

`print(내용)` 함수를 사용한다

`print("hello" "world")` 의 경우 concatenation과 동일하다

`print("hello", "world")` 의 경우 문자열 결합시 space를 넣는다

줄바꿈을 하지 않으려면 `end`인수를 사용한다: `print(x, end="줄끝문자")`

## 클래스

### 구조

```py
class MyClass:
    def __init__(self):
        #초기화
        pass

    def add(self, x, y):
        #수행문
        pass
```

클래스에서 객체를 생성할 때는 생성자를 호출한다

```py
x = MyClass()
```

클래스내의 함수 첫 번째 파라미터는 `self`이다
`self`는 함수를 호출한 인스턴스의 주소를 자동으로 전달 받는다.

### 상속

클래스 상속을 위해서는 클래스 이름 뒤 괄호로 상속할 클래스명을 넣는다

```py
class MyClass(AncestorClass):
    pass
```

메소드 오버라이딩은 별도의 키워드 없이 동일한 메소드를 정의하면 된다

### 클래스 변수

다른 객체지향 언어에서 static 상수 같은 기능이다

## 모듈

모듈이란 함수나 변수 또는 클래스를 모아놓은 파일이다

모듈을 불러올 때 사용하는 키워드는 다음과 같다: `from module import function` 또는 `import module`

## 패키지

패키지는 디렉토리 구조로 되어 있다.

## 예외처리

```py
try:
    # 수행문
except 발생예외 as 예외할당변수:
    # 예외처리
```

`else`절을 사용하면 예외가 발생하지 않았을 때 실행된다. 반드시 `except` 다음에 위치해야 한다

```py
try:
    # 수행문
except:
    # 예외처리
else:
    # 예외 미발생 시 수행
```

`finally`절을 사용하면 예외 발생여부와 관계없이 수행된다

```py
try:
    # 수행문
except:
    # 예외처리
finally:
    # 공통수행
```

여러 개의 오류를 처리하기 위해서는 `except`절을 여러개 사용하거나 예외를 튜플로 묶는다

```py
try:
    # 수행문
except FooError as e:
    # 예외처리
except BarError as e:
    # 예외처리

try:
    # 수행문
except (FooError, BarError) as e:
    # 예외처리
```

### 오류 회피

`pass` 키워드로 오류를 통과할 수 있다

```py
try:
    # 수행문
except:
    pass
```

### 오류 발생

`raise` 키워드로 오류를 생성할 수 있다

`raise FooError`

### 사용자 정의 에러

`Exception` 클래스를 상속받은 사용자 정의 에러를 만들 수 있다

```py
class MyError(Exception):
    pass
```

## 내장함수

- `abs(x)`: 숫자의 절댓값 반환. 인자는 정수 또는 실수이어야 하고, 인자가 복소수면 그 크기가 반환됨
- `all(iterable)`: iterable 의 모든 요소가 참이거나 비어있으면 True 반환
- `any(iterable)`: iterable 의 요소 중 어느 하나라도 참이면 True를 반환하고, 비어 있으면 False를 반환
- `chr(x)`: 아스키 코드값을 입력받아 해당하는 문자를 반환한다
- `dir(obj)`: 객체가 가지고 있는 변수나 함수 목록을 출력함
- `divmod(x, y)`: x를 y로 나눈 몫과 나머지를 튜플형태로 반환함
- `enumerate(iterable)`: 순서가 있는 자료형을 입력받아 인덱스 값을 포함한 enumerate 객체를 반환한다: `(index, value)`
- `eval(expression)`: 실행가능한 문자열을 받아 실행한다 (동적 처리)
- `filter(함수, iterable)`: 반복가능한 자료형을 함수에 넣어 반환되는 값만 모아 돌려줌 (lambda 사용하면 간결해짐)
- `hex(x)`: 정수 값을 받아 16진수를 반환함
- `id(obj)`: 객체를 입력받아 주소값을 반환함
- `input([prompt])`: 사용자 입력을 받아 반환함
- `int(x)`: 문자열을 받아 정수로 반환
- `isinstance(obj, class)`: 첫 번째로 받은 객체가 클래스의 인스턴스인지 여부를 반환
- `len(array)`: 입력값의 길이(개수)를 반환한다
- `list(iterable)`: 반복가능한 자료형을 받아 리스트로 반환한다
- `map(함수, iterable)`: 반복가능한 자료형을 함수로 처리한 뒤 반환받는다
- `max(iterable)`: 반복가능한 자료형에서 최대값을 반환한다
- `min(iterable)`: 반복가능한 자료형에서 최소값을 반환한다
- `oct(x)`: 정수를 받아 8진수를 반환한다
- `open(filename, [mode])`: 파일을 읽어 파일객체를 반환한다 (모드: w-쓰기, r-읽기, a-추가, b-바이너리(wra와 함께사용))
- `ord(c)`: 문자를 받아 아스키 코드를 반환한다
- `pow(x, y)`: x의 y제곱 값을 반환한다
- `range([start], stop, [step])`: 시작점 부터 마지막까지 단계에 맞춘 반폐구간 반복가능 객체를 반환한다
- `repr(obj)`: 객체를 문자열로 반환한다, 반환받은 결과값을 `eval()`로 처리하면 최초 타입 객체로 변환된다.
- `round(number, [digit])`: 숫자로 자리수를 받아 반올림하여 반환함
- `sorted(iterable)`: 입력값을 정렬된 리스트로 반환한다
- `str(obj)`: 객체를 문자열로 변환하여 반환한다
- `tuple(iterable)`: 반복가능한 객체를 튜플로 변환하여 반환한다
- `type(obj)`: 입력받은 객체의 자료형을 반환한다
- `zip(*iterable)`: 같은 개수로 이루어진 반복가능한 자료형을 묶는다 (순서쌍으로 묶음) -> `(a, b) zip (c , d)` -> `(a, c), (b, d)`
