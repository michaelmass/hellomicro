import { connect } from "https://esm.sh/@dagger.io/dagger@0.9.3";
import {
  build,
  publish,
} from "https://raw.githubusercontent.com/michaelmass/pipelines/master/dagger/docker.ts";
import { getInfinsical } from "https://raw.githubusercontent.com/michaelmass/pipelines/master/dagger/infisical.ts";
import { context } from "npm:@actions/github";

await connect(async (client) => {
  const infisical = getInfinsical({ client });

  const dockerTokenSecret = await infisical.get({
    name: "TOKEN",
    secretPath: "docker",
    secretName: "dockerToken",
  });

  const dockerUsernameSecret = await infisical.get({
    name: "USERNAME",
    secretPath: "docker",
    secretName: "dockerUsername",
  });

  const username = await dockerUsernameSecret.plaintext();

  const [container, ...platformVariants] = await Promise.all([
    build({ client, platform: 'linux/amd64' }),
    build({ client, platform: 'linux/arm64' }),
  ]);

  await publish({
    container,
    password: dockerTokenSecret,
    username,
    repository: `${username}/hellomicro`,
    tags: getTags(context.ref),
    platformVariants,
  });
});

function removePrefix(value: string, prefix: string) {
  return value.startsWith(prefix) ? value.slice(prefix.length) : value;
}

function getTags(ref: string) {
  return ref.startsWith("refs/tags/")
    ? ["latest", removePrefix(removePrefix(ref, "refs/tags/"), "v")]
    : ["latest"];
}
