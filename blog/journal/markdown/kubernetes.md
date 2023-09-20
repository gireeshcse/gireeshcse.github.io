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

- Computer Cluster
    - Machines are split into two groups.
        - Kubernetes Control Plane
            - Brain of cluster
        - Kubernetes Workload Plane.
            - Runs our apps

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

# Kubernetes Application Developer

