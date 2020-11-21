# README
## heroku
<https://github.com/heroku/heroku-buildpack-go#passing-a-symbol-and-optional-string-to-the-linker>
```
$ pack build \
  --env "GO_LINKER_SYMBOL=main.Commit" \
  --env "GO_LINKER_VALUE=`git rev-parse HEAD`" \
  myapp:heroku --builder heroku/buildpacks:18
```

## Google
<https://github.com/GoogleCloudPlatform/buildpacks#go-buildpacks>
```
$ pack build \
  --env "GOOGLE_GOLDFLAGS=-X main.Commit=`git rev-parse HEAD`" \
  myapp:google --builder gcr.io/buildpacks/builder
```

<https://paketo.io/docs/buildpacks/language-family-buildpacks/go/>
```
$ docker run --rm myapp:paketo-tiny  # aware of buildpack.yml
```
