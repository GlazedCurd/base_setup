# fluentd/Dockerfile

FROM fluent/fluentd:v1.15.0-debian-1.0
USER root
RUN apt-get update && apt-get install make -y && apt-get install gcc -y
RUN ["fluent-gem", "install", "fluent-plugin-mongo"]
RUN ["fluent-gem", "install", "fluent-plugin-prometheus"]
RUN ["fluent-gem", "install", "fluent-plugin-route"]
USER fluent