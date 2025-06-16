
# Usa un'immagine base per la build multi-stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copia i file go.mod e go.sum per scaricare le dipendenze
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copia il resto del codice sorgente
COPY . .

# Compila l'applicazione Go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Fase di runtime
FROM alpine:latest

WORKDIR /root/

# Copia l'eseguibile compilato dalla fase di build
COPY --from=builder /app/main .

# Espone la porta su cui il tuo backend ascolta (es. 8080)
EXPOSE 8080

# Comando per avviare l'applicazione
CMD ["./main"]
