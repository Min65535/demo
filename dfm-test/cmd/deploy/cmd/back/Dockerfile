FROM debian:latest

COPY publish /usr/bin/

COPY run_in_k8s.sh /

RUN chmod +x /run_in_k8s.sh

ENTRYPOINT ./run_in_k8s.sh