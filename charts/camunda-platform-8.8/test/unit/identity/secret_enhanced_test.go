// Test golden file for identity firstuser secret with new format
func (s *secretTest) TestIdentityFirstUserSecretGoldenFile() {
	// given
	options := &helm.Options{
		SetValues: map[string]string{
			"identity.firstUser.enabled":         "true",
			"identity.firstUser.secret.password": "test-password-123",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}

	s.templates = []string{
		"templates/identity/secret-firstuser.yaml",
	}

	// when
	output := helm.RenderTemplate(s.T(), options, s.chartPath, s.release, s.templates)

	// then
	// This would typically be compared against a golden file
	s.Require().Contains(output, "camunda-platform-test-identity-firstuser")
	s.Require().Contains(output, "identity-firstuser-password")
	s.Require().Contains(output, "io.camunda.zeebe/warning")
}