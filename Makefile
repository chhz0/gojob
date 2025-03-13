.PHONY: build
build:
	@bash build.sh

.PHONY: clean
clean:
	@rm -rf _output

.PHONY: run
run:
	@_output/job