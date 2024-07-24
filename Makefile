go_build_dir=pkg
#https://github.com/istio/istio/tree/master/manifests/charts/base/crds
#https://github.com/istio/istio/tree/release-1.22
istio_release=release-1.22
#https://github.com/cert-manager/cert-manager/tree/master/deploy/crds
#https://github.com/cert-manager/cert-manager/tree/release-1.15
cert_manager_release=release-1.15

.PHONY: clean
clean:
	rm -rf ${go_build_dir}

.PHONY: gen-istio
gen-istio:
	crd2pulumi --goPath=pkg/istio https://raw.githubusercontent.com/istio/istio/${istio_release}/manifests/charts/base/crds/crd-all.gen.yaml

.PHONY: gen-cert-manager
gen-cert-manager:
	crd2pulumi --goPath=pkg/certmanager https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-certificates.yaml
	crd2pulumi --goPath=pkg/certmanager https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-certificaterequests.yaml
	crd2pulumi --goPath=pkg/certmanager https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-challenges.yaml
	crd2pulumi --goPath=pkg/certmanager https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-clusterissuers.yaml
	crd2pulumi --goPath=pkg/certmanager https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-issuers.yaml
	crd2pulumi --goPath=pkg/certmanager https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-orders.yaml

.PHONY: gen
gen: clean gen-istio gen-cert-manager

.PHONY: go-deps
go-deps:
	go mod download
	go mod tidy

.PHONY: go-vet
go-vet:
	go vet ./...

.PHONY: go-fmt
go-fmt:
	go fmt ./...

.PHONY: build-go
build-go: go-deps go-vet go-fmt

.PHONY: build
build: clean gen build-go
