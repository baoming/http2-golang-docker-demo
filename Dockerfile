FROM scratch
MAINTAINER Bowen Denning <bdenning@gmail.com>

ADD webservice webservice
ADD localhost.crt localhost.crt
ADD localhost.key localhost.key

EXPOSE 80/tcp
ENTRYPOINT ["/webservice"]
