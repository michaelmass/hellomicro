import { connect } from "https://esm.sh/@dagger.io/dagger@0.8.7";
import {
  build,
  login,
  logout,
  deploy,
} from "https://raw.githubusercontent.com/michaelmass/pipelines/master/dagger/docker.ts";
import { getInfinsical } from "https://raw.githubusercontent.com/michaelmass/pipelines/master/dagger/infisical.ts";

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

  await login({
    username: dockerUsernameSecret,
    password: dockerTokenSecret,
  })

  const container = await build({ client });

  await deploy({
    container,
    repository: "docker.io/michaelmass/hellomicro",
    tags: ["latest"],
  });

  await logout();
});
