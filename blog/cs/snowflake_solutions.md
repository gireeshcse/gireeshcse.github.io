# Snowflake


* To help customers use their data by providing a flexible, scalable system while reducing the maintainance burden on IT teams.
* Software-as-a-Service

- Generally, sometimes we have spiky workloads with the need to scale up their compute resources instantly to cater for short bursts of high demand. We can take advantage of cloud for this.
- Snowflake decouples the storage and compute.

## Architecture

- Database Storage

    Reorganizes structured and unstructured data into its internally optimized, compressed, columnar format.

- Query processing

    Uses compute and memory resources provided by virtual warehouses to execute and process queries and data operations.

- Cloud services

    A collection of supporting services that coordinate activities across the plafrom such as Authentication and Access Control, Security, Infrastructure Manager, Optimizer etc..

### Database Storage.

* Organizing and tracking where data is stored for efficient retrieval.

#### Micro Partitions

Relatively small (50MB and 500MB of uncompressed data) blocks of storage that sit in the underlying cloud's provider's data store, whether that be AWS S3, Google GCS, Azure BLOB storage.

- Divide and map the incoming data into micro partitions using the ordering of the data as it is inserted/loaded.
- Compress the data
- Capture and store metadata.

This allows for finer-grain query pruning.Intead of scanning the entire table we can process only the micro partitions which are essential based on the query

* **Pruning:** The process of narrowing down a query to only read what is absolutely required to satisfy the query.

    - Based on WHERE on query and using metadata of micropartitions.
    - After identifying the micropartitions , the header of each partition is read to identify the relevant columns, futher negating the need to read any more data than is required.

Note: Metadata is like index page in the book.

Note: For semistructured data, sub-columns are used.

#### Data Clustering

Data stored in tables is sorted along natural dimensions such as data or geographic regions. This is done automatically as data is loaded into snowflake.

#### Virtual Warehouses.

* Bundle of compute resources and memory to perfrom DML activity on the data within Snowflake which includes loading data into tables.

##### Configuration

* Size
* Maximum Clusters
* Minimum Clusters
* Scaling Policy
    - The default is Standard, where performance will be optimized over saving credits.
    - In Economy, Snowflake will wait longer before choosing to provision additional clusters and will be more aggressive when looking to decommision clusters as demand for resources starts to fall.

* Auto Suspend
    - default 10 mins.
    - Also clears the local disk cache.
* Auto Resume (imp option)

The minimum and maximum values of clusters can be anything between 1 and 10.

Multi-Cluster warehouses are best suited if we have
    - workloads that can spike rapidly
    - Many concurrent users accessing the same data.

Note: Moving from a smaller to a larger virtual warehouse is referred to as scaling up while adding addition clusters is known as scaling out.

For predictable workloads without high concurrency
    - Create multiple virtual warehouses for different purposes.
        - Ex: One for data ingestion pipelines, one for our data visualization tool, one for finance department etc.


#### Caching

When a query is executed against a virtual warehouse for first time, the result set is pushed into the cache.If same query is received, it returns this result set to the user.

##### Result Cache.

- Stores the result set of every query executed in the past 24 hours.
- Note:
    - If underlying data that makes up the result set changes, this will invalidate the result set held in the cache.
    - If query contains a functions that needs to be evaluated at execution time(Non-Deterministic), then the cache cannot be used.
    - Not specific to an individual virtual warehouse.

##### Local Disk Cache

- used to hold the results of SQL queries. The data is fetched from the remote disk and cached within the local solid state disk(SSD) of virtual warehouse.

### Query Processing

The job of the query processing service is to take queries before passing them to cloud service layer.

The overall process
-   A query is executed on Snowflake and can originate from ODBC, JDBC or the web interface.
- The optimizer sees the query and first attempts to satisfy it from the result cache.
- The optimizer then employs data pruning to narrow down the location of the data to be processed and generated a query plan.
- The resources provided by the virtual warehouse are then used to scan only the data needed from the local SSD cache or database storage before processing and retrieving the results.
- Finally, the result is processed, returned to the user, and popped into the result cache for future use.

### Cloud Services.

* Authentication(Snowflake plaform)
    - Single sign-on 
        - SAML 2.0 compliant identity provider
        - Key-pair authentication
        - Simple username and password.

* Infrastructure management

    - Management of infrastructure including storage buckets, provisioning and decommissioning new clusters to support virtual warehouses.

* Metadata management
    - Collects metadata when various operations are carried out on the snowflake platform
        - Who is using what data?
        - Which users are logged in?
        - When was the table last loaded? is it being used?
        - Which columns are important in a table and which columns are least used.
        - Where is table data stored along with clustering information?

    - Exposed in layer of system-defined views and some table functions
    - The query optimizer relies on the statistics of tables data changes when generating query plan.

* Query parsing and optimization

    - Parses the query to ensure there are no syntactical errors and manages the query execution. 
    - Relies upon the resources provided by the query processing layer to execute.

* Access control

    - To ensure that users can only access or carry out operations on the objects and data they are permitted to access, basd on their privileges.
    - RBAC

## Data Movement

### Stages

- A Stage is an area to rest our data files prior to loading them into a table.
- Types
    - External
    - Internal

#### External Stages.

- It is pointer to external third-party cloud storage location. 
- Can be (Regardless of what platform we run our snowflake account on.)
    - Amazon S3
    - Google Cloud Storage
    - Microsoft Azure

```
CREATE STAGE "BUILDING_SS"."PUBLIC".S3_STAGE URL= s3://arn CREDENTIALS = (AWS_KEY_ID = '****' AWS_SECRET_KEY= '*****' )
```

Within our external stage we can have external tables, which holds metadata that tells snowflake where to locate the data files that relate to the table.This approach allows data to sit outside of snowflake but appears to users as if it resides within Snowflake.

* Data Lake
    - Used to store high volumes of raw data for cost effectiveness, while only a subset of high quality, refined data is loaded into snowflake.
    - Remains as the source of truth.
    - Allows us to expose a subset of key business data and model it into a format to support the business needs.
    - Can act as a layer of protection against 'vendor lock-in" 

#### Internal Stages

* Focus on data within Snowflake.
- Types
    - user
        * Allocated to each user
        * Used only one user needs to access the staging data before loading it into multiple target tables.
        * Refered using prefix : **@~**
    - table
        * Each table has a table stage associated with it. 
        * Multiple users can access the data files, but these files can only load into one target table
        * **@%**
    - named
        * Multipe users can access data in this stage and data files can also target multiple tables.
        * **@**

### File Formats

* Provides information to snowflake on how to interpret an incoming file.

```
CREATE [OR REPLACE ] FILE FORMAT [IF NOT EXISTS] <name>
                     TYPE = { CSV |JSON | AVRO | ORC | PARQUET | XML}
                     [ formatTypeOptions ]
                     [ COMMENT = '<string_literal>']
```

* Commonly-Used Parameters.

    - Compression
        - GZIP
    - RECORD_DELIMITER
        - ,
    - SKIP_HEADER
        - 1
    - TRIM_SPACE
        - TRUE
    - ERROR_ON_COLUMN_COUNT_MISMATCH
        - Whether to fail the load if the column count from the data file doesn't match the target table.
        - TRUE
    - Encoding
        - UTF8

### COPY INTO

* To load staged data into a table.

```
COPY INTO [<namespace>.]<table_name>
     FROM { internalStage | externalStage | externalLocation }
[ FILES ( '<file_name>' [, '<file_name>'] [, ...] ) ]
[ PATTERN = '<regex_pattern'> ]
[ FILE_FORMAT = ({ FORMAT_NAME = '[<namespae>.]<file_format_name>' |
                TYPE = { CSV |JSON | AVRO | ORC | PARQUET | XML}
                })]
```

* Namespace

    * The database and/or schema of the table
    * EDW.sales

* table_name

    * Target table to load the data into
    * sales_transactions

* From
    * s3://raw/sales_system/2021/05/11/

* Files (Optional)
    * sales_01.csv

* Pattern (Optional)
    * *.csv (loads all the csv files within the given location)

* File_Format (Optional)
    * ff_csv

* TYPE (Optional)
    - Specifies the type of files to load into the table
    - GZIP

* VALIDATION_MODE

    - Allows us to validate the data file without the time it takes to load the data into Snowflake first.
    - Useful for testing out staged data files during the early phases of development. 
    - It is possible to validate a specified number of rows or the entire file.
    - RETURN_10_ROWS

#### Transformations

* COPY INTO supports some basic, lightweight transformations which can simplify our downstream ETL operations.
* Possible to load a subset of table data and reorder, alias, cast, or concatenate columns 

```
COPY INTO home_sales(city,zip, sale_date, price)
    FROM (SELECT t.$1, t.$2, t.$6, t.$7 FROM @mystage/sales.csv.gz t)
    file_format = (format_name = mycsvformat);

COPY INTO home_sales(city,zip, sale_date, price, full_name)
    FROM (SELECT substr(t.$1,4), t.$2, t.$6, t.$7, concat(t.$8,t.$9) FROM @mystage/sales.csv.gz t)
    file_format = (format_name = mycsvformat);
```

### Data Loading Considerations

#### File Preparation

* 100-250 MB compressed files(Recommended)
* By spliting files, we can take snowflake's abiltiy to load data in parallel.

Note: The number and capacity of servers in a virtual warehouse influences how many files we can load in parallel.

* When working with CSV files, it is better to standardize some of that data before loading into Snowflake such as
    - Encoding the file in UTF-8 format
    - Escaping any single or double quotes in the data
    - Selecting uniqure delimiters 

#### Semistructured Data

* All semistructured data in snowflake is held in a VARIANT column in the target table.
* STRIP_NULL_VALUES

### Dedicated Virtual Warehouses

* It is common to have dedicated VWs to supply the resources for our data pipelines and bulk data loads
* Configuring multiple dedicated warehouses to line up with bussiness groups, departments, or applications protects these processes from fighting for same resources.

### Partitioning Staged Data

* It is good practice to partition data in the external cloud storage locations. 
* Separating the data into a logical structures using file paths.

### Loading Data

* When we load data into snowflake using COPY INTO, metadata is tracked and stored for 64 days.The metadata helps in preventing loading the same data files. To overide this behavior use **FORCE** copy option.

Note: Snowflake calculates and stores an MD5 checksum for all the files in the metadata. This is used subsequently to prevent the reloading of duplicate files within the 64-day window.

### Unloading Data from Snowflake.

- Exporting data from Snowflake using a process known as unloading.
- CSV and JSON
- We can unload them to an internal or external stage.

```
COPY INTO @ext_stage/result/data_
FROM 
(
    SELECT t1.column_a, t1.column_b, t1.column_c
    FROM table_one t1
    INNER JOIN table_two t2 on ti.id = t2.id
    WHERE t2.column_c = '2022-04-01'
)
```

### Continuous Data Loads using Snowpipe

* **Snowpipe** is a fully managed service designed to quickly and effiently move smaller amounts of frequently arriving data into a Snowflake table from a stage.

Note: We must create files first, we cannot stream data directly into Snowflake.

* We stream our data into an external stage, such as S3. S3 calls Snowpipe using an event notification to tell Snowpipe there is new data to process.
* The data is then pushed into a queue(SQS queue). The Snowpipe service then consumes and processes data from this queue and pushes the data to a target table in Snowflake.
* We can use REST API calls to invoke Snowpipe's API, which places the files in a queue before loading them to the target table.

* Snowpipe is serverless. No need to configure virtual warehouses. Snowflake takes care of this for us.**We only pay for the compute time used to load data.**

### Streams

