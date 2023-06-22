# Especificamos la imagen de GO para la APP
FROM golang:1.20
# Especificamos el WORKDIR
WORKDIR "/app/SateliteMostrador"
# Copiamos todo el proyecto
COPY . .
# Copiamos los packetes para correr la APP
RUN go get -u github.com/gin-gonic/gin
# Compilamos el binario exe para la APP
RUN go build -o main .
# Corremos el exe
CMD ["./main"]