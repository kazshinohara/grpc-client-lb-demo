# gRPC Client-side load balancing demo

A demo application shows gRPC Client-side load balancing with k8s headless services in DNS Round Robin manner.

## How to use

### 1. Preparation

Set your preferred Google Cloud region name.
```shell
export REGION_NAME={{REGION_NAME}}
```

Set your Google Cloud Project ID
```shell
export PROJECT_ID={{PROJECT_ID}}
```

Set your Artifact Registry repository name
```shell
export REPO_NAME={{REPO_NAME}}
```

Enable Google Cloud APIs
```shell
gcloud services enable \
  artifactregistry.googleapis.com \
  cloudbuild.googleapis.com 
```

### 2. build container images
Note: please make your own [Artifact Registry repo](https://cloud.google.com/artifact-registry/docs/docker/quickstart) in advance, if you don't have it yet.

#### Build the client image
```shell
git clone git@github.com:kazshinohara/grpc-client-lb-demo.git
```
```shell
cd grpc-client-lb-demo/hostinfo-client
```
```shell
gcloud builds submit --tag ${REGION_NAME}-docker.pkg.dev/${PROJECT_ID}/${REPO_NAME}/hostinfo-client:v1
```

#### Build the server image
```shell
cd ../hostifno-server
```
```shell
gcloud builds submit --tag ${REGION_NAME}-docker.pkg.dev/${PROJECT_ID}/${REPO_NAME}/hostinfo-server:v1
```

### 3. Replace the image paths with yours in k8s manifests
#### Update client.yaml
```shell
cd ../hostinfo-client
```
```shell
vim client.yaml
```

#### Update server.yaml
```shell
cd ../hostinfo-server
```
```shell
vim server.yaml
```

### 4. Deploy containers to your GKE cluster
#### Deploy server (must be done at first)
```shell
kubectl apply -f server.yaml
```

#### Deploy client
```shell
cd ../hostinfo-client
```
```shell
kubeclt apply -f client.yaml
```

### 5. Check the behavior
#### Confirm the client pod name
```shell
kubectl get pods | grep hostinfo-client
```

#### Check logs
```shell
kubectl logs {{client_pod_name}}
```

You will see the requests (Unary gRPC) from client to server have been load-balanced like below.
The server has 3 replicas and each request went to one of them in Round Robin manner.
```shell
kubectl logs hostinfo-client-5685c4946b-wcbmq -f
2021/06/23 07:56:40 hostinfo-server-6589b974dd-4828c
2021/06/23 07:56:41 hostinfo-server-6589b974dd-zlz6x
2021/06/23 07:56:42 hostinfo-server-6589b974dd-5cs69
2021/06/23 07:56:43 hostinfo-server-6589b974dd-4828c
2021/06/23 07:56:44 hostinfo-server-6589b974dd-zlz6x
2021/06/23 07:56:45 hostinfo-server-6589b974dd-5cs69
2021/06/23 07:56:46 hostinfo-server-6589b974dd-4828c
2021/06/23 07:56:47 hostinfo-server-6589b974dd-zlz6x
2021/06/23 07:56:48 hostinfo-server-6589b974dd-5cs69
2021/06/23 07:56:49 hostinfo-server-6589b974dd-4828c
2021/06/23 07:56:50 hostinfo-server-6589b974dd-zlz6x
2021/06/23 07:56:51 hostinfo-server-6589b974dd-5cs69
2021/06/23 07:56:52 hostinfo-server-6589b974dd-4828c
2021/06/23 07:56:53 hostinfo-server-6589b974dd-zlz6x
2021/06/23 07:56:54 hostinfo-server-6589b974dd-5cs69
2021/06/23 07:56:55 hostinfo-server-6589b974dd-4828c
2021/06/23 07:56:56 hostinfo-server-6589b974dd-zlz6x
2021/06/23 07:56:57 hostinfo-server-6589b974dd-5cs69
2021/06/23 07:56:58 hostinfo-server-6589b974dd-4828c
2021/06/23 07:56:59 hostinfo-server-6589b974dd-zlz6x
2021/06/23 07:57:00 hostinfo-server-6589b974dd-5cs69
2021/06/23 07:57:01 hostinfo-server-6589b974dd-4828c
2021/06/23 07:57:02 hostinfo-server-6589b974dd-zlz6x
2021/06/23 07:57:03 hostinfo-server-6589b974dd-5cs69
```
