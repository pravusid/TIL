# oh-my-zsh alias

## 디렉토리

<https://github.com/ohmyzsh/ohmyzsh/blob/master/lib/directories.zsh>

```sh
alias -g ...='../..'
alias -g ....='../../..'
alias -g .....='../../../..'
alias -g ......='../../../../..'

alias -- -='cd -'
alias 1='cd -1'
alias 2='cd -2'
alias 3='cd -3'
alias 4='cd -4'
alias 5='cd -5'
alias 6='cd -6'
alias 7='cd -7'
alias 8='cd -8'
alias 9='cd -9'

alias md='mkdir -p'
alias rd=rmdir

function d ()

# List directory contents
alias lsa='ls -lah'
alias l='ls -lah'
alias ll='ls -lh'
alias la='ls -lAh'
```

## clipboard

<https://github.com/ohmyzsh/ohmyzsh/blob/master/lib/clipboard.zsh>

- `<command> | clipcopy` : copies stdin to clipboard
- `clipcopy <file>` : copies a file's contents to clipboard
- `clippaste` : writes clipboard's contents to stdout
- `clippaste | <command>` : pastes contents and pipes it to another process
- `clippaste > <file>` : paste contents to a file

## sudo

<https://github.com/ohmyzsh/ohmyzsh/blob/master/lib/misc.zsh>

```sh
alias _="sudo"
```

## git

<https://github.com/ohmyzsh/ohmyzsh/blob/master/plugins/git/git.plugin.zsh>

- add

  - `ga`: `git add`
  - `gaa`: `git add --all`

- commit

  - `gcmsg`: `git commit -m`

- branch

  - `gb`: `git branch`
  - `gba`: `git branch -a`
  - `gbr`: `git branch --remote`
  - `gbd`: `git branch -d`
  - `gbD`: `git branch -D`
  - `gbnm`: `git branch --no-merged`
  - `gbda`: `git branch --no-color --merged | command grep -vE "^(*|\s*($(git_main_branch)|develop|dev)\s*$)" | command xargs -n 1 git branch -d`

- switch

  - `gsw`: `git switch`
  - `gswc`: `git switch -c`

- restore

  - `grs`: `git restore`
  - `grss`: `git restore --source`
  - `grst`: `git restore --staged`

- diff

  - `gd`: `git diff`
  - `gdv`: `git diff -w $@ | view -`

- remote

  - `grv`: `git remote -v`

- push

  - `gp`: `git push`
  - `gpv`: `git push -v`
  - `gpd`: `git push --dry-run`
  - `gpf`: `git push --force-with-lease`
  - `gpf!`: `git push --force`
  - `gpsup`: `git push --set-upstream origin $(git_current_branch)`

- fetch

  - `gf`: `git fetch`
  - `gfa`: `git fetch --all --prune`

- pull

  - `gl`: `git pull`

- merge

  - `gm`: `git merge`
  - `gmtvim`: `git mergetool --no-prompt --tool=vimdiff`

- rebase

  - `grb`: `git rebase`
  - `grbi`: `git rebase -i`
  - `grbd`: `git rebase develop`
  - `grbm`: `git rebase $(git_main_branch)`

- rm / clean

  - `grm`: `git rm`
  - `grmc`: `git rm --cached`
  - `gclean`: `git clean -id`

- status

  - `gst`: `git status`
  - `gsb`: `git status -sb`

- stash

  - `gsta`: `git stash save`
  - `gstp`: `git stash pop`
  - `gstd`: `git stash drop`
  - `gstc`: `git stash clear`
  - `gstl`: `git stash list`
  - `gsts`: `git stash show --text`

- log

  - `glols`: `git log --graph --pretty="%Cred%h%Creset -%C(auto)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset" --stat`
  - `glola`: `git log --graph --pretty="%Cred%h%Creset -%C(auto)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset" --all`
