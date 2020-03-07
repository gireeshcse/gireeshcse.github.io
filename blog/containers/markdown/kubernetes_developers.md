# Kubernetes 

Kubernetes (K8s) is an open-source system for automating deployment, scaling, and management of containerized applications.

## Advantages

* We can package an app and we can let kubernetes to manage it for us
* Management of containers
* Elimination of single points of failures
* Scales containers
* Updates containers without bringing down the application.
* Have a robust networking and persistent storage options.

**Conductor of containers**
**Provides a declarative way to define a cluster' state**
**Contains one or more master nodes and worker nodes(can be Physical servers,VMs).The workers nodes contains PODS which contains containers**

* POD         ----> Suit
* Container   ----> Person
* Store(etcd)(acts as a database for cluster), 
* controller manager(Takes requests and uses scheduler to perform/act upon actions),
* API Server(To interact with the cluster to give instruction to go from one state to other<--kubectl)

**Each node has a Kubelet to communicate with the master,Container runtime to run containers within the PODs and  Kube-Proxy ensures each pod has a IP address**

## Benefits

* Orchestrate Containers
* Zero-Downtime Deployemnts 
* Self Healings
* Scale Containers

### For Developers

* Emulate production locally
* Move from Docker Compose to Kubernetes
* Create an end-to-end testing environment
* To Ensure application scales properly
* To Ensure secrets/config are working properly.
* Performance testing scenarios
* Workload scenarios(CI/CD and more)
* Learn how to leverage deployment options
* Help DevOps create resources and solve problems

## Running Locally

* Install [Minikude](https://kubernetes.io/docs/tasks/tools/install-minikube/) 
* [Docker Desktop](https://www.docker.com/products/docker-desktop) available for Mac and Windows.

Note: (To use KVM driver for ubuntu 18.04)(https://www.linuxtechi.com/install-configure-kvm-ubuntu-18-04-server/)

[Minikube drivers](https://minikube.sigs.k8s.io/docs/reference/drivers/)

Note: sudo minikube start --vm-driver=none for Ubuntu 18.04 minikube version 1.6.2 worked [minikube release](https://github.com/kubernetes/minikube/releases/)

* To install latest minikube (Linux) Installed Binary

    curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
    && chmod +x minikube
    sudo install minikube /usr/local/bin/

## minikube Commands(Generally reuires sudo)

    
    # To starts a local kubernetes cluster Generally run with 
    # sudo minikube --vm-driver=none if running in host directly not in VM
    minikube start 
    minikube stop # stop a local kubernetes cluster
    minikube dashboard # access the kubernetes dashboard running within the minikube cluster
    minikube delete # deletes a local kubernetes cluster
    minikube start -p <name>' to create a new cluster, or 'minikube delete' to delete this one
    minikube status

## kubectl commands 

        kubectl version 
        kubectl cluster-info
        kubectl get all # all info about Kubernetes Pods,Deployments,Services, and more
        kubectl run [container-name]  --image=[image-name] # simple way to create a deployment for a POD
        kubectl port-forward [pod] [ports] # forward a port to allow external access
        kubectl expose [port] # expose a port for a Deployment/Pod
        kubectl create [resource]  # create a resource
        kubectl apply [resource]  # createor modify a resource 
        kubectl --help
        kubectl get pods
        kubectl get services
        
## Enabling Web UI Dashboard 

[For more Info](https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/)

        sudo minkube dashboard 

        or

        kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/\
        v2.0.0-beta8/aio/deploy/recommended.yaml
        kubectl describe secret -n kube-system

From above command copy the token of type **kubernetes.io/service-account-token**

        kubectl proxy

### Error:  first record does not look like a tls handshake kubernetes

Change (https to http)

http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

to 

http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/


## Featues

* Service discovery and load balancing 

    Kubernetes gives Pods their own IP addresses and a single DNS name for a set of Pods, and can load-balance across them.

* Storage Orchestration

    Automatically mount the storage system of our choice, whether from local storage, a public cloud provider such as GCP or AWS, or a network storage system such as NFS, iSCSI, Gluster, Ceph, Cinder, or Flocker.

* Self Healing 

    Restarts containers that fail, replaces and reschedules containers when nodes die, kills containers that don’t respond to your user-defined health check, and doesn’t advertise them to clients until they are ready to serve.

* Automating rollouts and rollbacks

    Kubernetes progressively rolls out changes to your application or its configuration, while monitoring application health to ensure it doesn’t kill all your instances at the same time. If something goes wrong, Kubernetes will rollback the change for you. Take advantage of a growing ecosystem of deployment solutions.

* Secret and configuration management

    Deploy and update secrets and application configuration without rebuilding our image and without exposing secrets in your stack configuration.

* Horizontal scaling 

    Scale your application up and down with a simple command, with a UI, or automatically based on CPU usage.

# Pods

* A Pod is the basic execution unit of a Kubernetes application-the smallest and simplest unit in the Kubernetes object model that you create or deploy.
* Pods run containers
* Pods acts as a environment for containers.
* As a Developer we need to organise the application "parts" into **Pods** (Server, caching, APIs, database, etc.)
* Pod IP, memory, volumes, etc. shared across containers
* we can scale horizontally by adding Pod replicas
* **Pods live and die but never come back to life.** New one is created

* Pod containers share the same Network namespace(share IP/port)
* Pod containers have the same loopback network interfaces(localhost)
* Containers processes need to bind to different ports within a Pod
* Ports can be reused by containers in separate Pods
* Pods never span nodes

Note: Pods have different ips (10.0.0.33, 10.0.0.43)

## Running a Pod

Use any of the following

* kubectl run command
* kubectl create/apply command with a yaml file.

        # Run the nginx:alpine container in a Pod
        kubectl run [podname]  --image=nginx:alpine
        kubectl create deployment [podname] --image=nginx:alpine

        # Examples
        kubectl run sample-nginx-alpine-pod --image=nginx:alpine # deprecated
        kubectl create deployment nginx-pod --image=nginx:alpine

        # list only pods
        kubectl get pods
        kubectl get pods --watch
        # list all resources 
        kubectl get all

**Pods and containers are only accessible within the kubernetes cluster by default**
**One way to expose a container port externally: kubectl port-forward**
**Image to run a pod is a docker image**

        # Enable Pod Container to be 
        # called externally
        kubectl port-forward [name-of-pd]  external_port:internal_port
        kubectl port-forward pod/nginx-pod-6fc99f67cd-h4zxr  8001:80 
        # 127.0.0.1:8001 to access the application

        # will cause pod to be recreated
        kubectl delete pod [name-of-pod]
        kubectl delete pod nginx-pod-6fc99f67cd-h4zxr

        # Delete Deployment that manages the pod
        # use kubectl get all to get deployment name of the pod
        kubectl delete deployment [name-of-deployment] 
        kubectl delete deployment nginx-pod

### YAML Review

* Composed of maps and lists
* Indentation matters (be consistent!)
* ALways use spaces
* Maps:
    - name:value pairs
    - Maps can contain other maps for more complex data structures
* Lists:
    - Sequence of items
    - Multipe maps can be defined in a list

#### Example

    key: value
    complexMap:
        key1: value
        key2:
        subKey: value

    items:
        - item1
        - item2
        itemsMap:
        - map1: value
            map1Prop: value
        - map2L value
            map2Prop: value

### nginx.pod.yml

    apiVersion: v1
    kind: Pod
    metadata:
      name: my-nginx
    spec:
      containers:
      - name: my-nginx
        image: nginx:alpine

### creating a Pod using YAML

    # perform a trial create and validate the YAML (validate true is default)
    kubectl create --filename file.pod.yml --dry-run --validate=true
    kubectl create -f nginx.prod.yml --dry-run --validate=true

    # Create a Pod from YAML
    # Will error if Pod already exists
    kubectl create -f file.pod.yml

    # altenative way to create or apply changes to a
    # Pod from YAML
    kubectl apply -f file.pod.yml
    kubectl apply -f nginx.prod.yml
    # above command creates a warning 
    # use --save-config when you what to use 
    # kubectl apply in future
    kubectl create -f file.prod.yml --save-config # Store current properties in resource's annotations

**--save-config** causes the resource's configuration settings to be saved in the **annotations**.Having this allows in-place changes to be made to a Pod in the future using **kubectl apply**

kubectl edit or kubectl patch can also be used to change small or subset of changes to a Pod.

    # delete a Pod
    kubectl delete pod [name-of-pod]

    # delete Pod using YAML file that created it
    kubectl delete -f file.pod.yml
    kubectl delete -f nginx.prod.yml

### nginx.pod.yml

    apiVersion: v1
    kind: Pod
    metadata:
      name: my-nginx
      labels:
        app: nginx
        rel: stable
    spec:
      containers:
      - name: my-nginx
        image: nginx:alpine
        ports:
        - containerPort: 80

Note: labels are used in deployments

    kubectl create -f nginx.prod.yml --save-config
    # shows output in YAML this is because of --save-config(added annotations to the o/p)
    kubectl get pod my-nginx -o yaml 
    # To trobleshoot the Pod this output is useful
    kubectl describe pod [pod-name]
    kubectl describe pod my-nginx
    kubectl apply -f nginx.prod.yml

    # to go into pod with interactive shell
    kubectl exec [pod-name] -it sh
    kubectl exec my-nginx -it sh # enter exit the shell

    kubectl edit -f nginx.pod.yml

    
 ### Error socat not found - minikube 

    kubectl port-forward my-nginx 8001:80
    
[Solution Link](https://github.com/kubernetes/minikube/issues/68#issuecomment-344346923)
  if you're running the none driver, you'll need a whole host of dependencies that kubernetes requires: docker, iptables, socat, certain kernel modules enabled, etc
  sudo apt-get install socat fixed the issue.


# Containers

## Benefits

* Accelerate Developer Onboarding 
* Eliminate App Conflicts 
* Environment Consistency
* Ship Software Faster

# Deployments

# Services

# Storage Options

# ConfigMaps and Secrets