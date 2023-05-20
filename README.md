# dynamitedb

A distributed key-value database

Inspired by DynamoDB 2007 paper

There is 2 major components:
    1. Coordinator: Responsible for coordinating read/write requests to data store 
    2. Gossiper: Responsible for knowing about the existence of other nodes as well as who is responsible for storing what keys
