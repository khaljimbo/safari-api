# Get first golang image
FROM golang:alpine as builder
# Allow Microscaler to get meta details from the container  
ARG VCS_REF
LABEL org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/khaljimbo/safari-api"
COPY api /src/api
WORKDIR /src/api
# Install git and get external dependancies
RUN apk add git coreutils \
    && GO111MODULE=off go get github.com/gorilla/mux
RUN go build -o /src/api
FROM alpine
WORKDIR /bin
COPY --from=builder /src/api/ ./
ENV PORT=8000 \
    LOG_FORMAT=text
EXPOSE 8000
CMD ["api"]