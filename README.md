# gcs-sync-sidecar ![example workflow](https://github.com/fsn-capital/gcs-sync-sidecar/workflows/CI/badge.svg)

## Getting Started

Install helm secrets [here](https://github.com/jkroepke/helm-secrets/wiki/Installation).

## Rendering Templates

``` helm secrets install --debug --dry-run -f gcs-sync/secrets.yaml -f gcs-sync/values.yaml gcs-sync gcs-sync ```

## Deployment 

For deployment another repository will be created.