#### End-to-End Testing

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

Build the Docker image for the test and push it to the registry using
Cloud Build

```shell
gcloud builds submit --config cloudbuild.yaml .
```

Deploy to Cloud Run

```shell
gcloud run deploy --image gcr.io/$PROJECT/e2etest --platform managed
```
