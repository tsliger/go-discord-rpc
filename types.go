package main

type handshake struct {
	V        int64  `json:"v"`
	ClientID string `json:"client_id"`
}

type internalActivity struct {
	Cmd   string       `json:"cmd"`
	Args  internalArgs `json:"args"`
	Nonce string       `json:"nonce"`
}

type internalArgs struct {
	Pid      int          `json:"pid"`
	Activity ActivityData `json:"activity"`
	Nonce    string       `json:"nonce"`
}

type ActivityData struct {
	State      string            `json:"state,omitempty"`
	Type       int               `json:"type,omitempty"`
	Details    string            `json:"details,omitempty"`
	Timestamps ActivityTimestamp `json:"timestamps,omitempty"`
	Assets     ActivityAssets    `json:"assets,omitempty"`
}

type ActivityTimestamp struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type ActivityButtons struct {
}
