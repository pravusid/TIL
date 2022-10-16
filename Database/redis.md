# Redis

- 레디스는 오픈소스, 인-메모리 자료구조 저장소이며 데이터베이스, 캐시, 메시지브로커, 스트리밍엔진으로 사용될 수 있다
- 레디스는 strings, hashes, lists, sets, sorted sets with range queries, bitmaps, hyperloglog, geospatial indexs 그리고 streams 자료구조를 지원한다
- 레디스는 내장 복제기능, lua scripting, LRU eviction, transactions 그리고 여러 단계의 디스크 저장을 지원한다
- 레디스 센티넬을 통해 고가용성을, 레디스 클러스터를 통해 자동 파티셔닝을 지원한다

> -- <https://redis.io/docs/about/>

- 레디스는 요청을 직렬처리한다

> -- <https://redis.io/docs/reference/optimization/latency/#single-threaded-nature-of-redis>

## Refs

- <https://redis.io/docs/>
- <https://meetup.toast.com/posts/224>
- <https://meetup.toast.com/posts/225>
- <https://meetup.toast.com/posts/226>
- <https://meetup.toast.com/posts/227>
- <https://meetup.toast.com/posts/251>
- <https://meetup.toast.com/posts/245>
- <https://tech.kakao.com/2020/11/10/if-kakao-2020-commentary-01-kakao/>
- <https://tech.kakao.com/2022/02/09/k8s-redis/>
- <https://www.youtube.com/watch?v=mPB2CZiAkKM>
- <https://deview.kr/2021/sessions/526>

## 키 (key)

- 문자열을 사용한다 (빈문자열도 허용)
- 최대 크기는 512Mb

## 자료구조 (data structure)

<https://redis.io/docs/data-types/>

### strings

- 기본 자료구조로 볼 수 있다
- 최대 크기는 512MB
- `incr` `getset` 등을 사용할 수 있다

### lists

- linked list 자료구조이다

### hashes

- key 값에 대응하는 value 구조가 `{ [field]: value }` 이다

### sets

- 중복되지 않고 정렬되지 않은 문자열 집합이다
- 집합연산을 수행할 수 있다

### sorted sets

- 중복되지는 않으나 score 값으로 정렬된 문자열 집합이다
- 정렬된 값들을 빠르게 조회할 수 있다

### 기타 자료구조

> bitmaps, hyperloglog, geospatial indexs, streams ...

## 캐시 전략

<https://docs.aws.amazon.com/ko_kr/AmazonElastiCache/latest/red-ug/Strategies.html>

### lazy loading

- 캐시에 데이터 요청
- 캐시 적중
  - 캐시는 데이터를 반환한다
- 캐시 누락
  - 캐시는 null을 반환
  - 애플리케이션은 원본 데이터를 조회
  - 애플리케이션은 조회한 데이터를 캐시에 업데이트 하고 사용함

pros & cons

- 장점
  - 요청된 데이터만 캐싱
  - 장애가 애플리케이션에 치명적인 영향을 주지 않음
- 단점
  - 캐시가 누락된 상황에서 성능상 손해가 발생
  - 캐시 누락 상황에서만 캐시를 업데이트 하면 outdated 발생할 수 있음

### write through

데이터베이스에 데이터를 작성할 때마다 캐시를 업데이트 함

pros & cons

- 장점
  - outdated 캐시가 없음
  - 캐시 누락상황의 페널티가 없음
- 단점
  - 데이터 쓰기를 할 때마다 페널티가 발생 (원본 + 캐시 중복입력)
  - 노드 장애등으로 데이터가 누락될 수 있음
  - 캐시 이탈 (대부분의 데이터는 사용될 일이 없음) -> TTL 사용

## Persistence (영구저장)

<http://redisgate.kr/redis/configuration/persistence.php>

### AOF(Append Only File)

- 입력/수정/삭제 명령이 실행될 때 마다 appendonly.aof 파일에 기록됨
- 기본적으로 파일에 append를 실행하지만, 계속 추가만하면 파일이 너무 커지므로 특정시점에 전체 데이터를 다시 쓴다

### RDB(snapshot)

- RDB는 특정 시점의 메모리에 있는 데이터 전체를 바이너리 파일로 저장
- AOF 파일보다 사이즈가 작다
- `BGSAVE` 또는 `SAVE` 명령으로 RDB 파일을 생성할 수 있다

## 보안

### ACL

<https://redis.io/docs/manual/security/acl/>

## Replication

<https://redis.io/docs/manual/replication/>

- 하나의 master 노드
- 다수의 replica 노드 (replica-replica 관계도 가능)

> This is how a full synchronization works in more details:
>
> The master starts a background saving process to produce an RDB file.
> At the same time it starts to buffer all new write commands received from the clients.
> When the background saving is complete, the master transfers the database file to the replica,
> which saves it on disk, and then loads it into memory.
> The master will then send all buffered commands to the replica.
> This is done as a stream of commands and is in the same format of the Redis protocol itself.

## Sentinel

<https://redis.io/docs/manual/sentinel/>

장애 상황에서 replica 노드를 master 노드로 승격시켜 failover 진행

## Cluster

<https://redis.io/docs/manual/scaling/>

## 활용사례

<https://devs0n.tistory.com/92>

### 좋아요

- key: `comment:like:{userId}`
- dataType: `set`

### 일일 순 방문자

- key `visitors:{YYYYMMDD}`
- dataType: `bitmap`
- userId -> offset으로 처리하고 `BITOP and` 연산

### 최근 검색 목록

- key: `history:{userId}`
- dataType: `sorted set`
- unixtime을 가중치로 사용
- n개만 관리한다면
  - 데이터 추가 후
  - `ZREMRANGEBYRANK` 명령으로 `-(n+1)`번째 데이터를 삭제함

## 주의사항

### 영구저장소로 사용하지 말것

- AOF, RDB 기능은 성능을 떨어뜨리고 장애 발생가능성을 높임
- 스토리지로 사용하려면 AOF를 사용하는 것이 나음

관련 `redis.conf` 설정

```conf
# RDB
save ""
stop-writes-on-bgsave-error no

# AOF
appendonly no
auto-aof-rewrite-percentage 100
```

### 메모리 관련설정을 잊지 말것

```conf
maxmemory <가용 메모리의 60~70%>
maxmemory-policy volatile-lru
```

### 장애발생 가능성이 있는 커맨드 사용하지 말 것

`O(N)` 시간복잡도로 작동하는 커맨드는 장애를 일으킬 수 있다

### 보안설정

- protected-mode: `bind` 옵션으로 지정한 ip 또는 loop-back만 접속 허용
- bind: 최대 16개 까지 접속을 허용할 ip를 지정할 수 있음
- requirepass: 접속에 사용할 비밀번호 지정

### 배포구성 선택

- 고가용필요

  - 샤딩필요 -> cluster
  - 샤딩불필요 -> sentinel

- 고가용불필요

  - master-replica
  - stand-alone
