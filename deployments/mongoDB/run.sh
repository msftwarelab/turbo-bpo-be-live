
#!/bin/bash

sudo docker build -t hello-mongo:latest .
sudo docker run -e MONGO_INITDB_ROOT_USERNAME=admin \
      -e MONGO_INITDB_ROOT_PASSWORD=@ccess.123 \ --name turboBPO-DB -d -v /tmp/mongodb:/data/db -p 27017:27017 hello-mongo