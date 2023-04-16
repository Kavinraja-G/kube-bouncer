package admission

import (
	"log"

	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	admissionReviewRequest admissionv1.AdmissionReview
	universalDeserializer  = serializer.NewCodecFactory(runtime.NewScheme()).UniversalDeserializer()
)

// returns admissionReviewRequest object based on the input request body
func GetAdmissionReviewRequest(body []byte) (admissionv1.AdmissionReview, error) {
	_, _, err := universalDeserializer.Decode(body, nil, &admissionReviewRequest)
	if err != nil {
		log.Printf("could not deserialize request: %v", err)
		return admissionReviewRequest, err
	}

	return admissionReviewRequest, nil
}
