// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Ksonnet Registry API documentation\n",
    "title": "Ksonnet Registry API",
    "contact": {
      "name": "bryan@heptio.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "0.1.0"
  },
  "host": "localhost:5000",
  "basePath": "/",
  "paths": {
    "/api/v1/packages": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "package"
        ],
        "summary": "List packages",
        "operationId": "listPackages",
        "parameters": [
          {
            "type": "string",
            "description": "Filter by namespace",
            "name": "namespace",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Lookup value for package search",
            "name": "query",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Packages"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/v1/packages/{namespace}/{package}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "package"
        ],
        "summary": "List all releases for a package",
        "operationId": "showPackageReleases",
        "parameters": [
          {
            "type": "string",
            "description": "namespace",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "package name",
            "name": "package",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/PackageManifest"
            }
          },
          "401": {
            "description": "Not authorized to read the package",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Package not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "package"
        ],
        "summary": "Push new package release to the registry",
        "operationId": "createPackage",
        "parameters": [
          {
            "type": "string",
            "description": "Namespace for package",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Package for release",
            "name": "package",
            "in": "path",
            "required": true
          },
          {
            "type": "boolean",
            "default": false,
            "description": "Force push the release (if allowed)",
            "name": "force",
            "in": "query"
          },
          {
            "description": "Package object to be added to the registry",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PostPackage"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Package"
            }
          },
          "401": {
            "description": "Not authorized to create the package",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Package not found (if force=true)",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "409": {
            "description": "Package already exists",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "422": {
            "description": "Bad version or name format",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/v1/packages/{namespace}/{package}/blobs/sha256/{digest}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/x-gzip"
        ],
        "tags": [
          "blobs"
        ],
        "summary": "Pull a package blob by digest",
        "operationId": "pullBlob",
        "parameters": [
          {
            "type": "string",
            "description": "namespace",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "package name",
            "name": "package",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "content digest",
            "name": "digest",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "file",
              "title": "package-targz"
            }
          },
          "401": {
            "description": "Not authorized to read the package",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Package not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/v1/packages/{namespace}/{package}/{release}": {
      "get": {
        "description": "show package release",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "package"
        ],
        "summary": "show package release",
        "operationId": "showPackageRelease",
        "parameters": [
          {
            "type": "string",
            "description": "namespace",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "package name",
            "name": "package",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "release",
            "name": "release",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Manifest"
            }
          },
          "401": {
            "description": "Not authorized to read the package",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Release not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "package"
        ],
        "summary": "Delete a package release",
        "operationId": "deletePackageRelease",
        "parameters": [
          {
            "type": "string",
            "description": "namespace",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "package name",
            "name": "package",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "release name",
            "name": "release",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "successful operation"
          },
          "401": {
            "description": "Not authorized to read the package",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Package not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/v1/packages/{namespace}/{package}/{release}/pull": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/x-gzip"
        ],
        "tags": [
          "package",
          "blobs"
        ],
        "summary": "Download the package",
        "operationId": "pullPackage",
        "parameters": [
          {
            "type": "string",
            "description": "namespace",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "package name",
            "name": "package",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "release name",
            "name": "release",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "file",
              "title": "package-targz"
            }
          },
          "401": {
            "description": "Not authorized to read the package",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Package not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/version": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "info"
        ],
        "summary": "Display api version",
        "operationId": "getVersion",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Dependency": {
      "type": "object",
      "title": "dependency",
      "properties": {
        "constraint": {
          "description": "Dependency constraint",
          "type": "string"
        },
        "name": {
          "description": "Dependency name",
          "type": "string"
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "description": "http status code",
          "type": "integer",
          "title": "http-code"
        },
        "details": {
          "description": "error extra data",
          "type": "object"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      }
    },
    "Manifest": {
      "type": "object",
      "title": "manifest",
      "properties": {
        "content": {
          "$ref": "#/definitions/PartDescriptor"
        },
        "created_at": {
          "description": "creation data",
          "type": "string",
          "format": "date-time",
          "title": "created-at"
        },
        "metadata": {
          "description": "KeyValue object to add complementary and format specific information",
          "type": "object",
          "title": "metadata"
        },
        "package": {
          "description": "package name",
          "type": "string",
          "title": "package-name"
        },
        "release": {
          "description": "release name",
          "type": "string",
          "title": "release-name"
        }
      }
    },
    "Package": {
      "description": "Package object",
      "type": "object",
      "title": "Package",
      "properties": {
        "content": {
          "$ref": "#/definitions/Manifest"
        },
        "created_at": {
          "description": "Package creation date",
          "type": "string",
          "format": "date-time",
          "title": "created_at"
        },
        "package": {
          "description": "Package name",
          "type": "string",
          "title": "package-name"
        },
        "release": {
          "description": "Package release",
          "type": "string",
          "title": "package-release"
        }
      }
    },
    "PackageManifest": {
      "description": "manifests",
      "type": "array",
      "title": "manifests",
      "items": {
        "$ref": "#/definitions/Manifest"
      }
    },
    "Packages": {
      "description": "List packages, short view",
      "type": "array",
      "title": "Packages",
      "items": {
        "description": "test",
        "type": "object",
        "properties": {
          "created_at": {
            "description": "Package creation date",
            "type": "string",
            "format": "date-time",
            "title": "created_at"
          },
          "default": {
            "description": "Default/latest release version",
            "type": "string",
            "title": "default-release"
          },
          "name": {
            "description": "Package name",
            "type": "string",
            "title": "package-name"
          },
          "releases": {
            "description": "All available releases",
            "type": "array",
            "title": "available-releases",
            "items": {
              "description": "Release name",
              "type": "string",
              "title": "release-name"
            }
          },
          "visibility": {
            "description": "package visibility (public or private)",
            "type": "string",
            "title": "visibility"
          }
        }
      }
    },
    "PartDescriptor": {
      "description": "part descriptor",
      "type": "object",
      "title": "descriptor",
      "properties": {
        "dependencies": {
          "description": "dependencies",
          "type": "array",
          "title": "dependencies",
          "items": {
            "$ref": "#/definitions/Dependency"
          }
        },
        "digest": {
          "description": "content digest",
          "type": "string",
          "title": "digest"
        },
        "size": {
          "description": "blob size",
          "type": "integer",
          "format": "int64",
          "title": "content-size"
        }
      }
    },
    "PostPackage": {
      "description": "Package object",
      "type": "object",
      "title": "Package",
      "properties": {
        "blob": {
          "description": "Package blob: a tar.gz that is b64",
          "type": "string",
          "title": "blob"
        },
        "package": {
          "description": "Package name",
          "type": "string",
          "title": "package-name"
        },
        "release": {
          "description": "Package version",
          "type": "string",
          "title": "package-version"
        }
      }
    },
    "PullJson": {
      "description": "Package content",
      "type": "object",
      "title": "PackageContent",
      "properties": {
        "blob": {
          "description": "Package blob: a tar.gz in b64-encoded",
          "type": "string",
          "title": "blob"
        },
        "filename": {
          "description": "suggested filename",
          "type": "string",
          "title": "filename"
        },
        "package": {
          "description": "Package name",
          "type": "string",
          "title": "package-name"
        },
        "release": {
          "description": "Package version",
          "type": "string",
          "title": "package-version"
        }
      }
    },
    "Version": {
      "type": "object",
      "properties": {
        "ksonnet-registry-api": {
          "type": "string"
        }
      }
    }
  }
}`))
}
