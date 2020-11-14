# Chinese Notes Project
## About the project
Chinese-English dictionary and digital library of literary Chinese classic
and historic texts for English speakers. It includes a framework that can be
re-used for other corpora including a template system,
a Chinese-English dictionary, a corpus management system written in Go, 
web pages for learning grammar and such, a collection of texts, and a system for
looking up the words in the texts by mouse over and popover by clicking on the
words. Please join the low volume email group
[chinesenotes-announce](https://groups.google.com/forum/#!forum/chinesenotes-announce)
for announcements of new features and other updates.

The software and dictionary powers several web sites with different corpora:

1. http://chinesenotes.com - for literary Chinese documents and a small amount
   of modern Chinese
2. http://ntireader.org - for the Taisho version of the Chinese Buddhist Canon
3. http://hbreader.org - for Venerable Master Hsing Yun's collected writings
4. http://www.primerbuddhism.org - A Primer in Chinese Buddhist Writings

The Chinese Notes software includes several components:

1. *cnreader* - a Go program to analyze the library corpus, performing text
   segmentation, matching Chinese text to dictionary entries, indexing the text,
   and generating HTML files for reading the texts. This utility is something
   like [Hugo](https://gohugo.io/) or the [Sphinx Python documentation 
   generator](http://www.sphinx-doc.org/en/master/).
2. *cnweb* - a Go web application for reading and searching the texts and
   looking up dictionary entries. Dictionary data, library metadata, and a 
   text retrieval index is loaded into a SQL database to support the web site.

Web application software for searching the dictionary and corpus is at
https://github.com/alexamies/chinesenotes-go

A JavaScript library to help in presenting the web application is at
https://github.com/alexamies/chinesedict-js

Python ulitities for analysis of text in the structure here are at
https://github.com/alexamies/chinesenotes-python

For a description of how the framework can be used for other corpora see
[FRAMEWORK.md](FRAMEWORK.md).

## Acknowldegements
Major sources used directly in the dictionary whose professional and freely
shared work is gratefully acknowledged include:

- [CC-CEDICT Chinese - English dictionary](http://cc-cedict.org/wiki/), shared
under the [Creative Commons Attribution-Share Alike 3.0 
License](http://creativecommons.org/licenses/by-sa/3.0/)
- [Chinese Wikisource](https://zh.wikisource.org/wiki/Wikisource) from the
Wikimedia Foundation, aslo under a Creative Commons license
- [Unihan Database](http://www.unicode.org/charts/unihan.html) from the Unicode
Consortium under a freely reusable license
- [教育部國語辭典](http://resources.publicense.moe.edu.tw/dict_reviseddict_download.html)
Republic of China Ministry of Education Standard Chinese Dicionary, aslo under a
Creative Commons license

## Installing the Chinese Notes web site

### Environment
Installation instructions are for Debian.

### Install git on the host and checkout the code base
```
git clone git://github.com/alexamies/chinesenotes.com

cd chinesenotes.com
export CNREADER_HOME=`pwd`
```

### Go command line tool
Generates markup for HTML page popovers

Install go (see https://golang.org/doc/install)

For more details on the corpus organization and command line tool to process
it see corpus/CORPUS-README.md and go/src/cnreader/README-go.md. Basic use:

```
$ cd go/src/cnreader

$ go build

$ ./cnreader -all
```

### Project Organization

/bin
 - Wrappers for command line Go programs, sich as the bin/cnreadher.sh script

/corpus
 - raw text files for making up the text corpus

/data
 - dictionary data files

/data/corpus
 - metadata files describing the structure of the corpus

/go
 - source code for the command line tool for analysis of the corpus and
   generation of HTML pages

/html
 - raw HTML content minus headers, footers, menus, etc. This is the source
   HTML before application of templates for pages that are not considered part
   of the corpus. The home, about, references, and relates pages are here.

/html/material-templates
  - Go templates for generation of HTML files using material design lite
    styles. This directory can be overridden by setting an environment variable
    named TEMPLATE_HOME with the path relative to the project home.

/index
 - Output from corpus analysis

/web-resources
 - Static resources, including CSS, JavaScript, image, and sound files

/web-staging
 - Generated HTML files. Many but not all files are generated with the Go
  command line tool cnreader.
  This is a default that can be overridden by setting an environment varialbe 
  named WEB_DIR with the path relative to the project home.


## Containerization
Containerization has now replaced the old system of deployment directly on a
virtual machine.

### Local Development Environment or Build Machine

The build machine builds the source code and Docker images. Get the project
files from GitHub:

```
git clone https://github.com/alexamies/chinesenotes.com.git
```

### Install Docker
[Docker Documentation](https://docs.docker.com/engine/installation/linux/docker-ce/ubuntu/)

Instance name: BUILD_VM

Image type: Ubuntu Zesty 17.04

```
gcloud compute --project $PROJECT ssh --zone $ZONE $BUILD_VM
sudo apt-get remove docker docker-engine docker.io
sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

sudo apt-get update
sudo apt-get install docker-ce

sudo usermod -a -G docker ${USER}

# Run a test
docker run hello-world
```

## Database
For database setup see 
https://github.com/alexamies/chinesenotes-go

Compile the library document files and tiles into a tab separated file for
loading into the database with the Python program

```shell
bin/doc_list.sh
```

### Text Files for Full Text Search
Copy the text files to an object store. If you prefer not to use GCS, then you
can use the local file system on the application server. The instructions here
are for GCS. See [Authenticating to Cloud Platform with Service
Accounts](https://cloud.google.com/kubernetes-engine/docs/tutorials/authenticating-to-cloud-platform)
for detailed instructions on authentication.

```
TEXT_BUCKET={your txt bucket}chinesenotes-text
# First time
gsutil mb gs://$TEXT_BUCKET
gsutil -m rsync -d -r corpus gs://$TEXT_BUCKET
```

To enable the web application to access the storage system, create a service
account with a GCS Storage Object Admin role and download the JSOn credentials
file, as described in [Create service account
credentials](https://cloud.google.com/kubernetes-engine/docs/tutorials/authenticating-to-cloud-platform).
Assuming that you saved the file in the current working directory as 
credentials.json, create a local environment variable for local testing
```
export GOOGLE_APPLICATION_CREDENTIALS=$PWD/credentials.json
```

go get -u cloud.google.com/go/storage

#### Make and Save Go Application Image
The Go app is not needed for chinesenotes.com at the moment but it is use for
other sites (eg. hbreader.org).

Build the Docker image for the Go application:

```
docker build -t cn-app-image .
```

Run it locally with minimal features (C-E dictionary lookp only) enabled
```
docker run -it --rm -p 8080:8080 --name cn-app \
  --mount type=bind,source="$(pwd)",target=/cnotes \
  cn-app-image
```

Test basic lookup with curl
```
curl http://localhost:8080/find/?query=你好
curl http://localhost:8080/findsubstring?query=男&topic=Idiom
```

Run it locally with all features enabled
```
DBUSER=app_user
DBPASSWORD="***"
DATABASE=cse_dict
docker run -itd --rm -p 8080:8080 --name cn-app --link mariadb \
  -e DBHOST=mariadb \
  -e DBUSER=$DBUSER \
  -e DBPASSWORD=$DBPASSWORD \
  -e DATABASE=$DATABASE \
  -e SENDGRID_API_KEY="$SENDGRID_API_KEY" \
  -e GOOGLE_APPLICATION_CREDENTIALS=/cnotes/credentials.json \
  -e TEXT_BUCKET="$TEXT_BUCKET" \
  --mount type=bind,source="$(pwd)",target=/cnotes \
  cn-app-image
```

Debug
```shell
docker exec -it cn-app bash 
```

Test locally by sending a Curl command. Home page

```shell
curl http://localhost:8080
```

Translation memory:

```shell
curl http://localhost:8080/findtm?query=結實
```

Push to Google Container Registry

```shell
docker tag cn-app-image gcr.io/$PROJECT/cn-app-image:$TAG
docker -- push gcr.io/$PROJECT/cn-app-image:$TAG
```

Or use Cloud Build

```shell
BUILD_ID=r162
gcloud builds submit --config cloudbuild.yaml . \
  --substitutions=_IMAGE_TAG="$BUILD_ID"
```

Check that the expected image has been added with the command

```shell
gcloud container images list-tags gcr.io/$PROJECT_ID/nti-image
```

### Web Front End

#### Material Design Web
See the section web-resources/README.md for compiling and testing JavaScript
and CSS files, including ther Material Design resources.

#### HTML File Generation
To generate all HTML files, from the top level project directory
```
export CNREADER_HOME=`pwd`
export DEV_HOME=`pwd`
bin/cnreader.sh
```

For production, copy the files to the storage system.

#### Testing

See e2etest/README.md

## Deploying to Production

### Store HTML Files in Cloud Storage Bucket
This is not stored to a container, rather the web files are uploaded to
Google Cloud Storage. These command will run faster if executed from a build
server in the cloud

```
export BUCKET={your bucket}
# First time
gsutil mb gs://$BUCKET
bin/push.sh
gsutil web set -m index.html -e 404.html gs://$BUCKET
```

### Set up a Cloud SQL Database
New: Replacing management of the Mariadb database in a Kubernetes cluster
Follow instructions in 
[Cloud SQL Quickstart](https://cloud.google.com/sql/docs/mysql/quickstart) using
the Cloud Console.

Connect to the instance from a VM
```
DB_INSTANCE=[your database instance]
gcloud sql connect $DB_INSTANCE --user=root
```

Execute statements in first_time_setup.sql and corpus_index.ddl to define
database and tables as per instructions at
https://github.com/alexamies/chinesenotes-go

Don't forget to switch to the proper database

```sql
use cse_dict
```

Load the proper files into the words table as per data/load_data.sql, the
proper corpus files into the database as per data/load_index.sql, and the
proper index files as per index/load_word_freq.sql.

### Deploy to Cloud run

Deploy the web app to Cloud Run

```shell
PROJECT_ID=[Your project]
IMAGE=gcr.io/${PROJECT_ID}/cn-app-image:${BUILD_ID}
SERVICE=cnreader
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
--allow-unauthenticated \
--add-cloudsql-instances $INSTANCE_CONNECTION_NAME \
--set-env-vars INSTANCE_CONNECTION_NAME="$INSTANCE_CONNECTION_NAME" \
--set-env-vars DBUSER="$DBUSER" \
--set-env-vars DBPASSWORD="$DBPASSWORD" \
--set-env-vars DATABASE="$DATABASE" \
--set-env-vars TEXT_BUCKET="$TEXT_BUCKET" \
--set-env-vars CNREADER_HOME="/" \
--set-env-vars AVG_DOC_LEN="4497"
```

Test it with the command

```shell
curl $URL/find/?query=你好
```

You should see a JSON reply.

### Create and configure the load balancer

Configure a backend bucket

```shell
BACKEND_BUCKET=cnotes-web-bucket-prod
gcloud compute backend-buckets create $BACKEND_BUCKET --gcs-bucket-name $BUCKET
```

Create a serverless Network Endpoint Group for the Cloud Run deployment

```shell
NEG=[name of NEG]
gcloud compute network-endpoint-groups create $NEG \
    --region=$REGION \
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

Create a URL map

```shell
URL_MAP=[your url map name]
gcloud compute url-maps create $URL_MAP \
    --default-backend-bucket $BACKEND_BUCKET
```

Create a matcher for the NEG backend service

```shell
MATCHER_NAME=[your matcher name]
gcloud compute url-maps add-path-matcher $URL_MAP \
    --default-backend-bucket $BACKEND_BUCKET \
    --path-matcher-name $MATCHER_NAME \
    --path-rules="/find/*=$LB_SERVICE,/findadvanced/*=$LB_SERVICE,/findmedia/*=$LB_SERVICE,/findsubstring=$LB_SERVICE,/findtm=$LB_SERVICE"
```

Configure the target proxy

```shell
TARGET_PROXY=cnotes-lb-proxy-prod
gcloud compute target-http-proxies create $TARGET_PROXY \
    --url-map $URL_MAP

STATIC_IP=cnotes-web-prod
gcloud compute addresses create $STATIC_IP --global

FORWARDING_RULE=cnotes-content-rule-prod
gcloud compute forwarding-rules create $FORWARDING_RULE \
    --address $STATIC_IP \
    --global \
    --target-http-proxy $TARGET_PROXY \
    --ports 80,443
```

Setting a named port may take manual editing in the cloud console since the
GKE cluster does not create a named port for the instance group.

Check the name of the forwarding-rule and url-map
```
gcloud compute forwarding-rules list
gcloud compute url-maps list
gcloud compute url-maps describe $URL_MAP
```

To update the load balancer

```shell
gcloud compute url-maps edit $URL_MAP
```

### HTTPS Setup

Create an SSL cert

```shell
SSL_CERTIFICATE_NAME=[your cert name]
DOMAIN=[your domain]
gcloud compute ssl-certificates create $SSL_CERTIFICATE_NAME --domains=$DOMAIN
```

Create a target proxy

```shell
SSL_TARGET_PROXY=[your proxy]
gcloud compute target-https-proxies create $SSL_TARGET_PROXY \
    --url-map=$URL_MAP \
    --ssl-certificates=$SSL_CERTIFICATE_NAME
```

Create a forwarding rule

```shell
SSL_FORWARDING_RULE=[your forwarding rule]
gcloud compute forwarding-rules create $SSL_FORWARDING_RULE \
    --address $STATIC_IP \
    --ports=443 \
    --global \
    --target-https-proxy $SSL_TARGET_PROXY
```
