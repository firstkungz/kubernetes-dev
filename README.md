# kubernetes-dev

### Connect to Kubernetes on GCP

Import service account

```
gcloud auth activate-service-account --key-file=training-kube-sa.json
```

Verify that service account is active on gcloud command

```
ACTIVE ACCOUNT
* training-kube@poc-innovation-iot.iam.gserviceaccount.com
```

Import config from GKE

```
gcloud container clusters get-credentials poc-cluster-cmmn --zone asia-southeast1-b --project poc-innovation-iot
```

verfiy that kubectl command is connected with selected GKE cluster

```
kubectl get nodes
```

---

## Kubernetes Basic command

### Example 1 - Try to get list resources with command

Component in every namespaces

```
kubectl get {component} -A
```

Specific type and name spaces

```
kubectl get {component} -n {namespace}
```

Describe information of component

```
kubectl describe {component} {component_name}
```

### Example 2 - Interaction with component

Execute command pods

```
kubectl exec -it {pods_name} {command}
```

Entering pods (bash or others depend on how docker built)

```
kubectl exec -it {pods_name} bash
```

Create / Update / Delete component (more information in deployment part)

```
    kubectl apply -f deployment.yaml
    kubectl delete -f deployment.yaml
```

Scale pods from deployment

```
kubectl scale --replicas=3 deployment/test
```

Edit resource on the fly

```
kubectl edit svc/docker-registry
```

---

## Apply kustomize for deployment

Deploy microservice to kubernetes by kustomize

```
kustomize build {path_to_kustomize} | kubectl apply -n {namespace} -f -
```

After this you need to try yourself for editing some of kustomize following readme.MD in example-app-kustomize
