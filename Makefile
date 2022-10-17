BINARY_NAME=terraform-provider-lxd
TF_TEST_DIR="tests/lxd_provider_tf"

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm -f ${BINARY_NAME}

test:
	go test -v ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

fmt:
	go fmt

lint:
	golangci-lint run --enable-all

install-provider: build
	mkdir -p ~/.terraform.d/plugins/lxdprovider.com/lxdprovider/lxd/1.0.0/linux_amd64
	cp terraform-provider-lxd ~/.terraform.d/plugins/lxdprovider.com/lxdprovider/lxd/1.0.0/linux_amd64

test-provider-apply:
	cd $(TF_TEST_DIR) && ls -A1 | grep terraform | xargs rm -rf
	cd $(TF_TEST_DIR) && terraform init && terraform apply --auto-approve

test-provider-destroy:
	cd $(TF_TEST_DIR) && ls -A1 | grep terraform | xargs rm -rf
	cd $(TF_TEST_DIR) && terraform init && terraform destroy --auto-approve

doc:
	echo "Lxd doco can be accessed here: http://localhost:6060/pkg/github.com/lrwx00t/lxd_provider/lxd/"
	godoc -http=:6060