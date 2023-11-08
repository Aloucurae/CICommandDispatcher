#!/bin/bash
docker  build  --tag registry.alexjonas.com.br/command-dispatcher:latest .
# docker buildx build --platform linux/amd64,linux/arm64  --tag registry.alexjonas.com.br/command-dispatcher:latest --tag registry.alexjonas.com.br/command-dispatcher:v1  --push .