FROM golang:1.17-alpine

RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go mod download
RUN go build -o shipment

EXPOSE 8801

CMD [ "/shipment" ]