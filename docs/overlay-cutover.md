# Form3 overlay cutover (upstream v2.8.3)

This branch (`feature/upstream-v2.8.3-overlay`) is a clean overlay on upstream tag `v2.8.3`.
Form3 `master` in the sibling clone remains the read-only patch source; do not merge it here.

## Release tags

- Overlay releases use `v2.8.3-f3-<shortsha>` (short git SHA of the overlay commit).
- Multi-arch images are published by `.github/workflows/f3_upload_image.yml` on tag push.

## Helm upgrade

1. Inventory live Form3 CRDs in the cluster (`kubectl get crd | grep chaos-mesh.org`).
2. Upgrade the chart from this branch or from published images tagged `v2.8.3-f3-*`.
3. Compare `helm/chaos-mesh/values.yaml` with your current Form3 values (topology spread, logger, extra CA, optional daemon, etc.).

```bash
helm upgrade chaos-mesh ./helm/chaos-mesh \
  --namespace chaos-mesh \
  --values your-values.yaml
```

## CRDs

Upstream v2.8 may require applying updated CRDs before or with the controller upgrade. Apply CRDs from this tree when adding overlay-only kinds:

- `certificatechaos`
- `ciliumchaos`
- `cloudstackhostchaos`
- `cloudstackvmchaos`

Plus generic overlay kinds from earlier phases (`k8schaos`, `resourcescalechaos`, `rollingrestartchaos`, `podpvcchaos`, `nodeselectorchaos`).

```bash
kubectl apply -f helm/chaos-mesh/crds/
```

Re-run if the controller logs webhook or schema errors after upgrade.

## What is not in this overlay

- Form3 CI rework from legacy `master` (upstream workflows + `f3_upload_image.yml` only).
- Unmerged remote types (`AWSAZChaos`, `GCPAZChaos`, `GKENodePoolChaos`) unless added later.
- Task `PodSpec` extensions deferred from Form3 `master`.

## Rollback

Keep the previous Form3 chart revision and image tags. CRD downgrades are generally unsafe; prefer restoring the prior controller/chart while leaving newer CRD schema in place only if workloads still validate.
