# Docker: no space left on device

dangling volume을 삭제한다

```sh
docker volume rm $(docker volume ls -qf dangling=true)
```
