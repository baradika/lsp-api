## Daftar Endpoint

### Autentikasi

#### Register

- **URL**: `/api/v1/auth/register`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "Baradika",
    "full_name": "Fase Rais Baradika",
    "email": "pakau@gmail.com",
    "password": "password123"
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "message": "User registered successfully",
    "data": {
      "id": 1,
      "username": "Baradika",
      "email": "pakau@gmail.com"
    }
  }
  ```

#### Login

- **URL**: `/api/v1/auth/login`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "email": "pakau@gmail.com",
    "password": "password123"
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "message": "Login successful",
    "data": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
  }
  ```

#### Logout

- **URL**: `/api/v1/auth/logout`
- **Method**: `POST`
- **Headers**: `Authorization: Bearer {token}`
- **Response**:
  ```json
  {
    "success": true,
    "message": "Logged out successfully"
  }
  ```

### Manajemen Asesor

#### Membuat Asesor Baru

- **URL**: `/api/v1/assessors`
- **Method**: `POST`
- **Headers**: `Authorization: Bearer {token}`
- **Request Body**:
  ```json
  {
    "full_name": "Jane Smith",
    "registration": "ASR-001",
    "email": "jane@example.com",
    "phone": "08123456789",
    "competencies": [
      {
        "id": 1,
        "name": "Web Development",
        "code": "WD-001",
        "description": "Competency in web development"
      }
    ]
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "message": "Assessor created successfully",
    "data": {
      "id": 1,
      "full_name": "Jane Smith",
      "registration": "ASR-001",
      "email": "jane@example.com",
      "phone": "08123456789",
      "competencies": [
        {
          "id": 1,
          "name": "Web Development",
          "code": "WD-001",
          "description": "Competency in web development"
        }
      ],
      "is_complete": false,
      "created_at": "2023-10-15T10:30:00Z",
      "updated_at": "2023-10-15T10:30:00Z"
    }
  }
  ```

#### Mendapatkan Semua Asesor

- **URL**: `/api/v1/assessors`
- **Method**: `GET`
- **Headers**: `Authorization: Bearer {token}`
- **Response**:
  ```json
  {
    "success": true,
    "message": "Assessors retrieved successfully",
    "data": [
      {
        "id": 1,
        "full_name": "Jane Smith",
        "registration": "ASR-001",
        "email": "jane@example.com",
        "phone": "08123456789",
        "competencies": [
          {
            "id": 1,
            "name": "Web Development",
            "code": "WD-001",
            "description": "Competency in web development"
          }
        ],
        "is_complete": false,
        "created_at": "2023-10-15T10:30:00Z",
        "updated_at": "2023-10-15T10:30:00Z"
      }
    ]
  }
  ```

#### Mendapatkan Asesor Berdasarkan ID

- **URL**: `/api/v1/assessors/{id}`
- **Method**: `GET`
- **Headers**: `Authorization: Bearer {token}`
- **Response**:
  ```json
  {
    "success": true,
    "message": "Assessor retrieved successfully",
    "data": {
      "id": 1,
      "full_name": "Jane Smith",
      "registration": "ASR-001",
      "email": "jane@example.com",
      "phone": "08123456789",
      "competencies": [
        {
          "id": 1,
          "name": "Web Development",
          "code": "WD-001",
          "description": "Competency in web development"
        }
      ],
      "is_complete": false,
      "created_at": "2023-10-15T10:30:00Z",
      "updated_at": "2023-10-15T10:30:00Z"
    }
  }
  ```

#### Memperbarui Asesor

- **URL**: `/api/v1/assessors/{id}`
- **Method**: `PUT`
- **Headers**: `Authorization: Bearer {token}`
- **Request Body**:
  ```json
  {
    "full_name": "Jane Smith Updated",
    "registration": "ASR-001",
    "email": "jane.updated@example.com",
    "phone": "08123456789",
    "competencies": [
      {
        "id": 1,
        "name": "Web Development",
        "code": "WD-001",
        "description": "Competency in web development"
      },
      {
        "id": 2,
        "name": "Mobile Development",
        "code": "MD-001",
        "description": "Competency in mobile development"
      }
    ]
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "message": "Assessor updated successfully",
    "data": {
      "id": 1,
      "full_name": "Jane Smith Updated",
      "registration": "ASR-001",
      "email": "jane.updated@example.com",
      "phone": "08123456789",
      "competencies": [
        {
          "id": 1,
          "name": "Web Development",
          "code": "WD-001",
          "description": "Competency in web development"
        },
        {
          "id": 2,
          "name": "Mobile Development",
          "code": "MD-001",
          "description": "Competency in mobile development"
        }
      ],
      "is_complete": false,
      "created_at": "2023-10-15T10:30:00Z",
      "updated_at": "2023-10-15T11:15:00Z"
    }
  }
  ```

#### Menghapus Asesor

- **URL**: `/api/v1/assessors/{id}`
- **Method**: `DELETE`
- **Headers**: `Authorization: Bearer {token}`
- **Response**:
  ```json
  {
    "success": true,
    "message": "Assessor deleted successfully"
  }
  ```

#### Memperbarui Status Kelengkapan Asesor

- **URL**: `/api/v1/assessors/{id}/completeness`
- **Method**: `PATCH`
- **Headers**: `Authorization: Bearer {token}`
- **Request Body**:
  ```json
  {
    "is_complete": true
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "message": "Assessor completeness updated successfully"
  }
  ```
