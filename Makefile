# build
build: ## Build the app
	go build -o bin/main

# run
run-register: ## Run the register app
	./bin/main -mode reg

run-publish: ## Run the publish app
	./bin/main -mode pub