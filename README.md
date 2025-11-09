# Simple REST API – User CRUD (Gin + In-Memory Storage)

This is a simple REST API project built with **Golang** using the **Gin** framework.  
The application follows a **Clean Architecture** approach and demonstrates a minimal, maintainable backend setup — without any database, using only in-memory storage.

---

## ✨ Features

### Gin Framework
- RESTful API built using **Gin**
- Follows clean and modular structure: `handler`, `service`, `repository`, and `validation`

### In-Memory Repository
- Stores all user data directly in memory (RAM)
- No database dependency (lightweight & fast)

### Validation Layer
- Custom request validation before data is processed
- Ensures that user input (like name and email) meets required rules

### Clean Architecture
- Clear separation between:
  - **Handler** – HTTP logic (Gin routes)
  - **Service** – Business logic
  - **Repository** – Data layer (in-memory)
  - **Validation** – Request validation

### CRUD Operations
- Create new users  
- Read (get all users or a single user by ID)  
- Update existing users  
- Delete users  

