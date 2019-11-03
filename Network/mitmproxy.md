# mitmproxy

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
