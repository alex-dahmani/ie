package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("patient", func() {
	DefaultMedia(PatientMedia)
	BasePath("/patients")
	Response(OK)
	Response(NotFound)
	Response(BadRequest)
	Response(InternalServerError)
	Action("show", func() {
		Description("Get patient by id.")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Patient ID")
		})
	})
	Action("list", func() {
		Routing(GET(""))
		Params(func() {
			Param("sort_by", String)
			Param("page", Integer, func() {
				Minimum(1)
			})
			Param("per_page", Integer, func() {
				Minimum(1)
			})
		})
		Description("List all patients.")
		Response(OK, func() {
			Media(CollectionOf(PatientMedia, func() {
				View("default")
			}))
		})
	})
})

var _ = Resource("care_team", func() {
	DefaultMedia(CareTeamMedia)
	BasePath("/care_teams")
	Response(OK)
	Response(NotFound)
	Response(BadRequest)
	Response(InternalServerError)
	Action("show", func() {
		Description("Get care team by id.")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", String, "Care team ID")
		})
	})
	Action("list", func() {
		Routing(GET(""))
		Description("List all care teams.")
		Response(OK, func() {
			Media(CollectionOf(CareTeamMedia, func() {
				View("default")
			}))
		})
	})
	Action("create", func() {
		Routing(POST(""))
		Payload(CareTeamPayload, func() {
			Required("name", "leader")
		})
		Description("Create care team.")
	})
	Action("update", func() {
		Routing(PUT("/:id"))
		Payload(CareTeamPayload, func() {
			Required("name", "leader")
		})
		Description("Update care team.")
	})
	Action("delete", func() {
		Routing(DELETE("/:id"))
		Description("Delete care team.")
	})
})

var _ = Resource("swagger", func() {
	Description("The API Swagger specification")
	Files("/swagger.json", "/swagger/swagger.json")
	Files("/swagger-ui/*filepath", "swagger-ui/")
})