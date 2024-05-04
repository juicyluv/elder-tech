// Package handlers provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package handlers

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Course Course
type Course struct {
	About        string         `json:"about"`
	AuthorId     int64          `json:"author_id"`
	Categories   []int16        `json:"categories"`
	CourseBlocks *[]CourseBlock `json:"course_blocks,omitempty"`
	CoverImage   *int64         `json:"cover_image,omitempty"`

	// CreatedAt A timestamp representing a date and time in RFC3339 format
	CreatedAt           Timestamp `json:"created_at"`
	Description         string    `json:"description"`
	Difficulty          int16     `json:"difficulty"`
	ForWho              string    `json:"for_who"`
	Id                  int32     `json:"id"`
	Progress            *float32  `json:"progress,omitempty"`
	Rating              *float32  `json:"rating,omitempty"`
	Requirements        string    `json:"requirements"`
	TimeToCompleteHours int16     `json:"time_to_complete_hours"`
	Title               string    `json:"title"`

	// UpdatedAt A timestamp representing a date and time in RFC3339 format
	UpdatedAt *Timestamp `json:"updated_at,omitempty"`
}

// CourseBlock defines model for CourseBlock.
type CourseBlock struct {
	CourseId    int32  `json:"course_id"`
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Number      int    `json:"number"`
	Title       string `json:"title"`
}

// CourseCategory defines model for CourseCategory.
type CourseCategory struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
}

// Error Error
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// SignInResponse defines model for SignInResponse.
type SignInResponse struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Patronymic *string `json:"patronymic,omitempty"`
	Surname    *string `json:"surname,omitempty"`
	Token      string  `json:"token"`
}

// Timestamp A timestamp representing a date and time in RFC3339 format
type Timestamp = time.Time

// User defines model for User.
type User struct {
	Age *int16 `json:"age,omitempty"`

	// CreatedAt A timestamp representing a date and time in RFC3339 format
	CreatedAt Timestamp `json:"created_at"`
	Email     *string   `json:"email,omitempty"`
	Gender    *int16    `json:"gender,omitempty"`
	Id        int64     `json:"id"`
	ImageId   *int64    `json:"image_id,omitempty"`

	// LastOnline A timestamp representing a date and time in RFC3339 format
	LastOnline *Timestamp `json:"lastOnline,omitempty"`
	Name       string     `json:"name"`
	Patronymic *string    `json:"patronymic,omitempty"`
	Phone      string     `json:"phone"`
	Surname    *string    `json:"surname,omitempty"`
}

// SignInJSONBody defines parameters for SignIn.
type SignInJSONBody struct {
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// SignUpJSONBody defines parameters for SignUp.
type SignUpJSONBody struct {
	Email    *string `json:"email,omitempty"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Phone    string  `json:"phone"`
	Surname  *string `json:"surname,omitempty"`
}

// CreateCourseJSONBody defines parameters for CreateCourse.
type CreateCourseJSONBody struct {
	About                 string  `json:"about"`
	Categories            []int16 `json:"categories"`
	Description           string  `json:"description"`
	Difficulty            int16   `json:"difficulty"`
	ForWho                string  `json:"forWho"`
	Requirements          string  `json:"requirements"`
	TimeToCompleteMinutes int16   `json:"timeToCompleteMinutes"`
	Title                 string  `json:"title"`
}

// AddCourseBlockJSONBody defines parameters for AddCourseBlock.
type AddCourseBlockJSONBody interface{}

// UpdateUserJSONBody defines parameters for UpdateUser.
type UpdateUserJSONBody struct {
	Age        *int16  `json:"age,omitempty"`
	Email      *string `json:"email,omitempty"`
	Gender     *int16  `json:"gender,omitempty"`
	Name       string  `json:"name"`
	Patronymic *string `json:"patronymic,omitempty"`
	Phone      string  `json:"phone"`
	Surname    *string `json:"surname,omitempty"`
}

// UpdateUserImageMultipartBody defines parameters for UpdateUserImage.
type UpdateUserImageMultipartBody struct {
	File *openapi_types.File `json:"file,omitempty"`
}

// SignInJSONRequestBody defines body for SignIn for application/json ContentType.
type SignInJSONRequestBody SignInJSONBody

// SignUpJSONRequestBody defines body for SignUp for application/json ContentType.
type SignUpJSONRequestBody SignUpJSONBody

// CreateCourseJSONRequestBody defines body for CreateCourse for application/json ContentType.
type CreateCourseJSONRequestBody CreateCourseJSONBody

// AddCourseBlockJSONRequestBody defines body for AddCourseBlock for application/json ContentType.
type AddCourseBlockJSONRequestBody AddCourseBlockJSONBody

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody UpdateUserJSONBody

// UpdateUserImageMultipartRequestBody defines body for UpdateUserImage for multipart/form-data ContentType.
type UpdateUserImageMultipartRequestBody UpdateUserImageMultipartBody
