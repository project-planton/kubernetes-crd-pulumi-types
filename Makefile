go_build_dir=pkg
#https://github.com/istio/istio/tree/master/manifests/charts/base/crds
#https://github.com/istio/istio/tree/release-1.22
istio_release=release-1.22
#https://github.com/cert-manager/cert-manager/tree/master/deploy/crds
#https://github.com/cert-manager/cert-manager/tree/release-1.15
cert_manager_release=release-1.15

#https://github.com/apache/solr-operator/tree/v0.8.1/helm/solr-operator/crds
#https://github.com/apache/solr-operator/releases/tag/v0.8.1
solr_operator_release=v0.8.1

.PHONY: clean
clean:
	rm -rf ${go_build_dir}

.PHONY: gen-istio
gen-istio:
	crd2pulumi --force --goPath=pkg/istio https://raw.githubusercontent.com/istio/istio/${istio_release}/manifests/charts/base/crds/crd-all.gen.yaml
	#https://github.com/pulumi/crd2pulumi/issues/89
	mv pkg/kubernetes pkg/istio

.PHONY: gen-cert-manager
gen-cert-manager:
	crd2pulumi --force --goPath=pkg/certmanager https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-certificates.yaml \
		https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-certificaterequests.yaml \
		https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-challenges.yaml \
		https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-clusterissuers.yaml \
		https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-issuers.yaml \
		https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-orders.yaml
	#https://github.com/pulumi/crd2pulumi/issues/89
	mv pkg/kubernetes pkg/certmanager

.PHONY: gen-solr-operator
gen-solr-operator:
	crd2pulumi --force --goPath=pkg/istio https://raw.githubusercontent.com/apache/solr-operator/${solr_operator_release}/helm/solr-operator/crds/crds.yaml
	#https://github.com/pulumi/crd2pulumi/issues/89
	mv pkg/kubernetes pkg/solroperator

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

.PHONY: gen
gen: clean gen-istio gen-cert-manager gen-solr-operator
