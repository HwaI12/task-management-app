FROM node:17.7.2
WORKDIR /app
COPY package*.json ./
RUN yarn install
COPY . .
EXPOSE 3000
ENTRYPOINT [ "npm", "start" ]
