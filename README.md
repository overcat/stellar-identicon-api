# Stellar Identicon Generate API

The API is based on [stellar-identicon-go](https://github.com/StellarCN/stellar-identicon-go/), you can use the API to
generate identicons online.

## API

```JSON
{
  "_link": {
    "avatar": {
      "href": "/avatar/{account_id}/{?width,height}"
    },
    "self": {
      "href": "/"
    }
  },
  "sep_0033_file": "https://github.com/stellar/stellar-protocol/blob/master/ecosystem/sep-0033.md",
  "stellar_identicon_api_source_code": "https://github.com/overcat/stellar-identicon-api",
  "stellar_identicon_go_source_code": "https://github.com/overcat/stellar-identicon-go"
}
```

## Deploy

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/overcat/stellar-identicon-api)