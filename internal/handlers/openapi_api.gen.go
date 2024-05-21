// Package handlers provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /auth/sign-in)
	SignIn(w http.ResponseWriter, r *http.Request)

	// (POST /auth/sign-up)
	SignUp(w http.ResponseWriter, r *http.Request)

	// (POST /courses)
	CreateCourse(w http.ResponseWriter, r *http.Request)

	// (GET /courses/author/{id})
	GetAuthorCourses(w http.ResponseWriter, r *http.Request, id int64)
	// Delete course block
	// (DELETE /courses/blocks/{id})
	DeleteCourseBlock(w http.ResponseWriter, r *http.Request, id int64)
	// Update course block
	// (PATCH /courses/blocks/{id})
	UpdateCourseBlock(w http.ResponseWriter, r *http.Request, id int64)
	// Get course block lessons
	// (GET /courses/blocks/{id}/lessons)
	GetCourseBlockLessons(w http.ResponseWriter, r *http.Request, id int64)
	// Add course block lesson
	// (POST /courses/blocks/{id}/lessons)
	AddCourseBlockLesson(w http.ResponseWriter, r *http.Request, id int64)

	// (GET /courses/categories)
	GetCourseCategories(w http.ResponseWriter, r *http.Request)
	// Delete course block lesson content
	// (DELETE /courses/content/{id})
	DeleteCourseBlockLessonContent(w http.ResponseWriter, r *http.Request, id int64)
	// Update course block lesson content
	// (PATCH /courses/content/{id})
	UpdateCourseBlockLessonContent(w http.ResponseWriter, r *http.Request, id int64)
	// Delete course block lesson
	// (DELETE /courses/lessons/{id})
	DeleteCourseBlockLesson(w http.ResponseWriter, r *http.Request, id int64)
	// Update course block lesson
	// (PATCH /courses/lessons/{id})
	UpdateCourseBlockLesson(w http.ResponseWriter, r *http.Request, id int64)
	// Get course block lesson contents
	// (GET /courses/lessons/{id}/content)
	GetCourseBlockLessonContents(w http.ResponseWriter, r *http.Request, id int64)
	// Add course block lesson content
	// (POST /courses/lessons/{id}/content)
	AddCourseBlockLessonContent(w http.ResponseWriter, r *http.Request, id int64)

	// (GET /courses/user/{id})
	GetUserCourses(w http.ResponseWriter, r *http.Request, id int64)

	// (DELETE /courses/{id})
	DeleteCourse(w http.ResponseWriter, r *http.Request, id int32)

	// (GET /courses/{id})
	GetCourse(w http.ResponseWriter, r *http.Request, id int32)

	// (PATCH /courses/{id})
	UpdateCourse(w http.ResponseWriter, r *http.Request, id int32)
	// Get course blocks
	// (GET /courses/{id}/blocks)
	GetCourseBlocks(w http.ResponseWriter, r *http.Request, id int32)
	// Add course block
	// (POST /courses/{id}/blocks)
	AddCourseBlock(w http.ResponseWriter, r *http.Request, id int32)

	// (POST /courses/{id}/join)
	JoinCourse(w http.ResponseWriter, r *http.Request, id int32)

	// (POST /courses/{id}/leave)
	LeaveCourse(w http.ResponseWriter, r *http.Request, id int32)

	// (GET /courses/{id}/members)
	GetCourseMembers(w http.ResponseWriter, r *http.Request, id int32)
	// Get image content by ID
	// (GET /images/{id})
	GetImage(w http.ResponseWriter, r *http.Request, id int64)

	// (GET /users/{id})
	GetUser(w http.ResponseWriter, r *http.Request, id int64)

	// (PATCH /users/{id})
	UpdateUser(w http.ResponseWriter, r *http.Request, id int64)

	// (PATCH /users/{id}/image)
	UpdateUserImage(w http.ResponseWriter, r *http.Request, id int64)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (POST /auth/sign-in)
func (_ Unimplemented) SignIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /auth/sign-up)
func (_ Unimplemented) SignUp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /courses)
func (_ Unimplemented) CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /courses/author/{id})
func (_ Unimplemented) GetAuthorCourses(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete course block
// (DELETE /courses/blocks/{id})
func (_ Unimplemented) DeleteCourseBlock(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update course block
// (PATCH /courses/blocks/{id})
func (_ Unimplemented) UpdateCourseBlock(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get course block lessons
// (GET /courses/blocks/{id}/lessons)
func (_ Unimplemented) GetCourseBlockLessons(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add course block lesson
// (POST /courses/blocks/{id}/lessons)
func (_ Unimplemented) AddCourseBlockLesson(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /courses/categories)
func (_ Unimplemented) GetCourseCategories(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete course block lesson content
// (DELETE /courses/content/{id})
func (_ Unimplemented) DeleteCourseBlockLessonContent(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update course block lesson content
// (PATCH /courses/content/{id})
func (_ Unimplemented) UpdateCourseBlockLessonContent(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete course block lesson
// (DELETE /courses/lessons/{id})
func (_ Unimplemented) DeleteCourseBlockLesson(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update course block lesson
// (PATCH /courses/lessons/{id})
func (_ Unimplemented) UpdateCourseBlockLesson(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get course block lesson contents
// (GET /courses/lessons/{id}/content)
func (_ Unimplemented) GetCourseBlockLessonContents(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add course block lesson content
// (POST /courses/lessons/{id}/content)
func (_ Unimplemented) AddCourseBlockLessonContent(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /courses/user/{id})
func (_ Unimplemented) GetUserCourses(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (DELETE /courses/{id})
func (_ Unimplemented) DeleteCourse(w http.ResponseWriter, r *http.Request, id int32) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /courses/{id})
func (_ Unimplemented) GetCourse(w http.ResponseWriter, r *http.Request, id int32) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (PATCH /courses/{id})
func (_ Unimplemented) UpdateCourse(w http.ResponseWriter, r *http.Request, id int32) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get course blocks
// (GET /courses/{id}/blocks)
func (_ Unimplemented) GetCourseBlocks(w http.ResponseWriter, r *http.Request, id int32) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Add course block
// (POST /courses/{id}/blocks)
func (_ Unimplemented) AddCourseBlock(w http.ResponseWriter, r *http.Request, id int32) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /courses/{id}/join)
func (_ Unimplemented) JoinCourse(w http.ResponseWriter, r *http.Request, id int32) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /courses/{id}/leave)
func (_ Unimplemented) LeaveCourse(w http.ResponseWriter, r *http.Request, id int32) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /courses/{id}/members)
func (_ Unimplemented) GetCourseMembers(w http.ResponseWriter, r *http.Request, id int32) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get image content by ID
// (GET /images/{id})
func (_ Unimplemented) GetImage(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /users/{id})
func (_ Unimplemented) GetUser(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (PATCH /users/{id})
func (_ Unimplemented) UpdateUser(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (PATCH /users/{id}/image)
func (_ Unimplemented) UpdateUserImage(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// SignIn operation middleware
func (siw *ServerInterfaceWrapper) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SignIn(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// SignUp operation middleware
func (siw *ServerInterfaceWrapper) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SignUp(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateCourse operation middleware
func (siw *ServerInterfaceWrapper) CreateCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateCourse(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetAuthorCourses operation middleware
func (siw *ServerInterfaceWrapper) GetAuthorCourses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAuthorCourses(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteCourseBlock operation middleware
func (siw *ServerInterfaceWrapper) DeleteCourseBlock(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteCourseBlock(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateCourseBlock operation middleware
func (siw *ServerInterfaceWrapper) UpdateCourseBlock(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateCourseBlock(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetCourseBlockLessons operation middleware
func (siw *ServerInterfaceWrapper) GetCourseBlockLessons(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCourseBlockLessons(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddCourseBlockLesson operation middleware
func (siw *ServerInterfaceWrapper) AddCourseBlockLesson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddCourseBlockLesson(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetCourseCategories operation middleware
func (siw *ServerInterfaceWrapper) GetCourseCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCourseCategories(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteCourseBlockLessonContent operation middleware
func (siw *ServerInterfaceWrapper) DeleteCourseBlockLessonContent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteCourseBlockLessonContent(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateCourseBlockLessonContent operation middleware
func (siw *ServerInterfaceWrapper) UpdateCourseBlockLessonContent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateCourseBlockLessonContent(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteCourseBlockLesson operation middleware
func (siw *ServerInterfaceWrapper) DeleteCourseBlockLesson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteCourseBlockLesson(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateCourseBlockLesson operation middleware
func (siw *ServerInterfaceWrapper) UpdateCourseBlockLesson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateCourseBlockLesson(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetCourseBlockLessonContents operation middleware
func (siw *ServerInterfaceWrapper) GetCourseBlockLessonContents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCourseBlockLessonContents(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddCourseBlockLessonContent operation middleware
func (siw *ServerInterfaceWrapper) AddCourseBlockLessonContent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddCourseBlockLessonContent(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUserCourses operation middleware
func (siw *ServerInterfaceWrapper) GetUserCourses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserCourses(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteCourse operation middleware
func (siw *ServerInterfaceWrapper) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteCourse(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetCourse operation middleware
func (siw *ServerInterfaceWrapper) GetCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCourse(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateCourse operation middleware
func (siw *ServerInterfaceWrapper) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateCourse(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetCourseBlocks operation middleware
func (siw *ServerInterfaceWrapper) GetCourseBlocks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCourseBlocks(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddCourseBlock operation middleware
func (siw *ServerInterfaceWrapper) AddCourseBlock(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddCourseBlock(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// JoinCourse operation middleware
func (siw *ServerInterfaceWrapper) JoinCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.JoinCourse(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// LeaveCourse operation middleware
func (siw *ServerInterfaceWrapper) LeaveCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LeaveCourse(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetCourseMembers operation middleware
func (siw *ServerInterfaceWrapper) GetCourseMembers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCourseMembers(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetImage operation middleware
func (siw *ServerInterfaceWrapper) GetImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetImage(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUser operation middleware
func (siw *ServerInterfaceWrapper) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUser(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUser(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateUserImage operation middleware
func (siw *ServerInterfaceWrapper) UpdateUserImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, JWTAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUserImage(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/sign-in", wrapper.SignIn)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/sign-up", wrapper.SignUp)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/courses", wrapper.CreateCourse)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/courses/author/{id}", wrapper.GetAuthorCourses)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/courses/blocks/{id}", wrapper.DeleteCourseBlock)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/courses/blocks/{id}", wrapper.UpdateCourseBlock)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/courses/blocks/{id}/lessons", wrapper.GetCourseBlockLessons)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/courses/blocks/{id}/lessons", wrapper.AddCourseBlockLesson)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/courses/categories", wrapper.GetCourseCategories)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/courses/content/{id}", wrapper.DeleteCourseBlockLessonContent)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/courses/content/{id}", wrapper.UpdateCourseBlockLessonContent)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/courses/lessons/{id}", wrapper.DeleteCourseBlockLesson)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/courses/lessons/{id}", wrapper.UpdateCourseBlockLesson)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/courses/lessons/{id}/content", wrapper.GetCourseBlockLessonContents)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/courses/lessons/{id}/content", wrapper.AddCourseBlockLessonContent)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/courses/user/{id}", wrapper.GetUserCourses)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/courses/{id}", wrapper.DeleteCourse)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/courses/{id}", wrapper.GetCourse)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/courses/{id}", wrapper.UpdateCourse)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/courses/{id}/blocks", wrapper.GetCourseBlocks)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/courses/{id}/blocks", wrapper.AddCourseBlock)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/courses/{id}/join", wrapper.JoinCourse)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/courses/{id}/leave", wrapper.LeaveCourse)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/courses/{id}/members", wrapper.GetCourseMembers)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/images/{id}", wrapper.GetImage)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users/{id}", wrapper.GetUser)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/users/{id}", wrapper.UpdateUser)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/users/{id}/image", wrapper.UpdateUserImage)
	})

	return r
}
