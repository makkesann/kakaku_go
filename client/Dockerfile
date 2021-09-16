FROM node:latest as build-stage

WORKDIR /app
COPY . .

RUN yarn add yarn
RUN yarn install
CMD [ "yarn", "serve" ]