# oh-my-zsh alias

## 디렉토리

```sh
alias -g ...='../..'
alias -g ....='../../..'
alias -g .....='../../../..'
alias -g ......='../../../../..'

alias -- -='cd -'
alias 1='cd -'
alias 2='cd -2'
alias 3='cd -3'
alias 4='cd -4'
alias 5='cd -5'
alias 6='cd -6'
alias 7='cd -7'
alias 8='cd -8'
alias 9='cd -9'

alias md='mkdir -p'
alias rd='rmdir'
alias d='dirs -v | head -10'

# List directory contents
alias l='ls -lah'
alias ll='ls -lh'
alias la='ls -lAh'
```

## sudo

```sh
alias _='sudo'
```

## git

- `ga`: `git add`
- `gaa`: `git add --all`

- `gcmsg`: `git commit -m`

- `gb`: `git branch`
- `gba`: `git branch -a`
- `gbr`: `git branch --remote`
- `gbd`: `git branch -d`
- `gbD`: `git branch -D`
- `gbnm`: `git branch --no-merged`
- `gbda`: `git branch --no-color --merged | command grep -vE "^(*|\s*(master|develop|dev)\s*$)" | command xargs -n 1 git branch -d`

- `gco`: `git checkout`
- `gcd`: `git checkout develop`
- `gcm`: `git checkout master`
- `gcb`: `git checkout -b`

- `gd`: `git diff`
- `gdv`: `git diff -w $@ | view -`

- `grv`: `git remote -v`

- `gp`: `git push`
- `gpv`: `git push -v`
- `gpd`: `git push --dry-run`
- `gpf`: `git push --force-with-lease`
- `gpf!`: `git push --force`
- `gpsup`: `git push --set-upstream origin $(git_current_branch)`

- `gf`: `git fetch`
- `gfa`: `git fetch --all --prune`

- `gl`: `git pull`

- `gm`: `git merge`
- `gmtvim`: `git mergetool --no-prompt --tool=vimdiff`

- `grb`: `git rebase`
- `grbi`: `git rebase -i`
- `grbd`: `git rebase develop`
- `grbm`: `git rebase master`

- `grh`: `git reset`
- `grhh`: `git reset --hard`

- `grm`: `git rm`
- `grmc`: `git rm --cached`
- `gclean`: `git clean -id`

- `gst`: `git status`
- `gsb`: `git status -sb`

- `glols`: `git log --graph --pretty='%Cred%h%Creset -%C(auto)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --stat`
- `glola`: `git log --graph --pretty='%Cred%h%Creset -%C(auto)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --all`
