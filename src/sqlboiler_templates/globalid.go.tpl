{{- $alias := .Aliases.Table .Table.Name}}

{{- range $column := .Table.Columns -}}
	{{- $colAlias := $alias.Column $column.Name -}}
	{{- if eq $colAlias "ID" -}}
		// convert id -> globalid
		func (o *{{$alias.UpSingular}}) GlobalID() string {
			return globalid.ToGlobalID("{{$alias.UpSingular}}", o.ID)
		}
	{{- end -}}
{{- end }}
