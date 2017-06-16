// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "api": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/intervention-engine/ie/design
// --out=$(GOPATH)/src/github.com/intervention-engine/ie
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// CareTeamController is the controller interface for the CareTeam actions.
type CareTeamController interface {
	goa.Muxer
	AddPatient(*AddPatientCareTeamContext) error
	Create(*CreateCareTeamContext) error
	Delete(*DeleteCareTeamContext) error
	Huddles(*HuddlesCareTeamContext) error
	List(*ListCareTeamContext) error
	RemovePatient(*RemovePatientCareTeamContext) error
	Schedule(*ScheduleCareTeamContext) error
	Show(*ShowCareTeamContext) error
	Update(*UpdateCareTeamContext) error
}

// MountCareTeamController "mounts" a CareTeam resource controller on the given service.
func MountCareTeamController(service *goa.Service, ctrl CareTeamController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/care_teams/:id/patients/:patient_id", ctrl.MuxHandler("preflight", handleCareTeamOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/care_teams", ctrl.MuxHandler("preflight", handleCareTeamOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/care_teams/:id", ctrl.MuxHandler("preflight", handleCareTeamOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/care_teams/:id/huddles", ctrl.MuxHandler("preflight", handleCareTeamOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddPatientCareTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.AddPatient(rctx)
	}
	h = handleCareTeamOrigin(h)
	service.Mux.Handle("POST", "/api/care_teams/:id/patients/:patient_id", ctrl.MuxHandler("add_patient", h, nil))
	service.LogInfo("mount", "ctrl", "CareTeam", "action", "AddPatient", "route", "POST /api/care_teams/:id/patients/:patient_id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateCareTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateCareTeamPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleCareTeamOrigin(h)
	service.Mux.Handle("POST", "/api/care_teams", ctrl.MuxHandler("create", h, unmarshalCreateCareTeamPayload))
	service.LogInfo("mount", "ctrl", "CareTeam", "action", "Create", "route", "POST /api/care_teams")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteCareTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleCareTeamOrigin(h)
	service.Mux.Handle("DELETE", "/api/care_teams/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "CareTeam", "action", "Delete", "route", "DELETE /api/care_teams/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewHuddlesCareTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Huddles(rctx)
	}
	h = handleCareTeamOrigin(h)
	service.Mux.Handle("GET", "/api/care_teams/:id/huddles", ctrl.MuxHandler("huddles", h, nil))
	service.LogInfo("mount", "ctrl", "CareTeam", "action", "Huddles", "route", "GET /api/care_teams/:id/huddles")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListCareTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleCareTeamOrigin(h)
	service.Mux.Handle("GET", "/api/care_teams", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "CareTeam", "action", "List", "route", "GET /api/care_teams")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRemovePatientCareTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.RemovePatient(rctx)
	}
	h = handleCareTeamOrigin(h)
	service.Mux.Handle("DELETE", "/api/care_teams/:id/patients/:patient_id", ctrl.MuxHandler("remove_patient", h, nil))
	service.LogInfo("mount", "ctrl", "CareTeam", "action", "RemovePatient", "route", "DELETE /api/care_teams/:id/patients/:patient_id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewScheduleCareTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ScheduleCareTeamPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Schedule(rctx)
	}
	h = handleCareTeamOrigin(h)
	service.Mux.Handle("POST", "/api/care_teams/:id/huddles", ctrl.MuxHandler("schedule", h, unmarshalScheduleCareTeamPayload))
	service.LogInfo("mount", "ctrl", "CareTeam", "action", "Schedule", "route", "POST /api/care_teams/:id/huddles")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowCareTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleCareTeamOrigin(h)
	service.Mux.Handle("GET", "/api/care_teams/:id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "CareTeam", "action", "Show", "route", "GET /api/care_teams/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateCareTeamContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CareTeamPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleCareTeamOrigin(h)
	service.Mux.Handle("PUT", "/api/care_teams/:id", ctrl.MuxHandler("update", h, unmarshalUpdateCareTeamPayload))
	service.LogInfo("mount", "ctrl", "CareTeam", "action", "Update", "route", "PUT /api/care_teams/:id")
}

// handleCareTeamOrigin applies the CORS response headers corresponding to the origin.
func handleCareTeamOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateCareTeamPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateCareTeamPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createCareTeamPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalScheduleCareTeamPayload unmarshals the request body into the context request data Payload field.
func unmarshalScheduleCareTeamPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &scheduleCareTeamPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateCareTeamPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateCareTeamPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &careTeamPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// HuddleController is the controller interface for the Huddle actions.
type HuddleController interface {
	goa.Muxer
	Cancel(*CancelHuddleContext) error
}

// MountHuddleController "mounts" a Huddle resource controller on the given service.
func MountHuddleController(service *goa.Service, ctrl HuddleController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/huddles/:id/patients/:patient_id", ctrl.MuxHandler("preflight", handleHuddleOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCancelHuddleContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Cancel(rctx)
	}
	h = handleHuddleOrigin(h)
	service.Mux.Handle("DELETE", "/api/huddles/:id/patients/:patient_id", ctrl.MuxHandler("cancel", h, nil))
	service.LogInfo("mount", "ctrl", "Huddle", "action", "Cancel", "route", "DELETE /api/huddles/:id/patients/:patient_id")
}

// handleHuddleOrigin applies the CORS response headers corresponding to the origin.
func handleHuddleOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// PatientController is the controller interface for the Patient actions.
type PatientController interface {
	goa.Muxer
	List(*ListPatientContext) error
	Show(*ShowPatientContext) error
}

// MountPatientController "mounts" a Patient resource controller on the given service.
func MountPatientController(service *goa.Service, ctrl PatientController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/patients", ctrl.MuxHandler("preflight", handlePatientOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/patients/:id", ctrl.MuxHandler("preflight", handlePatientOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListPatientContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handlePatientOrigin(h)
	service.Mux.Handle("GET", "/api/patients", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Patient", "action", "List", "route", "GET /api/patients")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowPatientContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handlePatientOrigin(h)
	service.Mux.Handle("GET", "/api/patients/:id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Patient", "action", "Show", "route", "GET /api/patients/:id")
}

// handlePatientOrigin applies the CORS response headers corresponding to the origin.
func handlePatientOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger-ui/*filepath", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger-ui/*filepath", "swagger-ui/")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger-ui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/", "route", "GET /swagger-ui/*filepath")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger-ui/", "swagger-ui/index.html")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger-ui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/index.html", "route", "GET /swagger-ui/")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
