package domain

type FormStatus string

const (
	FormStatusDraft     FormStatus = "DRAFT"
	FormStatusPublished FormStatus = "PUBLISHED"
	FormStatusArchived  FormStatus = "ARCHIVED"
	FormStatusDeleted   FormStatus = "DELETED"
)

var FormStatusOptions = []FormStatus{
	FormStatusDraft,
	FormStatusPublished,
	FormStatusArchived,
	FormStatusDeleted,
}
