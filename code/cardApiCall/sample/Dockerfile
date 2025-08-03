FROM golang:1.23.5 AS build
WORKDIR /cardApiCall
# Copy dependencies list
COPY go.mod go.sum ./
# Build with optional lambda.norpc tag
COPY main.go .
RUN go build -tags lambda.norpc -o main main.go
# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /cardApiCall/main ./main
ENTRYPOINT [ "./main" ]