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