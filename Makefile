TEST?=$$(go list ./... |grep -v 'vendor')
TARGETS=darwin linux windows
SENSU_VERSION=5.3.0

docker:
	rm -rf /var/lib/sensu
	for i in $$(docker ps -q); do docker rm -f $$i; done
	docker pull sensu/sensu:$(SENSU_VERSION)
	docker run -v /var/lib/sensu:/var/lib/sensu -d --name sensu-backend -p 2380:2380 -p 3000:3000 -p 8080:8080 -p 8081:8081 sensu/sensu:$(SENSU_VERSION) sensu-backend start
	docker run -v /var/lib/sensu:/var/lib/sensu -d --name sensu-agent sensu/sensu:$(SENSU_VERSION) sensu-agent start --backend-url ws://localhost:8081 --subscriptions webserver,system --cache-dir /var/lib/sensu

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

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
	GOOS=$@ GOARCH=amd64 go build -o "dist/$@/terraform-provider-sensu_${TRAVIS_TAG}_x4"
	zip -j dist/terraform-provider-sensu_${TRAVIS_TAG}_$@_amd64.zip dist/$@/terraform-provider-sensu_${TRAVIS_TAG}_x4
