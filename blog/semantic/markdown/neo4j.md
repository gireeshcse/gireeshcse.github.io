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


