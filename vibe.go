package ordinex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// VibeSorter implements Vibe Sort. It sends the slice to a Large Language Model
// and trusts that the model returns it in order. If the API call fails or the
// model returns something that cannot be parsed, the original slice is returned
// unsorted.
//
// Time: O($). Space: O(☁).
type VibeSorter struct {
	// APIKey is the OpenAI API key. If empty, the OPENAI_API_KEY environment
	// variable is used.
	APIKey string

	// Model is the model name to query. If empty, it defaults to "gpt-4o-mini".
	Model string
}

// Name returns the algorithm's name, "Vibe Sort".
func (VibeSorter) Name() string { return "Vibe Sort" }

// Sort returns input sorted by a Large Language Model. The input is not
// modified. If the request fails, the response cannot be parsed, or input has
// fewer than two elements, a copy of input is returned unchanged. The result is
// whatever the model produces and is not guaranteed to be sorted.
func (v VibeSorter) Sort(input []int) []int {
	out := copySlice(input)
	if len(out) <= 1 {
		return out
	}

	key := v.APIKey
	if key == "" {
		key = os.Getenv("OPENAI_API_KEY")
	}
	model := v.Model
	if model == "" {
		model = "gpt-4o-mini"
	}

	raw, _ := json.Marshal(input)
	prompt := fmt.Sprintf(
		"Sort this JSON integer array in ascending order. "+
			"Respond with only a valid JSON array, no explanation: %s", raw)

	body, _ := json.Marshal(map[string]any{
		"model": model,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	})

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions",
		bytes.NewReader(body))
	if err != nil {
		return out
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return out
	}
	defer resp.Body.Close()

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil || len(result.Choices) == 0 {
		return out
	}

	var sorted []int
	if err := json.Unmarshal([]byte(result.Choices[0].Message.Content), &sorted); err != nil {
		return out
	}

	return sorted
}
