package main

import "github.com/goadesign/goa"

// WebsiteController implements the website resource.
type WebsiteController struct {
	*goa.Controller
}

// NewWebsiteController creates a website controller.
func NewWebsiteController(service *goa.Service) *WebsiteController {
	return &WebsiteController{Controller: service.NewController("WebsiteController")}
}
