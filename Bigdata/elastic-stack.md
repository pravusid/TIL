# Elastic Stack

## ElasticSearch

### ElasticSearch 설치

```sh
[1]: max file descriptors [4096] for elasticsearch process is too low, increase to at least [65536]
[2]: max number of threads [1024] for user [howtobiz] is too low, increase to at least [4096]
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

### ElasticSearch 설정

<https://www.elastic.co/guide/en/elasticsearch/reference/current/important-settings.html>

`config/elasticsearch.yml`

- network.host: [ "_local_", "a.b.c.d" ]

## Kibana

### Kibana 설치

### Kibana 설정

<https://www.elastic.co/guide/en/kibana/current/settings.html>
