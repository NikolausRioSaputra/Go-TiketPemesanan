# Go-TiketPemesanan

Go-TiketPemesanan adalah aplikasi web berbasis API yang ditulis dalam Golang untuk mengelola pemesanan tiket secara online. Aplikasi ini menyediakan layanan CRUD (Create, Read, Update, Delete) untuk entitas seperti pengguna, event, dan pesanan tiket. Dengan menggunakan framework Gin dan PostgreSQL sebagai database, aplikasi ini berjalan di port 8080 dan menawarkan berbagai endpoint untuk memudahkan manajemen data dan transaksi tiket secara efisien.

## 1. Folder event
Folder event berisi file yang mengelola operasi terkait event atau acara.

### create event
![create Event Screenshot](./images/EventSuccess/createEven.png)

### get all event
![get all Event Screenshot](./images/EventSuccess/getAllEvent.png)

### find by Id event
![find by id Event Screenshot](./images/EventSuccess/findByIdEvent.png)

### Log Event
![log Event Screenshot](./images/EventSuccess/logEvent.png)

### Folder Validation: Mengandung validasi data untuk memastikan input pengguna sesuai dengan aturan sebelum diproses.
![validation price Screenshot](./images/EventSuccess/validation/priceValidation.png)
![validation stock tiket Screenshot](./images/EventSuccess/validation/tiketStockEmpty.png)
![validation type empty Screenshot](./images/EventSuccess/validation/typeEventEmpty.png)
![validation date empty Screenshot](./images/EventSuccess/validation/validationDateEmty.png)
![validation location empty Screenshot](./images/EventSuccess/validation/ValidationlocationEmpty.png)
![validation Name empty Screenshot](./images/EventSuccess/validation/validationNameEmpty.png)



## 2. Folder user
### Folder user berisi file yang mengelola operasi terkait pengguna.

### create user
![Create User Screenshot](./images/UserSuccess/createUser.png)

### find by id user
![FindbyId User Screenshot](./images/UserSuccess/userfindById.png)

### get all user
![get all User Screenshot](./images/UserSuccess//getAllUser.png)

### update user
![Update User Screenshot](./images/UserSuccess/userUpdate.png)

### setelah update
![after Update Screenshot](./images/UserSuccess/afterUpdate.png)

### delete user
![Delete User Screenshot](./images/UserSuccess/deleteUser.png)
### setelah Delete
![After Delete User Screenshot](./images/UserSuccess/afterDelete.png)

### Log User
![log aman](./images/UserSuccess/statusSuccesUser.png)
![log kena validation](./images/UserSuccess/validation/LogUserFailed.png)

### File user_validation.go: Mengandung validasi data pengguna untuk memastikan input sesuai dengan aturan sebelum diproses.

![User addres empty Screenshot](./images/UserSuccess/validation/validationAddressEmpty.png)
![user balance empty Screenshot](./images/UserSuccess/validation/validationBalance.png)
![user name empty Screenshot](./images/UserSuccess/validation/validationEmtyName.png)



## 3. Folder orders
### Folder orders berisi file yang mengelola operasi terkait pesanan atau transaksi.

### create order
![create order Screenshot](./images/Orders/createOrder.png)

### get all orders
![get all orders Screenshot](./images/Orders/allOrders.png)

### setelah order balance berkuang 
![order balance berkuang Screenshot](./images/Orders/balanceUpdate.png)

### setelah order stock berkurang 
![order stock berkurang Screenshot](./images/Orders/stockUpdate.png)

### Log orders
![log orders Screenshot](./images/Orders/Validation/logOrders.png)

### File order_validation.go: Mengandung validasi data pesanan untuk memastikan transaksi berjalan sesuai aturan.

![Create User Screenshot](./images/Orders/Validation/orderEventNotFound.png)
![Create User Screenshot](./images/Orders/Validation/OrdersinsufficientBalance.png)
![Create User Screenshot](./images/Orders/Validation/OrderUsernotfound.png)
![Create User Screenshot](./images/Orders/Validation/stockNotEnough.png)
![Create User Screenshot](./images/Orders/Validation/typeTiketNotFound.png)

