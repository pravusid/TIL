# 리눅스 기본

## 기본 명령어

- `ls`: 파일 목록 출력

- `cd`: 디렉토리 이동

- `pwd`: print working directory

- `rm`: 파일이나 디렉토리를 삭제

- `cp`: 파일이나 디렉토리 복사

- `touch`: 크기가 0인 새파일 생성

- `mv`: 파일이나 디렉토리의 이름을 변경하거나 옮길때 사용

- `mkdir`: 새로운 디렉토리를 생성한다

- `rmdir`: 비어있는 디렉토리를 삭제

- `cat` 파일의 내용을 화면에 보여준다. 여러 파일을 나열하면 파일을 연결해서 보여준다

- `head`: 텍스트 파일의 앞 10행을 출력

- `tail`: 텍스트 파일의 마지막 10행을 출력

- `more`: 텍스트 형식으로 작성된 파일을 페이지 단위로 화면에 출력한다

  - Space: 다음페이지로 이동
  - B: 앞페이지로 이동
  - Q: 종료

- `less`: more 명령과 비슷하지만 보다 확장된 기능

  - more에서 사용하는 키
  - PageUp / PageDown

- `file`: 해당 파일의 종류를 출력

- `clear`: 터미널 화면을 지운다

- `man <명령어>`: 해당 명령어의 매뉴얼을 출력함

## 사용자와 그룹

`/etc/passwd`: `사용자이름:사용자ID:사용자소속그룹ID:전체이름:홈디렉토리:기본셸`

`/etc/group`: `그룹이름:비밀번호:그룹ID:그룹에속한사용자이름`

### 사용자

- `useradd`: 새로운 사용자를 추가한다

- `passwd`: 사용자의 비밀번호 지정/변경

- `usermod`: 사용자의 속성변경 (useradd와 동일 옵션)

- `userdel`: 사용자를 삭제한다

  - `-r`: 사용자 삭제와 동시에 홈 디렉토리 삭제

- `change`: 사용자 암호를 주기적으로 변경하도록 한다

### 그룹

- `groups`: 사용자가 속한 그룹을 보여준다

- `groupadd`: 새로운 그룹을 생성한다

- `groupmod`: 그룹 속성을 변경한다

- `groupdel`: 그룹을 삭제한다

- `gpasswd`: 그룹의 암호를 설정하거나 그룹 관리 수행

## 파일

- `chmod`: 파일 허가권 변경
- `chown`: 파일 소유권 변경

### 링크의 종류

- Hard Link: 하드링크를 생성하면 링크파일이 하나 생성되고 원본과 같은 inode를 사용한다
- Symbolic Link: 새로운 inode를 생성하고 데이터는 원본 파일과 연결된다

### inode

리눅스/유닉스의 파일 시스템에서 사용하는 자료구조로 파일이나 디렉토리의 여러 정보가 있다.

모든 파일이나 디렉토리는 각자 하나씩의 inode가 있고,
각 inode에는 파일의 소유권, 허가권, 종류 등의 정보와 파일의 실제 데이터가 위치하는 주소도 있다.

### 파일 압축/묶기

#### 압축

- xz: xz 압축/압축해제

  - `xz -k <filename>`: filename을 filename.xz로 압축
  - `xz -d <filename>`: filename.xz를 filename으로 압축해제

- bzip2: bz2 압축/압축해제

- gzip: gz 압축/압축해제

- zip: zip 압축

  - `zip <filename>.zip <target>`

- unzip: zip 압축해제

  - `unzip <filename>.zip`

#### 묶기

`tar`

- c: 묶음을 만든다

- x: 묶음을 푼다

- t: 묶인 경로를 보여준다

- C: 지정된 디렉토리에 묶음을 푼다

- v: 진행과정 출력

- 압축 (묶음 과정과 별도)

  - z: gzip
  - j: bzip2
  - J: xz

- f: 파일이름 지정 (필수)

일반적으로 사용하는 명령어는 다음과 같다

- 묶음: `tar zcvf <filename>.tar.gz <target>`
- 풀기: `tar zxvf <filename>.tar.gz [-C 경로선택]`

> 버전에 따라 입력대상(target) 위치가 가장 뒤에 오지 않으면 오류 발생할 수도 있음.  
> 입력대상은 전체(`*`), 숨김파일 포함 전체(`.`), 상대경로(`./foo`)등을 사용하고 디렉토리에 trailing slash 제외

파일/디렉토리 제외: `--exclude <경로>`

`tar: ./<파일명>: file changed as we read it` 에러

- tar 대상 경로 내에 tar 결과파일이 포함되지 않도록 exclude 처리
- (버전에 따라, 파일이 생성되며 구조가 변한것으로 인식하므로) 파일 사전 생성: `touch <파일명>.tar.gz`

### 파일 위치 검색

`find <경로> <옵션>`

- name: 파일이름/확장자
- user: 사용자
- perm: 퍼미션
- size: `-size +10k -size -100k`: 10kb~100kb
- exec: `-exec <명령어> {} \;`찾은파일에 명령실행

`which <실행명령>`: PATH에 정의된 경로만 검색하여 위치를 절대경로로 반환함

`whereis <실행명령>`: 실행파일, 소스, man 페이지를 대상으로 검색

`locate <파일이름>`: `updatedb` 명령으로 갱신된 파일 목록 데이터베이스에서 검색함

## CRON

반복작업을 자동으로 실행할 수 있도록 예약하는 것을 `cron`이라 하고 `crond` 서비스가 이를 수행한다

`/etc/crontab`: `분 시 일 월 요일 <사용자> <명령>`

- 분: 0 ~ 59
- 시: 0 ~ 23
- 일: 1 ~ 31
- 월: 1 ~ 12
- 요일: 0 ~ 6
- every: \*

crontab은 주기에 따라 하위 디렉토리 내용을 호출한다

- cron.hourly
- cron.daily
- cron.weekly
- cron.monthly

## Network

### 명령어

- `nmtui`: network manager text user interface (여러가지 네트워크 설정을 할 수 있다)
- `systemctl <start|stop|restart|status> network`: 네트워크 설정 적용
- `ifup <device>`: 해당 네트워크 장치 작동
- `ifdown <device>`: 해당 네트워크 장치 종료
- `nslookup`: dns 서버와 통신확인
- `ping <ip|URL>`: icmp 프로토콜로 해당 위치와 통신확인

### 설정파일

- `/etc/sysconfig/network`: 네트워크 기본 설정 정보
- `/etc/sysconfig/network-scripts/ifcfg-ens32`: 해당 장치의 네트워크 설정정보
- `/etc/resolv.conf`: DNS 서버 정보와 호스트 이름이 들어있음
- `/etc/hosts`: FQDN(fully qualified domain name) 정보가 들어있음

## 파이프, 리다이렉션

- 파이프(`|`): 파이프 앞의 출력을 파이프 뒤로 전달함

- 리다이렉션: 표준 입출력의 방향을 지정함

  - 하나는 입/출력, 두개는 추가(append)
    - 출력(`>`|`>>`)
    - 입력(`<`|`<<`)

- 에러 포함(`&`)

- 디스크립터

  - 표준 입력 0
  - 표준 출력 1
  - 표준 에러 2

## GRUB

설정파일: `/boot/grub2/grub.cfg`
