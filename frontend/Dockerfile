FROM node:22

WORKDIR /app

COPY package*.json .

RUN npm i

COPY . .

ENTRYPOINT ["npm", "run", "dev", "--", "--host"]
