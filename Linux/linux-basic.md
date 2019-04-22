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
- J: tar + xz
- z: tar + gzip
- j: tar + bzip2
- f: 파일이름 지정 (필수)

> `tar cvf <filename>.tar <target>`

### 파일 위치 검색
