# elasticTraining
在k8s集群上，使用HPA进行pytorch弹性训练，基于GPU的数量进行增加和减少任务pod


curl 10.0.102.46:46178/metrics

apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name:  prometheus-elastic-sm 
  namespace: monitoring 
spec:
  endpoints:
  - interval: 5s
    port: http
    path: /metrics
  namespaceSelector:
    matchNames:
    - default
  selector:
    matchLabels:
      app: prometheus-elastic