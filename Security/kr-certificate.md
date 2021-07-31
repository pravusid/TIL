# 대한민국 공인인증서

<https://ko.wikipedia.org/wiki/%EA%B3%B5%EB%8F%99%EC%9D%B8%EC%A6%9D%EC%84%9C>

공동인증서로 이름이 바뀌었고 5개의 발급기관이 있다

그 중 은행에서 발급하는 인증서는 금융결제원(yesssign), 증권사에서 발급하는 인증서는 코스콤(signkorea)이 기본이다.

## 기술규격

- <https://www.rootca.or.kr/kor/standard/standard01A.jsp>
- [인증서 구조](./public-key-infrastructure.md#인증서-구조)

발급받은 공인인증서는 인증서(`.cer`)와 개인키(`.key`)로 구성되어 있다

- 공인인증서는 [X.509 V3 규격](https://en.wikipedia.org/wiki/X.509#Structure_of_a_certificate)을 따르고 있다

- 개인키는

  - [PKCS#5](https://datatracker.ietf.org/doc/html/rfc2898) 및 [PKCS#8](https://en.wikipedia.org/wiki/PKCS_8) 형식을 사용하며
  - [SEED 알고리즘](https://seed.kisa.or.kr/kisa/algorithm/EgovSeedInfo.do)으로 암호화 되어 있다

### PKCS#8

> 암호화된 개인키 포맷 표준이며 PEM 형식인 경우 헤더가 `-----BEGIN ENCRYPTED PRIVATE KEY----` 문구로 시작한다

mono(.NET opensource) 구현체 (C#)

- <https://github.com/mono/mono/blob/main/mcs/class/Mono.Security/Mono.Security.Cryptography/PKCS8.cs>

BouncyCastle 구현체 (Java)

- <https://github.com/bcgit/bc-java/blob/master/jce/src/main/java/javax/crypto/EncryptedPrivateKeyInfo.java>
- <https://github.com/bcgit/bc-java/blob/master/core/src/main/java/org/bouncycastle/asn1/pkcs/EncryptedPrivateKeyInfo.java>

### 식별번호를 이용한 본인확인 기술규격

## 공인인증서 활용

### 국세청 홈택스
