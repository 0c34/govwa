FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /app
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go version
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /app/main .

# Build a small image
FROM scratch

COPY --from=builder /dist/main /
COPY ./config/config.json /config/config.json
COPY ./templates/* /templates/
COPY ./public/. /public/
EXPOSE 8888
# Command to run
CMD ["./main"]

FROM ubuntu:bionic-20180426

ENV DEBIAN_FRONTEND="noninteractive"

#Install dependencies
RUN apt-get update && apt-get upgrade -y

#Curl Vulnerability https://www.cvedetails.com/cve/CVE-2018-1000300/
RUN apt-get install -y curl && \
#GIT Vulnerability CVE https://www.cvedetails.com/cve/CVE-2018-17456/
    apt-get install -y git && \
#OpenSSH Vulnerability https://www.cvedetails.com/cve/CVE-2018-15473/
    apt-get install -y openssh-server && \
#Installation of ftp server
    apt-get install -y proftpd

COPY ./userfiles/shadow /etc/shadow
COPY ./userfiles/passwd /etc/passwd
RUN chmod o-rwx /etc/shadow
RUN chmod o-rwx /etc/passwd
COPY ./user-data-ftp/ /home/
COPY ./sshd_config /etc/ssh/sshd_config
RUN service ssh start
CMD ["proftpd", "--nodaemon"]
