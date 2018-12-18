#!/usr/bin/env sh

# abort on errors
set -e

rm -rf docs/*
cp -r ../* docs

node create.index.js
cat index.md >> docs/README.md

# build
npm run docs:build

# navigate into the build output directory
cd docs/.vuepress/dist

git init
git add -A
git commit -m 'deploy'

git push -f git@github.com:pravusid/wiki.git master:gh-pages

cd -
