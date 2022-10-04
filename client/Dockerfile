FROM node:18-alpine3.16 as installer

    WORKDIR /app

    COPY package.json .

    RUN npm install

FROM node:18-alpine3.16 as builder

    WORKDIR /app

    COPY --from=installer /app/node_modules /app/node_modules

    COPY . .

    COPY src src

    RUN npm run build

FROM nginxinc/nginx-unprivileged:stable-alpine-slim

    COPY --from=builder /app/dist /usr/share/nginx/html
    COPY nginx.conf /etc/nginx/conf.d/default.conf
