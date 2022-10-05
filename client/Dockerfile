FROM node:16-alpine3.15 as builder

FROM builder

WORKDIR /app
COPY ./client/* .

RUN npm install

ENTRYPOINT npm run dev -- --host --port ${CLIENT_PORT}