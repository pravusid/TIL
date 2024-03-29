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
    - /boot/efi: 200mb, 파일시스템을 fat32 & ESP(EFI System Partition)로 설정

- `/`: 최소 5gb (보통 50gb 정도를 할당하였음)

- `/home`: 나머지 공간

## swap

- 4GB 이하 RAM: 최소 2GB의 스왑 공간
- 4GB에서 16GB RAM: 최소 4GB 스왑 공간
- 16GB에서 64GB RAM: 최소 8GB 스왑 공간
- 64GB에서 256GB RAM: 최소 16GB 스왑 공간
- 256GB에서 512GB RAM: 최소 32GB 스왑 공간

### 스왑 파일 생성

[[swapfile]]
