#  Currency Converter 

A Currency Converter API built with **Go (Gin)** that uses **Keploy** for e2e testing and external API mocking.

- Fetches real-time exchange rates via the [Frankfurter API](https://www.frankfurter.app/).
- Uses a public, open-source exchange rate provider.
- Integrated with **Keploy** to record and replay API interactions, ensuring tests pass even if exchange rates change or the internet is down.


# Installation

1. Clone the repo

```
git clone https://github.com/cnu1812/currency-converter.git
```

2. Run the application

```
sudo -E keploy record -c "go run main.go"
```

3. Open a new terminal and send requests to the app.

example

```
curl -X POST http://localhost:8080/convert \
-H "Content-Type: application/json" \
-d '{"from": "USD", "to": "EUR", "amount": 100}'
```

Expected output

```
{"converted_amount":85.419,"exchange_rate":0.85419,"message":"Calculated using live rates!","original_amount":100,"original_currency":"USD"}

```