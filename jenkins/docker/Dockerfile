FROM jenkins:latest

# Temp fix for creating volumes
USER root
RUN chown -R 1000:1000 /var/jenkins_home

RUN apt-get update
RUN apt-get -y install python-software-properties
RUN apt-get -y install software-properties-common

# Install Gradle
RUN add-apt-repository ppa:cwchien/gradle
# RUN ls /etc/apt/sources.list.d/cwchien-gradle-jessie.list
RUN sed -i 's/jessie/trusty/g' /etc/apt/sources.list.d/cwchien-gradle-jessie.list
RUN apt-get update
RUN apt-get -y install gradle
