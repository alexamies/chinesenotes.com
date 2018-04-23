# Chinese Notes Project
## About the project
This is a project for study of modern and literary Chinese language. It includes
a Chinese-English dictionary, a corpus management system written in Go, 
web pages for learning grammar and such, a collection of texts, and a system for
looking up the words in the texts by mouse over and popover by clicking on the
words.

The software and dictionary power three web sites with different corpora:

1. http://chinesenotes.com - for literary Chinese documents and a small amount of modern
   Chinese
2. http://ntireader.org - for the Taisho version of the Chinese Buddhist Canon
3. http://hbreader.org - for Venerable Master Hsing Yun's collected writings

## Acknowldegements
The dictionary includes many words from the 

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
Installation instructions are for Debian LAMP. The web site will work on any
environment that runs PHP.

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

### Database Setup
```
sudo apt-get -y install mysql-server mysql-client
```

Follow instructions in dictionary-readme.txt to set up the database

### Web Server Setup
```
$ sudo apt-get install -y apache2 php5 php5-mysql

Set the Apache home directory to the web directory for the project

$ sudo vi /etc/apache2/sites-enabled/000-default
```

$ sudo apachectl restart

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

/web
 - HTML and PHP files. Many but not all files are generated with the Go command 
  line tool cnreader. HTML files are written in HTML 5 (See <a 
  href='https://developers.google.com/web/fundamentals/'>Web Fundamentals</a>).
  This is a default that can be overridden by setting an environment variale 
  named WEB_DIR with the path relative to the project home.

 /web/script
  - JavaScript files

 /web/analysis
  - Corpus analysis files (generated)

 /web/images
 - static images

/web/inc
 - PHP includes

 /web/erya, /web/laoshe, etc
  - corpus files (generated)

/web-staging
 - The bin/cnreadher.sh script places the generated files in this directory.

## Containerization
Containerization is new, under development, and not yet deployed to prod.

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

The application uses a Mariadb database. 

### Mariadb Docker Image
[Mariadb Image Documentation](https://hub.docker.com/r/library/mariadb/)

To start a Docker container with Mariadb and connect to it from a MySQL command
line client execute the command below. First, set environment variable 
`MYSQL_ROOT_PASSWORD`.

```
docker run --name mariadb -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD -d \
  --mount type=bind,source="$(pwd)"/data,target=/cndata \
  mariadb:10.3
docker exec -it mariadb bash
mysql --local-infile=1 -h localhost -u root -p
```

The data in the database is persistent unless the container is deleted. To
restart the database use the command

```
docker restart  mariadb
```

To load data from other sources connect to the database container
or start up a mysql-client

```
docker build -f docker/client/Dockerfile -t mysql-client-image .
docker run -itd --rm --name mysql-client --link mariadb \
  --mount type=bind,source="$(pwd)"/data,target=/cndata \
  mysql-client-image

docker exec -it mariadb bash

# In the container command line
cd cndata
mysql --local-infile=1 -h localhost -u root -p

# In the mysql client
# Edit password in the script
# source first_time_setup.sql
source drop.sql
source notes.ddl
source load_data.sql
source corpus_index.ddl
source load_index.sql
source library/digital_library.sql

```

### Go Application
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
export DBDBHOST=localhost
export DBUSER={database user}
export DBPASSWORD={the password}
cd go/src/cnweb/identity
go test
```

### Email configuration
Email is sent with [SendGrid](https://sendgrid.com). 
[SendGrid Go Client Library Documentation](https://github.com/sendgrid/sendgrid-go)

Follow the following steps to get it set up
echo "export SENDGRID_API_KEY='YOUR_API_KEY'" > sendgrid.env
echo "sendgrid.env" >> .gitignore
source ./sendgrid.env
go get github.com/sendgrid/sendgrid-go

#### Make and Save Go Application Image
The Go app is not needed for chinesenotes.com at the moment but it is use for
other sites (eg. hbreader.org).

Build the Docker image for the Go application:

```
docker build -f docker/go/Dockerfile -t cn-app-image .
```

Run it locally
```
export DBDBHOST=mariadb
export DBUSER=app_user
export DBPASSWORD="***"
export DATABASE=cse_dict
docker run -itd --rm -p 8080:8080 --name cn-app --link mariadb \
  -e DBDBHOST=$DBDBHOST \
  -e DBUSER=$DBUSER \
  -e DBPASSWORD=$DBPASSWORD \
  -e DATABASE=$DATABASE \
  -e SENDGRID_API_KEY="$SENDGRID_API_KEY" \
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

# Test it locally

docker run -itd --rm -p 80:80 --name cn-web --link cn-app \
  --mount type=bind,source="$(pwd)"/$WEB_DIR,target=/usr/local/apache2/htdocs \
  cn-web-image

#Attach to a local image for debugging, if needed
docker exec -it cn-web bash
```

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

Set the load balancer up after creating the Kubernetes cluster

### Set Up Kubernetes Cluster and Deployment
[Container Engine Quickstart](https://cloud.google.com/container-engine/docs/quickstart)
The dynamic part of the app run in a Kubernetes cluster using Google Kubernetes Engine.
The Maria DB runs on a persistent volume. The password for the root user and 
application user for the database are stored in Kubernetes secrets.

```
gcloud config set project $PROJECT
gcloud config set compute/zone $ZONE
gcloud container clusters create $CLUSTER \
  --zone=$ZONE \
  --disk-size=500 \
  --machine-type=n1-standard-1 \
  --num-nodes=1 \
  --enable-cloud-monitoring

gcloud compute disks create --size 200GB cnotesdb-disk

gcloud container clusters get-credentials $CLUSTER --zone=$ZONE

kubectl create secret generic mysqlroot --from-literal=MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD
kubectl create secret generic mysql --from-literal=DBPASSWORD=$DBPASSWORD

kubectl create --save-config -f kubernetes/db-deployment.yaml
kubectl create --save-config -f kubernetes/db-service.yaml
```

The application tier is dependent on an operational database, which takes
some manual configuration to achieve. 

```
kubectl get pods
POD_NAME=
tar -zcf  cndata.tar.gz data
kubectl cp cndata.tar.gz $POD_NAME:.
kubectl exec -it $POD_NAME bash
rm -rf data
rm -rf cndata/*
tar -zxf cndata.tar.gz
#mkdir cndata
mv data/* cndata/.
cd cndata
mysql --local-infile=1 -h localhost -u root -p
source notes.ddl
source corpus_index.ddl
source load_data.sql
source load_index.sql
source library/digital_library.sql
```

Execute the database configuration steps above and continue to configure the
app and web tiers.

```
# Deploy the app tier
kubectl apply -f kubernetes/app-deployment.yaml 
kubectl apply -f kubernetes/app-service.yaml

# Configure public ingress
kubectl apply -f kubernetes/web-ingress.yaml
```

Find the name of the forwarding-rule and url-map created with the ingress object
```
gcloud compute instance-groups list
gcloud compute forwarding-rules list
gcloud compute url-maps list
gcloud compute url-maps describe URL_MAP
```

Create and configure the load balancer
```
# Allow the load balancer to reach the VM
gcloud compute firewall-rules create cnotes-app-rule \
    --allow tcp:30080 \
    --source-ranges 130.211.0.0/22,35.191.0.0/16
gcloud compute health-checks create http cnotes-app-check --port=30080 \
     --request-path=/healthcheck/
gcloud compute backend-services create cnotes-service \
     --protocol HTTP \
     --health-checks cnotes-app-check \
     --global
INSTANCE_GROUP=gke-cnotes-cluster-default-pool-42ceda90-grp
gcloud compute backend-services add-backend cnotes-service \
    --balancing-mode UTILIZATION \
    --max-utilization 0.8 \
    --capacity-scaler 1 \
    --instance-group $INSTANCE_GROUP \
    --instance-group-zone $ZONE \
    --global
gcloud compute backend-buckets create cnotes-web-bucket --gcs-bucket-name $BUCKET
gcloud compute url-maps create cnotes-map \
    --default-backend-bucket cnotes-web-bucket
gcloud compute url-maps add-path-matcher cnotes-map \
    --default-backend-bucket cnotes-web-bucket \
    --path-matcher-name cnotes-matcher \
    --path-rules="/find/*=cnotes-service,/findmedia/*=cnotes-service"
gcloud compute target-http-proxies create cnotes-lb-proxy \
    --url-map cnotes-map
gcloud compute addresses create cnotes-web --global
gcloud compute forwarding-rules create cnotes-content-rule \
    --address cnotes-web \
    --global \
    --target-http-proxy cnotes-lb-proxy \
    --ports 80
```

Setting a named port may take manual editing in the cloud console since the
GKE cluster does not create a named port for the instance group.

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
TAG=prototype19
kubectl set image deployment/hsingyundl-web hsingyundl-web=gcr.io/$PROJECT/hsingyundl-web-image:$TAG
kubectl set image deployment/hsingyundl-app hsingyundl-app=gcr.io/$PROJECT/hsingyundl-app-image:$TAG

# If configuration changes outside image are needed:
kubectl get deployment hsingyundl-web -o yaml > kubernetes/web-deployment.yaml
kubectl get deployment hsingyundl-app -o yaml > kubernetes/app-deployment.yaml

# Make edits
kubectl apply -f kubernetes/db-deployment.yaml
kubectl apply -f kubernetes/db-service.yaml
kubectl apply -f kubernetes/app-deployment.yaml
kubectl apply -f kubernetes/app-service.yaml
kubectl apply -f kubernetes/web-deployment.yaml
kubectl apply -f kubernetes/web-ingress.yaml
```