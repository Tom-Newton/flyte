/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type CoreArtifactBindingData struct {
	Index int64 `json:"index,omitempty"`
	PartitionKey string `json:"partition_key,omitempty"`
	Transform string `json:"transform,omitempty"`
}
