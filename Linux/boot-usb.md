# Linux Boot USB 생성 in Linux

> [Ventoy](https://www.ventoy.net/en/index.html) 사용 추천

- ISO 파일 다운로드

- 다운로드한 ISO 파일 체크섬 확인 `$ sha1sum -c <linux-image>.iso.sha1`

- USB 장치 경로 확인
  - df 사용: `$ df`
  - ls 사용: `$ ls -l /dev/disk/by-id`
  - fdisk 사용: `$ sudo fdisk -l`

- USB 장치 언마운트
  - `$ sudo umount /device/name`
  - 예를 들어: `$ sudo umount /dev/sdb1`

- [dd](https://ko.wikipedia.org/wiki/Dd_(%EC%9C%A0%EB%8B%89%EC%8A%A4)) 사용하여 boot usb 생성
  - `$ sudo dd if=<linux-image>.iso of=/dev/sdb bs=4M oflag=sync status=progress`
