import { connect } from "https://esm.sh/@dagger.io/dagger@0.9.3";
import {
  build,
} from "https://raw.githubusercontent.com/michaelmass/pipelines/master/dagger/docker.ts";

await connect(async (client) => {
  await build({ client });
});
