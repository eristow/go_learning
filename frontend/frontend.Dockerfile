FROM node:23.9 AS build

WORKDIR /staging

COPY package.json pnpm-lock.yaml ./
RUN npm install -g pnpm
RUN pnpm install --frozen-lockfile

COPY . .

RUN pnpm lint

RUN pnpm check

RUN pnpm build

RUN pnpm prune --production


FROM node:23.9 AS frontend

ENV NODE_ENV=production

WORKDIR /app
COPY --from=build /staging/.env ./.env
COPY --from=build /staging/build ./build
COPY --from=build /staging/package.json ./package.json
COPY --from=build /staging/node_modules ./node_modules

ENTRYPOINT ["node", "--env-file=.env", "build"]
