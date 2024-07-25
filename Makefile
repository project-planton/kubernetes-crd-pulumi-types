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

#https://github.com/strimzi/strimzi-kafka-operator/tree/main/install/cluster-operator
#https://github.com/strimzi/strimzi-kafka-operator/tree/release-0.42.x
strimzi_operator_release=release-0.42.x

#https://github.com/zalando/postgres-operator/tree/v1.12.2/manifests
#https://github.com/zalando/postgres-operator/tree/v1.12.2
zalando_operator_release=v1.12.2

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
	crd2pulumi --force --goPath=pkg/solroperator https://raw.githubusercontent.com/apache/solr-operator/${solr_operator_release}/helm/solr-operator/crds/crds.yaml
	#https://github.com/pulumi/crd2pulumi/issues/89
	mv pkg/kubernetes pkg/solroperator

.PHONY: gen-strimzi-operator
gen-strimzi-operator:
	crd2pulumi --force --goPath=pkg/strimzioperator https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/main/install/cluster-operator/040-Crd-kafka.yaml \
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/041-Crd-kafkaconnect.yaml \
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/042-Crd-strimzipodset.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/043-Crd-kafkatopic.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/044-Crd-kafkauser.yaml \
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/045-Crd-kafkamirrormaker.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/046-Crd-kafkabridge.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/047-Crd-kafkaconnector.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/048-Crd-kafkamirrormaker2.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/049-Crd-kafkarebalance.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/04A-Crd-kafkanodepool.yaml
	#https://github.com/pulumi/crd2pulumi/issues/89
	mv pkg/kubernetes pkg/strimzioperator


.PHONY: gen-zalando-operator
gen-zalando-operator:
	crd2pulumi --force --goPath=pkg/zalandooperator https://raw.githubusercontent.com/zalando/postgres-operator/v1.12.2/manifests/operatorconfiguration.crd.yaml \
		https://raw.githubusercontent.com/zalando/postgres-operator/v1.12.2/manifests/postgresql.crd.yaml \
		https://raw.githubusercontent.com/zalando/postgres-operator/v1.12.2/manifests/postgresteam.crd.yaml
	#https://github.com/pulumi/crd2pulumi/issues/89
	mv pkg/kubernetes pkg/zalandooperator

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
gen: clean gen-istio gen-cert-manager gen-solr-operator gen-strimzi-operator gen-zalando-operator
