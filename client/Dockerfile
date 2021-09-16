FROM node:latest as build-stage

WORKDIR /app
COPY . .

RUN yarn global add @vue/cli-service
RUN yarn install
CMD [ "yarn", "serve" ]