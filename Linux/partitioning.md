# Partitioning

개인 사용자용 리눅스(server 배포판이 아닌 desktop 배포판) 기준

서버 배포판의 경우에는 `/var`, `/tmp` 등의 파티션을 고려해야 한다

## 나누기를 권장하는 파티션

- `swap`: 아래 참조

- `/boot`: 512mb ~ 1gb

  - legacy bios
    - /boot: 512mb
  - UEFI
    - /boot: 512mb
    - /boot/efi: 200mb (파일시스템을 EFI System Partition으로 설정)

- `/`: 최소 5gb (보통 50gb 정도를 할당하였음)

- `/home`: 나머지 공간

## swap

- 4GB 이하 RAM: 최소 2GB의 스왑 공간
- 4GB에서 16GB RAM: 최소 4GB 스왑 공간
- 16GB에서 64GB RAM: 최소 8GB 스왑 공간
- 64GB에서 256GB RAM: 최소 16GB 스왑 공간
- 256GB에서 512GB RAM: 최소 32GB 스왑 공간

### 스왑 파일 생성

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

# /etc/fstab 파일을 편집하여 부팅 시 스왑 파일을 활성화
vim /etc/fstab
/swapfile swap swap defaults 0 0
```
