# Register

```bash
curl -X POST http://localhost:8000/api/v1/auth/register \
-H "Content-Type: application/json" \
-d '{
  "email": "test@test.com",
  "password": "123456",
  "role": "student"
}'
```

# Login
```bash
curl -X POST http://localhost:8000/api/v1/auth/login \
-H "Content-Type: application/json" \
-d '{
  "email": "test@test.com",
  "password": "123456"
}'
```

# Protected
## `export TOKEN=<<Token_access>>`

export TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJkNzZkMzZiNy0xNTI5LTQwZDQtOGFjZi05OTFmYzUzZjRiYWEiLCJSb2xlIjoic3R1ZGVudCIsImV4cCI6MTc3NDI0OTk1MSwiaWF0IjoxNzc0MjQ5MzUxfQ.4YEkeMr-KPxLvFC1_oPo9J6nvY-zDeuXnGhm2U3kcSM

# Get me
```bash
curl -X GET http://localhost:8000/api/v1/users/me \
-H "Authorization: Bearer $TOKEN"
```

# Student List
```bash
curl -X GET http://localhost:8000/api/v1/students \
-H "Authorization: Bearer $TOKEN"
```

# Employer List
```bash
curl -X GET http://localhost:8000/api/v1/employers \
-H "Authorization: Bearer $TOKEN"
```

# Student profile
```bash
curl -X GET http://localhost:8000/api/v1/students/1 \
-H "Authorization: Bearer $TOKEN"
```

# Employer profile
```bash
curl -X GET http://localhost:8000/api/v1/employers/1 \
-H "Authorization: Bearer $TOKEN"
```

# Student remake
```bash
curl -X PUT http://localhost:8000/api/v1/students/1 \
-H "Authorization: Bearer $TOKEN" \
-H "Content-Type: application/json" \
-d '{
  "name": "Artem",
  "skills": ["Go", "Docker"],
  "about": "Backend developer"
}'
```

# Get Resume
```bash
curl -X GET http://localhost:8000/api/v1/students/1/resume \
-H "Authorization: Bearer $TOKEN"
```

curl: (26) Failed to open/read local data from file/application
# Upload Resume
```bash
curl -X POST http://localhost:8000/api/v1/students/USER_ID/resume \
-H "Authorization: Bearer $TOKEN" \
-F "file=@resume.pdf"
```

# Delete Resume
```bash
curl -X DELETE http://localhost:8000/api/v1/students/USER_ID/resume \
-H "Authorization: Bearer $TOKEN"
```