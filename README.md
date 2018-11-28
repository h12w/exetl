msa: micro service assignment
=============================

MSA provides a general purpose CSV file ingester that supports multiple backend storages.

## Code Layout

```
msa/ all domain types and constants that are needed by interactions between its sub-packages
    cmd/ contain all main packages of services or cli
        storage/ storage service
        ingester/ ingester service
        storage-cli/  a cli tool for storage service
        ingester-cli/ a cli tool for ingester service  
    db/memdb an in-memory implementation of msa.DB interface
    proto/ gRPC generated code and utility functions
    service/ contain logic of all services
        storage/ storage service logic
        iingester/ ingester service logic
    testdata/  data for testing
```

## Install

```bash
go get github.com/h12w/msa
```

## Testing

```bash
make test
```


### Docker setup

```bash
docker-compose up
```
