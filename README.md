# NETDOCTOR

## usage

### create config.json
```
netctl init
```
netctl while load config.json as default config

#### config.json demo
```
{
 "namespace": "kosmos-system",
 "version": "0.2.1",
 "protocol": "tcp",
 "podWaitTime": 30,
 "port": "8889",
 "maxNum": 3,
 "cmdTimeout": 10,
 "srcKubeConfig": "~/.kube/config",
 "srcImageRepository": "ghcr.io/kosmos-io"
}
```

### check
```
netctl check
```

### resume 

```
netctl resume
```
Detect nodes that failed last time