#!/bin/bash
export GOOGLE_CLOUD_PROJECT=$1

docker build . --tag "gcr.io/$GOOGLE_CLOUD_PROJECT/go-pets-teva"
docker push "gcr.io/$GOOGLE_CLOUD_PROJECT/go-pets-teva:latest"

cd terraform
terraform init
terraform apply -auto-approve
cd ..