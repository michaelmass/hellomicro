import { connect } from "https://esm.sh/@dagger.io/dagger@0.8.7";
import {
  build,
  deploy,
} from "https://raw.githubusercontent.com/michaelmass/pipelines/master/dagger/docker.ts";

await connect(async (client) => {
  const container = await build({ client });
  await deploy({
    client,
    container,
    repository: "michaelmass/hellomicro",
    tags: ["latest"],
    token: client.setSecret("a", "a"),
  });
});
