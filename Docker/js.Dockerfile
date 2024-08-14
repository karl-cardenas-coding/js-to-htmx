# Copyright (c) karl-cardenas-coding
# SPDX-License-Identifier: MIT

FROM node:20 AS builder

ADD . /source
WORKDIR /source

RUN npm ci && npm run build


FROM caddy:latest

COPY --from=builder /source/Caddyfile /etc/caddy/Caddyfile
COPY --from=builder /source/build /usr/share/caddy


EXPOSE 8080

CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile", "--adapter", "caddyfile"]