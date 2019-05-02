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

There are also some Python ulitities for processing text.

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
sudo apt-get update
sudo apt-get install -y git

cd $HOME/chinesenotes.com
-- Substitute for your own location and user name
export CN_HOME=/disk1
sudo mkdir $CN_HOME/chinesenotes.com
sudo chown alex:alex $CN_HOME/chinesenotes.com
git clone git://github.com/alexamies/chinesenotes.com $CN_HOME/chinesenotes.com

cd $CN_HOME/chinesenotes.com
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

/colab
- Colab notebooks for exploring the library text data

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

 /html/templates
  - Go templates for generation of HTML files. This is a default that can be 
    overridden by setting an environment variale named TEMPLATE_HOME with the
    path relative to the project home.

/html/material-templates
  - Go templates for generation of HTML files using material design lite
    styles.

/index
 - Output from corpus analysis

/kubernetes
 - For production deployment

/python
 - Jupyter notebooks and Python utilities for additional of new terms to the
   dictionary.

/web-resources
 - Static resources, including CSS, JavaScript, image, and sound files

/web-staging
 - Generate HTML files. Many but not all files are generated with the Go command 
  line tool cnreader. HTML files are written in HTML 5 (See <a 
  href='https://developers.google.com/web/fundamentals/'>Web Fundamentals</a>).
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
[Mariadb Documentation](https://mariadb.org/)

The local development environment uses a Mariadb database. 

### Mariadb Docker Image
See the documentation at [Mariadb Image 
Documentation](https://hub.docker.com/_/mariadb/) and [Installing and using 
MariaDB via Docker](https://mariadb.com/kb/en/library/installing-and-using-mariadb-via-docker/).

To start a Docker container with Mariadb and connect to it from a MySQL command
line client execute the command below. First, set environment variable 
`MYSQL_ROOT_PASSWORD`. Also, create a directory outside the container to use as a
permanent Docker volume for the database files. In addition, mount volumes for
the tabe separated data to be loaded into Mariadb. See 
[Manage data in Docker](https://docs.docker.com/storage/) and 
[Use volumes](https://docs.docker.com/storage/volumes/) for details on volume
and mount management with Docker.

```
MYSQL_ROOT_PASSWORD=[your password]
mkdir mariadb-data
docker run --name mariadb -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD -d \
  -v "$(pwd)"/mariadb-data:/var/lib/mysql \
  --mount type=bind,source="$(pwd)"/data,target=/cndata \
  --mount type=bind,source="$(pwd)"/index,target=/cnindex \
  mariadb:10.3
```

The data in the database is persistent even if the container is deleted. To
restart the database use the command

```
docker restart  mariadb
```

Compile the library document files and tiles into a tab separated file for
loading into the database with the Python program
```
bin/doc_list.sh
```

To load data from other sources connect to the database container
or start up a mysql-client
```
docker exec -it mariadb bash
```

In the container command line
```
cd cndata
mysql --local-infile=1 -h localhost -u root -p
```

In the mysql client
Edit password in the script
```
# source first_time_setup.sql
# source drop.sql
source delete_index.sql
source notes.ddl
source load_data.sql
source corpus_index.ddl
source load_index.sql
source library/digital_library.sql
quit
```
The word frequency index files are large and may take a long time to load.
Exit MySQL and change to the index directory for loading the word and bigram
frequency files
```
cd ../index
mysql --local-infile=1 -h localhost -u root -p
source delete_word_freq.sql
source load_word_freq.sql
```

### Go Applications
[Go Documentation](https://golang.org)

The indexing command tool and new generation of the web application are written
in Go. 

##### Install Go
[Go Install Documentation](https://golang.org/doc/install)

```
wget https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go*

vi .profile
```

Add 

```
export PATH=$PATH:/usr/local/go/bin

source .profile
```

Build Software locally

```
cd chinesenotes.com
cd go
source path.bash.inc
cd src/cnweb
go build
```

For Go unit tests set in a local development environment first export environment
variables

```
export DBHOST=localhost
export DBUSER={database user}
export DBPASSWORD={the password}
cd go/src/cnweb/identity
go test
```

Run locally:
```
export DBHOST=localhost
export DBUSER=app_user
export DBPASSWORD="***"
export DATABASE=cse_dict
./cnweb
```

In another window  send a HTTP request:
```
curl http://localhost:8080/find/?query=hello
```

### Email configuration (Optional)
Optional, used for password recovery in translation portal.

Email is sent with [SendGrid](https://sendgrid.com). 
[SendGrid Go Client Library Documentation](https://github.com/sendgrid/sendgrid-go)

Follow the following steps to get it set up
echo "export SENDGRID_API_KEY='YOUR_API_KEY'" > sendgrid.env
echo "sendgrid.env" >> .gitignore
source ./sendgrid.env
go get github.com/sendgrid/sendgrid-go

### Text Files for Full Text Search
Copy the text files to an object store. If you prefer not to use GCS, then you
can use the local file system on the application server. The instructions here
are for GCS. See [Authenticating to Cloud Platform with Service
Accounts](https://cloud.google.com/kubernetes-engine/docs/tutorials/authenticating-to-cloud-platform)
for detailed instructions on authentication.

```
TEXT_BUCKET={your txt bucket}
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
docker build -f docker/go/Dockerfile -t cn-app-image .
```

Run it locally
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
```
docker exec -it cn-app bash 
```

Push to Google Container Registry

```
docker tag cn-app-image gcr.io/$PROJECT/cn-app-image:$TAG
gcloud docker -- push gcr.io/$PROJECT/cn-app-image:$TAG
```

### Web Front End
To generate all HTML files, from the top level project directory
```
export CNREADER_HOME=`pwd`
export DEV_HOME=`pwd`
bin/cnreader.sh
```

For production, copy the files to the storage system.

For development, use the web Docker container. It for developming versions of 
the chinesenotes.com web front end. It is an Apache web server that serves 
static content and proxies dynamic requests to the Go app via Javascript AJAX
code.

To build the docker image:

```
docker build -f docker/web/Dockerfile -t cn-web-image .
```
Test it locally
```
WEB_DIR=web-staging
docker run -itd --rm -p 80:80 --name cn-web --link cn-app \
  --mount type=bind,source="$(pwd)"/$WEB_DIR,target=/usr/local/apache2/htdocs \
  cn-web-image
```

To attach to a local image for debugging, if needed:

```
docker exec -it cn-web bash
```
Set the load balancer up after creating the Kubernetes cluster

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
cd data
INSTANCE=cnotes
gcloud sql connect $INSTANCE --user=root
```

Execute statements in first_time_setup.sql and corpus_index.ddl to define
database and tables.

Consolidate the document titles into a single file
```
bin/doc_list.sh 
```

Import the data for table word_freq_doc via the Cloud Console using the import
function. The other tables can be imported using the MySQL client (much faster):
```
INSTANCE=cnotes
gcloud sql connect $INSTANCE --user=root
#source data/notes.ddl
#source data/corpus_index.ddl
#source data/drop.sql
#source data/delete_index.sql
source data/load_data.sql
source data/load_index.sql
source index/load_word_freq.sql
#source data/library/digital_library.sql
```

### Set Up Kubernetes Cluster and Deployment
[Container Engine Quickstart](https://cloud.google.com/container-engine/docs/quickstart)
The dynamic part of the app run in a Kubernetes cluster using Google
Kubernetes Engine. To create the cluster and authenticate to it:
```
gcloud container clusters create $CLUSTER \
  --zone=$ZONE \
  --disk-size=500 \
  --machine-type=n1-standard-1 \
  --num-nodes=1 \
  --enable-cloud-monitoring
gcloud container clusters get-credentials $CLUSTER --zone=$ZONE
```

Configure access to Cloud SQL using instructions in
[Connecting from Kubernetes Engine](https://cloud.google.com/sql/docs/mysql/connect-kubernetes-engine).
Save the JSON key file. Create the proxy user:

```
PROXY_PASSWORD=[Your value]
gcloud sql users create proxyuser cloudsqlproxy~% --instance=$INSTANCE \
  --password=$PROXY_PASSWORD
```

Get the instance connection name:
```
gcloud sql instances describe $INSTANCE
```

Create secrets
```
PROXY_KEY_FILE_PATH=[JSON file]
kubectl create secret generic cloudsql-instance-credentials \
    --from-file=credentials.json=$PROXY_KEY_FILE_PATH
kubectl create secret generic cloudsql-db-credentials \
    --from-literal=username=proxyuser --from-literal=password=$PROXY_PASSWORD
```

Deploy the app tier
```
kubectl apply -f kubernetes/app-deployment.yaml 
kubectl apply -f kubernetes/app-service.yaml
```

Test from the command line
```
kubectl get pods
POD_NAME=[your pod name]
kubectl exec -it $POD_NAME bash
apt-get update
apt-get install curl
curl http://localhost:8080/find/?query=hello
```

The load balancer connects to the Kubernetes NodePort with a managed instance
group named port. To get the list of named ports use the command
```
gcloud compute instance-groups list
MIG=[your managed instance group]
gcloud compute instance-groups managed get-named-ports $MIG
```

To add a new named port use the command
```
PORTNAME=hsingyunport
PORT=30080
gcloud compute instance-groups managed set-named-ports $MIG \
  --named-ports="$PORTNAME:$PORT" \
  --zone=$ZONE
```
Be careful with this command that you do not accidentally clear the already
existing named ports that other apps in the cluster may be depending on.

### Create and configure the load balancer

Add a firewall rule to allow the load balancer to reach the managed instance
group.
```
FW_RULE_NAME=cnotes-allow-lb
gcloud compute firewall-rules create $FW_RULE_NAME \
    --allow tcp:$PORT \
    --source-ranges 130.211.0.0/22,35.191.0.0/16
```

Add a health check
```
HEALTH_CHECK_NAME=cnotes-app-check
gcloud compute health-checks create http $HEALTH_CHECK_NAME --port=$PORT \
     --request-path=/healthcheck/
```

Configure the load balancer backend service
```

BACKEND_NAME=cnotes-backend
gcloud compute backend-services create $BACKEND_NAME \
     --protocol HTTP \
     --health-checks $HEALTH_CHECK_NAME \
     --global
gcloud compute backend-services add-backend $BACKEND_NAME \
    --balancing-mode UTILIZATION \
    --max-utilization 0.8 \
    --capacity-scaler 1 \
    --instance-group $MIG \
    --instance-group-zone $ZONE \
    --global
```

Configure the backend bucket

```
BACKEND_BUCKET=cnotes-web-bucket
gcloud compute backend-buckets create $BACKEND_BUCKET --gcs-bucket-name $BUCKET
```

Configure the load balancer
```
URL_MAP=[your url-map]
gcloud compute url-maps create $URL_MAP \
    --default-backend-bucket $BACKEND_BUCKET
MATCHER_NAME=cnotes-url-matcher
gcloud compute url-maps add-path-matcher $URL_MAP \
    --default-backend-bucket $BACKEND_BUCKET \
    --path-matcher-name $MATCHER_NAME \
    --path-rules="/find/*=$BACKEND_NAME,/findadvanced/*=$BACKEND_NAME,/findmedia/*=$BACKEND_NAME"

TARGET_PROXY=cnotes-lb-proxy
gcloud compute target-http-proxies create $TARGET_PROXY \
    --url-map $URL_MAP

STATIC_IP=cnotes-web
gcloud compute addresses create $STATIC_IP --global

FORWARDING_RULE=cnotes-content-rule
gcloud compute forwarding-rules create $FORWARDING_RULE \
    --address $STATIC_IP \
    --global \
    --target-http-proxy $TARGET_PROXY \
    --ports 80

```

Setting a named port may take manual editing in the cloud console since the
GKE cluster does not create a named port for the instance group.

Check the name of the forwarding-rule and url-map
```
gcloud compute forwarding-rules list
gcloud compute url-maps list
gcloud compute url-maps describe $URL_MAP
```

### Service Account Access to Text Files for Full Text Search
Save the credentials.json file created above to a Kubernetes secret with the
command

```
kubectl create secret generic cnotes-app-key \
  --from-file=key.json=credentials.json
```

This should match the ```GOOGLE_APPLICATION_CREDENTIALS``` environment variable
and also the ```volumes``` and ```volumeMounts``` entries in the 
app-deployment.yaml file.

### Troubleshooting
SSH to another VM and try sending a HTTP request via curl to the internal IP
of the VM hosting the GKE cluster:
```
INTERNAL_IP={The IP}
curl http://$INTERNAL_IP:30080/find/?query=hello
```

Check that the instance group has a named port, so that the load balancer can
send traffic to it, and that the port name in the LB configuration matches
the instance group port name.

### Update App in Kubernetes Cluster

To update an existing deployment using the GCP Container Registry

```
# Make edits
kubectl apply -f kubernetes/app-deployment.yaml
kubectl apply -f kubernetes/app-service.yaml
```