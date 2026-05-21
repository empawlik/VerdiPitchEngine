package ai

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"cloud.google.com/go/vertexai/genai"
)

func TestNewEnterpriseClient_RequiredParameters(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{
			name: "Missing ProjectID",
			cfg: Config{
				ProjectID:       "",
				CredentialsPath: "dummy.json",
			},
			wantErr: true,
		},
		{
			name: "Missing CredentialsPath",
			cfg: Config{
				ProjectID:       "dummy-project",
				CredentialsPath: "",
			},
			wantErr: true,
		},
		{
			name: "Non-existent CredentialsPath",
			cfg: Config{
				ProjectID:       "dummy-project",
				CredentialsPath: "does-not-exist-at-all.json",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewEnterpriseClient(ctx, tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEnterpriseClient() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && client == nil {
				t.Error("Expected non-nil client when error is nil")
			}
		})
	}
}

func TestNewEnterpriseClient_Success(t *testing.T) {
	ctx := context.Background()

	// Generate a small RSA private key for testing
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}
	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		t.Fatalf("Failed to marshal PKCS8: %v", err)
	}
	pemBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privBytes,
	}
	pemString := string(pem.EncodeToMemory(pemBlock))

	// Construct fully valid Service Account JSON
	credMap := map[string]interface{}{
		"type":           "service_account",
		"project_id":     "dummy-project",
		"private_key_id": "dummy-key-id",
		"private_key":    pemString,
		"client_email":   "dummy@dummy-project.iam.gserviceaccount.com",
	}

	dummySA, err := json.Marshal(credMap)
	if err != nil {
		t.Fatalf("Failed to marshal credentials map: %v", err)
	}

	tempDir := t.TempDir()
	credPath := filepath.Join(tempDir, "cred.json")
	if err := os.WriteFile(credPath, dummySA, 0600); err != nil {
		t.Fatalf("Failed to write dummy SA file: %v", err)
	}

	cfg := Config{
		ProjectID:       "dummy-project",
		CredentialsPath: credPath,
	}

	client, err := NewEnterpriseClient(ctx, cfg)
	if err != nil {
		t.Fatalf("Expected client to initialize successfully with valid dummy credentials, got: %v", err)
	}
	if client == nil {
		t.Fatal("Expected client not to be nil")
	}
	defer func() { _ = client.Close() }()
}

func TestRunAgent_FailsUnderTesting(t *testing.T) {
	ctx := context.Background()
	// RunAgent has hardcoded paths, so it is expected to fail in local unit tests.
	err := RunAgent(ctx)
	if err == nil {
		t.Error("Expected RunAgent to fail due to missing hardcoded credentials file, but it succeeded")
	}
}

type mockClient struct {
	mockModel generativeModel
	closeErr  error
}

func (m *mockClient) GenerativeModel(name string) generativeModel {
	return m.mockModel
}

func (m *mockClient) Close() error {
	return m.closeErr
}

type mockGenerativeModel struct {
	temp         float32
	generateResp *genai.GenerateContentResponse
	generateErr  error
}

func (m *mockGenerativeModel) SetTemperature(temp float32) {
	m.temp = temp
}

func (m *mockGenerativeModel) GenerateContent(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
	return m.generateResp, m.generateErr
}

func TestRunAgent_Scenarios(t *testing.T) {
	// Save original clientCreator and restore it
	origCreator := clientCreator
	defer func() {
		clientCreator = origCreator
	}()

	ctx := context.Background()

	t.Run("Client creation error", func(t *testing.T) {
		clientCreator = func(ctx context.Context, cfg Config) (client, error) {
			return nil, fmt.Errorf("mock client creation error")
		}
		err := RunAgent(ctx)
		if err == nil || err.Error() != "mock client creation error" {
			t.Errorf("expected mock client creation error, got %v", err)
		}
	})

	t.Run("GenerateContent error", func(t *testing.T) {
		mockM := &mockGenerativeModel{
			generateErr: fmt.Errorf("mock generate error"),
		}
		mockC := &mockClient{
			mockModel: mockM,
		}
		clientCreator = func(ctx context.Context, cfg Config) (client, error) {
			return mockC, nil
		}
		err := RunAgent(ctx)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		expectedErr := "generation failure: mock generate error"
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err.Error())
		}
	})

	t.Run("Empty candidates response", func(t *testing.T) {
		mockM := &mockGenerativeModel{
			generateResp: &genai.GenerateContentResponse{
				Candidates: []*genai.Candidate{},
			},
		}
		mockC := &mockClient{
			mockModel: mockM,
		}
		clientCreator = func(ctx context.Context, cfg Config) (client, error) {
			return mockC, nil
		}
		err := RunAgent(ctx)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		expectedErr := "empty model response candidates"
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err.Error())
		}
	})

	t.Run("Nil candidate content response", func(t *testing.T) {
		mockM := &mockGenerativeModel{
			generateResp: &genai.GenerateContentResponse{
				Candidates: []*genai.Candidate{
					{
						Content: nil,
					},
				},
			},
		}
		mockC := &mockClient{
			mockModel: mockM,
		}
		clientCreator = func(ctx context.Context, cfg Config) (client, error) {
			return mockC, nil
		}
		err := RunAgent(ctx)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		expectedErr := "empty model response candidates"
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err.Error())
		}
	})

	t.Run("Empty candidate parts response", func(t *testing.T) {
		mockM := &mockGenerativeModel{
			generateResp: &genai.GenerateContentResponse{
				Candidates: []*genai.Candidate{
					{
						Content: &genai.Content{
							Parts: []genai.Part{},
						},
					},
				},
			},
		}
		mockC := &mockClient{
			mockModel: mockM,
		}
		clientCreator = func(ctx context.Context, cfg Config) (client, error) {
			return mockC, nil
		}
		err := RunAgent(ctx)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		expectedErr := "empty model response candidates"
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err.Error())
		}
	})

	t.Run("Success response", func(t *testing.T) {
		mockM := &mockGenerativeModel{
			generateResp: &genai.GenerateContentResponse{
				Candidates: []*genai.Candidate{
					{
						Content: &genai.Content{
							Parts: []genai.Part{
								genai.Text("Confirm telemetry connection."),
							},
						},
					},
				},
			},
		}
		mockC := &mockClient{
			mockModel: mockM,
		}
		clientCreator = func(ctx context.Context, cfg Config) (client, error) {
			return mockC, nil
		}
		err := RunAgent(ctx)
		if err != nil {
			t.Fatalf("expected success, got error: %v", err)
		}
		if mockM.temp != 0.2 {
			t.Errorf("expected temperature 0.2, got %f", mockM.temp)
		}
	})
}
