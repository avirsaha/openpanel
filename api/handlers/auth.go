package handlers

import (
    "encoding/json"
    "net/http"
    "gopanel/internal/auth"
    "gopanel/api/response"
)

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        response.Error(w, http.StatusBadRequest, "Invalid input")
        return
    }

    token, err := auth.Authenticate(req.Username, req.Password)
    if err != nil {
        response.Error(w, http.StatusUnauthorized, "Login failed")
        return
    }

    response.JSON(w, http.StatusOK, map[string]string{"token": token})
}
