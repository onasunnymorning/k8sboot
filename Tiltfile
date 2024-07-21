allow_k8s_contexts('default')

# Build Image
docker_build('gprins/k8sboot', '.', dockerfile='cmd/api/Dockerfile')

# Read the Helm Chart and convert to k8s YAML
yaml = helm(
  'helmcharts/api',
  name='api',
  values=['./helmcharts/api/values.yaml']
)

# Apply the k8s YAML
k8s_yaml(yaml)