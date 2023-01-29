# Docker Volume

volume 사용현황 및 binding 조회

```sh
docker ps -a --format '{{ .ID }}' | xargs -I {} docker inspect -f '{{ .Name }}{{ range .Mounts }}{{ printf "\n\t" }}{{ .Type }} {{ if eq .Type "bind" }}{{ .Source }}{{ end }}{{ .Name }} => {{ .Destination }}{{ end }}{{ printf "\n" }}' {}
```

dangling volume을 삭제한다

```sh
docker volume rm $(docker volume ls -qf dangling=true)
```
