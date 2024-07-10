# API - GO

### GOAL
- Create an API to manage students of a course 

### Routes
- GET /students - List all students
- POST /students - Create student
- GET /students/:id - Get info from a specific student
- PUT /students/:idd - Update student
- DELETE /students/:id - Delete student
- GET /students?actie=<true/false> - Get a active / non-active user

### Struct Student
- Name
- CPF
- Email
- Age
- Active