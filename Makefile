.PHONY: tidy
build:
	@bash build.sh

.PHONY: build
run:
	@_output/job

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: clean
clean:
	@rm -rf _output