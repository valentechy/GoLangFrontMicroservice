# GoLangFrontMicroservice

Microservicio de prueba utilizado para artículo en www.valentinpalacios.com

## Construir imagen

```
docker build -t your-docker-repo/frontend-microservice:1.0.0 .
```

## Subir imagen

```
docker push your-docker-repo/frontend-microservice:1.0.0
```

## Desplegar micro en Kubernetes

```
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f ingress.yaml
```