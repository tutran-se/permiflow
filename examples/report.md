# Permiflow RBAC Audit Report
## 📊 Summary
- Total bindings scanned: **240**
- Found 2 cluster-admin binding(s)
- Found 3 wildcard verb usage(s)
- Found 8 subject(s) with secrets access
- Found 0 privilege escalation(s)
- Found 16 exec access(es)
- Found 16 config read secrets access(es)

---
## 🚦 Risk Levels
- **HIGH**: Wildcard verbs or resources, privilege escalation risks
- **MEDIUM**: Sensitive resources with non-wildcard verbs
- **LOW**: Non-sensitive resources with non-wildcard verbs

---
## 📘 Table of Contents
- [clusterrole-aggregation-controller (ServiceAccount)](#clusterrole-aggregation-controller-serviceaccount)
- [system:masters (Group)](#systemmasters-group)
- [system:kube-proxy (User)](#systemkube-proxy-user)
- [kindnet (ServiceAccount)](#kindnet-serviceaccount)
- [system:bootstrappers:kubeadm:default-node-token (Group)](#systembootstrapperskubeadmdefault-node-token-group)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [system:masters (Group)](#systemmasters-group)
- [coredns (ServiceAccount)](#coredns-serviceaccount)
- [kube-proxy (ServiceAccount)](#kube-proxy-serviceaccount)
- [ttl-controller (ServiceAccount)](#ttl-controller-serviceaccount)
- [service-controller (ServiceAccount)](#service-controller-serviceaccount)
- [local-path-provisioner-service-account (ServiceAccount)](#local-path-provisioner-service-account-serviceaccount)
- [local-path-provisioner-service-account (ServiceAccount)](#local-path-provisioner-service-account-serviceaccount)
- [route-controller (ServiceAccount)](#route-controller-serviceaccount)
- [resourcequota-controller (ServiceAccount)](#resourcequota-controller-serviceaccount)
- [pod-garbage-collector (ServiceAccount)](#pod-garbage-collector-serviceaccount)
- [persistent-volume-binder (ServiceAccount)](#persistent-volume-binder-serviceaccount)
- [node-controller (ServiceAccount)](#node-controller-serviceaccount)
- [attachdetach-controller (ServiceAccount)](#attachdetach-controller-serviceaccount)
- [namespace-controller (ServiceAccount)](#namespace-controller-serviceaccount)
- [horizontal-pod-autoscaler (ServiceAccount)](#horizontal-pod-autoscaler-serviceaccount)
- [horizontal-pod-autoscaler (ServiceAccount)](#horizontal-pod-autoscaler-serviceaccount)
- [generic-garbage-collector (ServiceAccount)](#generic-garbage-collector-serviceaccount)
- [endpointslice-controller (ServiceAccount)](#endpointslice-controller-serviceaccount)
- [daemon-set-controller (ServiceAccount)](#daemon-set-controller-serviceaccount)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [expand-controller (ServiceAccount)](#expand-controller-serviceaccount)
- [token-cleaner (ServiceAccount)](#token-cleaner-serviceaccount)
- [cloud-provider (ServiceAccount)](#cloud-provider-serviceaccount)
- [bootstrap-signer (ServiceAccount)](#bootstrap-signer-serviceaccount)
- [kube-scheduler (ServiceAccount)](#kube-scheduler-serviceaccount)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [kube-controller-manager (ServiceAccount)](#kube-controller-manager-serviceaccount)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [system:nodes (Group)](#systemnodes-group)
- [system:bootstrappers:kubeadm:default-node-token (Group)](#systembootstrapperskubeadmdefault-node-token-group)
- [system:bootstrappers:kubeadm:default-node-token (Group)](#systembootstrapperskubeadmdefault-node-token-group)
- [system:nodes (Group)](#systemnodes-group)
- [system:bootstrappers:kubeadm:default-node-token (Group)](#systembootstrapperskubeadmdefault-node-token-group)
- [bootstrap-signer (ServiceAccount)](#bootstrap-signer-serviceaccount)
- [system:anonymous (User)](#systemanonymous-user)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [persistent-volume-binder (ServiceAccount)](#persistent-volume-binder-serviceaccount)
- [statefulset-controller (ServiceAccount)](#statefulset-controller-serviceaccount)
- [route-controller (ServiceAccount)](#route-controller-serviceaccount)
- [deployment-controller (ServiceAccount)](#deployment-controller-serviceaccount)
- [deployment-controller (ServiceAccount)](#deployment-controller-serviceaccount)
- [deployment-controller (ServiceAccount)](#deployment-controller-serviceaccount)
- [replicaset-controller (ServiceAccount)](#replicaset-controller-serviceaccount)
- [disruption-controller (ServiceAccount)](#disruption-controller-serviceaccount)
- [disruption-controller (ServiceAccount)](#disruption-controller-serviceaccount)
- [disruption-controller (ServiceAccount)](#disruption-controller-serviceaccount)
- [disruption-controller (ServiceAccount)](#disruption-controller-serviceaccount)
- [disruption-controller (ServiceAccount)](#disruption-controller-serviceaccount)
- [disruption-controller (ServiceAccount)](#disruption-controller-serviceaccount)
- [disruption-controller (ServiceAccount)](#disruption-controller-serviceaccount)
- [disruption-controller (ServiceAccount)](#disruption-controller-serviceaccount)
- [endpoint-controller (ServiceAccount)](#endpoint-controller-serviceaccount)
- [endpoint-controller (ServiceAccount)](#endpoint-controller-serviceaccount)
- [endpoint-controller (ServiceAccount)](#endpoint-controller-serviceaccount)
- [endpoint-controller (ServiceAccount)](#endpoint-controller-serviceaccount)
- [attachdetach-controller (ServiceAccount)](#attachdetach-controller-serviceaccount)
- [endpointslice-controller (ServiceAccount)](#endpointslice-controller-serviceaccount)
- [endpointslice-controller (ServiceAccount)](#endpointslice-controller-serviceaccount)
- [endpointslice-controller (ServiceAccount)](#endpointslice-controller-serviceaccount)
- [endpointslicemirroring-controller (ServiceAccount)](#endpointslicemirroring-controller-serviceaccount)
- [endpointslicemirroring-controller (ServiceAccount)](#endpointslicemirroring-controller-serviceaccount)
- [endpointslicemirroring-controller (ServiceAccount)](#endpointslicemirroring-controller-serviceaccount)
- [endpointslicemirroring-controller (ServiceAccount)](#endpointslicemirroring-controller-serviceaccount)
- [endpointslicemirroring-controller (ServiceAccount)](#endpointslicemirroring-controller-serviceaccount)
- [ephemeral-volume-controller (ServiceAccount)](#ephemeral-volume-controller-serviceaccount)
- [ephemeral-volume-controller (ServiceAccount)](#ephemeral-volume-controller-serviceaccount)
- [ephemeral-volume-controller (ServiceAccount)](#ephemeral-volume-controller-serviceaccount)
- [ephemeral-volume-controller (ServiceAccount)](#ephemeral-volume-controller-serviceaccount)
- [expand-controller (ServiceAccount)](#expand-controller-serviceaccount)
- [expand-controller (ServiceAccount)](#expand-controller-serviceaccount)
- [expand-controller (ServiceAccount)](#expand-controller-serviceaccount)
- [expand-controller (ServiceAccount)](#expand-controller-serviceaccount)
- [expand-controller (ServiceAccount)](#expand-controller-serviceaccount)
- [deployment-controller (ServiceAccount)](#deployment-controller-serviceaccount)
- [expand-controller (ServiceAccount)](#expand-controller-serviceaccount)
- [attachdetach-controller (ServiceAccount)](#attachdetach-controller-serviceaccount)
- [generic-garbage-collector (ServiceAccount)](#generic-garbage-collector-serviceaccount)
- [horizontal-pod-autoscaler (ServiceAccount)](#horizontal-pod-autoscaler-serviceaccount)
- [horizontal-pod-autoscaler (ServiceAccount)](#horizontal-pod-autoscaler-serviceaccount)
- [horizontal-pod-autoscaler (ServiceAccount)](#horizontal-pod-autoscaler-serviceaccount)
- [horizontal-pod-autoscaler (ServiceAccount)](#horizontal-pod-autoscaler-serviceaccount)
- [horizontal-pod-autoscaler (ServiceAccount)](#horizontal-pod-autoscaler-serviceaccount)
- [attachdetach-controller (ServiceAccount)](#attachdetach-controller-serviceaccount)
- [attachdetach-controller (ServiceAccount)](#attachdetach-controller-serviceaccount)
- [horizontal-pod-autoscaler (ServiceAccount)](#horizontal-pod-autoscaler-serviceaccount)
- [job-controller (ServiceAccount)](#job-controller-serviceaccount)
- [job-controller (ServiceAccount)](#job-controller-serviceaccount)
- [job-controller (ServiceAccount)](#job-controller-serviceaccount)
- [job-controller (ServiceAccount)](#job-controller-serviceaccount)
- [job-controller (ServiceAccount)](#job-controller-serviceaccount)
- [resourcequota-controller (ServiceAccount)](#resourcequota-controller-serviceaccount)
- [namespace-controller (ServiceAccount)](#namespace-controller-serviceaccount)
- [attachdetach-controller (ServiceAccount)](#attachdetach-controller-serviceaccount)
- [attachdetach-controller (ServiceAccount)](#attachdetach-controller-serviceaccount)
- [node-controller (ServiceAccount)](#node-controller-serviceaccount)
- [node-controller (ServiceAccount)](#node-controller-serviceaccount)
- [node-controller (ServiceAccount)](#node-controller-serviceaccount)
- [node-controller (ServiceAccount)](#node-controller-serviceaccount)
- [node-controller (ServiceAccount)](#node-controller-serviceaccount)
- [node-controller (ServiceAccount)](#node-controller-serviceaccount)
- [persistent-volume-binder (ServiceAccount)](#persistent-volume-binder-serviceaccount)
- [persistent-volume-binder (ServiceAccount)](#persistent-volume-binder-serviceaccount)
- [persistent-volume-binder (ServiceAccount)](#persistent-volume-binder-serviceaccount)
- [persistent-volume-binder (ServiceAccount)](#persistent-volume-binder-serviceaccount)
- [persistent-volume-binder (ServiceAccount)](#persistent-volume-binder-serviceaccount)
- [persistent-volume-binder (ServiceAccount)](#persistent-volume-binder-serviceaccount)
- [persistent-volume-binder (ServiceAccount)](#persistent-volume-binder-serviceaccount)
- [persistent-volume-binder (ServiceAccount)](#persistent-volume-binder-serviceaccount)
- [deployment-controller (ServiceAccount)](#deployment-controller-serviceaccount)
- [system:authenticated (Group)](#systemauthenticated-group)
- [persistent-volume-binder (ServiceAccount)](#persistent-volume-binder-serviceaccount)
- [persistent-volume-binder (ServiceAccount)](#persistent-volume-binder-serviceaccount)
- [pod-garbage-collector (ServiceAccount)](#pod-garbage-collector-serviceaccount)
- [system:authenticated (Group)](#systemauthenticated-group)
- [pod-garbage-collector (ServiceAccount)](#pod-garbage-collector-serviceaccount)
- [pv-protection-controller (ServiceAccount)](#pv-protection-controller-serviceaccount)
- [pv-protection-controller (ServiceAccount)](#pv-protection-controller-serviceaccount)
- [pvc-protection-controller (ServiceAccount)](#pvc-protection-controller-serviceaccount)
- [pvc-protection-controller (ServiceAccount)](#pvc-protection-controller-serviceaccount)
- [pvc-protection-controller (ServiceAccount)](#pvc-protection-controller-serviceaccount)
- [replicaset-controller (ServiceAccount)](#replicaset-controller-serviceaccount)
- [certificate-controller (ServiceAccount)](#certificate-controller-serviceaccount)
- [certificate-controller (ServiceAccount)](#certificate-controller-serviceaccount)
- [disruption-controller (ServiceAccount)](#disruption-controller-serviceaccount)
- [replicaset-controller (ServiceAccount)](#replicaset-controller-serviceaccount)
- [replication-controller (ServiceAccount)](#replication-controller-serviceaccount)
- [replication-controller (ServiceAccount)](#replication-controller-serviceaccount)
- [replication-controller (ServiceAccount)](#replication-controller-serviceaccount)
- [replication-controller (ServiceAccount)](#replication-controller-serviceaccount)
- [replication-controller (ServiceAccount)](#replication-controller-serviceaccount)
- [local-path-provisioner-service-account (ServiceAccount)](#local-path-provisioner-service-account-serviceaccount)
- [namespace-controller (ServiceAccount)](#namespace-controller-serviceaccount)
- [deployment-controller (ServiceAccount)](#deployment-controller-serviceaccount)
- [kube-scheduler (ServiceAccount)](#kube-scheduler-serviceaccount)
- [root-ca-cert-publisher (ServiceAccount)](#root-ca-cert-publisher-serviceaccount)
- [local-path-provisioner-service-account (ServiceAccount)](#local-path-provisioner-service-account-serviceaccount)
- [resourcequota-controller (ServiceAccount)](#resourcequota-controller-serviceaccount)
- [route-controller (ServiceAccount)](#route-controller-serviceaccount)
- [service-account-controller (ServiceAccount)](#service-account-controller-serviceaccount)
- [service-account-controller (ServiceAccount)](#service-account-controller-serviceaccount)
- [service-controller (ServiceAccount)](#service-controller-serviceaccount)
- [service-controller (ServiceAccount)](#service-controller-serviceaccount)
- [kube-proxy (ServiceAccount)](#kube-proxy-serviceaccount)
- [service-controller (ServiceAccount)](#service-controller-serviceaccount)
- [statefulset-controller (ServiceAccount)](#statefulset-controller-serviceaccount)
- [statefulset-controller (ServiceAccount)](#statefulset-controller-serviceaccount)
- [replicaset-controller (ServiceAccount)](#replicaset-controller-serviceaccount)
- [statefulset-controller (ServiceAccount)](#statefulset-controller-serviceaccount)
- [statefulset-controller (ServiceAccount)](#statefulset-controller-serviceaccount)
- [statefulset-controller (ServiceAccount)](#statefulset-controller-serviceaccount)
- [statefulset-controller (ServiceAccount)](#statefulset-controller-serviceaccount)
- [statefulset-controller (ServiceAccount)](#statefulset-controller-serviceaccount)
- [statefulset-controller (ServiceAccount)](#statefulset-controller-serviceaccount)
- [ttl-after-finished-controller (ServiceAccount)](#ttl-after-finished-controller-serviceaccount)
- [ttl-after-finished-controller (ServiceAccount)](#ttl-after-finished-controller-serviceaccount)
- [kube-proxy (ServiceAccount)](#kube-proxy-serviceaccount)
- [ttl-controller (ServiceAccount)](#ttl-controller-serviceaccount)
- [coredns (ServiceAccount)](#coredns-serviceaccount)
- [kube-proxy (ServiceAccount)](#kube-proxy-serviceaccount)
- [coredns (ServiceAccount)](#coredns-serviceaccount)
- [system:authenticated (Group)](#systemauthenticated-group)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [daemon-set-controller (ServiceAccount)](#daemon-set-controller-serviceaccount)
- [daemon-set-controller (ServiceAccount)](#daemon-set-controller-serviceaccount)
- [daemon-set-controller (ServiceAccount)](#daemon-set-controller-serviceaccount)
- [system:nodes (Group)](#systemnodes-group)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [system:bootstrappers:kubeadm:default-node-token (Group)](#systembootstrapperskubeadmdefault-node-token-group)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [kube-dns (ServiceAccount)](#kube-dns-serviceaccount)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:bootstrappers:kubeadm:default-node-token (Group)](#systembootstrapperskubeadmdefault-node-token-group)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:monitoring (Group)](#systemmonitoring-group)
- [system:kube-proxy (User)](#systemkube-proxy-user)
- [kindnet (ServiceAccount)](#kindnet-serviceaccount)
- [system:kube-proxy (User)](#systemkube-proxy-user)
- [system:kube-proxy (User)](#systemkube-proxy-user)
- [system:authenticated (Group)](#systemauthenticated-group)
- [system:unauthenticated (Group)](#systemunauthenticated-group)
- [system:serviceaccounts (Group)](#systemserviceaccounts-group)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [daemon-set-controller (ServiceAccount)](#daemon-set-controller-serviceaccount)
- [attachdetach-controller (ServiceAccount)](#attachdetach-controller-serviceaccount)
- [bootstrap-signer (ServiceAccount)](#bootstrap-signer-serviceaccount)
- [bootstrap-signer (ServiceAccount)](#bootstrap-signer-serviceaccount)
- [daemon-set-controller (ServiceAccount)](#daemon-set-controller-serviceaccount)
- [daemon-set-controller (ServiceAccount)](#daemon-set-controller-serviceaccount)
- [daemon-set-controller (ServiceAccount)](#daemon-set-controller-serviceaccount)
- [cronjob-controller (ServiceAccount)](#cronjob-controller-serviceaccount)
- [cronjob-controller (ServiceAccount)](#cronjob-controller-serviceaccount)
- [cronjob-controller (ServiceAccount)](#cronjob-controller-serviceaccount)
- [cronjob-controller (ServiceAccount)](#cronjob-controller-serviceaccount)
- [system:kube-controller-manager (User)](#systemkube-controller-manager-user)
- [cronjob-controller (ServiceAccount)](#cronjob-controller-serviceaccount)
- [kube-controller-manager (ServiceAccount)](#kube-controller-manager-serviceaccount)
- [cronjob-controller (ServiceAccount)](#cronjob-controller-serviceaccount)
- [system:kube-scheduler (User)](#systemkube-scheduler-user)
- [certificate-controller (ServiceAccount)](#certificate-controller-serviceaccount)
- [root-ca-cert-publisher (ServiceAccount)](#root-ca-cert-publisher-serviceaccount)
- [certificate-controller (ServiceAccount)](#certificate-controller-serviceaccount)
- [replicaset-controller (ServiceAccount)](#replicaset-controller-serviceaccount)
- [certificate-controller (ServiceAccount)](#certificate-controller-serviceaccount)
- [certificate-controller (ServiceAccount)](#certificate-controller-serviceaccount)
- [token-cleaner (ServiceAccount)](#token-cleaner-serviceaccount)

---
## <a name="clusterrole-aggregation-controller-serviceaccount"></a>Subject: clusterrole-aggregation-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:clusterrole-aggregation-controller`
- Verbs: `[escalate get list patch update watch]`
- Resources: `[clusterroles]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Can update clusterroles (potential privilege escalation)


## <a name="systemmasters-group"></a>Subject: system:masters (Group)
- Namespace: ``
- Role: `cluster-admin`
- Verbs: `[*]`
- Resources: `[]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Wildcard verb or resource detected


## <a name="systemkube-proxy-user"></a>Subject: system:kube-proxy (User)
- Namespace: ``
- Role: `system:node-proxier`
- Verbs: `[get list watch]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="kindnet-serviceaccount"></a>Subject: kindnet (ServiceAccount)
- Namespace: `kube-system`
- Role: `kindnet`
- Verbs: `[list watch]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="systembootstrapperskubeadmdefault-node-token-group"></a>Subject: system:bootstrappers:kubeadm:default-node-token (Group)
- Namespace: ``
- Role: `kubeadm:get-nodes`
- Verbs: `[get]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[get list watch]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[list watch]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Wildcard verb or resource detected


## <a name="systemmasters-group"></a>Subject: system:masters (Group)
- Namespace: ``
- Role: `cluster-admin`
- Verbs: `[*]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Wildcard verb or resource detected


## <a name="coredns-serviceaccount"></a>Subject: coredns (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:coredns`
- Verbs: `[get]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="kube-proxy-serviceaccount"></a>Subject: kube-proxy (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:node-proxier`
- Verbs: `[get list watch]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="ttl-controller-serviceaccount"></a>Subject: ttl-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:ttl-controller`
- Verbs: `[list patch update watch]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="service-controller-serviceaccount"></a>Subject: service-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:service-controller`
- Verbs: `[list watch]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="local-path-provisioner-service-account-serviceaccount"></a>Subject: local-path-provisioner-service-account (ServiceAccount)
- Namespace: `local-path-storage`
- Role: `local-path-provisioner-role`
- Verbs: `[get list watch]`
- Resources: `[nodes persistentvolumeclaims configmaps]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="local-path-provisioner-service-account-serviceaccount"></a>Subject: local-path-provisioner-service-account (ServiceAccount)
- Namespace: `local-path-storage`
- Role: `local-path-provisioner-role`
- Verbs: `[*]`
- Resources: `[endpoints persistentvolumes pods]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Wildcard verb or resource detected


## <a name="route-controller-serviceaccount"></a>Subject: route-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:route-controller`
- Verbs: `[list watch]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="resourcequota-controller-serviceaccount"></a>Subject: resourcequota-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:resourcequota-controller`
- Verbs: `[list watch]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Wildcard verb or resource detected


## <a name="pod-garbage-collector-serviceaccount"></a>Subject: pod-garbage-collector (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:pod-garbage-collector`
- Verbs: `[get list]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="persistent-volume-binder-serviceaccount"></a>Subject: persistent-volume-binder (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[get list]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="node-controller-serviceaccount"></a>Subject: node-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:node-controller`
- Verbs: `[delete get list patch update]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="attachdetach-controller-serviceaccount"></a>Subject: attachdetach-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:attachdetach-controller`
- Verbs: `[get list watch]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="namespace-controller-serviceaccount"></a>Subject: namespace-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:namespace-controller`
- Verbs: `[delete deletecollection get list]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Wildcard verb or resource detected


## <a name="horizontal-pod-autoscaler-serviceaccount"></a>Subject: horizontal-pod-autoscaler (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:horizontal-pod-autoscaler`
- Verbs: `[get list]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Wildcard verb or resource detected


## <a name="horizontal-pod-autoscaler-serviceaccount"></a>Subject: horizontal-pod-autoscaler (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:horizontal-pod-autoscaler`
- Verbs: `[get list]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Wildcard verb or resource detected


## <a name="generic-garbage-collector-serviceaccount"></a>Subject: generic-garbage-collector (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:generic-garbage-collector`
- Verbs: `[delete get list patch update watch]`
- Resources: `[*]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Wildcard verb or resource detected


## <a name="endpointslice-controller-serviceaccount"></a>Subject: endpointslice-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpointslice-controller`
- Verbs: `[get list watch]`
- Resources: `[nodes pods services]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="daemon-set-controller-serviceaccount"></a>Subject: daemon-set-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:daemon-set-controller`
- Verbs: `[list watch]`
- Resources: `[nodes]`
- Scope: `Cluster`
- Risk Level: **HIGH**
- Reason: Access to nodes (host-level access)


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[update]`
- Resources: `[secrets serviceaccounts]`
- Scope: `Cluster`
- Risk Level: **MEDIUM**
- Reason: Can access secrets (sensitive credential exposure)


## <a name="expand-controller-serviceaccount"></a>Subject: expand-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:expand-controller`
- Verbs: `[get]`
- Resources: `[secrets]`
- Scope: `Cluster`
- Risk Level: **MEDIUM**
- Reason: Can access secrets (sensitive credential exposure)


## <a name="token-cleaner-serviceaccount"></a>Subject: token-cleaner (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:token-cleaner`
- Verbs: `[delete get list watch]`
- Resources: `[secrets]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can access secrets (sensitive credential exposure)


## <a name="cloud-provider-serviceaccount"></a>Subject: cloud-provider (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:cloud-provider`
- Verbs: `[create get list watch]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="bootstrap-signer-serviceaccount"></a>Subject: bootstrap-signer (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:bootstrap-signer`
- Verbs: `[get list watch]`
- Resources: `[secrets]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can access secrets (sensitive credential exposure)


## <a name="kube-scheduler-serviceaccount"></a>Subject: kube-scheduler (ServiceAccount)
- Namespace: `kube-system`
- Role: `system::leader-locking-kube-scheduler`
- Verbs: `[get update]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system::leader-locking-kube-scheduler`
- Verbs: `[get update]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="kube-controller-manager-serviceaccount"></a>Subject: kube-controller-manager (ServiceAccount)
- Namespace: `kube-system`
- Role: `system::leader-locking-kube-controller-manager`
- Verbs: `[get update]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system::leader-locking-kube-controller-manager`
- Verbs: `[get update]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `extension-apiserver-authentication-reader`
- Verbs: `[get list watch]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `extension-apiserver-authentication-reader`
- Verbs: `[get list watch]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="systemnodes-group"></a>Subject: system:nodes (Group)
- Namespace: ``
- Role: `kubeadm:nodes-kubeadm-config`
- Verbs: `[get]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="systembootstrapperskubeadmdefault-node-token-group"></a>Subject: system:bootstrappers:kubeadm:default-node-token (Group)
- Namespace: ``
- Role: `kubeadm:nodes-kubeadm-config`
- Verbs: `[get]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="systembootstrapperskubeadmdefault-node-token-group"></a>Subject: system:bootstrappers:kubeadm:default-node-token (Group)
- Namespace: ``
- Role: `kubeadm:kubelet-config`
- Verbs: `[get]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="systemnodes-group"></a>Subject: system:nodes (Group)
- Namespace: ``
- Role: `kubeadm:kubelet-config`
- Verbs: `[get]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="systembootstrapperskubeadmdefault-node-token-group"></a>Subject: system:bootstrappers:kubeadm:default-node-token (Group)
- Namespace: ``
- Role: `kube-proxy`
- Verbs: `[get]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="bootstrap-signer-serviceaccount"></a>Subject: bootstrap-signer (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:bootstrap-signer`
- Verbs: `[get list watch]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="systemanonymous-user"></a>Subject: system:anonymous (User)
- Namespace: ``
- Role: `kubeadm:bootstrap-signer-clusterinfo`
- Verbs: `[get]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **MEDIUM**
- Reason: Can read configmaps (often contains credentials or API keys)


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[get]`
- Resources: `[configmaps namespaces secrets serviceaccounts]`
- Scope: `Cluster`
- Risk Level: **MEDIUM**
- Reason: Can access secrets (sensitive credential exposure)


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[delete]`
- Resources: `[secrets]`
- Scope: `Cluster`
- Risk Level: **MEDIUM**
- Reason: Can access secrets (sensitive credential exposure)


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[create]`
- Resources: `[secrets serviceaccounts]`
- Scope: `Cluster`
- Risk Level: **MEDIUM**
- Reason: Can access secrets (sensitive credential exposure)


## <a name="persistent-volume-binder-serviceaccount"></a>Subject: persistent-volume-binder (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[get]`
- Resources: `[secrets]`
- Scope: `Cluster`
- Risk Level: **MEDIUM**
- Reason: Can access secrets (sensitive credential exposure)


## <a name="statefulset-controller-serviceaccount"></a>Subject: statefulset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:statefulset-controller`
- Verbs: `[update]`
- Resources: `[statefulsets/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="route-controller-serviceaccount"></a>Subject: route-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:route-controller`
- Verbs: `[patch]`
- Resources: `[nodes/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="deployment-controller-serviceaccount"></a>Subject: deployment-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:deployment-controller`
- Verbs: `[create delete get list patch update watch]`
- Resources: `[replicasets]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="deployment-controller-serviceaccount"></a>Subject: deployment-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:deployment-controller`
- Verbs: `[get list update watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="deployment-controller-serviceaccount"></a>Subject: deployment-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:deployment-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="replicaset-controller-serviceaccount"></a>Subject: replicaset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:replicaset-controller`
- Verbs: `[create delete list patch watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="disruption-controller-serviceaccount"></a>Subject: disruption-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:disruption-controller`
- Verbs: `[get list watch]`
- Resources: `[replicasets]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="disruption-controller-serviceaccount"></a>Subject: disruption-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:disruption-controller`
- Verbs: `[get list watch]`
- Resources: `[replicationcontrollers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="disruption-controller-serviceaccount"></a>Subject: disruption-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:disruption-controller`
- Verbs: `[get list watch]`
- Resources: `[poddisruptionbudgets]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="disruption-controller-serviceaccount"></a>Subject: disruption-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:disruption-controller`
- Verbs: `[get list watch]`
- Resources: `[statefulsets]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="disruption-controller-serviceaccount"></a>Subject: disruption-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:disruption-controller`
- Verbs: `[update]`
- Resources: `[poddisruptionbudgets/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="disruption-controller-serviceaccount"></a>Subject: disruption-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:disruption-controller`
- Verbs: `[get]`
- Resources: `[*/scale]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="disruption-controller-serviceaccount"></a>Subject: disruption-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:disruption-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="disruption-controller-serviceaccount"></a>Subject: disruption-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:disruption-controller`
- Verbs: `[patch]`
- Resources: `[pods/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="endpoint-controller-serviceaccount"></a>Subject: endpoint-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpoint-controller`
- Verbs: `[get list watch]`
- Resources: `[pods services]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="endpoint-controller-serviceaccount"></a>Subject: endpoint-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpoint-controller`
- Verbs: `[create delete get list update]`
- Resources: `[endpoints]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="endpoint-controller-serviceaccount"></a>Subject: endpoint-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpoint-controller`
- Verbs: `[create]`
- Resources: `[endpoints/restricted]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="endpoint-controller-serviceaccount"></a>Subject: endpoint-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpoint-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="attachdetach-controller-serviceaccount"></a>Subject: attachdetach-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:attachdetach-controller`
- Verbs: `[get list watch]`
- Resources: `[csidrivers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="endpointslice-controller-serviceaccount"></a>Subject: endpointslice-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpointslice-controller`
- Verbs: `[update]`
- Resources: `[services/finalizers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="endpointslice-controller-serviceaccount"></a>Subject: endpointslice-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpointslice-controller`
- Verbs: `[create delete get list update]`
- Resources: `[endpointslices]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="endpointslice-controller-serviceaccount"></a>Subject: endpointslice-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpointslice-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="endpointslicemirroring-controller-serviceaccount"></a>Subject: endpointslicemirroring-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpointslicemirroring-controller`
- Verbs: `[get list watch]`
- Resources: `[endpoints services]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="endpointslicemirroring-controller-serviceaccount"></a>Subject: endpointslicemirroring-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpointslicemirroring-controller`
- Verbs: `[update]`
- Resources: `[services/finalizers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="endpointslicemirroring-controller-serviceaccount"></a>Subject: endpointslicemirroring-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpointslicemirroring-controller`
- Verbs: `[update]`
- Resources: `[endpoints/finalizers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="endpointslicemirroring-controller-serviceaccount"></a>Subject: endpointslicemirroring-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpointslicemirroring-controller`
- Verbs: `[create delete get list update]`
- Resources: `[endpointslices]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="endpointslicemirroring-controller-serviceaccount"></a>Subject: endpointslicemirroring-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:endpointslicemirroring-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="ephemeral-volume-controller-serviceaccount"></a>Subject: ephemeral-volume-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:ephemeral-volume-controller`
- Verbs: `[get list watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="ephemeral-volume-controller-serviceaccount"></a>Subject: ephemeral-volume-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:ephemeral-volume-controller`
- Verbs: `[update]`
- Resources: `[pods/finalizers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="ephemeral-volume-controller-serviceaccount"></a>Subject: ephemeral-volume-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:ephemeral-volume-controller`
- Verbs: `[create get list watch]`
- Resources: `[persistentvolumeclaims]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="ephemeral-volume-controller-serviceaccount"></a>Subject: ephemeral-volume-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:ephemeral-volume-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="expand-controller-serviceaccount"></a>Subject: expand-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:expand-controller`
- Verbs: `[get list patch update watch]`
- Resources: `[persistentvolumes]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="expand-controller-serviceaccount"></a>Subject: expand-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:expand-controller`
- Verbs: `[patch update]`
- Resources: `[persistentvolumeclaims/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="expand-controller-serviceaccount"></a>Subject: expand-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:expand-controller`
- Verbs: `[get list watch]`
- Resources: `[persistentvolumeclaims]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="expand-controller-serviceaccount"></a>Subject: expand-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:expand-controller`
- Verbs: `[get list watch]`
- Resources: `[storageclasses]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="expand-controller-serviceaccount"></a>Subject: expand-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:expand-controller`
- Verbs: `[get]`
- Resources: `[endpoints services]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="deployment-controller-serviceaccount"></a>Subject: deployment-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:deployment-controller`
- Verbs: `[update]`
- Resources: `[deployments/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="expand-controller-serviceaccount"></a>Subject: expand-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:expand-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="attachdetach-controller-serviceaccount"></a>Subject: attachdetach-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:attachdetach-controller`
- Verbs: `[create delete get list watch]`
- Resources: `[volumeattachments]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="generic-garbage-collector-serviceaccount"></a>Subject: generic-garbage-collector (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:generic-garbage-collector`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="horizontal-pod-autoscaler-serviceaccount"></a>Subject: horizontal-pod-autoscaler (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:horizontal-pod-autoscaler`
- Verbs: `[get list watch]`
- Resources: `[horizontalpodautoscalers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="horizontal-pod-autoscaler-serviceaccount"></a>Subject: horizontal-pod-autoscaler (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:horizontal-pod-autoscaler`
- Verbs: `[update]`
- Resources: `[horizontalpodautoscalers/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="horizontal-pod-autoscaler-serviceaccount"></a>Subject: horizontal-pod-autoscaler (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:horizontal-pod-autoscaler`
- Verbs: `[get update]`
- Resources: `[*/scale]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="horizontal-pod-autoscaler-serviceaccount"></a>Subject: horizontal-pod-autoscaler (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:horizontal-pod-autoscaler`
- Verbs: `[list]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="horizontal-pod-autoscaler-serviceaccount"></a>Subject: horizontal-pod-autoscaler (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:horizontal-pod-autoscaler`
- Verbs: `[list]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="attachdetach-controller-serviceaccount"></a>Subject: attachdetach-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:attachdetach-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="attachdetach-controller-serviceaccount"></a>Subject: attachdetach-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:attachdetach-controller`
- Verbs: `[list watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="horizontal-pod-autoscaler-serviceaccount"></a>Subject: horizontal-pod-autoscaler (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:horizontal-pod-autoscaler`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="job-controller-serviceaccount"></a>Subject: job-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:job-controller`
- Verbs: `[get list patch update watch]`
- Resources: `[jobs]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="job-controller-serviceaccount"></a>Subject: job-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:job-controller`
- Verbs: `[update]`
- Resources: `[jobs/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="job-controller-serviceaccount"></a>Subject: job-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:job-controller`
- Verbs: `[update]`
- Resources: `[jobs/finalizers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="job-controller-serviceaccount"></a>Subject: job-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:job-controller`
- Verbs: `[create delete list patch watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="job-controller-serviceaccount"></a>Subject: job-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:job-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="resourcequota-controller-serviceaccount"></a>Subject: resourcequota-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:resourcequota-controller`
- Verbs: `[update]`
- Resources: `[resourcequotas/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="namespace-controller-serviceaccount"></a>Subject: namespace-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:namespace-controller`
- Verbs: `[update]`
- Resources: `[namespaces/finalize namespaces/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="attachdetach-controller-serviceaccount"></a>Subject: attachdetach-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:attachdetach-controller`
- Verbs: `[patch update]`
- Resources: `[nodes/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="attachdetach-controller-serviceaccount"></a>Subject: attachdetach-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:attachdetach-controller`
- Verbs: `[list watch]`
- Resources: `[persistentvolumeclaims persistentvolumes]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="node-controller-serviceaccount"></a>Subject: node-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:node-controller`
- Verbs: `[patch update]`
- Resources: `[nodes/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="node-controller-serviceaccount"></a>Subject: node-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:node-controller`
- Verbs: `[patch update]`
- Resources: `[pods/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="node-controller-serviceaccount"></a>Subject: node-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:node-controller`
- Verbs: `[delete list]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="node-controller-serviceaccount"></a>Subject: node-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:node-controller`
- Verbs: `[create get list update]`
- Resources: `[clustercidrs]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="node-controller-serviceaccount"></a>Subject: node-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:node-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="node-controller-serviceaccount"></a>Subject: node-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:node-controller`
- Verbs: `[get]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="persistent-volume-binder-serviceaccount"></a>Subject: persistent-volume-binder (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[create delete get list update watch]`
- Resources: `[persistentvolumes]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="persistent-volume-binder-serviceaccount"></a>Subject: persistent-volume-binder (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[update]`
- Resources: `[persistentvolumes/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="persistent-volume-binder-serviceaccount"></a>Subject: persistent-volume-binder (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[get list update watch]`
- Resources: `[persistentvolumeclaims]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="persistent-volume-binder-serviceaccount"></a>Subject: persistent-volume-binder (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[update]`
- Resources: `[persistentvolumeclaims/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="persistent-volume-binder-serviceaccount"></a>Subject: persistent-volume-binder (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[create delete get list watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="persistent-volume-binder-serviceaccount"></a>Subject: persistent-volume-binder (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[get list watch]`
- Resources: `[storageclasses]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="persistent-volume-binder-serviceaccount"></a>Subject: persistent-volume-binder (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[create delete get update]`
- Resources: `[endpoints]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="persistent-volume-binder-serviceaccount"></a>Subject: persistent-volume-binder (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[create delete get]`
- Resources: `[services]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="deployment-controller-serviceaccount"></a>Subject: deployment-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:deployment-controller`
- Verbs: `[get list update watch]`
- Resources: `[deployments]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemauthenticated-group"></a>Subject: system:authenticated (Group)
- Namespace: ``
- Role: `system:basic-user`
- Verbs: `[create]`
- Resources: `[selfsubjectreviews]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="persistent-volume-binder-serviceaccount"></a>Subject: persistent-volume-binder (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[watch]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="persistent-volume-binder-serviceaccount"></a>Subject: persistent-volume-binder (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:persistent-volume-binder`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="pod-garbage-collector-serviceaccount"></a>Subject: pod-garbage-collector (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:pod-garbage-collector`
- Verbs: `[delete list watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemauthenticated-group"></a>Subject: system:authenticated (Group)
- Namespace: ``
- Role: `system:basic-user`
- Verbs: `[create]`
- Resources: `[selfsubjectaccessreviews selfsubjectrulesreviews]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="pod-garbage-collector-serviceaccount"></a>Subject: pod-garbage-collector (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:pod-garbage-collector`
- Verbs: `[patch]`
- Resources: `[pods/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="pv-protection-controller-serviceaccount"></a>Subject: pv-protection-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:pv-protection-controller`
- Verbs: `[get list update watch]`
- Resources: `[persistentvolumes]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="pv-protection-controller-serviceaccount"></a>Subject: pv-protection-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:pv-protection-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="pvc-protection-controller-serviceaccount"></a>Subject: pvc-protection-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:pvc-protection-controller`
- Verbs: `[get list update watch]`
- Resources: `[persistentvolumeclaims]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="pvc-protection-controller-serviceaccount"></a>Subject: pvc-protection-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:pvc-protection-controller`
- Verbs: `[get list watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="pvc-protection-controller-serviceaccount"></a>Subject: pvc-protection-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:pvc-protection-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="replicaset-controller-serviceaccount"></a>Subject: replicaset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:replicaset-controller`
- Verbs: `[get list update watch]`
- Resources: `[replicasets]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="certificate-controller-serviceaccount"></a>Subject: certificate-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:certificate-controller`
- Verbs: `[create]`
- Resources: `[subjectaccessreviews]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="certificate-controller-serviceaccount"></a>Subject: certificate-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:certificate-controller`
- Verbs: `[update]`
- Resources: `[certificatesigningrequests/approval certificatesigningrequests/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="disruption-controller-serviceaccount"></a>Subject: disruption-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:disruption-controller`
- Verbs: `[get list watch]`
- Resources: `[deployments]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="replicaset-controller-serviceaccount"></a>Subject: replicaset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:replicaset-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="replication-controller-serviceaccount"></a>Subject: replication-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:replication-controller`
- Verbs: `[get list update watch]`
- Resources: `[replicationcontrollers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="replication-controller-serviceaccount"></a>Subject: replication-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:replication-controller`
- Verbs: `[update]`
- Resources: `[replicationcontrollers/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="replication-controller-serviceaccount"></a>Subject: replication-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:replication-controller`
- Verbs: `[update]`
- Resources: `[replicationcontrollers/finalizers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="replication-controller-serviceaccount"></a>Subject: replication-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:replication-controller`
- Verbs: `[create delete list patch watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="replication-controller-serviceaccount"></a>Subject: replication-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:replication-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="local-path-provisioner-service-account-serviceaccount"></a>Subject: local-path-provisioner-service-account (ServiceAccount)
- Namespace: `local-path-storage`
- Role: `local-path-provisioner-role`
- Verbs: `[get list watch]`
- Resources: `[storageclasses]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="namespace-controller-serviceaccount"></a>Subject: namespace-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:namespace-controller`
- Verbs: `[delete get list watch]`
- Resources: `[namespaces]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="deployment-controller-serviceaccount"></a>Subject: deployment-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:deployment-controller`
- Verbs: `[update]`
- Resources: `[deployments/finalizers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="kube-scheduler-serviceaccount"></a>Subject: kube-scheduler (ServiceAccount)
- Namespace: `kube-system`
- Role: `system::leader-locking-kube-scheduler`
- Verbs: `[watch]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="root-ca-cert-publisher-serviceaccount"></a>Subject: root-ca-cert-publisher (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:root-ca-cert-publisher`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="local-path-provisioner-service-account-serviceaccount"></a>Subject: local-path-provisioner-service-account (ServiceAccount)
- Namespace: `local-path-storage`
- Role: `local-path-provisioner-role`
- Verbs: `[create patch]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="resourcequota-controller-serviceaccount"></a>Subject: resourcequota-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:resourcequota-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="route-controller-serviceaccount"></a>Subject: route-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:route-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="service-account-controller-serviceaccount"></a>Subject: service-account-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:service-account-controller`
- Verbs: `[create]`
- Resources: `[serviceaccounts]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="service-account-controller-serviceaccount"></a>Subject: service-account-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:service-account-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="service-controller-serviceaccount"></a>Subject: service-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:service-controller`
- Verbs: `[get list watch]`
- Resources: `[services]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="service-controller-serviceaccount"></a>Subject: service-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:service-controller`
- Verbs: `[patch update]`
- Resources: `[services/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="kube-proxy-serviceaccount"></a>Subject: kube-proxy (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:node-proxier`
- Verbs: `[list watch]`
- Resources: `[endpointslices]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="service-controller-serviceaccount"></a>Subject: service-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:service-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="statefulset-controller-serviceaccount"></a>Subject: statefulset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:statefulset-controller`
- Verbs: `[list watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="statefulset-controller-serviceaccount"></a>Subject: statefulset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:statefulset-controller`
- Verbs: `[get list watch]`
- Resources: `[statefulsets]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="replicaset-controller-serviceaccount"></a>Subject: replicaset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:replicaset-controller`
- Verbs: `[update]`
- Resources: `[replicasets/finalizers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="statefulset-controller-serviceaccount"></a>Subject: statefulset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:statefulset-controller`
- Verbs: `[update]`
- Resources: `[statefulsets/finalizers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="statefulset-controller-serviceaccount"></a>Subject: statefulset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:statefulset-controller`
- Verbs: `[create delete get patch update]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="statefulset-controller-serviceaccount"></a>Subject: statefulset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:statefulset-controller`
- Verbs: `[create delete get list patch update watch]`
- Resources: `[controllerrevisions]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="statefulset-controller-serviceaccount"></a>Subject: statefulset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:statefulset-controller`
- Verbs: `[create get]`
- Resources: `[persistentvolumeclaims]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="statefulset-controller-serviceaccount"></a>Subject: statefulset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:statefulset-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="statefulset-controller-serviceaccount"></a>Subject: statefulset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:statefulset-controller`
- Verbs: `[delete update]`
- Resources: `[persistentvolumeclaims]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="ttl-after-finished-controller-serviceaccount"></a>Subject: ttl-after-finished-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:ttl-after-finished-controller`
- Verbs: `[delete get list watch]`
- Resources: `[jobs]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="ttl-after-finished-controller-serviceaccount"></a>Subject: ttl-after-finished-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:ttl-after-finished-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="kube-proxy-serviceaccount"></a>Subject: kube-proxy (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:node-proxier`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="ttl-controller-serviceaccount"></a>Subject: ttl-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:ttl-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="coredns-serviceaccount"></a>Subject: coredns (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:coredns`
- Verbs: `[list watch]`
- Resources: `[endpoints services pods namespaces]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="kube-proxy-serviceaccount"></a>Subject: kube-proxy (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:node-proxier`
- Verbs: `[list watch]`
- Resources: `[endpoints services]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="coredns-serviceaccount"></a>Subject: coredns (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:coredns`
- Verbs: `[list watch]`
- Resources: `[endpointslices]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemauthenticated-group"></a>Subject: system:authenticated (Group)
- Namespace: ``
- Role: `system:discovery`
- Verbs: `[get]`
- Resources: `[]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[create]`
- Resources: `[leases]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[get update]`
- Resources: `[leases]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[create]`
- Resources: `[endpoints]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[get update]`
- Resources: `[endpoints]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="daemon-set-controller-serviceaccount"></a>Subject: daemon-set-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:daemon-set-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="daemon-set-controller-serviceaccount"></a>Subject: daemon-set-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:daemon-set-controller`
- Verbs: `[create delete get list patch update watch]`
- Resources: `[controllerrevisions]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="daemon-set-controller-serviceaccount"></a>Subject: daemon-set-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:daemon-set-controller`
- Verbs: `[create]`
- Resources: `[pods/binding]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemnodes-group"></a>Subject: system:nodes (Group)
- Namespace: ``
- Role: `system:certificates.k8s.io:certificatesigningrequests:selfnodeclient`
- Verbs: `[create]`
- Resources: `[certificatesigningrequests/selfnodeclient]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[create]`
- Resources: `[tokenreviews]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[create]`
- Resources: `[subjectaccessreviews]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systembootstrapperskubeadmdefault-node-token-group"></a>Subject: system:bootstrappers:kubeadm:default-node-token (Group)
- Namespace: ``
- Role: `system:certificates.k8s.io:certificatesigningrequests:nodeclient`
- Verbs: `[create]`
- Resources: `[certificatesigningrequests/nodeclient]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system:kube-controller-manager`
- Verbs: `[create]`
- Resources: `[serviceaccounts/token]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="kube-dns-serviceaccount"></a>Subject: kube-dns (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:kube-dns`
- Verbs: `[list watch]`
- Resources: `[endpoints services]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[create]`
- Resources: `[leases]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[get update]`
- Resources: `[leases]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[create]`
- Resources: `[endpoints]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[get update]`
- Resources: `[endpoints]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systembootstrapperskubeadmdefault-node-token-group"></a>Subject: system:bootstrappers:kubeadm:default-node-token (Group)
- Namespace: ``
- Role: `system:node-bootstrapper`
- Verbs: `[create get list watch]`
- Resources: `[certificatesigningrequests]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[delete get list watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[create]`
- Resources: `[bindings pods/binding]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[patch update]`
- Resources: `[pods/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[get list watch]`
- Resources: `[replicationcontrollers services]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[get list watch]`
- Resources: `[replicasets]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[get list watch]`
- Resources: `[statefulsets]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[get list watch]`
- Resources: `[poddisruptionbudgets]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[get list watch]`
- Resources: `[persistentvolumeclaims persistentvolumes]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[create]`
- Resources: `[tokenreviews]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[create]`
- Resources: `[subjectaccessreviews]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[get list watch]`
- Resources: `[csinodes]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[get list watch]`
- Resources: `[namespaces]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[get list watch]`
- Resources: `[csidrivers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:kube-scheduler`
- Verbs: `[get list watch]`
- Resources: `[csistoragecapacities]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemmonitoring-group"></a>Subject: system:monitoring (Group)
- Namespace: ``
- Role: `system:monitoring`
- Verbs: `[get]`
- Resources: `[]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-proxy-user"></a>Subject: system:kube-proxy (User)
- Namespace: ``
- Role: `system:node-proxier`
- Verbs: `[list watch]`
- Resources: `[endpoints services]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="kindnet-serviceaccount"></a>Subject: kindnet (ServiceAccount)
- Namespace: `kube-system`
- Role: `kindnet`
- Verbs: `[use]`
- Resources: `[podsecuritypolicies]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-proxy-user"></a>Subject: system:kube-proxy (User)
- Namespace: ``
- Role: `system:node-proxier`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-proxy-user"></a>Subject: system:kube-proxy (User)
- Namespace: ``
- Role: `system:node-proxier`
- Verbs: `[list watch]`
- Resources: `[endpointslices]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemauthenticated-group"></a>Subject: system:authenticated (Group)
- Namespace: ``
- Role: `system:public-info-viewer`
- Verbs: `[get]`
- Resources: `[]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemunauthenticated-group"></a>Subject: system:unauthenticated (Group)
- Namespace: ``
- Role: `system:public-info-viewer`
- Verbs: `[get]`
- Resources: `[]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemserviceaccounts-group"></a>Subject: system:serviceaccounts (Group)
- Namespace: ``
- Role: `system:service-account-issuer-discovery`
- Verbs: `[get]`
- Resources: `[]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:volume-scheduler`
- Verbs: `[get list patch update watch]`
- Resources: `[persistentvolumes]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:volume-scheduler`
- Verbs: `[get list watch]`
- Resources: `[storageclasses]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system:volume-scheduler`
- Verbs: `[get list patch update watch]`
- Resources: `[persistentvolumeclaims]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="daemon-set-controller-serviceaccount"></a>Subject: daemon-set-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:daemon-set-controller`
- Verbs: `[create delete list patch watch]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="attachdetach-controller-serviceaccount"></a>Subject: attachdetach-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:attachdetach-controller`
- Verbs: `[get list watch]`
- Resources: `[csinodes]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="bootstrap-signer-serviceaccount"></a>Subject: bootstrap-signer (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:bootstrap-signer`
- Verbs: `[update]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="bootstrap-signer-serviceaccount"></a>Subject: bootstrap-signer (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:bootstrap-signer`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Namespaced`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="daemon-set-controller-serviceaccount"></a>Subject: daemon-set-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:daemon-set-controller`
- Verbs: `[update]`
- Resources: `[daemonsets/finalizers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="daemon-set-controller-serviceaccount"></a>Subject: daemon-set-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:daemon-set-controller`
- Verbs: `[update]`
- Resources: `[daemonsets/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="daemon-set-controller-serviceaccount"></a>Subject: daemon-set-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:daemon-set-controller`
- Verbs: `[get list watch]`
- Resources: `[daemonsets]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="cronjob-controller-serviceaccount"></a>Subject: cronjob-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:cronjob-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="cronjob-controller-serviceaccount"></a>Subject: cronjob-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:cronjob-controller`
- Verbs: `[delete list]`
- Resources: `[pods]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="cronjob-controller-serviceaccount"></a>Subject: cronjob-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:cronjob-controller`
- Verbs: `[update]`
- Resources: `[cronjobs/finalizers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="cronjob-controller-serviceaccount"></a>Subject: cronjob-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:cronjob-controller`
- Verbs: `[update]`
- Resources: `[cronjobs/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-controller-manager-user"></a>Subject: system:kube-controller-manager (User)
- Namespace: ``
- Role: `system::leader-locking-kube-controller-manager`
- Verbs: `[watch]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="cronjob-controller-serviceaccount"></a>Subject: cronjob-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:cronjob-controller`
- Verbs: `[create delete get list patch update watch]`
- Resources: `[jobs]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="kube-controller-manager-serviceaccount"></a>Subject: kube-controller-manager (ServiceAccount)
- Namespace: `kube-system`
- Role: `system::leader-locking-kube-controller-manager`
- Verbs: `[watch]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="cronjob-controller-serviceaccount"></a>Subject: cronjob-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:cronjob-controller`
- Verbs: `[get list update watch]`
- Resources: `[cronjobs]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="systemkube-scheduler-user"></a>Subject: system:kube-scheduler (User)
- Namespace: ``
- Role: `system::leader-locking-kube-scheduler`
- Verbs: `[watch]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="certificate-controller-serviceaccount"></a>Subject: certificate-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:certificate-controller`
- Verbs: `[delete get list watch]`
- Resources: `[certificatesigningrequests]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="root-ca-cert-publisher-serviceaccount"></a>Subject: root-ca-cert-publisher (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:root-ca-cert-publisher`
- Verbs: `[create update]`
- Resources: `[configmaps]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="certificate-controller-serviceaccount"></a>Subject: certificate-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:certificate-controller`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="replicaset-controller-serviceaccount"></a>Subject: replicaset-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:replicaset-controller`
- Verbs: `[update]`
- Resources: `[replicasets/status]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="certificate-controller-serviceaccount"></a>Subject: certificate-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:certificate-controller`
- Verbs: `[sign]`
- Resources: `[signers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="certificate-controller-serviceaccount"></a>Subject: certificate-controller (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:certificate-controller`
- Verbs: `[approve]`
- Resources: `[signers]`
- Scope: `Cluster`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


## <a name="token-cleaner-serviceaccount"></a>Subject: token-cleaner (ServiceAccount)
- Namespace: `kube-system`
- Role: `system:controller:token-cleaner`
- Verbs: `[create patch update]`
- Resources: `[events]`
- Scope: `Namespaced`
- Risk Level: **LOW**
- Reason: No known high-risk permissions


