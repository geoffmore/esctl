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
version=$(grep -Eo "${semver_regex}" ${version_file})
echo ${version}
#if [[ ${version_file} =~ $semver_regex ]]; then
#  version="${BASH_REMATCH[1]}"
#  
#  echo ${version}
#fi

# Determine semver operation. Fail on garbage input
# Options are minor and patch
case ${1} in
  "minor")
    semver_op="minor"
    ;;
  "patch")
    semver_op="patch"
    ;;
  *)
    echo "Invalid input. Use argument 'minor' or 'patch'"
    exit 1
esac


# minor
# patch

## Bump version in pkg/version/version.go

## Bump version in go.mod?

## Make an empty commit "bump version to {new_version}"
#git commit --allow-empty -m "Bump version to ${NEW_VERSION}"
#git tag "${NEW_VERSION}"
