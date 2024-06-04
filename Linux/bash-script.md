# Bash Shell Script

## 환경변수

> - <https://www.gnu.org/software/bash/manual/html_node/Bourne-Shell-Variables.html>
> - <https://www.gnu.org/software/bash/manual/html_node/Bash-Variables.html>

## 변수

- 변수를 선언하지 않으며, 처음 변수에 값이 할당되면 생성된다
- 변수값의 타입은 모두 문자열(string)이다
- 변수이름은 대소문자를 구분한다
- 변수 할당시 `=` 좌우 공백이 없어야 한다
- 변수를 큰따옴표로 묶어도 되고 묶지 않아도 되지만, 공백이 있다면 묶어야 함
- `$` 문자를 사용하기 위해서는 작은따옴표로 묶거나 escape(`\`)문자로 사용해야 함
- 즉, 작은따옴표 내부는 모두 문자로 처리한다
- 큰 따옴표 안에서 사용되는 다음 세 문자는 무시하지 않는다.
  - dollar (변수실행)
  - backtick (명령실행)
  - reverse-slash (escape)

### 변수와 curly braces

<https://stackoverflow.com/questions/8748831/when-do-we-need-curly-braces-around-shell-variables>

```bash
$var      # use the variable
${var}    # same as above
${var}bar # expand var, and append "bar" too
$varbar   # same as ${varbar}, i.e expand a variable called varbar, if it exists.
```

### 변수 내 숫자연산

- 변수 내부에서 숫자와 연산기호를 사용하려면 `expr` 키워드와 backtick(`)을 사용한다
- 괄호 및 곱하기(`*`) 기호도 escape(`\`)문자로 사용해야함

### 파라미터 변수

해당 스크립트 명령을 다음과 같이 실행했다면: `<명령> param1 param2 param2 ... param10`

파라미터 변수가 기본적으로 할당된다

- `$*`=전체 매개변수
- `$@`=전체 매개변수
- `$0`=명령
- `$1`=param1
- `$2`=param2
- `$3`=param3
- ...
- `$10`=param10
- `$#`=매개변수 개수

### 배열

```bash
array=("a" "b" "c")
array[3]="variable"

echo "배열 전체 ${array[@]}"
echo "배열 전체 개수 ${#array[@]}"

# 4번째 요소 삭제
unset array[3]

# 배열 전체 삭제
unset array
```

### 변수 기본값 할당(대체)

```bash
echo 1 ${FOO:="hello"}
echo 2 $FOO
# 1 hello
# 2 hello

echo 1 ${BAR:-"world"}
echo 2 $BAR
# 1 world
# 2
```

<https://unix.stackexchange.com/questions/122845/using-a-b-for-variable-assignment-in-scripts>

|                    | Parameter set and not null | Parameter set but null | Parameter unset |
| ------------------ | -------------------------- | ---------------------- | --------------- |
| ${parameter:-word} | substitute parameter       | substitute word        | substitute word |
| ${parameter-word}  | substitute parameter       | substitute null        | substitute word |
| ${parameter:=word} | substitute parameter       | assign word            | assign word     |
| ${parameter=word}  | substitute parameter       | substitute null        | assign word     |
| ${parameter:?word} | substitute parameter       | error, exit            | error, exit     |
| ${parameter?word}  | substitute parameter       | substitute null        | error, exit     |
| ${parameter:+word} | substitute word            | substitute null        | substitute null |
| ${parameter+word}  | substitute word            | substitute word        | substitute null |

### special parameters

<https://www.gnu.org/software/bash/manual/html_node/Special-Parameters.html>

## shell expansion

<https://www.gnu.org/software/bash/manual/html_node/Shell-Expansions.html>

### export

외부 변수로 선언함 (현재 쉘 세션이 남아있는 동안 유효)

## 제어문

<https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Conditional-Constructs>

### if문

`[ 조건 ]`의 각 단어 사이는 모두 공백이 있어야 한다

```bash
if [ 조건1 ]; then
  조건1이 참인 경우 실행
elif [ 조건2 ]
then
  조건2가 참인 경우 실행
else
  이외의 경우 실행
fi
```

### case문

```bash
case <변수> in
  조건1)
    조건1의 경우 실행;;
  조건2 | 조건3)
    조건2 또는 조건3의 경우 실행;;
  [nN]*)
    앞에 n 또는 N이 들어간 경우 실행;;
  조건4)
    조건4의 경우 무엇인가 실행하고 추가 실행
    마지막으로 실행;;
  *)
    이외의 경우 실행;;
esac
```

### 조건문 상의 비교연산자

> 괄호가 하나인 경우 `test` 명령으로 작동한다 `[ expression ]`

<https://en.wikipedia.org/wiki/Test_(Unix)>

> 확장 비교연산자를 사용하기 위해서는 괄호를 두개 써야한다 `[[ expression ]]`

<https://www.gnu.org/software/bash/manual/html_node/Bash-Conditional-Expressions.html>

- 문자열 비교

  - `"A" == "A"`: 같은경우
  - `"A" != "B"`: 다른경우
  - `-n "str"`: NULL이 아닌경우
  - `-z "str"`: NULL인 경우

- 산술 비교 (`수식1 <산술연산자> 수식2`)

  - `-eq`: 같은경우
  - `-ne`: 다른경우
  - `-gt`: 수식1이 큰경우
  - `-ge`: 수식1이 크거나 같은경우
  - `-lt`: 수식1이 작은경우
  - `-le`: 수식1이 작거나 같은경우

- 파일 처리 (`연산자 <파일이름>`)

  - `-d`: 파일이 디렉토리
  - `-e`: 파일이 존재
  - `-f`: 일반 파일인 경우
  - `-u`: set-user-id가 설정된 경우
  - `-g`: set-group-id가 설정된 경우
  - `-s`: 파일 크기가 0이 아닌 경우
  - `-r`: 파일 읽기 가능
  - `-w`: 파일 쓰기 가능
  - `-x`: 파일 실행 가능

- 패턴매칭

  - <https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Pattern-Matching>
  - `X == pattern`

- 정규표현식

  - `X =~ 표현식`

### 조건문 상의 논리연산자

- `-a` 또는 `&&`
- `-o` 또는 `||`
- `!변수`: 변수가 거짓인경우

`if [ 조건1 ] && [ 조건2 ]`

`-a`나 `-o`는 조건(`[]`) 내부에서 사용할 수 있으나 각 조건을 괄호로 묶어야 함

`if [ \( 조건1 \) -a \( 조건2 \) ]`

## 반복문

### for-in문

```bash
for 변수 in 값1 값2
do
  반복할 명령
done
```

### while문

```bash
while [ 조건 ]
do
  반복할 명령
done
```

### until문

조건이 거짓인 동안 계속 반복한다

```bash
until [ 조건 ]
do
  반복할 명령
done
```

### break, continue

- break: 반복문 종료
- continue: 키워드 아래영역 실행 생략하고 반복문 조건검사로 돌아감

## 함수

함수를 선언하고 호출시 인자를 넘길 수 있다. 넘긴 인자는 함수내부에서 파라미터로 호출 가능하다.

```bash
함수명 () {
  echo `expr $1 + $2`
}

함수명 100 200
```

## eval

문자열을 명령문으로 해석하여 실행함

`eval "ls -al"`

## subshell

`$(<script>)` subshell로 괄호 내부에 스크립트를 작성해 결과값을 사용할 수 있음

`set $(...)`을 사용해서 결과값을 파라미터로 할당할 수 있음 (subshell의 결과를 공백 단위로 $1 부터 할당)

## shift

인자는 9개까지 사용가능(1~9, 0은 명령어) 하므로 10번째 인자부터는 shift를 통해서 모든 파라미터 변수를 한단계식 낮춰 사용해야 한다.

## 예제

### process id

```bash
#!/bin/bash

PROC_NAME=$1
PID=$(pgrep -f $PROC_NAME)
echo "$PID"
```

## 병렬실행

wait 사용

> With no parameters it waits for all background processes to finish. With a PID it waits for the specific process.

```bash
process &
process &
process &
wait
```

[GNU parallel](http://www.gnu.org/software/parallel/) 사용

```bash
echo {1..5} | parallel docker build -t ...{} ...{}
```

> <https://stackoverflow.com/questions/24843570/concurrency-in-shell-scripts>

## 실행 상대경로

```bash
# 스크립트가 실행된 상대경로를 구한다
current_dir=$(dirname $BASH_SOURCE)
# if symlink (GNU)
current_dir=$(dirname $(readlink -f $BASH_SOURCE))

# 실행할 스크립트 내부에서 아래 명령 이후 내용은 wd가 변경된 상태로 수행됨
cd $current_dir
```

> - <https://stackoverflow.com/questions/24112727/relative-paths-based-on-file-location-instead-of-current-working-directory>
> - <https://stackoverflow.com/questions/35006457/choosing-between-0-and-bash-source>
