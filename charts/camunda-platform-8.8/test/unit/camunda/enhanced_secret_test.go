// Copyright 2024 Camunda Services GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package camunda

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	coreV1 "k8s.io/api/core/v1"
)

type enhancedSecretTest struct {
	suite.Suite
	chartPath string
	release   string
	namespace string
	templates []string
}

func TestEnhancedSecretTemplate(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../../../")
	require.NoError(t, err)

	suite.Run(t, &enhancedSecretTest{
		chartPath: chartPath,
		release:   "camunda-platform-test",
		namespace: "camunda-platform-" + strings.ToLower(random.UniqueId()),
		templates: []string{},
	})
}

// Test identity.firstUser new secret format
func (s *enhancedSecretTest) TestIdentityFirstUserNewSecretFormatPassword() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"identity.firstUser.enabled":       "true",
			"identity.firstUser.secret.password": "test-password-123",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	s.templates = []string{
		"templates/identity/secret-firstuser.yaml",
	}

	// when
	output := helm.RenderTemplate(s.T(), options, s.chartPath, s.release, s.templates)
	var secret coreV1.Secret
	helm.UnmarshalK8SYaml(s.T(), output, &secret)

	// then
	s.Require().Equal("camunda-platform-test-identity-firstuser", secret.ObjectMeta.Name)
	s.Require().Contains(secret.Annotations, "io.camunda.zeebe/warning")
	s.Require().Equal("test-password-123", string(secret.Data["identity-firstuser-password"]))
}

func (s *enhancedSecretTest) TestIdentityFirstUserNewSecretFormatExistingSecret() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"identity.firstUser.enabled":              "true",
			"identity.firstUser.secret.existingSecret": "my-existing-secret",
			"identity.firstUser.secret.existingSecretKey": "my-password-key",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	s.templates = []string{
		"templates/identity/secret-firstuser.yaml",
	}

	// when
	_, err := helm.RenderTemplateE(s.T(), options, s.chartPath, s.release, s.templates)

	// then - should not render secret when using existing secret
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "could not find template")
}

func (s *enhancedSecretTest) TestIdentityFirstUserFallbackToLegacyFormat() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"identity.firstUser.enabled":  "true",
			"identity.firstUser.password": "legacy-password",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	s.templates = []string{
		"templates/identity/secret-firstuser.yaml",
	}

	// when
	output := helm.RenderTemplate(s.T(), options, s.chartPath, s.release, s.templates)
	var secret coreV1.Secret
	helm.UnmarshalK8SYaml(s.T(), output, &secret)

	// then
	s.Require().Equal("camunda-platform-test-identity-firstuser", secret.ObjectMeta.Name)
	s.Require().Equal("legacy-password", string(secret.Data["identity-firstuser-password"]))
}

// Test OAuth client secrets new format
func (s *enhancedSecretTest) TestConsoleSecretNewFormat() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"global.identity.auth.enabled":       "true",
			"global.identity.auth.console.secret.password": "console-secret-123",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	s.templates = []string{
		"templates/camunda/secret-console.yaml",
	}

	// when
	output := helm.RenderTemplate(s.T(), options, s.chartPath, s.release, s.templates)
	var secret coreV1.Secret
	helm.UnmarshalK8SYaml(s.T(), output, &secret)

	// then
	s.Require().Equal("camunda-platform-test-console-identity-secret", secret.ObjectMeta.Name)
	s.Require().Contains(secret.Annotations, "io.camunda.zeebe/warning")
	s.Require().Equal("console-secret-123", string(secret.Data["identity-console-client-token"]))
}

func (s *enhancedSecretTest) TestConnectorsSecretNewFormat() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"global.identity.auth.enabled":          "true",
			"global.identity.auth.connectors.secret.password": "connectors-secret-123",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	s.templates = []string{
		"templates/camunda/secret-connectors.yaml",
	}

	// when
	output := helm.RenderTemplate(s.T(), options, s.chartPath, s.release, s.templates)
	var secret coreV1.Secret
	helm.UnmarshalK8SYaml(s.T(), output, &secret)

	// then
	s.Require().Equal("camunda-platform-test-connectors-identity-secret", secret.ObjectMeta.Name)
	s.Require().Contains(secret.Annotations, "io.camunda.zeebe/warning")
	s.Require().Equal("connectors-secret-123", string(secret.Data["identity-connectors-client-token"]))
}

// Test license secret new format
func (s *enhancedSecretTest) TestLicenseSecretNewFormat() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"global.license.secret.password": "test-license-key",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	s.templates = []string{
		"templates/camunda/secret-camunda-license.yaml",
	}

	// when
	output := helm.RenderTemplate(s.T(), options, s.chartPath, s.release, s.templates)
	var secret coreV1.Secret
	helm.UnmarshalK8SYaml(s.T(), output, &secret)

	// then
	s.Require().Equal("camunda-platform-test-license", secret.ObjectMeta.Name)
	s.Require().Contains(secret.Annotations, "io.camunda.zeebe/warning")
	s.Require().Equal("test-license-key", string(secret.Data["CAMUNDA_LICENSE_KEY"]))
}

func (s *enhancedSecretTest) TestLicenseSecretFallbackToLegacy() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"global.license.key": "legacy-license-key",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	s.templates = []string{
		"templates/camunda/secret-camunda-license.yaml",
	}

	// when
	output := helm.RenderTemplate(s.T(), options, s.chartPath, s.release, s.templates)
	var secret coreV1.Secret
	helm.UnmarshalK8SYaml(s.T(), output, &secret)

	// then
	s.Require().Equal("camunda-platform-test-license", secret.ObjectMeta.Name)
	s.Require().Equal("legacy-license-key", string(secret.Data["CAMUNDA_LICENSE_KEY"]))
}

func (s *enhancedSecretTest) TestLicenseSecretExistingSecret() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"global.license.secret.existingSecret": "my-license-secret",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	s.templates = []string{
		"templates/camunda/secret-camunda-license.yaml",
	}

	// when
	_, err := helm.RenderTemplateE(s.T(), options, s.chartPath, s.release, s.templates)

	// then - should not render secret when using existing secret
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "could not find template")
}

// Test precedence: new format takes priority over legacy
func (s *enhancedSecretTest) TestNewFormatTakesPrecedenceOverLegacy() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"global.license.key":            "legacy-license-key",
			"global.license.secret.password": "new-license-key",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	s.templates = []string{
		"templates/camunda/secret-camunda-license.yaml",
	}

	// when
	output := helm.RenderTemplate(s.T(), options, s.chartPath, s.release, s.templates)
	var secret coreV1.Secret
	helm.UnmarshalK8SYaml(s.T(), output, &secret)

	// then
	s.Require().Equal("new-license-key", string(secret.Data["CAMUNDA_LICENSE_KEY"]))
}

// Test helper functions
func (s *enhancedSecretTest) TestSecretHelperFunctions() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"global.license.secret.existingSecret":    "my-license-secret",
			"global.license.secret.existingSecretKey": "my-license-key",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	s.templates = []string{
		"templates/core/statefulset.yaml",
	}

	// when
	output := helm.RenderTemplate(s.T(), options, s.chartPath, s.release, s.templates)

	// then - verify that helper functions are working
	s.Require().Contains(output, "my-license-secret")
	s.Require().Contains(output, "my-license-key")
}