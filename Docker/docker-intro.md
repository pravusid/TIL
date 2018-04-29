# Docker

## Docker 소개

컨테이너 기반의 오픈소스 가상화 플랫폼

서버에서 이야기하는 컨테이너도 이와 비슷한데 다양한 프로그램, 실행환경을 컨테이너로 추상화하고 동일한 인터페이스를 제공하여 프로그램의 배포 및 관리를 단순하게 함

이미지는 컨테이너 실행에 필요한 파일과 설정값등을 포함하고 있는 것으로 불변이다(Immutable). 추가되거나 변하는 값은 컨테이너에 다시 저장된다. 같은 이미지에서 여러개의 컨테이너를 생성할 수 있다. 이미지는 컨테이너의 상태가 바뀌거나 컨테이너가 삭제되더라도 변하지 않고 유지된다.

이미지는 여러개의 읽기 전용 레이어로 구성되고 파일이 추가되거나 수정되면 새로운 레이어가 생성.

ubuntu 이미지가 A + B + C의 집합이라면, ubuntu 이미지를 베이스로 만든 nginx 이미지는 A + B + C + nginx

## Docker 설치

Docker는 리눅스 컨테이너 기술을 기반으로 두기 때문에 Windows(Hyper-V)나 MacOS에서는 VM을 사용함

### Linux

설치: `curl -s https://get.docker.com/ | sudo sh`

우분투 기본 저장소: `sudo apt install docker.io`

사용자를 docker 그룹에 추가함

```sh
sudo usermod -aG docker $USER # 현재 접속중인 사용자에게 권한주기
sudo usermod -aG docker your-user # your-user 사용자에게 권한주기
```

### Windows, MacOS

설치파일을 받아서 실행

## Docker Composer

 설정 관리를 위해 YAML방식의 설정파일 이용

### 설치

리눅스 이외의 환경에서는 기본설치되어 있음

```sh
curl -L "https://github.com/docker/compose/releases/download/1.9.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
docker-compose version
```
