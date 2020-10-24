TARGET := builder

.PHONY: builder
builder: ./builder.toml pkgs
	pack create-builder my-builder:distroless --config ./builder.toml

pkgs: ./buildpacks/empty.cnb \
      ./buildpacks/python.cnb \
      ./buildpacks/chatty.cnb

buildpacks/empty.cnb: ./buildpacks/empty/buildpack.toml \
	              ./buildpacks/empty/package.toml \
                      ./buildpacks/empty/bin/*
	cd ./buildpacks; \
	pack package-buildpack empty.cnb --config ./empty/package.toml --format file

buildpacks/chatty.cnb: ./buildpacks/chatty/buildpack.toml \
	               ./buildpacks/chatty/package.toml \
	               ./buildpacks/chatty/bin/*
	cd ./buildpacks; \
	pack package-buildpack chatty.cnb --config ./chatty/package.toml --format file

buildpacks/python.cnb: ./buildpacks/python/buildpack.toml \
	               ./buildpacks/python/package.toml \
	               ./buildpacks/python/bin/*
	cd ./buildpacks; \
	pack package-buildpack python.cnb --config ./python/package.toml --format file

.PHONY: build-images
stack:
	cd images; \
	docker build -t distroless:builder-run -f ./Dockerfile.run .; \
	docker build -t distroless:builder-build -f ./Dockerfile.build .

.PHONY: inspect-builder
inspect-builder:
	pack inspect-builder python:distroless

.PHONY: clean
clean:
	rm -f buildpacks/*.cnb
	docker image rm distroless:builder-build
	docker image rm distroless:builder-run
