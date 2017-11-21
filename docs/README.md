# REST JSON API Documentation

This REST API defines the interface exposed by the `go-wt` server, in order to interact with Winding Tree smartcontracts published on the ethereum blockchain.

Queries that implies a change on the blochain (eg: publishing smartcontract, modifying state) infer costs in ETH or LIF and therefore require prior user identification.

## General

* The API is based on HTTP REST with JSON as data object modeling.
* All fields are optional unless mentionned otherwise.

## Ressources

The API is split by ressources:

* [User](User.md): Exposes methods to interact with the user on the blockchain
* [Hotel](Hotel.md): Exposes methods to interact with hotel smartcontracts