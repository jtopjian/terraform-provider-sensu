TEST?=$$(go list ./... |grep -v 'vendor')
TARGETS=darwin linux windows
SENSU_VERSION=6.11.0

docker:
	docker-compose down || true
	docker-compose up -d

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m -count 1

build:
	go install

fmtcheck:
	echo "==> Checking that code complies with gofmt requirements..."
	files=$$(find . -name '*.go' | grep -v 'vendor' ) ; \
	gofmt_files=`gofmt -l $$files`; \
	if [ -n "$$gofmt_files" ]; then \
		echo 'gofmt needs running on the following files:'; \
		echo "$$gofmt_files"; \
		echo "You can use the command: \`make fmt\` to reformat code."; \
		exit 1; \
	fi

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

targets: $(TARGETS)

$(TARGETS):
	CGO_ENABLED=0 GOOS=$@ GOARCH=amd64 go build -o "dist/$@/terraform-provider-sensu_${TRAVIS_TAG}_x4"
	zip -j dist/terraform-provider-sensu_${TRAVIS_TAG}_$@_amd64.zip dist/$@/terraform-provider-sensu_${TRAVIS_TAG}_x4
