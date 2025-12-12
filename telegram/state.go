package telegram

import (
	"errors"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/IteratorInnovator/git-gram/config"
)

var stateSecret []byte = []byte(config.AppCfg.STATE_SECRET)

func generateStateToken(chat_id int64) (string, error) {
	chatId := strconv.FormatInt(chat_id, 10)

	mac := hmac.New(sha256.New, stateSecret)
	_, err := mac.Write([]byte(chatId))
	if err != nil {
		return "", err
	}

	signature := mac.Sum(nil)
	raw := append([]byte(chatId+":"), signature...)
    token := base64.RawURLEncoding.EncodeToString(raw)

	return token, nil
}

func parseAndVerifyStateToken(token string) (int64, error) {
	raw, err := base64.RawURLEncoding.DecodeString(token)
    if err != nil {
        return 0, errors.New("invalid state token")
    }

    parts := strings.SplitN(string(raw), ":", 2)
    if len(parts) != 2 {
        return 0, errors.New("invalid state token")
    }

    payload := parts[0]
    sigBytes := []byte(parts[1])

    mac := hmac.New(sha256.New, stateSecret)
    mac.Write([]byte(payload))
    expectedSig := mac.Sum(nil)

    if !hmac.Equal(sigBytes, expectedSig) {
        return 0, errors.New("invalid state token")
    }

    chatID, err := strconv.ParseInt(payload, 10, 64)
    if err != nil {
        return 0, errors.New("invalid state token")
    }

    return chatID, nil
}