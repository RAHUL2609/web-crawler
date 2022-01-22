FROM alpine:latest

RUN mkdir -p demo
RUN chmod -R 777 demo

ADD main demo/

COPY urls.json demo/

WORKDIR demo
CMD ["./main"]