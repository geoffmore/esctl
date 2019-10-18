#!/bin/bash

# Argument count check
if [[ ${#@} -lt 1 ]]; then
  echo "Insufficient number of arguments for command"
  exit 1
fi

PKG="${1}"
OUTPUT_BIN="tmp-esctl"

# Argument validation
grep -E "^[a-z]{3,}$"  <<< "$PKG"
if [[ "$?" -ne 0 ]]; then
  echo "Package name needs to be lower case, with no under_scores or mixedCaps (https://blog.golang.org/package-names) and have at least 3 characters"
  exit 1
fi

echo "Attemping to template files..."
# Add the pkg to pkg dir
 mkdir -p pkg/${PKG}
# Add the template value to file
echo "pkg: ${PKG}" > pkg.yml
# Create and place template files (safely)
set -o noclobber
gotpl build/extension/cmd.tmpl   < pkg.yml > cmd/${PKG}.go
gotpl build/extension/main.tmpl  < pkg.yml > pkg/${PKG}/${PKG}.go
gotpl build/extension/types.tmpl < pkg.yml > pkg/${PKG}/types.go
rm pkg.yml
set +o noclobber
echo "Done!"


# Test the new package in a build
echo "Building package with new module..."
go build -tags ${PKG} -o "./${OUTPUT_BIN}" || echo "Compilation was not successful with new package"

# Test functionality with commands
# test-int assumes connectivity to an existing cluster. This is a bad test and
# should rely on default configuration and a test container. Without the
# context feature, this may not be possible without moving files around
"./${OUTPUT_BIN}" ${PKG} test-cmd > /dev/null 2>&1 \
  && "./${OUTPUT_BIN}" ${PKG} test-int > /dev/null 2>&1

case $? in
  0)
    echo "Tests were successful. Enjoy extending esctl!"
    ;;
  *)
    echo "Tests were not successful. Further debugging is needed..."
    ;;
esac

# Clean up (remove binary)
rm "./${OUTPUT_BIN}"
