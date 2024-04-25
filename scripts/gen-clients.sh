#!/bin/bash

set -auexo pipefail

# This script requres a checkout of https://github.com/kubernetes/code-generator release-1.21 in ../
# For more information read https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/
# To generate CRDs, run this in the base directory of AGIC repo. Generated files will be in ~/go/src/ dir. Copy them over to ./pkg folder.

# Commands to copy:
# cp ~/go/src/github.com/Azure/application-gateway-kubernetes-ingress/pkg/apis           ./pkg -r
# cp ~/go/src/github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client     ./pkg -r

# echo -e "Cleanup previously generated code..."
# rm -rf pkg/client $(find ./pkg -name 'zz_*.go')
source ../code-generator/kube_codegen.sh
echo -e "Generate Application Gateway CRDs..."
kube::codegen::gen_helpers \
    --boilerplate ../code-generator/examples/hack/boilerplate.go.txt \
    --extra-peer-dir "../pkg/apis/agic" \
    --extra-peer-dir "../pkg/apis/multicluster" \
    ../pkg

kube::codegen::gen_client \
    --with-watch \
    --output-dir ../pkg/crd_client/agic_crd_client \
    --output-pkg github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/agic_crd_client \
    --boilerplate ../code-generator/examples/hack/boilerplate.go.txt \
    ../pkg/apis/agic

echo -e "Generate Azure Multi-Cluster CRDs..."
kube::codegen::gen_client \
    --with-watch \
    --output-dir ../pkg/crd_client/azure_multicluster_crd_client \
    --output-pkg github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/azure_multicluster_crd_client \
    --boilerplate ../code-generator/examples/hack/boilerplate.go.txt \
    ../pkg/apis/multicluster

istio_version=$(grep 'istio.io/client-go' ../go.mod | cut -d' ' -f2)
go get istio.io/client-go@$istio_version

echo -e "Generate Istio CRDs..."
kube::codegen::gen_client \
    --with-watch \
    --one-input-api networking \
    --output-dir ../pkg/crd_client/istio_crd_client \
    --output-pkg github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/istio_crd_client \
    --boilerplate ../code-generator/examples/hack/boilerplate.go.txt \
    /home/wsl/go/pkg/mod/istio.io/client-go@$istio_version/pkg/apis
