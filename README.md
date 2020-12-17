UNDERCONSTRUCTION (doesnt works)

---

- [SRS Origin Cluster](https://github.com/ossrs/srs/wiki/v4_EN_OriginCluster) for a Large Number of Streams: SRS origin cluster is designed to serve a large number of streams. (StatefulSet).
``` yaml
    apiVersion: streaming.srs/v1alpha1
    kind: SRSCluster
    metadata:
        name: srscluster-sample
    spec:
    # SRS Origin servers
        size: 3
        serviceName: rtmp

```

- Nginx reads and delivers HLS. (Deployment and HPA).
``` yaml
    apiVersion: streaming.srs/v1alpha1
    kind: SRSClusterNginx
    metadata:
        name: srsnginx-sample
    spec:
        clusterName: srscluster-sample
        # SRS Nginx servers
        minReplicas: 10
        maxReplicas: 25
        metrics:
        - type: Resource
            resource:
            name: memory
            target:
                type: Utilization
                averageUtilization: 80
```