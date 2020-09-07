# Docker og Kubernetes presentasjon/workshop for UiA

- apps: i denne mappen ligger applikasjonene vi skal jobbe med i denne presentasjonen
- kubernetes: manifest filer for å deploye applikasjon til kubernetes
- src/docs/asciidocs: her ligger slidene som asciidoc filer

## Hvordan bygge slidene til en revlewjs presentasjon

`./gradlew`

Slidene blir da lagt i `build/docs` mappen


## Ressurser

### Docker

I denne presentasjonen viser vi kun hvordan man bygger docker images med Dockefile på den grunnlegende måten. Det finnes høyere
nivå abstraksjoner som det kan være verd å se på:

 - https://buildpacks.io/
 - https://github.com/containers/buildah
 - https://spring.io/blog/2020/01/27/creating-docker-images-with-spring-boot-2-3-0-m1
 - https://github.com/GoogleContainerTools/jib
 - https://github.com/GoogleContainerTools/kaniko
 - https://github.com/Skatteetaten/architect 
### Kubernetes


 - [Intoduksjon til kubernetes](https://www.digitalocean.com/community/tutorials/an-introduction-to-kubernetes)
 - [Kuberetes dokumentasjon](https://kubernetes.io/docs/home/)
  
Det finnes mange forskjellige teknikker og verktøy for å konfigurere kubernetes. Under er en liten, men ikke utfyllende, liste med relevant materiale for de som er interessert. 

 - https://blog.argoproj.io/the-state-of-kubernetes-configuration-management-d8b06c1205
 - https://www.giantswarm.io/blog/application-configuration-management-in-kubernetes/   
 - https://oam.dev/
 - https://www.weave.works/technologies/gitops/
 - https://www.pulumi.com/
 - https://skatteetaten.github.io/aurora/ 