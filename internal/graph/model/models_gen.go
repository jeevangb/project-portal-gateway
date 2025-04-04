// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type HealthStatus struct {
	Status  int32  `json:"status"`
	Message string `json:"message"`
}

type Mutation struct {
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Project struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	TechnologyStack []string `json:"technologyStack"`
	MentorName      string   `json:"mentorName"`
	Status          string   `json:"status"`
}

type ProjectInput struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	MentorName      string   `json:"mentorName"`
	TechnologyStack []string `json:"technologyStack"`
	Status          string   `json:"status"`
}

type Query struct {
}

type UpdateProjectInput struct {
	Title           *string  `json:"title,omitempty"`
	Description     *string  `json:"description,omitempty"`
	MentorName      *string  `json:"mentorName,omitempty"`
	TechnologyStack []string `json:"technologyStack,omitempty"`
	Status          *string  `json:"status,omitempty"`
}

type UserResponse struct {
	Message string  `json:"message"`
	Token   *string `json:"token,omitempty"`
}
