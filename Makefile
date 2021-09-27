# ============================================================================ #
# HELPERS
# ============================================================================ #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ============================================================================ #
# DEVELOPMENT
# ============================================================================ #

## run/pmtools: run the core pmtools application web server
.PHONY: run
run:
	@go run ./cmd/pmtools/

## clean: remove all mock data files
.PHONY: clean
clean:
	@rm -f credits.csv ssrs.csv gar.csv roles.csv comments.csv

## run/mock: generate mock data files
.PHONY: run/mock
run/mock: clean/mock
	@python3 ./tools/mock/main.py 200

# ============================================================================ #
# QUALITY CONTROL
# ============================================================================ #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo "Running tests..."
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ============================================================================ #
# BUILD
# ============================================================================ #

current_time = $(shell date --iso-8601=seconds)
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build/pmtools: build the cmd/pmtools application
.PHONY: build/pmtools
build:
	@echo 'Cleaning previous build...'
	@rm -rf ./bin/
	@echo 'Building cmd/pmtools...'
	go build -ldflags=${linker_flags} -o=./bin/pmtools ./cmd/pmtools
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/pmtools ./cmd/pmtools