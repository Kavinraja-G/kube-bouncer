package webhooks

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	admission "github.com/Kavinraja-G/kube-bouncer/pkg/admission"
	corev1 "k8s.io/api/core/v1"
)

// handler for podBouncer validation webhooks
func PodBouncer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("could not read the request body: %v", err)
	}

	admissionReviewRequest, err := admission.GetAdmissionReviewRequest(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	// allows the request by default
	admissionReviewResponse := admission.AllowRequest(admissionReviewRequest)

	// gets raw request body and tries to translate into the pod object
	var reqPod corev1.Pod
	rawPod := admissionReviewRequest.Request.Object.Raw
	if err := json.Unmarshal(rawPod, &reqPod); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("failed to unmarshall the requested pod object: %v", err)
	}

	// Check 1: deny the response if the pod doesn't have the readiness or liveness probe configured
	for _, container := range reqPod.Spec.Containers {
		if container.ReadinessProbe == nil || container.LivenessProbe == nil {
			errMsg := "Requested pod " + reqPod.Name + " does not contain readiness or liveness probes. Please add them and retry."
			admissionReviewResponse = admission.DenyRequest(errMsg, admissionReviewRequest)
		}
	}

	bytes, err := json.Marshal(&admissionReviewResponse)
	if err != nil {
		log.Printf("error while marshaling the admission review response: %v", err)
	}

	w.Write(bytes)
}
