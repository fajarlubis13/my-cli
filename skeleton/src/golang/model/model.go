package model

{{ if .TableStatus }}import "time"{{ end }}

// {{ toCamel .ProjectName }} ...
type {{ toCamel .ProjectName }} struct {
	{{- range $e := .Entities }}
	{{ toCamel $e.Name }} {{ $e.Type }} `json:"{{ toDelimeted $e.Name 95 }}"{{ if $e.Binding }} binding:"required"{{ end }}`
	{{- end -}}
	{{ if .TableStatus }}
	
	CreatedAt            *time.Time `json:"-"`
	CreatedBy            int64      `json:"-"`
	UpdatedAt            *time.Time `json:"-"`
	UpdatedBy            int64      `json:"-"`
	DeletedAt            *time.Time `json:"-"`
	DeletedBy            int64      `json:"-"`	
	{{- end }}
}
