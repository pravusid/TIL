# Git 명령어

## Ignore / Untracking / Remove

- 파일을 인덱스에서 삭제(등록취소)
  - `git rm -r --cached <directory>`
  - `git rm --cached <filename>`

## Add / Staging

## Commit

## Reset / Revert

## Merge / Rebase

## Branch / Checkout

- 브랜치 생성: `git checkout -b <branch>`
- 브랜치 삭제: `git branch --delete <branch>`
- 브랜치 삭제(강제): `git branch -D <branch>`

## Remote

- 생성한 브랜치 원격 저장소 최초 커밋시 push
  - `git push --set-upstream <remote> <branch>`
  - `git push -u <remote> <branch>`

- 원격 저장소 브랜치와 연결
  - `git branch --set-upstream-to <remote>/<branch>`
  - alias `git branch -u <remote>/<branch>`

- 삭제한 로컬 브랜치 원격 저장소 반영: `git push origin :<branch>`
