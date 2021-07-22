# local-s3-management

Tool to manage AWS S3 buckets and objects locally

Created for CIS-2790-NET01 Systems Analyst Simulation

_*POC/Demo application only, not for production use_

## Functionality

Complete CRUD ability on AWS Buckets and their Objects

Get all Buckets
Create a Bucket
Delete a Bucket

Get all Objects
Download an Object
Upload an Object
Delete an Object

## How To Run

Add a `.env` file to the root of the project with your AWS credentials

```txt
AWS_ACCESS_KEY_ID=xxx
AWS_SECRET_ACCESS_KEY=xxx
AWS_DEFAULT_REGION=us-east-1
```

Run the following commands to launch the application

```sh
cd s3-local-management/view
npm run build
cd ..
docker-compose up
```

##### You can now access the application from a browser at [localhost]

[localhost]: <http://localhost>
