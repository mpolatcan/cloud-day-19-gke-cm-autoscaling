### Install Helm and Tiller

```
wget https://storage.googleapis.com/kubernetes-helm/helm-v2.11.0-linux-amd64.tar.gz
tar -zxvf helm-v2.11.0-linux-amd64.tar.gz
sudo mv linux-amd64/helm /usr/local/bin/helm
kubectl create clusterrolebinding user-admin-binding --clusterrole=cluster-admin --user=$(gcloud config get-value account)
kubectl create serviceaccount tiller --namespace kube-system
kubectl create clusterrolebinding tiller-admin-binding --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
helm init --service-account=tiller
helm update
```

### Install Prometheus

```
kubectl create namespace monitoring
helm repo add coreos https://s3-eu-west-1.amazonaws.com/coreos-charts/stable/
helm install coreos/prometheus-operator --name prometheus-operator --namespace monitoring
helm install coreos/kube-prometheus --name kube-prometheus --namespace monitoring
```

### Install Prometheus Adapter

```
helm install stable/prometheus-adapter --name prometheus-adapter --set prometheus.url=http://kube-prometheus.monitoring.svc --namespace monitoring
```

### Deploy Sample Application

```
kubectl create ns test
kubectl apply -f sample_app.yaml
```

## Create Service Monitor

```
kubectl apply -f service_monitor.yaml
```

### Create Horizontal Pod Autoscaler

```
kubectl apply -f sample_app_hpa.yaml
```

### Running stress test with your favorite stress tool. In there, I have used LoadImpact's k6 tool 

    k6 run --vus 100 --duration 3m stress.js

#### stress.js
```javascript
import http from "k6/http";
import { sleep } from "k6";

export default function() {
    http.get("http://34.95.102.51/stress");
};
```

### References
 
[1] - https://github.com/DirectXMan12/k8s-prometheus-adapter

[2] - https://github.com/kubeless/kubeless/tree/master/manifests/autoscaling