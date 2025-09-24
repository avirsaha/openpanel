package handlers

import (
    "encoding/json"
    "net/http"

    "gopanel/internal/users"
    "gopanel/api/response"
)

// Struct to handle POST /api/users
type CreateUserRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// GET /api/users
func GetUsers(w http.ResponseWriter, r *http.Request) {
    userList, err := users.ListUsers()
    if err != nil {
        response.Error(w, http.StatusInternalServerError, "Failed to list users")
        return
    }

    response.JSON(w, http.StatusOK, map[string]any{
        "users": userList,
    })
}

// POST /api/users
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        response.Error(w, http.StatusBadRequest, "Invalid request body")
        return
    }

    // Basic validation
    if req.Username == "" || req.Password == "" {
        response.Error(w, http.StatusBadRequest, "Username and password required")
        return
    }

    err := users.CreateUser(req.Username, req.Password)
    if err != nil {
        response.Error(w, http.StatusBadRequest, err.Error())
        return
    }

    response.JSON(w, http.StatusCreated, map[string]string{
        "message": "User created successfully",
        "username": req.Username,
    })
}
