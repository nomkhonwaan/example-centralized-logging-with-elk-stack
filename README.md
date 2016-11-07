# Example Centralized Logging with ELK and Filebeat

## Installation
To run this prohect you need to install Docker and make sure your `docker` and `docker-compose` already installed.
Then clone this project via `git` command

```
$ git clone https://github.com/nomkhonwaan/example-centralized-logging-with-elk-stack.git
```

## Getting Started
### Master
On the master server before going next step you need to verify the firewall rules already allowed
on ports 5044 (Filebeat) and 5601 (Kibana). Then go to master folder and type following these commands

```
master$ cd /path/to/example-centralized-logging-with-elk-stack/master
master$ docker-compose up
```

At the first time it will build all images are not exists, wait until it run (up to your network)
after that open browser and goto http://localhost:5601 it will show Kibana 
then goto Management > Saved Objects > Import and choose kibana.json on ./master/kibana this is pre-config of Kibana dashboard and virtualizations objects.
But there are nothing until you start your client.

**Jot down your master IP, this will required for setup Filebest on client**

### Client
On the client you need to change the Filebeat host IP follow you master IP at ./slave/filebeat/filebeat.yml change 
`PUT_YOUR_LOGSTASH_IP_HERE` to your master IP make sure it can be connect by using 

```
client$ telnet YOUR_LOGSTASH_IP 5044
```

You can change target hosts at `./config.yml`, not supported only `GET` method!
then do the same as master

```
client$ cd /path/to/example-centralized-logging-with-elk-stack/slave
client$ docker-compose up
```

if everything worked you will see logs appear on Kibana and dashboard like this

![Kibana Dashboard](https://raw.github.com/nomkhonwaan/example-centralized-logging-with-elk-stack/master/screenshot.png)