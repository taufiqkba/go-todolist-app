# Go-todolist-app

Model
- Todo
- Message string
- Checked bool

Usecase
- Create todo Items
- Get All Items
- Checked list Items

#### 7 steps to create todo-list app

1. Create Domain (Using Domain Driven Design)
   - todocore

2. Create Model
   - Todo

3. Create UseCase
   - RunTodoCreate
   - GetAllGoto
   - RunTodoCheck

4. Create Repository/services
   - to access database via interface

5. Create Gateway
   - implementation of interface from repository/services

6. Create Controller
   - RunTodoCreate: POST /todo
   - GetAllGoto: GET /todo
   - RunTodoCheck: PUT /todo/:todo_id

7. Create application
   - Usecase + Gateway + Controller