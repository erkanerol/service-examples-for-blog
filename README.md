# Example services to understand how to debug k8s services locally


## Deployments of base setup
```
kubectl run service-back --image=erkanerol/service-back:v1 --port=8080 --expose=true
kubectl run service-middle --image=erkanerol/service-middle:v1 --port=8081 --expose=true
kubectl run service-front --image=erkanerol/service-front:v1 --port=8082 --expose=true
```