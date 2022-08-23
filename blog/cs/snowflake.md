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