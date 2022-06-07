### Database

#### Relational Database

* Designed to capture and record data for OLTP(Online Transaction Processing)
* Live, real-time data
* Data stored in tables with rows and columns
* Flexible Schema(How data is organised) - We can change as we go
* Vertically Scalable (Increase Size of Instance - RAM, CPU etc.)
* Example: MySQL, Postgresql, Microsoft SQL Server etc..
* Have unique identifiers (unique/primary keys)

##### Uses

* Stores structured data
* Better understood data integrity
* Flexible quering
* Handles medium level of transactions well.

Note: Horizontal Scaling - Add more instances.It is proven to quite efficient because we are not putting more work load on a single instance.

#### Data Warehouse

* Designed for analytical processing - OLAP (Online Analytical Processing) - To analyze huge amounts of data
* ETL - Extracts, Transforms, Loads
* Data is refreshed for source systems through ETL - Stores current and historical
* Data is summarized and refreshed periodically
* Rigid Schema (How the data is organised) - We have  to plan ahead 


Note: Databases work slowly for quering large amounts of data and can slow down transactional processes, Data Warehouses don't interfere with any processes and generally faster.


#### Data Lake

* Designed to capture raw data (structured, semi-structured, unstructured) - Video, Picture etc any kind of data.
* Made for large amounts of data
* Used for ML and AI in its current state or for Analytics with processing.
* Can organize and put into Databases or Data Warehouses.

#### NoSQL - Not Only SQL

* Doesn't store data in tables with columns/rows (Semi/unstructured data)
* Horizontally Scalable.
* Different types of NoSQL - Document-Type, Key-Value, Graph, Wide-Column
* Examples: MongoDB, Cassandra, BigTable, Redis, HBase, Neo4j and CouchDB.

##### Uses

* Stores any unstructured data
* Highly scalable
* Handles large amounts of transactions well
* Low latency for data retrieval.
