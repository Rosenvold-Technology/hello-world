
# 🌟 World's Most Over‑Engineered Hello Website 🌟

This repository contains a **Kubernetes‑native**, **Go‑powered**, and **visually stunning** hello world application
designed to be far more complicated than necessary.

## Features

* **Gin‑powered HTTP server** with graceful shutdown and structured logs.
* **Prometheus metrics** and **OpenTelemetry tracing** baked in.
* **Multi‑stage Docker build** producing a tiny distroless image.
* **Responsive gradient UI** with canvas‑based confetti celebrations.
* **Kubernetes manifests** featuring liveness/readiness probes, HPA,
  Ingress, and ConfigMap‑driven configuration.
* **Makefile** shortcuts for local kind clusters.

## Quick Start

```bash
# build and load into kind
make kind

# deploy
make deploy

# celebrate
kubectl port-forward svc/hello-website 8080:80
open http://localhost:8080
```

Enjoy your drastically over‑complicated hello world! 🎉
