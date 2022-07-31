# Cloud Platforms such AWS, GCP and Azure

Secure cloud services platform, offering compute power, database storage, content delivery and other functionality to help businesses scale and grow.

Pay as per your usage model.

## DevOps

A set of practices intended to reduce the time between committing a change to system and the change being placed into normal production, while ensuring high quality.

### Popular DevOps Tools

* ANSIBLE
* splunk
* git
* maven
* graddle
* Jenkins
* JIRA


### CICD (Continious Integration and Continious Delivery)

#### AWS CodePipeline

Continuous delivery service which helps us to model, visualize and automate the steps required to release our software.

*Uses*

* Monitoring our processes in real time
* Ensures consistent Release process
* Speed up delivery while improving quality.
* View pipeline history details

**Architecture**

Source Code -> Build -> Staging (Deployed and Tested) -> Manual Approval -> Production (Code Deployed)

*Stages*

* Code changes trigger pipeline to run
* Source (Runs source action with source provider) - output(.zip) stored in Amazon S3 Artifact Bucket
* Build (Runs build action with build provider) - output(.zip) stored in Amazon S3 Artifact Bucket
* Staging (Runs Deploy action with deployment provider)
* Applications deploy to instances in the AWS cloud; container based applications deploy services in AWS Cloud.


##### Code Commit

##### Code Build

##### Code Deploy9
