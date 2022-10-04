FROM golang:latest as build
WORKDIR /app
COPY . .
RUN go build src/main.go

FROM gcr.io/distroless/cc
WORKDIR /app
COPY --from=build /app/main .
CMD [ "/app/main" ]
EXPOSE 8000
