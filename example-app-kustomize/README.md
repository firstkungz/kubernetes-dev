# How to Update Overlays Configuration

Development team must be update application-kubernetes.yaml, config.env, secret.env, and/or set_resources.yaml
* overlays/env/configs/application-kubernetes.yaml - for provide configmap files.
* overlays/env/configs/config.env - for provide configmap environment variables.
* overlays/env/secrets/secret.env - for provide secret environment variables.
* overlays/env/patches/set_resources.yaml - for set replicas number and cpu/memory limit and request.

## Application-kubernetes.yaml
application-kubernetes.yaml file under the overlays/env/configs directory will be using to generate configmap as a file in container volume.  
Example of application-kubernetes.yaml:
```bash
spring:
  application:
    name: mockup-services
```

## Config.env
config.env file under the overlays/env directory will be using to generate configmap as an environment variable.  
Example of config.env:
```bash
MYSQL_HOST=xxx-platform-mysql-sit.xxx.nonprod.gcp.ktbcloud
REDIS_HOST=xxx-platform-redis-sit.xxx.nonprod.gcp.ktbcloud
```

## Secret.env
secret.env file under the overlays/env directory will be using to generate secret as an environment variable.  
Example of secret.env:
```bash
KAFKA_SASL_USERNAME=sit_msg_xxxx
KAFKA_SASL_PASSWORD=xxxxxx
KAFKA_SSL-CA-PEM=-----BEGIN CERTIFICATE-----\xxxxxx\n-----END CERTIFICATE-----
```

## Set_resources.yaml
In some cases, development team might need to update *replicas* number or *cpu/memory* for performance testing. Development team must be update in set_resources.yaml under the overlays/env that you wish to update.  
Example of set_resources.yaml:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app
spec:
  replicas: 1
  template:
    spec:
      containers:
        - name: app 
          resources:
            limits:
              cpu: 500m
              memory: 256Mi
            requests:
              cpu: 250m
              memory: 128Mi

```
# Project's Additional Request 
Note: Development team must not edit any files in *base* directory.

* base/resources/deployment.yaml - livenessProbe
* base/resources/deployment.yaml - readinessProbe
* base/resources/deployment.yaml - resources

```yaml
apiVersion: apps/v1
kind: Deployment
spec:
  template:
    spec:
      livenessProbe:
        httpGet:
          path: /actuator/health
          port: http
        timeoutSeconds: 10
        initialDelaySeconds: 10
      readinessProbe:
        httpGet:
          path: /actuator/health/health
          port: http
        timeoutSeconds: 10
        initialDelaySeconds: 10
      resources:
        limits:
          cpu: 500m
          memory: 128Mi
        requests:
          cpu: 400m
          memory: 64Mi
```
