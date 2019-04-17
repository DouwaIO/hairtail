package webhook

import (
	"net/url"

	"github.com/DouwaIO/hairtail/src/task"
)

type Opts struct {
	URL         string
}

func New(opts Opts) (task.Task, error) {
	url, err := url.Parse(opts.URL)
	if err != nil {
		return nil, err
	}
	return &WebHook{
		URL:         opts.URL,
	}, nil
}

type WebHook struct {
	URL          string
}

func (g *WebHook) User(token string, t int) (error) {
	// client := NewClient(g.URL, token, g.SkipVerify)
	// login, err := client.CurrentUser(t)
	// if err != nil {
	// 	return nil, err
	// }

	// user := &model.User{}

	// if strings.HasPrefix(login.AvatarUrl, "http") {
	// 	user.Avatar = login.AvatarUrl
	// } else {
	// 	user.Avatar = g.URL + "/" + login.AvatarUrl
	// }

	return nil
}
