package main

var src = `
// Code generated DO NOT EDIT
// Generated by the Indigo gentype tool
package {{ .PackageName }}

import (
    "reflect"
    "fmt"

    "github.com/ezachrisen/indigo/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"


	{{ range $key, $value :=  .Imports }}
	{{ $value }} "{{ $key }}"
	{{ end }}
)

{{$packageName:=.PackageName}}

{{ range .Objects }}


    var {{ .Name }}Type = types.NewTypeValue("{{$packageName}}.{{ .Name }} ", traits.IndexerType)

    func (v {{ .Name }}) ConvertToNative(typeDesc reflect.Type) (interface{}, error) {
	//log.Println("{{ .Name }}ConvertToNative")
	return nil, fmt.Errorf("cannot convert attribute message to native types")
    }
    
    func (v {{ .Name }}) ConvertToType(typeValue ref.Type) ref.Val {
	//log.Println("{{ .Name }}.ConvertToType")
	return types.NewErr("cannot convert attribute message to CEL types")
    }
    
    func (v {{ .Name }}) Equal(other ref.Val) ref.Val {
	//log.Println("{{ .Name }}.Equal")
	return types.NewErr("attribute message does not support equality")
    }
    
    func (v {{ .Name }}) Type() ref.Type {
	//log.Println("{{ .Name }}.Type")
	return {{ .Name }}Type
    }
    
    func (v {{ .Name }}) Value() interface{} {
	//log.Println("{{ .Name }}.Value")
	return v
    }

    func (v {{ .Name }}) Get(index ref.Val) ref.Val {
	//log.Println("{{ .Name }}.Get ", index, index.Type(), index.Value(), index.Type().TypeName())
	field, ok := index.Value().(string)

	if !ok {
		return types.NewErr("Field %v not found in type %s", index.Value(), "{{ .Name }}")
	}

	switch field {

	{{ range .Fields  }}
	    
	/* Code generation debug information
	    Type: {{ .Type }}
	    Name: {{ .Type.TypeName }}
	    Mult: {{ .Type.Multiple }}
	    ID  : {{ .Type.TypeID }}		
	    Obj : {{ .Type.IsObject }}
	 */

	    case "{{ .Name }}":
	     {{ if eq .Type.TypeName  "string"  }}
	     	 return types.String(v.{{ .Name }})
	     {{ else if eq .Type.TypeName  "int" }}
	     	 return types.Int(v.{{ .Name }})
	     {{ else if eq .Type.TypeName  "float64" }}
	     	 return types.Double(v.{{ .Name }})
	     {{ else }}
	     	return v.{{ .Name }}
	     {{ end }}
 
   {{ end }} 

	default:
		return nil
        }
     }

    func (v {{ .Name }}) ProvideStructDefintion() cel.StructDefinition {
    	 return cel.StructDefinition{
	   Name: "{{$packageName}}.{{ .Name }}",
	   Fields: map[string]*ref.FieldType{
	      {{ range .Fields }}
	      "{{ .Name }}": 
	      	     {{ if eq .Type.TypeName  "string"  }}
	     	     	 &ref.FieldType{Type: decls.String},
	              {{ else if eq .Type.TypeName  "int" }}
	     	         &ref.FieldType{Type: decls.Int},
	              {{ else if eq .Type.TypeName  "float64" }}
     		         &ref.FieldType{Type: decls.Double}, 
				{{ end }}
		   {{ end }}			 
		  },		   			   	   
	   }
    }

{{end}}
`