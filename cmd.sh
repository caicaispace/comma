#!/bin/bash
echo "$1" "$2"
if [ "$1" == "git" ]; then
  branch=$(git symbolic-ref --short -q HEAD)
  time=$(date '+%F %T')
  msg='auto push at '${time}
  echo "$branch"
  echo "$msg"
  git add .
  git commit -m "${msg}"
  if [ "$2" == "push" ]; then
    echo "------------ push ${branch}  ------------"
    # git push origin "${branch}"
    git push
  elif [ "$2" == "clear" ]; then
    echo "------------ clear commit ------------"
    remoteUrl=$(git config --get remote.origin.url)
    log "$remoteUrl"
    rm -rf .git
    git init
    git add .
    git commit -am "init"
    git remote add origin "${remoteUrl}"
    git push origin master --force
  fi
fi
