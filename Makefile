TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=github.com
NAMESPACE=Skeen
NAME=os2mo
VERSION=0.3.1
OS_ARCH=linux_amd64

default: install

terraform-provider-os2mo: *.go */*.go go.mod test
	go build .

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: terraform-provider-os2mo
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv $+ ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	-rm examples/.terraform.lock.hcl
	cd examples && terraform init

example: install
	cd examples && terraform apply --auto-approve

test:
	terraform fmt -recursive
	go fmt ./...
	go vet .
	go test ./...

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
