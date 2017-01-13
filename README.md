# graphql-sample

This is a sample application that uses the golang
implementation of the graphql specification, Its a simple application
that was used for testing new things in the language and testing out
connections to several databases and how joins can be done on the
application side of things.


This sample application uses three different models and represents a sample message board.

There are three different model and each is stored used a different database:
1. User  -> Mysql
2. Message  -> Mongodb
3. Comment -> Cassandra (Note this is actually really inefficent the way I have it)


###Setup:

On first boot of the application run:
```
docker-compose up -d mysql mongo cassandra
```

Wait about 10 seconds for the databases to startup the run the following to seed the databases:
```
bash install.sh
```



###Requesting:

the gql server is running and exposing port 8090 to localhost.

Post a message to http://localhost:8090/graphql using the "Content-type:application/graphql" header
```
 {
   post (id:2){
     id
     content
     comments{
        content
        user{
          id
          name
        }
     }
   }
 }
```