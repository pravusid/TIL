# Public Key Infrastructure

## PKCS (public-key cryptography standards)

<https://en.wikipedia.org/wiki/PKCS>

> 암호화 알고리즘 목록 출력: `openssl enc --list`

### PKCS#1

- <https://en.wikipedia.org/wiki/PKCS_1>
- <https://datatracker.ietf.org/doc/html/rfc8017>

RSA 암호표준이며 공개키와 비밀키(ASN.1) 규격, 암복호화/서명에 필요한 알고리즘과 인코딩 등을 정의한다

PEM 형식인 경우 헤더는 다음과 같다

- 비밀키: `-----BEGIN RSA PRIVATE KEY-----`
- 공개키: `-----BEGIN RSA PUBLIC KEY-----`

### PKCS#8

개인키 포맷 표준이며 PEM 형식인 경우 헤더는 다음과 같다

- 암호화된 경우: `-----BEGIN ENCRYPTED PRIVATE KEY----`
- 암호화되지 않은 경우: `-----BEGIN PRIVATE KEY----`

암호화된 경우 본문은 두 개의 Sequence로 구성되어 있다

<https://datatracker.ietf.org/doc/html/rfc5208#section-6>

- `EncryptionAlgorithmIdentifier`: 암호화 알고리즘 Identifier(OID)
- `EncryptedData`: 암호화된 개인키

```asn1
EncryptedPrivateKeyInfo ::= SEQUENCE {
  encryptionAlgorithm  EncryptionAlgorithmIdentifier,
  encryptedData        EncryptedData
}

EncryptionAlgorithmIdentifier ::= SEQUENCE {
  algorithm       OBJECT IDENTIFIER,
  parameters      ANY DEFINED BY algorithm OPTIONAL
}

EncryptedData ::= OCTET STRING
```

암호화되지 않은 경우 다음의 Sequence로 구성되어 있다

<https://datatracker.ietf.org/doc/html/rfc5208#section-5>

```asn1
PrivateKeyInfo ::= SEQUENCE {
  version                   Version,
  privateKeyAlgorithm       PrivateKeyAlgorithmIdentifier,
  privateKey                PrivateKey,
  attributes           [0]  IMPLICIT Attributes OPTIONAL
}
```

> `AlgorithmIdentifier`: <https://datatracker.ietf.org/doc/html/rfc5280#section-4.1.1.2>

mono(.NET opensource) 구현체 (C#)

- <https://github.com/mono/mono/blob/main/mcs/class/Mono.Security/Mono.Security.Cryptography/PKCS8.cs>

BouncyCastle 구현체 (Java)

- <https://github.com/bcgit/bc-java/blob/master/jce/src/main/java/javax/crypto/EncryptedPrivateKeyInfo.java>
- <https://github.com/bcgit/bc-java/blob/master/core/src/main/java/org/bouncycastle/asn1/pkcs/EncryptedPrivateKeyInfo.java>

### PKCS#7

<https://en.wikipedia.org/wiki/PKCS_7>

공인인증서 역시 PKCS#7 표준에 따라 개인키를 사용해 서명한 내용을 공개키와 함께 배포하면 해당 내용의 소유권을 인증할 수 있다

또한 상위기관에서 서명한 CRL(Certificate Revocation List, 인증서 폐기 목록)을 확인하는 용도로도 사용된다.

CRL을 사용한 인증서 검증과정은 다음과 같다

- 인증서의 일련번호
- 인증서 CRL Distribution Points 확인
- CRL 서명 검증
- CRL 내용에서 인증서 일련번호 검색

## ASN.1 (Abstract Syntax Notation One)

- <https://en.wikipedia.org/wiki/ASN.1>
- <https://en.wikipedia.org/wiki/X.690>
- <https://letsencrypt.org/docs/a-warm-welcome-to-asn1-and-der/>

ASN.1은 직렬화/역직렬화 하여 크래스플랫폼에서 사용가능한 추상 구문 구조를 기술하는 표준 인터페이스 표현 언어이다

ASN.1은 X.609 표준에 서술된 BER(Basic Encoding Rules), CER(Canonical Encoding Rules), DER(Distinguished Encoding Rules) 형식으로 인코딩하여 사용한다.

<https://en.wikipedia.org/wiki/X.690#BER,_CER_and_DER_compared>

PKCS 및 X.509 표준에 사용되는 문법이다

## PKI 인증서 구조

- <https://en.wikipedia.org/wiki/X.509>
- <https://ko.wikipedia.org/wiki/X.509>
- <https://datatracker.ietf.org/doc/html/rfc5280>

X.509는 공개키 기반(PKI) 인증구조의 ITU-T 표준이다

### Root CA (certificate authorities)

최상위 인증기관

### Intermediate certificate

<https://en.wikipedia.org/wiki/Chain_of_trust>

> Root CA는 이름 그대로 인증서 신뢰 관계의 원점(Trust anchor)이기 때문에 Root CA가 훼손될 경우 신뢰 구조 전체가 무너지게 됩니다.
> 따라서 네트워크에서 분리된(Air Gapped 또는 오프라인) 환경에서 안전하게 관리되어야 합니다(이런 부분이 잘 관리되고 있는지 검증받아야 CA가 될 수 있기도 합니다).
> 이렇게 Root CA는 중간 CA를 발급(서명)하는 경우에만 제한적으로 사용되고,
> Root CA를 이용해 발급한 중간 CA만 시스템 상에 온라인으로 두고 Leaf 인증서 발급(서명)에 사용하는 위와 같은 구조가 되는 것입니다.

from: <https://engineering.linecorp.com/ko/blog/best-practices-to-secure-your-ssl-tls/>

Root CA의 경우 브라우저나 기본 KeyStore에 포함되어 있다.
하지만 Intermediate Certificate는 포함되지 않을 수 있으므로 TLS 제공자의 서버에서 제공해야 한다.

### Domain certificate

> CA로 부터 인증을 받을 수 있다, Https 인증서는 Domain 인증이지만 공인인증서라면 개인인증서에 해당한다

예를 들어, Let’s Encrypt 경우 인증을 받으면 다음 파일들이 생성된다

- `privkey.pem`: 도메인 인증서의 개인키
- `cert.pem`: 서명된 도메인 인증서

CA 정보를 담은 파일도 함께 생성된다

- `chain.pem`: Let’s Encrypt의 중간 인증서(intermediate certificate)
- `fullchain.pem`: `cert + chain`

## Two-way SSL communication

클라이언트 인증이라고도 불린다

- 클라이언트의 인증서와 서버의 인증서를 각각 생성
- 클라이언트와 서버 각각에서 신뢰할 수 있는 CA에 상호의 인증서를 추가
- 양쪽 모두가 신뢰할 수 있는 상황이 아니면 연결을 거부

<https://www.ibm.com/support/knowledgecenter/SSRMWJ_7.0.1.13/com.ibm.isim.doc/securing/cpt/cpt_ic_security_ssl_scenario.html>

## Self-signed Certificate

OpenSSL을 사용하여 Self Signed Certificate를 생성할 수 있다.

### Self-signed Root CA

Root CA의 개인키를 생성한다: `openssl genrsa -aes256 -out root-ca.key 2048`

CSR로 Root CA 인증서 생성: `openssl req -x509 -new -nodes -key root-ca.key -sha256 -days 1825 -out root-ca.crt`

대화형으로 진행됨

```sh
Enter pass phrase for myCA.key:
-----
Country Name (2 letter code)
State or Province Name (full name)
Locality Name (eg, city)
Organization Name (eg, company)
Organizational Unit Name (eg, section)
Common Name (e.g. server FQDN or YOUR name)
Email Address
```

### Self-signed CA-signed Cert

서버 비밀키 생성: `openssl genrsa -out myserver.key 2048`

CSR 생성(대화형으로 진행됨): `openssl req -new -key myserver.key -out myserver.csr`

CSR 검증: `openssl req -in myserver.csr -noout -text`

#### 서버 인증서 생성: FDQN

Fully Qualified Domain Name (FDQN): 기존의 고유이름 방식 (`[hostname].[domain].[tld]`)

```sh
openssl x509 -req -in myserver.csr -CA root-ca.crt -CAkey root-ca.key -CAcreateserial \
  -out myserver.crt -days 1825 -sha256
```

#### 서버 인증서 생성: SAN

Subject Alternative Name (SAN): 멀티 도메인 인증을 위한 방식

```sh
openssl x509 -req -in myserver.csr -CA root-ca.crt -CAkey root-ca.key -CAcreateserial \
  -extfile <(printf "subjectAltName=DNS:example.com,DNS:example.net,IP:10.0.0.1") \
  -out myserver.crt -days 1825 -sha256
```

### 검증

`openssl x509 -in myserver.crt -text -noout`

### 통신

자체서명 인증서 과정에서 다음파일들이 생성된다

- root CA 비밀키
- root CA 인증서
- 서버 개인키
- 서버 인증서

서버에서는 다음 파일을 가지고 클라이언트 요청을 받는다

- 서버 개인키
- 서버 인증서

> 자체 서명은 사설 root CA 직접 서명이므로 인증서 chain은 필요없다

클라이언트는 다음 파일을 가지고 자체서명 인증서를 사용하는 서버에 요청한다

- root CA 인증서

또는 클라이언트에서 `rejectUnauthorized = false`를 사용할 수도 있다 (인증서 검증하지 않음)

## 인증서 형식

### 인증서 인코딩

<https://www.ssl.com/guide/pem-der-crt-and-cer-x-509-encodings-and-conversions/>

#### `PEM` (Privacy Enhanced Mail)

본문은 Base64 ASCII 인코딩 되어 있으며, 첫 줄과 마지막 줄(헤더, 푸터)은 어떤 종류의 데이터인지 표시한다

```txt
-----BEGIN *****-----
BASE64 ENCODED TEXT
-----END *****_____
```

#### BINARY

X.609(BER, CER, DER) 인코딩된 바이너리(이진) 데이터이다

### 인증서 파일 확장자

#### `.pem`

**PEM** 형식으로 되어있는 파일이며 어떠한 인증서 관련 파일인지는 내용을 확인해 보아야 한다

#### `.crt` or `.cer` (Certificate)

일반적으로 인증서 파일의 확장자로 많이 쓰이고 주로 **PEM** 형식으로 되어있다.
unix 계열에서는 주로 `crt`, 윈도우즈에서는 `cer` 확장자를 사용한다

PEM 인증서인 경우 `-----BEGIN CERTIFICATE-----` 헤더로 시작한다

#### `.csr` (Certificate Signing Request)

인증서 발급 신청을 위해 CA에 요청할 내용이 담겨 있는 파일이며 주로 **PEM** 형식으로 되어있다.

PEM 인증서인 경우 `-----BEGIN NEW CERTIFICATE REQUEST-----` 헤더로 시작한다

#### `.key`

일반적으로 개인키 파일에 사용하는 확장자이다(공개키에 쓸 수도 있다...). 주로 **PEM** 형식으로 되어 있으나, 바이너리 형식일 수도 있다

#### `.pfx` or `.p12` (PKCS#12)

<https://en.wikipedia.org/wiki/PKCS_12>

일반적으로 X.509 인증서와 개인키를 (혹은 다수의 묶음을) 하나의 파일에 포함한 **바이너리** 형식의 _archive file format_ 이다.

PKCS#12 역시 서명하고 암호화 할 수 있으며 이 경우 "SafeBags"라 불린다.

#### `.jks` (Java Key Store)

일종의 자바에서 사용하는 _p12_ 포맷이지만 다른 시스템과 호환성은 없다. 자바9 이후로는 PKCS#12 포맷이 기본이 되었다.

## 비대칭 키 생성 (openssl)

<https://www.openssl.org/docs/manpages.html>

```sh
# private key
openssl genrsa -out private.pem 2048
# private key with cipher
openssl genrsa -aes-256-cbc -out private.pem 2048
# public key
openssl rsa -in private.pem -out public.pem -pubout
```

## 인증서 인코딩 변환

### PEM -> 바이너리(DER)

```sh
# 인증서 변환
openssl x509 -in cert.pem -out cert.der -outform der

# 개인키 변환 (암호화 파일은 암호 입력후 복호화 출력됨 -> 다시 암호화 필요)
openssl rsa -in priv.pem -out priv.key -outform der
```

### 바이너리(DER) -> PEM

```sh
# 인증서 변환
openssl x509 -in cert.der -inform der -out cert.pem

# 개인키 변환 (암호화 파일은 암호 입력후 복호화 출력됨 -> 다시 암호화 필요)
openssl rsa -in priv.key -inform der -out priv.pem
```

## 인증서 유형 변환

### 개인키 -> PKCS#8

```sh
# Convert a private key to PKCS#8 format using default parameters (AES with 256 bit key and hmacWithSHA256)
openssl pkcs8 -in private.pem -topk8 -out private.key
```

### PKCS#8 -> PKCS#1

```sh
openssl rsa -in pkcs8.pem -out pkcs1.pem
```

### PKCS#12 -> PEM

```sh
# 인증서 추출
openssl pkcs12 -in cert.p12 -out cert.pem -clcerts -nokeys

# 복호화된 개인키 추출
openssl pkcs12 -in cert.p12 -out pri.pem -nocerts -nodes

# 암호화된 개인키 추출
openssl pkcs12 -in cert.p12 -out encPri.pem -nocerts
# 개인키 복호화
openssl rsa -in encPri.pem -out pri.pem
```

### PEM -> PKCS#12

```sh
# cert.pem 파일로 통합 (full-chain)
cat domain.crt chain1.crt chain2.crt root.crt > cert.pem

# .pfx/.p12 파일로 저장 (cert + privkey)
openssl pkcs12 -export -out unified.p12 -in cert.pem -inkey priv.key [-name example.com] [-certfile root-chain.cer]
```

### PKCS#12 -> JKS

`keytool -importkeystore -srckeystore unified.p12 -srcstoretype pkcs12 -destkeystore unified.jks -deststoretype jks`

```sh
대상 키 저장소 비밀번호 입력: ******
새 비밀번호 다시 입력: ******
소스 키 저장소 비밀번호 입력: ****** (pfx 생성시 지정한 암호)
```

> <https://www.securesign.kr/guides/SSL-Certificate-Convert-Format>

#### jks 오류 (Cannot recover key)

> Solution: The KeyStore password and The Key password should be the same.

Changing both passwords using keytool

Change KeyStore password

```sh
keytool -storepasswd -new %newpassword% -keystore %YourKeyStore%.jks
# replace %newpassword% with your actual password, same with YourKeyStore
```

Change Alias key Password

```sh
keytool -keypasswd -alias %MyKeyAlias% -new %newpassword% -keystore KeyStore.jks
# Note: supply old passwords for both keystore and alias when asked for them
```

<https://stackoverflow.com/questions/14606837/cannot-recover-key>
