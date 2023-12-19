// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	acl "github.com/mbbush/provider-kafka-jet/internal/controller/kafka/acl"
	quota "github.com/mbbush/provider-kafka-jet/internal/controller/kafka/quota"
	topic "github.com/mbbush/provider-kafka-jet/internal/controller/kafka/topic"
	userscramcredential "github.com/mbbush/provider-kafka-jet/internal/controller/kafka/userscramcredential"
	providerconfig "github.com/mbbush/provider-kafka-jet/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acl.Setup,
		quota.Setup,
		topic.Setup,
		userscramcredential.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
