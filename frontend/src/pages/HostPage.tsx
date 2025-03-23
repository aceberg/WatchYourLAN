import { useParams } from "@solidjs/router";
import { createSignal, onMount } from "solid-js";

import { apiGetHost } from "../functions/api";

import HostCard from "../components/HostPage/HostCard";
import Ping from "../components/HostPage/Ping";
import HistCard from "../components/HostPage/HistCard";
import { emptyHost, Host } from "../functions/exports";

function HostPage() {

  const [currentHost, setCurrentHost] = createSignal<Host>(emptyHost);

  onMount(async () => {
    const params = useParams();
    const host = await apiGetHost(params.id);

    setCurrentHost(host);
  });

  return (
    <>
    <div class="row">
      <div class="col-md">
        <HostCard host={currentHost()}></HostCard>
      </div>
      <div class="col-md">
        <Ping IP={currentHost().IP}></Ping>
      </div>
    </div>
    <div class="row mt-4">
      <div class="col-md">
        <HistCard mac={currentHost().Mac}></HistCard>
      </div>
    </div>
    </>
  )
}

export default HostPage