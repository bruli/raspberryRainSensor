package rs

func New(serveURL string, cl HTTPClient) Handler {
	cli := client{
		cl:        cl,
		serverURL: serveURL,
	}
	return Handler{ReadRain: ReadRain(cli)}
}
