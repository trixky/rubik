FROM node:16-alpine3.15

WORKDIR /app
COPY ./client/* .
RUN rm -rf /app/node_modules

RUN npm install

ENTRYPOINT npm install && npm run dev -- --host --port ${CLIENT_PORT}