import { useLocation, useNavigate, useParams } from "@solidjs/router";
import { createEffect, createSignal, onCleanup, Show } from "solid-js";

import { apiGetHost } from "../functions/api";

import HostCard from "../components/HostPage/HostCard";
import Ping from "../components/HostPage/Ping";
import HostActivityCard from "../components/HostPage/HostActivityCard";
import HistCard from "../components/HostPage/HistCard";
import { emptyHost, emptyPageContext, Host, setPageContext } from "../functions/exports";

function HostPage() {

  const [currentHost, setCurrentHost] = createSignal<Host>(emptyHost);
  const [loadError, setLoadError] = createSignal("");
  const params = useParams();
  const location = useLocation();
  const navigate = useNavigate();
  const previousTitle = document.title;
  let requestId = 0;

  const isEditMode = () => new URLSearchParams(location.search).get("edit") === "1";
  const setEditMode = (editing: boolean) => {
    if (!params.id) {
      return;
    }

    navigate("/host/" + params.id + (editing ? "?edit=1" : ""));
  };

  createEffect(() => {
    const id = params.id;

    if (!id) {
      return;
    }

    const activeRequest = ++requestId;
    setLoadError("");
    setCurrentHost(emptyHost);
    setPageContext({ kind: "host", hostName: "" });
    document.title = "Host · LANventory";

    apiGetHost(id)
      .then((host) => {
        if (activeRequest !== requestId) {
          return;
        }

        setCurrentHost(host);
      })
      .catch(() => {
        if (activeRequest !== requestId) {
          return;
        }

        setPageContext({ kind: "host", hostName: "" });
        document.title = "Host · LANventory";
        setLoadError("Host details could not be loaded. The device may have been deleted or the backend may be unavailable.");
      });
  });

  onCleanup(() => {
    requestId++;
    setPageContext(emptyPageContext);
    document.title = previousTitle;
  });

  createEffect(() => {
    const host = currentHost();

    if (host.ID === 0) {
      return;
    }

    const hostName = host.Name.trim();
    setPageContext({ kind: "host", hostName });
    document.title = (hostName || "Host") + " · LANventory";
  });

  return (
    <div class="host-page">
    <Show
      when={!loadError()}
      fallback={
        <div class="data-load-warning" role="alert">
          <i class="bi bi-exclamation-triangle-fill" aria-hidden="true"></i>
          <span>{loadError()}</span>
        </div>
      }
    >
      <div class="row g-3 mx-0 host-page-row">
        <div class="col-md">
          <HostCard host={currentHost()} editMode={isEditMode()} onEditModeChange={setEditMode} onHostChange={setCurrentHost}></HostCard>
        </div>
        <div class="col-md">
          <Ping IP={currentHost().IP}></Ping>
        </div>
      </div>
      <div class="row g-3 mx-0 mt-1 host-page-row">
        <div class="col-md">
          <HostActivityCard host={currentHost()}></HostActivityCard>
        </div>
      </div>
      <div class="row g-3 mx-0 mt-1 host-page-row">
        <div class="col-md">
          <HistCard mac={currentHost().Mac}></HistCard>
        </div>
      </div>
    </Show>
    </div>
  )
}

export default HostPage
