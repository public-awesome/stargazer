FROM alpine:3.13
RUN apk add -U --no-cache ca-certificates


COPY ./build/stargazer /usr/bin/stargazer
EXPOSE 8080

CMD ["stargazer"]
