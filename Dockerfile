# ---------- base (modules cache) ----------
FROM golang:1.25.3-alpine AS base
WORKDIR /src
RUN apk add --no-cache ca-certificates tzdata
COPY go.mod go.sum ./
RUN go mod download

# ---------- build api (net/http + mux) ----------
FROM base AS build-api
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags "-s -w" -o /out/api ./cmd/main.go

# ---------- build project_fiber ----------
FROM base AS build-project
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags "-s -w" -o /out/project_fiber ./cmd/project_fiber/main.go

# ---------- runtime: api ----------
FROM gcr.io/distroless/static-debian12 AS api-runtime
WORKDIR /app
COPY --from=build-api /out/api /app/app
EXPOSE 22001
USER nonroot:nonroot
ENTRYPOINT ["/app/app"]

# ---------- runtime: project_fiber ----------
FROM gcr.io/distroless/static-debian12 AS project-runtime
WORKDIR /app
COPY --from=build-project /out/project_fiber /app/app
EXPOSE 22002
USER nonroot:nonroot
ENTRYPOINT ["/app/app"]
