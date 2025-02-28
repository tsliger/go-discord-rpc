package discordrpc

const PLAYING_TYPE = 0
const STREAMING_TYPE = 1
const LISTENTING_TYPE = 2
const WATCHING_TYPE = 3
const CUSTOM_TYPE = 4
const COMPETING_TYPE = 5

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
	Pid      int           `json:"pid"`
	Activity *ActivityData `json:"activity"`
	Nonce    string        `json:"nonce"`
}

type ActivityData struct {
	Type          int               `json:"type,omitempty"`
	Url           string            `json:"url,omitempty"`
	Timestamps    ActivityTimestamp `json:"timestamps,omitempty"`
	ApplicationId string            `json:"application_id,omitempty"`
	Details       string            `json:"details,omitempty"`
	State         string            `json:"state,omitempty"`
	Emoji         ActivityEmoji     `json:"emoji,omitempty"`
	Party         ActivityParty     `json:"party,omitempty"`
	Assets        ActivityAssets    `json:"assets,omitempty"`
	Secrets       ActivitySecrets   `json:"secrets,omitempty"`
	Instance      bool              `json:"instance,omitempty"`
	Buttons       []ActivityButton  `json:"buttons,omitempty"`
}

type ActivityTimestamp struct {
	Start int64 `json:"start,omitempty"`
	End   int64 `json:"end,omitempty"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type ActivityButton struct {
	Label string `json:"label,omitempty"`
	Url   string `json:"url,omitempty"`
}

type ActivityEmoji struct {
	Name     string `json:"name,omitempty"`
	Id       string `json:"id,omitempty"`
	Animated bool   `json:"animated,omitempty"`
}

type ActivityParty struct {
	Id   string `json:"id,omitempty"`
	Size []int  `json:"size,omitempty"`
}

type ActivitySecrets struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}
