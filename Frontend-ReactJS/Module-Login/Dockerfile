FROM golang:1.19-alpine

WORKDIR /app

RUN apk update && apk upgrade

RUN apk add --update nodejs npm

RUN npm install -g pnpm

COPY ./ ./

ENV CHOKIDAR_USEPOLLING=true

CMD pnpm run play

EXPOSE 3000