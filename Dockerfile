FROM golang:1.19 as builder
WORKDIR /app
ADD . /app/
RUN go build -o /amidtoken main.go


# FROM alpine:3.17
# COPY --from=builder /amidtoken /
EXPOSE 1212
CMD [ "/amidtoken" ]