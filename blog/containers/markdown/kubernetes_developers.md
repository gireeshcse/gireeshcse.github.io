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

        kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/recommended.yaml
        kubectl describe secret -n kube-system

From above command copy the token of type **kubernetes.io/service-account-token**

        kubectl proxy

Error:  first record does not look like a tls handshake kubernetes

Change

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

# PODS

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