rm -rf bid-tracker; 

echo "============================================================================"
echo "=======================Now Building Static Binaries========================="
echo "============================================================================"
CC=/usr/local/bin/musl-gcc CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bid-tracker  -a -ldflags '-extldflags "-static" -s -w'  .

echo "============================================================================"
echo "=======================Now Building Docker Images==========================="
echo "============================================================================"
docker build --network=host --build-arg VCS_REF=`git rev-parse --short HEAD` --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"`  -t ansrivas/bid-tracker:latest -f Dockerfile .
