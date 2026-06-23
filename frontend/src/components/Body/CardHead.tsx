import { Show, createSignal } from "solid-js";
import { editNames, selectedIDs, setEditNames } from "../../functions/exports";
import Filter from "../Filter";
import Search from "../Search";
import { getHosts } from "../../functions/atstart";
import { apiDelHost, apiRescan } from "../../functions/api";

function CardHead() {
  const [rescanning, setRescanning] = createSignal(false);

  const handleEditNames = (toggle: boolean) => {
    if (!toggle) {
      getHosts();
    }
    setEditNames(toggle);
  };

  const handleDel = async () => {
    const ids = selectedIDs();
    
    for (let id of ids) {
      await apiDelHost(id);
    }
    
    window.location.href = '/';
  };

  const handleRescan = async () => {
    if (rescanning()) {
      return;
    }

    setRescanning(true);
    try {
      await apiRescan();
      await getHosts();
    } finally {
      setRescanning(false);
    }
  };

  return (
    <div class="row">
      <div class="col-md mt-1 mb-1">
        <div class="d-flex justify-left">
        <Filter></Filter>
        </div>
      </div>
      <div class="col-md mt-1 mb-1">
        <div class="d-flex justify-content-between">
        <div class="d-flex gap-2">
          <Search></Search>
          <button
            class="btn btn-outline-primary"
            title="Rescan now"
            onClick={handleRescan}
            disabled={rescanning()}
            style="width: 2.5rem;"
          >
            <i class={rescanning() ? "spinner-border spinner-border-sm" : "bi bi-arrow-clockwise"}></i>
          </button>
        </div>
        <Show
          when={editNames()}
          fallback={<button class="btn btn-outline-primary" title="Toggle edit" onClick={[handleEditNames, true]}>Edit</button>}
        >
          <button type="button" onClick={handleDel} title="Delete selected hosts" class="btn btn-outline-danger">Delete selected</button>
          <button class="btn btn-primary" title="Toggle edit" onClick={[handleEditNames, false]}>Edit</button>
        </Show>
        </div>
      </div>
    </div>
  )
}

export default CardHead
