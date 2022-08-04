### ETL Concepts

* ETL(Extract, Transform, Load) is a process of data integration that encompasses three steps - Extraction, Transformation, and Loading.
* ETL systems take large volumes of raw data from multiple sources converts it for analysis, and loads that data into our warehouse.

#### Uses

* To optimize our data for analytics in an optimized environment.
* Cross-domain analysis: Joining data from disparate data sources

#### Tools

* Amazon Redshift
    - Provides access to a variety of data analytics tools, compliance features and even artificial intelligence and machine learning applications.

* Google BigQuery

#### Extraction

* Extracted data sets come from a source into a staging area
* Staging area acts as a buffer b/w the data warehouse and the source data.
* Staging area is used for data cleansing and organisation.

#### Transformation

* Data cleaning and organisation stage is the transformation stage.
* Data from multiple source systems will be normalized and converted to a single system format - improving data quality and complaince.
* ETL yields transformed data through these methods.
    - Cleaning
        Example: Mapping NULL to 0 or "Male" to "M" and "Female" to "F", date format consistency etc
    - Filtering
    - Joining
        - Linking data from multiple sources 
    - Sorting
    - Spliting
        - Splitting a single column into multiple columns.
    - Deduplication

        - Process that elimates excessive copies of data and significatly decreases storage capacity requirements.
        - Identifying and removing duplicate records.

    - Summarization

        - Values are summarized to obtain total figures which are calculated and stored at multiple levels as business metrics.
        - Example: Adding all purchases a customer has made to build a customer lifetime value(CLV) metric.

#### Loading

* Data that has been extracted to a staging area and transformed is loaded into our data warehouse.
* Depending upon our business needs, data can be loaded in batches or all at once.


### OLTP - Online Transaction Processing

### OLAP - Online Analytical Processing

### Relational Database Management System

Data stored in relations(Tables)

#### ACID

* Atomicity
* Consistency
* Isolation
* Durability

### SQL - Structured Query Language

Core of relational database which is used for accessing and managing database.

#### Subsets

* DDL - Data Definition Language

    * Create
    * Alter
    * Drop

* DML - Data Manipulation Language

    * insert
    * update
    * delete
    * select

* DCL - Data Control Language

    * Grant
    * Revoke

#### SELECT

* WHERE
* ORDER BY
* GROUP BY
* HAVING


* UNION
    - Used to combine the results of two tables while also removing duplicate entries.

```
SELECT * FROM log WHERE log.refererurl is NULL
UNION
SELECT * FROM log WHERE log.user_id is NULL;
```

* MINUS
    - Used to return rows from the first query but not from the second query.

```
SELECT NAME, AGE , GRADE
FROM Table1
MINUS 
SELECT NAME, AGE, GRADE 
FROM Table2
```

* INTERSECT
    - Used to combine the results of both queries into a single row.

```
SELECT * FROM log WHERE log.refererurl is NULL
INTERSECT
SELECT * FROM log WHERE log.user_id is NULL;
```

#### SQL and  BigQuery

```
from google.cloud import bigquery
client = bigquery.Client()
```

In BigQuery, each dataset is contained in a corresponding project. 

```
# hacker_news dataset.
dataset_ref = client.dataset("hacker_news",project="bigquery-public-data")

# API request - fetch the dataset
dataset = client.get_dataset(dataset_ref)
```

Every dataset is just a collection of tables. 

```
tables = list(client.list_tables(dataset))
for table in tables:
    print(table.table_id)

# contruct a reference to the "full" table
table_ref = dataset_ref.table('full')

# API request- fetch the table
table = client.get_table(table_ref)
```

* **Client** Object hold projects and a connection to the BigQuery service
* **project** is a collection of datasets
* **Dataset** is a collection of tables.

```
# print information on all the columns in the "full" table 
table.schema 

[
    SchemaField('title', "STRING", 'NULLABLE', 'Story Title', (),None),


    SchemaField('deleted', 'BOOLEAN', 'NULLABLE', 'Is deleted?', (), None)

]
```

```
SchemaField(name,field_type,mode,description)

# To prieview first five rows
client.list_rows(table,max_results=5).to_dataframe()
client.list_rows(table,selected_rows=table.schema[:1],max_results=5).to_dataframe()
```

```
from google.cloud import bigquery

# create client object
client = bigquery.Client()

# contruct a reference to the "chicago_crime" dataset
dataset_ref = client.dataset("chicage_crime", project="bigquery-public-data")

# API request - fetch the dataset
dataset = client.get_dataset(dataset_ref)

# get the tables in the dataset
tables = list(client.list_tables(dataset))
num_tables = len(tables) # to get the number of tables

# get columns in 'crime' table have "TIMESTAMP" data

## get the table ref 
table_ref = dataset_ref.table('crime')

## get table 
table = client.get_table(table_ref)

count = 0
for field in table.schema:
    if field.field_type == "TIMESTAMP":
        count += 1

df = client.list_rows(table,max_results=5).to_dataframe()
```

Queries

```
query ="""
SELECT city FROM `bigquery-public-data.openaq.global_air_quality` 
WHERE country = 'US'
"""

client = bigquery.Client()
query_job = client.query(query)
us_cities = query_job.to_dataframe()
us_cities.city.value_counts().head()

query_popular = """
SELECT parent, COUNT(id) AS NumPosts
FROM `bigquery-public-data.hacker_news.comments`
GROUP BY parent
HAVING count(id) > 10
"""
```
Note: If we have any GROUP BY clause, then all variables must be passed to rither a GROUP_BY or an aggregation function.


```
SELECT ID, Name,Animal FROM table
ORDER BY Animal DESC

SELECT Name, EXTRACT(DAY from Date) AS Day
FROM table

SELECT Name, EXTRACT(WEEK from Date) AS Day
FROM table
```

Aggregate functions

* COUNT
* SUM
* AVG
* MIN
* MAX


```
SELECT Animal, COUNT(ID) AS Number FROM
`bigquery-public-data.pet_records.pets`
GROUP BY Animal
HAVING COUNT(ID) > 1

# CTE - Common table expression is a temporary table that we return within our query.
query = """
    WITH Seniors AS
    (
        SELECT ID, Name FROM `bigquery-public-data.pet_records.pets`
        WHERE Years_old > 5
    )
    SELCT ID FROM Seniors
"""

WITH time AS
(
    SELECT DATE(block_timestamp) AS trans_date
    FROM `transactions`
)
SELECT COUNT(1) AS transactions, trans_date 
FROM time
GROUP BY trans_date
ORDER BY trans_date

# List all the tables in the dataset
tables = list(client.list_tables(dataset))

# Print names of all tables in the dataset (there is only one!)
for table in tables:  
    print(table.table_id)

table_name = 'taxi_trips'
table_ref = dataset_ref.table(table_name)
table = client.get_table(table_ref)
df = client.list_rows(table,max_results=5).to_dataframe()
```
Write a query that counts the number of trips in each year. 
```
# Your code goes here
rides_per_year_query = """
SELECT EXTRACT(YEAR FROM trip_start_timestamp) AS year, 
        COUNT(1) AS num_trips
FROM `bigquery-public-data.chicago_taxi_trips.taxi_trips`
GROUP BY year
ORDER BY year
"""
rides_per_year_query = """
WITH trips as (
 SELECT EXTRACT(YEAR FROM trip_start_timestamp) year from `bigquery-public-data.chicago_taxi_trips.taxi_trips`
)
SELECT year, COUNT(1) as num_trips from `trips`
GROUP BY year
ORDER BY year
"""

# Set up the query (cancel the query if it would use too much of 
# your quota)
safe_config = bigquery.QueryJobConfig(maximum_bytes_billed=10**10)
rides_per_year_query_job = client.query(rides_per_year_query) # Your code goes here

# API request - run the query, and return a pandas DataFrame
rides_per_year_result = rides_per_year_query_job.to_dataframe() # Your code goes here

# View results
print(rides_per_year_result)
```

Rides per month in 2017

```
rides_per_month_query = """
WITH trips as (
 SELECT EXTRACT(MONTH FROM trip_start_timestamp) AS month
 from `bigquery-public-data.chicago_taxi_trips.taxi_trips`
 WHERE EXTRACT(YEAR FROM trip_start_timestamp) = 2017
)
SELECT month, COUNT(1) as num_trips from `trips`
GROUP BY month
ORDER BY month
""" 

rides_per_month_query = """
    SELECT EXTRACT(MONTH FROM trip_start_timestamp) AS month, 
            COUNT(1) AS num_trips
    FROM `bigquery-public-data.chicago_taxi_trips.taxi_trips`
    WHERE EXTRACT(YEAR FROM trip_start_timestamp) = 2017
    GROUP BY month
    ORDER BY month
"""
```

 Write a query that shows, for each hour of the day in the dataset, the corresponding number of trips and average speed.

 ```
WITH RelevantRides AS
(
    SELECT EXTRACT(HOUR FROM trip_start_timestamp) AS hour_of_day, 
            trip_miles, 
            trip_seconds
    FROM `bigquery-public-data.chicago_taxi_trips.taxi_trips`
    WHERE trip_start_timestamp > '2017-01-01' AND 
            trip_start_timestamp < '2017-07-01' AND 
            trip_seconds > 0 AND 
            trip_miles > 0
)
SELECT hour_of_day, 
        COUNT(1) AS num_trips, 
        3600 * SUM(trip_miles) / SUM(trip_seconds) AS avg_mph
FROM RelevantRides
GROUP BY hour_of_day
ORDER BY hour_of_day
 ```

JOIN

```
# Data is retrieved if values are present in both tables.
SELECT p.Name as pet_name, o.Name as owner_name
FROM `pets` AS p
INNER JOIN `owners` as o on p.ID = o.Pet_ID

# to get list of tables
temp = list(client.list_tables(dataset))
list_of_tables = [table.table_id for table in temp ] 

# select answer for a topic related to bigquery
SELECT a.id, a.body, a.owner_user_id from 
`bigquery-public-data.stackoverflow.posts_answers` AS a
INNER JOIN `bigquery-public-data.stackoverflow.posts_questions` AS q on q.id = a.parent_id
WHERE q.tags LIKE '%bigquery%'

# Write a new query that has a single row for each user who answered at least one question with a tag that includes the string "bigquery"
SELECT a.owner_user_id AS user_id, COUNT(1) as number_of_answers from 
`bigquery-public-data.stackoverflow.posts_answers` AS a
INNER JOIN `bigquery-public-data.stackoverflow.posts_questions` AS q on q.id = a.parent_id
WHERE q.tags LIKE '%bigquery%'
GROUP BY user_id
HAVING number_of_answers > 0
```

### pandas

* Data analysis(Use)
* Core Objects - DataFrame and Series.

#### DataFrame

* Table
* Contains an array of individual entries.
* Each entry corresponds to a row and a column


```
pd.DataFrame({'Yes':[50,21],'No':[131,2]})

	Yes 	No
0 	50 	    131
1 	21 	    2

```

* Row labels are known as **Index**

```
pd.DataFrame({
                'Ram': ['Awesome','Average'],
                'Sita': ['Good','Bad']
            },
            index=['A','B'])

	Ram 	    Sita
A 	Awesome     Good
B 	Average     Bad	    
```

#### Series

* A sequence of data values.
* Like a Single column of a DataFrame
* Does not have a column name.

```
pd.Series([1,2,3,4])


0    1
1    2
2    3
3    4
4    5
dtype: int64

ingredients = pd.Series(["4 cups","1 cup","2 large","1 can"],index=["Flour","Milk","Eggs","Spam"],name="Dinner")
```

#### Other topics

Reading a csv

```
fd = pd.read_csv("file.csv")
fd.shape # (no_of_rows,no_of_columns)
fd.head() # grabs the first five rows
fd = pd.read_csv("file.csv",index_col=0) # the first column is treated as row index

animals = pd.DataFrame({'Cows': [12, 20], 'Goats': [22, 19]}, index=['Year 1', 'Year 2'])
animals.to_csv("cows_and_goats.csv")

# accessing columns
animals.Cows
animals["Cows]

# index base selection
animals["Cows"][0] # 12
animals.iloc[0] # selecting first row of data in a DataFrame.
animals.iloc[:,0] # retrive first column data
animals.iloc[:3,0] # retrive first column data from first 3 rows
animals.iloc([1,2],0) # retirve from the list.

# label based selection
animals.loc[0,"Cows"]
animals.loc[:,["Cows","Goats"]]
```

* **iloc** uses stdlib indexing scheme (0:10 => 0,1...,9 )
* **loc** uses indexes inclusively (0:10 => 0,1...,9,10 )

```
reviews.set_index("title") # manipulating the index
reviews.country == "Italy" # conditional selection. Produces a series of True/False booleans based on the Country
reviews.loc[reviews.country == 'Italy'] # selects rows which have the country as Italy.
reviews.loc[(reviews.country == 'Italy') & (reviews.points >= 90)]
reviews.loc[(reviews.country == 'Italy') | (reviews.points >= 90)]
reviews.loc[reviews.country.isin(["Italy", "France"])]
reviews.loc[reviews.price.notnull()]

reviews["critic"] = "everyone" # sets all rows value of column 
reviews["index_backwards"] = range(len(reviews),0,-1)

sample_reviews = reviews.iloc[[1,2,3,5,8]] # Select the records with index labels 1, 2, 3, 5, and 8
first_descriptions = reviews.loc[:9,"description"] # Select the first 10 values from the description
df = reviews.loc[[0,1,10,100],["country",'province','region_1','region_2']]

df = reviews.loc[:99,["country","variety"]]
df = reviews.iloc[:100,[0,11]]

# top_oceania_wines containing all reviews with at least 95 points (out of 100) for wines from Australia or New Zealand
top_oceania_wines = reviews.loc[((reviews.country == "Australia") | (reviews.country == "New Zealand")) & (reviews.points >= 95)]
```

Summary functions.

```
reviews.points.describe() # count, mean, max etc it is type aware
reviews.taster_name.describe() # count,unique, top,freq
reviews.points.mean()
median_points = reviews.points.median()
reviews.taster_name.unique()
reviews.taster_name.value_counts() # list of unique values and how often they occur in the dataset.

```

Maps

```
review_points_mean = reviews.points.mean()
reviews.points.map(lambda p: p - review_points_mean) # returns Series where all values are transformed
```

apply() is the equivalent method if we want to transform a whole DataFrame by calling a custom method on each row.

```
def remean_points(row):
    row.points = row.points - review_points_mean
    return row

reviews.apply(remean_points,axis='columns')
```

reviews.apply() with axis=index, we would need to give a function to transform each column instead of row.

map() and apply() return new, transformed Series and DataFrames, respectively.

```
# pandas built-ins for many common mapping operations.
review_points_mean = reviews.points.mean()
review.points - review_points_mean # returns a series
# for combining two seriess of equal length
reviews.country + " - " + reviews.region_1

#  Create a variable bargain_wine with the title of the wine with the highest points-to-price ratio in the dataset.
def add_points_price_ratio(row):
    row["points_price_ratio"] = row.points / row.price
    return row
bargain_wine = reviews.apply(add_points_price_ratio,axis="columns")
max_value = bargain_wine.points_price_ratio.max()
bargain_wine = bargain_wine.loc[bargain_wine.points_price_ratio == max_value]
bargain_wine = bargain_wine.iloc[0].title # to get title

# above answer can be 
bargain_idx = (reviews.points / reviews.price).idxmax()
bargain_wine = reviews.loc[bargain_idx, 'title'] # to get only title


# creating a series with "tropical" or "fruity" in description
def add_counts_ratio(row):
    data = {
        "tropical":False,
        "fruity": False
    }
    for key in data:
        if row["description"].find(key) != -1:
            data[key] = True
        row[key] = data[key]
    return row

reviews_new = reviews.apply(add_counts_ratio,axis="columns")
fr_df = reviews_new.fruity.value_counts()
tr_df = reviews_new.tropical.value_counts()
descriptor_counts = pd.Series([tr_df.get(key=True),fr_df.get(key=True)],index=["tropical","fruity"])

# above problem can be solved simply by
n_trop = reviews.description.map(lambda desc: "tropical" in desc).sum()
n_fruity = reviews.description.map(lambda desc: "fruity" in desc).sum()
descriptor_counts = pd.Series([n_trop, n_fruity], index=['tropical', 'fruity'])

# rating problem
def add_ratings(row):
    rating = 1
    if row.points > 94:
        rating = 3
    elif row.points > 84:
        rating = 2
    
    if row.country == "Canada":
        rating = 3
    
    row["rating"] = rating
    return row

reviews_rat = reviews.apply(add_ratings,axis="columns")
star_ratings = reviews_rat.rating
```

**value_counts()** replication

```
reviews.groupby('points').points.count()
```
To get cheapest wine in each point value category

```
reviews.groupby('points').price.min()
```
To get the first restraurant review in that country

```
reviews.groupby('country').apply(lambda df: df.title.iloc[0]) 
# here apply gets df for each group
```

To get best wine by country and province
```
reviews.groupby(['country','province']).apply(lambda df: df.loc[df.points.idxmax()])
```

**agg()** lets us run a bunch of different functions on our DataFrame simulaneously.

```
reviews.groupby(['country']).price.agg([len,min,max])
```

Multi-Indexes

```
countries_reviewed = reviews.groupby(['country','province']).description.agg([len])

		                        len
country	    province	
Argentina	Mendoza Province	3264
            Other	            536
...	        ...	                ...
Uruguay	    San Jose	        3
            Uruguay	            24

countries_reviewed.reset_index()

	country	    province	        len
0	Argentina	Mendoza Province	3264
1	Argentina	Other	            536
...	...	...	...
423	Uruguay	    San Jose	        3
424	Uruguay	    Uruguay	            24

countries_reviewed.sort_values(by='len')
countries_reviewed.sort_values(by='len', ascending=False)
countries_reviewed.sort_index() # to sort by index values.
countries_reviewed.sort_values(by=['country','len'])
```

Who are the most common wine reviewers in the dataset? Create a Series whose index is the taster_twitter_handle category from the dataset, and whose values count how many reviews each person wrote.

```

# below codes both return same values
count = reviews.groupby('taster_name').apply(lambda df: df.taster_name.count()) 
count = reviews.groupby('taster_name').taster_name.count()

# solutions
reviews_written = reviews.groupby(['taster_twitter_handle']).taster_name.count()
reviews_written = reviews.groupby('taster_twitter_handle').size()
reviews_written = reviews.groupby('taster_twitter_handle').taster_twitter_handle.count()

```
What is the best wine I can buy for a given amount of money? Create a Series whose index is wine prices and whose values is the maximum number of points a wine costing that much was given in a review. Sort the values by price, ascending (so that 4.0 dollars is at the top and 3300.0 dollars is at the bottom).

```
best_rating_per_price = reviews.groupby(['price']).points.max().sort_index()

best_rating_per_price = reviews.groupby('price')['points'].max().sort_index()
```


```
price_extremes = reviews.groupby('variety')['price'].agg([min,max])
sorted_varieties = price_extremes.sort_values(['min','max'],ascending=False)
reviewer_mean_ratings = reviews.groupby('taster_name').points.mean()

# below both are same
country_variety_counts = reviews.groupby(['country','variety']).variety.count().sort_values(ascending=False)
country_variety_counts = reviews.groupby(['country', 'variety']).size().sort_values(ascending=False)
```

Data Type: How pandas is storing data internally. columns consisting entirely of strings do not get their own type they are of object type.
```
review.price.dtype
review.dtypes
reviews.points.astype('float64') # convert data type.
reviews.index.dtype

# missing data = NaN 
# get reviews which does not have country information
reviews[pd.isnull(reviews.country)]
reviews.region_2.fillna("unknown") # filling NaN with unknown
reviews.taster_twitter_handle.replace("@kerinokeefe","@kerino")
```
Sometimes the price column is null. How many reviews in the dataset are missing a price?

```
missing_price_reviews = reviews[reviews.price.isnull()]
n_missing_prices = len(missing_price_reviews)
# Cute alternative solution: if we sum a boolean series, True is treated as 1 and False as 0
n_missing_prices = reviews.price.isnull().sum() # false is 0 and true is 1 
# or equivalently:
n_missing_prices = pd.isnull(reviews.price).sum()
```

```
reviews_per_region = reviews.region_1.fillna("Unknown").value_counts().sort_values(ascending=False)

reviews.rename(columns={'points':'score'})
reviews.rename(index={0:'firstEntry',1:'secondEntry'})

reviews.set_index("wines")
reviews.rename_axis("wines",axis='rows').rename_axis("fields",axis='columns')
```
Combining two data frames having same data fields

```
df1 = pd.read_csv("1.csv")
df2 = pd.read_csv("2.csv")

df = pd.concat([df1,df2])
```

Combine data frame objects which are having an index in common

```
left.join(right,lsuffix="_left",rsuffix="_right")
```

```
renamed = reviews.rename(columns={"region_1":"region",'region_2':"locale"})
reindexed = reviews.rename_axis('wines',axis='rows')
combined_products = pd.concat([gaming_products,movie_products])

```

Both DFs include references to a MeetID, a unique key for each meet (competition) included in the database. Using this, generate a dataset combining the two tables into one.
```
powerlifting_combined = powerlifting_meets.set_index('MeetID').join(powerlifting_competitors.set_index("MeetID"))
```

### APACHE KAFKA 

* Is an event streaming platform to collect, store and process real time data streams at scale.
* Collection of immutable append-only logs.

##### CMD Line interface

* Install Confluent CLI 

```
confluent login --save

confluent environment list
Confluent environment use env-qd252
confluent kafka cluster list
confluent kafka cluster use lkc-dokxq7
confluent api-key create --resource lkc-dokxq7
confluent api-key use <API KEY> --resource lkc-dokxq7
confluent kafka topic list
confluent kafka topic consume --from-beginning poems

# new terminal 
confluent kafka topic produce poems --parse-key
5:"From the ashes a fire shall awaken"
6:"A light from the shadows shall spring"
```

#### Uses

* Distributed logging
* Stream processing
* Pub-Sub messaging

#### Events

* Internet of Things
* Business process change
* User interaction
* Microservice output

An event is Notification + State

Key - Value Pair

Key is an some identity in system.


#### Topics 

* The primary component of storage in Kafka is called a topic which typically represents a particular data entity.
* Topics are broken into even smaller components called **partitions**
* When we write a message to a Kafka topic that message is actually stored in one of the topics partitions. To partition to which message is routed is based on the key of that message.

* Producing to and consuming from a topic is done through console, producer API from our application or even Kafka Connect.

* Name container for events.
    - Systems contain lots of topics
    - Can duplicate data b/w topics
* Durable logs of events.
    * Append only 
    * Can only seek by offset, not indexed
* Events are immutable.
* Retention period is configurable (Can be configured to expire the message)

##### Partitioning

* Takes single topic log and breaks it into multiple logs each of which is stored in different paritions and each partition can reside in separate node in cluster.
* If message has no key, the messages are distributed using round robin equally distributed across the partitions.
* If message has a key, we run the hash on the key and mod it with the number of partitions. Same key message always lands in the same partition.
* Default paritions for a topic - **6**
* Partions are useful in allowing us to break up our topic into more managable chunks which can stored in multiple nodes in the cluster.

```
confluent kafka topic list
confluent kafka topic describe poems
confluent kafka topic create --partitions 1 poems_1
confluent kafka topic create --partitions 4 poems_4
```
To write messages

```
confluent kafka topic produce poems_1 --parse-key
5:"A light from the shadows shall spring"
8:"The crownless again shall be king"

^c

confluent kafka topic produce poems_4 --parse-key
1:"All that is gold does not glitter"
2:"Not all who wander are lost"
```

#### Brokers

* An computer, instance or container running the Kafka process (Broker process)
* Manage partitions (Each brokers has some partitions)
* Handles write and read requests
* Manage replication of partions.
* Intenstionally simeple 

#### Replication

* Copying of partitioning data to several other brokers to keep it safe. Those are called follower replicas. Main partition is Leader.
* Copies of data for fault tolerance.
* One lead partition and N-1 followers.
* In general, writes and reads happen to the leader.

#### Producers

* Kafka producer is connected to the cluster.
* In Java,
    - KafkaProducer - used to connect to the cluster
    - ProducerRecord - To hold key value which we want to send.
* Produces that takes the decision in which message has to go in which partition.

#### Consumers

* KafkaConsumer - used to connect to the cluster
* We will subscribe to the topic
* ConsumerRecords - contains individual instance of message.
* Reads messages from topics
* Reading message does not destroy the message.
* Scaling consumer groups is more or less automatic. A single instance of consuming application will always receive the all the messages from all those partitions.If there exists another consuming instance with same configuation the cluster will distribute the partitions within these instance(Rebalancing) with ordering message guarantee 

```
# python3
# pip install confluent-kafka
# in Terminal run -> confluent kafka cluster describe 
# Get Endpoint, API Key and Secret
# confluent api-key create --resource {ID}
# confluent api-key use {API Key} --resource {ID}
```
create config.ini with following content.
```
[default]
bootstrap.servers=< Endpoint >
security.protocol=SASL_SSL
sasl.mechanisms=PLAIN
sasl.username=< API Key >
sasl.password=< API Secret >

[consumer]
group.id=python_kafka101_group_1
# 'auto.offset.reset=earliest' to start reading from the beginning of
# the topic if no committed offsets exist.
auto.offset.reset=earliest
```
Create another file name **consumer.py** 
```
#!/usr/bin/env python
from argparse import ArgumentParser, FileType
from configparser import ConfigParser
from confluent_kafka import Consumer
if __name__ == '__main__':
    # Parse the command line.
    parser = ArgumentParser()
    parser.add_argument('config_file', type=FileType('r'))
    args = parser.parse_args()

    # Parse the configuration.
    config_parser = ConfigParser()
    config_parser.read_file(args.config_file)
    config = dict(config_parser['default'])
    config.update(config_parser['consumer'])

    # Create Consumer instance
    consumer = Consumer(config)

    # Subscribe to topic
    topic = "poems"
    consumer.subscribe([topic])

    # Poll for new messages from Kafka and print them.
    try:
        while True:
            msg = consumer.poll(1.0)
            if msg is None:
                print("Waiting...")
            elif msg.error():
                print("ERROR: %s".format(msg.error()))
            Else:
                # Extract the (optional) key and value, and print.
                print("Consumed event from topic {topic}: key = {key:12} value = {value:12}".format(topic=msg.topic(), key=msg.key().decode('utf-8'), value=msg.value().decode('utf-8')))
    except KeyboardInterrupt:
        pass
    finally:
        # Leave group and commit final offsets
        consumer.close()
```
Run the script

```
chmod u+x consumer.py
./consumer.py config.ini
```

#### Kafka Connect

* Data integration system and ecosystem
* Because some other systems are not Kafka
* External client process; does not run on Brokers.
* Horizontally scalable
* Fault tolerant
* Declarative

We can configure to stream data from Kafka to Elastic Search

##### Connectors

* Pluggable software component
* Interfaces to external system and to Kafka
* Also exist as runtime entities.
* Source connectors act as Producers.
* Sink connectors act as Consumers.
* Available on Confluent Hub.

##### Datagen source connector

* Can generate sample data by defining the schemas.

#### Schema Registry

* Schemas of messages of topic may evolves over a period of time.
* Standalone server process external to Kafka brokers.
* maintains a database of schemas
* HA(High Availability) deployment option available
* Consumer/Producer API component
* Defines the schema compatibility for the topic message.
* Producer API prevents the incompatible messages from being produced.
* Consumer API prevents the incompatible messages from being consumed.
* JSON, Avro and Protocol Buffers are supported seriazable formats.
* Schema IDs with schema are cached for performance in both producer and consumer side.

* With Avro, we have explicitly defined Schemas.

```
confluent kafka topic consume --value-format avro --sr-api-key <API_KEY> --sr-api-secret <API SECRET> orders
```

#### Kafka Streams

* Functional Java API
* Filtering, grouping, aggregating, joining and more (Computational API-Stream processing computations)
* Scalable, fault-tolerant state management (Persist state to local disk and to internal cluster topics)
* Scalable computation based on Consumer Groups.
* Exposes REST API
* Integrates within our services as a library.
* Runs in the context of our application
* Does not require special infrastructure

#### ksqlDB

* A database optimized for stream processing
* Runs on its own scalable, falut tolerant cluster adjacent to the Kafka Cluster.(REST API is exposed)
* Stream processing programs written SQL

```
CREATE TABLE rated_movies AS
    SELECT title, SUM(rating)/COUNT(rating) as avg_rating from ratings
    INNER JOIN movies
        ON ratings.movie_id = movies.movie_id
    GROUP BY title EMIT CHANGES;
```
* Command line interface
* REST API for application integration
* Java library
* Kafka connect integration.
* Standalone SQL powered stream processing cluster

### Credits

[Pandas](https://www.kaggle.com/learn/pandas)

[SQL](https://www.kaggle.com/learn/intro-to-sql)

[ETL Database](https://www.stitchdata.com/etldatabase/)

[Datawarehousing and ETL tools](https://www.integrate.io/blog/etl-data-warehousing-explained-etl-tool-basics/)

[Apache Kafka](https://developer.confluent.io/learn-kafka/apache-kafka/events/)