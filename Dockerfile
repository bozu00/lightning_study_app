FROM golang:1.10

# ENV GOPATH /go
# ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
# RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN go get -u github.com/go-sql-driver/mysql 
RUN go get -u github.com/jmoiron/sqlx                  
RUN go get -u github.com/labstack/echo                 
RUN go get -u gopkg.in/gorp.v1                         
RUN go get -u bitbucket.org/liamstask/goose/cmd/goose  
RUN go get -u github.com/gorilla/sessions              
RUN go get -u github.com/labstack/echo-contrib 
RUN go get -u github.com/ipfans/echo-session           
RUN go get -u github.com/dgrijalva/jwt-go              
RUN go get -u github.com/boj/redistore
RUN go get -u github.com/google/uuid
RUN go get -u github.com/alecthomas/participle
RUN go get -u cloud.google.com/go/storage
RUN go get -u golang.org/x/net/context
RUN go get -u github.com/BurntSushi/toml
RUN go get -u github.com/githubnemo/CompileDaemon
RUN go get -u github.com/Shaked/gomobiledetect 
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/jessevdk/go-assets-builder

COPY google_application_credentials.json google_application_credentials.json
ENV GOOGLE_APPLICATION_CREDENTIALS google_application_credentials.json

WORKDIR $GOPATH
