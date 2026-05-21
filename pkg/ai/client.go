package ai

import (
	"context"
	"fmt"

	"cloud.google.com/go/vertexai/genai"
	"google.golang.org/api/option"
)

// +agentlint:enforce

// Config dictates the strict environmental requirements for the AI gateway.
type Config struct {
	ProjectID       string
	Location        string
	CredentialsPath string
}

// NewEnterpriseClient initializes a Vertex AI client bound to your high-assurance Service Account.
func NewEnterpriseClient(ctx context.Context, cfg Config) (*genai.Client, error) {
	if cfg.ProjectID == "" || cfg.CredentialsPath == "" {
		return nil, fmt.Errorf("initialization aborted: ProjectID and CredentialsPath are strictly required")
	}

	// Default routing to us-central1 if unspecified
	if cfg.Location == "" {
		cfg.Location = "us-central1"
	}

	// option.WithAuthCredentialsFile deterministically locks authentication to your downloaded JSON.
	// This prevents the system from accidentally using a developer's personal gcloud CLI login.
	client, err := genai.NewClient(
		ctx,
		cfg.ProjectID,
		cfg.Location,
		option.WithAuthCredentialsFile(option.ServiceAccount, cfg.CredentialsPath),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to provision Vertex AI client: %w", err)
	}

	return client, nil
}

// client abstracts the SDK client behavior for testing.
type client interface {
	GenerativeModel(string) generativeModel
	Close() error
}

// generativeModel abstracts the SDK's generative model behavior for testing.
type generativeModel interface {
	SetTemperature(float32)
	GenerateContent(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error)
}

// defaultClient wraps the concrete SDK client to implement the client interface.
type defaultClient struct {
	client *genai.Client
}

// GenerativeModel provisions the generative model wrapper.
func (c *defaultClient) GenerativeModel(name string) generativeModel {
	return &defaultGenerativeModel{model: c.client.GenerativeModel(name)}
}

// Close closes the underlying client connection.
func (c *defaultClient) Close() error {
	return c.client.Close()
}

// defaultGenerativeModel wraps the concrete SDK model to implement the generativeModel interface.
type defaultGenerativeModel struct {
	model *genai.GenerativeModel
}

// SetTemperature updates the model's temperature settings.
func (m *defaultGenerativeModel) SetTemperature(temp float32) {
	m.model.SetTemperature(temp)
}

// GenerateContent requests content generation from the model wrapper.
func (m *defaultGenerativeModel) GenerateContent(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
	return m.model.GenerateContent(ctx, parts...)
}

// clientCreator is an internal hook overridden in tests to mock the Vertex AI client.
var clientCreator = func(ctx context.Context, cfg Config) (client, error) {
	c, err := NewEnterpriseClient(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &defaultClient{client: c}, nil
}

// RunAgent boots the agent generative pipeline for telemetry checking.
func RunAgent(ctx context.Context) error {
	cfg := Config{
		ProjectID:       "antigravity-core-497020",
		Location:        "us-central1", // Standard enterprise region
		CredentialsPath: "/absolute/path/to/antigravity-core-497020-xxxxxx.json",
	}

	client, err := clientCreator(ctx, cfg)
	if err != nil {
		return err
	}
	// Per 900-GO, defer cleanup to prevent memory leaks on the underlying gRPC connection
	defer func() { _ = client.Close() }()

	// Provision the model (Gemini 1.5 Pro is standard for complex routing/reasoning)
	model := client.GenerativeModel("gemini-1.5-pro-001")

	// Enforce deterministic temperature for high-assurance logic if needed
	model.SetTemperature(0.2)

	resp, err := model.GenerateContent(ctx, genai.Text("Confirm telemetry connection."))
	if err != nil {
		return fmt.Errorf("generation failure: %w", err)
	}

	if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil && len(resp.Candidates[0].Content.Parts) > 0 {
		fmt.Println(resp.Candidates[0].Content.Parts[0])
	} else {
		return fmt.Errorf("empty model response candidates")
	}
	return nil
}
