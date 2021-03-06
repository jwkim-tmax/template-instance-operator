/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
type V1beta2DaemonSetUpdateStrategy struct {

	// Rolling update config params. Present only if type = \"RollingUpdate\".
	RollingUpdate *V1beta2RollingUpdateDaemonSet `json:"rollingUpdate,omitempty"`

	// Type of daemon set update. Can be \"RollingUpdate\" or \"OnDelete\". Default is RollingUpdate.
	Type_ string `json:"type,omitempty"`
}
