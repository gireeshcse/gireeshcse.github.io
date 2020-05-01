## Graphs

* A graph is just a collection of vertices and edges.
* A set of nodes and the relationships that connect them.
*  Graphs represent entities as nodes and the ways in which those entities relate to the world as relationships.

### The Labeled Property Graph Model

* It contains nodes and relationships.
* Nodes contain properties (key-value pairs).
* Nodes can be labeled with one or more labels.
* Relationships are named and directed, and always have a start and end node.
* Relationships can also contain properties

### A High-Level View of the Graph Space

* Technologies used primarily for **transactional online graph persistence**, typically accessed directly in real time from an application
* Technologies used primarily for **offline graph analytics**, typically performed as a series of batch steps

## Graph Database

* A graph database management system is an online database management system with Create, Read, Update, and Delete (CRUD) methods that expose a graph data model. 
* Graph databases are generally built for use with transactional (OLTP) systems. Accordingly, they are normally optimized for transactional performance, and engineered with transactional integrity and operational availability in mind.

**Two properties of graph databases**

* The underlying storage
* The processing engine

### Graph Compute Engines

A graph compute engine is a technology that enables global graph computational algorithms to be run against large datasets.

### The Power of Graph Databases

* Performance

    - Queries are localized to a portion of the graph. As a result, the execution time for each query is proportional only to the size of the part of the graph traversed to satisfy that query, rather than the size of the overall graph.

* Flexibility

    - The graph data model expresses and accommodates business needs in a way that enables IT to move at the speed of business.
    - Graphs are naturally additive, meaning we can add new kinds of relationships, new nodes, new labels, and new subgraphs to an existing structure without disturbing existing queries and application functionality.

* Agility

    - The schema-free nature of the graph data model, coupled with the testable nature of a graph database’s application programming interface (API) and query language, empower us to evolve an application in a controlled manner


### Relational Databases Lack Relationships

* Relational databases are designed to codify paper forms and tabular structures.
* Relationships do exist in the vernacular of relational databases, but only at modeling time, as a means of joining tables
* The relational model becomes burdened with large join tables, sparsely populated rows.

Relational schema for storing customer orders in a customer-centric, transactional application

**User Schema**

    user_id     email    Phone      name

**Order Schema**

    order_id    user_id

**Line Item**

    order_id    product_id      quantity

**Product**

    product_id      description     handling


The application exerts a tremendous influence over the design of this schema, making some queries very easy, and others more difficult:

* Join tables add accidental complexity; they mix business data with foreign key metadata.
* Foreign key constraints add additional development and maintenance overhead just to make the database work.
* Sparse tables with nullable columns require special checking in code, despite the presence of a schema.
* Several expensive joins are needed just to discover what a customer bought.
* Reciprocal queries are even more costly. “What products did a customer buy?” is relatively cheap compared to “which customers bought this product?”, which is the basis of recommendation systems. We could introduce an index, but even with an index, recursive questions such as “which customers buying this product also bought that product?” quickly become prohibitively expensive as the degree of recursion increases.

Relational databases struggle with highly connected domains.

Example: Social Domains

**Person Schema**

    ID      Person

**PersonFriend**

    PersonID    FriendID

**Bob’s friends Query**


    SELECT p1.Person
    FROM Person p1 JOIN PersonFriend
        ON PersonFriend.FriendID = p1.ID
    JOIN Person p2
        ON PersonFriend.PersonID = p2.ID
    WHERE p2.Person = 'Bob'

**Who is friends with Bob?**

    SELECT p1.Person
    FROM Person p1 JOIN PersonFriend
        ON PersonFriend.PersonID = p1.ID
    JOIN Person p2
        ON PersonFriend.FriendID = p2.ID
    WHERE p2.Person = 'Bob'

This reciprocal query is still easy to implement, but on the database side it’s more expensive, because the database now has to consider all the rows in the PersonFriend table.

**Alice’s friends-of-friends**

    SELECT p1.Person AS PERSON, p2.Person AS FRIEND_OF_FRIEND
    FROM PersonFriend pf1 JOIN Person p1
        ON pf1.PersonID = p1.ID
    JOIN PersonFriend pf2
        ON pf2.PersonID = pf1.FriendID
    JOIN Person p2
        ON pf2.FriendID = p2.ID
    WHERE p1.Person = 'Alice' AND pf2.FriendID <> p1.ID



### NOSQL Databases Also Lack Relationships

Most NOSQL databases—whether key-value-, document-, or column-oriented— store sets of disconnected documents/values/columns. This makes it difficult to use them for connected data and graphs.

Aggregate stores are not functionally equivalent to graph databases with respect to connected data.
Aggregate stores do not maintain consistency of connected data, nor do they support what is known as index-free adjacency, whereby elements contain direct links to their neighbors. As a result, for connected data problems, aggregate stores must employ inherently latent methods
for creating and querying relationships outside the data model.


A small social network encoded in an aggregate store

        User: Ram
        Friends: [Sita]

        User: Sita
        Friends: [Hanuman, Rohini]

        User: Hanuman
        Friends: [Ram]

        User: Bharat
        Friends: [Rohini, Ram]

Requires numerous index lookups but no brute-force scans of the entire dataset.

Here friendship isn’t always symmetric.**"who is friends with Bob?”** is difficult rather than **“who are Bob’s friends?”** because it requires brute-force scan across the whole dataset looking for friends entries that contain Bob.

For above probelm, a graph database provides constant order lookup for the same query. In this case, we simply find the node in the graph that represents Bob, and then follow any incoming friend relationships; these relationships lead to nodes that represent people who consider Bob to be their friend. This is far cheaper than brute-forcing the result because it considers far fewer members of the network; that is, it considers only those that are connected to Bob. Of course, if everybody is friends with Bob, we’ll still end up considering the entire dataset.

To avoid having to process the entire dataset, we could denormalize the storage model by adding backward links. Adding a second property, called perhaps **friended_by** , to each user, we can list the incoming friendship relations associated with that user. But this doesn’t come for free. For starters, we have to pay the initial and ongoing cost of increased write latency, plus the increased disk utilization cost for storing the additional metadata. On top of that, traversing the links remains expensive, because each hop requires an index lookup. This is because aggregates have no notion of locality, unlike graph databases, which naturally provide index-free adjacency through real—not reified—relationships. By implementing a graph structure atop a nonnative store, we get some of the benefits of partial connectedness, but at substantial cost.

This substantial cost is amplified when it comes to traversing deeper than just one hop. Friends are easy enough, but imagine trying to compute—in real time—friends-of-friends, or friends-of-friends-of-friends.

Many systems try to maintain the appearance of graph-like processing, but inevitably it’s done in batches and doesn’t provide the real-time interaction that users demand.

### Graph Databases Embrace Relationships

Relationships in a graph naturally form paths. Querying—or traversing—the graph involves following paths. Because of the fundamentally path-oriented nature of the data model, the majority of path-based graph database operations are highly aligned with the way in which the data is laid out, making them extremely efficient

Graph databases are the best choice for connected data.

Finding extended friends in a relational database versus efficient finding in Neo4j

`
Depth   | RDBMS execution time(s)    | Neo4j execution time(s)   |  Records returned
2       |  0.016                     | 0.01                      | ~2500
3       | 30.267                     | 0.168                     |  ~110,000
4       | 1543.505                   | 1.359                     | ~600,000
5       | Unfinished                 | 2.132                     | ~800,000
`

Neo4j’s response time remains relatively flat: just a fraction of a second to perform the query—definitely quick enough for an online system.

An R-Tree is a graph-like index that describes bounded boxes around geographies. Using such a structure we can describe overlapping hierarchies of locations.

From the data practitioner’s point of view, it’s clear that the graph database is the best technology for dealing with complex, variably structured, densely connected data— that is, with datasets so sophisticated they are unwieldy when treated in any form other than a graph.

### Data Modeling with Graphs

* A labeled property graph is made up of nodes, relationships, properties, and labels.
* Nodes contain properties. Think of nodes as documents that store properties in the form of arbitrary key-value pairs. In Neo4j, the keys are strings and the values are the Java string and primitive data types, plus arrays of these types.
* Nodes can be tagged with one or more labels. Labels group nodes together, and indicate the roles they play within the dataset.
* Relationships connect nodes and structure the graph. A relationship always has a direction, a single name, and a start node and an end node—there are no dangling relationships. Together, a relationship’s direction and name add semantic clarity to the structuring of nodes.
* Like nodes, relationships can also have properties. The ability to add properties to relationships is particularly useful for providing additional metadata for graph algorithms, adding additional semantics to relationships (including quality and weight), and for constraining queries at runtime.

### Querying Graphs: An Introduction to Cypher

Cypher is an expressive (yet compact) graph database query language.