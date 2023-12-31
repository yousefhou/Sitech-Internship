# docker:
is responsible for creating and managing objects such as imeges , containers ,volumes, and network on a host.

*the rest API provieds an interface to manage objects in Docker.
 
 # the docker CLI:
 is the command line interface that we use to run command to mange objects in Docker.
 
# libcontainer :
is written in go the same programming language that docker is written in and reduce docker dependency on the kernal LXC technology with libcontainer docker could now directly interact with the linux kernal features such as namespaces and cgroups (the default execution environment of docker)

 # docker image : 
 is a read only template used for creating a docker container (could be base images of OS or apps like webservers or databses )	
 # docker container : 
 running instance of an image (read write template ) ceate start stop move delete a container using Docker CLI or the api .
 # docker network :
 help in implementing various solutions for containers to be able to communicate with each other 
 when the container dies all the data died .
 # docker volumes :
 help in persisting data across container restarts 
 image registry :stores docker images and is used publish and share images publicly or withna private organization , docker hub is a public registry that anyone can use , docker is config to look for images on docker hub by defalut 
 
 # unix socket :
 is an ipc or inter-process communication mechanism that used between different processes on the same host *Docker deamon is only accessible within the same host and the docker CLi .
 # the docker deamon :
 is not accessible outside of the docker host because its only listening on th unix socket  so we cane make it to listen on a TCP interface on the docker host by adding the host option while running the docker deamin command line  in this case our host happens to have an interface .
 
 *the command line utility used to manage docker objests is called docker itself . 
 
 # how to run docker contianer command :
 docker container run -d --name=blabla ubuntu

# some questions :
Q: Which is the default data directory for Docker?
>/var/lib/docker

Q:What component of the docker engine manages the images, containers, volumes and networks on a host?
>docker deamon

Q:By default, data stored inside the container is always persistent
>false

Q:What is the command to view the version of docker engine installed?
>docker version

Q:What are the components of the Docker Engine?
>Docker CLI , Docker Deamon , REST API

Q:What component of the docker architecture is responsible for managing containers on Linux on version 1.15 of Docker Engine? 
>libcontainer

Q:What are the 2 specifications from OCI?
>here are currently two specifications in development and in use: Runtime Specification (runtime-spec) and the Image Specification (image-spec)

Q:By default, Docker is configured to look for images on Google Cloud Registry
>false

Q:We can run containers without installing Docker?
>true

Q:What are the primary objects that Docker engine manages?
>images , container ,Volumes , networks .

Q:Which component is responsible for keeping the containers alive when the Docker Daemon goes down?
>containered-Shim

Q:Which component is a read-only template used for creating a Docker container?
>docker image

Q:What does OCI stand for?
>open container initiative 


Q:Each container gets a CPU share of …. assigned by default.
>1024

Q:What will happen if the --memory-swap is set to -1?
>the container is allowed to use unlimited swap 

Q:What is a linux feature that prevents a process within the container to access raw sockets?
>kernel capabilities.

Q:by default a container runs with unlimited CPU and memory resource
>True

Q:By default, all containers get the same share of CPU cycles. How to modify the shares?
>docker container run --cpu-shares=512nginx


# when you install Docker it creates 3 networks automaticly (Bridge , None , Host ):
*Bridge:  is the default network a container gets attached to and its a private internal network created by Docker  on the Host  All containers attached to this network by default and they get an internal IP address usually in the range 172.17 series . the container can access each other using this intrenal IP if required to access any of these containers from the outside world map the ports of these containers to ports on the docker Host 
another way to access to these containers is to associate the container to the host network this takes out any network isolation between the docker host and the docker container 

*None : the container are not attachedto any network and dosent have any access to the external network or other containers they run in an isolated network 

*containers can reach each other using their names 

# when you want to connect  for example to SQL use the Docker container name 

how does Docker implement networking? and what the thecnlogy behind it? how are the container isolated within the host ? Docker uses network namespaces that creates separate namespace for each container it then uses virtual ethernet pairs to connect containers together. 

# to list the default availbale network we run : 
> docker network ls 

# how do we see the network settings and IP address assigned to a network ?
>docker network inspect XXXXXXXXXX .

(the XXXXXXXXX is the Id or the name of the container 
in the IPAM section thats for IP address managment and there you can see the type of network  and the subnet assigend to that network .)

# to connect a container to a custom network :
>docker network connect custom-net my container .
*you may connect a single container to multiple networks 

# to disconnect run:
>docker network disconnect custom-net my container .

# to remove a network run:
> docker network rm custom-net . 

# to remove all unused network run :
> docker network prune .

# How to get the subnet, gateway of the network c0a0b59a3807?
>docker network inspect c0a0b59a3807

# What is the default network driver used on a container if you haven’t specified one?
>bridge 

# What is the command to connect a running container with name myapp to the existing bridge network my-net?
>docker network connect  my-net myapp

# Which of the following commands would create a user-defined bridge network called my-net?
>docker network create --driver bridge my-net

# Overlay networks connect multiple Docker daemons together and enable swarm services to communicate with each other.
>true

# Which command is used to see the network settings and IP address assigned to a container with id c164825bb3d3 that uses the myapp image?
>docker inspect c164825bb3d3

# If you use the …… network mode for a container, that container’s network stack is not isolated from the Docker host (the container shares the host’s networking namespace), and the container does not get its own IP-address allocated.
> Host

# What is the command to remove all unused networks?
>docker network prune 

# What is the command to remove the my-net network?
>docker network rm my-net 

# Which command is used to list the default available networks?
>docker network ls


# Network namespaces:- 
are used by container like docker to implement network isolation .

# Namespace for example : 
if the host is your house so rooms is the namespaces within the house that you assigned to each for youe children the room helps in providing privacy to each Child , each child can only see whats  within his/her room they cant see whats happen outside the room as parents you have visibility  into all the rooms in the house as well as other areas in the house if you wish  you can  eatsblish connectivity  between two rooms in the house 
when you amke a container  you want to make sure that it is isolated that is does not see any other proccess on the host or any other containers so we create a special room for it on our host using namespace 
as far as the container is concerned it only sees the proccess run by it and things thas it is on its own host 
the undelying host however has visibility into all of the proccesses including those rinnig insde conatiners this can be seen when you list the proccesses from within the container you see a single process with a process id of one when you list  the same processes as a root user from  the underlying host you see all the other processes along with the process running inside the container this time with a diffrent process ID 

*when the continer is created we create a network namespaces for it that way it has no visibility to any network related information  on the host and the container can have its own virtual interfaces routing and ARP table 

*container are separated from the underlying host using namespaces  


# to create a new network namespace on linux host run:
> ip netns add red
> ip netns add blue 
# to list network namespaces run :
> ip netns 

# to list interface on my host run :
> ip link 

>ip netns exec red ip link 
# to orignal ip link run : 
>ip -n red link 
*to connect two namespaces its like useing vitual cable (pipe) 
# to connect two namespaces run 
> ip link add veth-red type veth peer name veth-blue 
# and then attach the interface to the namespace using this command:
> ip link set veth-red netns red 

>ip link set veth-blue netns blue 

# and to assign ip to them run :
> ip -n red addr add 192.168.15.1 dev veth-red 

>ip -n blue addr add 192.168.15.2 dev veth-blue .
# `to they can reach other run :
> ip -n red link set veth-red up

>ip -n blue link set veth-blue up

*if you have more than two use these comman > photo in screenshots commited name is "more than two vitual cable "

Q:Which command is used to see the IP address and other network settings assigned to a container with id 33373b1ccc3f that uses the wordpress image?

>Docker inspect 33373b1ccc3f

Q:Which command is used to disconnect the my-net network from the redis container?

>docker network disconnect my-net redis

Q:Which of the following commands would create a user-defined bridge network called dev-net?

>docker network create --driver bridge dev-net

Q:What is the default network driver used on a container if you haven’t specified one?

>bridge

Q:What is the command to list all available networks?

>docker network ls 

Q:What is the command to connect a running container with name myapp to the existing bridge network dev-net?

>docker network connect dev-net myapp

Q:Which command is used to see the details of the subnet and gateway of network id ce982a9edf65?

>docker network inspect ce982a9edf65

Q:What is the command to remove all unused networks from the Docker host?

>docker network prune 

Q:If you use the … network mode for a container, that container’s network stack is not isolated from the Docker host (the container shares the host’s 
networking namespace), and the container does not get its own IP-address allocated.

>Host


Q:What is the command to delete the network named connector?

>docker network rn connector 

Q:Which command is used to remove the my-vol volume?

>docker volume remove my-vol

Q:You can remove a vol1 which is in use by a container using the command docker volume rm --force vol1.

>false

Q:By default, all files inside an image are in a writable layer.

>false

Q:What is the command to get details of the volume my-vol such as the driver, mountpoint, volumename, …etc?

>docker volume inspect my-vol


Q:Which option is used to mount a volume ?

>mount

Q:By default, all files created inside a container are stored on a writable container layer.

>true 
Q:What is the command to create a volume with the name my-vol?

>docker volume create my-vol

Q:What is the command to remove unused volumes?

>docker volume prune 

Q:Volumes are the preferred mechanism for persisting data generated by and used by Docker containers.

>true

Q:What is the command to list volumes?

>docker volume ls

Q:The volumes are mounted as “readonly” by default inside the container if no options are specified.

>false

Q:Volumes are the preferred mechanism for persisting data generated by and used by Docker containers.

>true

Q:Which option is used to mount a volume ?

>mount


# docker image managment:-
where the images stored and how do we get them ?

images stored in a central image repository known as image registry (docker hub )

you can  create public and private images on docker hub .


you can create your own images and push it to docker hub  and even make the private .

# docker trusted registry service:-
>official images 

>verified images 

>users images 

*when you pull image that doesn't mean to create a container from it .

# why i need to create my own images ??
becuse mabye you cant find a service that you want or you and your tam decided the app you developing will be dockerized  fro shipping and deployment . 

# A build’s context is the set of files located in the specified PATH or URL, Which kind of resources can the URL parameter refer to ?
>git repo

# What is a recommended approach for installing packages and libraries while building an image?
>update and install commands in the same instraction.

# What is the file used to exclude temporary files such as log files or builds from the context during a build?
>.dockerignor

# Whenever a build is initiated by running the Docker build command, the files under the build context are transferred to the Docker daemon, at a temporary directory under the docker’s filesystem. Which directory are these files stored in?
>var/lib/docker/tmp


# If the build fails at a particular stage, it repurposes the previous layers from the cache and does not really rebuild them.
>true

# If you do not specify a tag name, you can’t build the image.
>false

# Which option can be used to disable the cache while building a docker image?
>--no-cache

# Build an image using a context build under path /tmp/docker and name it webapp.
>docker build /tmp/docker -t webapp

# forces the build to install a particular version of package regardless of what’s in the cache. This technique can also reduce failures due to unanticipated changes in required packages.
>version pinning 









