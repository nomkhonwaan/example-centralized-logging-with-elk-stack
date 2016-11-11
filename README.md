# Centralized Logging System wth ELK Stack 

## Table of Contents
- [Installation](#installation)
- [Getting Started](#getting-started)
  - [Master](#master)
  - [Client](#client)

## Installation
To install this project you need to install Docker and make sure `docker` and `docker-compose` already installed,
then using `git` to clone this project 

```
$ git clone https://github.com/nomkhonwaan/example-centralized-logging-with-elk-stack.git
```

then following these steps on master and client 

## Getting Statred
Normally master and client should stay in the same network but sending logs across public network is fine due to your log info.
BTW I recommend to secure with SSL/TLS while using public network. 

### Master 
On the master server should whitelist port 5044 for sending logs via Filebeat and port 5601 as public for Kibana.
Then go to the master folder and type following commands

```
master$ cd /path/to/example-centralized-logging-with-elk-stack/master
master$ docker-compose up 
```

Once it done it should have 3 services running in the container, BTW you can't see anything in the Kibana
because the client didn't setup. So move to the client and running it.

### Client 
On the client you need to update `filebeat.yml` on `YOUR_LOGSTASH_IP` to your Logstash IP 
it should be master IP in the default.

```
filebeat.prospectors:
- input_type: log
  paths:
    - /opt/stresser/logs/*.log
output.logstash:
  hosts: ["YOUR_LOGSTASH_IP:5044"]
```

and same as the master type following these commands to bring it up

```
client$ cd /path/to/example-centralized-logging-with-elk-stack/slave
client$ docker-compose up
```

if everything work fine you will see 2 servics are running. Then go back to Kibana
and setup an "Index Patterns" with the "Time-field name" you need to choose "received-date" and then
go to Management > Saved Objects > Import and choose `kibana.json` file inside `./master/kibana/kibana.json` 
this is pre-config for Kibana dashboard and virtualization objects.

![Centralized Logging Dashboard](https://raw.github.com/nomkhonwaan/example-centralized-logging-with-elk-stack/master/screenshot.png)

## FYI 
To run the client alongside the master you need to add `networks` config like this
after `services` section in the `docker-compose.yml` file

```
networks:
  default:
    external:
      name: master_default
```

this will let client services attach to the master network (it should named master_default by default) 
and then you need to config `filebeat.yml` in the `YOUR_LOGSTASH_IP` to use `logstash` name instead.
The `logstash` name is an local domain name on Docker Compose system, declared in the `docker-compose.yml` file at `container_name` field.

> For long run testing please change the request targets to your own server instead of http://httpbin.org 
