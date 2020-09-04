FROM golang:1.15-alpine3.12 as builder

WORKDIR /repeater

COPY *.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o repeater .


FROM alpine:3.12

# Oppdaterer pakker, og legge til sertifikater og til slutt sletter cache fra package manager.
RUN apk update && apk add curl ca-certificates && rm -rf /var/cache/apk/*

# Oppretter en gruppe og en bruker som er medlem av gruppen.
RUN addgroup -S uia && adduser -S student -G uia

# Oppretter mappe hvor applikasjonen skal ligge.
RUN mkdir -p /home/student/app

# Setter hvilken mappe containeren skal jobbe i.
WORKDIR /home/student/app

COPY --from=builder /repeater/repeater .

# Alle kommandoer som blir kjørt fra nå av kjøres som brukeren appuser
USER student

ENTRYPOINT [ "./repeater" ]