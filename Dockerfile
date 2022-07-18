FROM golang:1.18.1-bullseye
RUN apt-get update && apt-get -y upgrade
RUN mkdir -p /home/agent
COPY . /home/agent/webapp
WORKDIR /home/agent/webapp
RUN go build -o valorant_agent
CMD ["./agent"]
