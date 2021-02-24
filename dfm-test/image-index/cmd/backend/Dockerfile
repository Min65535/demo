FROM xl.image:5000/debian:stable

COPY image-index /usr/bin/

COPY ./static /usr/bin/static

COPY run_in_k8s.sh /

RUN chmod +x /run_in_k8s.sh

ENTRYPOINT ./run_in_k8s.sh