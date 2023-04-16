package webhooks

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	admission "github.com/Kavinraja-G/kube-bouncer/pkg/admission"
	utils "github.com/Kavinraja-G/kube-bouncer/pkg/utils"
	"k8s.io/utils/strings/slices"
)

var (
	denyNamespaces = strings.Split(utils.GetEnv("DENY_NAMESPACES", ""), ",")
)

// handler for namespaceBouncer validation webhooks
func NamespaceBouncer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("could not read the request body: %v", err)
	}

	admissionReviewRequest, err := admission.GetAdmissionReviewRequest(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	reqNamespace := admissionReviewRequest.Request.Namespace
	admissionReviewResponse := admission.AllowRequest(admissionReviewRequest)

	// deny the response if the namespace exists in the denyNamespace list
	log.Printf("Validating the requested namespace: %v", reqNamespace)
	if slices.Contains(denyNamespaces, reqNamespace) {
		errMsg := "requested namespace " + reqNamespace + " found in deny namepsace list " + strings.Join(denyNamespaces, ",")
		admissionReviewResponse = admission.DenyRequest(errMsg, admissionReviewRequest)
	}

	bytes, err := json.Marshal(&admissionReviewResponse)
	if err != nil {
		log.Printf("error while marshaling the admission review response: %v", err)
	}

	w.Write(bytes)
}
