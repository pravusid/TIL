# Go Language 문자열 / 자료구조

## 문자열

문자열은 바이트가 연속적으로 나열되어 있는 것으로 Go 에서는 `string`이라는 자료형을 사용한다.
연속된 바이트를 나타내는 방식으로 `[]byte`가 있는데 이것과 달리 `string`은 읽기 전용이다.

### 유니코드 처리

Go 언어의 기본 인코딩은 UTF-8로 되어있다. 따라서 문자 처리에 1바이트 부터 최대 6바이트 까지 사용될 수 있다.

```go
for i, r := range "가나다" {
    fmt.Println(i, r)
}
fmt.Println(len("가나다"))

// output
// 0 44032
// 3 45208
// 6 45796
// 9
```

index는 글자당 3바이트씩이라 0, 3, 6이 출력되고 전체 글자는 9바이트 이므로 마지막에 9가 출력됨

### 테스트

테스트 파일은 테스트 대상 파일과 같은 이름 끝에 `_test`를 붙인다(`pikachu.go` -> `pikachu_test.go`)

그리고 난 뒤 콘솔에 테스트할 내용을 출력하고, 예상되는 결과를 `// Output:` 주석 다음줄 부터 주석으로 쓴다.

```go
package hangul

import "fmt"

func ExampleHasConsonantSuffix() {
    fmt.Println(HasConsonantSuffix("Go 언어"))
    fmt.Println(HasConsonantSuffix("그럼"))
    fmt.Println(HasConsonantSuffix("피카피카츄"))
    // Output:
    // false
    // true
    // false
}
```

### 바이트 단위 처리

문자열을 반복문에서 어떻게 사용하는지에 따라 유니코드 문자단위 혹은 바이트 단위로 동작한다.

문자열을 바이트 단위로 출력하려면 `fmt.Printf`를 사용할 수 있다.

`fmt.Printf("%x\n", "가나다")`
`fmt.Printf("% x\n", "가나다")`

바이트 단위로 스트링을 16진수로 출력하며 `% x`는 바이트 단위 사이 공백을 넣어 출력한다.

문자열은 읽기 전용이기 때문에 바이트를 조작하는 것은 불가능하다

```go
s := "가나다"
s[2]++ // 에러 발생
```

문자열을 바이트 단위의 슬라이스로 변환할 수 있다.
슬라이스는 배열을 유연하게 만든 자료구조로 자바의 ArrayList, C++의 vector와 비슷하다.

`rune`이 `int32`의 별칭이듯 `byte`는 `uint8`의 별칭이다.

문자열을 byte 슬라이스로, byte 슬라이스를 문자열로 형변환 할 수 있다.

```go
func Example_modifyBytes() {
    b := []byte("가나다")
    b[2]++
    fmt.Println(string(b))
    // Output:
    // 각나다
}
```

어떤 문자열이 들어있는지가 중요하다면 string, 실제 바이트 표현이 중요하다면 []byte를 이용하는 것이 좋다.
