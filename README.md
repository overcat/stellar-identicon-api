# Stellar Identicon Generate API


The API is based on [stellar-identicon-go](github.com/overcat/stellar-identicon-go),
you can use the API to generate identicons online.

https://stellar-identicon.herokuapp.com/

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

## Deploy

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/overcat/stellar-identicon-api)