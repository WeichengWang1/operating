/*
Copyright 2023 The KusionStack Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"strings"

	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var (
	// HandlerMap contains all admission webhook handlers.
	HandlerMap = map[string]admission.Handler{}
)

// Add adds itself to the manager
func Add(mgr manager.Manager) error {
	server := mgr.GetWebhookServer()

	// register admission handlers
	for name, handler := range HandlerMap {
		if len(name) == 0 {
			klog.Warningf("Skip no-name handler.")
			continue
		}

		path := name
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		server.Register(path, &webhook.Admission{Handler: handler})
		klog.V(3).Infof("Registered webhook handler %s", path)
	}

	return nil
}