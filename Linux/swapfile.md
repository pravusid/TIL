# Swapfile 생성

> <https://aws.amazon.com/ko/premiumsupport/knowledge-center/ec2-memory-swap-file/>

```sh
# dd 명령을 사용하여 루트 파일 시스템에 스왑 파일생성. "bs"는 블록 크기이고 "count"는 블록 수
sudo dd if=/dev/zero of=/swapfile bs=1G count=4

# 스왑 파일의 읽기 및 쓰기 권한을 업데이트
chmod 600 /swapfile

# Linux 스왑 영역을 설정
mkswap /swapfile

# 스왑 공간에 스왑 파일을 추가하여 스왑 파일을 즉시 사용할 수 있도록 함
swapon /swapfile

# 프로시저가 성공적인지 확인
swapon -s
```

`/etc/fstab` 파일을 편집하여 부팅 시 스왑 파일을 활성화

```txt
/swapfile swap swap defaults 0 0
```
