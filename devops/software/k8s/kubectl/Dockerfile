FROM roffe/kubectl:latest

RUN mkdir -p /root/.kube

ADD kubectl /usr/local/bin
COPY .kube /root/.kube
RUN kubectl version

