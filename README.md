# data-wallet-go

## mfile 

### Upload

#### Create

Upload file to get cid

Method: POST

Body Param:
- data: string (required) - The data to be uploaded.
- address: string (required) - The address associated with the data.

Response:
```
{
    "cid":""
}
```
#### Confirm
Sign the file did and generate file-did

Method: POST

Body Param:
- sig: string (required) - The signature used to confirm the upload.(Sign the cid)
- address: string (required) - The address associated with the upload.

Response:
```
{
    "mdid":""
}
```

### Download
According to the file did, download the file

Method: GET

Query Param:
- mdid: string (required) - The unique identifier for the file to be downloaded.
- address: string (required) - The address associated with the file.

## domain

### bind
Domain name binding did

Method: POST

Body Param:
- domain: string (required) - The domain to be bound.
- address: string (required) - The address to which the domain is to be bound.

### renewal
Renewal Domain Name 

Method: POST

Body Param:
- domain: Domain name to be renewed


