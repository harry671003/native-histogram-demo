## Native Histogram Demo

A simple demo of native histograms in Prometheus.

### Building
`make build`

### Running
`make run`


## Demo App
After running just visit `http://localhost:8081/metrics` to see the metrics emitted. 
Bonus tip, use `Accept: application/openmetrics-text` for native histogram metrics.

## Prometheus
Prometheus runs in localhost:9091

### PromQL queries

Use this query for classic histogram.

```
histogram_quantile(0.9, sum by (le) (request_latency_bucket))
```

Use this query for native histogram.

```
histogram_quantile(0.9, native_request_latency)
```
