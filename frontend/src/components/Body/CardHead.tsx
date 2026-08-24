import { Show } from "solid-js";
import { editNames, selectedIDs } from "../../functions/exports";
import Filter from "../Filter";
import Search from "../Search";
import { apiDelHost } from "../../functions/api";

function CardHead() {

  const handleDel = async () => {
    const ids = selectedIDs();
    
    for (let id of ids) {
      await apiDelHost(id);
    }
    
    window.location.href = '/';
  };

  return (
    <div class="device-toolbar">
      <div class="device-toolbar-filters">
        <Filter></Filter>
      </div>
      <div class="device-toolbar-actions">
        <Search></Search>
        <Show when={editNames() && selectedIDs().length > 0}>
          <button type="button" onClick={handleDel} title="Delete selected hosts" class="btn btn-sm wyl-button device-delete-button">
            <i class="bi bi-trash3-fill" aria-hidden="true"></i>
            <span>DELETE ({selectedIDs().length})</span>
          </button>
        </Show>
      </div>
    </div>
  )
}

export default CardHead
