FROM golang:latest 
RUN mkdir /app
RUN dep ensure 
ADD ./webapp /app/
ADD ./vendor /app/ 
WORKDIR /app
EXPOSE 8080
ENV DATABASE onboarding
ENV DB_USER root
ENV DB_PASSWORD root
ENV DB_HOST unix
RUN go build webapp-crud.go 
CMD ["/app/webbapp-crud"]