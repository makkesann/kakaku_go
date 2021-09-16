FROM node:latest

WORKDIR /app
COPY . .

# RUN apk update
# RUN yarn add yarn
# RUN yarn add @vue/cli
RUN yarn install
CMD [ "yarn", "serve" ]