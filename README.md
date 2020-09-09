# Docker og Kubernetes presentasjon/workshop for UiA

Innholdet i dette repo:
- apps: i denne mappen ligger applikasjonene vi skal jobbe med i denne presentasjonen
- kubernetes: manifest filer for å deploye applikasjon til kubernetes
- src/docs/asciidocs: her ligger slidene som asciidoc filer

Det er fint om man før man kommer i workshop/forelesning kan installere følgende vertøy:
 * https://www.docker.com/products/docker-desktop
   * fint om man har prøvd å starte opp kubernetes fra docker desktop
   * på linux installer docker, kubernetes-cli og microk8s eller en annen lokal kubernetes distribusjon
 * https://github.com/digitalocean/doctl
   * Hvis man vil gjøre handson demo her er det fint om man har en konto på digital ocean også

Hvis du har noen spørsmål før, under eller etter workshop og foretrekker å stille dem skriftlig så lag en issue/sak i dette github repo
så skal vi svare så godt vi kan.

## Hvordan bygge slidene til en  presentasjon

`make doc`

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
 
 Verktøy for å jobbe med kubernetes som kan være fine å vite om er:
  - https://github.com/lensapp/lens
    - anbefalt GUI 
  - https://github.com/derailed/k9s
    - anbefalt terminal/curses gui
  - https://github.com/kubernetes-sigs/krew
    - installere utvidelser til kubectl
  - https://github.com/wercker/stern
    - taile logge
  - https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/
    - denne er litt knotete å sette opp 
    - må trikse litt for å få auth til å funke på Docker Desktop https://forums.docker.com/t/docker-for-mac-kubernetes-dashboard/44116/9
  
#### Kjekke kubectl kommandoer
Kubectl kommandoen har veldig god hjelp, så lurer du på noe ville jeg startet med å se der. 

 - `alias k=kubectl` hvis man er lat og ikke vil skrive kubectl hele tiden 
 - `k explain pod` for å forklare hvordan en ressurs ser ut
 - `k explain pod.status` for å forklare hvordan en underressurs ser ut
 - `k api-resources` for å liste alle ressurser man har
