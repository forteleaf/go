package main

import (
	"flag"
	"os"

	"github.com/JSila/packtFree"
	"github.com/JSila/packtFree/log"
	"github.com/JSila/packtFree/notify"
	"github.com/JSila/packtFree/notify/email"
	"github.com/JSila/packtFree/notify/email/options"
	"github.com/JSila/packtFree/notify/stdout"
)

func main() {
	var notifier = flag.String("notifier", "stdout", "How to notify user(s) - via email or stdout?")
	var cwd string
	flag.StringVar(&cwd, "cwd", "./", "Directory where config files and email template are")

	flag.Parse()

	os.Chdir(cwd)

	if err := packtFree.Get(); err != nil {
		log.Log.Panic(err)
	}

	var n notify.Notifier
	switch *notifier {
	case notify.Email:
		options, err := options.FromFile("config/mailgun.json")
		if err != nil {
			log.Log.Panic(err)
		}
		n, err = email.New(packtFree.Data, options)
		if err != nil {
			log.Log.Panic(err)
		}
	default:
		n = stdout.New(packtFree.Data)
	}

	users, err := notify.UsersFromFile("config/recipients.json")
	if err != nil {
		log.Log.Fatal(err)
	}
	for _, user := range users {
		if err := n.Notify(user); err != nil {
			log.Log.Error(err)
			log.Log.Fatal("Can't notify via ", n.Name())
		}
	}
}
