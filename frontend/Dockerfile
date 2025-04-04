FROM node:22-slim@sha256:bac8ff0b5302b06924a5e288fb4ceecef9c8bb0bb92515985d2efdc3a2447052

WORKDIR /app

RUN corepack enable
RUN corepack prepare pnpm@10.7.0 --activate

COPY package.json pnpm-lock.yaml ./

RUN pnpm install --frozen-lockfile

COPY server.cjs tsconfig.json index.html vite.config.ts ./
COPY ./src ./src
COPY ./public ./public

RUN pnpm build

CMD ["pnpm", "serve"]
