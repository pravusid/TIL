# cURL(Client URL)

`man curl`

## Cheat Sheet

- 상세 내용 출력(--verbose): `-v`
- 진행상황 감춤(--silent): `-s`
- maximum time(--max-time): `-m <seconds>`
- not check SSL certificates(--insecure): `-k`

- 파일로 출력(--output): `-o <file>`
- remote-name으로 파일 출력(--remote-name): `-O`
- 이어받기(--continue-at): `-C <offset>`

- method(--request): `-X "METHOD"`

- data(해당 옵션을 사용하면 POST method로 요청함)
  - ascii(--data-ascii alias --data): `-d "data"`, `-d "key1=val1&key2=val2"`
  - binary(--data-binary): 줄바꿈을 제거하고 보냄: `--data-binary @file`
  - urlencode: `--data-urlencode`
  - multipart/form-data(--form): `-F username=kim;profile=@file`
  - PUT File Upload(--upload-file): `-T`

- follow redirects(301, 302)(--location): `-L`

- auth(--user): `-u user:password`

- with Headers(--include): `-i`
- only Headers(--head): `-I`
- headers(--header): `-H "name: value"`
- request a compressed response: `--compressed`
- dump header(--dump-header): `-D <file>`

- cookies
  - read cookies(--cookie ): `-b <file>` `-b "key1=val1; key2=val2"`
  - read and write cookies(--junk-session-cookies): `-b -j <file>`
  - write cookies(--cookie-jar): `-c <file>`
  - send cookies: `-b "c=1; d=2"`

- user-agent: `-A "string"`
- use proxy(--proxy): `-x [protocol://]host[:port]`

## Example

- GET

  ```sh
  curl "http://pravusid.kr/data"
  ```

- POST

  ```sh
  curl -H -X POST -F "key=value" -F "filename=@/file/path" "http://pravusid.kr/upload"
  ```

- PUT

  ```sh
  curl -H 'Content-Type: application/json' \
  -X PUT -d '<"key": "val", ...">' "http://pravusid.kr/data"
  ```

- DELETE

  ```sh
  curl -H 'Content-Type: application/json' \
  -X DELETE -d '<"key":"val", ...>' "http://pravusid.kr/data"
  ```
