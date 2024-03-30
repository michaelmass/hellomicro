import { connect } from "https://esm.sh/@dagger.io/dagger@0.9.3";
import {
  build,
  publish,
} from "https://raw.githubusercontent.com/michaelmass/pipelines/master/dagger/docker.ts";
import { getInfinsical } from "https://raw.githubusercontent.com/michaelmass/pipelines/master/dagger/infisical.ts";
import { context } from 'npm:@actions/github'

await connect(async (client) => {
  const infisical = getInfinsical({ client });

  console.log(`Event: ${context.eventName}`)
  console.log(`Ref: ${context.ref}`)

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
  const container = await build({ client });

  await publish({
    container,
    password: dockerTokenSecret,
    username,
    repository: `${username}/hellomicro`,
    tags: ["latest"],
  });
});
