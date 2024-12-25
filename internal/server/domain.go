package server

import "github.com/gin-gonic/gin"

func loadDIDDomainMoudles(r *gin.RouterGroup, h *handle) {
	r.POST("/bind", h.domainBind)
	r.POST("/renewal", h.domainRenewal)
}

// @Summary		DomainBind
// @Description	DomainBind
// @Tags			domain
// @Accept			json
// @Produce		json
// @Param			domain	body		string	true	"domain"
// @Param			address	body		string	true	"address"
// @Success		200		{string}	string	"ok"
// @Router			/domain/bind [post]
func (h *handle) domainBind(c *gin.Context) {}

// @Summary		DomainRenewal
// @Description	DomainRenewal
// @Tags			domain
// @Accept			json
// @Produce		json
// @Param			domain	body		string	true	"domain"
// @Success		200		{string}	string	"ok"
// @Router			/domain/renewal [post]
func (h *handle) domainRenewal(c *gin.Context) {}
