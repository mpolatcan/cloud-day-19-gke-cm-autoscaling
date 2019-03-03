
## Custom Metric Autoscaling with Stackdriver on GKE

### Create cluster-rolebinding for cluster-admin
    sudo kubectl create clusterrolebinding cluster-admin-binding --clusterrole cluster-admin --user $(gcloud config get-value account)

### Deploy Stackdriver Adapter to GKE cluster
	sudo kubectl create -f https://raw.githubusercontent.com/GoogleCloudPlatform/k8s-stackdriver/master/custom-metrics-stackdriver-adapter/deploy/production/adapter.yaml

### Create a secret to store your GCP service account file in secret
    sudo kubectl create secret generic gke-auth-file-secret --from-file=[YOUR_AUTH.json]

### Deploy API Server and PubSub Consumer
    sudo kubectl apply -f api_server.yaml
    sudo kubectl apply -f consumer.yaml
    
### Deploy Horizontal Pod Autoscaler (HPA) for API Server and Pubsub Consumer
    sudo kubectl apply -f api_server_hpa.yaml
    sudo kubectl apply -f consumer_hpa.yaml
    
### Running test with your favorite load test tool.

In there, I have used LoadImpact's k6 tool that executes Javascript test scenarios.
    
#### stress.js
```javascript
import http from "k6/http";
import { sleep } from "k6";

export default function() {
    http.get("http://[INGRESS_EXTERNAL_IP]/stress");
};
```

You can learn Ingress external ip by executing below command:

    kubectl get ing 

Start test:

    k6 run --vus 100 --duration 3m stress.js


