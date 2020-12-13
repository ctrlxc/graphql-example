{{- $alias := .Aliases.Table .Table.Name}}

// WARNING: required ID column
func (o *{{$alias.UpSingular}}) GlobalID() string {
	return util.ToGlobalID("{{$alias.UpSingular}}", o.ID)
}
