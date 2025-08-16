.PHONY: run
run:
	templ generate .
	go run ./cmd/server

.PHONY: run/generator
run/generator:
	templ generate .
	go run ./cmd/generator

.PHONY: live/templ
live/templ:
	templ generate --watch --proxy="http://localhost:3000" --proxybind="0.0.0.0" --open-browser=false -v

.PHONY: live/server
live/server:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "go build  -o tmp/bin/main ./cmd/server" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

.PHONY: live/tailwind
live/tailwind:
	npx --yes @tailwindcss/cli -i ./src/css/input.css -o ./assets/css/tailwind.css --minify --watch

.PHONY: live/sync_assets
live/sync_assets:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "assets" \
	--build.include_ext "js,css,md"

.PHONY: live
live:
	make -j3 live/templ live/server live/sync_assets

.PHONY: build/tailwind
build/tailwind:
	npx --yes @tailwindcss/cli -i ./src/css/input.css -o ./assets/css/tailwind.css --minify

.PHONY: build/server
build/server:
	go build -o main ./cmd/server

.PHONY: build
build:
	make build/tailwind build/server
