# Git 명령어

<https://git-scm.com/book/ko/v2/Git의-기초-수정하고-저장소에-저장하기>

git이 인식하는 파일 상태는 다음과 같다

- new file(`A`)
- modified(`M`)
- unmodified
- untracked(`??`)

```sh
$ git status -s
 M README               # 내용 변경했으나 Staged에 추가 하지 않음
MM Rakefile             # 변경점을 Staged에 추가하고, 이후 추가로 변경함
A  lib/git.rb           # 새로 추가되어 추적하지 않던 파일을 Staged에 등록하여 추적함
M  lib/simplegit.rb     # 내용을 변경하고 Staged에 추가함
?? LICENSE.txt          # 아직 추적하지 않는 파일
```

Working Directory 각 파일들은 상태에 따라 다음에 위치한다

- Tracked

  - Commited (스냅샷에 포함됨 == 이후 커밋에서 unmodified로 시작)

  - Staged (Staging Area, Changes to be committed)

    - modified
    - new file
    - deleted
    - renamed
    - `...`

  - UnStaged (Changes not staged for commit)

    - modified
    - new file
    - deleted
    - renamed
    - `...`

  - Unmodified

- Untracked

- Ignored (`.gitignore`)

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

### Commit 수정

방법1. autosquash

```sh
git commit --fixup <수정할_커밋_ID>
git rebase --autosquash
```

방법2. interactive rebase

```sh
git rebase -i <수정할_커밋_직전커밋_ID>

# 수정할 커밋을 edit으로 변경하고 저장
# 체크아웃 상태에서 코드 수정

git commit --amend
git rebase --continue
```

## Tag

- 목록

  - 태그 목록: `git tag`
  - 태그 검색: `git tag -l '검색어'` i.e. `git tag -l 'v1.4.2.*'`

- 생성

  - Lightweight 태그 붙이기: `git tag <tag>`
  - 이전 커밋에 태그 붙이기: `git tag <tag> <commit_checksum>`
  - Annotated 태그 붙이기: `git tag -a <tag> -m '<message>'`
  - 서명한 태그 붙이기: `git tag -s <tag> -m '<message>'`

- 검증

  - 서명한 태그 검증: `git tag -v [태그 이름]` (서명자의 공개키가 Keyring에 있어야 함)

- 삭제

  - 태그 삭제: `git tag -d <tag>`

## Reflog

- <https://git-scm.com/docs/git-reflog>
- <https://git-scm.com/book/ko/v2/Git의-내부-운영-및-데이터-복구>

> Reference logs, or "reflogs", record when the tips of branches and other references were updated in the local repository.
> Reflogs are useful in various Git commands, to specify the old value of a reference.
> For example, HEAD@{2} means "where HEAD used to be two moves ago"

- reflog 출력: `git reflog [ref:-HEAD]`
- reflog 남기기: `git update-ref`

## Reset

<https://git-scm.com/book/ko/v2/Git-도구-Reset-명확히-알고-가기>

- commit level: `git reset [--soft|--hard] [commit:-HEAD]`
- file level: `git reset [commit:-HEAD] [paths]`

reset 명령은 정해진 순서대로 세 개의 트리를 덮어써 나가다가 옵션에 따라 지정한 곳에서 멈춘다.

- HEAD가 가리키는 브랜치를 옮긴다. (--soft 옵션이 붙으면 여기까지)
- Index를 HEAD가 가리키는 상태로 만든다. (--hard 옵션이 붙지 않았으면 여기까지)
- 워킹 디렉토리를 Index의 상태로 만든다.

## Merge

development 브랜치를 main 브랜치로 merge

```sh
git switch main
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

`RE-BASE`는 커밋의 부모커밋을 변경한다는 개념이다

즉, A 브랜치로부터 생성된 B브랜치에서 작업 이후 A브랜치를 대상으로 rebase를 실행한다면,
B브랜치를 생성한 A브랜치의 커밋 이후, A 브랜치에서 발생한 변경점을 B브랜치에 적용한다는 것이다

(A브랜치의 마지막 커밋을 B브랜치의 생성지점으로 변경).

`git rebase main development`

또는

```sh
git switch development
git rebase main
```

실행 결과는 다음과 같다

```txt
Before                              After
A---B---C---F---G (main)            A---B---C---F---G (main)
         \                                           \
          D---E (HEAD development)                   D'---E' (HEAD development)
```

rebase 진행도중 conflict가 발생한다면, merge시 conflict 해결과 같은 방법을 적용한 뒤

`git rebase --continue` 명령어를 입력한다

이후 development 브랜치는 main 브랜치로 이동하여 `FF` 가능하므로

```sh
git switch main
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

#### 특정커밋에 적용

<https://fle.github.io/git-tip-keep-your-branch-clean-with-fixup-and-autosquash.html>

```sh
git commit --fixup <commit_hash> #특정 커밋에 변경사항을 적용한다는 표시를 해둔다
git rebase -i --autosquash #fixup, squash 표시한 커밋을 자동 적용한다
```

### `git rebase --onto`

기본 rebase 명령어

- 지정한 브랜치의 도달할 수 있는 마지막 커밋 -> 현재 브랜치(HEAD 위치 브랜치)의 **base로** 설정한다

<https://womanonrails.com/git-rebase-onto>

#### `--onto` with 2 args

2개의 인자를 사용하는 onto rebase 명령어

- `git rebase --onto <newparent_commit> <oldparent_commit>`
- `newparent_commit` -> `oldparent_commit.child ~ HEAD`의 **base로** 설정한다
- 세 번째 인자로 현재 브랜치를 사용하면, 2개의 인자를 사용한 것과 동일하게 작동한다

`git rebase --onto F D` == `git rebase --onto F D my-branch`

실행 결과는 다음과 같다

```txt
Before                                    After
A---B---C---F---G (branch)                A---B---C---F---G (branch)
         \                                             \
          D---E---H---I (HEAD my-branch)                E'---H'---I' (HEAD my-branch)
```

#### `--onto` with 3 args

3개의 인자를 사용하는 onto rebase 명령어

- `git rebase --onto <newparent_commit> <oldparent_commit> <until_commit>`
- 브랜치(2개의 인자 사용시) 대신 커밋이 세 번째 인자로 전달하고, 브랜치의 base가 아닌 해당 커밋 hierarchy base가 변경된다
- `oldparent_commit.child ~ until_commit` 범위로 새로운 detached 커밋 hierarchy가 생성하고 `newparent_commit` 커밋을 base로 삼는다

`git rebase --onto F D H`

```txt
Before                                    After
A---B---C---F---G (branch)                A---B---C---F---G (branch)
         \                                        |    \
          D---E---H---I (HEAD my-branch)          |     E'---H' (HEAD)
                                                   \
                                                    D---E---H---I (my-branch)
```

### rebase를 통한 원격 작업 예시

작업시작을 위한 명령

```sh
git switch develop
git pull --rebase=preserve origin develop
git switch -c feature-foobar
# 작업을 한다
git add --all
git commit
```

원격 저장소에 업데이트 된 내용이 있을 수 있으므로 우선 동기화하고 rebase 실행

> push 이후에는 rebase 하지 않는다

```sh
git switch develop
git pull --rebase=preserve
git switch feature-foobar
git rebase develop
```

rebase 이후에는 PR 생성

```sh
git push origin feature-foobar
# PR & 코드 리뷰 ...
```

또는 직접 merge 실행

```sh
git merge --no-ff feature-foobar main
```

과정이 끝나면 branch는 닫는다

```sh
git branch --delete feature-foobar
# PR시 branch 삭제하거나 직접 삭제
git push --delete origin feature-foobar
```

## Branch / Checkout (Switch & Restore)

- 로컬 브랜치 목록: `git branch`
- 원격 브랜치 목록: `git branch -r`
- 전체 브랜치 목록: `git branch -a`
- 브랜치 삭제: `git branch --delete(-d) <branch>`
- 브랜치 (강제)삭제: `git branch -D <branch>`
- 브랜치 이름 변경: `git branch -m <before> <after>`
- 병합된 로컬 브랜치 모두 삭제: `git branch --merged | egrep -v "(^\*|main|development|제외할브랜치)" | xargs git branch -d`

> git 2.23 버전부터 checkout 명령이 switch, restore 명령으로 분리되었다

<https://git-scm.com/docs/git-switch>

- 브랜치 전환: `git switch <branch>`
- 브랜치 생성: `git switch -c <branch> [from-commit]`
- HEAD 이동 (detach): `git switch -d [<start-point>]`

<https://git-scm.com/docs/git-restore>

- 되돌리기(unstaged 파일 변경점 되돌리기, `--worktree` default true): `git restore <filename>`
- 되돌리기(staged -> unstaged): `git restore --staged <filename>`

## Sparse Checkout

<https://git-scm.com/docs/git-sparse-checkout>

> git 2.25 버전에 추가됨

대규모 repository에서 (특히 monorepo) 일부분만 checkout 할 때 사용

```bash
git clone --no-checkout git@github.com:pravusid/TIL.git
cd TIL

git sparse-checkout init --cone
git sparse-checkout set <path> [...paths]
git sparse-checkout list
```

## Clone

- <https://git-scm.com/docs/git-clone/en>
- <https://git-scm.com/docs/git-clone/en#URLS>

```sh
git clone [--branch <name>] [--depth n] <git-url> [dir-to-clone-into]

# branch → 특정 브랜치(태그)를 가져온다
# depth  → 마지막 n개의 커밋만 가져오는 shallow clone 실행
#          커밋을 추가로 가져오려면 git fetch --depth n --force 실행
```

## Remote

- 생성한 브랜치 원격 저장소 최초 커밋시 push

  - `git push --set-upstream <remote> <branch>`
  - `git push -u <remote> <branch>`

- 원격 저장소 브랜치와 연결

  - `git branch --set-upstream-to <remote>/<branch>`
  - alias `git branch -u <remote>/<branch>`

### Fetch

remote에서 데이터 가져오기(local에 반영하지는 않음): `git fetch`

remote에서 삭제한 branch 로컬 반영(로컬에서도 브랜치 삭제): `git fetch --all(-a) --prune(-p)`

- update all remote references (`--all`)
- drop deleted ones (`--prune`)

### Pull

`git pull`

`git pull --rebase`

#### 다른 브랜치 pull

> HEAD 위치가 아닌 다른 브랜치 pull

```sh
git fetch <remote> <local-branch>:<remote-branch>
git pull <remote> <local-branch>:<remote-branch>
```

#### pull force

<https://stackoverflow.com/questions/1125968/how-do-i-force-git-pull-to-overwrite-local-files>

```sh
git fetch origin main
git reset --hard origin/main
```

### Push

#### Remote Branch

branch push

`git push <remote> <branch>`

branch 삭제

`git push --delete origin feature-foobar`

`git push origin :<branch>`

> alias `[empty-local-branch]:[remote-branch]` === 비어있는 local-branch(empty reference)를 원격의 remote-branch로 push

#### Remote Tag

- Remote에 태그 Push: `git push <origin> <태그이름>`

- Remote에 없는 태그 모두 Push: `git push origin --tags`

- Remote 태그 삭제

  - `git push --delete origin tagname`
  - `git push origin :tagname`

## `.gitignore`

- 아무것도 없는 라인이나, `#`로 시작하는 라인은 무시
- 프로젝트 전체에 적용되는 표준 Glob 패턴을 사용
- 슬래시(`/`)로 시작하면 하위 디렉토리에 적용되지(Recursivity) 않음
- 디렉토리는 슬래시(`/`)를 끝에 사용하는 것으로 표현
- 느낌표(`!`)로 시작하는 패턴의 파일은 무시하지 않음

## 특정 기록 완전 삭제 (데이터 손상 위험)

<https://git-scm.com/docs/git-filter-branch#_examples>

> 파일 이동이 있었다면, 파일이 존재했던 모든 경로의 기록을 삭제 해야함

```sh
git filter-branch --force --index-filter \
  'git rm --cached --ignore-unmatch <URL_TO_FILE_OR_DIR>' \
  --prune-empty --tag-name-filter cat -- --all
```

git repo 정리

```sh
rm -rf .git/refs/original/
git reflog expire --expire=now --all
git gc --prune=now
git gc --aggressive --prune=now
```

원격 저장소 반영

```sh
git push --all --force
```

## Bisect

<https://git-scm.com/docs/git-bisect>

이진탐색을 통해 오류지점을 찾는 기능

```bash
git bisect start
git bisect bad HEAD
git bisect good 246320d

# 이동한 위치의 결과 입력 (둘 중 하나)
git bisect bad
git bisect good

# bisect 종료
git bisect reset
```

## Troubleshootings

<https://github.com/k88hudson/git-flight-rules>

> 뭔가 잘못 됐을 때 뭘 해야할지에 대한 우주비행사를 위한 가이드 (여기선 깃을 쓰는 개발자를 위한)
