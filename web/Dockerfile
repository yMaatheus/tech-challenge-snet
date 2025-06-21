FROM node:20-alpine AS builder

WORKDIR /app

COPY . .

RUN npm install

RUN npm run build

FROM node:20-alpine

WORKDIR /app

COPY --from=builder /app/.output .output
COPY --from=builder /app/node_modules node_modules
COPY --from=builder /app/package.json .
COPY --from=builder /app/nuxt.config.ts .
COPY --from=builder /app/.env .env

ENV NUXT_PORT=3000
EXPOSE 3000

CMD ["npx", "nuxi", "preview"]
