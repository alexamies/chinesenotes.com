# End-to-End Testing

First follow the procedure in the main README.md to generate the HTML files for
the site.

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
