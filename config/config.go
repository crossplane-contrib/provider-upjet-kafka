package config

import (
	"github.com/crossplane/upjet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("kafka_acl", func(r *config.Resource) {
		r.TerraformResource.Description = "The ACL resource allows a user to manage ACLs for Kafka."
		r.TerraformResource.Schema["acl_principal"].Description = "The principal to which the ACL applies."
		r.TerraformResource.Schema["acl_principal"].Description = "The principal to which the ACL applies."
		r.TerraformResource.Schema["acl_host"].Description = "The host from which the ACL principal will have access."
		r.TerraformResource.Schema["acl_operation"].Description = "The operation that is being allowed or denied."
		r.TerraformResource.Schema["acl_permission_type"].Description = "The type of permission. One of Unknown, Any, Allow, Deny."
		r.TerraformResource.Schema["resource_type"].Description = "The type of resource. One of Unknown, Any, Topic, Group, Cluster, TransactionalID."
		r.TerraformResource.Schema["resource_name"].Description = "The name of the resource to which the ACL applies."
		r.TerraformResource.Schema["resource_pattern_type_filter"].Description = "The pattern type of the resource. One of Any, Match, Literal, Prefixed."
	})
}
