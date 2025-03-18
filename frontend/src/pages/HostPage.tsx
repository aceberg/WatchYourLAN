import { useParams } from "@solidjs/router";
import { apiGetHost } from "../functions/api";

import HostPageHost from "../components/HostPageHost";
import { setCurrentHost } from "../functions/exports";
import { onMount } from "solid-js";
import HostPagePing from "../components/HostPagePing";

function HostPage() {

  onMount(async () => {
    const params = useParams();
    const host = await apiGetHost(params.id);

    setCurrentHost(host);
  });

  return (
    <div class="row">
      <div class="col-md">
        <HostPageHost></HostPageHost>
      </div>
      <div class="col-md">
        <HostPagePing></HostPagePing>
      </div>
    </div>
  )
}

export default HostPage