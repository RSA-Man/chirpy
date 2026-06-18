package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Cleaned_body string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	const redact = "****"
	badWords := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Split(params.Body, " ")
	for i, word := range words {
		lowerCaseWord := strings.ToLower(word)
		for _, badWord := range badWords {
			if lowerCaseWord == badWord {
				words[i] = redact
			}
		}
	}
	redactedBody := strings.Join(words, " ")

	respondWithJSON(w, http.StatusOK, returnVals{
		Cleaned_body: redactedBody,
	})
}
