FROM python:3.8.6
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY allure /usr/local
RUN mkdir -p /usr/local/java
ADD jdk.tar.gz /usr/local/java
#ENV 设置环境变量
ENV PATH /usr/local/allure/bin:$PATH
ENV JAVA_HOME /usr/local/java/jdk
ENV JRE_HOME ${JAVA_HOME}/jre
ENV CLASSPATH .:${JAVA_HOME}/lib:${JRE_HOME}/lib
ENV PATH ${JAVA_HOME}/bin:$PATH
RUN allure --version