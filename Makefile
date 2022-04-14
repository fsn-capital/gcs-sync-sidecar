TAG ?= v0.0.1-alpha

.PHONY: build
build: 
	docker buildx build --platform linux/amd64,linux/arm64 --push -t gcr.io/fsn-insight/gcs-sync:${TAG} .