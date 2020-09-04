# Base imaget vi bruker for denne containeren.
FROM alpine:latest

# Oppdaterer pakker, og legge til sertifikater og til slutt sletter cache fra package manager.
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Oppretter en gruppe og en bruker som er medlem av gruppen.
RUN addgroup -S uia && adduser -S student -G uia

# Oppretter mappe hvor applikasjonen skal ligge.
RUN mkdir -p /home/student/app

# Setter hvilken mappe containeren skal jobbe i.
WORKDIR /home/student/app

# Kopierer vår applikasjon til vår WORKDIR som er /home/student/app
COPY greeter .

# Mulighet for å sende med --build-arg til docker build.
ARG version=1

# Setter miljøvariabel, bruker ARG her
ENV VERSION=${version}

# Forteller containeren at vår applikasjon skal kunne nås på port 8080. 
EXPOSE 8080

# Alle kommandoer som blir kjørt fra nå av kjøres som brukeren appuser
USER student

# Forteller containeren hva som skal startes/kjøres
ENTRYPOINT [ "./greeter" ] 
