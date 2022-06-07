# Computer Networks

A computer network is  system of peripherals or computers interconnected with each other and has a standard communication channel established between them to exchange different types of information and data.

* **Node:** Any communicating device in a network.
* **Link:**  A link or edge refers to the connectivity between two nodes in the network. It includes the type of connectivity(wired or wireless) between the nodes and protocols used for one node to able to communicate with the other.

#### Network Topology

Physical layout of the network, connecting the different nodes using the links. It depicts the connectivity between the computers, devices, cables etc.

* Bus Topology

    - All the nodes are connected using the central link known as the bus.
    - It is useful to connect a smaller number of devices.
    - If the main cable gets damaged, it will damage the whole network.

* Star Topology

    - All the nodes are connected to one single node known as the central node.
    - If the central node fails the complete network is damagaed.
    - Easy to troubleshoot.
    - Mainly used in home and office networks.

* Ring Topology.

    - Each node is connected to exactly two nodes forming a ring structure.
    - If one of the nodes are damaged, it will damage the whole network.
    - It is used very rarely used as it is expensive and hard to install and manage.

* Mesh Topology

    - Each node is connected to one or many nodes.
    - It is robust as failure in one link only disconnects that node
    - It is rarely used and installation and management are difficult

* Tree Topology

    - A combination of star and bus topology also know as an extended bus topology.
    - All the smaller star networks are connected to a single bus.
    - If main bus fails, the whole network is damaged.

* Hybrid.

    - It is a combination of different topologies to form a new topology.


##### Network Types.

Classified based on the area of distribution of the network.

- Personal Area Network
- Local Area Network
- Metropolitan Area Network
- Wide Area Network
- Internet (Global Area Network)

### DNS

* Domain Name System. Default Port 53. 
* Translates the domain names to their corresponding IPs.


### VPN

* Virtual Private Network is a private WAN built on internet.
* It allows the creation of a secured tunnel(protected network) between different networks using the internet(public network).
* Using VPN, a client can connect to the organization's network remotely.

#### Advantages

* Used to connect offices in different geographical locations remotely and is cheaper when compared to WAN connections.
* Keeps an organization's information secured against any potential threats or intrusions by using virtualization.
* Encrypts the internet traffic and disguises the online identity.


### Layers of OSI

* OSI - **Open Systems Interconnection**
* 7 Layers - Work collaboratively to transmit the data from one person to another across the globe.
* Acts as a reference model
* Layers

    - Application Layer     - Software Layer
    - Presentation Layer    - Software Layer
    - Session Layer         - Software Layer
    - Transport Layer       - Heart of OSI
    - Network Layer         - Hardware Layer
    - Data Link Layer       - Hardware Layer
    - Physical Layer        - Hardware Layer


#### Physical Layer

* Lowest layer
* Reponsible for the actual physical connection between the devices.
* Contains information in the forms of bits.
* Responsible for transmitting individual bits from one node to the next.
* When receiving data, this layer will get the signal received and convert it into 0s and 1s and send them to the Data Link Layer, which will put the frame back together.

##### Functions of Physical Layer

- **Bit Synchronization:** Provides clock which controls both sender and receiver thus providing the synchronization at bit level.
- **Bit rate control:** Defines the transmission rate i.e no. of bits sent per second.
- **Physical Topologies:** Specifies the way in which the different, devices/nodes are arranged in a network i.e bus, star or mesh topology.
- **Transmission mode:** Defines the way in which data flows between the two connected devices.(Simplex, half-duplex, and full-duplex)


#### Data Link Layer

* Responsible for node-to-node delivery of the message.Uses MAC address to transmit. 
* Main function of this layer is to make sure data transfer is error-free from one node to another over the physical layer.

##### Sublayers

- Logical Link Control (LLC)
- Media Access Control (MAC)

##### Functions

- **Framing:** provides a way for a sender to transmit a set of bits that are meaningful to the receiver. This can be accomplished by attaching special bit patterns to the beginning and end of the frame.
- **Physical Addressing:** After creating frames, The DLL adds physical Address(MAC address) of the sender and/or receiver in the header of each frame.
- **Error Control:** Provides mechanism of error control in which it detects and retransmits damaged or lost frames
- **Flow Control:** The data rate must be constant on both sides else data may get corrupted thus, flow control coordinates the amount of data that can be sent before receiving acknowledgement.
- **Access Control:** When a single communication channel is shared by multiple devices, the MAC sub-layer of the data link layer helps to determine which device has control over the channel at a given time.

#### Network Layer

* Works for the transmission of data from one host to the other located in different networks.
* It also takes care of packet routing i.e. selection of the shortest path to transmit the packet, from the number of routes available.
* The sender and receiver's IP addresses are placed in the header by the network layer.

##### Functions 

* **Routing:** Determines which route is suitable from source to destination.
* **Logical Addressing:** Defines the addressing scheme, to identify each device on the internet/network work uniquely.The sender and receiver's IP addresses are placed in the header by the network layer.

#### Transport Layer

Provides services to the application layer and takes services from the network layer. 
The data in transport layer is referred as Segments.
Responsible for End to End delivery of the complete message.
Provides the acknowledgement of successful data transmission and re-transmits the data if an error is found.

**At Senders side:**

Receives formatted data from the upper layers, performs Segmentation.
Adds Source and Destination port numbers in its header and forwards the segmented data to the Network Layer

**At receiver's side:**

Reads the port number from its header and forwards the data which it has received to the respective application.
Performs sequencing and reassembling of the segmented data.

##### Functions

**Segmentation and Reassembly:** Accepts message from session layer, breaks the message into smaller units and adds header associated with it. At destination, it reassembles the message.

**Service Point Addressing:** Includes service point address or port address. This make sure the message is delivered to correct process.

##### Services

* Connection-Oriented Service (three phase process)

    - Connection Establishment
    - Data Transfer
    - Termination/ Disconnection

Here, the receiving device sends an acknowledgement, back to the source after a packet or group of packets of packets are received. This type of transmission is reliable and secure.

* Connectionless Service

One-Phase process and includes data transfer. No receiptant of acknowledgement of packets. 
Allows for much faster communication between devices.

### Session Layer

Responsible for the establishment of connection, maintainance of sessions, authentication, and also ensures security.

Also responsible for ensuring data is synchronized and consistent before transmission. Example: Streaming of multimedia audio and video

### Presentation Layer

The data from the application layer is extracted here and manipulated as per the required format to transmit over the network.

#### Functions

* **Translation** example: ASCII to EBCDIC
* **Encryption / Decryption**
* **Compression**

### Application Layer

Implemented by the network applications.

These applications produce the data, which has to be transferred over the network.

This layer serves as a window for the application services to access the network and for displaying the received information to the user.

### TCP/IP Model

* Stands for Transmission Control Protocol / Internet Protocol.

* Layers

    - Process/Application Layer
    - Host-to-Host/ Transport Layer
    - Internet Layer
    - Network Access/ Link Layer

#### Network Access Layer

Corresponds to the combination of Data Link Layer and Physical Layer of OSI

#### Internet Layer

Parallels OSI's Network layer.

Responsible for logical transmission of data over the entire network.

##### Main Protocols 

- **IP:** Internet Protocol responsible for delivering packets from the source host to the destination host by looking at the IP addresses in the packet headers. IP has 2 versions - IPV4 and IPV6

- **ICMP:** Internet Control Message Protocol. It is encapsulated within IP datagrams and is responsible for providing hosts with information about network problems.Uses Port 7 by default.

- **ARP:** Address Resolution Protocol. Its job is to find the hardware address of a host from a known IP address. ARP has several types: Reverse ARP, Proxy ARP, Gratuitous ARP and Inverse ARP.

### Transport or Host-to-Host Layer

Responsible for end-to-end communication and error-free delivery of data. 

#### Protocols

- **Transmission Control Protocol(TCP):** 
    
    * Provide reliable and error-free communication between end systems. 
    * It performs sequencing and segmentation of data.
    * It also has acknowledgment feature and controls the flow of the data through flow control mechanism.
    * Protocols like HTTP, FTP, Telnet, SMTP, HTTPS, etc use TCP at the transport layer.

- **User Datagram Protocol (UDP)**

    * Connectionles
    * For application that does not require reliable transport as it is very cost-effective since there are no overheads
    * Mainly used for multicasting and broadcasting.
    * Protocols like DNS, RIP(Routing Information Protocol), SNMP(Simple Network Management Protocol), RTP, NIP etc use UDP at transport layer.

#### Application Layer

* Performs functions of top three layers of OSI Model: Application, Presentation and Session Layer
* Responsible for node-to-node communication and controls user-interface specifications

#### Protocols

* **HTTP and HTTPS:** Hypertext transfer protocol. Defines the set of rules and standards on how the information can be transmitted on the WWW. It is a stateless protocol. HTTPS is combination of HTTP with SSL.
* **SSH:** Secure Shell. Its sets up a secure session over TCP/IP connection. Encrypted Connection
* **NTP:** Network Time Protocol. Used to synchronize the clocks on our computer to one standard time source.


### OSI vs TCP/IP

* 7 Layered architecture                            | 4 Layered architecture.
* Fixed boundaries and functionality for each Layer | Flexible architecture with no strict boundaries between layers.
* Low Reliability                                   | High Reliability
* Vertical Layer Approach                           | Horizontal layer Approach

### IP Network

* An IP network is a communication network that uses Internet Protocol (IP) to send and receive messages between one or more computers.
* An IP network requires that all hosts or network nodes be configured with the TCP/IP suite.
* An IP network requires that all connected devices - such as servers, switches, routers and other devices - be configured with the TCP/IP suite and have a valid IP address to perform any network communication.

### IP security (IPSec)

- The IP security (IPSec) is an Internet Engineering Task Force (IETF) standard suite of protocols between 2 communication points across the IP network that provide data authentication, integrity, and confidentiality. 
- It also defines the encrypted, decrypted and authenticated packets. The protocols needed for **secure key exchange** and **key management** are defined in it.

#### Uses of IPSec

* To encrypt application layer data.
* To provide security for routers sending routing data across the public internet.
* To provide authentication without encryption, like to authenticate that the data originates from a known sender.
* To protect network data by setting up circuits using IPsec tunneling in which all data is being sent between the two endpoints is encrypted, as with a Virtual Private Network(VPN) connection.

#### Components of IP Security 

* **Encapsulating Security Payload(ESP)**

    It provides data integrity, encryption, authentication and anti replay. It also provides authentication for payload.

* **Authentication Header(AH)**

    It also provides data integrity, authentication, and anti replay and it does not provide encryption. The anti replay protection, protects against unauthorized transmission of packets. It does not protet data's confidentiality.

* **Internet Key Exchange (IKE)**

    It is a network security protocol designed to dynamically exchange encryption keys and find a way over Security Association(SA) between 2 devices.

Security Association (SA) establishes shared security attributes between 2 network entities to support secure communication.

The Key Management Protocol (KMP) and Internet Security Association (ISA) provides a framework for authentication and key exchange. These tells how the set up of the Security Associations (SAs) and how direct connections between two hosts that are using IPSec.

Provides message content protection and also an open frame for implementing standard algorithms such as SHA and MD5.

The algorithms IPSec produces a unique identifier for each packet.This identifier then allows a device to determine whether a packet has been correct or not. Packets which are not authorized are discarded and not given to receiver.

#### Working of IP Security

1. The host checks if the packet should be transmitted using IPSec or not. These packet traffic triggers the security policy themselves. This is done when the system sending the packet apply an appropriate encryption. The incoming packets are also checked by the host that are encrypted properly or not.
2. Then the **IKE Phase 1** starts in which 2 hosts using IPSec authenticate themselves to each other to start a secure channel. It has 2 modes. The **Main Mode** which provides the greater security and the **Aggressive mode** which enables the host to establish an IPSec circuit more quickly.
3. The channel created in the last step is then used to securely negotiate the way the IP circuit will encrypt data across the IP circuit.
4. Now, the **IKE Phase 2** is conducted over the secure channel in which the two hosts negotiate the type of cryptographic algorithms to use on the session and agreeing on secret keying material to be used with those algorithms,
5. Then the data is exchanged across the newly created IPSec encrypted tunnel. These packets are encrypted and decrypted by the hosts using IPSec SAs.
6. When the communication between the hosts is completed or the session times out then the IPSec tunnel is terminated by discarding the keys by both the hosts.

### TLS Handshake.

#### Cryptography

* Method of using advanced mathematical principles in storing and transmitting data in a particular form so that only whom it is intended can read and process it.

#### Encryption

* Process whereby a message is encoded in a format that cannot be read or understood by an eavesdropper.

#### Data Integrity

* Hashing is used to provide Integrity.
    * General Assumptions.
        - Sender calculates Digest from the Message.
        - Sender sends Message and Digest.
        - Receiver calculates Digest from received Message.
        - Receiver compates both digests
            - If digests are identical, the message was not modified in transit.

        - This may result in Man in the middle attackers, here the hacker/attacker changes both digest and message.

    * Actual
        - Both parties establish a mutual Secret Key.
        - Sender combines Message + Secret Key to create Digest.
        - Sender sends the message and digest but not the key.
        - Receiver verifies by calculating hash of Message + Secret Key
            - Message was not modified in transit
            - Sender had the identical Secret Key.

#### Symmetric Encryption

* Same Key on both sides.
* Technique where same key is used to both encrypt and decrypt the data.
* Blowfish, AES, DES are examples 
* AES-128, AES-198 and AES-256 widely used.

* Disdavantage: All parties involved have to exchange the key used to encrypt the data before they can decrypt it.

* Example 1

    * Value: 5
    * Secret Key: 3
    * Encrypted value: 15 (5*3)

* Example 2

    * Value - hello
    * Secret Key - 3 (which shifts the original value 3 times)
    * Encrypted Value- khoor

#### Asymmetric Encryption

* Asymmetric encryption uses a mathematically related pair of keys for encryption and decryption: Public Key and Private Key.
* A message that is encrypted using a public key can only be decrypted using a private key while a message encrypted using a private key can be decrypted using a public key(For integrity).
* Keys are mathimatically linked 

* RSA, DSA, PKCS


* Simple Example

    - Encrytion
        - Value - hello
        - Encryption Key = 5
        - Encrypted Value - mjqqt
    - Decrytion
        - Value - mjqqt
        - Decryption key = 21
        - Decrypted Value = Hello



#### Message Authentication Code.

* Concept combining Message + Secret Key when calculating Digest.
* Provides Integrity and Authentication for Bulk Data transfer.
* Message + Secret Key must be combined in the same way
* Industry Standard: HMAC
    - Hash based Message Authentication Code (RFC 2104)


#### SSL/TLS have 3 goals 

    * Confidentiality  - Provided by Symmetric Encryption
    * Integrity (Not to tamper) - Provided by M.A.C (Hashing) - Message Authentication Code
    * Authentication - Provided by Certificates/PKI

* Symmetric Encryption / M.A.C Require Secret Keys

    * Provided by Key Exchange (Asymmetric Encryption)

* Cipher Suite picks four protocols.

    - Key Exchange
    - Authentication
    - Encryption
    - Hashing

#### Handshake

* Client Hello

    - Version

        - Highest version of TLS/SSL client supports.

    - Random Number - 32 bytes / 256 bits

        - Timetamp encoded in first four bytes

    - Session ID - 8 bytes / 64 bits

        - 0000... All 0's in initial Client Hello

    - Cipher Suites

        - List of Cipher Suites Client supports.

    - Extensions 

        - Optional additional features added to TLS/SSL. 

* Server Hello

    * Same info as client that server supports.

* Server sends the Certificate Chanin to the Client (Public Key)
* Server sends Hello done.
* Client Key Exchange.

    - Establish Mutual Keying Material (ie SEED Value)
    - Proves Server is true owner of Certificate.

    - Client Generates Pre-Master-Secret
        - 2 Bytes - TLS/SSL version
        - 46 bytes - Random 

        - Encrypted with Server's Public Key
            - Server can only decrypt if it has server Private Key.
        - Pre-Master-Secret is used to generate derive Master Secret.
            - PreMasterSecret
            - "master secret"
            - Client Random
            - Server Random
        - Master Secret is used to generate Session Keys.
            - Master Secret
            - "key expansion"
            - Client Random
            - Server Random

        - Session Keys (Symmetric Keys)

            - Client Encryption Key
            - Client HMAC Key

            - Server Encryption Key
            - Server HMAC Key

    - SSL generates 2 secure tunnels which are unidirectional. If client wants to send data to server it uses client keys for confidentiality and integrity.

    - Caliculations involve a PRF - Pseudo Random Function
        - Hashing algorithm that generates the digests at any desired length.


Note: At his point, both parties have identical session keys. But client/server don't know whether the other has the same keys.

- Client sends Change Cipher Spec
    - Indicates Client has everything necessary to speak securely.

- Client sends Handshake: Finished.

    - Proves to the Server that client has correct Session Keys.
    - Client calulates Hash of all handshake Records seen so far.Uses the below to generate Handshake Hash
        - Client Hello
        - Server Hello
        - Certificate
        - Server Hello Done
        - Client Key Exchange
    - Using all the info such as Client Hello to generate Handshake Hash prevents the man-in-the-middle attack.
    - Combined with other values client creates Verification Data using PRF
        - Master Secret
        - "client finished"
        - Handshake Hash
    - Verification Data is encrypted with Client Session Keys.
    - Encrypted Verification data is sent to the server.

- After verifying server sends the Change Cipher Spec

    - Indicates Server has everything necessary to speak securely.

- Handshake: Finished

    - Proves to the Client that Server has correct Session Keys.
    - Client calulates Hash of all handshake Records seen so far.Uses the below to generate Handshake Hash
        - Client Hello
        - Server Hello
        - Certificate
        - Server Hello Done
        - Client Key Exchange
        - CLient Finished
    - Combined with other values server creates Verification Data using PRF
        - Master Secret
        - "server finished"
        - Handshake Hash
    - Verification Data is encrypted with Server Session Keys.
    - Encrypted Verification data is sent to the Client.

- Now application data is sent using the session keys.


### Secure Socket Layer (SSL)

Provides security to the data that is transferred between web browser and server. SSL encrypts the link between a web server and a browser which ensures that all data passed between them remain private and free from attack.

SSL is designed to make use of TCP to provide reliable end-to-end secure service.


#### SSL Protocols

* SSL record protocol.


    SSL Record provides 2 sevices for SSL connection.

    - Confidentiality
    - Message Integrity.


    In the SSL Record Protocol application data is divided into fragments. The fragment is compressed and then encrypted MAC (Message Authentication Code) generated by algorithms like SHA (Secure Hash Protocol) and MD5 (Message Digest) is appended. After that encryption of the data is done and in last SSL header is appended to the data.

* Handshake protocol.

     Used to establish sessions. This protocol allows the client and server to authenticate each other by sending series of messages to each other. Uses four phases to complete its cycle.

     - **Phase-1**: both client and server send hello-packets to each other, In this IP session, cipher suite and protocol version are exchanged for security purposes.
     - **Phase-2**: Server sends his certificate and Server-key exchange. The server end phase-2 by sending the Server hello-end packet.
     - **Phase-3**: Client replies to the server by sending his ceritificate and Client-exchange-key.
     - **Phase-4**: Change-cipher suite occured and after this handshake protocol ends.

* Change-cipher spec protocol.

    This protocol uses the SSL record protocol. Unless Handshake Protocol is completed, the SSL record Output will be in a pending state.After handshake protocol, the Pending state is converted into the current state.
    Change-cipher protocol consists of a single message which is 1 byte in length and can have only one value. This protocol's purpose is to cause the pending state to be copied into the current state.

* Alert Protocol

    Used to convey SSL-related alerts to the peer entity. Each message in this protocol contains 2 bytes. - Level(1 byte) and Alert (1 byte).

    The level is further classified into two parts.

    **Warning (level = 1)**

    This alert has no impact on the connection between sender and receiver. Some of them are:

    Bad certificate: When received certificate is corrupt.
    No Certificate: When an appropriate certificate is not available.
    Certificate expired: When a certificate has expired.
    Certificate unknown: When some other unspecified issue arose in processing the certificate, rendering it unacceptable.
    Close notify: It notifies that the sender will no longer send any messages in the connection.

    **Fatal Error (level = 2)**

    This alert breaks the connection between sender and receiver. Some of them are:

    Handshake failure: When the sender is unable to negotiate an acceptable set of security parameters given the options available.
    Decompression failure: When the decompression function receives improper input.
    Illegal parameters: When a field is out of range or inconsistent with other fields.
    Bad record MAC: When an incorrect MAC was received.
    Unexpected message: When an inappropriate message is received.

    The second byte in the Alert protocol describes the error.

### Transport Layer Security (TLS)

Designed to provide security at transport layer. Derived from SSL.
It ensures that no third party may eavesdrop or tampers with any message.

#### Benefits of TLS

* Encryption

    TLS/SSL can help to secure transmitted data using encryption.

* Interoperability

    TLS/SSL works with most web browsers and on most operating systems and web servers

* Algorithm flexibility

    Provides operations for authentication mechanism, encryption algorithms and hashing algorithm that are used during the secure session.

* Ease of Deployment.

* Ease of Use.

    Because we implement TLS/SSL beneath the application layer, most of its operations are completely invisible to client.


#### Working of TLS

The client connect to server using TCP. The clients sends the following.

* Version of SSL/TLS
* which cipher suites, compression method it wants to use.

The server checks what the highest SSL/TLS version that is supported by them both, picks a cipher suite from one of the client option and optionally picks a compression method. After this basic setup is done, the server provides its certificate. This certificate must be trusted either by the client itself or a party that the client trusts.  Having verified the certificate and being certain this server really is who he claims to be (and not a man in the middle), a key is exchanged. This can be a public key, "PreMasterSecret" or simply nothing depending upon cipher suite.

Both the server and client can now compute the for symmetric encryption. The handshake is finished and the two hosts can communicate securely. 

### SSL VS TLS

The main differences between Secure Socket Layer and Transport Layer Security is that. In SSL (Secure Socket Layer), **Message digest** is used to create master secret and It provides the basic security services which are Authentication and confidentiality. while In TLS (Transport Layer Security), **Pseudo-random function** is used to create master secret.

*  TLS is the successor of the SSL protocol. 

* SSL is a cryptographic protocol that uses explicit connections to establish secure communication between web server and client
* TLS is also a cryptographic protocol that provides secure communication between web server and client via implicit connections. It’s the successor of SSL protocol. 
* SSL (Secure Socket Layer) is less secured as compared to TLS(Transport Layer Security)
* TLS (Transport Layer Security) is just an updated, more secure, version of SSL


### Classful IP Addressing

The 32 bit IP address is divided into 5 sub-classes. These are:

* Class A - Network Bits - 8 Bits
* Class B - Network Bits - 16 Bits
* Class C - Network Bits - 24 Bits
* Class D - Multicast Address
* Class E - Reserved

Class D and E are reserved for the multicast and experimental purposes respectively.

IPV4 address is divided into 2 parts.

- Network ID
- Host ID

IP addresses are globally managed by Internet Assigned Numbers Authority (IANA) and regional internet registries (RIR)

Note: While finding the total no. of host IP addresses, 2 IP addresses are not counted and are therefore, decreased from the total count because the first IP address of any network is network number and where as last last IP address is reserved for broadcast IP.

#### Class A

These are assinged to the networks that contain a large number of hosts.

The higher order bit of the first octet in Class A is always set to 0.The remaining 7 bits are used to determine netword ID. Rest 24 bits Host ID

* 2^7 - 2 = 126 Network ID (0.0.0.0 and 127.x.y.z are special address)
* 2^24 - 2 = 16777214 Host ID

Default Subnet Mask - 255.x.x.x

IP Addresses Range : **1.x.x.x - 126.x.x.x**

#### Class B

These are assigned to networks that ranges from medium-sized to large-sized networks.

The higher order bits of the first octet of IP addresses of class B are always set to 10. The remaining 14 bits are used to determine network ID

Default Subnet Mask - 255.255.x.x

* 2^14 = 16384 Network Address
* 2^16 - 2 = 65534 host address.

IP Address Range: **128.0.x.x - 191.255.x.x**

#### Class C

These are assigned to small-sized networks.

The higher order bits of the first octet of IP addresses of class C are always set to 110. The remaining 21 bits are used to determine network ID

Default Subnet Mask - 255.255.255.x

* 2^21 = 2097152 network address
* 2^8 - 2 = 254 host address

IP Addresses Range: **192.0.0.x - 223.255.255.x**

#### Class D

Reserved for multi-casting.

The higher order bits of the first octet of IP addresses are always set to 1110. The remaining bits are for that interested hosts.

No Subnet-mask.

IP Range: **224.0.0.0-239.255.255.255**

#### Class E

Reserved for experimental and research purposes.

IP Range: **240.0.0.0-255.255.255.254**

Higher order bits of first octet of class E are always set to 1111.


#### Range of special IP addresses

169.254.0.0 - 169.254.0.16 : Link Local Addresses

127.0.0.0 - 127.0.0.8 : Loop-back addresses

0.0.0.0 - 0.0.0.8 : used to communicate within the current network.

#### Problems with Classful Addressing.

Millions of class A address are wasted.
Many of the class B addresses are wasted 
No. of addresses available in class C is so small that it cannot cater the needs of the organization. 
Class D addresses are used for multicast routing and therefore available as a single block only.
Class E are reserved.

**Because of above problems, Classful networking was replaced by Classless Inter-Domain Routing(CIDR) in 1993**

#### Important Info

- **Address:** The unique number ID assigned to one host or interface in a network.
- **Subnet:** A portion of a network that shares a particular subnet address.
- **Subnet mask:** A 32-bit combination used to describe which portion of an address refers to the subnet and which part refers to the host.
- **Interface:** A network connection.

#### Private Address

For each class, There are specific IPs that are reserved specifically for private use only. This IP address cannot be used for devices on the Internet as thery are non-routable.

```
IPv4 Class 	Private IPv4 Start Address 	Private IPv4 End Address
A 	        10.0.0.0 	                10.255.255.255
B 	        172.16.0.0 	                172.31.255.255
C 	        192.168.0.0 	            192.168.255.255
```

### Subnetting

* Allows us to create multiple logical networks in a single network.
* Any device or gateway that connects n networks/subnetworks has n distinct ip addresses, one for each network/subnetwork that it interconnects

In order to subnet a network, extend the natural mask with some of the bits from the host id portion of address in order to create a subnetwork ID.

Example: By extending the mask to be 255.255.255.224, we have taken 3 bits from original host portion of the address and used them to make subnets.With these 3 bits , it is possible to create 8 subnets. With remaining five host id bits, each subnet can have upto 32 host addresses, 30 of which can be actually assigned to a device, since host ids of all zeros is network id and all ones is broadcast address in that particular network.

Note: 204.17.5.32/27 denotes the network 204.17.5.32 with subnet mask 255.255.255.224.This is CIDR notation.

#### Determining the subnet of given IPs

DeviceA : 172.16.17.30/20

DeviceB : 172.16.28.15/20

Determining Subnet of DeviceA:

```
172.16.17.30  -   10101100.00010000.00010001.00011110
255.255.240.0 -   11111111.11111111.11110000.00000000
                  -----------------| sub|------------
subnet =          10101100.00010000.00010000.00000000 = 172.16.16.0
```

Determine the Subnet for DeviceB:

```
172.16.28.15  -   10101100.00010000.00011100.00001111
255.255.240.0 -   11111111.11111111.11110000.00000000
                  -----------------| sub|------------
subnet =          10101100.00010000.00010000.00000000 = 172.16.16.0
```
From these determinations, DeviceA and DeviceB have addresses that are part of the same subnet.

#### Creation of Subnets for a network

Given the Class C network of 204.15.5.0/24, create below subnets with required hosts for each subnet

netA: 14 hosts

netB: 28 hosts

netC: 2 hosts

netD: 7 hosts

netE: 28 hosts

Largest subnet must support 28 host addresses. Also In order to create 5 subnets, we need to use three bits from tthe Class C host bits which will leave 5 bits for the host portion of the address.with  5 bits, no of hosts each subnet support are 2^5 = 32 (30 usable). This meets our requirement.

An example of how we might assign the subnetworks is:

netA: 204.15.5.0/27      host address range 1 to 30

netB: 204.15.5.32/27     host address range 33 to 62

netC: 204.15.5.64/27     host address range 65 to 94

netD: 204.15.5.96/27     host address range 97 to 126

netE: 204.15.5.128/27    host address range 129 to 158

#### VLSM

* **Variable Length Subnet Masks.**
* Having same subnet masks for all the subnets ends up wasting the address space, in above example we have split the network into eight equal-size subnets; however, each subnet did not utilize all available host addresses, which resultant in wasted address space.

In above unsed subnets 

204.15.5.160/27

204.15.5.192/27

204.15.5.224/27

And also NetA, NetB, NetC have a lot of unused host address space.
VLSM allows us to use different masks for each subnet, thereby using address space efficiently.

The following masks are required by above subnetworks

netA: requires a /28 (255.255.255.240) mask to support 14 hosts

netB: requires a /27 (255.255.255.224) mask to support 28 hosts

netC: requires a /30 (255.255.255.252) mask to support 2 hosts

netD*: requires a /28 (255.255.255.240) mask to support 7 hosts

netE: requires a /27 (255.255.255.224) mask to support 28 hosts


Note:  a /29 (255.255.255.248) would only allow 6 usable host addresses therefore netD requires a /28 mask.

The easiest way to assign the subnets is to assign the largest first. For example, we can assign in this manner:

netB: 204.15.5.0/27  host address range 1 to 30

netE: 204.15.5.32/27 host address range 33 to 62

netA: 204.15.5.64/28 host address range 65 to 78

netD: 204.15.5.80/28 host address range 81 to 94

netC: 204.15.5.96/30 host address range 97 to 98

#### Special subnets

##### 31 Bit subnets.

A 30-bit subnet mask allows for four IPv4 addresses: two host addresses, one all-zeros networks, and one all-ones broadcast address. A point-to-point link can only have two host addresses. There is not real need to have the broadcast and all-zeros addresses with point-to-point links. A 31 bit submask will allow for exactly two host addresses, and eliminates the broadcastt and all-zeros addresses, thus conserving the use of IP addresses to the minimum for point-to-point links.

Using 31-Bit Prefixes on IPv4 Point-to-Point Links.The mask is 255.255.255.254 or /31.

##### 32-bit Subnets

A subnet mask of 255.255.255.255 (a /32 subnet) describes a subnet with only one IPv4 host address. These subnets cannot be used for assigning address to network links, because they always need more than one address per link. The use of /32 is strictly reserved for use on links that can have only one address. The example for Cisco routers is the loopback interface. These interfaces are internal interfaces and do not connect to other devices. As such, they can have a /32 subnet.


### CIDR

* Classless Interdomain Routing
* In CIDR, an IP network is represented by a prefix, which is an IP address and some indication of the length of the mask.


### Use of a router and how it is different from a gateway?

#### Router

- Network device that forwards packets from one network to another. Based on internal routing tables, routers read each incoming packet and decide how to forward it. Routers work at the network layer of the protocol.

- Primary Goal : Route traffic from one network to other.
- Feature Support : Provide additional features like DHCP server, NAT, Static Routing and Wireless Networking/IPv6 Address, Mac Address
- Hosted On: Dedicated Appliance (Router Hardware)
- Working Principle: Works by installing routing information for various networks and routes traffic based on destination address.



#### Gateway

- Device that converts one protocol or format to another. A network gateway converts packets from one protocol to another. The gateway functions as an entry/exit point to the network.
- Primary Goal : Translate from one protocol to other.
- Protocol Conversion like VoIP to PSTN(Public Switched Telephone Network) or Network Access Control etc.
- Dedicated/Virtual Appliance or physical Server
- Related Terms: Proxy Server, Gateway Router, Voice Gateway.
- Working Principle: Works by differentiating what is inside network and what is outside network.

A gateway can be implemented as a router or switch using software, hardware or a combination of both, depending on the types of protocols to be used across the network.


#### Switch

- Connect devices within a single network, transfer incoming and outgoing internet traffic between connected devices.While a router is designed to connect across multiple computer networks. 
- Can be used in wired network only.
- Maintains MAC address in a lookup table.
- Switches are mostly used in large computer networks connecting many computers within a single local area network. Depending on the size of a network, multiple switches may be required to connect different groups of computer networks.

To Provide internet access to the connected networks , switches are further connected to a router. The router has the advantage of getting internet access by using a single modem for all the devices in the network.

### DHCP

* Dynamic Host Configuration Protocol.
* It is an application layer protocol used to auto-configure devices on IP networks enabling them to use the TCP and UDP-based protocols. The DHCP server auto-assigns the IPs and other network configurations to devices individually which enables them to communicate over the IP network. It uses port 67 by default.

### MAC Address

* Media Access Control Address.
* It is a 48 bit or 64 bit unique identifier
* Used at Data Link Layer

### Hub vs Switch

```
Hub                                         |   Switch 

Operates at Physical Layer                  |   Operates at Data Link Layer
Half-Duplex transmission Mode               |   Full duplex transmission mode.
No software support for the administration. |   LAN devices can be connected
Less speeds up to 100 MBPS                  |   Supports high speed in GBPS.

```

### ipconfig and ifconfig

Used to get the TCP/IP summarry and allows to change the DHCP and DNS settings

* **ipconfig:** Internet Protocol Configuration. Used in windows to view and configure network interfaces.
* **ifconfig:** Interface Configuration. Used in MAC and Linux to view and configure network interfaces.

### Firewall

* Network security system that is used to monitor the incoming and outgoing traffic and blocks the same based on the firewall security policies. Its acts as a wall between the internet(public network) and network devices(a private network). It is either a hardware device, software program, or a combination of both. It adds a layer of security to the network.

### Unicasting, Anycasting, Multicasting and Broadcasting

* **Unicasting:** If the message is sent to a single node from the source then it is known as unicasting. This is commonly used in networks to establish a new connection.
* **Anycasting:** If the message is sent to any of the nodes from the source then it is known as anycasting. It is mainly used to get the content from any of the servers in the Content Delivery System.
* **Multicasting:** If the message is sent to a subset of nodes from the source then it is known as multicasting. Used to send the same data to multiple receivers. 
* **Broadcasting:** If the message is sent to all the nodes in a network from a source then it is known as broadcasting. DHCP and ARP in the local network use broadcasting.

### What happens when you request a url in the web browser?

Below are the steps that are being followed:

* Check the browser cache first if the content is fresh and present in cache display the same.
* If not, the browser checks if the IP of the URL is present in the cache (browser and OS) if not then request the OS to do a DNS lookup using UDP to get the corresponding IP address of the URL from the DNS server to establish a new TCP connection.
* A new TCP connection is set between the browser and the server using three-way handshaking.
* An HTTP request is sent to the server using the TCP connection.
* The web servers running on the Servers handle the incoming HTTP request and send the HTTP response.
* The browser process the HTTP response sent by the server and may close the TCP connection or reuse the same for future requests.
* If the response data is cacheable then browsers cache the same.
* Browser decodes the response and renders the content.

#### In Layers of OSI or TCP/IP

We don't use OSI in real work networks. We use TCP/IP Network Model.

* We type a web address in web browser. Web browser uses HTTP which is an **Application Layer Protocol**
* Behind the scenes the web browser gets the IP address of the URL using Domain Name System which is also an **Application Layer Protocol**
* Now the browser creates a HTTP packet and gives this packet to Transport Layer Protocol.Now this layer establishes a session creating a TCP connection. Now TCP some info required to maintain session/pipe to HTTP packet.
* Now TCP handovers the TCP packet to an Internet Protocol Layer. Its main job is addressing and routing. Now IP puts its own info on top of TCP packet which is required for routing in the internet
* Now IP handovers packet to Network interface Layer, which defines the protocols and hardware required to deliver data across some physical network. Most of the PC uses Ethernet.Here the IP packet is encapsulated with Ethernet header and Ethernet Trailer, creating Ethernet Frame. IT contains some thing called MAC address which is used to send frame locally(LAN). It transmits data i.e data leaves the PC and eventually reaches the Server.
* The web server physically receives the electrical signal over a cable, and recreates the same bits by interpreting the meaning of electrical signals.
* Web server now de-encapsulates the IP packet from the Ethernet frame by removing and discarding the Ethernet header and trailer. Similarly it reads the TCP information and finally hands it over to HTTP process which understands the HTTP request.
* Now webserver now sends back to us the web page in similar steps.

### Ethernet

* Traditional technology for connecting devices in a wired LAN or WAN. It enables devices to communicate with each other via a protocol, which is set of rules.

* Describes how network devices format and transmit data so other devices on the same LAN or network can recognize, receive and process the information. 

#### Working

* IEEE specifies in the family of standards called IEEE 802.3 that the Ethernet Protocol touches both Layer 1 (physical layer) and Layer 2(Data link layer) on OSI

* Units of transimission
    - Packet

        * Frame is wrapped in a packet that contains several bytes of information to establish the connection and mark where the frame starts.
        
    - Frame

        * Includes Payload of data being transmitted
        * MAC address of both sender and receiver
        * Error Correction information to detect transmission problems


### NAT

* Network Address Translation
* This protocol provides a way for multiple computers on a common network to share a single connection to the internet.

### RIP

* Routing Information Protocol.
* Used by routers to send data from one network to another. It efficiently manages routing data by broadcasting its routing table to all other routers within the network. 
* It determines the network distance in units of hop.

### How do proxy servers protect computer networks ?
 
 * Proxy servers primarily prevent external users who are identifying the IP address of an internal network.
 * Can make a network virtually invisible to external users.


### Denial of Service (DOS)

Is overloading the system server so it cannot process legitimate traffic and will be forced to reset.

### TOKEN Ring

* Devices are physically arranged to form a ring. 
* A token is passed among stations. If a station wants to send data, it must wait and capture the token. Only the token holders are permitted to transmit frame. Token ring allows each station to send one frame per turn.

### Digital Data Communication Techniques

#### Parallel.

* All bits of a byte are transferred simultaneously on separate parallel wires. Synchronization between multiple bits is required which becomes difficult over large distance.
* Parallel communication gives large bandwidth but expensive, possible only for devices which are close to each other.

### Serial 

Bits transferred serially one after other.Serial communication gives less bandwidth but cheaper, suitable for transmission over long distances.

### QUIC

* Quick UDP Internet Connections.
* General-purpose transport layer network protocol 
* New generation Internet Protocol that speeds online web applications that are susceptible to delay, such as searching, video streaming etc., by reducing the RTT needed to connect to a server.

Effectively QUIC provides a TCP-like connection using UDP, where all packet handling is performed by the application, and not the network stack of an OS/router/etc.

### HTTP/1.x vs HTTP/2

* HTTP2 is much faster and more reliable than HTTP1. 
* HTTP1 loads a single request for every TCP connection, while HTTP2 avoids network delay by using multiplexing.
* HTTP2 Uses multiplexing, where over a single TCP connection resources to be delivered are interleaved and arrive at the client almost at the same time. It is done using streams which can be prioritized, can have dependencies and individual flow control. It also provides a feature called server push that allows the server to send data that the client will need but has not yet requested

### TCP - 1981

TCP is reliable. One of them to ensure this SYN numbers. 
TCP is all about the data(regarding connection). It does not care about the payload.
TCP options are communicated only in SYN-ACK in 3 way handshake. To capture the options we need to capture the initial 3 way handshake packets.
TCP header Length: 44 bytes.


#### Ephemeral Port

Communication endpoint(port) of a transport layer protocol of the Internet protocol suite that is used for only a short period of time for the duration of a communication session.

#### TTL - Time To Live

TTL is a value for the period of time that a packet, or data, should exist on a computer or network before being discarded.

Most common values: 64, 128, 255.Increases with increase in no. of hops between them

#### RTT - Round Trip Time

#### NAT - Network Address Translation

Process that enables one, unique IP address to represent an entire group of computers. 
In NAT, a network device, often a router or NAT firewall, assigns a computer or computers inside a private network a public address.

####  Three way handshake (SYN, SYN-ACK,ACK)

* SYN (Client) - TCP

    - Sequence Number: 0 (Initial Sequence Number) Randomly generated
    - Source Port: 54846 (Ephemeral Port)
    - Destination Port: 80
    - Acknowledgement Number: 0

* SYN ACK (Server) Accepting Connection -TCP

    - Source Port: 80
    - Destination Port: 54846
    - Sequence Number: 0 (Initial Sequence Number) Randomly generated
    - Acknowledgement Number: 1

* ACK - TCP (Client)

    - Source Port: 54846
    - Destination Port: 80
    - Sequence Number: 1
    - Acknowledgement Number: 1

* HTTP (Client) Sends Get Data (/online)

    - Source Port: 54846
    - Destination Port: 80
    - Sequence Number: 1
    - Acknowledgement Number: 1
    - TCP payload (501 bytes)

* ACK (Server)

    - Source Port: 80
    - Destination Port: 54846
    - Sequence Number: 1
    - Acknowledgement Number: 502 (1 + 501 (client tcp payload))

* PSH,ACK (Server)

    - Source Port: 80
    - Destination Port: 54846
    - Sequence Number: 1
    - Acknowledgement Number: 502 
    - TCP payload (534 bytes)

* ACK (Client)

    - Source Port: 54846
    - Destination Port: 80
    - Sequence Number: 502
    - Acknowledgement Number: 535

* PSH,ACK (Server)

    - Source Port: 80
    - Destination Port: 54846
    - Sequence Number: 535
    - Acknowledgement Number: 502 
    - TCP payload (1186 bytes)

* HTTP/1.1 200 OK (server)

    - Source Port: 80
    - Destination Port: 54846
    - Sequence Number: 1721(535 + 1186)
    - Acknowledgement Number: 502 
    - TCP Payload (5 bytes)

* ACK (Client)

    - Source Port: 54846
    - Destination Port: 80
    - Sequence Number: 502
    - Acknowledgement Number: 1726

* GET /favicon.ico HTTP/1.1

    - Source Port: 54846
    - Destination Port: 80
    - Sequence Number: 502
    - Acknowledgement Number: 1726
    - TCP Payload (439 bytes)

* ACK (server)

    - Source Port: 80
    - Destination Port: 54846
    - Sequence Number: 1726
    - Acknowledgement Number: 941(502+439) 

* ACK (server)

    - Source Port: 80
    - Destination Port: 54846
    - Sequence Number: 1726
    - Acknowledgement Number: 941
    - TCP Payload (500 bytes)

* ACK (server)

    - Source Port: 80
    - Destination Port: 54846
    - Sequence Number: 2226(1726 + 500)
    - Acknowledgement Number: 941
    - TCP Payload (124 bytes)

* ACK (Client)

    - Source Port: 54846
    - Destination Port: 80
    - Sequence Number: 941
    - Acknowledgement Number: 2350(2226 + 124)

#### Important TCP Flags

```
000.            = Reserved
...0            = Nonce
.... 0          = Congestion Window Reduced (CWR)
.... .0         = ECn_Echo
.... ..0        = Urgent
.... ...0       = Acknowledgement
.... .... 0     = Push
.... .... .0    = Reset
.... .... ..1   = Syn (Set it is syn packet)
.... .... ...0  = Fin
```

#### Window : 65535 (2 Bytes initially)

How much data we can send/receive at once.
Latency can destroy the network throughout.

Now TCP header options, has window scale . If it is 9 multiple by 2^9(512) * 65535 

### Throughput 

Amount of data traveling successfully across a network.

### Bandwidth

Maximum data volume capacity of a network.

#### MSS

* Maximum Segment Size
* TCP (Segments)
*  Payload - 1460 (Does not include TCP and IP Headers)

#### MTU

* Maximum Transmission Unit
* Ethernet (Frames)
* Payload - 1500 (Includes TCP and IP headers)
* MTU is per hop, while MSS, since it is on TCP, it's end-to-end.



### Credits

[geeksforgeeks](https://www.geeksforgeeks.org)

[IP Addressing and Subnetting for New Users](https://www.cisco.com/c/en/us/support/docs/ip/routing-information-protocol-rip/13788-3.html)

[interviewbit](https://www.interviewbit.com/networking-interview-questions/)

[Gateway vs Router](https://www.router-switch.com/faq/gateway-router-difference.html)

[Routers vs. Switches vs. Access Points](https://www.baeldung.com/cs/routers-vs-switches-vs-access-points)

[Symmetric Vs Asymmetric Encryption](https://www.ssl2buy.com/wiki/symmetric-vs-asymmetric-encryption-what-are-differences)