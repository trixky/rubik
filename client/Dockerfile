FROM node:16-alpine3.15

# copy client files
WORKDIR /app
COPY ./client .

# remove local node modules
RUN rm -rf /app/node_modules

# install client dependencies
RUN npm install

ENTRYPOINT npm install && npm run dev -- --host --port ${CLIENT_PORT}