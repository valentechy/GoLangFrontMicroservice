# GoLangFrontMicroservice

Microservicio de prueba utilizado para artículo en www.valentinpalacios.com

## Construir imagen

```bash
docker build -t your-docker-repo/frontend-microservice:1.0.0 .
```

## Subir imagen

```bash
docker push your-docker-repo/frontend-microservice:1.0.0
```

## Desplegar micro en Kubernetes

```bash
kubectl apply -f K8s/deployment.yaml
kubectl apply -f K8s/service.yaml
kubectl apply -f K8s/ingress.yaml
```