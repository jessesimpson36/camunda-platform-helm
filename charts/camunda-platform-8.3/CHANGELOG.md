# Changelog

## [8.3.25](https://github.com/camunda/camunda-platform-helm/compare/camunda-platform-8.3-8.3.24...camunda-platform-8.3-8.3.25) (2025-03-10)


### Bug Fixes

* **deps:** limit elasticsearch helm chart version for 8.3 and 8.4 ([#3063](https://github.com/camunda/camunda-platform-helm/issues/3063)) ([2f32232](https://github.com/camunda/camunda-platform-helm/commit/2f3223235466eed70dbbfe144c8ee7a31d527792))
* disable secret autoGenerated flag since it causes race condition ([#2906](https://github.com/camunda/camunda-platform-helm/issues/2906)) ([ddbccd9](https://github.com/camunda/camunda-platform-helm/commit/ddbccd9089c517ba12cf401e1f2617ffda55738e))
* ensure app configs rendered correctly in ConfigMap ([#3071](https://github.com/camunda/camunda-platform-helm/issues/3071)) ([36fcfe3](https://github.com/camunda/camunda-platform-helm/commit/36fcfe3d7eef93b4d613ca6891ac18161e3add37))
* remove unused test connection pod ([#3001](https://github.com/camunda/camunda-platform-helm/issues/3001)) ([9d2309a](https://github.com/camunda/camunda-platform-helm/commit/9d2309ab50c3bc1e3bb0fb2d0b7e6a27ed587200))


### Documentation

* update keywords of charts ([#3027](https://github.com/camunda/camunda-platform-helm/issues/3027)) ([7ce4275](https://github.com/camunda/camunda-platform-helm/commit/7ce4275968bb4ba4504a254ac4f02d2318be47d7))


### Dependencies

* update camunda-platform-8.3 (patch) ([#3055](https://github.com/camunda/camunda-platform-helm/issues/3055)) ([f7d7017](https://github.com/camunda/camunda-platform-helm/commit/f7d70171900e891a18f0df416f94e68498aa5cea))
* update camunda/tasklist docker tag to v8.3.23 ([#3096](https://github.com/camunda/camunda-platform-helm/issues/3096)) ([07420e8](https://github.com/camunda/camunda-platform-helm/commit/07420e8e44c74d5abdabe81fce116d8cd0d89d63))
* update elasticsearch docker tag to v19.21.2 ([#2885](https://github.com/camunda/camunda-platform-helm/issues/2885)) ([3fcadb0](https://github.com/camunda/camunda-platform-helm/commit/3fcadb0ab91b830e9442fd5cd5170de2da64f460))
* update module github.com/gruntwork-io/terratest to v0.48.2 ([#2664](https://github.com/camunda/camunda-platform-helm/issues/2664)) ([6ceef68](https://github.com/camunda/camunda-platform-helm/commit/6ceef685236ac41506ff3ce742759b1d3cbfde36))
