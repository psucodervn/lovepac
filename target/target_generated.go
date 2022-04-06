// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2022-04-07 02:03:24.193461 +0700 +07 m=+0.001588834
// TODO add the commit hash in here too

package target

import (
	"text/template"
)

var loveTemplate = template.Must(template.New("love").Parse(`local quads = {}

{{range .Sprites -}}
quads['{{.Name}}'] = love.graphics.newQuad({{.Left}},{{.Top}},{{.Width}},{{.Height}},{{$.Width}},{{$.Height}})
{{end}}
return quads
`))

var spineTemplate = template.Must(template.New("spine").Parse(`{{.ImageFilename}}
size:{{.Width}},{{.Height}}
scale:{{.Scale}}
{{- range .Sprites}}
{{.DisplayName}}
bounds:{{.Left}},{{.Top}},{{.Width}},{{.Height}}
{{- end}}

`))

var starlingTemplate = template.Must(template.New("starling").Parse(`<TextureAtlas imagePath="{{.ImageFilename}}">
{{- range .Sprites}}
    <SubTexture name="{{.Name}}" x="{{.Left}}" y="{{.Top}}" width="{{.Width}}" height="{{.Height}}"/>
{{- end}}
</TextureAtlas>
`))
