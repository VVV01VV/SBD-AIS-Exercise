# Software Architecture for Big Data - Exercise 2

This exercise we're going to implement a REST server in Golang with Chi, that exposes the following routes:

- GET `/api/menu`
  - Returns a slice (array) of drinks
- GET `/api/order/all`
  - Returns a slice of orders
- GET `/api/order/totalled`
    - Returns a map of order ID + how many of these drinks have been ordered
- POST `/api/order`
  - Accepts a single order JSON and stores it in memory
- GET `/openapi/*`
    - Serves OpenAPI documentation and can be used for testing
- GET `/`
    - Serves a simple dashboard showing your orders

The server is reachable via port :3000, i.e. http://localhost:3000.
A scaffold is provided to get you started in the `skeleton/` folder.
You need to complete the webservice by resolving all comments including a `todo` keyword.

Once you've created your REST routes, execute the `build-openapi-docs.sh` script to create an OpenAPI compatible
REST documentation.

## Model Definitions
```
Drink
{
    description	string
    id	        integer
    name        string
    price       number
}
```

``` Example
{
  "id": 1,
  "name": "Beer",
  "price": 2.5,
  "description": "A refreshing beer"
}
```

```
Order
{
    amount      integer
    created_at  string
    drink_id    integer // foreign key
}
```

```Example
{
  "drink_id": 1,
  "amount": 2,
  "created_at": "2025-10-16T19:10:47Z"
}
```


## **Running the Server**

1. Open a terminal inside the **skeleton** directory:

   ```
   D:\Big Data Class\SBD-AIS-Exercise\Exc_2\skeleton
   ```

2. Make sure all dependencies are ready:

   ```
   go mod tidy
   ```

3. Generate or update Swagger documentation :

   ```
   & "D:\Big Data Class\go\bin\swag.exe" init --pd --st -g main.go --dir "." -o docs
   ```

4. Start the server:

   ```
   go run .
   ```

5. Open these URLs in browser:

   * **Dashboard:** [http://localhost:3000](http://localhost:3000)
   * **Swagger UI:** [http://localhost:3000/openapi/index.html](http://localhost:3000/openapi/index.html)

---

## **Creating Orders (via Swagger)**

Orders need to be inserted one by one using Swagger UI:

1. Go to [http://localhost:3000/openapi/index.html]
2. Expand **POST /api/order**
3. Click **Try it out**
4. Paste one by one the orders in JSONs format and click **Execute** after each

### Order 1 – 2 Beers

```
{"drink_id": 1,"amount": 2}
```

### Order 2 – 3 Spritzers

```
{"drink_id": 2,"amount": 3}
```

### Order 3 – 1 Coffee

```
{"drink_id": 3,"amount": 1}
```

After each **Execute**, you should see:

```
Response body: "ok"
Response code: 200
```

---

## **Checking Totals and History**

After inserting the orders:

1. Expand **GET /api/order/totalled**

   * Click **Try it out → Execute**
   * You should see a response showing you the amount of ordered drinks of each id

2. Expand **GET /api/order/all**

   * Click **Execute**
   * You’ll see all orders with timestamps

3. Refresh the **dashboard** at [http://localhost:3000](http://localhost:3000)

   * The **Total Orders** bar chart and **Order History** line chart should now update and display data.

---

## **Notes**

* All data is stored in memory — restarting the server clears the orders.

* The OpenAPI documentation (Swagger) is generated automatically using swaggo.

