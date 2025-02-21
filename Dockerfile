FROM golang:1.23

WORKDIR /app

# Instalacja Air
RUN go install github.com/air-verse/air@latest

# Kopiowanie plików go.mod i go.sum
COPY go.mod go.sum ./

# Pobieranie zależności
RUN go mod download

# Kopiowanie kodu źródłowego
COPY . .

# Utworzenie katalogu na przesyłane pliki
RUN mkdir -p uploads

# Exposowanie portu
EXPOSE 8080

# Uruchomienie aplikacji z Air
CMD ["air"]