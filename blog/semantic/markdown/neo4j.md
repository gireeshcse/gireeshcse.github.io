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

MATCH (ram:Person{email:"ram@example.com"})<-[:FRIEND_OF*2]-(connectionFriend:Person)
RETURN DISTINCT connectionFriend

```

##### Syntax best to use

* Node aliases, Property names - lower camel case.
* Labels - upper camel case.
* Relationships - Upper snake case
* Keyboards - Uppercase.