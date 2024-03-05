# Go Implementation

## Install SQLite3 on Ubuntu

```sh
sudo apt update && sudo apt upgrade -y
```

```sh
sudo apt install sqlite3
```

```sh
sqlite3 --version
```

## Bazel

### Buildozer

```sh
buildozer 'use_repo_add @gazelle//:extensions.bzl go_deps com_github_tailscale_sqlite' //MODULE.bazel:all
```

### Gotchas

### Remove SQLite database after 'bazel run'

```sh
find . -name "sqlite.db"
```

```sh
rm ./src/backend_brawl_/backend_brawl.runfiles/_main/sqlite*.*
```
