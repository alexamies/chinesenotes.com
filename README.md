# Chinese Notes Project
## About the project
This is a project for study of modern and literary Chinese language. It includes
a Chinese-English dictionary, a corpus management system written in Go, 
web pages for learning grammar and such, a collection of texts, and a system for
looking up the words in the texts by mouse over and popover by clicking on the
words.

## Acknowldegement
The dictionary is based on the the [CC-CEDICT Chinese - English dictionary]
(http://cc-cedict.org/wiki/), shared under the 
[Creative Commons Attribution-Share Alike 3.0 License]
(http://creativecommons.org/licenses/by-sa/3.0/).

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

### Go command line tool
For more details on the corpus organization and command line tool to process
it see corpus/CORPUS-README.md and go/src/cnreader/README-go.md.

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
 - command line tool for analysis of the corpus

/html
 - raw HTML content minus headers, footers, menus, etc

 /html/templates
  - Go template for generation of HTML files

/web
 - HTML and PHP files. Many but not all files are generated with the Go command line tool. HTML files are written in HTML 5 (See <a href='https://developers.google.com/web/fundamentals/'>Web Fundamentals</a>).

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

# Copy files as needed
tar -zcf cndata.tar.gz data
docker cp cndata.tar.gz mariadb:/cndata.tar.gz

docker exec -it mysql-client bash

# In the container command line
tar -zxf cndata.tar.gz
mv data cndata
cd cndata
mysql --local-infile=1 -h localhost -u root -p

# In the mysql client
# Edit password in the script
# source first_time_setup.sql
source drop.sql
source notes.ddl
source load_data.sql
source corpus_index.ddl
source load_index.ddl

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
  cn-app-image
```

Push to Google Container Registry

```
docker tag cn-app-image gcr.io/$PROJECT/cn-app-image:$TAG
gcloud docker -- push gcr.io/$PROJECT/cn-app-image:$TAG
```

### Experimental Web Front End
The experimental web front end is not used in the chinesenotes.com web site. It
is used for testing the cnweb Go web application and as a prototype for a
future version of chinesenotes.com. It is an Apache web server that serves 
static content and proxies dynamic requests to the Go app via Javascript AJAX
code.

To build the docker image:

```
docker build -f docker/web/Dockerfile -t cn-web-image .

# Test it locally

docker run -itd --rm -p 80:80 --name cn-web --link cn-app  cn-web-image

#Attach to a local image for debugging, if needed
docker exec -it cn-web bash
```

Push to Google Container Registry

```
docker tag hsingyundl-web-image gcr.io/$PROJECT/hsingyundl-web-image:$TAG
gcloud docker -- push gcr.io/$PROJECT/hsingyundl-web-image:$TAG

```

### PHP Web Front End
[Example Project PHP Documentation](https://cloud.google.com/container-engine/docs/tutorials/guestbook)

The PHP web front end is currently used on the chinesenotes.com web site.

Make the web front end PHP Docker image

The web front end is an Apache web server that serves static content and 
uses PHP to serve dynamic requests.

```
docker build -f docker/php/Dockerfile -t chinesenotes-web-image .

# Test it locally
First, export environment variables `DBUSER` and `DBPASSWORD` to connect to the 
database, as per unit tests above.

```
docker run -itd --rm -p 80:80 --name chinesenotes-web --link mariadb -e DBUSER=$DBUSER -e DBPASSWORD=$DBPASSWORD chinesenotes-web-image
```

# Attach to a local image for debugging, if needed
docker exec -it chinesenotes-web bash
```

Push to Google Container Registry
[Google Container Registry Quickstart](https://cloud.google.com/container-registry/docs/quickstart)

```
TAG=prototype19
docker tag chinesenotes-web-image gcr.io/$PROJECT/chinesenotes-web-image:$TAG
gcloud docker -- push gcr.io/$PROJECT/chinesenotes-web-image:$TAG

```

### Set Up Kubernetes Cluster and Deployment
[Container Engine Quickstart](https://cloud.google.com/container-engine/docs/quickstart)
The digital library runs in a Kubernetes cluster using Google Container Engine.
The Maria DB runs on a persistent volume. The password for the root user and 
application user for the database are stored in Kubernetes secrets.

```
gcloud container clusters create $CLUSTER --zone=$ZONE --disk-size=500 --machine-type=n1-standard-1 --num-nodes=1 --enable-cloud-monitoring

gcloud compute disks create --size 200GB mariadb-disk

kubectl create secret generic mysqlroot --from-literal=MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD
kubectl create secret generic mysql --from-literal=DBPASSWORD=$DBPASSWORD

kubectl create --save-config -f kubernetes/db-deployment.yaml
kubectl create --save-config -f kubernetes/db-service.yaml
```

The application tier is dependent on an operational database, which takes
some manual configuration to achieve. 

```
kubectl get pods
kubectl exec -it {POD_NAME} bash
```

Execute the database configuration steps above and continue to configure the
app and web tiers.

```
# Deploy the app tier
kubectl create --save-config -f kubernetes/app-deployment.yaml 
kubectl create --save-config -f kubernetes/app-service.yaml

# Deploy the web tier
kubectl create --save-config -f kubernetes/web-deployment.yaml
kubectl expose deployment hsingyundl-web --target-port=80  --type=NodePort

# Check that the service is available
kubectl get service hsingyundl-web

# Configure public ingress
kubectl create --save-config -f kubernetes/web-ingress.yaml
```

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