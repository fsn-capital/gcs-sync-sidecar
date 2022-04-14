FROM golang:1.18.1-buster as builder
WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod download 
RUN go build -o app main.go


FROM golang:1.18.1-buster as final 
WORKDIR /app
COPY --from=builder /app/app ./app
# Downloading gcloud package
RUN curl https://dl.google.com/dl/cloudsdk/release/google-cloud-sdk.tar.gz > /tmp/google-cloud-sdk.tar.gz

# Installing the package
RUN mkdir -p /usr/local/gcloud \
    && tar -C /usr/local/gcloud -xvf /tmp/google-cloud-sdk.tar.gz \
    && /usr/local/gcloud/google-cloud-sdk/install.sh

# Adding the package path to local
ENV PATH $PATH:/usr/local/gcloud/google-cloud-sdk/bin
CMD [ "./app" ]