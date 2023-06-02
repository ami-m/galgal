#builder: compile phase
FROM golang:stretch AS builder
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src/dropit" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/dropit
COPY .  ./

#install dependencies
RUN go mod vendor

RUN useradd -s /bin/bash -m -d /usr/src/app dropit
USER dropit
########################### phase 2 #####################

FROM builder AS app-builder
# Compile the dropit application each time container is booting
RUN { \
      echo '#!/bin/bash'; \
      echo 'go install app/app.go'; \
      echo '$GOPATH/bin/app'; \
    } | tee /usr/src/app/entrypoint.sh
RUN chmod a+x /usr/src/app/entrypoint.sh
EXPOSE 8000
ENTRYPOINT [ "/usr/src/app/entrypoint.sh" ]

FROM builder AS db-builder
# Compile the dropit db builder each time container is booting
RUN { \
      echo '#!/bin/bash'; \
      echo 'go install databases/database.go'; \
      echo '$GOPATH/bin/database --migrationAction=destroy'; \
      echo '$GOPATH/bin/database --migrationAction=create --shouldRunSeeds=true'; \
    } | tee /usr/src/app/entrypoint.sh
RUN chmod a+x /usr/src/app/entrypoint.sh
ENTRYPOINT [ "/usr/src/app/entrypoint.sh" ]

