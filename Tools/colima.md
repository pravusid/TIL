# CoLima

<https://github.com/abiosoft/colima>

## 사용방법

- 시작 (기본설정)
  - `colima start --cpu 1 --memory 2 --disk 10`
  - `colima start --edit --disk 30`
- 시작 (추가설정)
  - `colima start --cpu 1 --memory 2 --disk 30 --arch aarch64 --vm-type vz --vz-rosetta --mount-type virtiofs`
- 자동실행
  - `brew services start colima`
- 오류발생시
  - `limactl stop -f colima`
  - <https://github.com/abiosoft/colima/issues/381>
- 볼륨 연결했을 때 permission 오류 발생하는 경우
  - `시스템 설정 → 개인정보 보호 및 보안 → 파일 및 폴더`에서 볼륨이 포함된 위치 권한 추가
- options
  - <https://lima-vm.io/docs/reference/limactl_start/>
  - `--arch string` [limactl create] (virtual) machine architecture (x86_64, aarch64, riscv64)
  - `--vm-type string` [limactl create] virtual machine type (qemu, vz)
  - `--rosetta (vz-rosseta)` [limactl create] enable Rosetta (for vz instances)
  - `--mount-type string` [limactl create] mount type (reverse-sshfs, 9p, virtiofs)
    - <https://virtio-fs.gitlab.io/>

## 참고

- <https://x.com/darjeelingt/status/1816702772047413309?s=46&t=zQaFJ1iBc-ZF5GKv5CVZpQ>
- <https://shubham.codes/blog/2023-08-18-analysis-docker-dsktop-colima-on-mac-m1/>
- <https://medium.com/@guillem.riera/the-most-performant-docker-setup-on-macos-apple-silicon-m1-m2-m3-for-x64-amd64-compatibility-da5100e2557d>
- <https://dev.to/dkechag/improving-docker-file-performance-on-macos-537k>
- <https://www.tyler-wright.com/blog/using-colima-on-an-m1-m2-mac/>
