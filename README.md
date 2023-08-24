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


- kubectl command
    ```
    brew install kubectl
    ```

    Check that kubectl is installed

    ```
    kubectl version
    ```

## Exercise 1 Kubernetes Basic command line

This exerise willl run under example-kubernetes folder

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

## Exercise 2 Try kustomize for deployment

This exerise willl run under example-kustomize folder

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


## Exercise 3 From code to running service
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

	// Try to Dump all config from file by viper
	fmt.Println(viper.AllSettings())

	// Loading env config of key SOMECONFIGKEY
	someConfig := os.Getenv("SOMECONFIGKEY")
	fmt.Printf("Env Config is : %s", someConfig)
```


### Build docker image

The thing before we can run service in kubernetes, we need to packaging it inside the container, so we need to build the dockerfile for packaging our service.

You can see some explanation in Dockerfile

```
docker build -t asia.gcr.io/poc-innovation-iot/training-kube/example-app:v1.5 .
```

After you built the docker image, then we need to push the docker image to cloud

```
docker push asia.gcr.io/poc-innovation-iot/training-kube/example-app:v1.5
```


### Deployment by kustomize

You need to ensure that image, tag that we push in the kustomization.yaml is correct

file example-app-kustomize/overlays/sit/kustomization.yaml
```
images:
  - name: asia.gcr.io/poc-innovation-iot/training-kube/example-app
    newTag: v1.4
```

Then the configuration you need to enter the require configuration
- Env - go to example-app-kustomize/overlays/sit/configs/config.env file
- File - go to example-app-kustomize/overlays/sit/configs/config.yaml file

For specifies the resources to be use example-app-kustomize/overlays/sit/patches/set_resources.yaml
```
resources:
    limits:
        cpu: 200m
        memory: 512Mi
    requests:
        cpu: 200m
        memory: 512Mi
```

Then apply the kustomize to kubernetes

```
kustomize build ./example-app-kustomize/overlays/sit | kubectl apply -n {namespace} -f -
```

Then we finished on deployment service into kubernetes