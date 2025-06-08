# mitmproxy

<https://mitmproxy.org/>

## 설치

```sh
docker run --rm -it \
  [-v ~/.mitmproxy:/home/mitmproxy/.mitmproxy] \
  -p 8080:8080 mitmproxy/mitmproxy \
  [mitmdump]
```

for mitmweb

```sh
docker run --rm -it \
  -v ~/.mitmproxy:/home/mitmproxy/.mitmproxy \
  -p 8080:8080 \
  -p 127.0.0.1:8081:8081 \
  mitmproxy/mitmproxy
  mitmweb --web-iface 0.0.0.0
```

설치 후 작동 테스트

```sh
http_proxy=http://localhost:8080/ curl http://example.com/
https_proxy=http://localhost:8080/ curl -k https://example.com/
```

## 사용방법

### <https://scrapfly.io/blog/how-to-install-mitmproxy-certificate/>

- Run `mitmproxy` in the terminal and it'll start a proxy on `localhost:8080` on your machine.
- Start Chrome instance with `mitmproxy` proxy attached to it:
  - Linux: `google-chrome --proxy-server="localhost:8080"`
  - MacOs: `open -a "Google Chrome for Testing" --args --proxy-server="localhost:8080"`
  - Windows: `chrome.exe --proxy-server="localhost:8080"`
- Open `http://mitm.it` in the browser and download the certificate for your system.
- 크롬브라우저에 인증서 등록
  - 기존 방식
    - Open `chrome://settings/security` in the browser.
    - Click on `Authorities` tab.
    - Click on `Import` button and select the certificate from the step 4.
  - macOS 키체인 처리 변경 후
    - macOS 키체인 접근 → mitmproxy 인증서 신뢰 선택
- Now `mitmproxy` will be able to intercept and decrypt all `https` traffic going through it.
- This enables it being used with headless browser tools like Selenium, Playwright or Puppeteer.

### unsafe legacy renegotiation disabled

<https://stackoverflow.com/a/72245418>

```conf
openssl_conf = openssl_init

[openssl_init]
ssl_conf = ssl_sect

[ssl_sect]
system_default = system_default_sect

[system_default_sect]
Options = UnsafeLegacyRenegotiation
```

```bash
OPENSSL_CONF=$FOOBAR ./mitmproxy
```

### 테스트용 크롬 설치

```bash
npm init
npm i @puppeteer/browsers
npx @puppeteer/browsers install chrome@stable
```
