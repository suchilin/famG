FROM node:10.17-alpine
WORKDIR /app
COPY package.json yarn.lock /app/
RUN yarn install
COPY . /app
