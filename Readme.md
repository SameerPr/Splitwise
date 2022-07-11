# Splitwise

This is a prototype for Splitwise application.

How to run :

```text
go test -v ./...
go run main.go
```

APIs:

1. Add Users
   > POST `http://localhost:8080/users/`
    1. data format ```{
       "Id": 1,
       "Name": "Sameer",
       "Email": "sameerprajapati84@gmail.com"
       }```
2. Remove User
   > DELETE `http://localhost:8080/users/1`
3. Add Expense
   > POST ```http://localhost:8080/expense```
    1. data format

```text
    {
     "id": 1,
     "amount": 200,
     "paidBy": 1,
     "type": "Equal",
     "users": [1,2,3]
    }
    
    or
     
    { 
    "id": 2,
    "amount": 500,
    "paidBy": 2,
    "type": "Exact",
    "split": [
             {"id": 1,"amount": 200 },
             {"id": 2,"amount": 300 }
            ]
   }
   
   or 
   
   {
    "id": 2,
    "amount": 500,
    "paidBy": 2,
    "type": "Percentage",
    "percentSplit": [
        {
            "id": 1,
            "percentage": 40
        },
        {
            "id": 2,
            "percentage": 60
        }
    ]
}
```
4. Remove Expense
   > DELETE `http://localhost:8080/expense/1`

5. Show User Balance
   > GET `http://localhost:8080/users/balance/1`
6. Show all balances
   > GET `http://localhost:8080/users/balances`

included:  
- Postman collection with all these APIs and data format




#### Author:  Sameer Prajapati 