# kubernetes-dev

Prerequisites for this training
Homebrew
https://brew.sh/

Docker
https://docs.docker.com/desktop/install/mac-install/

```
gcloud command
```

Install the gcloud CLI | Google Cloud
ทดลอง run gcloud command ว่าใช้งานได้ปกติ จะแสดงว่าเรายังไม่ login อะไรเลย (ถ้าพึ่งลง)

```
gcloud auth list
```

และทำการ login ด้วย account ตัวเอง

gcloud auth login
ให้เลือก Y เพื่อ login ด้วย account google ของเราและเลือก account google ที่เราจะใช้งานในหน้า website (หรือ login ใน website ถ้าไม่เคย login with google)
เมื่อเรียบร้อยทดสอบ run gcloud auth list อีกทีจะแสดง email ที่เราพึ่ง login ไป เป็นอันเสร็จสิ้น

kubectl command

```
brew install kubectl
```

Check version command

```
kubectl version
```

ให้ check เลข version ให้ตัว GitVersion ของ Client >= 1.19.1 น่าจะ ok แล้ว

```
Client Version: version.Info{Major:"1", Minor:"19", GitVersion:"v1.19.1", GitCommit:"206bcadf021e76c27513500ca24182692aabd17e", GitTreeState:"clean", BuildDate:"2020-09-09T19:10:21Z", GoVersion:"go1.15.1", Compiler:"gc", Platform:"darwin/amd64"}
```

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

Build kustomize into script for deployment
Example for overlays sit

```
kustomize build ./example-app-kustomize/overlays/sit
```

Deploy microservice to kubernetes by kustomize

```
kustomize build {path_to_kustomize} | kubectl apply -n {namespace} -f -
```

After this you need to try yourself for editing some of kustomize following readme.MD in example-app-kustomize
and you're finish this :)
