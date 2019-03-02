
## Custom Metric Autoscaling with Stackdriver on GKE

### Create cluster-rolebinding for cluster-admin
    sudo kubectl create clusterrolebinding cluster-admin-binding --clusterrole cluster-admin --user MutluPolatcan@gmail.com

### Deploy Stackdriver Adapter to GKE cluster
	sudo kubectl create -f https://raw.githubusercontent.com/GoogleCloudPlatform/k8s-stackdriver/master/custom-metrics-stackdriver-adapter/deploy/production/adapter.yaml

### Deploy Horizontal Pod Autoscaler (HPA) for API Server and Pubsub Consumer
```yaml
apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: api-server-hpa
spec:
  minReplicas: 3
  maxReplicas: 10
  metrics:
    - external:
        metricName: loadbalancing.googleapis.com|https|request_count
        targetAverageValue: 100
      type: External
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: api-server
```

```yaml
apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: pubsub-consumer-hpa
spec:
  minReplicas: 10
  maxReplicas: 30
  metrics:
    - external:
        metricName: pubsub.googleapis.com|subscription|num_undelivered_messages
        metricSelector:
          matchLabels:
            resource.labels.subscription_id: test-subscription
        targetAverageValue: "50"
      type: External
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: pubsub-consumer
````

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