package main

import (
	"fmt"
	"os"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/mymmrac/telego"
)

const envPrefix = "TG_BOT_"

func main() {
	botToken := env("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Printf("] bot.Token %v: %v\n", bot.Token(), botToken)

	// Demo how to set up a webhook on Telegram side (done below)
	// _ = bot.SetWebhook(&telego.SetWebhookParams{
	// 	URL: env("WEBHOOK_BASE") + "/bot" + bot.Token(),
	// })

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo()
	fmt.Printf("Webhook Info: %+v\n", info)

	// Get an update channel from webhook, also all options are optional.
	// Note: For one bot, only one webhook allowed.
	updates, err := bot.UpdatesViaWebhook("/bot"+bot.Token(),
		// Set chan buffer (default 128)
		telego.WithWebhookBuffer(128),

		// Set fast http server that will be used to handle webhooks (default telego.FastHTTPWebhookServer)
		// Note: If SecretToken is non-empty, it will be verified on each request
		telego.WithWebhookServer(telego.FastHTTPWebhookServer{
			Logger: bot.Logger(),
			Server: &fasthttp.Server{},
			Router: router.New(),
			//SecretToken: "token",
		}),

		// Calls SetWebhook before starting webhook
		telego.WithWebhookSet(&telego.SetWebhookParams{
			URL: env("WEBHOOK_BASE") + "/bot" + bot.Token(),
		}),
	)
	assert(err == nil, "UpdatesViaWebhook error", err)

	// Start server for receiving requests from the Telegram
	go func() {
		fmt.Println("Starting Webhook server.")
		err = bot.StartWebhook("localhost:443")
		assert(err == nil, "StartWebhook error", err)
		fmt.Println("Webhook server stopped.")
	}()

	// Stop reviving updates from update channel and shutdown webhook server
	defer func() {
		err = bot.StopWebhook()
		assert(err == nil, "StopWebhook error", err)
	}()

	fmt.Println("Handling updates...")
	// Loop through all updates when they came
	for update := range updates {
		fmt.Printf("Update: %+v\n", update)
	}
}

func env(name string) string {
	value, ok := os.LookupEnv(envPrefix + name)
	assert(ok, "Environment variable "+envPrefix+name+" not found")
	//fmt.Printf("] %v: %v\n", name, value)
	return value
}

func assert(ok bool, args ...interface{}) {
	if !ok {
		fmt.Println(append([]interface{}{"FATAL:"}, args...)...)
		os.Exit(1)
	}
}
