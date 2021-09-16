FROM node:latest as build-stage

WORKDIR /app
COPY . .

# RUN yarn add @vue/cli
RUN yarn install
# CMD [ "yarn", "serve" ]