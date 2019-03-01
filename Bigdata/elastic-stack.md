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
network.host: [ "_local_", "0.0.0.0" ]
```

## Kibana

### Kibana 설치

### Kibana 설정

<https://www.elastic.co/guide/en/kibana/current/settings.html>

### Management: index pattern

### Discover

### Visualize

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

### input

### filter

### output

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

### 실행

- 목록보기: `curator_cli --config ./curator.yml show_indices --verbose`
- 액션 실행
  - `curator --config curator.yml [--run-dry] delete.yml`
  - `--run-dry` 옵션으로 테스트 해볼 수 있음

## SlackAction

<https://www.elastic.co/guide/en/elastic-stack-overview/current/actions-slack.html>
