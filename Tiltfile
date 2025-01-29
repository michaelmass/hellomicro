local_resource(
  'generator',
  'buf generate',
  deps=['./proto'],
  labels=[
    'generator',
  ],
)

local_resource(
  'backend',
  serve_cmd='go run . serve',
  readiness_probe=probe(http_get=http_get_action(port=8080, path='/v1/ping'), period_secs=1, failure_threshold=10),
  deps=['./pkg', './main.go', './api'],
  links=[
    'http://127.0.0.1:8080/v1/ping',
  ],
  labels=[
    'service',
  ],
)

local_resource(
  'frontend',
  'cd ./client && pnpm install',
  serve_cmd='cd ./client && pnpm start',
  readiness_probe=probe(http_get=http_get_action(port=3000, path='/'), period_secs=1, failure_threshold=10),
  ignore=['./client/node_modules'],
  links=[
    'http://127.0.0.1:3000',
  ],
  labels=[
    'service',
  ],
)
