.PHONY: tidy
build:
	@bash build.sh

.PHONY: build
run:
	@_output/tasks

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: clean
clean:
	@rm -rf _output

.PHONY: gen-model
gen-model:
	@gentool -db mysql -dsn 'root:root@tcp(127.0.0.1:3306)/gotasks' -onlyModel -modelPkgName internal/tasks/model