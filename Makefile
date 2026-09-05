TAILWIND_IN=./static/css/input.css
TAILWIND_OUT=./static/css/output.css
MAIN_PATH=./cmd/web/main.go

.PHONY: dev build watch-templ watch-tailwind run-app clean

dev:
	./tailwindcss -i $(TAILWIND_IN) -o $(TAILWIND_OUT)
	@make -j 3 watch-templ watch-tailwind run-app

watch-templ:
	templ generate --watch

watch-tailwind:
	./tailwindcss -i $(TAILWIND_IN) -o $(TAILWIND_OUT) --watch
	./tailwindcss -i $(TAILWIND_IN) -o $(TAILWIND_OUT) --content "./cmd/web/views/**/*.templ" --watch

run-app:
	air

build:
	templ generate
	./tailwindcss -i $(TAILWIND_IN) -o $(TAILWIND_OUT) --minify
	go build -o bin/app $(MAIN_PATH)

clean:
	rm -rf bin tmp $(TAILWIND_OUT)