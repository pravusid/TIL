# OpenSSL

OpenSSL을 사용하여 Self Signed Certificate를 생성할 수 있다.

## Root CA

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

## CA-signed Cert for My Server

서버 비밀키 생성: `openssl genrsa -out myserver.key 2048`

CSR 생성(대화형으로 진행됨): `openssl req -new -key myserver.key -out myserver.csr`

CSR 검증: `openssl req -in myserver.csr -noout -text`

### 서버 인증서 생성: FDQN

Fully Qualified Domain Name (FDQN): 기존의 고유이름 방식 (`[hostname].[domain].[tld]`)

```sh
openssl x509 -req -in myserver.csr -CA root-ca.crt -CAkey root-ca.key -CAcreateserial \
  -out myserver.crt -days 1825 -sha256
```

### 서버 인증서 생성: SAN

Subject Alternative Name (SAN): 멀티 도메인 인증을 위한 방식

```sh
openssl x509 -req -in myserver.csr -CA root-ca.crt -CAkey root-ca.key -CAcreateserial \
  -extfile <(printf "subjectAltName=DNS:example.com,DNS:example.net,IP:10.0.0.1") \
  -out myserver.crt -days 1825 -sha256
```

## 검증

`openssl x509 -in myserver.crt -text -noout`
