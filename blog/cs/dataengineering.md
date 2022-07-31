### ETL Concepts

* ETL(Extract, Transform, Load) is a process of data integration that encompasses three steps - Extraction, Transformation, and Loading.
* ETL systems take large volumes of raw data from multiple sources converts it for analysis, and loads that data into our warehouse.

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
### pandas

* Data analysis(Use)
* Core Objects - DataFrame and Series.

#### DataFrame

* Table
* Contains an array of individual entries.
* Each entry correspons to a row and a column


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

### Credits

[Pandas](https://www.kaggle.com/learn/pandas)
[SQL](https://www.kaggle.com/learn/intro-to-sql)