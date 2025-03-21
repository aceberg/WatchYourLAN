import { useParams } from "@solidjs/router";
import { onMount } from "solid-js";

import { apiGetHost } from "../functions/api";
import { setCurrentHost } from "../functions/exports";

import HostCard from "../components/HostPage/HostCard";
import Ping from "../components/HostPage/Ping";

function HostPage() {

  onMount(async () => {
    const params = useParams();
    const host = await apiGetHost(params.id);

    setCurrentHost(host);
  });

  return (
    <div class="row">
      <div class="col-md">
        <HostCard></HostCard>
      </div>
      <div class="col-md">
        <Ping></Ping>
      </div>
    </div>
  )
}

export default HostPage