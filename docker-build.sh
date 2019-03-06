IMAGE="gorm-hello-world"
VERSION="1.0"

echo "Building..."
docker build -t ${IMAGE}:${VERSION} .
docker tag ${IMAGE}:${VERSION} ${IMAGE}:latest

#echo "Running docker container..."
#docker run -p 4531:4531 ${IMAGE}:latest