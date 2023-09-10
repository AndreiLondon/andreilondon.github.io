FROM golang:1.19
LABEL vesion="1.0"
LABEL maintaner="Aliaksei Vidaseu and Andrei Martynenko"
LABEL description="Project Forum"
# "Project is a forum. We can register,log in, save post, 
# make comment on post, like or dislike posts and comments, filter posts"
LABEL port="8080"
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /docker-forum
CMD [ "/docker-forum" ]
#Build Image
# docker build --tag forum .
#Run image
#docker run forum
#Publish
#docker run --publish 8080:8080 forum
#Publish in detached mode
#docker run -d -p 8080:8080 forum
#List of containes
#docker ps -all
#Remove Container
#docker rm 150bfea2f440
#Remove Image
#docker rmi ascii-web
#Export container
#docker export suspect-container > suspect-container.tar
#Show Lables
#docker image inspect forum | jq '.[0].ContainerConfig.Labels'
#Build image 
#docker image build -f Dockerfile -t <name_of_the_image> .
## docker image build -f Dockerfile -t forum-image .
#Run container
#docker container run -p <port_you_what_to_run> --detach --name <name_of_the_container> <name_of_the_image>
##  docker container run -p 8080:8080 --detach --name forum-container forum-image  
# file system
#docker exec -it <container_name> /bin/bash
## docker exec -it forum-container /bin/sh