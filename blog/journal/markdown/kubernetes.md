# Kubernetes in Action

- **Kubernetes**
    * It steers our applications and reports on their status while we decide where we want the system to go.

    * Software system for automating the deployment and management of complex, large scale application systems composed of computer processes running in containers.

    * Provides an abstraction layer over the underlying hardware to both users and applications.

    * Deploying applications declaratively

        - We provide a description of the applications design
        - K8s turn the description into a running set of components.
        - K8s ensures the components keep running by restarting those that fail and recreating those that disappear.

    * It is an interface between distributed applications and the cluster of machines they run on.

    * Features
        - Service discovery
            - a mechanism that allows applications to find other applications and use the services they provide,
        - horizontal scaling
            - replicating our application to adjust to fluctuations in load.
        - load-balancing 
            - distributing load across all the application replicas
        - self-healing
            - keeping the system healthy by automatically restarting failed applications and moving them to healthy nodes after their nodes fail,
        - leader election
            - a mechanism that decides which instance of the application should be active while the others remain idle but ready to take over if the active instance fails.

    * Benefits

        - Automatically adjusting to changing load.
        - Simplifying Application Development
            - Applications can query Kubernetes API to obtain info about their environment.

- Computer Cluster
    - Machines are split into two groups.
        - Kubernetes Control Plane
            - Brain of cluster
            - Master node(s)
            - Components that control the cluster run on these nodes.
            - One or more master nodes.
        - Kubernetes Workload Plane.
            - Runs our apps
            - Worker nodes(s)
            - Kubernetes components on worker nodes communicate with those running on the master, but never with each other.
            - Each node also runs several kubernetes components that manage the apps running on the node.

    - Components of the Kubernetes Control Plane
        - etcd

            * Distributed datastore persists the objects we create through API. The server is the only component that talks to etcd.

        - Kubernetes API Server
            * Developers and operators create objects via the API
            * Kubernetes Components on worker nodes communicate only with API Server.
            * API Server stores objects in distributed data store. No other components access the store directly.

        - Scheduler
            * Decides on which worker node eaach application instance should run.

        - Controllers
            * Each controller has a different task. Most of them create objects from the objects we create.

    - Worker Node Components

        * Several Kubernetes components also run on these nodes. They perform the task of running, monitoring and providing connectivity between our applications.

        - Kubelet
            - Nodes agent.
            - Manages the applications running on the node.
            - Instructs the container runtime to run the apps and get their status.

        - Kubelet Proxy

            - Creates a virtual load balancer for each application that needs it.

        - Container Runtime
            - Typically Docker, but other runtimes are increasingly common.
            - Applications run in containers and are started by the Container Runtime.


    - Non-production clusters can use a single master node, but highly avaliable clusters use at least three physical master nodes to host the Control Plane. 
    - The number of worker nodes depends on the number of application we will deploy.
    - Applications are deployed via the Kubernetes API.
    - Note: **Each application must be small enough to fit on one of the worker nodes.**

* Scalability

    - System's ability to grow. We can scale down, scale up and scale out.

    * Horizontal Scaling
        - Involves adding more machnines or nodes to a system.
        - Workload distribution
            - Workload is distributed across multiple nodes. Parts of the workload reside on these different nodes.

    * Veritical Scaling
        - Involves adding more power (CPU, RAM, storage etc.) to an existing machine.
        - Workload distribution
            - Single node handles the entire workload.

* Monolithic Application

    - Consists of one or a very small number of applications.

* Microservices

    - Each microservice is an application which can be installed, configured and managed separately.
    - Require automated management due to their multiplicity and distributed nature.

* How Kubernetes runs an application

    - Everything in kubernetes is represented by objects.
    - We create and retrieve these objects via Kubernetes API.
    - Our application consists of several types of these objects.
        - One type represents the application deployment as a whole.
        - Another represents a running instance of our application.
        - Another represents the service provided by a set of these instances and allows reaching them at a single IP address, and there are many others.

    - Deploying the application
        - We submit application manifest to the Kubernetes API. The API server writes the objects defined in the manifest to etcd. In addition, it notifies all interested components that these objects have been created.

        - A controller notices the newly created objects and creates several new objects - one for each appplication instance.

        - The scheduler assigns a node to each instance. It selects the best worker node for each new application instance object and assigns it to the instance - by modifying the object via the API.

        - The Kubelet notices that an instance is assigned to the Kubelet's node. It runs the application instance via the container runtime.

        - The Kube Proxy notices that the application instances are ready to accept connections from clients and configures a load balancer for them.Because an application deployment can consist of multiple application instances, a load balancer is required to expose them at a single IP address.


        - The Kubelets and the controllers monitor the system and keep the applications running.

* Kubernetes can run directly on our bare-metal machines or in virtual machines running in our data center.
* Note: If our system has more than 20 microservices, we will most likely benefit from the integration of kubernetes. 

* Understanding containers

    * Containers are isolated processes running in the existing OS that consumes only the resources the app consumes.
    * Containers start the application faster.
    * Containers make system calls on the single kernel running in the host os.
    - Differences
        - Applications running on bare metal directly
            - All applications use the same kernel and see the same hardware. They are not isolated.
        - Applications running in virtual machines
            - Applications A and B running one VM is strongly isolated from Application C running in another VM and each VM run separate kernels.
            - Hypervisor splits the physical hardware resources into smaller sets of virtual resources that the operating system in each VM can use.
        - Applications running in isolated containers
            - Applications use the same kernel, but kernel isolates them from each other.
            - Only one application per container.
    - Security implications
        - VMS provide complete isolation
        - Containers share memory space, where as each VM uses its own chunk of memory.Therefore, if we don't limit the amount of memory that a container can use, this could cause other containers to run out of memory or cause their data to be swapped out to disk.
    - Docker Platform
        - It is a platform for packaging, distributing and running applications.
        - It allows us to package our application along with its entire environment.
        - Docker images
            - Something we can package our application and its environment.
            - It contains whole filesystem that the application will use and additional metadata, such as path to the executable file to run when the image is executed, the ports the application listen on, and other information about the image.
        - Registries
            - Repository of container images
        - Containers
            - A container is instantiated from a container image.
            - Running containr is a normal process running in the host OS, but its environment is isolated from that of the host and the environments of other processes.
            - Image Layers
                - These layers are smaller and can be shared and reused across multiple images.
                - The filesystems are isolated by Copy-on-Write(Cow) mechanism.The filesystem of a container consists of read-only layers from the container image and an additional read/write layer stacked on top.When an application running in cotainerer A changes a file in one of the read-only layers, the entire file is copied into the container's read/write layers and the file contents are changed there. Since each container has its own writable layer, changes to shared files are not visible in any other container. When we delete a file, it is only marked as deleted in the read/write layer, but it's still present in one or more of the layers below.
            - Portability Limitations
                - containers don't have their own kernel, If a containerized application requires a particular kernel version(pariticular kernel module), it may not work on every computer.
                - Containerized app built for a specific hardware architecture can only run on computers with the same architecture. For this, we need a VM to emulate the architecture.
            - **VMS are enabled through virualization support in CPU and by virtualization software on the host, containers are enabled by the Linux kernel itself.**

    - Open Container Initiative (OCI)
        - To create open industry standards around container formats and runtime.
        - OCI Image Format Specification.
            - Defines a standard interface for container runtimes with the aim of standardizing the creation, configuration and execution of containers.
    
    - Running hello world container
        ```
        docker run busybox echo "Hello World"
        ```
        - The above commands tells the Docker CLI to run the container.
        - Docker CLI sends an API request to Docker Daemon
        - Docker daemon checks if the busybox image is in the local cache
        - If the busybox image isn't available locally, Docker pulls it from the image registry.
        - Docker creates the hello-world container and runs the echo "Hello World" command inside it.
    
        - Note: If the local computer runs a Linux OS, the Docker CLI tool and the daemon both run in host OS. If it runs macOS or Windows, the daemon and the containers run in the Linux VM.
    - Dockerfile

        ```
        FROM node:18
        ADD app.js /app.js
        ENTRYPOINT ["node","app.js"]
        ```
    - Building the image

        ```
        docker build -t myapp:latest .
        ```

        We tell Docker to build an image called myapp based on the contents of the current directory. Docker reads Dockerfile in the directory and builds the image based on the directive in the file.

    - Listing locally stored images

        ```
        docker images
        ```
    - To see the layers of an image and their size by running

        ```
        docker history myapp:latest
        ```

        * ADD and ENTRYPOINT will be uppermost layers.

    - Running the container image

        ```
        docker run --name myapp-container -p 7080:8080 -d myapp
        ```

        * -d flag : detached from the console.
        * Port 7080 on the host computer is mapped to port 8080 in the container

    - Listing all running containers

        ```
        docker ps
        ```

    - To see the additional information about the running container

        ```
        docker inspect myapp-container
        ```

    - Inspecting the application logs (standard o/p and error streams written by the application)

        ```
        docker logs myapp-container
        ```

    - Tagging an image under an additional tag

        ```
        docker tag myapp gireeshcse/myapp:1.0

        docker images | head
        ```

    - Pushing the image to Docker Hub

        ```
        docker login -u myid -p mypassword docker.io
        ```

    - Once logged in, push the myid/myapp:1.0 image to docker hub

        ```
        docker push gireeshcse/myapp:1.0
        ```

    - Running the app on other hosts

        ```
        docker run -p 7090:8080 -d gireeshcse/myapp:1.0
        ```

    - Stopping and deleting containers

        ```
        docker stop myapp-container
        docker rm myapp-container

        docker ps -a
        docker rmi myapp:latest
        docker image prune # remove all dangling images
        ```
    
    - Linux Namespaces
        - Ensures that each process has its own view of the system.
        - Process running in a container will only see some of the files, processes, and n/w interfaces on the system, as well as a different system hostname, just as if it is were running in a separate virtual machine.

    - Running a Shell inside an existing container

        ```
        docker exec -it myapp-container bash
        ```

    - List the processes running in a container

        ```
        ps aux
        ```
    If you need even finer control over what sys-calls a program can make, you can use seccomp
(Secure Computing Mode). You can create a custom seccomp profile by creating a JSON file
that lists the sys-calls that the container using the profile is allowed to make. You then
provide the file to Docker when you create the container.
    - Limiting a process resource usage with Linux Control Groups
        - Linux Namespaces make it possible for processes to access only some of the host's resources, but they don't limit how much of a single resource each process can consume.Ex: It can allow a process to access only a particular n/w interface, but we can't limit the n/w bandwidth the process consumes.
        - **CGROUPS**
            - Second Linux kernel feature that makes containers possible - **Linux Control Groups**
            - It limits, accounts for and isolates system resources such as CPU, memory and disk or n/w bandwidth.
            - Limiting a container's use of the resources
                - To allow container to only use cores one and two.
                    ```
                    docker run --cpuset-cpus="1,2" ...
                    ```
                - To allow container to use only half of a cpu core,
                    ```
                    docker run --cpus="0.5" ...
                    ```
                - Options to limit container memory and swap usage:
                    - --memory
                    - --memory-reservation
                    - --kernel-memory
                    - --memory-swap
                    - --memory-swappiness
                    ```
                    docker run --memory="100m" ...
                    ```

# Kubernetes Application Developer

