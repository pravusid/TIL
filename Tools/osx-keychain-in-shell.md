# OSX keychain in Shell

```sh
# 추가
security add-generic-password -U -a ${USER} -D "environment variable" -s "key" -w "secret"

# 조회
security find-generic-password -w -a ${USER} -D "environment variable" -s "key"
```
