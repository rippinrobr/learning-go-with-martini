package services

import (
  "net/url"

  "github.com/codegangsta/martini-contrib/binding"
)

type Service struct {
  Name string `json:"name"`
  URL string `json:"name"`
  Description string `json:"description"`
}

func (service *Service) Validate( errors *binding.Errors, req *http.Request ) {
  if service.name == "" {
    errors.Overall["missing-requirement"] = "name is a required field";
  }

  if service.URL == "" {
    errors.Overall["missing-requirement"] = "URL is a required field";
  } else {
    _, err = url.Parse( service.URL )
    if err != nil {
      errors.Overall["missing-requirement"] = service.URL + " is not a valid URL";
    }
  }
}

func RegisterService( service Service ) bool {
}
