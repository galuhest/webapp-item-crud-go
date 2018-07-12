FROM golang:latest 
RUN mkdir /apps
ADD . /apps/
WORKDIR /apps
EXPOSE 8080
ENV DATABASE onboarding
ENV DB_USER root
ENV DB_PASSWORD root
ENV DB_HOST unix
RUN go get -u -insecure github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build app/webapp/main.go 
CMD ["/app/webapp/main"]