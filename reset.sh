#!/bin/bash
git fetch origin
git push -d origin $(git branch -r -l |cut -c 3- | sed -e 's@origin/@refs/heads/@g' | grep -v upstream | grep -v main | grep -v patch | grep -v base)
