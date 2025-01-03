The changelog is automatically generated and it follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) format.

## [camunda-platform-9.4.5](https://github.com/jessesimpson36/camunda-platform-helm/releases/tag/camunda-platform-9.4.5) (2024-12-10)

### Fixes

- Renovate disable elasticsearch minor upgrades and revert elasticsearch upgrade (#2666)

<!-- generated by git-cliff -->
### Release Info

Supported versions:

- Camunda applications: [8.4](https://github.com/camunda/camunda-platform/releases?q=tag%3A8.4&expanded=true)
- Helm values: [9.4.5](https://artifacthub.io/packages/helm/camunda/camunda-platform/9.4.5#parameters)
- Helm CLI: [3.16.3](https://github.com/helm/helm/releases/tag/v3.16.3)

Camunda images:

- docker.io/camunda/connectors-bundle:8.4.15
- docker.io/camunda/identity:8.4.15
- docker.io/camunda/operate:8.4.15
- docker.io/camunda/optimize:8.4.12
- docker.io/camunda/tasklist:8.4.15
- docker.io/camunda/zeebe:8.4.14
- registry.camunda.cloud/web-modeler-ee/modeler-restapi:8.4.12
- registry.camunda.cloud/web-modeler-ee/modeler-webapp:8.4.12
- registry.camunda.cloud/web-modeler-ee/modeler-websockets:8.4.12

Non-Camunda images:

- docker.io/bitnami/elasticsearch:8.9.2
- docker.io/bitnami/keycloak:22.0.5
- docker.io/bitnami/os-shell:12-debian-12-r16
- docker.io/bitnami/postgresql:14.5.0-debian-11-r35
- docker.io/bitnami/postgresql:15.10.0-debian-12-r2

### Verification

To verify the integrity of the Helm chart using [Cosign](https://docs.sigstore.dev/signing/quickstart/):

```shell
cosign verify-blob camunda-platform-9.4.5.tgz \
  --bundle camunda-platform-9.4.5.cosign.bundle \
  --certificate-oidc-issuer "https://token.actions.githubusercontent.com" \
  --certificate-identity "https://github.com/jessesimpson36/camunda-platform-helm/.github/workflows/chart-release-chores.yml@refs/pull/2560/merge"
```
