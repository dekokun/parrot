# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o parrot

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/parrot /app/
ENTRYPOINT ["./parrot"]
CMD ["-port", "80"]
