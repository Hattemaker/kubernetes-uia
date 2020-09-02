# Base imaget vi bruker for denne containeren.
FROM alpine:latest

# Oppdaterer pakker, og legge til sertifikater og til slutt sletter cache fra package manager.
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Oppretter en gruppe og en bruker som er medlem av gruppen.
RUN addgroup -S uia && adduser -S student -G uia

# Oppretter mappe hvor applikasjonen skal ligge.
RUN mkdir -p /api

# Setter hvilken mappe containeren skal jobbe i.
WORKDIR /api

# Kopierer vår applikasjon til /api
COPY greeter .

# Forteller containeren at vår applikasjon skal kunne nås på port 8080. 
EXPOSE 8080

# Alle kommandoer som blir kjørt fra nå av kjøres som brukeren appuser
USER student

# Forteller containeren hva som skal startes/kjøres.
ENTRYPOINT [ "./greeter" ]
