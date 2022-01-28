# backend-challenge-2022

Base URL: https://shopify-challenge-ahmad.herokuapp.com

Make requests using any tool, i.e. cURL, Postman, Insomnia, etc.

## Endpoints:

- Creating inventory (you don't need the /create but I thought it'd be good to be explicit, although it's not best practice):

  `PUT` `https://shopify-challenge-ahmad.herokuapp.com/inventory/create`

    ```json
    "title": string,
    "description" string,
    "price" float,
    "quantity": int,
    ```

- View All Inventory:

  `GET` `https://shopify-challenge-ahmad.herokuapp.com/inventory/view/all`

- View Specific Inventory:

  `GET` `https://shopify-challenge-ahmad.herokuapp.com/inventory/view/{Item ID}`

- Delete Inventory:

  `DELETE` `https://shopify-challenge-ahmad.herokuapp.com/inventory/delete/{Item ID}`

  ```json
    "reason": string
    ```

- Undo Delete of Inventory (this should not be a DELETE request, oversight...):

  `DELETE` `https://shopify-challenge-ahmad.herokuapp.com/inventory/undo-delete/{Item ID}`


- Editing inventory:

  `PUT` `https://shopify-challenge-ahmad.herokuapp.com/inventory/edit/{Item ID}`

    ```json
    "title": string, - optional
    "description" string, - optional
    "price" float, - optional
    "quantity": int, - optional
    ```

  There is an error with quantity where you are not able to set it to 0, this is due to the parsing in go, it can be
  fixed but I was in a rush.
