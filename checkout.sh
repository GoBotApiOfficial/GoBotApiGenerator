#!/bin/bash
shopt -s dotglob

git clone -j12 https://github.com/GoBotApiOfficial/gobotapi botapi
# shellcheck disable=SC2164
cd botapi/
for filename in *; do
    case "$filename" in
      "README.md" | "LICENSE.md" | "logo1.svg" | "logo2.svg" | ".github" | ".gitignore" | "examples" | "CODE_OF_CONDUCT.md" | "CONTRIBUTING.md")
        if [ -d "$filename" ]; then
          cp -r "$filename" "../gobotapi/$filename"
        else
          cp "$filename" "../gobotapi/$filename"
        fi
        ;;
    esac
done
