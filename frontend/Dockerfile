FROM node:17.7.2 as node
WORKDIR /app
COPY package*.json yarn.lock ./
RUN yarn install
COPY . .
EXPOSE 3000
ENTRYPOINT [ "npm", "start" ]
