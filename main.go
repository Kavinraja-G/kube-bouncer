package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	webhooks "github.com/Kavinraja-G/kube-bouncer/pkg/webhooks"
)

type WebhookParameters struct {
	certFile string
	keyFile  string
	port     int
}

// Driver function to handle the webhooks
func main() {
	var webhookParams WebhookParameters

	// setup flags for the webhook server
	flag.IntVar(&webhookParams.port, "port", 8443, "Server port for the Webhook")
	flag.StringVar(&webhookParams.certFile, "tlsCert", "/etc/kubebouncer/certs/tls.crt", "File containing the certficate required for HTTPS communication")
	flag.StringVar(&webhookParams.keyFile, "tlsKey", "/etc/kubebouncer/certs/tls.key", "Private key file for the TLS certificate")

	// validate routes
	http.HandleFunc("/validate-namespace", webhooks.NamespaceBouncer)
	http.HandleFunc("/validate-pods", webhooks.PodBouncer)

	// start the server with TLS
	log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(webhookParams.port), webhookParams.certFile, webhookParams.keyFile, nil))
}
