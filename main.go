package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	admission "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
	"k8s.io/utils/strings/slices"
)

type WebhookParameters struct {
	certFile string
	keyFile  string
	port     int
}

var (
	denyNamespaces        = []string{"prod"}
	config                *rest.Config
	webhookParams         WebhookParameters
	universalDeserializer = serializer.NewCodecFactory(runtime.NewScheme()).UniversalDeserializer()
)

func HandleValidate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var admissionReviewRequest admission.AdmissionReview

	if _, _, err := universalDeserializer.Decode(body, nil, &admissionReviewRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("could not deserialize request: %v", err)
	}

	reqNamespace := admissionReviewRequest.Request.Namespace
	admissionReviewResponse := admission.AdmissionReview{
		TypeMeta: admissionReviewRequest.TypeMeta,
		Response: &admission.AdmissionResponse{
			UID: admissionReviewRequest.Request.UID,
		},
	}
	log.Printf("Validating the requested namespace: %v is in the deny list %v", reqNamespace, denyNamespaces)

	// deny the response if the namespace exists in the denyNamespace list
	if slices.Contains(denyNamespaces, reqNamespace) {
		errMsg := fmt.Sprintf("requested namespace %s found in deny namepsace list %v", reqNamespace, denyNamespaces)
		admissionReviewResponse.Response.Allowed = false
		admissionReviewResponse.Response.Result = &metav1.Status{Message: errMsg}
	} else {
		admissionReviewResponse.Response.Allowed = true
	}

	bytes, err := json.Marshal(&admissionReviewResponse)
	if err != nil {
		log.Printf("error while marshaling response: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func main() {
	// setup flags for the webhook server
	flag.IntVar(&webhookParams.port, "port", 8443, "Server port for the Webhook")
	flag.StringVar(&webhookParams.certFile, "tlsCert", "/etc/nsbouncer/certs/tls.crt", "File containing the certficate required for HTTPS communication")
	flag.StringVar(&webhookParams.keyFile, "tlsKey", "/etc/nsbouncer/certs/tls.key", "Private key file for the TLS certificate")

	// validate route
	http.HandleFunc("/validate", HandleValidate)

	// start the server with TLS
	log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(webhookParams.port), webhookParams.certFile, webhookParams.keyFile, nil))
}
