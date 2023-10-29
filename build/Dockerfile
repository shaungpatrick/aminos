# docker build -t amino -f build/Dockerfile .
FROM ubuntu

WORKDIR /app

COPY app/aminos.json /app
COPY build/aminos /app

EXPOSE 3333

CMD /app/aminos