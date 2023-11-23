FROM golang:latest

WORKDIR /workspace

COPY elastic elastic

EXPOSE 7777

RUN chmod +x elastic
 
CMD ["./elastic"]

# registry.cnbita.com:5000/wuyiqiang/elastic:v1