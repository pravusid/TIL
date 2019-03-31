# Git 명령어

git이 인식하는 파일 상태는 다음과 같다

- unmodified
- new file(`A`)
- modified(`M`)
- untracked(`??`)

각 파일들이 상태에 따라 위치할 수 있는 영역은 다음곽 같다

- Commited (스냅샷에 포함됨 == unmodified가 됨)
- Staging Area (Changes to be committed)
- **Working Directory** (Changes not staged for commit)
- Ignored (`.gitignore`에 선언됨)

```sh
$ git status -s
 M README               # 내용 변경했으나 Staged에 추가 하지 않음
MM Rakefile             # 변경점을 Staged에 추가하고, 이후 추가로 변경함
A  lib/git.rb           # 새로 추가되어 추적하지 않던 파일을 Staged에 등록하여 추적함
M  lib/simplegit.rb     # 내용을 변경하고 Staged에 추가함
?? LICENSE.txt          # 아직 추적하지 않는 파일
```

## Ignore / Untracking / Remove

- 파일을 인덱스에서 삭제: Staging Area에서만 제거하고 워킹 디렉토리에 있는 파일은 지우지 않고 남겨둠
  - `git rm -r --cached <directory>`
  - `git rm --cached <filename>`

## Add / Staging

- Staging: `git add <filename>`
- Staging All: `git add .`

## Commit

- Staging Area의 파일 커밋(메시지 포함): `git commit -m '<message>'`
- Staging Area의 데이터를 마지막 커밋에 추가: `git commit --amend`
- Tracking 파일 전체를 Staging 생략하고 커밋: `git commit -a`

## Log / Reflog

- reflog 출력: `git reflog`
- reflog 남기기: `git update-ref`

## Reset / Revert

- Unstaging: `git reset HEAD <filename>`

## Merge

development 브랜치를 master 브랜치로 merge

```sh
git checkout master
git merge development
```

merge는 `--ff`(fast foward)가 기본 설정이다

- `git merge --no-ff`
  - ff가 가능하더라도 3-way-merge 실행

- `git merge --squash`
  - 대상 브랜치로부터 merge할 내용만 반영하고 merge는 실행하지 않음
  - 즉, 대상 브랜치의 모든 변경점을 하나로 합쳐서 병합할 브랜치 HEAD에 반영함

merge시 발생하는 conflict 해결

- 수동
- `git mergetool`

## Rebase

rebase는 RE-BASE로 커밋의 부모커밋을 변경한다는 개념이다

즉, A 브랜치로부터 생성된 B브랜치에서 작업 이후 A브랜치를 대상으로 rebase를 실행한다면,
B브랜치를 생성한 A브랜치의 커밋 이후, A 브랜치에서 발생한 변경점을 B브랜치에 적용한다는 것이다
(현재 A브랜치의 HEAD를 B브랜치의 생성지점으로 변경).

`git rebase master development`

또는

```sh
git checkout development
git rebase master
```

rebase 진행도중 conflict가 발생한다면, merge시 conflict 해결과 같은 방법을 적용한 뒤

`git rebase --continue` 명령어를 입력한다

이후 development 브랜치는 master HEAD로부터 ff가 가능하므로

```sh
git checkout master
git merge development
```

### 대화형 rebase: `git rebase -i <수정을_시작할_커밋_직전커밋>`

명령을 실행하면 커밋 로그가 날짜순 오름차순으로 기록된 에디터가 실행되고 아래의 명령을 지정할 수 있다.

- pick: 해당 커밋 사용
- reword: 커밋 메시지 변경
- edit: 커밋 메시지 및 내용 변경
- squash: 커밋을 직전 커밋과 합친다
- fixup: 커밋을 직전 커밋과 합치지만 메시지는 합치지 않는다(직전 커밋의 메시지만 남긴다)
- drop: 커밋 히스토리에서 해당 커밋을 삭제함

해당 커밋에 적용할 명령을 쓴 뒤 에디터를 저장하고 나가면 작업을 수행할 수 있다

필요한 작업을 모두 수행하였다면 [`git commit --amend`], `git rebase --continue` 실행

### `git rebase --interactive --autosquash`

작업이 모두 끝나고 커밋 히스토리를 정리할 계획을 갖고 있다면 미리 squash나 fixup 할 커밋에 메시지를 표기해둘 수 있다

커밋 메시지 접두사로 `squash!` 또는 `fixup!`을 사용한다면
대화형 rebase가 실행됨과 동시에 해당 커밋들은 squash와 fixup 상태인 것을 볼 수 있다.

### `git rebase --onto`

#### `--onto` with 2 args

rebase는 HEAD에 의해서 참조된 브랜치의 새로운 base를 대상 브랜치의 도달할 수 있는 마지막 커밋으로 변경한다.

마지막 커밋 대신 특정 커밋을 참조하게 하려면 2개의 인자를 사용하는 `--onto` 옵션을 사용한다

`git rebase --onto <newparent> <oldparent>`

#### `--onto` with 3 args

3개의 인자를 사용하는 `rebase --onto`는 임의의 범위 커밋들의 base 커밋을 변경할 수 있다

`git rebase --onto <newparent> <oldparent> <until>`

## Branch / Checkout

- 로컬 브랜치 목록: `git branch -a`
- 원격 브랜치 목록: `git branch -r`
- 브랜치 생성: `git checkout -b <branch>`
- 브랜치 삭제: `git branch --delete <branch>`
- 브랜치 삭제(강제): `git branch -D <branch>`
- 브랜치 이름 변경: `git branch -m <before> <after>`
- 체크아웃(Not staged 파일 변경점 되돌리기): `git checkout -- <filename>`

## Remote

- 생성한 브랜치 원격 저장소 최초 커밋시 push
  - `git push --set-upstream <remote> <branch>`
  - `git push -u <remote> <branch>`

- 원격 저장소 브랜치와 연결
  - `git branch --set-upstream-to <remote>/<branch>`
  - alias `git branch -u <remote>/<branch>`

- 삭제한 로컬 브랜치 원격 저장소 반영: `git push origin :<branch>`

### Fetch

remote에서 데이터 가져오기(합치지는 않음): `git fetch`

### Pull

`git pull`

`git pull --rebase`

### Push

## `.gitignore`

- 아무것도 없는 라인이나, `#`로 시작하는 라인은 무시
- 프로젝트 전체에 적용되는 표준 Glob 패턴을 사용
- 슬래시(`/`)로 시작하면 하위 디렉토리에 적용되지(Recursivity) 않음
- 디렉토리는 슬래시(`/`)를 끝에 사용하는 것으로 표현
- 느낌표(`!`)로 시작하는 패턴의 파일은 무시하지 않음
