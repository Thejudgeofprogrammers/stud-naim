# 🚀 Trampolin Platform API

## Backend API для платформы «Трамплин» — системы взаимодействия студентов и работодателей

---

## 📌 Базовая информация

**Base URL:**
http://localhost:8000/api/v1

**Авторизация:**
Authorization: Bearer <access_token>

**Формат данных:**
Content-Type: application/json

---

# 🔐 AUTH

## 🟢 Регистрация

POST /auth/register  
Авторизация: ❌ Не требуется  

### Request
```json
{
  "email": "test@test.com",
  "password": "123456",
  "role": "student"
}
```

### Response
`201 created`

# 🔐 AUTH

## 🟢 Логин

POST /auth/login  
Авторизация: ❌ Не требуется  

### Request
```json
{
  "email": "test@test.com",
  "password": "123456"
}
```

### Response
```json
{
  "access_token": "JWT",
  "refresh_token": "REFRESH"
}
```
## 👤 USERS
**🟡 Получить текущего пользователя**

`GET /users/me`

Авторизация: ✅

### Response
```json
{
  "id": "user_id",
  "role": "student"
}
```

**🟡 Список студентов**

`GET /students`

## Авторизация: ✅

### Response
```json
[
  {
    "id": "1",
    "name": "Artem",
    "skills": ["Go", "Docker"],
    "about": "Backend developer",
    "resume_url": ""
  }
]
```

**🟡 Профиль студента**

`GET /students/:id`

## Авторизация: ✅

### Response
```json
{
  "id": "1",
  "name": "Artem",
  "skills": ["Go"],
  "about": "Backend developer",
  "resume_url": ""
}
```

### 🔒 Обновить студента

`PUT /students/:id`

### Авторизация: ✅ (owner)

### Request
```json
{
  "name": "Artem",
  "skills": ["Go", "Docker"],
  "about": "Backend developer"
}
```

### Response

`204 No Content`

**🟡 Список работодателей**

`GET /employers`

## Авторизация:

### Response
```json
[
  {
    "id": "1",
    "company_name": "Company"
  }
]
```

**🟡 Профиль работодателя**

`GET /employers/:id`

## Авторизация: ✅

### Response
```json
{
  "id": "1",
  "company_name": "Company",
  "description": "IT company",
  "representative": "John Doe",
  "vacancies": []
}
```

### 🔒 Обновить работодателя

`PUT /employers/:id`

## Авторизация: ✅ (owner)

### Request
```json
{
  "company_name": "New Company",
  "description": "Updated",
  "representative": "CEO"
}
```

## 📄 RESUME
### 🔒 Получить резюме

`GET /students/:id/resume`

## Авторизация: ✅ (owner)

### Response
```json
{
  "url": "/uploads/file.pdf"
}
```

### 🔒 Загрузить резюме

`POST /students/:id/resume`

#### Авторизация: ✅ (owner)

```json 
Form-data
file: resume.pdf
```

### Response
```json
{
  "url": "/uploads/userid_resume.pdf"
}
```

### 🔒 Удалить резюме

`DELETE /students/:id/resume`

#### Авторизация: ✅ (owner)

## 💼 OPPORTUNITIES

**🟡 Список возможностей**

`GET /opportunities`

### Авторизация: ✅

### Response
```json
[
  {
    "id": "1",
    "title": "Go Developer",
    "description": "Backend dev",
    "company_id": "123",
    "type": "job",
    "format": "remote",
    "location": "Riga",
    "salary": 1000,
    "tags": ["Go"],
    "created_at": "...",
    "expired_at": "..."
  }
]
```

**🟡 Фильтр**

`GET /opportunities/filter?tag=go&format=remote`

### Авторизация: ✅

**🟡 Получить одну**

`GET /opportunities/:id`

### Авторизация: ✅

### 🔒 Создать

`POST /opportunities`

### Авторизация: ✅ (employer)

### Request

```json
{
  "title": "Go Dev",
  "description": "Backend",
  "type": "job",
  "format": "remote",
  "location": "Riga",
  "salary": 1000,
  "tags": ["Go", "Docker"]
}
```

### 🔒 Обновить

`PUT /opportunities/:id`

### Авторизация: ✅ (owner)

## 🔒 Удалить

`DELETE /opportunities/:id`

### Авторизация: ✅ (owner)

## 🔐 Роли


| Роль |    Описание |
| ---- | ---- |
| student |	поиск работы |
| employer |	публикация вакансий |
| curator |	модерация |
| admin |	управление кураторами |

### ⚠️ Ошибки

```json
{
  "error": "message"
}
```

### 🚀 Запуск

`make run-dev`
