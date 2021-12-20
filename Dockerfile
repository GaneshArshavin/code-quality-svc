From golang:1.14
RUN apt-get update || exit 0
RUN apt-get upgrade -y
RUN apt-get install vim sudo dbus curl bc supervisor -y

RUN mkdir -p /go/src/github.com/CodeQualityProject
WORKDIR /go/src/github.com/CodeQualityProject
COPY . /go/src/github.com/CodeQualityProject/
ENV GOBIN=/go/bin/

RUN go install findlanguage.go