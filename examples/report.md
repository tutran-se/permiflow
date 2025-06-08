# Permiflow RBAC Audit Report

## Subject: horizontal-pod-autoscaler

- Namespace: `kube-system`
- Role: `system:controller:horizontal-pod-autoscaler`
- Verbs: `[get list]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**

## Subject: local-path-provisioner-service-account

- Namespace: `local-path-storage`
- Role: `local-path-provisioner-role`
- Verbs: `[*]`
- Resources: `[endpoints persistentvolumes pods]`
- Scope: `Cluster`
- Risk Level: **HIGH**

## Subject: generic-garbage-collector

- Namespace: `kube-system`
- Role: `system:controller:generic-garbage-collector`
- Verbs: `[delete get list patch update watch]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**

## Subject: horizontal-pod-autoscaler

- Namespace: `kube-system`
- Role: `system:controller:horizontal-pod-autoscaler`
- Verbs: `[get list]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**

## Subject: resourcequota-controller

- Namespace: `kube-system`
- Role: `system:controller:resourcequota-controller`
- Verbs: `[list watch]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**

## Subject: system:masters

- Namespace: ``
- Role: `cluster-admin`
- Verbs: `[*]`
- Resources: `[]`
- Scope: `Cluster`
- Risk Level: **HIGH**

## Subject: system:masters

- Namespace: ``
- Role: `cluster-admin`
- Verbs: `[*]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**

## Subject: namespace-controller

- Namespace: `kube-system`
- Role: `system:controller:namespace-controller`
- Verbs: `[delete deletecollection get list]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**

## Subject: system:kube-controller-manager

- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[list watch]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**

## Subject: token-cleaner

- Namespace: `kube-system`
- Role: `system:controller:token-cleaner`
- Verbs: `[delete get list watch]`
- Resources: `[secrets]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**

## Subject: system:kube-controller-manager

- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[delete]`
- Resources: `[secrets]`
- Scope: `Cluster`
- Risk Level: **MEDIUM**

## Subject: persistent-volume-binder

- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[get]`
- Resources: `[secrets]`
- Scope: `Cluster`
- Risk Level: **MEDIUM**

## Subject: system:kube-controller-manager

- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[get]`
- Resources: `[configmaps namespaces secrets serviceaccounts]`
- Scope: `Cluster`
- Risk Level: **MEDIUM**

## Subject: system:kube-controller-manager

- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[update]`
- Resources: `[secrets serviceaccounts]`
- Scope: `Cluster`
- Risk Level: **MEDIUM**

## Subject: bootstrap-signer

- Namespace: `kube-system`
- Role: `system:controller:bootstrap-signer`
- Verbs: `[get list watch]`
- Resources: `[secrets]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**

## Subject: system:kube-controller-manager

- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[create]`
- Resources: `[secrets serviceaccounts]`
- Scope: `Cluster`
- Risk Level: **MEDIUM**

## Subject: expand-controller

- Namespace: `kube-system`
- Role: `system:controller:expand-controller`
- Verbs: `[get]`
- Resources: `[secrets]`
- Scope: `Cluster`
- Risk Level: **MEDIUM**

## Subject: disruption-controller

- Namespace: `kube-system`
- Role: `system:controller:disruption-controller`
- Verbs: `[get]`
- Resources: `[*/scale]`
- Scope: `Cluster`
- Risk Level: **LOW**

## Subject: statefulset-controller

- Namespace: `kube-system`
- Role: `system:controller:statefulset-controller`
- Verbs: `[create delete get patch update]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**

## Subject: attachdetach-controller

- Namespace: `kube-system`
- Role: `system:controller:attachdetach-controller`
- Verbs: `[get list watch]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **LOW**

## Subject: attachdetach-controller

- Namespace: `kube-system`
- Role: `system:controller:attachdetach-controller`
- Verbs: `[patch update]`
- Resources: `[nodes/status]`
- Scope: `Cluster`
- Risk Level: **LOW**

## Subject: attachdetach-controller

- Namespace: `kube-system`
- Role: `system:controller:attachdetach-controller`
- Verbs: `[list watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**

## Subject: attachdetach-controller

- Namespace: `kube-system`
- Role: `system:controller:attachdetach-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
