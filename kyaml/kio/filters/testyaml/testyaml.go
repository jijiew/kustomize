// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

// Package testyaml contains test data and libraries for formatting
// Kubernetes configuration
package testyaml

var UnformattedYaml1 = []byte(`
spec: a
status:
  conditions:
  - 3
  - 1
  - 2
apiVersion: example.com/v1beta1
kind: MyType
`)

var UnformattedYaml2 = []byte(`
spec2: a
status2:
  conditions:
  - 3
  - 1
  - 2
apiVersion: example.com/v1beta1
kind: MyType2
`)

var UnformattedYaml3 = []byte(`
apiVersion: v1
items:
- apiVersion: v1
  kind: Namespace
  metadata:
    creationTimestamp: "2020-06-04T06:50:09Z"
    name: default
    resourceVersion: "157"
    selfLink: /api/v1/namespaces/default
    uid: a1e97d69-a62f-11ea-b632-42010a8000a7
  spec:
    finalizers:
    - kubernetes
  status:
    phase: Active
- apiVersion: v1
  kind: Namespace
  metadata:
    creationTimestamp: "2020-06-04T06:53:26Z"
    labels:
      name: kit
    name: kit
    resourceVersion: "1702"
    selfLink: /api/v1/namespaces/kit
    uid: 17235129-a630-11ea-b632-42010a8000a7
  spec:
    finalizers:
    - kubernetes
  status:
    phase: Active
- apiVersion: v1
  kind: Namespace
  metadata:
    annotations:
      gke.io/cluster: gke://cpa-kit-dev/us-central1/cpa-kit-dev-us-central1
      kubectl.kubernetes.io/last-applied-configuration: |
        {"apiVersion":"v1","kind":"Namespace",
         "metadata":{"annotations":{"gke.io/cluster":"gke://cpa-kit-dev/us-central1/cpa-kit-dev-us-central1"},
         "name":"kit-server-dev"}}
    creationTimestamp: "2020-06-04T07:17:23Z"
    name: kit-server-dev
    resourceVersion: "7881"
    selfLink: /api/v1/namespaces/kit-server-dev
    uid: 6f8e0a69-a633-11ea-b632-42010a8000a7
  spec:
    finalizers:
    - kubernetes
  status:
    phase: Active
- apiVersion: v1
  kind: Namespace
  metadata:
    creationTimestamp: "2020-06-04T06:50:06Z"
    name: kube-node-lease
    resourceVersion: "40"
    selfLink: /api/v1/namespaces/kube-node-lease
    uid: 9fa5de74-a62f-11ea-b632-42010a8000a7
  spec:
    finalizers:
    - kubernetes
  status:
    phase: Active
- apiVersion: v1
  kind: Namespace
  metadata:
    creationTimestamp: "2020-06-04T06:50:05Z"
    name: kube-public
    resourceVersion: "26"
    selfLink: /api/v1/namespaces/kube-public
    uid: 9f6f664f-a62f-11ea-b632-42010a8000a7
  spec:
    finalizers:
    - kubernetes
  status:
    phase: Active
- apiVersion: v1
  kind: Namespace
  metadata:
    annotations:
      kubectl.kubernetes.io/last-applied-configuration: |
        {"apiVersion":"v1","kind":"Namespace","metadata":{"annotations":{},"name":"kube-system"}}
    creationTimestamp: "2020-06-04T06:50:05Z"
    name: kube-system
    resourceVersion: "143"
    selfLink: /api/v1/namespaces/kube-system
    uid: 9f4b398b-a62f-11ea-b632-42010a8000a7
  spec:
    finalizers:
    - kubernetes
  status:
    phase: Active
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
`)

var UnformattedJSON1 = []byte(`
{
  "spec": "a",
  "status": {"conditions": [3, 1, 2]},
  "apiVersion": "example.com/v1beta1",
  "kind": "MyType"
}
`)

var FormattedYaml1 = []byte(`apiVersion: example.com/v1beta1
kind: MyType
spec: a
status:
  conditions:
  - 3
  - 1
  - 2
`)

var FormattedYaml2 = []byte(`apiVersion: example.com/v1beta1
kind: MyType2
spec2: a
status2:
  conditions:
  - 3
  - 1
  - 2
`)


var FormattedYaml3 = []byte(`apiVersion: v1
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
items:
- apiVersion: v1
  kind: Namespace
  metadata:
    name: default
    creationTimestamp: "2020-06-04T06:50:09Z"
    resourceVersion: "157"
    selfLink: /api/v1/namespaces/default
    uid: a1e97d69-a62f-11ea-b632-42010a8000a7
  spec:
    finalizers:
    - kubernetes
  status:
    phase: Active
- apiVersion: v1
  kind: Namespace
  metadata:
    name: kit
    labels:
      name: kit
    creationTimestamp: "2020-06-04T06:53:26Z"
    resourceVersion: "1702"
    selfLink: /api/v1/namespaces/kit
    uid: 17235129-a630-11ea-b632-42010a8000a7
  spec:
    finalizers:
    - kubernetes
  status:
    phase: Active
- apiVersion: v1
  kind: Namespace
  metadata:
    name: kit-server-dev
    annotations:
      gke.io/cluster: gke://cpa-kit-dev/us-central1/cpa-kit-dev-us-central1
      kubectl.kubernetes.io/last-applied-configuration: |
        {"apiVersion":"v1","kind":"Namespace",
         "metadata":{"annotations":{"gke.io/cluster":"gke://cpa-kit-dev/us-central1/cpa-kit-dev-us-central1"},
         "name":"kit-server-dev"}}
    creationTimestamp: "2020-06-04T07:17:23Z"
    resourceVersion: "7881"
    selfLink: /api/v1/namespaces/kit-server-dev
    uid: 6f8e0a69-a633-11ea-b632-42010a8000a7
  spec:
    finalizers:
    - kubernetes
  status:
    phase: Active
- apiVersion: v1
  kind: Namespace
  metadata:
    name: kube-node-lease
    creationTimestamp: "2020-06-04T06:50:06Z"
    resourceVersion: "40"
    selfLink: /api/v1/namespaces/kube-node-lease
    uid: 9fa5de74-a62f-11ea-b632-42010a8000a7
  spec:
    finalizers:
    - kubernetes
  status:
    phase: Active
- apiVersion: v1
  kind: Namespace
  metadata:
    name: kube-public
    creationTimestamp: "2020-06-04T06:50:05Z"
    resourceVersion: "26"
    selfLink: /api/v1/namespaces/kube-public
    uid: 9f6f664f-a62f-11ea-b632-42010a8000a7
  spec:
    finalizers:
    - kubernetes
  status:
    phase: Active
- apiVersion: v1
  kind: Namespace
  metadata:
    name: kube-system
    annotations:
      kubectl.kubernetes.io/last-applied-configuration: |
        {"apiVersion":"v1","kind":"Namespace","metadata":{"annotations":{},"name":"kube-system"}}
    creationTimestamp: "2020-06-04T06:50:05Z"
    resourceVersion: "143"
    selfLink: /api/v1/namespaces/kube-system
    uid: 9f4b398b-a62f-11ea-b632-42010a8000a7
  spec:
    finalizers:
    - kubernetes
  status:
    phase: Active
`)

var FormattedJSON1 = []byte(`{
  "apiVersion": "example.com/v1beta1",
  "kind": "MyType",
  "spec": "a",
  "status": {
    "conditions": [
      3,
      1,
      2
    ]
  }
}
`)
