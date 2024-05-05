FROM golang:1.20.2-alpine3.16 as build
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go

FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /app/bootstrap /bootstrap
ENTRYPOINT [ "/bootstrap" ]
