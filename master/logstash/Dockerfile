FROM logstash:latest

# Installing logstash-input-beats plugin
RUN /usr/share/logstash/bin/logstash-plugin install logstash-input-beats

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["-e", ""]