FROM ubuntu:14.04

RUN apt update && apt install -y python3-all python3-pip python3-setuptools ruby-dev build-essential
RUN gem install --no-ri --no-rdoc fpm
RUN pip3 install pip flake8 awscli setuptools --upgrade
