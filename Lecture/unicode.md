# 유니코드

유니코드(Unicode)는 전 세계의 모든 문자를 컴퓨터에서 일관되게 표현하고 다룰 수 있도록 설계된 산업 표준이며, 유니코드 협회(Unicode Consortium)가 제정한다.

이 표준에는 ISO 10646 문자 집합, 문자 인코딩, 문자 정보 데이터베이스, 문자들을 다루기 위한 알고리즘 등을 포함하고 있다.

## 유니코드 평면

<https://en.wikipedia.org/wiki/Plane_(Unicode)>

유니코드를 구획으로 나눈 것이며 0번(다국어 기본평면)에서 16번까지 17개로 나뉘어 있으며 각 편면은 2^16개 코드로 구성된다.

- 0: 다국어 기본 평면(BMP)

  - 범위: U+0000 ~ U+FFFF
  - 내용: 각 언어의 문자 및 특수문자(한글과 한중일 통합한자 포함)

- 1: 다국어 보충 평면(SMP)

  - 범위: U+10000 ~ U+1FFFF
  - 내용: 옛 문자나 음악기호, 수학기호

- 2: 상형문자 보충 평면(SIP)

  - 범위: U+20000 ~ U+2FFFF
  - 내용: 초기 유니코드에서 제외된 한중일 통합한자가 대부분

- 3: 상형문자 제 3평면

  - 범위: U+30000 ~ U+3FFFF
  - 내용: 갑골 문자, 금문 ... (아직 내용지정되어있지 않음)

- 4~13: 미지정 평면

  -범위: U+40000 ~ U+​DFFFF

- 14: 특수 목적 보충 평면(SSP)

  - 범위: U+E0000 ~ U+​EFFFF
  - 내용: 소수의 제어문자

- 15~16: 사용자 영역 평면

  - 범위: U+F0000 ~ U+​10FFFF
  - 내용: 특정 업체나 사용자별로 할당가능하며 호환성이 보장되지 않음

## Byte Order Mark (BOM)

`U+FEFF`: UTF-16으로 된 파일의 엔디언 식별을 위해서 파일의 맨 앞에 삽입한다.

유니코드 표준은 UTF-8의 BOM을 허용하지만, 그것의 사용은 필수가 아니며 UTF-8에서 바이트 순서는 어떤 의미도 없어서 권장되지도 않는다.

첫 두 바이트가 `FE FF`이면 빅 엔디언, `FF FE`이면 리틀 엔디언으로 식별한다.

## UTF

- 종류: 알파벳 / 한글
- EUC-KR: 1byte / 2bytes
- UTF-8: 1byte / 3bytes
- UTF-16: 2bytes / 2bytes
- UTF-32: 4bytes / 4bytes

### UTF-8

<https://ko.wikipedia.org/wiki/UTF-8>

> Universal Coded Character Set + Transformation Format – 8-bit

켄 톰슨과 롭 파이크가 만들었다. 가변 인코딩으로 한 글자를 1~4바이트로 표기한다.

여러 장점 덕분에 표준적으로 가장 많이 쓰이는 유니코드 인코딩이 되었다.

| 설명       | 코드 범위(십육진법) | UTF-8 표현(이진법)                  | UTF-16BE 표현(이진법)               |
| ---------- | ------------------- | ----------------------------------- | ----------------------------------- |
| ASCII 범위 | 000000-00007F       | 0xxxxxxx                            | 00000000 0xxxxxxx                   |
| 2바이트    | 000080-0007FF       | 110xxxxx 10xxxxxx                   | 00000xxx xxxxxxxx                   |
| 3바이트    | 000800-00FFFF       | 1110xxxx 10xxxxxx 10xxxxxx          | xxxxxxxx xxxxxxxx                   |
| 4바이트    | 010000-10FFFF       | 11110zzz 10zzxxxx 10xxxxxx 10xxxxxx | 110110yy yyxxxxxx 110111xx xxxxxxxx |

- 1바이트로 표시된 문자의 최상위 비트는 항상 0이다

- 2바이트 이상으로 표시된 문자의 경우, 첫 바이트의 상위 비트들이 그 문자를 표시하는 데 필요한 바이트 수를 결정한다

  - 예를 들어서 2바이트는 110으로 시작하고, 3바이트는 1110으로 시작한다

- 첫 바이트가 아닌 나머지 바이트들은 상위 2비트가 항상 10이다

#### UTF-8 오류처리

> U+003F(?, 물음표)나 U+FFFD(�, 유니코드 대치 문자) 같은 다른 문자를 집어 넣는다

### UTF-16

기본 다국어평면의 문자는 16비트 값으로 인코딩하고, 그 이상 범위의 문자는 32비트로 인코딩 한다.

### UTF-32

한 글자를 32비트(4bytes)로 인코딩한다.

## 한글평면

- <https://d2.naver.com/helloworld/19187>
- <https://d2.naver.com/helloworld/76650>

| 이름                                         | 처음 | 끝   | 개수  |
| -------------------------------------------- | ---- | ---- | ----- |
| 한글 자모 (Hangul Jamo)                      | 1100 | 11FF | 256   |
| 호환용 한글 자모 (Hangul Compatibility Jamo) | 3130 | 318F | 96    |
| 한글 자모 확장 A (Hangul Jamo Extended A)    | A960 | A97F | 32    |
| 한글 소리 마디 (Hangul Syllables)            | AC00 | D7AF | 11184 |
| 한글 자모 확장 B (Hangul Jamo Extended B)    | D7B0 | D7FF | 80    |

> 완성형 한글 범위(한글소리마디): '가'(U+AC00)부터 '힣'(U+D7A3)

## 각종 유니코드 공백 (space)

<https://en.wikipedia.org/wiki/Space_(punctuation)#Types_of_spaces>
