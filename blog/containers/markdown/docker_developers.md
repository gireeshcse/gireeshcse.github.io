# Docker

* Debug your app, not your environment
* Securely build and share any application, anywhere
* The fastest way to design and deliver containerized applications and microservices on the desktop and cloud.

To get started

    docker run -dp 80:80 docker/getting-started

* -d - run the container in detached mode (in the background)
* -p 80:80 - map port 80 of the host to port 80 in the container
* docker/getting-started - the image to use

## Container

* Container is simply another process on your machine that has been isolated from all other processes on the host machine.
* A standard unit of software.

### Common Features

* They have to run on **single** host
* Containers have a **root** process.
* Isolated

#### chroot in Linux

A **chroot** command on Unix operating systems is an operation that changes the apparent root directory for the current running process and its children. When you change root to another directory you cannot access files and commands outside that directory. This directory is called a **chroot jail**.

#### pivot_root in Linux

 Changes the **root mount** in the mount namespace of the calling process.  More precisely, it moves the root mount to the directory **put_old** and makes **new_root** the new root mount.

Has the benefit of putting the old mounts into a separate directory. These old mounts could be unmounted afterwards to make the filesystem completely invisible to broken out processes.

### Container Image

When running a container, it uses an isolated filesystem. This custom filesystem is provided by a container image. Since the image contains the container's filesystem, it must contain everything needed to run an application - all dependencies, configuration, scripts, binaries, etc. The image also contains other configuration for the container, such as environment variables, a default command to run, and other metadata.

## Example: 1

- Download the the app code from 

    https://github.com/docker/getting-started/tree/master/app

- Define **Dockerfile**

        FROM node:12-alpine
        WORKDIR /app
        COPY . .
        RUN yarn install --production
        CMD ["node", "/app/src/index.js"]
    
    * CMD directive specifies the default command to run when starting a container from this image.

- Run following command

        docker build -t getting-started .

    * -t flag tags our image

- Starting Container

    docker run -dp 3000:3000 getting-started

    * Access application [http://localhost:3000/](http://localhost:3000/) in browser. 

- docker ps
    
        CONTAINER ID  IMAGE            COMMAND CREATED        STATUS        PORTS                 
        5894edcec7d1  getting-started  "docker 2 minutes ago  Up 2 minutes  0.0.0.0:3000->3000/tcp
        ab95b5dcdb52  docker/getting-  "nginx  50 minutes ago Up 50 minutes 0.0.0.0:80->80/tcp

- Important Commands

        docker stop <the-container-id> # To stop container from running
        docker rm <the-container-id> # To remove a container
        docker rm -f <the-container-id> # force flag
        docker image ls # to list the images in system
        docker tag getting-started gireeshcse/getting-started # to create a tag 
        docker login -u YOUR-USER-NAME # to login to the account i.e one time activity
        docker exec <container-id> cat /data.txt 
        docker run -it ubuntu ls /
        docker logs -f <container-id> # We can watch logs 
        docker ps # shows running containers by default 
        docker ps -a #all containers which are running as well as stopped
        docker image history <image-name>
        docker inspect <container_id>
        #  matches containers with the color label regardless of its value.
        docker ps --filter "label=color" 
        # matches containers with the color label with the blue value.
        docker ps --filter "label=color=blue"
        docker ps --filter "name=nostalgic_stallman"
        docker ps --filter "name=nostalgic" # filter for a substring in a name 
        docker ps -a --filter 'exited=0'
        # exited with status of 137 meaning a SIGKILL(9) killed them.
        docker ps -a --filter 'exited=137'
        # created, restarting, running, removing, paused, exited and dead
        docker ps --filter status=running
        docker ps --filter publish=80
        docker run -d --publish=80 busybox top
        docker run -d --expose=8080 busybox top
        # all containers that have exposed TCP port in the range of 8000-8080
        docker ps --filter expose=8000-8080/tcp
        docker ps --filter publish=80/udp
        docker images -a | grep "pattern" | awk '{print $3}' | xargs docker rmi
        docker images -a | grep "^<none>.*" | awk '{print "\"image="$3"\""}'

- docker ps Formatting
  * Placeholders
    - .ID
    - .Image
    - .Command
    - .CreatedAt
    - .RunningFor
    - .Ports
    - .Status
    - .Size
    - .Names
    - .Labels
    - .Label
    - .Mounts
    - .Networks

                $ docker ps -a --format "{{.ID}}: {{.Image}}"
                338b9bea3cb0: node:12-alpine
                4fcc126a4a30: nicolaka/netshoot
                f0fd3dd0beb3: ubuntu
                701690bf6bdb: ubuntu

- Update App

    - Update Source code 
    - Build (docker build -t getting-started .)
    - Stop and remove the running container.

### Sharing the App

* Create an account in [Docker Hub](https://hub.docker.com/)
* Create Repo in Docker Hub (Public)
* Push the image (docker push gireeshcse/getting-started)

### Testing your images

[http://play-with-docker.com/](http://play-with-docker.com/)

Create instance and test your app.

### Persisting our DB

#### Containers Filesystem

When a container runs, it uses the various layers from an image for its filesystem. Each container also gets its own "scratch space" to create/update/remove files. Any changes won't be seen in another container, even if they are using the same image.

#### Container Volumes

Volumes provide the ability to connect specific filesystem paths of the container back to the host machine. If a directory in the container is mounted, changes in that directory are also seen on the host machine. If we mount that same directory across container restarts, we'd see the same files.

**Persisting Data**

By creating a volume and attaching (often called "mounting") it to the directory where the data is stored in , we can persist the data. 
If our container writes to the mounted directory, it will be persisted to the host in the volume.

**Named Volume**

Named Volume is simply a bucket of data. Docker maintains the physical location on the disk and we need to remember the name of the volume(Every time we use the volume)

     docker volume create todo-db
     docker run -dp 3000:3000 -v todo-db:/etc/todos getting-started
     
* -v flag to specify a volume mount

**Diving into our Volume**

            $ docker volume inspect todo-db
            [
                {
                    "CreatedAt": "2020-03-26T21:05:29+05:30",
                    "Driver": "local",
                    "Labels": {},
                    "Mountpoint": "/var/lib/docker/volumes/todo-db/_data",
                    "Name": "todo-db",
                    "Options": {},
                    "Scope": "local"
                }
            ]
Note: Generally We will need to have root access to access this directory from the host.

**Bind Mounts**

With bind mounts, we control the exact mountpoint on the host. We can use this to persist data, but is often used to provide additional data into containers. When working on an application, we can use a bind mount to mount our source code into the container to let it see code changes, respond, and let us see the changes right away.

**Starting a Dev-Mode Container**

* Mount our source code into the container
* Install all dependencies, including the "dev" dependencies
* Start nodemon to watch for filesystem changes(For nodejs applications)

            docker run -dp 3000:3000 \
                -w /app -v $PWD:/app \
                node:12-alpine \
                sh -c "yarn install && yarn run dev"


    * -w /app - sets the "working directory" or the current directory that the command will run from

* After making changes and testing changes, we can finally build our new image.

Note: *Bind mounts is very common for local development setups*

### Rule of Thumb

**each container should do one thing and do it well**

**Some of the Reasons**

* There's a good chance we'd have to scale APIs and front-ends differently than databases
* Separate containers let us version and update versions in isolation
* While we may use a container for the database locally, we may want to use a managed service for the database in production. We don't want to ship your database engine with your app then.
* Running multiple processes will require a process manager (the container only starts one process), which adds complexity to container startup/shutdown

### Container Networking

If two containers are on the same network, they can talk to each other. If they aren't, they can't.

* Creating a network

        docker network create todo-app

* Start MySQL Container

        docker run -d \
            --network todo-app --network-alias mysql \
            -v todo-mysql-data:/var/lib/mysql \
            -e MYSQL_ROOT_PASSWORD=secret \
            -e MYSQL_DATABASE=todos \
            mysql:5.7

    - -e to specify the environmental variables. FOr valid environmental varaibles for mysql,Refer [https://hub.docker.com/_/mysql/](https://hub.docker.com/_/mysql/)

* To Verify

        docker exec -it <mysql-container-id> mysql -p

Enter the password and give show databases

        mysql> show databases;
        +--------------------+
        | Database           |
        +--------------------+
        | information_schema |
        | mysql              |
        | performance_schema |
        | sys                |
        | todos              |
        +--------------------+
        5 rows in set (0.00 sec)

* To test and running another container in the same network

    docker run -it --network todo-app nicolaka/netshoot
    dig mysql # DNS Tool which we can look up the IP address fro hostname **mysql**

        ; <<>> DiG 9.14.8 <<>> mysql
        ;; global options: +cmd
        ;; Got answer:
        ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 14072
        ;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0

        ;; QUESTION SECTION:
        ;mysql.                         IN      A

        ;; ANSWER SECTION:
        mysql.                  600     IN      A       172.19.0.2

        ;; Query time: 88 msec
        ;; SERVER: 127.0.0.11#53(127.0.0.11)
        ;; WHEN: Thu Mar 26 17:12:55 UTC 2020
        ;; MSG SIZE  rcvd: 44

**MYSQl Environmental Variables**

* MYSQL_HOST - the hostname for the running MySQL server
* MYSQL_USER - the username to use for the connection
* MYSQL_PASSWORD - the password to use for the connection
* MYSQL_DB - the database to use once connected


* Run Our App using mysql containers

        $ docker run -dp 3000:3000 \
        -w /app -v $PWD:/app \
        --network todo-app \
        -e MYSQL_HOST=mysql \
        -e MYSQL_USER=root \
        -e MYSQL_PASSWORD=secret \
        -e MYSQL_DB=todos \
        node:12-alpine \
        sh -c "yarn install && yarn run dev"

        $ docker logs 0c48a4df697e
        yarn install v1.22.0
        [1/4] Resolving packages...
        success Already up-to-date.
        Done in 0.57s.
        yarn run v1.22.0
        $ nodemon src/index.js
        [nodemon] 1.19.2
        [nodemon] to restart at any time, enter `rs`
        [nodemon] watching dir(s): *.*
        [nodemon] starting `node src/index.js`
        Waiting for mysql:3306.
        Connected!
        Connected to mysql db at host mysql
        Listening on port 3000

        $ docker exec -it 4934e3a2c10d mysql -p todos

        mysql> select * from todo_items;
        +--------------------------------------+-------------------------+-----------+
        | id                                   | name                    | completed |
        +--------------------------------------+-------------------------+-----------+
        | 344ba928-f463-45d4-8015-2e571fcc7948 | Wake up @5:00AM         |         0 |
        | 1c3833af-4cde-45bc-9d21-663ce045e332 | Say Good Night to Finch |         0 |
        +--------------------------------------+-------------------------+-----------+
        2 rows in set (0.00 sec)

## Docker Compose

* Docker Compose is a tool that was developed to help define and share multi-container applications.
* With Compose, we can create a YAML file to define the services and with a single command, can spin everything up or tear it all down.

        $ docker-compose version
        docker-compose version 1.24.1, build 4667896b
        docker-py version: 3.7.3
        CPython version: 3.6.8
        OpenSSL version: OpenSSL 1.1.0j  20 Nov 2018

* Create **docker-compose.yml**

        version: "3.7"

        services:
          app:
            image: node:12-alpine
            command: sh -c "yarn install && yarn run dev"
            ports:
              - 3000:3000
            working_dir: /app
            volumes:
              - ./:/app
            environment:
              MYSQL_HOST: mysql
              MYSQL_USER: root
              MYSQL_PASSWORD: secret
              MYSQL_DB: todos

          mysql:
            image: mysql:5.7
            volumes:
              - todo-mysql-data:/var/lib/mysql
            environment: 
              MYSQL_ROOT_PASSWORD: secret
              MYSQL_DATABASE: todos

        volumes:
          todo-mysql-data:

        $ docker-compose up -d
        Creating network "app_default" with the default driver
        Creating volume "app_todo-mysql-data" with default driver
        Creating app_mysql_1 ... done
        Creating app_app_1   ... done

**Note:** By default, Docker Compose automatically creates a network specifically for the application stack (which is why we didn't define one in the compose file).

TO Check the logs

        docker-compose logs -f

 The -f flag "follows" the log, so will give you live output as it's generated.

        docker-compose down
        docker-compose down --volumes # Remove volumes as well

## Best Practices

    $ docker image history getting-started

    IMAGE               CREATED       CREATED BY                                      SIZE  
    2fd3a9645cf6        5 hours ago   /bin/sh -c #(nop)  CMD ["node" "/app/src/ind…   0B                  
    ffa2a47db0b7        5 hours ago   /bin/sh -c yarn install --production            83.2MB              
    296afaf1a376        5 hours ago   /bin/sh -c #(nop) COPY dir:1bc092cfe17c580ca…   4.63MB              
    9c5526543a3f        6 hours ago   /bin/sh -c #(nop) WORKDIR /app                  0B                  
    f77abbe89ac1        2 days ago    /bin/sh -c #(nop)  CMD ["node"]                 0B                  
    <missing>           2 days ago    /bin/sh -c #(nop)  ENTRYPOINT ["docker-entry…   0B                  
    <missing>           2 days ago    /bin/sh -c #(nop) COPY file:238737301d473041…   116B                
    <missing>           2 days ago    /bin/sh -c apk add --no-cache --virtual .bui…   7.62MB              
    <missing>           2 days ago    /bin/sh -c #(nop)  ENV YARN_VERSION=1.22.0      0B                  
    <missing>           2 days ago    /bin/sh -c addgroup -g 1000 node     && addu…   74.9MB              
    <missing>           2 days ago    /bin/sh -c #(nop)  ENV NODE_VERSION=12.16.1     0B                  
    <missing>           2 days ago    /bin/sh -c #(nop)  CMD ["/bin/sh"]              0B                  
    <missing>           2 days ago    /bin/sh -c #(nop) ADD file:0c4555f363c2672e3…   5.6MB 

* docker image history --no-trunc getting-started 

        # Dockerfile
        FROM node:12-alpine
        WORKDIR /app
        COPY . .
        RUN yarn install --production
        CMD ["node", "/app/src/index.js"]

        # Dockerfile
        FROM node:12-alpine
        WORKDIR /app
        COPY package.json yarn.lock ./
        RUN yarn install --production
        COPY . .
        CMD ["node", "/app/src/index.js"]

**Once a layer changes, all downstream layers have to be recreated as well**

### Multi-Stage Builds

Multi-stage builds are an incredibly powerful tool to help use multiple stages to create an image. There are several advantages for them:

* Separate build-time dependencies from runtime dependencies
* Reduce overall image size by shipping only what your app needs to run

#### Maven/Tomcat Example

When building Java-based applications, a JDK is needed to compile the source code to Java bytecode. However, that JDK isn't needed in production. Also, you might be using tools like Maven or Gradle to help build the app. Those also aren't needed in our final image. Multi-stage builds help.

        FROM maven AS build
        WORKDIR /app
        COPY . .
        RUN mvn package

        FROM tomcat
        COPY --from=build /app/target/file.war /usr/local/tomcat/webapps 

In this example, we use one stage (called build) to perform the actual Java build using Maven. In the second stage (starting at FROM tomcat), we copy in files from the build stage. The final image is only the last stage being created (which can be overridden using the --target flag).

#### React Example
When building React applications, we need a Node environment to compile the JS code (typically JSX), SASS stylesheets, and more into static HTML, JS, and CSS. If we aren't doing server-side rendering, we don't even need a Node environment for our production build. Why not ship the static resources in a static nginx container?


        FROM node:12 AS build
        WORKDIR /app
        COPY package* yarn.lock ./
        RUN yarn install
        COPY public ./public
        COPY src ./src
        RUN yarn run build

        FROM nginx:alpine
        COPY --from=build /app/build /usr/share/nginx/html
        
Here, we are using a node:12 image to perform the build (maximizing layer caching) and then copying the output into an nginx container.

## Important Links

* [Why you shouldn't use ENV variables for secret data](https://diogomonica.com/2017/03/27/why-you-shouldnt-use-env-variables-for-secret-data/)