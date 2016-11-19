package config

func YourEmail() string {
	return "andrey.boar@gmail.com"
}

//These websites will be added to /etc/hosts file when focus command is run
func FavouriteWebSites() []string {
	w := make([]string, 0)

	w = append(w, "www.facebook.com")
	w = append(w, "news.ycombinator.com")

	return w
}
