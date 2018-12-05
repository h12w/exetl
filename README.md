exetl: an excercise of ETL service
==================================

MSA provides a general purpose CSV file ingester that supports multiple backend storages.

## Install

```bash
go get h12.io/exetl
```

## Testing

```bash
make test
```


### Start with docker-compose

```bash
docker-compose build --no-cache
docker-compose up --detach
```

## Design

### Code layout

```
exetl/ all domain types and constants that are needed by interactions between its sub-packages
    cmd/ contain all main packages of services or cli
        storage/ storage service
        ingester/ ingester service
        storage-cli/  a cli tool for storage service
        ingester-cli/ a cli tool for ingester service  
    db/memdb an in-memory implementation of exetl.DB interface
    proto/ gRPC generated code and utility functions
    service/ contain logic of all services
        storage/ storage service logic
        iingester/ ingester service logic
    testdata/  data for testing
```
