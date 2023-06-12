#!/bin/bash -ex

export licensing_namespace=ibm-common-services
kubectl create namespace ${licensing_namespace} || echo "namespace probably already exists"

# To add CRD:
kubectl apply -f config/crd/bases/operator.ibm.com_ibmlicensings.yaml
kubectl apply -f config/crd/bases/operator.ibm.com_ibmlicenseservicereporters.yaml
kubectl apply -f config/crd/bases/operator.ibm.com_ibmlicensingmetadatas.yaml
kubectl apply -f config/crd/bases/operator.ibm.com_ibmlicensingdefinitions.yaml
kubectl apply -f config/crd/bases/operator.ibm.com_ibmlicensingquerysources.yaml

# To add RBAC:
kubectl apply -f config/rbac/role.yaml
kubectl apply -f config/rbac/role_operands.yaml
kubectl apply -f config/rbac/service_account.yaml
kubectl apply -f config/rbac/role_binding.yaml

# To configure operator
kubectl apply -f config/manager/manager.yaml

# to deploy licensing service itself
kubectl apply -n ibm-common-services  -f licensing-service.yml
