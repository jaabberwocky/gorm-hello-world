IMAGE="gorm-hello-world"
VERSION="1.0"
GOOS="linux"

echo "Building Go binaries..."
env GOOS=linux go build main.go

echo "Building..."
docker build -t ${IMAGE}:${VERSION} .
docker tag ${IMAGE}:${VERSION} ${IMAGE}:latest

#echo "Running docker container..."
#docker run -p 4531:4531 ${IMAGE}:latest