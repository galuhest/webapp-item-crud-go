FROM golang:latest 
RUN mkdir /app
ADD . /app/
WORKDIR /app
EXPOSE 8080
ENV DATABASE onboarding
ENV DB_USER root
ENV DB_PASSWORD root
ENV DB_HOST unix
RUN ls vendor/github.com/galuhest
RUN go build app/webapp/main.go 
CMD ["/app/main"]