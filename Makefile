.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


## run/server: run server with go run
.PHONY: run/server
run/server:
	templ generate .
	go run ./cmd/server

## run/generator: run generator to generate html files
.PHONY: run/generator
run/generator:
	templ generate .
	go run ./cmd/generator

## live/templ: start templ proxy
.PHONY: live/templ
live/templ:
	templ generate --watch --proxy="http://localhost:3000" --proxybind="0.0.0.0" --open-browser=false -v

## live/generator: regenerate html files when content changes
.PHONY: live/generator
live/generator:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "go build  -o tmp/bin/main ./cmd/generator" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go,md" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true


## live/tailwind: live reload tailwindcss
.PHONY: live/tailwind
live/tailwind:
	npx --yes @tailwindcss/cli -i ./src/css/input.css -o ./assets/css/tailwind.css --minify --watch

## live/sync_assets: notify proxy on any assets changes
.PHONY: live/sync_assets
live/sync_assets:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.bin "true" \
	--build.cmd "sh -lc 'templ generate --notify-proxy'" \
	--build.delay "400" \
	--build.exclude_dir "" \
	--build.include_dir "assets,public,content" \
	--build.include_ext "js,css,html,md,go"


## live: start live reloading
.PHONY: live
live:
	make -j4 live/templ run/server live/generator live/sync_assets

## build/tailwind: compile tailwind
.PHONY: build/tailwind
build/tailwind:
	npx --yes @tailwindcss/cli -i ./src/css/input.css -o ./assets/css/tailwind.css --minify

## build/server: compile server
.PHONY: build/server
build/server:
	go build -o main ./cmd/server

## build/server: compile server and tailwind
.PHONY: build
build:
	make build/tailwind build/server
