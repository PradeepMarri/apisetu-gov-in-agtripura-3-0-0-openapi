package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// AcademicCertificateSchema represents the AcademicCertificateSchema schema from the OpenAPI specification
type AcademicCertificateSchema struct {
	Language string `json:"language"`
	Number int `json:"number"`
	Validfromdate string `json:"validFromDate"`
	Certificatedata map[string]interface{} `json:"CertificateData"`
	Issuedat string `json:"issuedAt"`
	Issuedate string `json:"issueDate"`
	Name string `json:"name"`
	Issuedto map[string]interface{} `json:"IssuedTo"`
	Status string `json:"status"`
	TypeField string `json:"type"`
	Issuedby map[string]interface{} `json:"IssuedBy"`
}

// ConsentArtifactSchema represents the ConsentArtifactSchema schema from the OpenAPI specification
type ConsentArtifactSchema struct {
	Consent map[string]interface{} `json:"consent"`
	Signature map[string]interface{} `json:"signature"`
}
