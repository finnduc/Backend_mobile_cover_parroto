package services

import (
	"context"
	"fmt"

	deepgram "github.com/deepgram/deepgram-go-sdk/v3/pkg/api/listen/v1/rest"
	interfaces "github.com/deepgram/deepgram-go-sdk/v3/pkg/client/interfaces/v1"
)

type ITranscriptionService interface {
	Transcribe(ctx context.Context, filePath string) (string, error)
}

type transcriptionService struct {
	dgClient *deepgram.Client
}

func NewTranscriptionService(dgClient *deepgram.Client) ITranscriptionService {
	return &transcriptionService{dgClient: dgClient}
}

func (s *transcriptionService) Transcribe(ctx context.Context, filePath string) (string, error) {
	options := &interfaces.PreRecordedTranscriptionOptions{
		Model:       "nova-3",
		Language:    "en-US",
		Punctuate:   true,
		SmartFormat: true,
	}

	res, err := s.dgClient.FromFile(ctx, filePath, options)
	if err != nil {
		return "", fmt.Errorf("deepgram transcription failed: %w", err)
	}

	if len(res.Results.Channels) == 0 || len(res.Results.Channels[0].Alternatives) == 0 {
		return "", fmt.Errorf("no transcription results")
	}

	transcript := res.Results.Channels[0].Alternatives[0].Transcript
	return transcript, nil
}
