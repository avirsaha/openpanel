package handlers

import (
    "encoding/json"
    "net/http"

    "gopanel/internal/web"
    "gopanel/api/response"
)

type CreateWebsiteRequest struct {
    Domain   string `json:"domain"`
    Username string `json:"username"`
}

func CreateWebsite(w http.ResponseWriter, r *http.Request) {
    var req CreateWebsiteRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        response.Error(w, http.StatusBadRequest, "Invalid request body")
        return
    }

    // Basic validation
    if req.Domain == "" || req.Username == "" {
        response.Error(w, http.StatusBadRequest, "Domain and username are required")
        return
    }

    err := web.CreateWebsite(web.WebsiteConfig{
        Domain:   req.Domain,
        Username: req.Username,
    })

    if err != nil {
        response.Error(w, http.StatusInternalServerError, err.Error())
        return
    }

    response.JSON(w, http.StatusCreated, map[string]string{
        "message": "Website created successfully",
        "domain":  req.Domain,
    })
}
