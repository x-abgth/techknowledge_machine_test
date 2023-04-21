# Task - 4

Run the below command from <strong>task_4</strong> directory - <br>
`go run cmd/api/main.go`

<br>

Test routes in `localhost:3000`

<br>

You can also this postman api collection link to ease the process - 
[Postman API Collection Link](https://www.postman.com/abgth/workspace/machine-test/collection/20732200-97053ba0-f54a-4997-806a-5a7759a97ad0?action=share&creator=20732200)

## Outputs

### Inserting an item
- `localhost:3000/items`<br>
method = "POST"<br><br>
![insert-one-item](outputs/insert-one-item.png)

- `localhost:3000/items`<br>
method = "GET"<br><br>
![get-all-items](outputs/get-all-items.png)

- `localhost:3000/items/{id}`<br>
method = "GET"<br><br>
![get-one-item](outputs/get-one-item.png)

- `localhost:3000/items/{id}`<br>
method = "PUT"<br><br>
![update-one-items](outputs/update-one-item.png)

- `localhost:3000/items/{id}`<br>
method = "DELETE"<br><br>
![delete-one-items](outputs/delete-one-item.png)