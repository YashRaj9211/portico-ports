package traefik

import (
	"fmt"
	"os"
	"text/template"
)


const CONFIG_DIR = "/etc/traefik/dynamic"

const ROUTE_TEMPLATE = `
http:
  routers:
    {{.Slug}}:
      rule: "Host(` + "`{{.Host}}`" + `)"
      service: "{{.Slug}}"
      entryPoints:
        - web

  services:
    {{.Slug}}:
      loadBalancer:
        servers:
          - url: "http://localhost:3000"
`

type RouteConfig struct {
	Slug string 
	Host string
}


func WriteRouteConfig(slug, host string) error {
	filename := fmt.Sprintf("%s/%s.yml", CONFIG_DIR, slug)

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot create traefik config file: %w", err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("route").Parse(ROUTE_TEMPLATE))
	return tmpl.Execute(f, RouteConfig{Slug: slug, Host: host})
}