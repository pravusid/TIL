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

A로 해결되지 않으면: `/etc/security/limits.d/90-nproc.conf`

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

<https://www.elastic.co/guide/en/elasticsearch/reference/current/settings.html>

`config/elasticsearch.yml`

```yml
path.data: /var/lib/elasticsearch
path.logs: /var/log/elasticsearch

network.host: [ "_local_", "0.0.0.0" ]
http.compression: true

xpack.security.enabled: false
xpack.monitoring.enabled: true
xpack.monitoring.collection.enabled: true

searchguard.enterprise_modules_enabled: false
```

`jvm.options`

```conf
# Xms represents the initial size of total heap space - 최소 시스템 메모리 1/2
# Xmx represents the maximum size of total heap space - 최대 32gb

-Xms4g
-Xmx4g
```

### APIs (kibana Dev Tools Console)

#### Index

<https://www.elastic.co/guide/en/elasticsearch/reference/current/indices.html>

#### Template

<https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html>

- 전체조회: GET /_template
- 삭제: DELETE /_template/{name}

### Data Types

#### Core datatypes

- String
  - [text](https://www.elastic.co/guide/en/elasticsearch/reference/current/text.html): 전체 텍스트 값을 인덱싱
  - [keyword](https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html): 구조화된 값(정확한 값으로만 검색)

- [Numeric](https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html)
  - long: 64bit 정수
  - integer: 32bit 정수
  - short: 16bit 정수
  - byte: 8bit 정수
  - double: double precision 64bit IEEE 754 실수
  - float: single precision 32bit IEEE 754 실수
  - half_float: half precision 16bit IEEE 754 실수
  - scaled_float: 고정 `double` scale factor에 의해 확장된 유한 실수

- [Date](https://www.elastic.co/guide/en/elasticsearch/reference/current/date.html)
  - 날짜 포맷으로 작성된 string
  - epoch miliseconds을 표현하는 long
  - epoch seconds를 표현하는 integer

- [Boolean](https://www.elastic.co/guide/en/elasticsearch/reference/current/boolean.html): 참, 거짓

#### Complex datatypes

- [Array](https://www.elastic.co/guide/en/elasticsearch/reference/current/array.html)
  - 모든 필드는 0개 이상의 값이 포함 될 수 있으나 모든 값은 동일한 타입이어야 한다
  - 객체의 배열은 예상 대로 작동하지 않으므로 중첩 데이터 유형을 사용해야 한다

- [Object](https://www.elastic.co/guide/en/elasticsearch/reference/current/object.html)
  - JSON 문서는 계층적이고 내부에 객체를 포함할 수 있다

- [Nested](https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html)
  - 중첩 타입은 객체 타입의 특수한 형태로 객체 배열을 서로 독립적으로 쿼리할 수 있도록 인덱싱 한다

#### Geo datatypes

<https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-point.html>

- Geo-point expressed as an object, with lat and lon keys
- Geo-point expressed as a string with the format: "lat,lon"
- Geo-point expressed as a geohash
- Geo-point expressed as an array with the format: [lon, lat]
- A geo-bounding box query which finds all geo-points that fall inside the box

## Kibana

### Kibana 설치

<https://www.elastic.co/guide/en/kibana/current/install.html>

### Kibana 설정

<https://www.elastic.co/guide/en/kibana/current/settings.html>

```yml
elasticsearch.hosts: ["https://localhost:9200"]
elasticsearch.username: "kibana"
elasticsearch.password: "password"

xpack.spaces.enabled: false
xpack.security.enabled: false
xpack.monitoring.enabled: true

elasticsearch.ssl.verificationMode: none
elasticsearch.ssl.certificateAuthorities: /etc/kibana/root-ca.pem
```

## Search Guard (Community)

### Search Guard 설치

<https://docs.search-guard.com/latest/search-guard-installation>

- 호환 버전 확인: <https://docs.search-guard.com/latest/search-guard-versions>
- 설치: `bin/elasticsearch-plugin install -b com.floragunn:search-guard-6:<version>`
- 라이선스 확인: `https://example.com:9200/_searchguard/license`
- 라이선스 제한(`elasticsearch.yml`): `searchguard.enterprise_modules_enabled: false`

### TLS 설정

<https://docs.search-guard.com/latest/offline-tls-tool>

<https://www.elastic.co/guide/en/elasticsearch/reference/6.6/certutil.html>

- TLS-Tool을 다운로드 받아 압축을 풀면 `config` 디렉토리에 `example.yml`과 `template.yml` 파일이 있음

- 해당 파일을 수정하여 `tlsconfig.yml` 파일을 생성함 (node URI)

- 이후 스크립트 실행: `<installation directory>/tools/` 디렉토리에서 `sgtlstool.sh -c ../config/tlsconfig.yml -ca -crt`

- 인증서 관련 설정파일이 생성됨: `<installation directory>/tools/out/<node>_elasticsearch_config_snippet.yml`

  - 생성되는 파일들

    - `root-ca.pem`: Root certificate
    - `root-ca.key`: Private key of the Root CA
    - `root-ca.readme`: Passwords of the root and intermediate CAs

    - `[nodename].pem`: Node(Server) certificate
    - `[nodename].key`: Private key of the node(server) certificate
    - `[nodename]_http.pem`: REST certificate, only generated if reuseTransportCertificatesForHttp is false
    - `[nodename]_http.key`: Private key of the REST certificate, only generated if reuseTransportCertificatesForHttp is false
    - `[nodename]_elasticsearch_config_snippet.yml`: Search Guard configuration snippet for this node, add this to elasticsearch.yml

    - `[name].pem`: Client certificate
    - `[name].key`: Private key of the client certificate
    - `client-certificates.readme`: Contains the auto-generated passwords for the certificates

  - 생성된 설정파일을 elastic 설정에 적용: `elasticsearch.yml`

- sgadmin 실행:

  ```sh
  ./sgadmin.sh -icl -nhnv -cd /usr/share/elasticsearch/plugins/search-guard-6/sgconfig \
    -cacert /etc/elasticsearch/certs/root-ca.pem \
    -cert /etc/elasticsearch/certs/<admin>.pem \
    -key /etc/elasticsearch/certs/<admin>.key -keypass <certificatepassword>
  ```

### 계정 설정

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

### kibana 설정

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
elasticsearch.username: "username"
elasticsearch.password: "password"

xpack.spaces.enabled: false
xpack.security.enabled: false

elasticsearch.ssl.verificationMode: none

searchguard.allow_client_certificates: true
elasticsearch.ssl.certificateAuthorities: ["root-ca.crt"]
elasticsearch.ssl.certificate: client.crt
elasticsearch.ssl.key: client.key
elasticsearch.ssl.key.passphrase: "PASSWORD"
```

### LogStash 설정

<https://docs.search-guard.com/latest/elasticsearch-logstash-search-guard>

```conf
input {
  beats {
    port => 9600
    ssl => true
    ssl_certificate_authorities => ["root-ca.crt"]
    ssl_certificate => "server.crt"
    ssl_key => "server.key"
    ssl_key_passphrase => "PASSWORD"
  }
}

output {
  elasticsearch {
    user => logstash
    password => logstash
    ssl => true
    ssl_certificate_verification => false
    ssl_certificate_authorities => ["root-ca.crt"]
    ssl_certificate => "client.crt"
    ssl_key => "client.key"
    ssl_key_passphrase => "PASSWORD"
  }
}
```

## X-Pack (Commercial)

### X-Pack 설정

ElasticSearch

- `bin/elasticsearch-plugin install x-pack`

- Confirm that you want to grant X-Pack additional permissions

- action.auto_create_index in elasticsearch.yml to allow X-Pack to create the following indices:
  - `action.auto_create_index: .security,.monitoring*,.watches,.triggered_watches,.watcher-history*,.ml*`

Kibana

- `bin/kibana-plugin install x-pack`

- `kibana.yml`
  - `elasticsearch.username: "kibana"`
  - `elasticsearch.password: "password"`

LogStash

- `bin/logstash-plugin install x-pack`

- `logstash.yml`
  - `xpack.monitoring.elasticsearch.username: logstash`
  - `xpack.monitoring.elasticsearch.password: password`

기본 계정/비밀번호

- elastic / changeme
- kibana / changme

### Disable X-Pack

elasticsearch.yml, kibana.yml and logstash.yml configuration files

`xpack.security.enabled: false`

| Setting                  | Description                                              |
| ------------------------ | -------------------------------------------------------- |
| xpack.graph.enabled      | Set to false to disable X-Pack graph features            |
| xpack.ml.enabled         | Set to false to disable X-Pack machine learning features |
| xpack.monitoring.enabled | Set to false to disable X-Pack monitoring features       |
| xpack.reporting.enabled  | Set to false to disable X-Pack reporting features        |
| xpack.security.enabled   | Set to false to disable X-Pack security features         |
| xpack.watcher.enabled    | Set to false to disable Watcher                          |

### X-Pack User Management APIs

<https://www.elastic.co/guide/en/elasticsearch/reference/6.6/security-api-users.html>

## Logstash

조건식: <https://www.elastic.co/guide/en/logstash/current/event-dependent-configuration.html#conditionals>

정규표현식: <https://www.elastic.co/guide/en/beats/filebeat/current/regexp-support.html>

실행 옵션: <https://www.elastic.co/guide/en/logstash/6.0/running-logstash-command-line.html>

```sh
nohup /usr/share/logstash/bin/logstash \
  --path.settings /etc/logstash/ \
  -f /etc/logstash/conf.d/ --config.reload.automatic \
  >> /var/log/logstash/logstash-stdout.log \
  2>> /v    ar/log/logstash/logstash-stderr.log &
```

### input

<https://www.elastic.co/guide/en/logstash/current/input-plugins.html>

#### beats

````conf
beats {
    port => 9600
    type => foo-type
}
````

### filter

<https://www.elastic.co/guide/en/logstash/current/filter-plugins.html>

#### grok 필터

구조화되지 않은 데이터의 구조를 선언하여 파싱함

#### json 필터

```conf
filter {
    json {
        source => "message"
        target => "message"
    }
    json {
        source => "[message][atts]"
        target => "[message][atts]"
    }
}
```

### output

<https://www.elastic.co/guide/en/logstash/current/output-plugins.html>

#### elastic search

```conf
output {
    elasticsearch {
        hosts => ["https://localhost:9200"]
        index => "logstash-name-%{+YYYY.MM.dd}"
        user => username
        password => password
        ssl => true
        ssl_certificate_verification => false
        cacert => "/etc/logstash/root-ca.pem"
    }
}
```

### stdout for debug

```conf
output {
    stdout { codec => rubydebug }
}
```

## Beats (FileBeat, MetricBeat, HeartBeat...)

`filebeat.yml`

```yml
# user inputs
filebeat.inputs:
- type: log
  enabled: true
  paths:
    - /var/log/*.log
  include_lines: ['^ERROR', '^WARN']
  exclude_lines: ['^DEBUG']
```

`__beat.yml`

```yml
output.elasticsearch:
  enabled: false

output.logstash:
  hosts: ["127.0.0.1:9600"]

queue.mem:
  events: 4096
  flush.min_events: 128
  flush.timeout: 30s
```

HeartBeat 설정: <https://www.elastic.co/guide/en/beats/heartbeat/current/heartbeat-reference-yml.html>

### 모듈

```sh
filebeat modules enable <module>
filebeat modules disable <module>
filebeat modules list

metricbeat modules enable <module>
metricbeat modules disable <module>
metricbeat modules list
```

### Dashboard 설정

```sh
filebeat setup -e \
metricbeat setup \
heartbeat setup -e \

  -E setup.kibana.host=localhost:5601 \
  -E setup.dashboards.index=customname-* \
  -E output.logstash.enabled=false \
  -E output.elasticsearch.hosts=['localhost:9200'] \
  -E output.elasticsearch.username=USERNAME \
  -E output.elasticsearch.password=PASSWORD \
  -E output.elasticsearch.ssl.verification_mode=none

  # 아래의 값을 명시하지 않으면 Elasticsearch output username/password 사용
  -E setup.kibana.username=<username> \
  -E setup.kibana.password=<password> \
```

### Template 적용

(그다지 추천하는 방법은 아님... -> 기본 템플릿 필드가 모두 들어가버림)

By default, Filebeat automatically loads the recommended template file, fields.yml,
if the Elasticsearch output is enabled.

```sh
filebeat setup --template \
metricbeat setup --template \
heartbeat setup -e \

  -E setup.template.name=customname \
  -E setup.template.pattern=customname-* \
  -E setup.kibana.host=localhost:5601 \
  -E setup.dashboards.index=customname-* \
  -E output.logstash.enabled=false \
  -E output.elasticsearch.index=customname-%{+yyyy.MM.dd} \
  -E output.elasticsearch.hosts=['localhost:9200'] \
  -E output.elasticsearch.username=USERNAME \
  -E output.elasticsearch.password=PASSWORD \
  -E output.elasticsearch.ssl.verification_mode=none
```

## Curator

### 설치

`pip install elasticsearch-curator`

`sg_roles.yml`

```yml
sg_curator:
  cluster:
    - CLUSTER_MONITOR  
    - CLUSTER_COMPOSITE_OPS_RO
  indices:
    'logstash-*':
      '*':
        - UNLIMITED
```

`sg_internal_users.yml`

```yml
curator:
  hash: $2y$12$Y7znAYZWqJBTJSrT8.iHreCyCVhRE5RQ4dKbbLKXtnutdTE2IP2n.
```

`sg_roles_mapping.yml`

```yml
sg_curator:
  users:
    - curator
```

`sg_config.yml`

```yml
clientcert_auth_domain:
  enabled: true
  order: 1
  http_authenticator:
    type: clientcert
    config:
      username_attribute: cn
    challenge: false
  authentication_backend:
    type: noop
```

### 설정

`curator.yml`

```yml
client:
  hosts:
    - 127.0.0.1
  port: 9200
  url_prefix:
  use_ssl: True
  certificate: /etc/elasticsearch/config/root-ca.pem
  ssl_no_validate: True
  http_auth: curator:curator
  timeout: 30
  master_only: False
```

### 실행

- 목록보기: `curator_cli --config ./curator.yml show_indices --verbose`
- 액션 실행
  - `curator --config curator.yml [--run-dry] delete.yml`
  - `--run-dry` 옵션으로 테스트 해볼 수 있음

### 주기적 삭제

`action-delete.yml`

```yml
actions:
  1:
    action: delete_indices
    description: Delete indices older than 30 days (based on index name)
    options:
      ignore_empty_list: True
      disable_action: False
    filters:
      - filtertype: pattern
        kind: prefix
        value: logstash-
      - filtertype: age
        source: name
        direction: older
        timestring: '%Y.%m.%d'
        unit: days
        unit_count: 30
```

## SlackAction

<https://www.elastic.co/guide/en/elastic-stack-overview/current/actions-slack.html>
