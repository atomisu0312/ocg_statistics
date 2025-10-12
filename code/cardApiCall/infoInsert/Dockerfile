FROM golang:1.23.5 AS build
ENV GOOS=linux
ENV GOARCH=arm64
WORKDIR /cardApiCall
# Copy dependencies list
COPY go.mod go.sum ./
RUN go mod download
# Build with optional lambda.norpc tag
COPY . .
RUN go build -tags lambda.norpc -o main main.go
# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /cardApiCall/main ./main
ENTRYPOINT [ "./main" ]