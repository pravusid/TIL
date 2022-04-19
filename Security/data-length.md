# 암호화 관련 데이터 길이

## 256bits(32bytes) to String

> 1byte 문자열 기준

<https://stackoverflow.com/questions/47412137/converting-bits-in-hexadecimal-to-bytes>

toHex: **64**

- 16진수 한 자리 출력위해(`2^4`) 4bits 필요
- 1byte당 2자리 출력가능
- 32bytes의 경우 길이는 64

toBase64: **44**

- 64진수 한 자리 출력위해(`2^6`) 6bits 필요
- 3bytes당 4자리 출력가능
- 32byes의 경우 길이는 43(+1; 마지막 `=` 문자처리)
