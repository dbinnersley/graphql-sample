#!/bin/bash

#Note this install assumes that the mysql and mongo containers are already running via docker-compose

#This initializes mysql with some sample data
docker-compose exec mysql mysql -h mysql -e 'CREATE DATABASE graphql_sample;'
docker-compose exec mysql mysql -h mysql -e 'CREATE TABLE graphql_sample.user (id VARCHAR(20) unique, name VARCHAR(20), height int, weight int);'
docker-compose exec mysql mysql -h mysql -e 'INSERT INTO graphql_sample.user (id, name, height, weight) VALUES
                                ("1","Derek", 10,100),
                                ("2","Cory", 11,90),
                                ("3","Brett", 12,80),
                                ("4","Jesse", 13,70);'

#This will initialize mongodb with some sample data
docker-compose exec mongo mongo graphql_sample --eval 'db.post.insert({"_id":"1", "content":"This is my content!", "userid": "1"})'
docker-compose exec mongo mongo graphql_sample --eval 'db.post.insert({"_id":"2", "content":"Shut up and dance!", "userid": "2"})'
docker-compose exec mongo mongo graphql_sample --eval 'db.post.insert({"_id":"3", "content":"Be vewwy vewwy quiet", "userid": "2"})'
docker-compose exec mongo mongo graphql_sample --eval 'db.post.insert({"_id":"4", "content":"Oh la la!!!", "userid": "2"})'
docker-compose exec mongo mongo graphql_sample --eval 'db.post.createIndex({"userid":1})'


#This will do a bunch of inserts into cassandra!!!
docker-compose exec cassandra cqlsh -e "CREATE KEYSPACE graphql_sample WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };"
docker-compose exec cassandra cqlsh -e "CREATE TABLE graphql_sample.comment_by_id (id VARCHAR PRIMARY KEY, content VARCHAR, userid VARCHAR, postid VARCHAR);"
docker-compose exec cassandra cqlsh -e "insert into graphql_sample.comment_by_id(id,content,userid,postid) values ('1','First content','1','1');"
docker-compose exec cassandra cqlsh -e "insert into graphql_sample.comment_by_id(id,content,userid,postid) values ('2','AWESOME!!!!!!','2','3');"
docker-compose exec cassandra cqlsh -e "insert into graphql_sample.comment_by_id(id,content,userid,postid) values ('3','Why are we doing this....','4','2');"
docker-compose exec cassandra cqlsh -e "insert into graphql_sample.comment_by_id(id,content,userid,postid) values ('4','Nothing better','4','3');"