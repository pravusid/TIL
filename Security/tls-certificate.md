# TLS Certificate

## Root CA(certificate authorities)

최상위 인증기관

## Intermediate certificate

<https://en.wikipedia.org/wiki/Chain_of_trust>

> Root CA는 이름 그대로 인증서 신뢰 관계의 원점(Trust anchor)이기 때문에 Root CA가 훼손될 경우 신뢰 구조 전체가 무너지게 됩니다.
> 따라서 네트워크에서 분리된(Air Gapped 또는 오프라인) 환경에서 안전하게 관리되어야 합니다(이런 부분이 잘 관리되고 있는지 검증받아야 CA가 될 수 있기도 합니다).
> 이렇게 Root CA는 중간 CA를 발급(서명)하는 경우에만 제한적으로 사용되고,
> Root CA를 이용해 발급한 중간 CA만 시스템 상에 온라인으로 두고 Leaf 인증서 발급(서명)에 사용하는 위와 같은 구조가 되는 것입니다.

from: <https://engineering.linecorp.com/ko/blog/best-practices-to-secure-your-ssl-tls/>

Root CA의 경우 브라우저나 기본 KeyStore에 포함되어 있다.
하지만 Intermediate Certificate는 포함되지 않을 수 있으므로 TLS 제공자의 서버에서 제공해야 한다.

## Domain certificate

> CA로 부터 인증을 받을 수 있다

예를 들어, Let’s Encrypt 경우 인증을 받으면 다음 파일들이 생성된다

- privkey.pem: 도메인 인증서의 개인키
- cert.pem: 서명된 도메인 인증서

CA 정보를 담은 파일도 함께 생성된다

- chain.pem: Let’s Encrypt의 중간 인증서(intermediate certificate)
- fullchain.pem: `cert + chain`

## Self-signed Certificate

OpenSSL을 사용하여 Self Signed Certificate를 생성할 수 있다.

### Root CA

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

### CA-signed Cert for My Server

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

## Two-way SSL communication

클라이언트 인증이라고도 불린다

- 클라이언트의 인증서와 서버의 인증서를 각각 생성
- 클라이언트와 서버 각각에서 신뢰할 수 있는 CA에 상호의 인증서를 추가
- 양쪽 모두가 신뢰할 수 있는 상황이 아니면 연결을 거부

<https://www.ibm.com/support/knowledgecenter/SSRMWJ_7.0.1.13/com.ibm.isim.doc/securing/cpt/cpt_ic_security_ssl_scenario.html>

## 인증서 변환

### PEM 을 PKCS#12 으로 변환

cert.pem 파일로 통합

`cat domain.crt chain1.crt chain2.crt root.crt > cert.pem`

.pfx 파일로 저장

`openssl pkcs12 -export -name example.com -in cert.pem -inkey private.key -out SecureSign.pfx`

### .pfx 에서 .jks 변환

`keytool -importkeystore -srckeystore SecureSign.pfx -srcstoretype pkcs12 -destkeystore SecureSign.jks -deststoretype jks`

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
