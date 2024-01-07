### Dockerfile builer pre-build images

```sh
docker build -t bitbus/paopao-ce-backend-builder:latest -f Dockerfile.backend-builder .
docker build -t bitbus/paopao-ce-backend-runner:latest -f Dockerfile.backend-runner .
docker build -t bitbus/paopao-ce-allinone-runner:latest -f scripts/docker/Dockerfile.allinone-runner .
```