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

  const cmd = new Deno.Command("docker", {
    stderr: "inherit",
    stdout: "inherit",
    stdin: "piped",
    args: ["login", "docker.io", "-u", await dockerUsernameSecret.plaintext(), "--password-stdin"],
  })

  const p = cmd.spawn();
  const writer = p.stdin.getWriter();

  await writer.write(new TextEncoder().encode(await dockerTokenSecret.plaintext()));
  await writer.close()

  const output = await p.output()

  if (output.code !== 0) {
    throw new Error("Failed to login into docker");
  }

  const container = await build({ client });

  await deploy({
    container,
    repository: "docker.io/michaelmass/hellomicro",
    tags: ["latest"],
  });

  await logout();
});
