#!/bin/bash

#Note this install assumes that the mysql and mongo containers are already running via docker-compose

#This initializes mysql with some sample data
docker-compose exec mysql mysql -e 'CREATE DATABASE graphql_sample;'
docker-compose exec mysql mysql -e 'CREATE TABLE graphql_sample.user (id int unique, name VARCHAR(20), height int, weight int);'
docker-compose exec mysql mysql -e 'INSERT INTO graphql_sample.user (id, name, height, weight) VALUES
                                (1,"Derek", 10,100),
                                (2,"Cory", 11,90),
                                (3,"Brett", 12,80),
                                (4,"Jesse", 13,70);'

#This will initialize mongodb with some sample data
docker-compose exec mongo mongo graphql_sample --eval 'db.post.insert({"_id":1, "content":"This is my content!", "userid": 1})'
docker-compose exec mongo mongo graphql_sample --eval 'db.post.insert({"_id":2, "content":"Shut up and dance!", "userid": 2})'
docker-compose exec mongo mongo graphql_sample --eval 'db.post.insert({"_id":3, "content":"Be vewwy vewwy quiet", "userid": 2})'
docker-compose exec mongo mongo graphql_sample --eval 'db.post.insert({"_id":4, "content":"Oh la la!!!", "userid": 4})'
docker-compose exec mongo mongo graphql_sample --eval 'db.post.createIndex({"userid":1})'