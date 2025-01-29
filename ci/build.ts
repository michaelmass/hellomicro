import { connect } from "https://esm.sh/@dagger.io/dagger@0.9.6";
import { build } from "https://raw.githubusercontent.com/michaelmass/pipelines/master/dagger/docker.ts";

await connect(async (client) => {
  await Promise.all([
    build({ client, platform: 'linux/amd64' }),
    build({ client, platform: 'linux/arm64' }),
  ]);
});
