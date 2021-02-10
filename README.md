# Ego Server

A simple HTTP echo server written in Go and inspired by the Google Cloud Run [Build and deploy a container](https://cloud.google.com/run/docs/quickstarts/build-and-deploy) quickstart.

* The PORT environment variable sets the server port (default is 8080)
* Returns the request headers and some additional information

## Google Cloud Run

[Cloud Run](https://cloud.google.com/run/) is Google Cloud's fully serverless compute offering that lets you run stateless HTTP-driven containers, without worrying about the infrastructure.

### Build

`gcloud builds submit --tag gcr.io/PROJECT-ID/egoserver` where **_PROJECT-ID_** is your GCP project ID.

### Deploy

`gcloud run deploy --image gcr.io/PROJECT-ID/egoserver --platform managed` where **_PROJECT-ID_** is your GCP project ID.

During deployment:
- You will be prompted for the service name: press Enter to accept the default name, `egoserver`
- You will be prompted for region: select the region of your choice, for example `europe-north1` (Finland)
- You will be prompted to allow unauthenticated invocations: respond `y`

### See it in action

On deployment (see above) success, the command line displays the service URL. Visit your deployed container by opening the service URL in a web browser or by using the `curl` command.

## 127.0.0.1

### Build
`$ docker build -t egoserver .`

### Run
`$ docker run -d -p 8080:8080 egoserver`

### See it in action
`$ curl localhost:8080` or head to http://localhost:8080 from your favorite web browser.