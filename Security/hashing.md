# Hash

해시함수는 임의길이 데이터를 16진수의 고정길이 데이터로 매핑하는 함수이다.

일반적으로 데이터검색을 위해 쓰이지만 해시를 역추적하기 어렵기 때문에 암호화에도 사용된다.
그러나 해시값과 입력값을 대조해놓은 레인보우 테이블을 사용하여 해독될 가능성이 있으므로 완전한 암호화라고 하기 어렵다.

## Hash 함수 종류

단위 (bit)

| 종류 | 최대 메시지크기 | 블록크기(입력) | 결과값 길이 |
| --- | --- | --- | --- |
| MD5 | 무제한 | 512 | 128 |
| SHA-1 | 2^64 - 1 | 512 | 160 |
| SHA-256 | 2^64 - 1 | 512 | 256 |
| SHA-512 | 2^128 - 1 | 1024 | 512 |

MD5와 SHA-1은 해시충돌 가능성이 보고되어 쓰지 않는다.

## Hash 함수 사용

### Java

#### SHA-256 사용

```java
public void hashing(String str) {
    MessageDigest md = null;
    try {
        md = MessageDigest.getInstance("SHA-256");
    } catch (NoSuchAlgorithmException e) {
        e.printStackTrace();
    }
    md.update(str.getBytes());
    byte byteData[] = md.digest();

    StringBuffer sb = new StringBuffer();
    for (byte b : byteData) {
        sb.append(String.format("%02x", b));
    }
    return sb.toString();
}
```

아래의 코드도 해시값을 16진법 문자열로 변환할 수 있다

```java
StringBuffer sb = new StringBuffer();
for (byte b : byteData) {
    sb.append(Integer.toString((b & 0xff) + 0x100, 16).substring(1));
}
```
