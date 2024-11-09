# GoLangFrontMicroservice

Microservicio de prueba utilizado para artículo en www.valentinpalacios.com

Este microservicio está escrito en GoLang y actúa como un frontend que recibe peticiones en dos rutas (`/primario` y `/secundario`). Cada ruta llama a un backend que devuelve un color en formato JSON, y el frontend muestra una página web sencilla con el color de fondo especificado.

## Requisitos

- Docker
- kind (Kubernetes IN Docker)
- kubectl

## Construir y ejecutar el microservicio

### 1. Construir la imagen Docker

```bash
docker build -t frontend-microservice:1.0.0 .
```

### 2. Cargar la imagen en el clúster de kind

```bash
kind load docker-image frontend-microservice:1.0.0
```

### 3. Desplegar el microservicio en Kubernetes

```bash
kubectl apply -f K8s/deployment.yaml
kubectl apply -f K8s/service.yaml
kubectl apply -f K8s/ingress.yaml
```

### 4. Verificar el despliegue

Puedes verificar que los pods están corriendo correctamente con el siguiente comando:

```bash
kubectl get pods
```

### 5. Acceder al microservicio

Para acceder al microservicio, necesitas configurar el Ingress Controller en tu clúster de kind. Puedes seguir la [documentación oficial de kind](https://kind.sigs.k8s.io/docs/user/ingress/) para configurar el Ingress Controller.

Una vez configurado, puedes acceder al microservicio a través de las rutas `/primario` y `/secundario` en el dominio que hayas configurado para el Ingress.

## Notas

- Asegúrate de que el backend esté corriendo y accesible desde el frontend para que las rutas funcionen correctamente.
- Puedes modificar los archivos de configuración de Kubernetes en el directorio `K8s` según tus necesidades.