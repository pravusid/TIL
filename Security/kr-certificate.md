# 대한민국 공인인증서

<https://ko.wikipedia.org/wiki/%EA%B3%B5%EB%8F%99%EC%9D%B8%EC%A6%9D%EC%84%9C>

공동인증서로 이름이 바뀌었고 5개의 발급기관이 있다

그 중 은행에서 발급하는 인증서는 금융결제원(yesssign), 증권사에서 발급하는 인증서는 코스콤(signkorea)이 기본이다.

## 기술규격

- 인증서 관련 전체적인 개념은 [[public-key-infrastructure]] 참고
- <https://www.rootca.or.kr/kor/standard/standard01A.jsp>

발급받은 공인인증서는 인증서(`.cer`)와 개인키(`.key`)로 구성되어 있다

- 공인인증서는

  - [X.509 V3 규격](https://en.wikipedia.org/wiki/X.509#Structure_of_a_certificate)을 따르고 있다
  - [전자서명 인증서 프로파일 규격](https://www.rootca.or.kr/kcac/down/TechSpec/1.1-KCAC.TS.CERTPROF.pdf)에서 인증서 구조를 확인할 수 있다

- 개인키는

  - [PKCS#5](https://datatracker.ietf.org/doc/html/rfc2898) 및 [PKCS#8](https://en.wikipedia.org/wiki/PKCS_8) 형식을 사용하며
  - PKCS#8 개인키를 ASN.1 구조로 파싱하면 [sequence[0]은 암호화 알고리즘, sequence[1]은 암호화된 개인키 정보이다](https://github.com/bcgit/bc-java/blob/f4ba48a0fab38264bce4d1898637de19fc787e9c/core/src/main/java/org/bouncycastle/asn1/pkcs/EncryptedPrivateKeyInfo.java#L21)
  - [암호 알고리즘 규격](https://www.rootca.or.kr/kcac/down/TechSpec/2.3-KCAC.TS.ENC.pdf)을 확인할 수 있다
  - 암호화 방식 (PBES)
    - 암호화 방식 OID 값 확인은 [전자서명인증체계 OID 가이드라인](https://www.rootca.or.kr/kcac/down/Guide/Object_Identifier_Guideline_for_the_Electronic_Signature_Certification_System.pdf) 참조
    - KEY: [PBKDF1 (OID: `1.2.410.200004.1.4`, `1.2.410.200004.1.15`)](https://seed.kisa.or.kr/kisa/algorithm/EgovSeedInfo.do), [PKCS#5 PBES2 (OID: `1.2.840.113549.1.5.13`)](https://datatracker.ietf.org/doc/html/rfc8018#section-6.2)
    - IV: 키 생성 방식에 따라 고정값 또는 키의 일부의 해시값을 사용
    - 암호화 알고리즘: SEED-CBC, ARIA-\*-CBC 등을 사용

### 공인인증서 Certificate Revocation List (인증서 폐기 목록) 확인

> 전자서명 인증서 프로파일 규격 6.2.12, 세부 규격은 <https://www.rootca.or.kr/kcac/down/TechSpec/1.2-KCAC.TS.CRLPROF.pdf> 참고

### 식별번호를 이용한 본인확인 기술규격

<http://www.rootca.or.kr/kcac/down/TechSpec/1.5-KCAC.TS.SIVID.pdf>

`VID = h(h(IDN, R)`

- `VID`: 가상식별번호
- `h`: 해쉬함수
- `IDN`: 개인식별정보 (주민번호, 사업자등록번호)
- `R`: 가입자를 식별할 수 있는 난수(RandomNum), 20Byte, (OID: `1.2.410.200004.10.1.1.3`)

가상 식별번호는 위와 같은 방식으로 클라이언트에서 생성되며
공인인증기관 공개키로 `VID` 및 `R` 값을 암호화하여 보내서 생성된 `VID` 값을 재검증하고 문제가 없으면 공인인증서가 발급된다

식별번호는 공인인증서를 사용하는 기관에서도 요구할 수 있는데 해당 내용은 기술규격 부록 B에 기재되어 있다

- (A) 유저

  - (A-1) 유저는 `IDN`과 `R`을 사용기관에 전달하며 이때 두 값은 안전한 방법으로 전달해야 한다
  - (A-2) A-1 단계의 `IDN` 값이 포함되어 있는 공인인증서를 사용기관에 전달한다

- (B) 사용기관

  - (B-1) A-2 단계에서 유저로 부터 전달받은 공인인증서에서 `VID` 값과 해쉬 알고리즘(`h`)을 추출한다
  - (B-2) A-1 단계에서 유저로 부터 전달받은 `IDN`과 `R`값과 B-1 단계에서 구한 해쉬 알고리즘(`h`)을 사용하여 `VID*` 값을 구한다
  - (B-3) B-1 단계의 `VID` 값과 B-2 단계의 `VID*` 값을 비교한다

> 상황에 따라 `IDN`, `R` 값 중 하나 혹은 전부를 전달하지 않을 수도 있다 (사용기관에서 이미 식별번호를 가지고 있거나 두 값의 해쉬값만 전달)

## 공인인증서 활용

### 국세청 홈택스 로그인

로그인에 필요한 데이터는 다음과 같다

- 공인인증서 일련번호
- 공인인증서 공개키를 PEM 인코딩 한 값
- 홈택스로부터 전자서명 요청받은 값 (pkcEncSsn)
- 개인키로 전자서명한 값 (signed pkcEncSsn)
- 개인키의 RandomNum
