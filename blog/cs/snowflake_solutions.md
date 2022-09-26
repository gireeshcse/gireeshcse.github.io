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


### Change Tracking Using Streams 

A stream object tracks any DML operations against the source table.This process is known as Change Data Capture.

* Change Data Capture (CDC) help streamline the movement of data.
    - If we have very large transaction table in our source system containing millions and millions of records, with 1000 new transactions being added per day, with no records change historically, we wouldn't want to reload the entire table each day. Using CDC against this table allows us to identify and target those 1000 records for extraction. This makes the entire process faster and more efficient.

* In  Snowflake, when we create a stream
    - a pair of hidden columns are added to the stream and they begin to store change tracking metadata.
    - A snapshot of source table is logically created which acts as baseline for all subsequent changes on the data can be identified.This baseline is referred as offset.
    - The stream also creates a change table, which stores both the before and after record between two points in time.This change table mirrors the structure of the source table along with the addition of some very handy change tracking columns.we can now point our processes to this change table and process the changed data.
    - Multiple queries can query the changed records from a stream without changing the offset. It is important to note that the offset is only moved when stream records are used within a DML transaction. Once the transaction commits, the offset moves forward, so we cannot reread thee same record again.

#### Stream Metadata Columns

These help us understand the nature of the changes applied to the source table, so we can process them correctly.

* METADATA$ACTION

    - Tells us what DML action was performed(INSERT or DELETE)
    - An update is effectively a DELETE followed by and INSERT

* METADATA$ISUPDATE

    - A Boolean value that indicated if the records were part of an UPDATE operation. When TRUE we would expect to see a pair of records, one with a DELETE and one with an INSERT

* METADATA$ROW_ID

    - A unique ID for the row.
    - Helpful for row level logging and auditability.

### Tasks

* Tasks execute a single SQL statement, giving us the ability to chain together a series of tasks and dependencies so we can execute them as required.

* A typical pattern is to have a task running every 10 minutes, which checks the stream for the presence of records to be processed using the system function **SYSTEM$STREAM_HAS_DATA('<stream_name>')**. If it returns FALSE, there are no records to process and task will exit.

* The task contains SQL logic that applies transformations or uses stored procedures or user-defined functions before merging those changes into a target table.

* We must specify a warehouse to use when we create a Task.
* Tasks require to have only one parent task, known as the root task. The root task must also have schedule associated with it. This can be a duration such as 5 minutes or a CRON expression.

A simple chain of tasks

![alt a Simple chain of tasks](images/snowflake_task_chain.png)

* Child tasks can then be chained together to execute by using the CREATE TASK ... AFTER and specifying the name of the preceding task.

```
CREATE [ OR REPLACE ] TASK [IF NOT EXISTS ] <name>
WAREHOUSE = <string>
[ SCHEDULE = '{ <num> MINUTE | USING CRON <expr> <timezone>}' ]
[ ALLOW_OVERLAPPING_EXECUTION = TRUE | FALSE ]
[ <session_parameter> = <value> [ ,<session_parameter> = <value> ...]]
[ USER_TASK_TIMEOUT_MS = <num>]
[COPY GRANTS]
[ COMMENT = 'string_literal>' ]
[ AFTER <string> ]
[ WHEN <boolean_expr>]
AS
    <sql>

```
High-Level diagram 
![alt a Simple chain of tasks](images/snowflake_streams_tasks_example.png)

```
--CREATE DATABASE

CREATE OR REPLACE DATABASE nse;

--SWITCH CONTEXT
USE DATABASE nse;

--CREATE SCHEMAS
CREATE SCHEMA STG;
CREATE SCHEMA CDC;
CREATE SCHEMA TGT;

--CREATE SEQUENCE
CREATE OR REPLACE SEQUENCE SEQ_01
START = 1
INCREMENT = 1;

--CREATE STAGING TABLE
CREATE OR REPLACE TABLE STG.EQ_MAIN
(
    eq_id VARCHAR(100),
    part_of_fo NUMBER(1),
    part_of_index NUMBER(1),
    face_value int,
    industry VARCHAR(100),
);

--CREATE TARGET TABLE
CREATE OR REPLACE TABLE TGT.EQ_MAIN
(
    eq_seq_key int default SEQ_01.nextval,
    eq_id VARCHAR(100),
    part_of_fo NUMBER(1),
    part_of_index NUMBER(1),
    face_value NUMBER(10),
    industry VARCHAR(100),
    date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
);

--Creating a stage
CREATE STAGE "nse"."PUBLIC".S3_STAGE URL = 's3://' CREDENTIALS = (AWS_KEY_ID = '' AWS_SECRET_KEY = '');

--GRANT PERMISSIONS ON STAGE
GRANT USAGE ON STAGE S3_STAGE TO SYSADMIN;

--SHOW STAGES
SHOW STAGES;

--UNLOAD DATA TO S3 EXTERNAL STAGE
COPY INTO @S3_STAGE/Customer FROM "SNOWFLAKE_SAMPLE_DATA"."TPCH_SF1"."CUSTOMER"
HEADER=TRUE;

--COPY INTO TABLE
COPY INTO STG.EQ_MAIN (eq_id, part_of_fo,part_of_index, face_value, industry)
FROM (SELECT $1, $2, $3, $4,$5 FROM @S3_STAGE/)
FILE_FORMAT=(TYPE = 'CSV' FIELD_DELIMITER = ',' SKIP_HEADER = 1 COMPRESSION = 'GZIP');

--CONFIRM WE HAVE SAME RECORDS AS IN CSV FILES
SELECT COUNT(*) FROM STG.EQ_MAIN;

--SEED TABLE
INSERT INTO TGT.EQ_MAIN (eq_id, part_of_fo,part_of_index, face_value, industry)
SELECT eq_id, part_of_fo,part_of_index, face_value, industry
FROM STG.EQ_MAIN;

--CREATE STREAM TO TRACK SUBSEQUENT CHANGES
CREATE OR REPLACE STREAM CDC.EQ_MAIN
ON TABLE STG.EQ_MAIN;

--SHOW STREAMS
SHOW STREAMS;

--CHECKING CHANGE TABLE FOR METADATA COLUMNS
SELECT * FROM CDC.EQ_MAIN;


--DML ON STAGE TABLE
UPDATE STG.EQ_MAIN
SET face_value = 100
WHERE eq_id = 'TCS';

--TO SEE TWO RECORDS, A DELETE AND AN INSERT IN CDC WITH VALUE METADATA$ISUPDATE = TRUE
SELECT * FROM CDC.EQ_MAIN;


--TO CHECK IT THERE ARE RECORDS IN THE STREAM WAITING TO BE PROCESSED(RETURNS BOOLEAN)
SELECT SYSTEM$STREAM_HAS_DATA('CDC.EQ_MAIN');

--CREATE A TASK TO MERGE THESE CHANGES 
CREATE OR REPLACE TASK CDC.MERGE_EQ_MAIN
WAREHOUSE = COMPUTE_WH_FINANCE
SCHEDULE = '5 minute'
WHEN 
    SYSTEM$STREAM_HAS_DATA('CDC.EQ_MAIN')
AS
MERGE INTO TGT.EQ_MAIN TGT
USING CDC.EQ_MAIN CDC
ON TGT.eq_id = CDC.eq_id

WHEN NOT MATCHED AND METADATA$ACTION = 'INSERT' AND METADATA$ISUPDATE = 'FALSE' THEN
INSERT (eq_id, part_of_fo,part_of_index, face_value, industry) VALUES (eq_id, part_of_fo,part_of_index, face_value, industry)

WHEN MATCHED AND METADATA$ACTION = 'INSERT' AND METADATA$ISUPDATE = 'TRUE' THEN
UPDATE SET TGT.part_of_fo = CDC.part_of_fo, TGT.part_of_index = CDC.part_of_index, TGT.face_valUE = CDC.face_value,TGT.industry = CDC.industry 

WHEN MATCHED AND METADATA$ACTION = 'DELETE' AND METADATA$ISUPDATE = 'FALSE' THEN DELETE;

--BY DEFAULT A TASK IS SET UP IN SUSPEND MODE
SHOW TASKS;

--ENSURE SYSADMIN CAN EXECUTE TASKS
USE ROLE accountadmin;
GRANT EXECUTE TASK ON ACCOUNT TO ROLE SYSADMIN;

--WE NEED TO RESUME THE TASK TO ENABLE IT
ALTER TASK CDC.MERGE_EQ_MAIN RESUME;

--TO MONITOR TASK_HISTORY
SHOW TASKS;
SELECT * FROM TABLE(INFORMATION_SCHEMA.TASK_HISTORY(TASK_NAME=>'MERGE_EQ_MAIN'));

--CHECKING TARGET TABLE
SELECT * FROM TGT.EQ_MAIN
WHERE eq_id = 'TCS';

--INSERT A NEW RECORD TO STAGING TABLE
INSERT INTO STG.EQ_MAIN (eq_id, part_of_fo,part_of_index, face_value, industry)
SELECT 'INFY',1,1,10,'IT';

--CHECKING TARGET TABLE
SELECT * FROM TGT.EQ_MAIN
WHERE eq_id = 'INFY';

DELETE FROM STG.EQ_MAIN WHERE eq_id = 'INFY';

DROP DATABASE nse;
ALTER WAREHOUSE COMPUTE_WH_FINANCE SUSPEND;
```

## CLONING

* Cloning is a metadata-only operation, meaning no data movement or additional storage is required. We can create a fully working test database, complete with data, in minutes, all with zero cost!

* Time effiency and cost savings are at the heart of what cloning is all about.

* Eases the creation of dev or test environments.
* We can use for experimental purposes.
    - Ex: Creating new features/data


### Performance Testing

* We can use 50% of scaled down of production environment for testing

#### Testing with Data

* Generally, generating data for non-production environments was almost a project in itself.

Testing a UAT environment against production data usually involves either restoring a backup of the production environment to UAT env or using ETL tool to extract data from production to populate the UAT environment.

#### Time Travel Feature.

* Snowflake automatically keeps a record of data(tables) versions in the background
* User-configuratble maximum retention period of up to 90 days for permannent objects. 
* Default is 24 hours

```
ALTER TABLE TABLE1 SET DATA_RETENTION_TIME_IN_DAYS = 60;

SELECT * FROM TABLE AT(timestamp => '2022-09-22 13:00:10.00 +0530'::timestamp_tz)
```

* No need to take a backup or snapshot.
* This feature is leveraged by cloning to replicate an object

#### Sensitive Data

* Risks are involved when using cloning to move data between environments.
    - Sensitive information (PII)
    - Ensuring General Data Protection Regulation(GDPR) and Health Insurance Portability and Accountability Act(HIPPA)

* Solution - Data masking in conjunction with cloning.

* **Data Masking**

    - Column-level security feature that uses policies to mask plain-text data dynamically depending upon the role accessing the data.
    - At query execution, the masking policy is applied to the column at every location where column appears.

### Working with Clones

* Cloned objects are also writable.Additional storage is needed to store these versions.(Performed Background)

```
CREATE OR REPLACE TABLE ClonedTable Clone Table; 
```
* List of objects that can be cloned.
    - Databases
    - Schemas
    - Tables
    - Streams
        - Any unconsumed records are not available for clone
        - Streams begins again at the time/point the clone was created
    - Extenal named stages
    - File formats
    - Sequences
        - when a table with a column with a default sequence is cloned, the cloned table still references the original sequence object.

        ```
        ALTER TABLE <table_name>
        ALTER COLUMN <column_name> SET DEFAULT <new_sequence>.nextval;
        ```
    - Tasks
        - Cloned tasks are suspended by default.we must execute ALTER TASK...RESUME

    - Pipes
        - A database or schema clone includes only pipe objects that reference external stages(AWS S3, GCP etc)

* The following objects cannot be cloned

    - Internal Stages
    - Pipes 
        - Internal pipes
        - cloned pipe is paused by default.
    - Views
        - Cannot be cloned directly
        - Cloned if it is contained within a database or schema we are cloning

#### Clone Permissions.

* Roles are not cloned with source database. Only the child objects of the original database (schemas and tables) would carry over their privileges to the cloned db.

```
CREATE OR REPLACE DATABASE DEMO;

GRANT USAGE ON DATABASE DEMO TO SYSADIM;

CREATE OR REPLACE DATABASE DEMO_CLONE CLONE DEMO;

--USAGE GRANT WILL NOT APPEAR FOR DEMO_CLONE
SHOW GRANTS ON DATABASE DEMO_CLONE;


CREATE OR REPLACE TABLE DEMO.PUBLIC.TABLE1 (COL VARCHAR);

GRANT SELECT ON TABLE ON DEMO.PUBLIC.TABLE1.SYSADMIN;

--RE-CREATE 
CREATE OR REPLACE DATABASE DEMO_CLONE CLONE DEMO;

--SHOWS THE GRANTS ON TABLE1 - PERMISSIONS PRESERVED
SHOW GRANTS ON TABLE DEMO_CLONE.PUBLIC.TABLE;

CREATE OR REPLACE TABLE DEMO.PUBLIC.TABLE_CLONE CLONE DEMO.PUBLIC.TABLE1;

-- SELECT PRIVILEGE DOESN'T CARRY OVER
SHOW GRANTS ON TABLE DEMO.PUBLIC.TABLE_CLONE;
```

#### Example

```
CREATE OR REPLACE DATABASE SALES_DEV;

CREATE TABLE ORDERS AS 
SELECT * FROM "SNOWFLAKE_SAMPLE_DATA"."TPCH_SF1".ORDERS;

CREATE OR REPLACE DATABASE SALES_PROD CLONE SALES_DEV;

USE DATABASE SALES_DEV;
UPDATE ORDERS
SET O_TOTALPRICE = O_TOTALPRICE * 1.1;

--AFTER VALIDATING THE CHANGES WE CAN PROMOTE THE CHANGE TO PRODUCTION
USE DATABASE SALES_PROD;
CREATE OR REPLACE TABLE ORDERS CLONE SALES_DEV.PUBLIC.ORDERS;


--FINAL CHECKS BETWEEN SALES_DEV AND SALES_PROD
SELECT DEV.O_ORDERKEY, DEV.O_TOTALPRICE, PROD.O_TOTALPRICE
FROM SALES_DEV.PUBLIC.ORDERS DEV
INNER JOIN SALES_PROD.PUBLIC.ORDERS PROD ON DEV.O_ORDERKEY =
PROD.O_ORDERKEY
LIMIT 10;

SELECT DEV.O_ORDERKEY, DEV.O_TOTALPRICE - PROD.O_TOTALPRICE AS
DIFFERENCE
FROM SALES_DEV.PUBLIC.ORDERS DEV
INNER JOIN SALES_PROD.PUBLIC.ORDERS PROD ON DEV.O_ORDERKEY =
PROD.O_ORDERKEY
HAVING DIFFERENCE <> 0;

```

* Address the following challenges

    - To avoid the need to pay for storing the same data more than once.
    - The time-consuming task of copying data to set up new environments for development and testing.
    - As a way to keep production and non-production environments in sync

### Advantages

- We can create a clone in seconds as it’s a metadata-only operation.
- We can update data in a clone independently of the source object.
- We can promote changes to a Production environment quickly and easily, with low risk.

## Managing Security and Access Control

### Roles

* Collection of privileges against one or more objects within our Snowflake account.
* Predefined system roles
    - ACCOUNTADMIN
        - administers Snowflake account.
        - can view all credit and billing information
        - Sub Roles
            - SYSADMIN
            - SECURITYADMIN
        - Configuring multi-factor authentication is strongly recommended for this account users.
        - Snowflake recommends limiting the use of this account and restricting access to a minimum set of users.
    - SYSADMIN
        - Can create warehouses and databases and objects within a database.
    - SECURITYADMIN
        - designed for the administration of security
        - management of granting or revoking privileges to roles
    - USERADMIN
        - For creating roles and users and managing the privileges assigned to them.
        - Child of SECURITYADMIN
    - PUBLIC
        - Default role that all the users end up in automatically.
        - Provides privileges to log into Snowflake and some basic object access

Note: In Snowflake, a user may be member of more than just one role. He has to switch to the role for which an action/query is having access/permission.

* Snowflake follows RBAC framework.

```
GRANT ROLE HR_SCHEMA_READ_ONLY TO ROLE MARKETING_INSIGHT;
REVOKE ROLE HR_SCHEMA_READ_ONLY FROM ROLE MARKETING_INSIGHT;
```

#### Logically Separating Out Roles

* Level 0
    - System Roles
* Level 1
    - Domain Roles
    - Used when we require groups of Level 2 and 3 roles separated for different environments, such as Dev, Test and Production.
* Level 2
    - Functional Roles
    - Used to contain users.This may be mapped to roles within an identity provider which contains groups of users.
* Level 3
    - Access Roles
    - Used to assign privileges on the actual objects within Snowflake.

### User and Application Authentication.


#### Multi-Factor Authentication

* An additional layer of security by requiring a user to provide a secondary form of authentication in addition to the standard login.


* MFA is recommended for ACCOUNTADMIN role.
* Bypassing MFA for a user for 60 minutes

```
ALTER USER MY_USER_NAME
SET MINS_TO_BYPASS_MFA = 60;
```

* Disabling MFA for a user (requires the user to re-enroll):

```
ALTER USER MY_USER_NAME
SET DISABLE = TRUE;
--TO verify MFA is disabled
DESCRIBE USER MY_USER_NAME;
```

#### Security Assertion Markup Language (SAML)

* An open standard which identity providers can use to provide a single
sign-on experience.

#### OAuth

* OAuth is an open standard to support authentication.

#### Key Pair Authentication

* Works on the basis of private and public keys

### Network Policies

By default, Snowflake allows users or applications from any IP address to access the service. To provide an additional layer of security and allow or deny specific IP addresses access to Snowflake, we can use network policies.

```
CREATE [ OR REPLACE ] NETWORK POLICY <name>
   ALLOWED_IP_LIST = ( '<ip_address>' [ , '<ip_address>' , ... ] )
   [ BLOCKED_IP_LIST = ( '<ip_address>' [ , '<ip_address>' , ... ] ) ]
   [ COMMENT = '<string_literal>' ]

ALTER ACCOUNT SET NETWORK_POLICY = <policy name>
```

### Handling PII Data

* Separately Storing PII Data
    - Separating the PII data away, we reduce the risk in breaking our data model or removing data unnecessarily in the future.
    - We should also consider your Time Travel settings in our PII data in Snowflake. Depending on our settings, Snowflake may have a version of this record stored beyond 30 days, which could allow this data to be mistakenly restored after it’s been removed

#### Auditing

* Provides greater control and transparency
* It is recommended to have metadata table to support the monitoring.
    - A unique identifier for the individual (obviously something that doesn’t identify the customer!)
    - The date the delete request was received
    - The date it was flagged for removal in the database
    - The date it was deleted

#### Controlling Access to PII Data

* We will need to create a minimum of two access roles
    - one to access sensitivedata and one that cannot

```
//ENSURE SYSADMIN IS USED
USE ROLE SYSADMIN;

//CREATE DATABASE
CREATE OR REPLACE DATABASE PII_DEMO;

//CREATE SCHEMAS
CREATE OR REPLACE SCHEMA SALES;
CREATE OR REPLACE SCHEMA CUSTOMER_PII;
CREATE OR REPLACE SCHEMA PRESENTATION;

//CREATE SALES TABLE TO STORE NON PII DATA
CREATE OR REPLACE TABLE SALES.ORDERS
(ORDER_ID INT,
CUSTOMER_ID INT,
ORDER_DATE TIMESTAMP_TZ);

//CREATE CUSTOMER TABLE TO STORE PII (CUSTOMER_EMAIL)
CREATE OR REPLACE TABLE CUSTOMER_PII.CUSTOMER
(CUSTOMER_ID INT,
CUSTOMER_EMAIL VARCHAR(50));

//INSERT SOME SAMPLE RECORDS TO BOTH TABLES
INSERT INTO SALES.ORDERS
SELECT 1, 1, '2021-06-07 12:21:10.00'
UNION
SELECT 2, 1, '2021-06-10 14:21:10.00';
INSERT INTO CUSTOMER_PII.CUSTOMER
SELECT 1, 'customer1@gmail.com';

//CREATE NORMAL UNSECURED VIEW
CREATE OR REPLACE VIEW PRESENTATION.CUSTOMER_LAST_ORDER_DATE
AS
SELECT C.CUSTOMER_ID,
        C.CUSTOMER_EMAIL,
        MAX(O.ORDER_DATE) AS MOST_RECENT_ORDER_DATE
FROM CUSTOMER_PII.CUSTOMER C
INNER JOIN SALES.ORDERS O on C.CUSTOMER_ID = O.CUSTOMER_ID
GROUP BY C.CUSTOMER_ID, C.CUSTOMER_EMAIL;

//INTRODUCE OBFUSCATION
CREATE OR REPLACE VIEW PRESENTATION.CUSTOMER_LAST_ORDER_DATE_PII
AS
SELECT C.CUSTOMER_ID,
        CASE WHEN CURRENT_ROLE() <> 'SENSITIVE_ALLOWED_ROLE' --NOTE THE USE OF CURRENT ROLE HERE
                THEN 'XXX-XX-XXXX' --THIS IS WHAT USERS OUTSIDE OF THE SENSITIVE_ALLOWED_ROLE WILL SEE RETURNED
                 ELSE CUSTOMER_EMAIL
            END AS CUSTOMER_EMAIL,
        MAX(O.ORDER_DATE) AS MOST_RECENT_ORDER_DATE
FROM CUSTOMER_PII.CUSTOMER C
INNER JOIN SALES.ORDERS O on C.CUSTOMER_ID = O.CUSTOMER_ID
GROUP BY C.CUSTOMER_ID, C.CUSTOMER_EMAIL;


//FINALLY TO PREVENT USERS BEING ABLE TO INFER THE UNDERLYING DATA STRUCTURES
//CREATE THE VIEW AS A SECURE VIEW
CREATE OR REPLACE SECURE VIEW PRESENTATION.CUSTOMER_LAST_ORDER_DATE_PII_SECURE
AS
---

//GRANT PERMISSIONS ON WAREHOUSE AND DATABASE TO SECURITYADMIN
GRANT USAGE ON WAREHOUSE COMPUTE_WH TO SECURITYADMIN;
GRANT USAGE ON DATABASE PII_DEMO TO SECURITYADMIN;

//SWITCH TO USE SECURITYADMIN ROLE TO CARRY OUT GRANTS AND CREATE ROLES
USE ROLE SECURITYADMIN;

//CREATE NEW ROLES
CREATE OR REPLACE ROLE SENSITIVE_ALLOWED_ROLE; --PII DATA ALLOWED
CREATE OR REPLACE ROLE SENSITIVE_DENIED_ROLE; --NO ACCESS TO PII DATA

//GRANT PERMISSIONS ON OBJECTS TO BOTH ROLES
GRANT USAGE ON WAREHOUSE COMPUTE_WH TO SENSITIVE_ALLOWED_ROLE;
GRANT USAGE ON WAREHOUSE COMPUTE_WH TO SENSITIVE_DENIED_ROLE;
GRANT USAGE ON DATABASE PII_DEMO TO SENSITIVE_ALLOWED_ROLE;
GRANT USAGE ON DATABASE PII_DEMO TO SENSITIVE_DENIED_ROLE;
GRANT USAGE ON SCHEMA CUSTOMER_PII TO SENSITIVE_ALLOWED_ROLE;
GRANT USAGE ON SCHEMA SALES TO SENSITIVE_ALLOWED_ROLE;
GRANT USAGE ON SCHEMA PRESENTATION TO SENSITIVE_ALLOWED_ROLE;
GRANT USAGE ON SCHEMA CUSTOMER_PII TO SENSITIVE_DENIED_ROLE;
GRANT USAGE ON SCHEMA SALES TO SENSITIVE_DENIED_ROLE;
GRANT USAGE ON SCHEMA PRESENTATION TO SENSITIVE_DENIED_ROLE;
GRANT SELECT ON VIEW  PRESENTATION.CUSTOMER_LAST_ORDER_DATE TO SENSITIVE_ ALLOWED_ROLE;
GRANT SELECT ON VIEW  PRESENTATION.CUSTOMER_LAST_ORDER_DATE TO SENSITIVE_DENIED_ROLE;
GRANT SELECT ON VIEW  PRESENTATION.CUSTOMER_LAST_ORDER_DATE_PII TO SENSITIVE_ALLOWED_ROLE;
GRANT SELECT ON VIEW  PRESENTATION.CUSTOMER_LAST_ORDER_DATE_PII TO SENSITIVE_DENIED_ROLE;
GRANT SELECT ON VIEW  PRESENTATION.CUSTOMER_LAST_ORDER_DATE_PII_SECURE TO   SENSITIVE_ALLOWED_ROLE;
GRANT SELECT ON VIEW  PRESENTATION.CUSTOMER_LAST_ORDER_DATE_PII_SECURE TO SENSITIVE_DENIED_ROLE;

//ADD BOTH THESE ROLES TO YOUR OWN USER TO MAKE IT EASY TO TEST OUT FOR THIS DEMO
GRANT ROLE SENSITIVE_ALLOWED_ROLE TO USER RAM;
GRANT ROLE SENSITIVE_DENIED_ROLE TO USER RAM;

//SWITCH CONTEXT TO SYSADMIN
USE ROLE SYSADMIN;

//OBSERVE RESULTS FROM NORMAL VIEW
SELECT * FROM PRESENTATION.CUSTOMER_LAST_ORDER_DATE;

//CHANGE THE CONTEXT OF THE ROLE
USE ROLE SENSITIVE_ALLOWED_ROLE;
SELECT CURRENT_ROLE();

//OBSERVE THE FACT ALL VALUES ARE RETURNED INCLUDING CUSTOMER EMAIL
SELECT * FROM PRESENTATION.CUSTOMER_LAST_ORDER_DATE_PII;

//SWITCH THE CONTEXT OF THE ROLE TO USERS WHO CANNOT VIEW PII
USE ROLE SENSITIVE_DENIED_ROLE;
SELECT CURRENT_ROLE();
//SELECT FROM THE VIEW AGAIN AND NOTE THE VALUE OF THE CUSTOMER EMAIL IS NOW MASKED
SELECT * FROM PRESENTATION.CUSTOMER_LAST_ORDER_DATE_PII;

//VIEW THE DDL FOR THE VIEW
SELECT GET_DDL('VIEW', 'CUSTOMER_LAST_ORDER_DATE_PII_SECURE', TRUE);
//CHANGE THE CONTEXT OF THE ROLE
USE ROLE SENSITIVE_ALLOWED_ROLE;
//OBSERVE THE FACT ALL VALUES ARE RETURNED INCLUDING CUSTOMER EMAIL
SELECT * FROM PRESENTATION.CUSTOMER_LAST_ORDER_DATE_PII_SECURE;
//ADDITIONALLY TRY AND VIEW THE DDL OF THE SECURE VIEW
SELECT GET_DDL('VIEW', 'CUSTOMER_LAST_ORDER_DATE_PII_SECURE', TRUE);
//SWITCH THE CONTEXT OF THE ROLE TO USERS WHO CANNOT VIEW PII
USE ROLE SENSITIVE_DENIED_ROLE;

//SELECT FROM THE VIEW AGAIN AND NOTE THE VALUE OF THE CUSTOMER EMAIL
SELECT * FROM PRESENTATION.CUSTOMER_LAST_ORDER_DATE_PII_SECURE;
//ADDITIONALLY TRY AND VIEW THE DDL OF THE SECURE VIEW
SELECT GET_DDL('VIEW', 'CUSTOMER_LAST_ORDER_DATE_PII_SECURE', TRUE);
```

#### Dynamic Data Masking

* Available in Eneterprise Edition and above.
* Allows us to create a dynamic masking policy as an object in the database
* This feature helps us control who can create (and modify) masking policies while centralizing the rules in one place.
* The masking policy is applied dynamically to any queries that reference the columns with a policy attached, so it’s very powerful.

```
//SWITCH ROLE
USE ROLE SYSADMIN;

//CREATE MASKING POLICY
CREATE OR REPLACE MASKING POLICY EMAIL_MASK AS (VAL STRING) RETURNS
STRING ->
  CASE
    WHEN CURRENT_ROLE() <> ('SENSITIVE_ALLOWED_ROLE') THEN 'XXX-XX-XXXX'
    ELSE VAL
  END;

//ATTACH THE MASKING POLICY TO THE CUSTOMER_EMAIL COLUMN ON THE NORMAL VIEW WE CREATED EARLIER
ALTER VIEW PRESENTATION.CUSTOMER_LAST_ORDER_DATE MODIFY COLUMN CUSTOMER_EMAIL SET MASKING POLICY EMAIL_MASK;

//CHANGE THE CONTEXT OF THE ROLE AND OBSERVE RESULTS
USE ROLE SENSITIVE_ALLOWED_ROLE;
SELECT * FROM PRESENTATION.CUSTOMER_LAST_ORDER_DATE;

USE ROLE SENSITIVE_DENIED_ROLE;
SELECT * FROM PRESENTATION.CUSTOMER_LAST_ORDER_DATE;

//NOTE THE THE DDL IS UNCHANGED AS THE MASKING IS APPLIED AT EXECUTION TIME
SELECT GET_DDL('VIEW', 'PRESENTATION.CUSTOMER_LAST_ORDER_DATE', TRUE);
```

### Row Access Policies

* Allow us to secure data at the individual row level based on the role executing the query

```
//ENSURE SYSADMIN IS USED
USE ROLE SYSADMIN;

//CREATE DATABASE
CREATE OR REPLACE DATABASE ROW_ACCESS;

USE DATABASE ROW_ACCESS;

//CREATE TABLE
CREATE OR REPLACE TABLE SALES
(ORDERNUMBER INTEGER,
QUANTITYORDERED INTEGER,
PRICEEACH INTEGER,
ORDERLINENUMBER INTEGER,
SALES INTEGER,
STATUS VARCHAR(100),
QTR_ID INTEGER,
MONTH_ID INTEGER,
YEAR_ID INTEGER,
PRODUCTLINE VARCHAR(100),
MSRP INTEGER,
PRODUCTCODE VARCHAR(100),
CUSTOMERNAME VARCHAR(100),
ADDRESSLINE1 VARCHAR(100),
ADDRESSLINE2 VARCHAR(100),
CITY VARCHAR(100),
STATE VARCHAR(100),
POSTALCODE VARCHAR(100),
COUNTRY VARCHAR(100),
TERRITORY VARCHAR(100),
CONTACTLASTNAME VARCHAR(100),
CONTACTFIRSTNAME VARCHAR(100),
DEALSIZE VARCHAR(100));

// file format required to load the sample sales data.
CREATE FILE FORMAT "ROW_ACCESS"."PUBLIC".csv TYPE = 'CSV' COMPRESSION =
'AUTO' FIELD_DELIMITER = ',' RECORD_DELIMITER = '\n' SKIP_HEADER = 1

//CREATE ROW ACCESS POLICY
CREATE OR REPLACE ROW ACCESS POLICY SALES_TERRITORY
    AS (TERRITORY STRING) RETURNS BOOLEAN ->
    CASE    WHEN 'SALES_MANAGER' = CURRENT_ROLE() THEN TRUE
            WHEN 'SALES_EMEA' = CURRENT_ROLE() AND TERRITORY = 'EMEA'
            THEN TRUE
            WHEN 'SALES_APAC' = CURRENT_ROLE() AND TERRITORY = 'APAC'
            THEN TRUE
    ELSE FALSE
END;

//APPLY THE ROW ACCESS POLICY TO THE TABLE
ALTER TABLE SALES
ADD ROW ACCESS POLICY SALES_TERRITORY
    ON (TERRITORY);

//GRANT PERMISSIONS ON OBJECTS TO ALL ROLES
GRANT USAGE ON WAREHOUSE COMPUTE_WH TO SALES_MANAGER;
GRANT USAGE ON WAREHOUSE COMPUTE_WH TO SALES_EMEA;
GRANT USAGE ON WAREHOUSE COMPUTE_WH TO SALES_APAC;
GRANT USAGE ON DATABASE ROW_ACCESS TO SALES_MANAGER;
GRANT USAGE ON DATABASE ROW_ACCESS TO SALES_EMEA;
GRANT USAGE ON DATABASE ROW_ACCESS TO SALES_APAC;
USE ROLE SYSADMIN;
GRANT USAGE ON SCHEMA PUBLIC TO SALES_MANAGER;
GRANT USAGE ON SCHEMA PUBLIC TO SALES_EMEA;
GRANT USAGE ON SCHEMA PUBLIC TO SALES_APAC;

GRANT SELECT ON TABLE SALES TO SALES_MANAGER;
GRANT SELECT ON TABLE SALES TO SALES_EMEA;
GRANT SELECT ON TABLE SALES TO SALES_APAC;


//ADD THESE ROLES TO YOUR OWN USER TO MAKE IT EASY TO TEST OUT
USE ROLE SECURITYADMIN;
GRANT ROLE SALES_MANAGER TO USER MYUSERNAME;
GRANT ROLE SALES_EMEA TO USER MYUSERNAME;
GRANT ROLE SALES_APAC TO USER MYUSERNAME;

//TEST OUT THE DIFFERENT ROLES AND OBSERVE THE RESULTS
USE ROLE SALES_MANAGER;
SELECT TERRITORY, COUNT(*)
FROM SALES
GROUP BY TERRITORY;

USE ROLE SALES_EMEA;
SELECT TERRITORY, COUNT(*)
FROM SALES
GROUP BY TERRITORY;

USE ROLE SALES_APAC;
SELECT TERRITORY, COUNT(*)
FROM SALES
GROUP BY TERRITORY;
```

### Advanced Snowflake Security Features

#### Future Grants

* Future grants allow us to put in place rules that tell Snowflake how to manage privileges for certain types of objects before they are created.
* Standardized layer of controls across common objects

```
GRANT USAGE ON FUTURE SCHEMAS IN DATABASE SALES TO ROLE MARKETING;
GRANT SELECT ON FUTURE TABLES IN SCHEMA SALES.BOOKS TO ROLE MARKETING;
```

#### Managed Access Schemas

* When users create objects in a schema, they become the object owner for that
object. To manage/restrict the grant ownerships priviliges to other roles, this is used.
* To centrailze the management of privileges to only certain roles, such as schema owner or anyone with MANAGE GRANTS privilege, while preventing other users from making decisions on who to grant access to.
* We use **WITH MANAGED ACCESS** keywords which execute the CREATE SCHEMA statement.
* The ownership of objects is moved away from the object owner to the schema owner (Or anyone with MANAGE GRANTS priviledge)