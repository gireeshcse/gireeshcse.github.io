### ETL Concepts

* ETL(Extract, Transform, Load) is a process of data integration that encompasses three steps - Extraction, Transformation, and Loading.
* ETL systems take large volumes of raw data from multiple sources converts it for analysis, and loads that data into our warehouse.

#### Extraction

* Extracted data sets come from a source into a staging area
* Staging area acts as a buffer b/w the data warehouse and he source data.
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
* **iloc** uses indexes inclusively (0:10 => 0,1...,9,10 )

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
```