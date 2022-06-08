#### Hosts

Any device which sends or receives the traffic.

    - Clients - Initiates the requests
    - Servers - Responds

#### IP Address

- Identity of the Host(Just like phone number or email address)
- 32 bit 
- Hierarchically assigned (Subnetting)

#### Network

- A Network is what transports traffic between Hosts
- Anytime two hosts are connected, we have a network.
- A logical grouping of hosts require similar connectivity.
- Networks can contain sub-networks or subnets
- Networks connect to other networks.
    - Internet.

#### Repeater

* Data crossing a wire decays as it travels.
* Repeaters regenerates signals. Allows communications across greater distances.

#### Hub

* Connecting hosts directly to each other doesn't scale.
* Hubs are simply multi-port Repeaters.
    - Facilitates scaling communication between additional hosts.
    - Everyone receives everyones else's data.

#### Bridge

* Bridges sit between Hub-Connected hosts.
* Have only two ports
* learns which hosts are on each side. 

#### Switch

* Combination of Hubs and Bridges
* Multiple Ports.
* Learns which hosts are on each port.


* Facilitate communication within a network.
* Hosts on  a network share the same IP address space.(192.168.1.x)
* Network: Grouping of hosts which require similar connectivity.

#### Router

* Facilitate communication between networks
* which may also connected to internet.
* Provides a traffic control point(security, filtering, redirecting)
* Sit on boundaries of network.
* Routers learn which networks they are attached to knowns as Routes - Stored in Routing Table.

* Routing Table - All networks a Router knows about.
* Have IP addresses in the Networks they are attacthed to 
* Gateway- Each hosts way out of their local network.
* Creates the Hierarchy in Networks and the entire internet.

- Routing is the process of moving data between networks.Router is a device whose primary purpose is Routing
- Switching is process of moving data within networks.Switch is a device who's primary purpose is Switching.

#### Other Network Devices

- Access Points
- Firewalls
- Load Balancers
- Layer 3 Switches
- IDS/IPS
- Proxies
- Virtual Switches
- Virtual Routers

All of above perform Routing and/or Switching

#### OSI Model

* Purpose Of Networking: Allow two hosts to share data with one another

* Hosts must follow a set of rules.

* Rules of Networking are divided into 7 layers (OSI)
    - Each layer serves a specific function

##### Layer 1 - Physical

* Computer data exists in form of Bits
* Something has to transport those bits b/w hosts
* L1 Technologies : Cables, Wifi, Repeates, Hubs

##### Layer 2 - Data Link - Hop to Hop

* Interacts with the Wire (i.e Physical Layer)
    - NIC / Wi-Fi access cards
* Hop to Hop (NIC to NIC)
* Addressing Scheme - MAC Address
    - 48 bits, represented as 12 hex digits.
    - Windows (94-65-9C-3B-8A-E5)
    - Linux (94:65:9C:3B:8A:E5)
    - CISCO (9465.9C3B.8AE5)
* L2 Technologies: NICS, Switches.

* This layer info is modified across the network by changing mac address

##### Layer 3 - Network - End to End

* Addressing Scheme - IP addresses
* L3 Technologies: Routers,Hosts (anything with IP address)

##### Layer 4 - Service to Service Delivery

* Distinguish data streams
* Addressing Scheme - Ports
    * 0-65535 - TCP - Favors Reliability
    * 0-65535 - UDP - Favors efficiency

* Servers listen for requests to pre-defined Ports.
* Clients select random port for each connection.
    - Response traffic will arrive on this port.
    - This also allows multiple requests for same server.

#### Layer 5,6,7 Session, Presentation, Application

* Distinction b/w these layers is somwhat vague.
* Other networking models combine these into one layer (ex: TCP/IP model)

Example:

* Client to Server
* Sending - Encapsulation.
* Transport Layer- Adds TCP header for each segment which contains the destination and source ports.
* Network Layer - Adds IP header for each packet which contains the ips of src and dst.
* Data Link Layer - Adds L2 header for each frame which contains the mac id's for hop to hop transfer.
* Physical Layers - Converts to 1 and 0's 


* Receiving  - De-Encapsulation
* Discards the each headers and transmits the data to application.

#### Rules

* Networking Devices operate at specific layers.

    * L2 devices only look into the datagram up to L2 header (Switches)
    * L3 Devices only look into the datagram up to L3 header (Routers)

* Network Protocols operate at specific layers

* Neither of above are strict rules - Exceptions exist. (Routers with ACL and ARP)
* OSI model is simply a model - not rigid rules everything adheres to
    - Conceptualization of what is required for data to flow through the internet.



### ARP - Address Resolution Protocol

* Links a L3 address to a L2 address.

### Host A and B are directly connected (Same Network - hubs/switches)

* Both hosts have a NIC, and therefore a MAC address.
* Both hosts are configured with an IP address and a Subnet Mask.
    - Subnet mask identifies the size of the IP network.
* Host A has some Data to send to the Host B
* Host A knows the IP address of Host B
    - ex: ping 10.1.1.33
    - Maybe IP address was acquired from DNS.
    - Host A knows a own ip address in its own IP network.
* Host A can create the L3 header to attach to the Data
* Host A does not know Host B's MAC address.
    - Host A mush use Address Resolution Protocol (ARP) to resolve targets MAC address.
    - ARP request includes sender's MAC address.
    - ARP request is a broadcast - sent to everyone on the network.
        * Destination MAC address: ffff.ffff.ffff
        * Reserved MAC address to send a packet to everyone on the local network.
    - ARP mappings are stored in an ARP cache.
        - When host B receives the broadcast message stores Host A mac address in its cache.
        - Host B responds by sending and ARP Response.
            - Response sent is Unicast (directly to Host A)
    - Host A populates it's ARP cache with Host B's IP/MAC mapping.
* Host A creates L2 header.
    - Layer 2 - Hop to Hop delivery.
* Data is sent to Host B
    * L2 header is discarded by Host B
    * L3 header is discarded by Host B
    * Data is processed by the Application.

* Host B has necessary information to respond
    - ARP cache is already populated.
* Further data exchange b/w hosts is simple.
    * Both hosts have what they need to create L3 and L2 headers.

### Host communication connected through a Router.(Foreign Network)

* Host A, Host C, and the Router have MAC and IP address.
* Host A has some data to send to Host C.
    - Host A knows Host C's Address is on a foreign network.
* Host A creates a L3 header
    - Layer 3 - End to End
* Host A needs to create a L2 header.
    * Layer 2 - Hop to Hop
    * Next hop is to router.
    * Host A uses ARP to reolve the MAC address of the Routers IP
        - Router IP address is configured as the Default Gateway.
    * Completes layer 2 header.
* Now data is sent to the router.
    * L2 header is discarded by router.
    * Router takes over from this point.
    * Host A's Job is done.
* ARP mapping generated by Host A can be used for ANY host in foreign networks.

* Note: Host A's first step when sending data is always the same.
    - Determine if targer IP is on local or foreign network.
        - Foreign: ARP for Default Gateway IP
        - Local: ARP for Target IP.

### Switches.

* Process of moving data within networks - Switching
* Switches primary purposes is Switching.
* Devices communicating through a switch belong to the same IP network.
* Switches are L2 devices - they only use L2 header to make decisions.
* Switches use and Maintain MAC Address Table
    * Mapping of Switch Ports to MAC addresses.
* Perform three actions
    - Learn
        - Update MAC address Table with mapping of Switch Port to Source MAC
    - Flood
        - Duplicate and send frame out all switch ports (except receiving port)
    - Forward
        - Use MAC address Table to deliver Frame to appropriate switch port.

* Process would be identical if Host was Router connected to the internet.

* Traffic going THROUGH the Switch
    - Switch has a MAC address

* Traffic going to Switch.
    - Switch has a MAC address and is configured with an IP address.
    - Switch essentially acts as a host in the network (follows all the host communication rules.)
    - Unicast Frame - destination MAC is another host.
        - Switch will flood only if MAC address is not in MAC address table.
    - Broadcast Frame - destination MAC address of ffff.ffff.ffff
        - Always Flooded

* Switches will only send Broadcasts if traffic is going TO or FROM the Switch.

#### VLANs

- Virtual Local Area Networks.
- Divides Switch Ports into multiple "mini-switches"
- Switches do all three actions within each VLAN.

#### Multiple Switches

- Switches maintain independent MAC address Tables.
- Switches perform switch actions independently.

### Routers

* Routers are connected to a network
    - Have an IP address and a MAC address.

* RFC 2460:
    - Internet Protocol, Version 6 Specification (IPv6)
    - Terminology
        - Node : a device that implements IPv6
        - Router : a node that forwards IPv6 packets not explicitly addressed to itself.
        - Host : any node that is not a router 
* Routers forward packets not destined to themselves (unlike hosts.)
* Routers maintain a map of all the networks they know about
    - Routing Table. (Used to send packets to a interface based on these)

* Routing Table can be populated via three methods:
    - Directly Connected - Routes for the Networks which are attached.
        * When Routers receive packets with an unknown Destination IP, packet is dropped
    - Static Routes - Routes manually provided by an Administrator.
        * When Routers receive packets with an unknown Destination IP, packet is dropped
    - Dynamic Routes - Routes learned automatically from other routers.
        * Routers learns routes from other Routers.
        * BGP, RIP OSPF, EIGRP, IS-IS etc..

* Routers also have ARP tables - Mapping of L3 to L2 address.
    * Everything with an IP address has an ARP table
    * Start Empty - populated as needed with network traffic.

* Example:
    - Host A has data to send to HOst C
    - Host A constructs the L3 header
    - Host A knows packet must be sent to Default Gateway (R1=Router 1) since it is a foreign network.
    - Host A cannot construct L2 header
        - No ARP entry for Gateway's IP address (R1)
        - Host A sends ARP request for Host C.
        - R1 populates its ARP table with entry for Host A
        - R1 sends ARP Response
        - Host A populates its ARP table with entry for its Gateway IP
    - R1 receives Packet.
    - R1 discards L2 header.
    - R1 looks up Destination IP in Routing Table.
        - Packets next hop is R2(Router 2)
    - R1 cannot construct L2 header
        - No ARP entry for R2
    - R1 sends ARP request for R2
    - R2 populates its ARP table with entry for R1
    - R2 sends ARP Response
    - R1 populates its ARP table with entry for R2
    - R1 constructs L2 header.
    - R1 sends packet to R2
    - R2 receives packet
    - R2 discards L2 header
    - R2 looks up Destination IP in Routing Table.
        - Packets final host is Host C
    - R2 sends ARP request for Host C
    * Host C populates its ARP Table with entry for R2
    - Host C sends RRP Response
    - R2 populates its ARP table with entry for Host C
    - R2 constructs L2 header.
    - R2 sends packet to Host C
    - Host C receives packet.
    - Host C discards L2 and L3 header
    - Host C processes the data. 
Note: Events b/w R2 and R1 would repeat for any amount of Routers in the Path.

* Each Router:
    - Looks up Destination IP in Routing Table to determine Next-Hop IP.
    * Adds a L2 header with Destination MAC next Router's MAC.
        - Performs ARP as necessary.

- Routers typically connected in a hierarchy.
    - Easier to scale - more consistent connectivity
    - Hierarchy allows for Route Summarization(/24 - CIDR subnetting)
        - Reduce number of Routes in Routing Table
        - Default Route - Ultimate Route summary
            - 0.0.0.0/0 - Every IPv4 address
            - For everything else, go here.

### Important Protocols

* Set of Rules and messages that form an Internet Standard - Protocol
    * ARP - Address Resolution Protocol - RFC 826
        - Resolves IP to MAC mappings.
        - ARP Requests / ARP Respones.

#### FTP - File Transfer Protocol.

* RETR file.pdf (Retrieve)


#### SMTP - Simple Mail Transfer Protocol.

* Client sends HELO client.com(smtp server)
* SMTP server reponses with 250 response code

#### HTTP - Hyper Text Transfer Protocol

#### SSL - Secure Sockets Layer

To establish a secure tunnel for communication

#### TLS - Transport Layer Security

To establish a secure tunnel for communication

#### HTTPS secured with SSL/TLS.

#### DNS - Domain Name System

- Converts Domain Names into IP addresses.
- Makes email service and http service possible.

- Every host needs four items for internet connectivity
    - IP address - Host identity on the internet.
    - Subnet mask - Sixe of Hosts network.
    - Default Gateway - Routers IP address.
    - DNS Server IP(s) - To translate domain names to IPs

#### DHCP - Dynamic Host Configuration Protocol.
    * DHCP server provides IP/SM/DG/DNS for clients.

### Important Info

* Data moves through networks based upon three Tables;
    - MAC Address Table - Mapping of Switchport to MAC address.
    - ARP Table/Cache - Mapping of IP address to MAC address.
    - Routing Table - Mapping of IP network to interface or Next Router.
    
#### Credits

[Networking Fundamentals](https://www.youtube.com/watch?v=bj-Yfakjllc&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi)