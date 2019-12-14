# End-to-End Testing

The end-to-end testing covers testing a selected set of web pages with
JavaScript and a backend that is not connected to a Database. The web app
provided allows serving static media, including HTML pages, from the same
container that AJAX calls are made to. This is simplified from the full
production setup with a load balancer fronting static media in GCS and
AJAX calls from the web application.

First follow the procedure in the main README.md to generate the HTML files for
the site.

## Run Locally

To run locally, first copy important static files to the static directory

```shell
./e2ecopy.sh
```

Build the Go web app

```shell
go build
```


## Deploy to Cloud Run
Deploy on [Cloud Run](https://cloud.google.com/run/) and use the pages in the
test folder to drive end to end tests, as described here. This is intended for
testing the JavaScript client with real web pages and mock data.

```shell
cd e2etest
```

Copy important files to the static directory

```shell
./e2ecopy.sh
```

Build the Docker image for the test and deploy to Cloud Run using
Cloud Build

```shell
gcloud builds submit --config cloudbuild.yaml .
```

Go to Cloud Run management page in the GCP Cloud Console to find the URL
of the server.