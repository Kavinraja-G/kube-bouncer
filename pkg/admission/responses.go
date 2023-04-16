package admission

import (
	admissionv1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DenyRequest returns an admission review response that denies the request with the specified error message
func DenyRequest(errMsg string, admissionReviewRequest admissionv1.AdmissionReview) admissionv1.AdmissionReview {
	admissionReviewResponse := admissionv1.AdmissionReview{
		TypeMeta: admissionReviewRequest.TypeMeta,
		Response: &admissionv1.AdmissionResponse{
			UID: admissionReviewRequest.Request.UID,
			Result: &metav1.Status{
				Message: errMsg,
			},
			Allowed: false,
		},
	}
	return admissionReviewResponse
}

// AllowRequest returns an admission review response that allows the request
func AllowRequest(admissionReviewRequest admissionv1.AdmissionReview) admissionv1.AdmissionReview {
	admissionReviewResponse := admissionv1.AdmissionReview{
		TypeMeta: admissionReviewRequest.TypeMeta,
		Response: &admissionv1.AdmissionResponse{
			UID:     admissionReviewRequest.Request.UID,
			Allowed: true,
		},
	}
	return admissionReviewResponse
}
