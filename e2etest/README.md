# End-to-End Testing

The end-to-end testing is based on staging test web pages, test
JavaScript, a test version of the web app, and the production database. 
This allows for serving static media, including HTML pages, from the same
container that AJAX calls are made to. This is simplified with nearly the full
production setup with a load balancer fronting static media in GCS and
AJAX calls from the web application as the prod setup.

## Generate HTML Pages and Build Web App

From the top level directoary, First follow the procedure in the main README.md
to generate the HTML files for the site.

```shell
nohup bin/cnreader.sh &
```

Check on status while doing other stuff

```shell
tail -f nohup.out
```

Build the web app, as per the main README.md instructions

```shell
export BUILD_ID=r160
gcloud builds submit --config cloudbuild.yaml . \
  --substitutions=_IMAGE_TAG="$BUILD_ID"
```

Deploy the web app to Cloud Run

```shell
PROJECT_ID=[Your project]
IMAGE=gcr.io/${PROJECT_ID}/cn-app-image:${BUILD_ID}
SERVICE=e2etest
REGION=us-central1
INSTANCE_CONNECTION_NAME=[Your connection]
DBUSER=[Your database user]
DBPASSWORD=[Your database password]
DATABASE=[Your database name]
MEMORY=400Mi
TEXT_BUCKET=[Your GCS bucket name for text files]
gcloud run deploy --platform=managed $SERVICE \
--image $IMAGE \
--region=$REGION \
--memory "$MEMORY" \
--add-cloudsql-instances $INSTANCE_CONNECTION_NAME \
--set-env-vars INSTANCE_CONNECTION_NAME="$INSTANCE_CONNECTION_NAME" \
--set-env-vars DBUSER="$DBUSER" \
--set-env-vars DBPASSWORD="$DBPASSWORD" \
--set-env-vars DATABASE="$DATABASE" \
--set-env-vars TEXT_BUCKET="$TEXT_BUCKET" \
--set-env-vars CNREADER_HOME="/"
```

The output will give the location of the URL that AJAX requests will be sent.
You can test it withe the command

```shell
curl $URL/find/?query=你好
```

Create a CS bucket for static files

```
export BUCKET={your bucket}
gsutil mb gs://$BUCKET
nohup e2etest/e2ecopy.sh &
gsutil web set -m index.html -e 404.html gs://$BUCKET
```

After the first time setup you only need to run 

```shell
nohup e2etest/e2ecopy.sh &
```

## Load balancer

Create a load balancer following instructions in 
[Setting up serverless NEGs](https://cloud.google.com/load-balancing/docs/negs/setting-up-serverless-negs)

Create a 
[Google-managed SSL certificate](https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs)

```shell
CERTIFICATE_NAME=[your cert name]
DOMAIN=[your domain]
gcloud compute ssl-certificates create $CERTIFICATE_NAME \
    --description="End to end testing" \
    --domains="$DOMAIN" \
    --global
```

Get an external IP address

```shell
IP_NAME=[your IP name]
gcloud compute addresses create $IP_NAME \
    --ip-version=IPV4 \
    --global
```

Configure a backend bucket for the LB

```shell
BACKEND_BUCKET=[your backend bucket]
gcloud compute backend-buckets create $BACKEND_BUCKET --gcs-bucket-name $BUCKET
```

Create a URL map with the backend bucket as default

```shell
URL_MAP=[your URL map name]
gcloud compute url-maps create $URL_MAP \
    --default-backend-bucket $BACKEND_BUCKET
```

Create a serverless Network Endpoint Group

```shell
NEG=[name of NEG]
gcloud beta compute network-endpoint-groups create $NEG \
    --region=us-central1 \
    --network-endpoint-type=serverless \
    --cloud-run-service=$SERVICE
```

Create an LB backend service

```shell
LB_SERVICE=[name of service]
gcloud compute backend-services create $LB_SERVICE --global
```

Add the NEG to the backend service

```shell
gcloud beta compute backend-services add-backend $LB_SERVICE \
    --global \
    --network-endpoint-group=$NEG \
    --network-endpoint-group-region=$REGION
```

Create a matcher for the NEG backend service
```shell
MATCHER_NAME=[your matcher name]
gcloud compute url-maps add-path-matcher $URL_MAP \
    --default-backend-bucket $BACKEND_BUCKET \
    --path-matcher-name $MATCHER_NAME \
    --path-rules="/find/*=$LB_SERVICE,/findadvanced/*=$LB_SERVICE,/findmedia/*=$LB_SERVICE,/findsubstring,/findtm=$LB_SERVICE"
```

Create a target HTTPS proxy

```shell
PROXY_NAME=[your proxy]
gcloud compute target-https-proxies create $PROXY_NAME \
    --ssl-certificates=$CERTIFICATE_NAME \
    --url-map=$URL_MAP
```

Create a HTTPS forwarding rule

```shell
FWD_RULE_NAME=[your forwarding rule]
gcloud compute forwarding-rules create $FWD_RULE_NAME \
    --address=$IP_NAME \
    --target-https-proxy=$PROXY_NAME \
    --global \
    --ports=443
```

It may take a few minutes for the LB to be configured. Testing the setup with
the command

```shell
curl https://$DOMAIN
```

or navigate with your browser.

## DEPRECATED Run Locally

To run locally, first copy important static files to the static directory

```shell
./e2ecopy.sh
```

Build the Go web app

```shell
go build
```

Start the web app

```shell
./e2etest
```

Navigate to http://localhost:8080/ to try out the web interface.
