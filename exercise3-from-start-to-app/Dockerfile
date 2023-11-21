FROM golang:1.21 as builder
############################
# STEP 1 Build golang images
############################

# Create appuser
ENV USER=appuser
ENV UID=10001
# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR $GOPATH/app

ENV GONOSUMDB gitdev.inno.ktb/*,gitdev.devops.krungthai.com/*
ENV GOPROXY https://proxy.golang.org,direct

COPY go.mod go.sum ./
# Fetch dependencies.
# Using go mod with go 1.11
RUN go mod download
RUN go mod verify
COPY . .
# Build the binary
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app ./

############################
# STEP 2 running service in a small image
############################
FROM alpine
ARG SERVICE
# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Add timezone thailand
RUN apk add --no-cache tzdata && \
    apk add --no-cache bash && \
    cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    echo "Asia/Bangkok" >/etc/timezone && \
    apk del tzdata
# Copy our static executable
COPY --from=builder /go/bin/app /go/bin/app
# Use an unprivileged user.
USER appuser:appuser
# Run the binary.
ENTRYPOINT ["/go/bin/app"]
