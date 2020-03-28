# PACMAN | YAY

- pacman: Arch Linux 패키지 관리자
- yay: AUR helper <https://github.com/Jguer/yay>

## 사용법

<https://wiki.archlinux.org/index.php/Pacman/Rosetta>

> 로그: `/var/log/pacman.log`

### 캐시삭제

- `pacman -Sc`
- `yay -Sc --aur`

## Troubleshooting

### `exists on filesystem` error

```sh
sudo pacman -S --overwrite \* <패키지>
```

### 깨진 의존성 제거

`sudo pacman -R $(pacman -Qdtq)`
