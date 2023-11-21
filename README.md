# kubernetes-dev

Prerequisites 
- Homebrew
https://brew.sh/

- Docker https://docs.docker.com/desktop/install/mac-install/

- gcloud command - Install the gcloud CLI | Google Cloud 

    Then you will sent to website, select the user need to login with google cloud.
    ```
    gcloud auth login
    gcloud auth list
    ```

- kubectl + kustomize command
    ```
    brew install kubectl
    brew install kustomize
    ```

    Check that kubectl, kustomize is installed, will get the version of installed command
    ```
    kubectl version
    kustomize version
    ```

## Exercise 1 Playing with Kubernetes config and Command

### Connect to Kubernetes on GCP

Import config from GKE

```
gcloud container clusters get-credentials poc-cluster-cmmn --zone asia-southeast1-b --project poc-innovation-iot
```

verfiy that kubectl command is connected with selected GKE cluster

```
kubectl get nodes
```

---

### 1.1 - Create the components

Try understanding what you will create in app1, app2

Apply (Create / Update) the component specified in yaml files
```
kubectl apply -f configmap.yaml
kubectl apply -f secret.yaml
kubectl apply -f service.yaml
kubectl apply -f deployment.yaml
```

### 1.2 - Try to get list resources with command

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

### 1.3 - Interaction with component

Execute command pods

```
kubectl exec -it {pods_name} {command}
```

Entering pods (bash or others depend on how docker built)
You can check either configuration type (env,file) and the service connection between them by try to ping.
```
kubectl exec -it {pods_name} bash
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


## Exercise 2 Troubleshooting

There're some of cases found, during the hands-on will showing about the investigation step.


## Exercise 3 Try kustomize for deployment

This exerise willl run under kustomize-config folder

At first you need to create namespace for simulate deploying to SIT / UAT
```
kubectl create namespace app-{yourname}-sit
kubectl create namespace app-{yourname}-uat
```


Preview apply script that kustomize built
```
kustomize build overlays/{env}
```

Example for overlays sit
```
kustomize build overlays/sit
```

Deploy microservice to kubernetes by kustomize

```
kubectl apply -n {namespace} -k overlays/{env}
```

So the rest of commands you need for deploy both SIT / UAT are
```
kubectl apply -n app-{yourname}-sit -k overlays/sit
kubectl apply -n app-{yourname}-uat -k overlays/uat
```


## Exercise 4 From code to running service
Now if we have the http service that can running, how we can make it from nothing to running in kubernetes

\* 
This exerise willl run under example-app, example-app-kustomize folder, This example is golang and echo.

### Health check
This example is for golang
The thing that we need is the http server that running and response HTTP 200 OK status.
Kubernetes will try to ping that the service is running, If not running it will restart the pods


```
e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Health Check!")
})
```
    
### Config injection

The next thing we need to do is how config injection into services

Service read config (contain both file and env example, you can select which ways to do)
```
	// Loading env config of path /configs/config.yaml
	viper.AddConfigPath("/configs")
	viper.AddConfigPath(".")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("yaml")   // Look for specific type
	viper.ReadInConfig()

	// Loading env config of key SOMECONFIGKEY
	someConfig := os.Getenv("SOMECONFIGKEY")
	fmt.Printf("Env Config is : %s", someConfig)
```


### Build docker image

The thing before we can run service in kubernetes, we need to packaging it inside the container, so we need to build the dockerfile for packaging our service.

You can see some explanation in Dockerfile

```
docker build -t asia.gcr.io/poc-innovation-iot/training-kube/example-app:v1.6 ./
```

After you built the docker image, then we need to push the docker image to cloud

```
docker push asia.gcr.io/poc-innovation-iot/training-kube/example-app:v1.6
```


### Deployment by kustomize

- You need to ensure that image, tag that we push in the kustomization.yaml is correct

file example-app-kustomize/overlays/sit/kustomization.yaml
```
images:
  - name: asia.gcr.io/poc-innovation-iot/training-kube/example-app
    newTag: v1.x
```

- Then ensure configuration you need to using is valid for the environment you will be deployed


Then verify the kustomize build is valid before deploy.

```
kustomize build ./kustomize-config/overlays/sit 
```

Then we will apply this by kustomize for deploy the service
```
kubectl apply -n {namespace} -k overlays/sit
```


(Bonus) Expose the service we already do to external by ingress.
```
kubectl apply -n {namespace} -f ingress.yaml
```