# GCS Sync [![CI](https://github.com/fsn-capital/gcs-sync-sidecar/workflows/CI/badge.svg)](https://github.com/fsn-capital/gcs-sync-sidecar/actions/workflows/ci.yml) [![Chart Builder](https://github.com/fsn-capital/gcs-sync-sidecar/workflows/Chart%20Builder/badge.svg)](https://github.com/fsn-capital/gcs-sync-sidecar/actions/workflows/chart.yml)

This repository contains a custom application to sync files from/to GCS buckets.

## Getting Started

* Install helm secrets [here](https://github.com/jkroepke/helm-secrets/wiki/Installation).
* To build/push docker container run ``` make build ```.
* Add chart repository using ``` helm repo add gcs-sync https://fsn-capital.github.io/gcs-sync-sidecar ```.
* Retrieve latest chart using ``` helm repo update ```

## Environment Variables 

* ``` SOURCE ```: Source directory or bucket
* ``` DESTINATION ```: Destination directory or bucket
* ``` GOOGLE_APPLICATION_CREDENTIALS ```: Points to the directory that contains service account json key. 

## Getting GPG Key

Please contact ``` Mert Kayhan ```.

## Rendering Templates

``` helm secrets install --debug --dry-run -f gcs-sync/secrets.yaml -f gcs-sync/values.yaml gcs-sync gcs-sync ```

## Deployment 

For deployment another repository will be created.
