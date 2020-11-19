# README
```
$ pack build \
  --env "GO_LINKER_SYMBOL=main.Commit" \
  --env "GO_LINKER_VALUE=`git rev-parse HEAD`" \
  myapp --builder heroku/buildpacks:18
```
