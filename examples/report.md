# Permiflow RBAC Audit Report

## 📘 Table of Contents

- [system:nodes (Group)](#subject-systemnodes-group)
- [generic-garbage-collector (ServiceAccount)](#subject-generic-garbage-collector-serviceaccount)

---

## <a name="subject-systemnodes-group"></a>Subject: system:nodes (Group)

- Namespace: ``
- Role: `kubeadm:kubelet-config`
- Verbs: `[get list]`
- Resources: `[configmaps]`
- Scope: `Namespaced`
- Risk Level: **LOW**
- ℹ️ Note: This is a cluster-wide subject (Group) granted access to this namespace.

## <a name="subject-generic-garbage-collector-serviceaccount"></a>Subject: generic-garbage-collector (ServiceAccount)

- Namespace: `kube-system`
- Role: `system:controller:generic-garbage-collector`
- Verbs: `[delete get list patch update watch]`
- Resources: `[*]`
- Scope: `Namespaced`
- Risk Level: **HIGH**
