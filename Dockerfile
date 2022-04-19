FROM golang:1.18.1-buster as builder
WORKDIR /app
COPY go.mod go.sum main.go ./
COPY common ./common
RUN go mod download 
RUN go build -o app main.go


FROM golang:1.18.1-buster as final 
WORKDIR /app
COPY --from=builder /app/app ./app
# download crcmod
RUN rm -rf /var/lib/apt/lists/* && apt-get update -y && apt-get install -y python3-pip && pip3 install --no-cache-dir -U crcmod && apt remove -y python3-pip && rm -rf /var/lib/apt/lists/*
# download sdk 
ADD https://dl.google.com/dl/cloudsdk/release/google-cloud-sdk.tar.gz /tmp/google-cloud-sdk.tar.gz 
# install sdk
RUN mkdir -p /usr/local/gcloud \
    && tar -C /usr/local/gcloud -xvf /tmp/google-cloud-sdk.tar.gz \
    && /usr/local/gcloud/google-cloud-sdk/install.sh

# Adding the package path to local
ENV PATH $PATH:/usr/local/gcloud/google-cloud-sdk/bin
CMD [ "./app" ]