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