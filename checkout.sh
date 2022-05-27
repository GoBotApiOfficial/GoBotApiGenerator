#!/bin/bash
shopt -s dotglob
git clone https://github.com/Squirrel-Network/gobotapi.git botapi
# shellcheck disable=SC2164
cd botapi/
for filename in *; do
    echo "$filename"
    case "$filename" in
      "README.md" | "LICENSE" | "logo1.png" | "logo2.png" | ".github" | ".gitignore" | "examples" | "CODE_OF_CONDUCT.md" | "CONTRIBUTING.md")
        if [ -d "$filename" ]; then
          cp -r "$filename" "../gobotapi/$filename"
        else
          cp "$filename" "../gobotapi/$filename"
        fi
        ;;
    esac
done