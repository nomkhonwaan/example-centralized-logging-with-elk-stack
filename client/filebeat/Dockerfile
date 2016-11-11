FROM ubuntu:latest

# Installing Filebeat
RUN apt-get update && \
    apt-get install -y curl && \
    curl -L -O https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-5.0.0-amd64.deb && \
    dpkg -i filebeat-5.0.0-amd64.deb

ENTRYPOINT ["/usr/share/filebeat/bin/filebeat"]
CMD ["-e" ""]