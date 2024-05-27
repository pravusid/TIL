# Swapfile 생성

> <https://aws.amazon.com/ko/premiumsupport/knowledge-center/ec2-memory-swap-file/>

```sh
# dd 명령을 사용하여 루트 파일 시스템에 스왑 파일생성. "bs"는 블록 크기이고 "count"는 블록 수
# 지정한 블록 크기가 인스턴스에서 사용 가능한 메모리보다 크면 "memory exhausted" 오류가 발생
# 이 예제 dd 명령에서 스왑 파일은 4GB(128MB x 32)
sudo dd if=/dev/zero of=/swapfile bs=128M count=32

# 스왑 파일의 읽기 및 쓰기 권한을 업데이트
sudo chmod 600 /swapfile

# Linux 스왑 영역을 설정
sudo mkswap /swapfile

# 스왑 공간에 스왑 파일을 추가하여 스왑 파일을 즉시 사용할 수 있도록 함
sudo swapon /swapfile

# 프로시저가 성공적인지 확인
sudo swapon -s
```

`/etc/fstab` 파일을 편집하여 부팅 시 스왑 파일을 활성화

```txt
/swapfile swap swap defaults 0 0
```

## Swappiness 설정

<https://wiki.archlinux.org/title/swap#Swappiness>
