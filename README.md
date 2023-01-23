# Go Service
|   Date    |   Change                  |    DoneBy     |
|-----------|---------------------------|---------------|      
|01-20-23   | Initial commit            | Sidharth      |
|01-23-23   | Cache functionality       | Selva         |
_________________________________________________________

## This go module consumes the running instance of https://github.com/mrds91/sb-av

- Clone go module https://github.com/mrds91/go-av & springboot module https://github.com/mrds91/sb-av
- Run the springboot module from within STS.
- Build the go module by `go build`
- Start the services by executing the command `go run .\cmd\Main.go` from the root directory 

### Use the below curls to test the services.

### 1> Address-validation service
curl --location --request POST 'http://mylocal.sprint.com:8082/address-validation' \
--header 'Content-Type: application/json' \
--data-raw '{
    "addresses": [
        {
            "address": {
                "state": "VA",
                "addressLine2": "172 Draper Ave",
                "city": "Reston",
                "addressLine1": "12524 Sunrise Valley Dr",
                "zipCode": "20191",
                "zipCodeExtension": "02760"
            },
            "addressTypes": [
                {
                    "addressUsageType": "BILLING"
                }
            ]
        }
    ],
    "flowType": "postpaid"
}


### 2> Post a cart Item
curl --location --request POST 'http://mylocal.sprint.com:8082/cart/456' \
--header 'Content-Type: application/json' \
--data-raw '{
    "addresses": [
        {
            "address": {
                "state": "VA",
                "addressLine2": "172 Draper Ave",
                "city": "Reston",
                "addressLine1": "12524 Sunrise Valley Dr",
                "zipCode": "20191",
                "zipCodeExtension": "02760"
            },
            "addressTypes": [
                {
                    "addressUsageType": "BILLING"
                }
            ]
        }
    ],
    "flowType": "postpaid"
}


### 3> Retrieve cart Item 
curl --location --request GET 'http://mylocal.sprint.com:8082/cart/456' \
--header 'Content-Type: application/json' \
--data-raw '{
    "addresses": [
        {
            "address": {
                "state": "VA",
                "addressLine2": "172 Draper Ave",
                "city": "Reston",
                "addressLine1": "12524 Sunrise Valley Dr",
                "zipCode": "20191",
                "zipCodeExtension": "02760"
            },
            "addressTypes": [
                {
                    "addressUsageType": "BILLING"
                }
            ]
        }
    ],
    "flowType": "postpaid"
}
