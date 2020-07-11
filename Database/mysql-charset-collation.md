# MySQL character sets and collations

<https://dev.mysql.com/doc/refman/8.0/en/charset-general.html>

character set은 기호와 인코딩의 세트이다.
collation은 character set에서 문자를 비교하기 위한 규칙모음이다.

두 문자값을 비교한다고 할 때 가장 간단한 방법은 인코딩내에서의 순서를 보는 것이다.
그러나 대소문자, 각국의 다양한 확장된 알파벳등을 정렬하려면 규칙이 필요하다.

MySQL에서는

- 다양한 character sets으로 문자열을 저장
- 여러가지 collations를 사용하여 문자열을 비교

## Character Sets In MySQL

MySQL의 기본 문자집합은 latin1(cp1252 West European)이다.

그외 다양한 문자집합을 지원하지만 일반적인 상황에서 사용하는 유니코드 문자집합은 다음과 같다.

<https://dev.mysql.com/doc/refman/8.0/en/charset-unicode-sets.html>

- utf8mb3: A UTF-8 encoding of the Unicode character set using one to three bytes per character.
- utf8: An alias for utf8mb3.
- utf8mb4: A UTF-8 encoding of the Unicode character set using one to four bytes per character.
- ucs2: The UCS-2 encoding of the Unicode character set using two bytes per character.
- utf16: The UTF-16 encoding for the Unicode character set using two or four bytes per character. Like ucs2 but with an extension for supplementary characters.
- utf16le: The UTF-16LE encoding for the Unicode character set. Like utf16 but little-endian rather than big-endian.
- utf32: The UTF-32 encoding for the Unicode character set using four bytes per character.

> `utf8mb3` 문자집합은 deprecated 이며 미래의 MySQL에서 제거될 것이다. 대신 `utf8mb4` 문자집합을 사용해야 한다.
> 현재 `utf8`은 `utf8mb3`의 별칭이지만, 미래 특정시점에 `utf8mb4`의 별칭이 될 것이다.
> `utf8`의 모호함을 피하기 위해서 문자열셋을 명시적으로(`utf8mb-`) 지정하는 것을 고려해야 한다.

- `utf8mb4`, `utf16`, `utf16le`, `utf32` 문자집합은 Basic Multilingual Plane (BMP) characters 및 BMP 외부의 supplementary characters를 지원한다
- `utf8`, `ucs2` 문자집합은 BMP characters만 지원한다

대부분의 유니코드 문자집합은 general collation이 있다. (`_general`으로 표시)
예를 들어, `utf8mb4` 문자집합은 `utf8mb4_general_ci`, `utf8mb4_bin`의 general 및 binary 정렬이 있다.
`utf8mb4_danish_ci` 경우 언어별 데이터 정렬이다.

대부분의 문자집합은 단일 이진 데이터 정렬만 있지만, `utf8mb4`는 `utf8mb4_bin` and (as of MySQL 8.0.17) `utf8mb4_0900_bin` 두가지 이다.

## Collations Sets In MySQL

유니코드를 사용한다면 다음 두 가지의 정렬방식을 주로 사용할 것이다

- utf8_general_ci (or utf8mb4_general_ci)
- utf8_unicode_ci (or utf8mb4_unicode_ci)

### collation naming conventions

| Suffix | Meaning            |
| ------ | ------------------ |
| `_ai`  | Accent-insensitive |
| `_as`  | Accent-sensitive   |
| `_ci`  | Case-insensitive   |
| `_cs`  | Case-sensitive     |
| `_ks`  | Kana-sensitive     |
| `_bin` | Binary             |

## 문자집합 및 정렬 & 인덱스

<https://stackoverflow.com/questions/59665045/will-existing-indexes-be-affected-when-changing-character-set-and-collation-of-m>

> An index is an ordered list of pointers to the rows of the table.
> The ordering is based on both the CHARACTER SET and COLLATION of the column(s) of the index.
> If you change either, the index must be rebuilt. A "pointer" (in this context) is a copy of the PRIMARY KEY.

특히 기본 utf8 문자집합이 변경되어 충돌발생 가능성이 존재

- utf8 charset + utf8_general_ci collation
- utf8mb4 charset + utf8mb4_unicode_ci collation
