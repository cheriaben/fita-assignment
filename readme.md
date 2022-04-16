# Simple Checkout REST API

### Endpoint

**[POST] /checkout**:

Request:
```json
{
  "data": {
    "checkout_items": [
      {
        "sku": "43N23P",
        "qty": 1
      },
      {
        "sku": "234234",
        "qty": 1
      }
    ]
  }
}
```

Response:
```json
{
  "data": {
    "checkout": {
      "final_price": 5399.99
    }
  },
  "metadata": {
    "path": "/checkout",
    "message": "POST /checkout [200] OK",
    "timestamp": "2022-04-15T01:00:03+07:00"
  }
}
```
