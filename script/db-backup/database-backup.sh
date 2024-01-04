#!/bin/bash

echo -e "Creating database backup.."
DATE=$(date +"%d-%m-%y")
TIME=$(date +"%H-%M-%S")
mongodump --uri "mongodb://dbAdmin:tUrb0DBpA55w0rD@148.72.144.75:27017/turboBpo?authSource=admin" --gzip --archive=/tmp/
cd /tmp/
echo -e "$DATE:$TIME"
mv  archive.gz turbo-db-$DATE:$TIME.gz
if aws s3 cp /tmp/turbo-db-$DATE:$TIME.gz  s3://turbo-bpo-backup
    then 
    curl -X POST -H 'Content-type: application/json' --data '{"text":"Successfully backup database"}' https://hooks.slack.com/services/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
    else
    curl -X POST -H 'Content-type: application/json' --data '{"text":"Failed to backup database"}' https://hooks.slack.com/services/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
fi
