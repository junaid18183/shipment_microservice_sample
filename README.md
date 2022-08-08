# Sample Shipment API

This is sample product API written in node using Express. For sale of simplicity it does not use any real DB and all data come directly from `fakeProductData.js`


# How to Run locally

Start the application

`go run main.go`

# Test the API

Get the single Shipment

```
curl  -s localhost:8801/api/shipments/1 | jq "."
{
  "id": 1,
  "manager": "Juned",
  "warehouse": "Pune",
  "shipment": "250320220500",
  "product": "7",
  "status": "",
  "approver": "",
  "quantity": "8",
  "created_on": "Fri, 25 Mar 2022 05:00:00 GMT"
}

```

Get all of the  single Shipments
```
curl -s localhost:8801/api/shipments | jq "."
```
Dumy commit
