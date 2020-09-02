# Greeter

Demo av Docker

## Quickstart

```sh
# Bygg applikasjonen først.
# Vi setter CGO_ENABLED til 0 for å lage en statisk binærfil (ellers ville den hatt referanser til go modules).
# Vi setter GOOS til linux for å forsikre oss om at applikasjonen er kompatibel med Docker base imaget.
CGO_ENABLED=0 GOOS=linux go build

# Bygg Docker Image
docker build -t uia-app .

# Evt bygg til en egen tag
# docker build -t uia:1 .

# Start Docker containeren
# Binder lokal port til container port. <lokal>:<container>
docker run -p 8080:8080 uia-app
```

## Kjekke kommandoer

```sh

# List images du har lokalt
docker images

# List kjørende containere
docker ps

# Kjør kommandoer i containeren
docker exec <container id> sh
```
