docker run -it --link skilltree-mongo:mongo --rm mongo sh -c 'exec mongo "$MONGO_PORT_27017_TCP_ADDR:$MONGO_PORT_27017_TCP_PORT/test"'
docker run --name skilltree-mongo -d -p 27017:27017 mongo