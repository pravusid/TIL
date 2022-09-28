# Docker Example: NGINX

## 설정

[[nginx]] 참고

## 실행

```sh
docker run -d --name my_nginx \
    -v $(pwd)/conf.d/:/etc/nginx/conf.d/ \
    -v $(pwd)/tls/:/etc/nginx/tls/ \
    --network host \
    nginx:latest;
```
