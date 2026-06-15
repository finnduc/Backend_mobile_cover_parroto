package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"go-cover-parroto/internal/configs"
)

type azurePronunciationAssessment struct {
	PronScore         float64 `json:"PronScore"`
	AccuracyScore     float64 `json:"AccuracyScore"`
	FluencyScore      float64 `json:"FluencyScore"`
	CompletenessScore float64 `json:"CompletenessScore"`
}

type pronunciationPhonemeResult struct {
	Phoneme                 string  `json:"Phoneme"`
	PronunciationAssessment struct {
		AccuracyScore float64 `json:"AccuracyScore"`
	} `json:"PronunciationAssessment"`
}

type pronunciationWordResult struct {
	Word                 string `json:"Word"`
	PronunciationAssessment struct {
		AccuracyScore float64 `json:"AccuracyScore"`
		ErrorType     string  `json:"ErrorType"`
	} `json:"PronunciationAssessment"`
	Phonemes []pronunciationPhonemeResult `json:"Phonemes"`
}

type pronunciationNBest struct {
	Lexical                 string                        `json:"Lexical"`
	ITN                     string                        `json:"ITN"`
	MaskedITN               string                        `json:"MaskedITN"`
	Display                 string                        `json:"Display"`
	PronunciationAssessment azurePronunciationAssessment  `json:"PronunciationAssessment"`
	Words                   []pronunciationWordResult      `json:"Words"`
}

type azureAssessmentResult struct {
	RecognitionStatus string                  `json:"RecognitionStatus"`
	NBest             []pronunciationNBest    `json:"NBest"`
}

type PronunciationResult struct {
	Text         string              `json:"text"`
	OverallScore float64             `json:"overallScore"`
	Scores       PronunciationScores `json:"scores"`
	Feedback     string              `json:"feedback"`
	Words        []WordResult        `json:"words"`
}

type PronunciationScores struct {
	Accuracy     float64 `json:"accuracy"`
	Fluency      float64 `json:"fluency"`
	Completeness float64 `json:"completeness"`
	Prosody      float64 `json:"prosody"`
}

type WordResult struct {
	Word         string          `json:"word"`
	Score        float64         `json:"score"`
	Feedback     string          `json:"feedback"`
	WeakPhonemes []PhonemeResult `json:"weakPhonemes"`
}

type PhonemeResult struct {
	Phoneme string  `json:"phoneme"`
	Score   float64 `json:"score"`
}

func assessPronunciation(audioData []byte, referenceText string) (*PronunciationResult, error) {
	cfg := configs.LoadAzureConfig()

	baseURL := fmt.Sprintf("https://%s.stt.speech.microsoft.com/speech/recognition/conversation/cognitiveservices/v1", cfg.SpeechRegion)

	assessmentConfig := map[string]interface{}{
		"ReferenceText":   referenceText,
		"GradingSystem":   "HundredMark",
		"Granularity":     "Phoneme",
		"Dimension":       "Comprehensive",
	}
	configJSON, _ := json.Marshal(assessmentConfig)
	encodedConfig := base64.StdEncoding.EncodeToString(configJSON)

	params := fmt.Sprintf("language=en-US&format=detailed&pronunciationAssessment=%s", encodedConfig)
	fullURL := fmt.Sprintf("%s?%s", baseURL, params)

	req, err := http.NewRequest("POST", fullURL, bytes.NewReader(audioData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Ocp-Apim-Subscription-Key", cfg.SpeechKey)
	req.Header.Set("Content-Type", "audio/wav")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("azure request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("azure returned status %d: %s", resp.StatusCode, string(body))
	}

	var azureResult azureAssessmentResult
	if err := json.NewDecoder(resp.Body).Decode(&azureResult); err != nil {
		return nil, fmt.Errorf("failed to decode azure response: %w", err)
	}

	if len(azureResult.NBest) == 0 {
		return nil, fmt.Errorf("no pronunciation results")
	}

	best := azureResult.NBest[0]
	pron := best.PronunciationAssessment

	prosodyScore := 0.0
	if len(best.Words) > 0 {
		prosodyScore = pron.PronScore
	}

	result := &PronunciationResult{
		Text:         best.Display,
		OverallScore: pron.PronScore,
		Scores: PronunciationScores{
			Accuracy:     pron.AccuracyScore,
			Fluency:      pron.FluencyScore,
			Completeness: pron.CompletenessScore,
			Prosody:      prosodyScore,
		},
		Words: []WordResult{},
	}

	result.Feedback = formatVietnameseFeedback(result.OverallScore)

	for _, w := range best.Words {
		wordPron := w.PronunciationAssessment
		wordResult := WordResult{
			Word:         w.Word,
			Score:        wordPron.AccuracyScore,
			Feedback:     formatWordFeedback(wordPron.AccuracyScore),
			WeakPhonemes: []PhonemeResult{},
		}
		for _, p := range w.Phonemes {
			if p.PronunciationAssessment.AccuracyScore < 80 {
				wordResult.WeakPhonemes = append(wordResult.WeakPhonemes, PhonemeResult{
					Phoneme: p.Phoneme,
					Score:   p.PronunciationAssessment.AccuracyScore,
				})
			}
		}
		result.Words = append(result.Words, wordResult)
	}

	return result, nil
}

func formatVietnameseFeedback(score float64) string {
	if score >= 80 {
		return "Bạn phát âm khá tốt. Ngữ điệu cần hơi đều, cần nhấn nhá tự nhiên hơn."
	}
	if score >= 60 {
		return "Bạn phát âm tạm được. Hãy luyện tập thêm để cải thiện."
	}
	return "Bạn cần luyện tập thêm. Hãy chú ý phát âm từng từ rõ ràng hơn."
}

func formatWordFeedback(score float64) string {
	if score >= 80 {
		return "Tốt"
	}
	if score >= 60 {
		return "Cần cải thiện"
	}
	return "Cần luyện tập thêm"
}

func encodeReferenceText(text string) string {
	return strings.ReplaceAll(text, " ", "+")
}
