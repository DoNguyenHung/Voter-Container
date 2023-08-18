#!/bin/bash
curl -d '{ "VoterId": 1, "FirstName": "John", "LastName": "Doe", "VoteHistory": [{"PollID": 59231, "VoteDate": "2021-08-15T14:30:45.00Z"}] }' -H "Content-Type: application/json" -X POST http://localhost:1080/voters/1
curl -d '{ "VoterId": 2, "FirstName": "Jane", "LastName": "Schmoe", "VoteHistory": [{"PollID": 12345, "VoteDate": "2021-08-16T14:30:45.00Z"}] }' -H "Content-Type: application/json" -X POST http://localhost:1080/voters/2
curl -d '{ "VoterId": 3, "FirstName": "Bob", "LastName": "Ross", "VoteHistory": [{"PollID": 54321, "VoteDate": "2021-08-17T14:30:45.00Z"}] }' -H "Content-Type: application/json" -X POST http://localhost:1080/voters/3
