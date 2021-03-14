# Stellar Identicon Generate API


The API is based on [stellar-identicon-go](github.com/overcat/stellar-identicon-go),
you can use the API to generate identicons online.

## API
```JSON
{
    "_link":{
        "avatar":{
            "href": "/avatar/{account_id}/{?width,height}",
            "templated": true
        }
    },
    "stellar_identicon_api_source_code": "https://github.com/overcat/stellar-identicon-api",
    "stellar_identicon_go_source_code": "https://github.com/overcat/stellar-identicon-go"
}
```