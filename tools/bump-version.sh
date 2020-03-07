#!/bin/bash

# Bump the version of stuff

# https://github.com/Masterminds/semver/blob/910aa146bd66780c2815d652b92a7fc5331e533c/version.go#L41
semver_regex=''
semver_regex+='v?([0-9]+)(\.[0-9]+)?(\.[0-9]+)?'
semver_regex+='(-([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?'
semver_regex+='(\+([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?'

version_file="pkg/version/version.go"

## Get the current version from pkg/version/version.go
# (Is this the correct source of tags?)
# https://stackoverflow.com/questions/1064499
# BASH regex doesn't work for now. grep does
old_version=$(grep -Eo "${semver_regex}" ${version_file})

echo "Old version: ${old_version}"

# Determine semver operation. Fail on garbage input
# Options are minor and patch
case ${1} in
  "minor")
    semver_index=1
    ;;
  "patch")
    semver_index=2
    ;;
  *)
    echo "Invalid input. Use argument 'minor' or 'patch'"
    exit 1
esac

# Split the semver string into an array
# https://askubuntu.com/questions/89995
# https://stackoverflow.com/questions/918886
IFS='.'; read -ra version_arr <<< "${old_version#"v"}"

# Bump part of semver based on user input
let version_arr[${semver_index}]++
# Join the semver array back into a string
# https://superuser.com/questions/461981
new_version="v$(IFS='.'; echo "${version_arr[*]}")"

echo "New version: ${new_version}"

## Bump version in pkg/version/version.go
#-i when I'm ready
sed  -i "s/$old_version/$new_version/g" "${version_file}"

# Stage updated version file, commit, and tag
git add "${version_file}" && \
git commit -m "Bump version ${old_version} -> ${new_version}" && \
git tag "${new_version}"

## Bump version in go.mod?
# go get new tag
# replace line in go.mod
# go mod tidy(?)
