# Drone Plugin for Koki Short

Define all your Kubernetes manifests using Koki Short, and use this drone plugin to automatically convert the Short manifests into K8s manifests on the fly

# Use Koki Short in your Projects

```
workspace:
  base: /go
  path: src/github.com/koki/short-drone-plugin

  pipeline:
    koki-short:
      image: kokster/short-drone-plugin:0.3.0
      files:
        example.yaml
```

This will convert `example.yaml` to `kube_example.yaml`, which is the kubernetes manifest equivalent to `example.yaml`. The next steps in the pipeline will now be able to use `kube_example.yaml` 
