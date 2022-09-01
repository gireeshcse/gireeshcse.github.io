# SNOWFLAKE

* Software-as-a-Service (SaaS)
* Enables data
    - Storage
    - Processing
    - Analytic Solutions
* Data Platform as a Cloud Service.

### Architecture - **Hybrid of traditional shared-disk and shared-nothing database architectures.**

- Uses central data repository for persisted data that is accessible from all the compute nodes in the platform.

- Processes queries using MPP (Massively parallel processing) compute clusters where each node in the cluster stores a portion of entire data set locally.

- 3 Layers
    - Database storage
    - Query Processing
    - Cloud Services

### User Management

* **USERADMIN** system role can create users using SQL (CREATE USER)

```
CREATE USER sam PASSWORD = 'sam@123tiger';
```

* Only the role with **OWNERSHIP** priviledge on a user, or higher role can modify most user properties using SQL (ALTER USER)

```
ALTER USER sam RENAME TO ram;
ALTER USER ram set PASSWORD = 'hello@124tiger';
ALTER USER ram set disabled = true; # disabling user
ALTER USER ram set mins_to_unlock = 0 ; # unlocking user
ALTER USER janesmith SET DEFAULT_WAREHOUSE = mywarehouse DEFAULT_NAMESPACE = mydatabase.myschema DEFAULT_ROLE = myrole DEFAULT_SECONDARY_ROLES = ('ALL');
```

* Snowflake uses roles to control the objects (virtual warehouses, databases, tables etc.) that users can access.

    - **ACCOUNTADMIN** role is required for performing account-level administrative tasks.

* GRANT ROLE command allows us to assign multiple roles to a single user.

```
GRANT ROLE myrole to user ram;
DROP USER ram;
DESC USER ram;
```

#### Databases

* Organization of uploaded structured and semi-structured data sets which are held for processing and analysis. 

* Snowflake automatically manages all parts of the data storage process, including organization, structure, metadata, file size, compression and statistics.

#### Warehouse

* One or more clusters that provide compute
* Snowflake virtual warehouse is a cluster of Database Servers that are deployed ondemand to handle user queries.

#### Worksheets

Provide a powerful and versatile method of running SQL queries as well as performing other Snowflake data loading, definition, and manipulation tasks.


### Installing SnowSQL


```
curl -O https://sfc-repo.snowflakecomputing.com/snowsql/bootstrap/1.2/linux_x86_64/snowsql-1.2.22-linux_x86_64.bash
curl -O https://sfc-repo.azure.snowflakecomputing.com/snowsql/bootstrap/1.2/linux_x86_64/snowsql-1.2.22-linux_x86_64.bash.sig
gpg --verify snowsql-1.2.22-linux_x86_64.bash.sig snowsql-1.2.22-linux_x86_64.bash
bash snowsql-1.2.22-linux_x86_64.bash
snowsql -v
```

#### Sign In to terminal and creating database

```
snowsql -a <account-name> -u <username>
create or replace database sf_tuts;
select current_database(),current_schema();

+------------------+----------------+                                           
| CURRENT_DATABASE | CURRENT_SCHEMA |
| ()               | ()             |
|------------------+----------------|
| SF_TUTS          | PUBLIC         |
+------------------+----------------+

create or replace table emp_basic (
  first_name string ,
  last_name string ,
  email string ,
  streetaddress string ,
  city string ,
  start_date date
);

+---------------------------------------+                                       
| status                                |
|---------------------------------------|
| Table EMP_BASIC successfully created. |
+---------------------------------------+
1 Row(s) produced. Time Elapsed: 0.299s
GIREESH#COMPUTE_WH@SF_TUTS.PUBLIC>
```
Make a Virtual Warehouse

```
create or replace warehouse sf_tuts_wh with
  warehouse_size='X-SMALL'
  auto_suspend = 180
  auto_resume = true
  initially_suspended=true;
```

Uploading the csv data 

```
put file:///tmp/employees0*.csv @<database-name>.<schema-name>.%<table-name>;
put file:///home/machine/snowflake/getting-started/employees0*.csv @sf_tuts.public.%emp_basic;

```
Note: The @ sign before the database and schema name @sf_tuts.public indicates that the files are being uploaded to an internal stage, rather than an external stage. The % sign before the table name %emp_basic indicates that the internal stage being used is the stage for the table

List staged files

```
list @<database-name>.<schema-name>.%<table-name>;
list @sf_tuts.public.%emp_basic;
```

Copying data into our table

```
copy into emp_basic
  from @%emp_basic
  file_format = (type = csv field_optionally_enclosed_by='"')
  pattern = '.*employees0[1-5].csv.gz'
  on_error = 'skip_file';
```

Querying data and inserting data and droping databases and warehouse

```
select * from emp_basic;
select * from emp_basic where first_name = 'Ron';
select email from emp_basic where email like '%.au';
insert into emp_basic values
  ('Clementine','Adamou','cadamou@sf_tuts.com','10510 Sachs Road','Klenak','2017-9-22') ,
  ('Marlowe','De Anesy','madamouc@sf_tuts.co.uk','36768 Northfield Plaza','Fangshan','2017-1-26');
drop database if exists sf_tuts;
drop warehouse if exists sf_tuts_wh;
```

To close connection

```
!exit
```

### Snowflake usecases

* Near-Zero matainenance

    - Reduces administrative work and frees up technical staff to focus on increasing analytics.
* Faster data ingestion and transformation.
* Snowflake delivers the Data Cloud - a global network where thousands of organizations mobilze data with near-unlimited scale, concurrency, and performance. Inside the Data Cloud, organizations unite their siloed data, easily discover and securely share governed data nad execute diverse analytic workloads.
* Multi-cluster shared data architecture with per-second pricing.

### Snowpark

* Developer framework for snowflake.
* Allows data engineers, data scientists and data developers to code in their language of choice and execute pipeline, ML workflow and data apps faster and more securely, in a single platform
* Designed to make building complex data pipelines much easier and to allow developers to interact with Snowflake directly without having to move data.
* Moving data outside of Snowflake introduces its own challenges; data security and privacy risks, data ingress egress charges , failed processing due to non-scalable resources and the outcomes are still inaccurate because they are based upon a sample of data. Processing the data inside snowflake using Snowpark UDF's is highly accurate, because outcomes are based upon the entire dataset not just sample.

#### Data Augmentation

* Techniques used to increase the amount of data by adding slightly modified copies of already existing data.
* Generating new data points from existing data. Includes minor alterations to data or using ML models to generate new points 
* Acts as a regularizer and helps reduce overfitting when training a machine learning model.
#### Possible usecases

- Using machine learning to augment data by hosting trained models
- Scanning for anomalies in our data
- Developing a routine to identify PII(personal identifiable information)

#### Usecase: Optimizing DataOps with StreamSets Engine for SNOWPARK

The StreamSets engine for Snowpark includes the benefits of the StreamSets DataOps Platform— built-in monitoring and orchestration of complex data pipelines at scale—all in the cloud and with no additional hardware required.

### Loading Data

* Best files sizes to be uploaded 100-250 MB in size.
* Loading large data sets can affect query performance. It is recommended using dedicated separate warehouses for loading and querying operations to optimize peformance for each.
* Organize/Partition data in stages by paths

* The data files contains load status meta data. (Expires after 64 days)
* Snowflake maintains detailed metadata for each table into which data is loaded, including 
  - Name of each file from which data was loaded
  - File Size
  - No. of rows parsed in the file
  - Timestamp of the last load for the file
  - Information about any errors encountered in the file during loading.

* JSON data : Removing "null" values
  - STIP_NULL_VALUES (option to TRUE)
* CSV Data : Trimming Leading Spaces.
  - fileds enclosed in quotes but has a leading space before opening quoattion character for each field , snowflake reads the leading space rather than the opening quotation character as the beginning of the field. The quotation characters are interpreted as string data.
* Removing loaded data files

  - Using PURGE copy option.
  - Using REMOVE command
    ```
    REMOVE internalStage [ PATTERN = '<regex_pattern>' ]
    ```
      - Remove all files from the path1/subpath2 path in a named internal stage named mystage
        ```
        REMOVE @mystage/path1/subpath2;
        ```
      - Remove all files from the stage for the orders table:
        ```
        REMOVE @%orders;
        ```

  ```
  "value1", "value2", "value3"

  
  COPY INTO mytable
  FROM @%mytable
  FILE_FORMAT = (TYPE = CSV TRIM_SPACE=true FIELD_OPTIONALLY_ENCLOSED_BY = '0x22');

  # the above cmd, trims the leading space and removes the quotation marks enclosing each field:

  SELECT * FROM mytable;

  +--------+--------+--------+
  | col1   | col2   | col3   |
  +--------+--------+--------+
  | value1 | value2 | value3 |
  +--------+--------+--------+
  ```
* Creating file format
  ```
  create or replace file format my_csv_format
  type = csv
  field_delimiter = "|"
  skip_header = 1
  null_if = ('NULL','null')
  empty_field_as_null = true
  compression = gzip;

  create or replace file format my_json_format 
  type = json;
  ```
* Validating our data
  -  Before loading our data, we can validate that the data in the uploaded files will load correctly.
  - **VALIDATION_MODE** parameter returns any errors that it encounters in a file. We can then modify the data in the file to ensure it loads without error.
  - **ON_ERROR** copy option to indicate what action to perform if errors are encountered in a file during loading.

* VALIDATE
  - Validates the files loaded in a past execution of **COPY INTO** command and returns all the errors encountered during load

  ```
  select * from table(validate(mytable, job_id => '_last'));
  select * from table(validate(mytable, job_id=>'5415fa1e-59c9-4dda-b652-533de02fdcf1'));
  CREATE OR REPLACE TABLE save_copy_errors AS SELECT * FROM TABLE(VALIDATE(mytable, JOB_ID=>'5415fa1e-59c9-4dda-b652-533de02fdcf1'));
  ``` 
#### Spliting csv files

```
split [-l line_count ] [-p pattern] [file [name]]

split -l 100000 file.csv pages
```

#### Snowpipe

* Designed to load new data.
* Creating smaller data files and staging them must not be more than once per minute.
* Snowflake's continuous data ingestion service.
* Loads data within minutes after files are added to a stage and submitted for ingestion.
* Snowflake manages load capacity, ensuring optimal compute resources to meet demand.
* Snowpipe provides "pipeline" for loading fresh data in micro-batches as soon as it is available.

* Different mechanisms for detecting the staged files are:
  - Automating Snowpipe using cloud messaging
  - Calling Snowpipe REST endpoints.

#### copy

```
COPY INTO load1 FROM @%load1/data1/ FILES=('test1.csv', 'test2.csv', 'test3.csv');

COPY INTO people_data FROM @%people_data/data1/
   PATTERN='.*person_data[^0-9{1,3}$$].csv';

```
#### Stages

A stage specifies where data files are stored (i.e "staged") so that data in the files can be loaded into a table.

Types of Internal Stages.
- User
  - Referenced using **@~**
  - Cannot be altered or dropped.
  ```
  PUT file:///data/data.csv @~/staged;
  LIST @~;
  COPY INTO mytable from @~/staged FILE_FORMAT = (FORMAT_NAME = 'my_csv_format');
  ```
- Table
  - Convenient option if our files need to be accessible to multiple users and only need to be copied into a single table.
  - Referenced using **@%mytable** 
  ```
  PUT file:///data/data.csv @%mytable;
  LIST @%mytable;
  ```
  - The following command loads all files in the stage for the **mytable** table.
    ```
    copy into mytable file_format = (type = csv field_delimiter = '|' skip_header = 1)
    ```
- Named
  - Provide greatest degree of flexibility for data loading.
  - Recommended when we plan regular data loads that could involve multiple users and/or tables.

    ```
    create or replace stage my_stage file_format = my_csv_format;
    PUT file:///data/data.csv @my_stage;
    LIST @my_stage;
    copy into mytable from @my_stage;
    ```

By default, each user and table in Snowflake is automatically allocated an internal stage for staging data files to be loaded.