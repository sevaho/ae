NAME := ae

help: ## Show this help message (default)]
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## build
	go build -o "$(NAME)"

build: ## build
	go build -o "$(NAME)"

install: ## build
	go build -o "$(NAME)"
	cp "$(NAME)" ~/.local/bin

clean: ## build
	rm -rf debug
	rm -rf a.out
