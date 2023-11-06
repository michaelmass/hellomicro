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

  const output = await new Deno.Command("docker", {
    stderr: "inherit",
    stdout: "inherit",
    args: ["login", "-u", await dockerUsernameSecret.plaintext(), "-p", await dockerTokenSecret.plaintext()],
  }).output();

  if (output.code !== 0) {
    throw new Error("Failed to login into docker");
  }

  const container = await build({ client });

  await deploy({
    container,
    repository: "michaelmass/hellomicro",
    tags: ["0.2.1"],
  });

  await logout();
});
