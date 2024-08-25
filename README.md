# Silk-Go
A Go library for using OlegDB
Exists to encapsulate and simplify an agreed usage pattern for OlegDB from Go for other projects. Specifically, it uses OlegDB to make a fancy linked list useable for graphs.

## Build
make all
then `make build` after initial use

## Usage
[import Silk
silk_db := *Silk.New()
silk_db.NewDatabase("/tmp")
node := silk.NewRelationalNode("Test Value")
silk.PushNode(node)
node_copy := silk.PullNode("Test Value")]