{{- $alias := .Aliases.Table .Table.Name}}

// WARNING: required ID column
func (o *{{$alias.UpSingular}}) GlobalID() string {
	return globalid.ToGlobalID("{{$alias.UpSingular}}", o.ID)
}
