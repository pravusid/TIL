# Elastic Stack

## ElasticSearch

### ES 설치

```sh
[1]: max file descriptors [4096] for elasticsearch process is too low, increase to at least [65536]
[2]: max number of threads [1024] for user [elasticsearch] is too low, increase to at least [4096]
[3]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]
[4]: system call filters failed to install; check the logs and fix your configuration or disable system call filters at your own risk
```

- A-1: `/etc/security/limits.conf`

  ```sh
  username hard nofile 65536
  username soft nofile 65536
  username hard nproc 65536
  username sort nproc 65536
  ```

- A-2: ulimit
  - ulimit -n 65536
  - ulimit -u 65536

- B-1: `/etc/sysctl.conf`
  - vm.max_map_count=262144

- B-2: `sysctl -w vm.max_map_count=262144`

A로 해결되지 않으면

`/etc/security/limits.d/90-nproc.conf`

```sh
# changed original soft nproc 1024 to 4096
* soft nproc 4096
```

여전히 해결되지 않으면: `/etc/init.d/elasticsearch`

```sh
start() {
    ulimit -u 4096
    # ...
}
```

### ES 설정

<https://www.elastic.co/guide/en/elasticsearch/reference/current/important-settings.html>

`config/elasticsearch.yml`

```yml
network.host: [ "_local_", "0.0.0.0" ]
```

### Curator

인덱스 limit

## Kibana

### Kibana 설치

### Kibana 설정

<https://www.elastic.co/guide/en/kibana/current/settings.html>

### 데이터 다루기

- Management: index pattern

- Discover

- Visualize

## Security

### Search Guard (Community)

#### Search Guard 설치

<https://docs.search-guard.com/latest/search-guard-installation>

- 호환 버전 확인: <https://docs.search-guard.com/latest/search-guard-versions>
- 설치: `bin/elasticsearch-plugin install -b com.floragunn:search-guard-6:<version>`
- 라이선스 확인: `https://example.com:9200/_searchguard/license`
- 라이선스 제한(`elasticsearch.yml`): `searchguard.enterprise_modules_enabled: false`

#### TLS 설정

<https://docs.search-guard.com/latest/offline-tls-tool>

- TLS-Tool을 다운로드 받아 압축을 풀면 `config` 디렉토리에 `example.yml`과 `template.yml` 파일이 있음
- 해당 파일을 수정하여 `tlsconfig.yml` 파일을 생성함
- 이후 스크립트 실행: `<installation directory>/tools/` 디렉토리에서 `./sgtlstool.sh -c ../config/tlsconfig.yml -ca -crt`
- 인증서 관련 설정파일이 생성됨: `<installation directory>/tools/out/<node>_elasticsearch_config_snippet.yml`
- 생성된 설정파일을 elastic 설정에 적용: `elasticsearch.yml`
- sgadmin 실행:

  ```sh
  ./sgadmin.sh -icl -nhnv -cd /usr/share/elasticsearch/plugins/search-guard-6/sgconfig \
  -cacert /etc/elasticsearch/certs/root-ca.pem \
  -cert /etc/elasticsearch/certs/<admin>.pem \
  -key /etc/elasticsearch/certs/<admin>.key -keypass <certificatepassword>
  ```

#### 계정 설정

- 기본 계정정보: `/usr/share/elasticsearch/plugins/search-guard-6/sgconfig/sg_internal_users.yml`
- 비밀번호 해싱: `plugins/search-guard-6/tools/hash.sh -p mycleartextpassword`
- internal user 사용을 위한 `sg_config.yml` 설정

  ```yml
  authc:
    basic_internal_auth_domain:
      enabled: true
      order: 1
      http_authenticator:
        type: basic
        challenge: true
      authentication_backend:
        type: internal

  authz:
    internal_authorization:
      enabled: true
      authorization_backend:
        type: internal
  ```

- 설정 업데이트: 스크립트(sgadmin 실행) 재실행

#### kibana 설정

플러그인 URL: <https://search.maven.org/search?q=g:com.floragunn%20AND%20a:search-guard-kibana-plugin>

```sh
NODE_OPTIONS="--max-old-space-size=8192" bin/kibana-plugin install https://url/to/search-guard-kibana-plugin-<version>.zip
```

`sg_internal_users.yml` 설정에서 kibana 에서 사용할 계정에 `kibanaserver` role을 부여함
(`sg_roles_mapping.yml`에 alias 되어있음)

```yml
sg_kibana_server:
  readonly: true
  backendroles:
    - kibanaserver
```

`sg_config.yml` 설정 후 스크립트(sgadmin 실행) 재실행

```yml
authc:
  kibana_auth_domain:
    enabled: true
    order: 0
    http_authenticator:
      type: basic
      challenge: false
    authentication_backend:
      type: internal
```

`/etc/kibana/kibana.yml` 설정

```yml
server.host: "pravusid.kr"
elasticsearch.hosts: ["https://pravusid.kr:9200"]
elasticsearch.username: "kibana"
elasticsearch.password: "password"

xpack.spaces.enabled: false
xpack.security.enabled: false

elasticsearch.ssl.verificationMode: none
elasticsearch.ssl.certificateAuthorities: /etc/kibana/root-ca.pem
```

### X-Pack (상용)

#### X-Pack 설정

ElasticSearch

- `bin/elasticsearch-plugin install x-pack`
- Confirm that you want to grant X-Pack additional permissions
- action.auto_create_index in elasticsearch.yml to allow X-Pack to create the following indices:
  - `action.auto_create_index: .security,.monitoring*,.watches,.triggered_watches,.watcher-history*,.ml*`

Kibana

- `bin/kibana-plugin install x-pack`
- `kibana.yml`
  - `elasticsearch.username: "kibana"`
  - `elasticsearch.password: "kibanapassword"`

LogStash

- `bin/logstash-plugin install x-pack`
- `logstash.yml`
  - `xpack.monitoring.elasticsearch.username: logstash_system`
  - `xpack.monitoring.elasticsearch.password: logstashpassword`

#### X-Pack 계정 관련

기본 계정/비밀번호

- elastic / changeme
- kibana / changme

계정 목록

```sh
GET /_xpack/security/user
```

계정 생성

```sh
POST /_xpack/security/user/<username>
{
  "password" : "비밀번호",
  "roles" : [ "superuser" ],
  "full_name" : "Gildong Hong",
  "email" : "hgd@foo.kr",
  "metadata" : {
    "intelligence" : 7
  },
  "enabled": true
}
```

계정 비활성화

```sh
PUT /_xpack/security/user/<username>/_disable
```

계정 삭제

```sh
DELETE /_xpack/security/user/<username>
```

#### Disable X-Pack

elasticsearch.yml, kibana.yml, and logstash.yml configuration files

`xpack.security.enabled: false`

| Setting                  | Description                                              |
| ------------------------ | -------------------------------------------------------- |
| xpack.graph.enabled      | Set to false to disable X-Pack graph features            |
| xpack.ml.enabled         | Set to false to disable X-Pack machine learning features |
| xpack.monitoring.enabled | Set to false to disable X-Pack monitoring features       |
| xpack.reporting.enabled  | Set to false to disable X-Pack reporting features        |
| xpack.security.enabled   | Set to false to disable X-Pack security features         |
| xpack.watcher.enabled    | Set to false to disable Watcher                          |

## Logstash

### input

### filter

### output

## FileBeat

`filebeat.yml`

## HeartBeat

## MetricBeat
