FROM golang:alpine AS builder

WORKDIR /output
COPY . .

# Generate templates
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -tags timetzdata -o /output/app .

FROM node:23-alpine AS assets

WORKDIR /output
COPY . .

RUN npm install
RUN npx @tailwindcss/cli --minify -i ./assets/inputs/input.css -o ./assets/css/output.css
RUN npx esbuild --bundle --minify --outfile=./assets/js/index.js ./assets/inputs/index.ts

FROM alpine

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=assets /output/assets/js /app/assets/js
COPY --from=assets /output/assets/css /app/assets/css

COPY --from=builder /output/app /app/app

ENV PORT="3000"
ENTRYPOINT ["/bin/sh", "-c", "/app/app --port=$PORT"]
