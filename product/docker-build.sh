echo '[*] build docker file'
docker build -t gabrielnavas/microservice-product .

# for testing container, not need, uses docker-compose for up container
#echo '[*] create a new container, puuling the image and starting the container'
#docker run -d gabrielnavas/microservice-product --name microservice-product 
