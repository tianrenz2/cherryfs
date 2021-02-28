# Cherryfs: A Distributed Object Storage

- [Cherryfs: A Distributed Object Storage](#cherryfs-a-distributed-object-storage)
  - [Introduction](#introduction)
  - [Prerequisite](#prerequisite)
  - [Components](#components)
    - [Meta Service](#meta-service)
    - [Chunk Service](#chunk-service)
  - [Tutorial](#tutorial)
    - [Chunk Service Setup](#chunk-service-setup)
    - [Meta Service Setup](#meta-service-setup)
    - [Client](#client)

## Introduction

Cherryfs is a file system, or also called object storage, it is designed to distributed environment to keep files safe and tolerate bad situations such as server failure. It implements some important features such as GET&PUT for file via gRPC, failover, data redundancy via replication, fault-tolerance.

## Prerequisite

Cherryfs uses [etcd](https://github.com/etcd-io/etcd) as the metadata storage in order to keep the cluster metadata safe. So before running the meta service, etcd cluster needs to
be deployed in the first place (maybe using goreman), for the instruction of doing this, checkout its official [documentation](https://etcd.io/docs/v3.1.12/dev-guide/local_cluster/).

## Components

There are two primary services of the cherryfs, meta service and chunk service.

### Meta Service

It is responsible for managing the meta data, such as host space info and object locations, of the whole cluster.

### Chunk Service

It is the local data plane, which is responsible for managing local file operations and gets requests from its peers.

## Tutorial

As it is stated above, there are two services that needs to be launched to form the whole file system in the cluster.

### Chunk Service Setup

Chunk service has to be run on all the nodes in order to provide the storage capability, also chunk service has to be started before the meta service.

There are some editable environment variables for the chunk service:

``` bash
# since the whole system is built upon etcd, the running etcd service addresses need 
# to be specified.
ETCDADDR=127.0.0.1:12380,127.0.0.1:22380,127.0.0.1:32380 

# each host has its own id, there needs to be a local path to store the host's id
HOST_ID_PATH=/etc/cherryfs/hostid

# this is the path of chunk's configuration file, a sample file is located in config/chunk1.json
CHUNK_CONFIG=/etc/chunk1.json

# sample launch command
ETCDADDR=127.0.0.1:12380,127.0.0.1:22380,127.0.0.1:32380 HOST_ID_PATH=/etc/cherryfs/hostid CHUNK_CONFIG=/etc/cherryfs/chunk1.json ./cherryfs-chunk
```

### Meta Service Setup

Meta service is like the conductor of the whole file system. This service is not required on all the nodes of the cluster, instead, at any moment, there is only one node that is taking this responsibility, so the number of nodes that needs to be running the meta serice should be at least 1. Keep in mind that if you run meta service only on 1 node, there would be no guarentee of the fault tolerance, the failure of this node could cause to the loss of metadata and thus the whole system will be unavailable.

Another important thing is that meta service depends on the etcd, be sure that meta service is running on the node which has been deployed in the etcd cluster.

There are some editable environment variables for the meta service:

``` bash
# since the whole system is built upon etcd, the running etcd service addresses need 
# to be specified.
ETCDADDR=127.0.0.1:12380,127.0.0.1:22380,127.0.0.1:32380 

# this is the path of chunk's configuration file, a sample file is located in config/cluster_config.json
META_CONFIG=/etc/cluster_config.json

# sample launch command
ETCDADDR=127.0.0.1:12380,127.0.0.1:22380,127.0.0.1:32380 META_CONFIG=/etc/cherryfs/cluster_config.json ./cherryfs-meta
```

### Client

To be updated...
