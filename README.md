# Cherryfs: A Distributed Object Storage

- [Cherryfs: A Distributed Object Storage](#cherryfs-a-distributed-object-storage)
  - [Introduction](#introduction)
  - [Prerequisite](#prerequisite)

## Introduction

Cherryfs is a file system, or also called object storage, it is designed to distributed environment to keep files safe and tolerate bad situations such as server failure. It implements some important features such as GET&PUT for file via gRPC, failover, data redundancy via replication, fault-tolerance.

## Prerequisite

Cherryfs uses [etcd](https://github.com/etcd-io/etcd) as the metadata storage in order to keep the cluster metadata safe. So before running the meta service, etcd cluster needs to
be deployed in the first place (maybe using goreman), for the instruction of doing this, checkout its official [documentation](https://etcd.io/docs/v3.1.12/dev-guide/local_cluster/). 
