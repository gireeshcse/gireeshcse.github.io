#### Spanner

* Google's Spanner, a scalable, multi-version, globally-distributed, and synchronously-replicated database.

Spanner is a distributed SQL database developed by Google. Spanner is a globally distributed database service and storage solution. It provides features such as global transactions, strongly consistent reads, and automatic multi-site replication and failover.

#### Important Pointers

* Large number of data records which does not require lots of joins or require a lot of aggregation such as summing, counting, averaging etc.. on these data, In such a scenarios relational databases are more effiecient than graph databases.

* Neo4j (Designed/Suited for)
    - Transactional ACID-compliant database.
    - Online Transaction Processing(OLTP)
    - Scalabillity

    - Complex join-intensive queries
    - Pathfinding queries


* **Cypher**

    - Declarative, pattern-matching query language

#### Data Modeling

* Nodes
* Relationships
* Properties
* Labels

* Design Model based on queries on data


##### Some Patterns

        
###### Example 1    
    
        name: "India"
        languages_spoken: ["Hindi", "English", "Telugu", "Tamil"]


        Language(name)  ----[:SPOKEN_IN]---> Country(name)

###### Example 2

        Country(name, website, language_name, currency_code, currency_name, number_of_words)

        Country(name, website) ----[:USES_CURRENCY]--> Currency(code,name)
        Country -----[:SPEAKS]---> Language(name,number_of_words)

###### Example 3

        BeerBrand(name) ----[:IS_A]----> BeerType(name)
        Brewery(name) ---[:BREWS]---> BeerBrand(name)
        BeerBrand(name) ----[:HAS_ALC_PERC]----> AlcoholPercentage(value)

        AlcoholPercentage(value:6.1)  ---[:PRECEDES]---> AlcoholPercentage(value:6.2) ---[:PRECEDES]---> AlcoholPercentage(value:6.3)

Note: Indexing can be used for AlcholPercentages

#### Dense node pattern

* Some parts of the graph are all connected to same node. It effects the graph traversals because all the connected relationships need to be evaluated to determine for next traversal from that node.

These are to modelled accordingly.

Example:

        Person(name:"sachin",dob:"1973-04-24") ------[:IS_A]---> Celebrity(type:"cricketer")
        FAN(id:"S01")---[:LIKES]---> Person(name:"sachin",dob:"1973-04-24")
        FAN(id:"S20M")---[:LIKES]---> Person(name:"sachin",dob:"1973-04-24")


        FAN(id:"S20M")---[:DENSE_LIKES]--->Meta(name:"fan")---[:LIKES]--->Person(name:"sachin",dob:"1973-04-24")
        FAN(id:"S20M")---[:DENSE_LIKES]--->Meta(name:"fan")---[:LIKES]--->Person(name:"dhoni",dob:"1981-07-07")

We can use meta nodes to connect to the top level nodes. We also limit number of relationships for meta node for effective traversal.


#### Cypher

##### CREATE

```
CREATE (ram:Person{name:"Sri Ram"}) - [:LOVES]-> (sita:Person{name:"Sita Devi"})
CREATE (sita)-[:LOVES]-> ram
```

The above creates 2 nodes but below one creates 4 nodes since there is no way to identify the nodes.

```
CREATE (:Person{name:"Sri Ram"}) - [:LOVES]-> (:Person{name:"Sita Devi"})
CREATE (:Person{name:"Sita Devi"}) - [:LOVES]-> (:Person{name:"Sri Ram"})
```

```
CREATE(ram:Royal:Person{name:"Sri Ram Chandra"})-[:LOVES{since:"a long time ago", till:"forever",where:"India"}]->(sita:Royal:Person{name:"Sita Devi"})
CREATE (ram) <-[:LOVES]-(sita)
```
Above Added 4 labels, created 2 nodes, set 5 properties, created 2 relationships,


##### READ
```
MATCH(n)
RETURN n;
```
Returns all the nodes.

```
MATCH(n:Person)
WHERE n.name="Sri Ram" or n.firstName="Sri Ram"
RETURN n;
```
firstName may exits for some nodes or not at all.

```
MATCH (n:Person)-[:LOVES]-()
WHERE toLower(n.name) = "sita devi"
RETURN n ORDER BY n.name ASC
SKIP 2 LIMIT 5;

MATCH (n:Person)-[:LOVES]->()
RETURN n;

MATCH (n:Person)<-[:LOVES]-()
RETURN n;
```

##### UPDATE

```
MATCH (n:Person{name:"Sita Devi"})
SET n.gender = "Male"
RETURN n;

MATCH (n:Person{name:"Sita Devi"})
WHERE n.gender = "Male"
SET n.gender = "Female"
SET n.age = 1000
RETURN n;

MATCH (n:Person{name:"Sita Devi"})
WHERE n.gender = "Male"
SET n.gender = "Female", n.age = 1000
RETURN n;

MATCH (n:Person)
SET n.age = n.age + 1
RETURN n;

MATCH (n:Person)
WHERE n.age >= 12 and n.age < 18
SET n.Teen
RETURN n;

MATCH (n:Person)
WHERE n.age > 18
REMOVE n.Teen
RETURN n;

MATCH (n:Person)
WHERE n.age IS NULL
SET n.age = 0
RETURN n;

MATCH (n:Person{name:"Sita Devi"})
REMOVE n.firstName
RETURN n;

```

##### DELETE

```
MATCH (r:Person{name:"Sri Ram"})
DELETE r;
```
Cannot delete node<0>, because it still has relationships. To delete this node, you must first delete its relationships.
```
MATCH (r:Person{name:"Sri Ram"})
DETACH DELETE r;
```

```
MATCH(me)
OPTIONAL MATCH (me)-[r]-()
DELETE me, r
```
Delete all nodes and relationships

##### Others

```
MERGE (me:Person{name:"Finch"})
ON MATCH SET me.accessed = timestamp()
ON CREATE SET me.age = 65
RETURN me.name,me.accessed, me.age;

MERGE (keanu:Person {name: 'Keanu Reeves'})
ON CREATE SET keanu.created = timestamp();

CREATE CONSTRAINT ON (p:Person) ASSERT p.identifier is UNIQUE;
DROP CONSTRAINT ON (p:Person) ASSERT p.identifier is UNIQUE;

MERGE (p:Person{name: "Keanu Reeves"})
ON MATCH SET p.identifier = "email@example.com"
RETURN p;

// Node(5) already exists with label `Person` and property `identifier` = 'email@example.com'
MERGE (p:Person{name: "Sita Devi"})
ON MATCH SET p.identifier = "email@example.com"
RETURN p;

// list friends of the friends of Ram
MATCH (ram:Person{email:"ram@example.com"})<-[:FRIEND_OF]-(immediateFriend:Person)<-[:FRIEND_OF]-(connectionFriend:Person)
RETURN connectionFriend

// equivalent to above but distinct as well
MATCH (ram:Person{email:"ram@example.com"})<-[:FRIEND_OF*2]-(connectionFriend:Person)
RETURN DISTINCT connectionFriend

```

##### Syntax best to use

* Node aliases, Property names - lower camel case.
* Labels - upper camel case.
* Relationships - Upper snake case
* Keywords - Uppercase.

##### APOC

* A Package of Components
* Awesome Procedures On Cypher

Collection of functions and procedures that are available for use in Cypher.

APOC Core can be installed by moving the APOC jar file from the $NEO4J_HOME/labs directory to the $NEO4J_HOME/plugins directory and restarting Neo4j.

Note: Latest release can be downloaded from the [Github](https://github.com/neo4j-contrib/neo4j-apoc-procedures/releases)


```
// count of functions and procedures of apoc. 252, 276(APOC 4.2.0.5)
CALL dbms.functions() YIELD name
WHERE name STARTS WITH 'apoc.'
RETURN count(name)
UNION
CALL dbms.procedures() YIELD name 
WHERE name STARTS WITH 'apoc.'
RETURN count(name)

// List functions
CALL dbms.functions()

// List procedures
CALL dbms.procedures()
```

* **Functions** are designed to return a single value after a computation that only reads the database.
* **Procedures**  can make changes to the database and return several results.

```
CREATE (p:Person{GUID: apoc.create.uuid()})
```

* For help

```
CALL apoc.help('sum') // finds all the procedures and function that contain sum in name or description
// "procedure", "apoc.trigger.resume", ... 
// "function",	"apoc.coll.sum", ..
// "function"	"apoc.coll.sumLongs", ..
```

Procedures are restricted by default, we need the set the following setting in  $NEO4J_HOME/conf/neo4j.conf

```
dbms.security.procedures.unrestricted=apoc.*
```

```
CALL apoc.meta.graph; // gives meta info about db.
```

* Random graph generator.

apoc.generate.ba(noNodes, edgesPerNode, label, type) - generates a random graph according to the Barabasi-Albert model"

```
CALL apoc.generate.ba(100,2,'Person','FRIEND_OF'); 

// To view all created nodes
MATCH(n)
RETURN n;

// person who is friend of him/her self
MATCH (p:Person)-[FRIEND_OF]-(p)
RETURN p
```

##### Datascience library 

* Latest can be downloaded from [Neo4J](https://neo4j.com/graph-data-science-software/) or [Github](https://github.com/neo4j/graph-data-science/releases)
* Copy the jar to  the $NEO4J_HOME/plugins directory and restart Neo4j.
* set the following setting in  $NEO4J_HOME/conf/neo4j.conf 
```
dbms.security.procedures.unrestricted=apoc.*,gds.*
```

* **Graph Catalog**

        - Graph algorithms run on a graph data model which is a projection of the Neo4j property graph data model. A graph projection can be seen as a view over the stored graph, containing only analytically relevant, potentially aggregated, topological and property information. Graph projections are stored entirely in-memory using compressed data structures optimized for topology and property lookup operations.

        - The graph catalog is a concept within the GDS library that allows managing multiple graph projections by name. Using its name, a created graph can be used many times in the analytical workflow. Named graphs can be created using either a Native projection or a Cypher projection. After usage, named graphs can be removed from the catalog to free up main memory.

* PageRank ([Ref](https://neo4j.com/docs/graph-data-science/current/algorithms/page-rank/))

   - It counts both number and quality of a relation and deduces an importance for the node.
   - Rules are
     - The more relations, the more important node
     * The more relations with important nodes, the more important the node.
   - Now present in Neo4J Data science library.


```
// creating thousand nodes
FOREACH (id IN range(0,1000) | CREATE (n:Page{id:id}) )

// create at most, a 100 thousand relations
MATCH(n1:Page), (n2:Page)
WITH n1,n2 LIMIT 1000000 
WHERE rand() <  0.1
CREATE (n1)-[:LINKS{weight:rand()}]-> (n2)
// Set 99673 properties, created 99673 relationships, completed after 641 ms.

// Creates a graph in the catalog using a Native projection.
CALL gds.graph.create(
  'myGraph',
  'Page',
  'LINKS',
  {
    relationshipProperties: 'weight'
  }
)
// Returns nodeProjection	relationshipProjection	graphName	nodeCount	relationshipCount	createMillis

// memory estimation for running the pagerank

CALL gds.pageRank.write.estimate('myGraph', {
  writeProperty: 'pageRank',
  maxIterations: 20,
  dampingFactor: 0.85
})
YIELD nodeCount, relationshipCount, bytesMin, bytesMax, requiredMemory
╒═══════════╤═══════════════════╤══════════╤══════════╤════════════════╕
│"nodeCount"│"relationshipCount"│"bytesMin"│"bytesMax"│"requiredMemory"│
╞═══════════╪═══════════════════╪══════════╪══════════╪════════════════╡
│1001       │99673              │24936     │24936     │"24 KiB"        │
└───────────┴───────────────────┴──────────┴──────────┴────────────────┘

// In the stream execution mode, the algorithm returns the score for each node.
CALL gds.pageRank.stream('myGraph')
YIELD nodeId, score
RETURN gds.util.asNode(nodeId).id AS name, score
ORDER BY score DESC, name ASC
╒══════╤══════════════════╕
│"name"│"score"           │
╞══════╪══════════════════╡
│609   │1.205301932941157 │
├──────┼──────────────────┤
│399   │1.1966946890971681│
├──────┼──────────────────┤
│202   │1.1890399693567406│
├──────┼──────────────────┤
│960   │1.1743968396459592│
├──────┼──────────────────┤

// Once we have finished using the named graph we can remove it from the catalog to free up memory.
CALL gds.graph.drop('myGraph') YIELD graphName;

// cartesian product
MATCH(n1:Node), (n2:Node)
return n1,n2
```

* Timeboxed execution of Cypher Statements

```
CALL apoc.cypher.runTimeBoxed(CypherStatement, params, timeoutInMs)
CALL apoc.cypher.runTimeBoxed('MATCH (n) return n', NULL, 2)
// Neo.TransientError.Transaction.Terminated
call apoc.cypher.runTimeboxed("MATCH (n:Page) RETURN n", {}, 1)
YIELD value
RETURN value.n.id; // RETURN value.n creating Neo.DatabaseError.Statement.ExecutionFailed
```

* Linking of a collection of nodes

Calling a single procedure will make it possible to create ONE(same) relation between all the nodes of a collection.
```
MATCH (p:Person)
WITH collect(p) AS persons
CALL apoc.nodes.link(persons,'BRO')
RETURN persons

MATCH (p:Person)-[r:BRO]->(m:Person)
RETURN p, count(r)
╒═══════════════════════╤══════════╕
│"p"                    │"count(r)"│
╞═══════════════════════╪══════════╡
│{"name":"User0","id":0}│1         │
├───────────────────────┼──────────┤
│{"name":"User1","id":1}│1         │
├───────────────────────┼──────────┤
│{"name":"User2","id":2}│1         │
├───────────────────────┼──────────┤
│{"name":"User3","id":3}│1         │
├───────────────────────┼──────────┤
```

[For More Info](https://neo4j.com/docs/labs/apoc/current/)
  - Math Functions
  - NLP
  - Graph Algorithms
  - Temporal (Date Time)
  - etc...