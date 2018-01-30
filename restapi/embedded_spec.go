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
          },
          {
            "type": "string",
            "description": "Filter by media-type",
            "name": "media_type",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Packages"
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
        "summary": "List all manifests for a package",
        "operationId": "showPackageManifests",
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
    "/api/v1/packages/{namespace}/{package}/blobs/sha256/{digest}/json": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "blobs"
        ],
        "summary": "Pull a package blob by digest",
        "operationId": "pullBlobJson",
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
          },
          {
            "type": "string",
            "default": "gzip",
            "description": "return format type(json or gzip)",
            "name": "format",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/PullJson"
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
    "/api/v1/packages/{namespace}/{package}/channels": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "channel"
        ],
        "summary": "List channels",
        "operationId": "listChannels",
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
              "type": "array",
              "title": "channels",
              "items": {
                "$ref": "#/definitions/Channel"
              }
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
          "channel"
        ],
        "summary": "Create a new channel",
        "operationId": "createChannel",
        "parameters": [
          {
            "type": "string",
            "description": "Channel name",
            "name": "name",
            "in": "query",
            "required": true
          },
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
              "$ref": "#/definitions/Channel"
            }
          },
          "401": {
            "description": "Not authorized to create the channel",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Package not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "409": {
            "description": "Channel already exists",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/v1/packages/{namespace}/{package}/channels/{channel}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "channel"
        ],
        "summary": "show channel",
        "operationId": "showChannel",
        "parameters": [
          {
            "type": "string",
            "description": "channel name",
            "name": "channel",
            "in": "path",
            "required": true
          },
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
              "type": "array",
              "title": "channels",
              "items": {
                "$ref": "#/definitions/Channel"
              }
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
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "channel"
        ],
        "summary": "Delete channel",
        "operationId": "deleteChannel",
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
            "description": "channel name",
            "name": "channel",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "full package name",
            "name": "package",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "title": "channels",
              "items": {
                "$ref": "#/definitions/Channel"
              }
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
    "/api/v1/packages/{namespace}/{package}/channels/{channel}/{release}": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "channel"
        ],
        "summary": "Add a release to a channel",
        "operationId": "createChannelRelease",
        "parameters": [
          {
            "type": "string",
            "description": "channel name",
            "name": "channel",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "namespace",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "full package name",
            "name": "package",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Release name",
            "name": "release",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Channel"
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
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "channel"
        ],
        "summary": "Remove a release from the channel",
        "operationId": "deleteChannelRelease",
        "parameters": [
          {
            "type": "string",
            "description": "channel name",
            "name": "channel",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "namespace",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "full package name",
            "name": "package",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Release name",
            "name": "release",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "title": "channels",
              "items": {
                "$ref": "#/definitions/Channel"
              }
            }
          },
          "401": {
            "description": "Not authorized to read the package",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Resource not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/api/v1/packages/{namespace}/{package}/{release}/{media_type}": {
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
        "summary": "Show a package",
        "operationId": "showPackage",
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
          },
          {
            "type": "string",
            "description": "content type",
            "name": "media_type",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Package"
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
        "operationId": "deletePackage",
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
          },
          {
            "type": "string",
            "description": "content type",
            "name": "media_type",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Package"
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
    "/api/v1/packages/{namespace}/{package}/{release}/{media_type}/pull": {
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
          },
          {
            "type": "string",
            "description": "content type",
            "name": "media_type",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "reponse format: json or blob",
            "name": "format",
            "in": "query"
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
    "/api/v1/packages/{namespace}/{package}/{release}/{media_type}/pull/json": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "package",
          "blobs"
        ],
        "summary": "Download the package",
        "operationId": "pullPackageJson",
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
          },
          {
            "type": "string",
            "description": "content type",
            "name": "media_type",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "reponse format: json or blob",
            "name": "format",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/PullJson"
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
    "Channel": {
      "description": "Organize releases into channel, eg: dev/beta/stable",
      "type": "object",
      "title": "Channel",
      "properties": {
        "current": {
          "description": "Current/latest release in the channel. The channel returns this release by default",
          "type": "string",
          "title": "Latest release"
        },
        "name": {
          "description": "Channel name",
          "type": "string",
          "title": "Channel name"
        },
        "releases": {
          "description": "All availables releases in the channel",
          "type": "array",
          "title": "Releases",
          "items": {
            "description": "Release name",
            "type": "string",
            "title": "Release name"
          }
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
          "$ref": "#/definitions/OciDescriptor"
        },
        "created_at": {
          "description": "creation data",
          "type": "string",
          "format": "date-time",
          "title": "created-at"
        },
        "mediaType": {
          "description": "manifest-type",
          "type": "string",
          "title": "media-type"
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
    "OciDescriptor": {
      "description": "OCI descriptor",
      "type": "object",
      "title": "descriptor",
      "properties": {
        "digest": {
          "description": "content digest",
          "type": "string",
          "title": "digest"
        },
        "mediaType": {
          "description": "content type",
          "type": "string",
          "title": "media-type"
        },
        "size": {
          "description": "blob size",
          "type": "integer",
          "format": "int64",
          "title": "content-size"
        },
        "urls": {
          "description": "download mirrors",
          "type": "array",
          "title": "urls",
          "items": {
            "description": "url",
            "type": "string",
            "title": "url"
          }
        }
      }
    },
    "Package": {
      "description": "Package object",
      "type": "object",
      "title": "Package",
      "properties": {
        "channels": {
          "type": "array",
          "title": "channels",
          "items": {
            "description": "List channels for that release",
            "type": "string",
            "title": "channel-name"
          }
        },
        "content": {
          "$ref": "#/definitions/Manifest"
        },
        "created_at": {
          "description": "Package creation date",
          "type": "string",
          "format": "date-time",
          "title": "created_at"
        },
        "mediaType": {
          "description": "manifest-type",
          "type": "string",
          "title": "media-type"
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
            "title": "created_at"
          },
          "default": {
            "description": "Default/latest release version",
            "type": "string",
            "title": "default-release"
          },
          "manifests": {
            "description": "All formats",
            "type": "array",
            "title": "available-manifests",
            "items": {
              "description": "format name",
              "type": "string",
              "title": "format-name"
            }
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
    "PostPackage": {
      "description": "Package object",
      "type": "object",
      "title": "Package",
      "properties": {
        "blob": {
          "description": "Package blob: a tar.gz in b64-encoded",
          "type": "string",
          "title": "blob"
        },
        "media_type": {
          "description": "mediatype of the blob",
          "type": "string",
          "title": "package-type"
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
        "appr-api": {
          "type": "string"
        }
      }
    }
  }
}`))
}
