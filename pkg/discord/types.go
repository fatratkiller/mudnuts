package main

// TokenAttributes defines the attributes of a cryptocurrency token.
type TokenAttributes struct {
	GTScore float64 `json:"gt_score"`
	// Add other attributes here as per your requirements.
}

// Token represents a single cryptocurrency token.
type Token struct {
	Attributes TokenAttributes `json:"attributes"`
}

// TokenResponse represents the response structure from the API.
type TokenResponse struct {
	Data []Token `json:"data"`
}
